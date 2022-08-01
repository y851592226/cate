package httpreq

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Client struct {
	client             *http.Client
	requestOptionFuncs []RequestOptionFunc
	mws                []Middleware
	debug              bool
	debugBody          bool
}

var DefaultTransport http.RoundTripper = &http.Transport{
	Proxy: http.ProxyFromEnvironment,
	DialContext: (&net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
		DualStack: true,
	}).DialContext,
	MaxIdleConns:          100,
	IdleConnTimeout:       90 * time.Second,
	TLSHandshakeTimeout:   10 * time.Second,
	ExpectContinueTimeout: 1 * time.Second,
}

func NewClient(options ...ClientOptionFunc) (*Client, error) {
	client := &Client{
		client: &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyFromEnvironment,
				DialContext: (&net.Dialer{
					Timeout:   30 * time.Second,
					KeepAlive: 30 * time.Second,
					DualStack: true,
				}).DialContext,
				MaxIdleConns:          100,
				IdleConnTimeout:       90 * time.Second,
				TLSHandshakeTimeout:   10 * time.Second,
				ExpectContinueTimeout: 1 * time.Second,
			},
		},
	}
	for _, option := range options {
		if err := option(client); err != nil {
			return nil, err
		}
	}
	return client, nil
}

var DefaultClient, _ = NewClient()

func (c *Client) Do(req *http.Request, options ...RequestOptionFunc) (*Response, error) {
	var err error
	for _, option := range c.requestOptionFuncs {
		if req, err = option(req); err != nil {
			return nil, err
		}
	}
	for _, option := range options {
		if req, err = option(req); err != nil {
			return nil, err
		}
	}
	ctx := req.Context()
	timeout, ok := GetTimeout(ctx)
	if ok {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, timeout)
		defer cancel()
	}
	query, ok := GetQuery(ctx)
	if ok {
		req.URL.RawQuery = query.Encode()
	}
	req = req.WithContext(ctx)

	mws, ok := GetMiddlewares(ctx)
	if !ok {
		mws = []Middleware{}
	}
	debugBody, ok := GetDebug(ctx)
	if ok {
		mws = append(mws, DebugMW(debugBody))
	} else if c.debug {
		mws = append(mws, DebugMW(c.debugBody))
	}

	// add global middlewares
	tmp := append(c.mws, globalUserDefinedMWs...)
	mws = append(tmp, mws...)
	endpoint := Chain(mws)(c.do)

	retryTimes, _ := GetRetryTimes(ctx)
	// do request
	var resp *Response
	for i := 0; i <= retryTimes; i++ {
		err = setBody(ctx, req)
		if err != nil {
			continue
		}
		resp, err = endpoint(req)
		if err != nil {
			continue
		}
		return resp, err
	}
	return nil, err
}

func setBody(ctx context.Context, req *http.Request) error {
	var err error
	body, ok := GetBody(ctx)
	rc, isReadCloser := body.(io.ReadCloser)
	if !isReadCloser && body != nil {
		rc2, isReader := body.(io.Reader)
		if isReader {
			req.Body = ioutil.NopCloser(rc2)
		}
	} else {
		req.Body = rc
	}
	if ok {
		switch v := body.(type) {
		case url.Values:
			b := v.Encode()
			req.ContentLength = int64(len(b))
			req.GetBody = func() (io.ReadCloser, error) {
				return ioutil.NopCloser(strings.NewReader(b)), nil
			}
		case *bytes.Buffer:
			req.ContentLength = int64(v.Len())
			buf := v.Bytes()
			req.GetBody = func() (io.ReadCloser, error) {
				r := bytes.NewReader(buf)
				return ioutil.NopCloser(r), nil
			}
		case *bytes.Reader:
			req.ContentLength = int64(v.Len())
			snapshot := *v
			req.GetBody = func() (io.ReadCloser, error) {
				r := snapshot
				return ioutil.NopCloser(&r), nil
			}
		case *strings.Reader:
			req.ContentLength = int64(v.Len())
			snapshot := *v
			req.GetBody = func() (io.ReadCloser, error) {
				r := snapshot
				return ioutil.NopCloser(&r), nil
			}
		case string:
			req.ContentLength = int64(len(v))
			req.GetBody = func() (io.ReadCloser, error) {
				return ioutil.NopCloser(strings.NewReader(v)), nil
			}
		case []byte:
			req.ContentLength = int64(len(v))
			req.GetBody = func() (io.ReadCloser, error) {
				return ioutil.NopCloser(bytes.NewReader(v)), nil
			}
		default:
			var data []byte
			data, err = json.Marshal(body)
			if err != nil {
				return err
			}
			req.Body = ioutil.NopCloser(bytes.NewReader(data))
			req.ContentLength = int64(len(data))
			req.Header.Set("Content-Type", "application/json")
			req.GetBody = func() (io.ReadCloser, error) {
				return ioutil.NopCloser(bytes.NewReader(data)), nil
			}
		}
		if req.GetBody != nil && req.ContentLength == 0 {
			req.Body = http.NoBody
			req.GetBody = func() (io.ReadCloser, error) { return http.NoBody, nil }
		} else {
			req.Body, err = req.GetBody()
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (c *Client) do(req *http.Request) (*Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	return newResponse(req, resp)
}

func Get(url string, options ...RequestOptionFunc) (*Response, error) {
	return DefaultClient.Get(url, options...)
}

func Post(url string, options ...RequestOptionFunc) (*Response, error) {
	return DefaultClient.Post(url, options...)
}

func Head(url string, options ...RequestOptionFunc) (*Response, error) {
	return DefaultClient.Head(url, options...)
}

func Put(url string, options ...RequestOptionFunc) (*Response, error) {
	return DefaultClient.Put(url, options...)
}

func Delete(url string, options ...RequestOptionFunc) (*Response, error) {
	return DefaultClient.Delete(url, options...)
}

func (c *Client) Get(url string, options ...RequestOptionFunc) (*Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	return c.Do(req, options...)
}

func (c *Client) Post(url string, options ...RequestOptionFunc) (*Response, error) {
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}
	return c.Do(req, options...)
}

func (c *Client) Head(url string, options ...RequestOptionFunc) (*Response, error) {
	req, err := http.NewRequest("HEAD", url, nil)
	if err != nil {
		return nil, err
	}
	return c.Do(req, options...)
}

func (c *Client) Put(url string, options ...RequestOptionFunc) (*Response, error) {
	req, err := http.NewRequest("PUT", url, nil)
	if err != nil {
		return nil, err
	}
	return c.Do(req, options...)
}

func (c *Client) Delete(url string, options ...RequestOptionFunc) (*Response, error) {
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}
	return c.Do(req, options...)
}
