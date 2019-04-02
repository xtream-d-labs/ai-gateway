package registry

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type httpStatusError struct {
	Response *http.Response
	Body     []byte // Copied from `Response.Body` to avoid problems with unclosed bodies later. Nobody calls `err.Response.Body.Close()`, ever.
}

func (err *httpStatusError) Error() string {
	return fmt.Sprintf("http: non-successful response (status=%v body=%q)", err.Response.StatusCode, err.Body)
}

var _ error = &httpStatusError{}

// ErrorTransport defines the data structure for returning errors from the round tripper.
type ErrorTransport struct {
	Transport http.RoundTripper
}

// RoundTrip defines the round tripper for the error transport.
func (t *ErrorTransport) RoundTrip(request *http.Request) (*http.Response, error) {
	resp, e1 := t.Transport.RoundTrip(request)
	if e1 != nil {
		return resp, e1
	}
	if resp.StatusCode >= 500 || resp.StatusCode == http.StatusUnauthorized {
		defer resp.Body.Close()
		body, e2 := ioutil.ReadAll(resp.Body)
		if e2 != nil {
			return nil, fmt.Errorf("http: failed to read response body (status=%v, err=%q)", resp.StatusCode, e2)
		}
		return nil, &httpStatusError{
			Response: resp,
			Body:     body,
		}
	}
	return resp, nil
}
