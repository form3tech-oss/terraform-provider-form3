package httputil

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"regexp"
)

// SecureDumpRequest does a security aware dump of a given HTTP request.
func SecureDumpRequest(req *http.Request) ([]byte, error) {
	var err error
	bodyCopy := req.Body
	bodyCopy, req.Body, err = drainBody(req.Body)
	if err != nil {
		return nil, fmt.Errorf("copying body failed: %w", err)
	}
	defer func() {
		req.Body = bodyCopy
	}()

	reqClone := req.Clone(req.Context())

	reqClone.Header.Del("Authorization")

	dump, err := httputil.DumpRequestOut(reqClone, true)
	return dump, err
}

// SecureDumpResponse does a security aware dump of a given HTTP response.
func SecureDumpResponse(res *http.Response) ([]byte, error) {
	data, err := httputil.DumpResponse(res, true)
	if err != nil {
		return nil, err
	}

	text := string(data)

	// TODO: cache regexps, make it configurable, support
	// variable whitespace between key and value, case insensitive
	// fields
	re := regexp.MustCompile(`"access_token":".*?"`)
	text = re.ReplaceAllString(text, `"access_token": "******"`)
	re = regexp.MustCompile(`"refresh_token":".*?"`)
	text = re.ReplaceAllString(text, `"refresh_token": "******"`)
	return []byte(text), nil
}

// drainBody copied from net/http/httputil/dump.go
func drainBody(b io.ReadCloser) (r1, r2 io.ReadCloser, err error) {
	if b == nil || b == http.NoBody {
		// No copying needed. Preserve the magic sentinel meaning of NoBody.
		return http.NoBody, http.NoBody, nil
	}
	var buf bytes.Buffer
	if _, err = buf.ReadFrom(b); err != nil {
		return nil, b, err
	}
	if err = b.Close(); err != nil {
		return nil, b, err
	}
	return ioutil.NopCloser(&buf), ioutil.NopCloser(bytes.NewReader(buf.Bytes())), nil
}
