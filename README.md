# cate

工具库，包含一些常用的组件或者库，把工作中常用的一些模块放到这个仓库下，减少重复开发工作，方便使用

> 有什么新的需求可以直接提，目地是为了让这个库更好用，服务开发，也欢迎大家提供新的模块

### cache/lrucache

> Golang LRUCache

简单的LRU 缓存

[文档](/cache/lrucache/README.md)

使用示例:

```go
opt := Options{
			Capacity: 3,
			OnDeleted: func(key, value interface{}) {
				fmt.Println(key, value)
			},
		}
cache, err := lrucache.NewCacher(&opt)
cache.Set("1", 1)
cache.Set("2", 2)
cache.Set("3", 3)
cache.Set("4", 4)
```

### cache/ttlcache

> Golang TTLCache

简单的TTL 缓存

[文档](/cache/ttlcache/README.md)

使用示例：

```go
var num int32
fetcher := func(ctx context.Context, key interface{}) (interface{}, error) {
			atomic.AddInt32(&num, 1)
			return num, nil
}
cacher := ttlcache.NewCache(ttlcache.Options{time.Second/2, fetcher})
cacher.Get(1) // 1
cacher.Get(1) // 1
cacher.Get(1) // 1
time.Sleep(time.Second)
cacher.Get(1)  // 2
cacher.Get(1)  // 2
cacher.Get(1)  // 2
```

### convert

> Golang type convert tool

包含的Golang的各种常见类型转换，免去每次类型转换时烦恼

[文档](/convert/README.md)

使用示例：

```go
i64,err := convert.ToInt64("1.23")
i64 := convert.AsInt64("1.23") 
i64 := convert.AsInt64Default("1.1.1" ,-1)
```

### crypto

> Encryption and decryption

易于使用的加解密工具

[文档](/crypto/README.md)

使用示例：

```go
type Payload struct {
			A int
		}
key := "1234567890abaaaa"
encoder := crypto.NewEncoder(key, aes.AES{})
payload := Payload{A: 12}
data, err := encoder.Encode(payload)
```

### encoding/json

对 `github.com/json-iterator/go` 的一个简单封装，同时包装了一些快捷操作

```
func MarshalDef(v interface{}, def []byte) []byte
func MarshalIndentString(v interface{}) string
func MarshalIndentStringDef(v interface{}, def string) string
func MarshalString(v interface{}) string
func MarshalStringDef(v interface{}, def string) string
```

### freejson

> Quickly get JSON values from map or slice

快速获取golang json结构的value，不需要预先定义结构体，同时也省去了类型转换的麻烦

[文档](/freejson/README.md)

使用示例：

```go
func ExampleFreejson() {
	msg := `{"PageNo":1,"Content":"test message","Data":[1,"12",1.1]}`
	o, err := freejson.ToObject(msg)
	if err != nil {
		panic(err)
	}
	fmt.Println(o.Float64("PageNo"))         // 1
	fmt.Println(o.String("Content"))         // "test message"
	fmt.Println(o.Array("Data").StringAt(1)) // "12"
	fmt.Println(o.AsString("PageNo"))        // "1"
	fmt.Println(o.AsInt("Content", -1))      // -1
	fmt.Println(o.Array("Data").AsIntAt(1))  // 12
}
```



### httpreq

> Easy-to-use http client for Golang

封装了Go的http client，更方便的发起http请求和解析请求数据，支持Debug和添加中间件

[文档](/httpreq/README.md)

使用示例：

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

### pool

> Golang persistent object pool，unlike sync.Pool

持久化的对象池，免去不停的创建和消费对象

[文档](/pool/README.md)

使用示例：

```
func NewConn() (interface{},error){
	return conn，nil
}
opt := Options{
		NewFunc:     newV,
		MaxSize:     2,
		MaxIdleSize: 1,
		MaxIdlelime: time.Second * 1,
	}
p := NewPool(&opt)
```

### rand

> Rand tools

提供常用的随机函数

[文档](/rand/README.md)

使用示例：

```go
r := rand.BetweenFloat64(-0.1, 0.1)
```

### ratelimit

> Rate Limiting based on Token Bucket

基于令牌桶的限流器

[文档](/ratelimit/README.md)

使用示例：

```go
tb := NewTokenBucket(5, 100, time.Second/10)
for i:=1; i<100; i++{
	ok := tb.Get()
	fmt.Println(ok)
	time.Sleep(time.Second/100)
}
```

### slice

> 切片的常用函数

包含了一些切片的常用函数

- **Type**Contains   判断是否包含某一个元素
- **Type**Index        返回元素的位置
- **Type**Reverse    切片反转

[文档](/slice/README.md)

使用示例：

```go
items = []interface{}{1, 2, "3", 4}
fmt.Println(slice.Index(items, "3")
```



### snowflake

> 分布式唯一ID

#### 原理

基于snowflake的思想进行开发，有以下几个优点

- 不需要手动配置节点ID
- 支持配置节点ID
- 支持设置各段长度
- 不会受到时钟回拨的影响

### task

> Concurrent task control

用来控制并发任务的执行，可以统一管理返回错误和设置超时时间

使用示例：

```go
func ExampleTasks() {
	var a int64
	task0 := func(ctx context.Context) error {
		time.Sleep(time.Second)
		atomic.AddInt64(&a,1)
		return nil
	}
	task1 := func(ctx context.Context) error {
		time.Sleep(time.Second*2)
		atomic.AddInt64(&a,2)
		return nil
	}
	task2 := func(ctx context.Context) error {
		time.Sleep(time.Second*3)
		atomic.AddInt64(&a,4)
		return nil
	}
	tasks := task.NewTasks()
	tasks.AddTaskFunc(task0)
	tasks.AddTaskFunc(task1)
	tasks.AddTaskFunc(task2)
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*2500)
	defer cancel()
	tasks.Run(ctx)
	err := tasks.Wait()
	fmt.Println(err) // task:2 error:context deadline exceeded
	fmt.Println(a) // 3
}
```

### 

