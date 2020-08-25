# httpreq

> A simple, easy-to-use HTTP client for Go

### Simple Example

GET request

```go
func ExampleGet() {
	resp, err := httpreq.Get("http://httpbin.org/get",
		httpreq.AddRequestQueryValue("key1", "value1"),
		httpreq.AddRequestQueryValue("key2", "value2"),
		httpreq.SetRequestHeader("User-Agent", "httpreq/1.0.0"),
		httpreq.SetRequestDebug(true, true),
		httpreq.SetRequestRetryTimes(3))
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.StatusCode())
	fmt.Println(resp.String())
}
```

POST request

```go
type Request struct {
	Content string
}

type Result struct {
	Data string `json:"data"`
}

func ExamplePost() {
	resp, err := httpreq.Post("http://httpbin.org/post",
		httpreq.SetRequestBody(Request{"post example"}),
		httpreq.SetRequestDebug(true, true),
		httpreq.SetRequestRetryTimes(3))
	if err != nil {
		panic(err)
	}
	var result Result
	err = resp.Bind(&result)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
```

### Client

Client is a wrapper of http.Client,  you can init a client with some client option functions

```go
type Client struct {
	client    *http.Client
	mws       []Middleware
	debug     bool
	debugBody bool
}
```

Use client option function

```
client,err := httpreq.NewClient(httpreq.SetClientTimeout(time.Second))
```

All client option function

- SetClientTimeout
- SetClientDebug
- SetClientTransport
- SetClientCheckRedirect
- SetClientCookieJar
- SetClientProxyURL
- SetClientProxy
- SetClientDial
- SetClientDialContext
- SetClientDialTLS
- AddClientMiddlewares

### Request

Requst is a alias of http.Request

```go
type Request = http.Request
```

Use request option function

```go
func ExampleGet2() {
	resp, err := httpreq.Get("http://httpbin.org/get",
		httpreq.SetRequestBasicAuth("username", "password"),
		httpreq.AddRequestCookie("sessionID", "12345"),
		httpreq.SetRequestRetryTimes(3))
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.String())
}
```

All request option func

- AddRequestCookie
- SetRequestBody
- AddRequestQueryValue
- SetRequestQueryValues
- AddRequestFormValue
- SetRequestFormValues
- SetRequestRetryTimes
- SetRequestDebug
- AddRequestMiddlewares
- SetRequestBasicAuth
- AddRequestHeader
- SetRequestContentType
- SetRequestHeader
- SetRequestKV
- SetRequestTimeout
- SetRequestContext

### Response

Response is a wrapper of http.Response, it is used to get response info from http.Response.

```go
type Response struct {
	RawRequest  *http.Request
	RawResponse *http.Response
	body        []byte
}
```

Bind response

```go
func ExampleBindResponse() {
	type Request struct {
		PageNo   int
		PageSize int
	}
	type Result struct {
		Json Request `json:"json"`
	}

	resp, err := httpreq.Post("http://httpbin.org/post",
		httpreq.SetRequestDebug(true, true),
		httpreq.AddRequestMiddlewares(httpreq.ExpectStatusCodeMW(200)),
		httpreq.SetRequestBody(Request{1, 10}))
	if err != nil {
		panic(err)
	}
	var result Result
	err = resp.Bind(&result)
	//err = resp.BindJSON(&result)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", result)
}
```

