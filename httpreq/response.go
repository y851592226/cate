package httpreq

import (
	"io/ioutil"
	"net/http"

	"github.com/y851592226/cate/freejson"
	"github.com/y851592226/cate/httpreq/binding"
)

type Response struct {
	RawRequest  *http.Request
	RawResponse *http.Response
	body        []byte
}

func newResponse(req *http.Request, resp *http.Response) (r *Response, err error) {
	r = &Response{
		RawRequest:  req,
		RawResponse: resp,
	}
	r.body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return r, nil
}

func (r *Response) Body() []byte {
	return r.body
}

func (r *Response) Status() string {
	return r.RawResponse.Status
}

func (r *Response) StatusCode() int {
	return r.RawResponse.StatusCode
}

func (r *Response) String() string {
	return string(r.body)
}

func (r *Response) Byte() []byte {
	return r.body
}

func (r *Response) ResponseHeader() http.Header {
	return r.RawResponse.Header
}

func (r *Response) Header(key string) string {
	return r.RawResponse.Header.Get(key)
}

func (r *Response) Cookies() []*http.Cookie {
	return r.RawResponse.Cookies()
}

func (r *Response) Bind(v interface{}) error {
	return binding.Default(r.Header("Content-Type")).BindBody(r.body, v)
}

func (r *Response) BindJSON(v interface{}) error {
	return binding.JSON.BindBody(r.body, v)
}

func (r *Response) BindXML(v interface{}) error {
	return binding.XML.BindBody(r.body, v)
}

func (r *Response) Object() freejson.Object {
	return freejson.AsObject(r.body)
}

func (r *Response) Array() freejson.Array {
	return freejson.AsArray(r.body)
}

func (r *Response) ToObject() (freejson.Object, error) {
	return freejson.ToObject(r.body)
}

func (r *Response) ToArray() (freejson.Array, error) {
	return freejson.ToArray(r.body)
}
