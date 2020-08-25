package httpreq

import (
	"fmt"
	"net/http/httputil"
	"time"
)

type EndPoint func(*Request) (*Response, error)

type Middleware func(EndPoint) EndPoint

var globalUserDefinedMWs = []Middleware{}

func AddGlobalMiddlewares(mws ...Middleware) {
	globalUserDefinedMWs = append(globalUserDefinedMWs, mws...)
}

func Chain(mws []Middleware) Middleware {
	if len(mws) == 0 {
		return EmptyMiddleware
	}
	return func(next EndPoint) EndPoint {
		return mws[0](Chain(mws[1:])(next))
	}
}

func EmptyMiddleware(next EndPoint) EndPoint {
	return next
}

func DebugMW(body bool) Middleware {
	return func(next EndPoint) EndPoint {
		return func(req *Request) (*Response, error) {
			begin := time.Now()
			dump, err1 := httputil.DumpRequest(req, body)
			if err1 != nil {
				fmt.Println("Error:", err1)
			} else {
				fmt.Printf("[http] HTTP Request: \n%s\n", string(dump))
			}

			resp, err := next(req)
			if err != nil {
				return resp, err
			}
			dump, err2 := httputil.DumpResponse(resp.RawResponse, false)
			if err2 != nil {
				fmt.Println("Error:", err2)
			} else {
				if body {
					fmt.Printf("[http] cost:%s \nHTTP Response: \n%s%s\n",
						time.Since(begin), string(dump), resp.String())
				} else {
					fmt.Printf("[http] cost:%s \nHTTP Response: \n%s\n",
						time.Since(begin), string(dump))
				}

			}
			return resp, err
		}
	}
}

func ExpectStatusCodeMW(code int) Middleware {
	return func(next EndPoint) EndPoint {
		return func(req *Request) (*Response, error) {
			resp, err := next(req)
			if err != nil {
				return nil, err
			}
			if resp.StatusCode() != code {
				return nil, fmt.Errorf("unexpect StatusCode:%s\n    Body:%s", resp.Status(), resp.Body())
			}
			return resp, nil
		}
	}
}
