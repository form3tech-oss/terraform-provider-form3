package httputil

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"regexp"
	"time"

	"github.com/goware/prefixer"
)

const secureMask = "******"

var (
	tokenRe = regexp.MustCompile(`"((?i)access_token|(?i)refresh_token)":\s*?".*?"`)
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

	if v := reqClone.Header.Get("Authorization"); len(v) > 0 {
		reqClone.Header.Set("Authorization", secureMask)
	}

	dump, err := httputil.DumpRequestOut(reqClone, true)
	if err != nil {
		return dump, err
	}

	dump = withEmptyLine(dump)

	prefixReader := prefixer.New(bytes.NewBuffer(dump), prefix("[REQ]"))
	return ioutil.ReadAll(prefixReader)
}

// SecureDumpResponse does a security aware dump of a given HTTP response.
func SecureDumpResponse(res *http.Response) ([]byte, error) {
	data, err := httputil.DumpResponse(res, true)
	if err != nil {
		return nil, err
	}

	text := string(data)
	text = tokenRe.ReplaceAllString(text, fmt.Sprintf(`"$1": "%s"`, secureMask))

	dump := withEmptyLine([]byte(text))

	prefixReader := prefixer.New(bytes.NewBuffer(dump), prefix("[RES]"))
	return ioutil.ReadAll(prefixReader)
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

func prefix(msg string) string {
	return fmt.Sprintf("%s [DEBUG] %s ", time.Now().Format("2006/01/02 15:04:05"), msg)
}

func withEmptyLine(data []byte) []byte {
	return append([]byte("\n"), data...)
}
