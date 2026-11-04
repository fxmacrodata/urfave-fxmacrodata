package fxmdcli

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const MaxBodyBytes = 4 << 20

type guardedTransport struct {
	base     http.RoundTripper
	key      string
	ctx      context.Context
	raw      []byte
	status   int
	complete bool
	calls    int
}

func newTransport(base http.RoundTripper, key string, ctx context.Context) *guardedTransport {
	if base == nil {
		base = &http.Transport{Proxy: nil, MaxIdleConns: 4, IdleConnTimeout: 30 * time.Second, TLSHandshakeTimeout: 10 * time.Second, ResponseHeaderTimeout: 30 * time.Second}
	}
	return &guardedTransport{base: base, key: key, ctx: ctx}
}
func (g *guardedTransport) RoundTrip(request *http.Request) (*http.Response, error) {
	g.calls++
	if g.calls != 1 {
		g.complete = false
		return nil, errors.New("one upstream request per command")
	}
	if request == nil || request.URL == nil || request.URL.Scheme != "https" || request.URL.Host != "api.fxmacrodata.com" || request.URL.User != nil || request.URL.Opaque != "" || request.URL.Fragment != "" || !strings.HasPrefix(request.URL.Path, "/v1/") {
		return nil, errors.New("request origin or path is not permitted")
	}
	if request.Method != "GET" && request.Method != "POST" {
		return nil, errors.New("only GET and POST are permitted")
	}
	for _, part := range strings.Split(request.URL.Path, "/") {
		if part == "." || part == ".." {
			return nil, errors.New("path traversal is forbidden")
		}
	}
	query, err := url.ParseQuery(request.URL.RawQuery)
	if err != nil {
		return nil, errors.New("invalid query encoding")
	}
	for name := range query {
		normalized := strings.NewReplacer("_", "", "-", "").Replace(strings.ToLower(name))
		switch normalized {
		case "apikey", "fxmdapikey", "fxmacrodataapikey", "accesstoken", "authorization":
			return nil, errors.New("credentials must not appear in query parameters")
		}
	}
	copy := request.Clone(g.ctx)
	copy.Host = ""
	copy.Header = make(http.Header)
	copy.Trailer = nil
	copy.TransferEncoding = nil
	copy.ContentLength = 0
	copy.GetBody = nil
	copy.Header.Set("Accept", "application/json")
	copy.Header.Set("User-Agent", "FXMacroData-Urfave/0.1.0")
	if contentType := request.Header.Get("Content-Type"); contentType != "" {
		copy.Header.Set("Content-Type", contentType)
	}
	if g.key != "" {
		copy.Header.Set("X-API-Key", g.key)
	}
	if request.Body != nil {
		defer request.Body.Close()
		raw, err := io.ReadAll(io.LimitReader(request.Body, MaxBodyBytes+1))
		if err != nil || len(raw) > MaxBodyBytes {
			return nil, errors.New("request body exceeds bounds")
		}
		copy.Body = io.NopCloser(bytes.NewReader(raw))
		copy.ContentLength = int64(len(raw))
		copy.GetBody = func() (io.ReadCloser, error) { return io.NopCloser(bytes.NewReader(raw)), nil }
	}
	response, err := g.base.RoundTrip(copy)
	if err != nil {
		return nil, errors.New("request transport failed")
	}
	if response == nil || response.Body == nil {
		return nil, errors.New("empty response")
	}
	defer response.Body.Close()
	g.status = response.StatusCode
	raw, err := io.ReadAll(io.LimitReader(response.Body, MaxBodyBytes+1))
	if err != nil || len(raw) > MaxBodyBytes {
		return nil, errors.New("response exceeds bounds")
	}
	g.raw = raw
	g.complete = true
	response.Body = io.NopCloser(bytes.NewReader(raw))
	response.ContentLength = int64(len(raw))
	return response, nil
}
