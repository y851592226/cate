# pool

> Golang persistent object pool，unlike sync.Pool

### Simple Example

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

### Doc

```go
type Options struct {
	NewFunc     func() (interface{}, error)
	MaxRetries  int
	MaxSize     int
	MaxIdleSize int
	MaxIdlelime time.Duration
}

type Pool interface {
	Get(ctx context.Context) (interface{}, error)
	Put(interface{}, error)
}

func NewPool(opt *Options) Pool

type Value struct {
	// Has unexported fields.
}

func (v *Value) Expired() bool
```

