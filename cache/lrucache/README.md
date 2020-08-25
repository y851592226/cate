# lrucache

> Golang LRUCache

### Simple Example

```go
opt := Options{
			Capacity: 3,
			OnDeleted: func(key, value interface{}) {
				fmt.Println(key, value)
			},
		}
cache, err := NewCacher(&opt)
cache.Set("1", 1)
cache.Set("2", 2)
cache.Set("3", 3)
cache.Set("4", 4)
```

### Doc

```go
type Cacher interface {
	Set(interface{}, interface{})
	Get(interface{}) (interface{}, bool)
	GetDefault(interface{}, interface{}) interface{}
	GetInt(interface{}) (int, bool)
	GetIntDefault(interface{}, int) int
	GetString(interface{}) (string, bool)
	GetStringDefault(interface{}, string) string
	Len() int
}

func NewCacher(opt *Options) (Cacher, error)

type Options struct {
	Capacity  int
	OnDeleted func(key, value interface{})
}
```

### Note

- The method Get**Type** will panic if the value isn‘t the type expect

  **for example**

  ```go
  cache.Set(1,1)
  cache.GetString(1) //will panic
  ```

  