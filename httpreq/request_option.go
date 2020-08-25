package httpreq

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/y851592226/cate/convert"
	"github.com/y851592226/cate/freejson"
)

const (
	KeyTimeout       = "$timeout"
	KeyMiddlewares   = "$middlewares"
	KeyDebug         = "$debug"
	KeyRetryTimes    = "$retryTimes"
	KeyQuery         = "$query"
	KeyForm          = "$form"
	KeyBody          = "$body"
	KeyAuthorization = "Authorization"
)

type RequestOptionFunc func(*http.Request) (*http.Request, error)

func AddRequestCookie(name, value string) RequestOptionFunc {
	return func(req *http.Request) (*http.Request, error) {
		cookie := &http.Cookie{
			Name:  name,
			Value: url.QueryEscape(value),
		}
		req.AddCookie(cookie)
		return req, nil
	}
}

func SetRequestBody(body interface{}) RequestOptionFunc {
	return SetRequestKV(KeyBody, body)
}

func SetRequestQuery(i interface{}) RequestOptionFunc {
	values := url.Values{}
	var err error
	switch v := i.(type) {
	case string:
		values, err = url.ParseQuery(v)
	case url.Values:
		values = v
	default:
		var o freejson.Object
		o, err = freejson.ToObject(i)
		if err != nil {
			break
		}
		err = o.Each(func(_ freejson.Object, key string, value interface{}) error {
			values.Add(key, convert.AsString(value))
			return nil
		})
	}
	if err != nil {
		return AddRequestError(err)
	}
	return SetRequestQueryValues(values)
}

func AddRequestError(err error) RequestOptionFunc {
	return func(req *http.Request) (*http.Request, error) {
		return nil, err
	}
}

func AddRequestQueryValue(key, value string) RequestOptionFunc {
	return func(req *http.Request) (*http.Request, error) {
		ctx := req.Context()
		values, ok := ctx.Value(KeyQuery).(url.Values)
		if !ok {
			values = url.Values{}
		}
		values.Add(key, value)
		return SetRequestQueryValues(values)(req)
	}
}

func SetRequestQueryValues(values url.Values) RequestOptionFunc {
	return SetRequestKV(KeyQuery, values)
}

func AddRequestFormValue(key, value string) RequestOptionFunc {
	return func(req *http.Request) (*http.Request, error) {
		ctx := req.Context()
		values, ok := ctx.Value(KeyForm).(url.Values)
		if !ok {
			values = url.Values{}
		}
		values.Add(key, value)
		return SetRequestFormValues(values)(req)
	}
}

func SetRequestPostForm(i interface{}) RequestOptionFunc {
	values := url.Values{}
	var err error
	switch v := i.(type) {
	case string:
		values, err = url.ParseQuery(v)
	case url.Values:
		values = v
	default:
		var o freejson.Object
		o, err = freejson.ToObject(i)
		if err != nil {
			break
		}
		err = o.Each(func(_ freejson.Object, key string, value interface{}) error {
			values.Add(key, convert.AsString(value))
			return nil
		})
	}
	if err != nil {
		return AddRequestError(err)
	}
	return SetRequestFormValues(values)
}

func SetRequestFormValues(values url.Values) RequestOptionFunc {
	return func(req *http.Request) (*http.Request, error) {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		return SetRequestKV(KeyForm, values)(req)
	}
}

func SetRequestRetryTimes(retryTimes int) RequestOptionFunc {
	return SetRequestKV(KeyRetryTimes, retryTimes)
}

func SetRequestDebug(debug, debugBody bool) RequestOptionFunc {
	if debug {
		return SetRequestKV(KeyDebug, debugBody)
	}
	return SetRequestKV(KeyDebug, nil)
}

func AddRequestMiddlewares(mws ...Middleware) RequestOptionFunc {
	return func(req *http.Request) (*http.Request, error) {
		ctx := req.Context()
		requestMiddlewares, ok := ctx.Value(KeyMiddlewares).([]Middleware)
		if !ok {
			requestMiddlewares = []Middleware{}
		}
		requestMiddlewares = append(requestMiddlewares, mws...)
		return SetRequestKV(KeyMiddlewares, requestMiddlewares)(req)
	}
}

func SetRequestBasicAuth(username, password string) RequestOptionFunc {
	return SetRequestHeader(KeyAuthorization, BasicAuthorization(username, password))
}

func AddRequestHeader(key, value string) RequestOptionFunc {
	return func(req *http.Request) (*http.Request, error) {
		req.Header.Add(key, value)
		return req, nil
	}
}

func SetRequestContentType(contentType string) RequestOptionFunc {
	return SetRequestHeader("Content-Type", contentType)
}

func SetRequestHeader(key, value string) RequestOptionFunc {
	return func(req *http.Request) (*http.Request, error) {
		req.Header.Set(key, value)
		return req, nil
	}
}

func SetRequestKV(key, value interface{}) RequestOptionFunc {
	return func(req *http.Request) (*http.Request, error) {
		ctx := req.Context()
		ctx = context.WithValue(ctx, key, value)
		req = req.WithContext(ctx)
		return req, nil
	}
}

func SetRequestTimeout(timeout time.Duration) RequestOptionFunc {
	return SetRequestKV(KeyTimeout, timeout)
}

func SetRequestContext(ctx context.Context) RequestOptionFunc {
	return func(req *http.Request) (*http.Request, error) {
		req = req.WithContext(ctx)
		return req, nil
	}
}
