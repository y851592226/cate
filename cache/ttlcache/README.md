# ttlcache

> Golang TTLCache

### Simple Example

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

### Doc

```go
type Cacher interface {
	Get(context.Context, interface{}) (interface{}, error)
	GetDefault(context.Context, interface{}, interface{}) interface{}
}

func NewCache(opt Options) Cacher

type Options struct {
	TTL     time.Duration
	Fetcher func(context.Context, interface{}) (interface{}, error)
}
```

