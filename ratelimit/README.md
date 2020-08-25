# ratelimit

> Rate Limiting based on Token Bucket

### Simple Example

```go
tb := NewTokenBucket(5, 100, time.Second/10)
for i:=1; i<100; i++{
	ok := tb.Get()
	fmt.Println(ok)
	time.Sleep(time.Second/100)
}
```

### Doc

```go
type Limiter interface {
	BlockingGet() //Blocking until get a ticket
	Get() bool    //return true if get a ticket
}

func NewTokenBucket(num, capacity int64, fillInterval time.Duration) Limiter

type TokenBucket struct {
	Num      int64 //每次生成令牌的数量
	Capacity int64 //bucket的总容量

	// Has unexported fields.
}

func (tb *TokenBucket) BlockingGet()

func (tb *TokenBucket) BlockingGetSlow()

func (tb *TokenBucket) Get() bool

func (tb *TokenBucket) SetCapacity(capacity int64)

func (tb *TokenBucket) SetNum(num int64)
```

