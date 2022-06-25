package utils

import (
	"crypto/tls"
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
	"time"
)

const (
	UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.5005.124 Safari/537.36 Edg/102.0.1245.44"

	HttpResultFormat = "StatusCode: [%d]; Body: [%s]"

	ErrorInvalidParameter = "Invalid parameter"

	URLSeparator = "/"
)

var (
	DefaultTransport *http.Transport
)

func init() {
	DefaultTransport = &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   5 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		MaxIdleConns:          20,
		IdleConnTimeout:       15 * time.Second,
		DisableCompression:    false,
		DisableKeepAlives:     false,
		ResponseHeaderTimeout: 10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		TLSHandshakeTimeout:   5 * time.Second,
		TLSClientConfig:       &tls.Config{InsecureSkipVerify: true},
	}
}

type HttpClientOpts struct {
	DisableCompression bool
	DisableKeepAlives  bool
	Headers            map[string]interface{}
	Cookies            []*http.Cookie
}

func NewHttpClient(method, url string, body io.Reader, opts ...*HttpClientOpts) (resp *http.Response, err error) {

	method = strings.TrimSpace(method)
	url = strings.TrimSpace(url)

	if method == "" || url == "" {
		err = fmt.Errorf("%s", ErrorInvalidParameter)
		return
	}

	method = strings.ToUpper(method)

	if len(opts) > 0 {
		DefaultTransport.DisableCompression = opts[0].DisableCompression
		DefaultTransport.DisableKeepAlives = opts[0].DisableKeepAlives
	}

	var httpClient = &http.Client{
		Transport: DefaultTransport,
	}

	var req *http.Request

	req, err = http.NewRequest(method, url, body)
	if err != nil {
		return
	}

	req.Header.Set("User-Agent", UserAgent)

	if len(opts) > 0 {
		if len(opts[0].Headers) > 0 {
			for k, v := range opts[0].Headers {
				req.Header.Set(k, ToStr(v))
			}
		}

		if len(opts[0].Cookies) > 0 {
			for _, cookie := range opts[0].Cookies {
				req.AddCookie(cookie)
			}
		}
	}

	return httpClient.Do(req)
}
