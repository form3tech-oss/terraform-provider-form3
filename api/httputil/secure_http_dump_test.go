package httputil_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/form3tech-oss/terraform-provider-form3/api"
	"github.com/form3tech-oss/terraform-provider-form3/api/httputil"
)

func TestSecureDumpRequest(t *testing.T) {
	cases := []struct {
		it                string
		req               *http.Request
		checkReqAfterDump func(*http.Request) error
		expectedInDump    []string
		unexpectedInDump  []string
	}{
		{
			it: "removes Authorization header from request",
			req: func() *http.Request {
				req, err := http.NewRequest(http.MethodPost, "", bytes.NewBuffer([]byte("the body")))
				if err != nil {
					t.Fatalf("create request failed: %v", err)
				}
				req.Header.Set("Authorization", "secret value")
				req.Header.Set("X-Special-Header", "special header")
				return req
			}(),
			checkReqAfterDump: func(req *http.Request) error {
				if req.Header.Get("Authorization") != "secret value" {
					return fmt.Errorf("expected header Authorization to have the value %q got %q", "secret value", req.Header.Get("Authorization"))
				}
				return nil
			},
			expectedInDump:   []string{"special header", "the body"},
			unexpectedInDump: []string{"secret value"},
		},

		{
			it: "works with request that does not have Authorization header",
			req: func() *http.Request {
				req, err := http.NewRequest(http.MethodPost, "", bytes.NewBuffer([]byte("the body")))
				if err != nil {
					t.Fatalf("create request failed: %v", err)
				}
				req.Header.Set("X-Special-Header", "special header")
				return req
			}(),
			expectedInDump: []string{"special header", "the body"},
		},
	}

	for _, tc := range cases {
		t.Run(tc.it, func(t *testing.T) {
			reqMadeCnt := 0
			mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				reqMadeCnt++
				defer req.Body.Close()
				w.WriteHeader(http.StatusOK)

				body, err := ioutil.ReadAll(req.Body)
				if err != nil {
					t.Fatalf("unexpected error %v", err)
				}

				if string(body) != "the body" {
					t.Errorf("expected %q got %q", "the body", string(body))
				}
			}))
			defer mockServer.Close()

			url, err := url.Parse(mockServer.URL)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			tc.req.URL = url

			dump, err := httputil.SecureDumpRequest(tc.req)
			if err != nil {
				t.Fatalf("unexpected error %v", err)
			}

			if tc.checkReqAfterDump != nil {
				err = tc.checkReqAfterDump(tc.req)
				if err != nil {
					t.Errorf("check of orginal request failed: %v", err)
				}
			}

			dumpString := string(dump)
			for _, v := range tc.expectedInDump {
				if !strings.Contains(dumpString, v) {
					t.Errorf("expected %q in dump:\n%s\n", v, dumpString)
				}
			}
			for _, uv := range tc.unexpectedInDump {
				if strings.Contains(dumpString, uv) {
					t.Errorf("unexpected %q in dump:\n%s\n", uv, dumpString)
				}
			}

			rewindableBody, err := api.NewReaderSeekerCloser(tc.req)
			if err != nil {
				t.Fatalf("unexpected error %v", err)
			}
			tc.req.Body = rewindableBody

			tryCnt := 3
			for i := 0; i < tryCnt; i++ {
				if tc.req.Body == nil {
					continue
				}

				body, ok := tc.req.Body.(api.ReadSeekerCloser)
				if !ok {
					t.Fatal("request body not satisfying the ReadSeekerCloser")
				}

				if _, err = body.Seek(0, 0); err != nil {
					t.Errorf("failed to seek request body: %s", err)
				}

				res, err := http.DefaultClient.Do(tc.req)
				if err != nil {
					t.Fatalf("unexpected error %v", err)
				}
				if res.StatusCode != http.StatusOK {
					t.Errorf("expected http status code %d got %d", http.StatusOK, res.StatusCode)
				}
			}

			if reqMadeCnt != tryCnt {
				t.Errorf("expected %d calls got %d", tryCnt, reqMadeCnt)
			}
		})
	}
}

func TestSecureDumpResponse(t *testing.T) {
	cases := []struct {
		it                 string
		req                *http.Request
		checkRespAfterDump func(*http.Response) error
		expectedInDump     []string
		unexpectedInDump   []string
	}{

		{
			it: "removes 'access_token' and 'refresh_token' fields from response",
			req: func() *http.Request {
				body := []byte(`{"access_token":"secret value","refresh_token":"secret value","other":"ok"}`)
				req, err := http.NewRequest(http.MethodPost, "", bytes.NewBuffer(body))
				if err != nil {
					t.Fatalf("create request failed: %v", err)
				}

				return req
			}(),
			checkRespAfterDump: func(resp *http.Response) error {
				return nil
			},
			expectedInDump:   []string{"ok"},
			unexpectedInDump: []string{"secret value"},
		},
	}

	for _, tc := range cases {
		t.Run(tc.it, func(t *testing.T) {
			mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				defer req.Body.Close()
				w.WriteHeader(http.StatusOK)
				req.Write(w)
			}))
			defer mockServer.Close()
			url, err := url.Parse(mockServer.URL)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			tc.req.URL = url

			resp, err := http.DefaultClient.Do(tc.req)
			if err != nil {
				t.Fatalf("unexpected error %v", err)
			}

			dump, err := httputil.SecureDumpResponse(resp)
			if err != nil {
				t.Fatalf("unexpected error %v", err)
			}

			if tc.checkRespAfterDump != nil {
				err = tc.checkRespAfterDump(resp)
				if err != nil {
					t.Errorf("check of orginal request failed: %v", err)
				}
			}

			dumpString := string(dump)
			for _, v := range tc.expectedInDump {
				if !strings.Contains(dumpString, v) {
					t.Errorf("expected %q in dump:\n%s\n", v, dumpString)
				}
			}
			for _, uv := range tc.unexpectedInDump {
				if strings.Contains(dumpString, uv) {
					t.Errorf("unexpected %q in dump:\n%s\n", uv, dumpString)
				}
			}

		})
	}
}
