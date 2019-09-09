package http

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/go-openapi/swag"
	"golang.org/x/net/context/ctxhttp"
)

// HTTPSend a http request
func HTTPSend(ctx context.Context, method, path string, query url.Values, body io.Reader, headers http.Header, contentLength int64) ([]byte, error) {
	cli := &http.Client{}
	req, err := build(cli, method, apipath(path, query), body, headers, contentLength)
	if err != nil {
		return nil, err
	}
	resp, err := request(ctx, cli, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func apipath(path string, query url.Values) string {
	u, _ := url.Parse(path)
	if len(query) > 0 {
		u.RawQuery = query.Encode()
	}
	return u.String()
}

func build(cli *http.Client, method, path string, body io.Reader, headers http.Header, contentLength int64) (*http.Request, error) {
	expectedPayload := (method == http.MethodPost || method == http.MethodPut)
	if expectedPayload && body == nil {
		body = bytes.NewReader([]byte{})
	}
	req, err := http.NewRequest(method, path, body)
	if err != nil {
		return nil, err
	}
	if headers != nil {
		for key := range headers {
			req.Header.Set(key, headers.Get(key))
		}
	}
	if expectedPayload && swag.IsZero(req.Header.Get("Content-Type")) {
		req.Header.Set("Content-Type", "text/plain")
	}
	if contentLength > 0 {
		req.ContentLength = contentLength
	}
	return req, nil
}

func request(ctx context.Context, cli *http.Client, req *http.Request) (*http.Response, error) {
	resp, err := ctxhttp.Do(ctx, cli, req)
	if err != nil {
		switch err {
		case context.Canceled, context.DeadlineExceeded:
			return nil, err
		}
		if nErr, ok := err.(*url.Error); ok {
			if nErr, ok := nErr.Err.(*net.OpError); ok {
				if os.IsPermission(nErr.Err) {
					return nil, errors.New("Permission denied while trying to connect to. " + err.Error())
				}
			}
		}
		if err, ok := err.(net.Error); ok {
			if err.Timeout() {
				return nil, errors.New("Connection failed. " + err.Error())
			}
			if !err.Temporary() {
				if strings.Contains(err.Error(), "connection refused") || strings.Contains(err.Error(), "dial unix") {
					return nil, errors.New("Connection failed. " + err.Error())
				}
			}
		}
		return nil, errors.New("error during connect. " + err.Error())
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 400 {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		if len(body) == 0 {
			return nil, fmt.Errorf("Error: request returned %s for API route and version %s, check if the server supports the requested API version", http.StatusText(resp.StatusCode), req.URL)
		}
		return nil, fmt.Errorf("Error response from server: %s", string(body))
	}
	return resp, nil
}
