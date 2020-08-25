package httpreq

import (
	"context"
	"net/http"
	"net/url"
	"time"
)

type Request = http.Request

var NewRequest = http.NewRequest

func GetTimeout(ctx context.Context) (time.Duration, bool) {
	value, ok := ctx.Value(KeyTimeout).(time.Duration)
	return value, ok
}

func GetQuery(ctx context.Context) (url.Values, bool) {
	value, ok := ctx.Value(KeyQuery).(url.Values)
	return value, ok
}

func GetBody(ctx context.Context) (interface{}, bool) {
	form, ok := ctx.Value(KeyForm).(url.Values)
	if ok {
		return form, ok
	}
	body := ctx.Value(KeyBody)
	return body, body != nil
}

func GetMiddlewares(ctx context.Context) ([]Middleware, bool) {
	value, ok := ctx.Value(KeyMiddlewares).([]Middleware)
	return value, ok
}

func GetDebug(ctx context.Context) (bool, bool) {
	value, ok := ctx.Value(KeyDebug).(bool)
	return value, ok

}

func GetRetryTimes(ctx context.Context) (int, bool) {
	value, ok := ctx.Value(KeyRetryTimes).(int)
	return value, ok
}
