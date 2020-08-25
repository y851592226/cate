# freejson

> quickly get JSON values from map or slice

### Simple Example

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



### Array

```go
func (a Array) ArrayAt(i int) Array
func (a Array) AsArrayAt(i int, def ...Array) Array
func (a Array) AsBoolAt(i int, def ...bool) bool
func (a Array) AsFloat32At(i int, def ...float32) float32
func (a Array) AsFloat64At(i int, def ...float64) float64
func (a Array) AsInt32At(i int, def ...int32) int32
func (a Array) AsInt64At(i int, def ...int64) int64
func (a Array) AsIntAt(i int, def ...int) int
func (a Array) AsInterfaceAt(i int, def ...interface{}) interface{}
func (a Array) AsObjectAt(i int, def ...Object) Object
func (a Array) AsStringAt(i int, def ...string) string
func (a Array) AsTimeAt(i int, def ...time.Time) time.Time
func (a Array) At(i int) interface{}
func (a Array) BoolAt(i int) bool
func (a Array) Each(f ArrayProcessor) error
func (a Array) Exists(i int) bool
func (a Array) Filter(f ArrayFilter) Array
func (a Array) First() interface{}
func (a Array) Float32At(i int) float32
func (a Array) Float64At(i int) float64
func (a Array) Int32At(i int) int32
func (a Array) Int64At(i int) int64
func (a Array) IntAt(i int) int
func (a Array) InterfaceAt(i int) interface{}
func (a Array) Last() interface{}
func (a Array) Len() int
func (a Array) ObjectAt(i int) Object
func (a Array) StringAt(i int) string
func (a Array) TimeAt(i int) time.Time
func (a Array) ToArrayAt(i int) (Array, error)
func (a Array) ToBoolAt(i int) (bool, error)
func (a Array) ToFloat32At(i int) (float32, error)
func (a Array) ToFloat64At(i int) (float64, error)
func (a Array) ToInt32At(i int) (int32, error)
func (a Array) ToInt64At(i int) (int64, error)
func (a Array) ToIntAt(i int) (int, error)
func (a Array) ToInterfaceAt(i int) (interface{}, error)
func (a Array) ToObjectAt(i int) (Object, error)
func (a Array) ToStringAt(i int) (string, error)
func (a Array) ToTimeAt(i int) (time.Time, error)
```

### Object

```go
func (o Object) Array(key string) Array
func (o Object) AsArray(key string, def ...Array) Array
func (o Object) AsBool(key string, def ...bool) bool
func (o Object) AsFloat32(key string, def ...float32) float32
func (o Object) AsFloat64(key string, def ...float64) float64
func (o Object) AsInt(key string, def ...int) int
func (o Object) AsInt32(key string, def ...int32) int32
func (o Object) AsInt64(key string, def ...int64) int64
func (o Object) AsInterfaceAt(key string, def ...interface{}) interface{}
func (o Object) AsObject(key string, def ...Object) Object
func (o Object) AsString(key string, def ...string) string
func (o Object) AsTimeAt(key string, def ...time.Time) time.Time
func (o Object) Bool(key string) bool
func (o Object) Each(f ObjectProcessor) error
func (o Object) Exists(key string) bool
func (o Object) Filter(f ObjectFilter) Object
func (o Object) Float32(key string) float32
func (o Object) Float64(key string) float64
func (o Object) Int(key string) int
func (o Object) Int32(key string) int32
func (o Object) Int64(key string) int64
func (o Object) Interface(key string) interface{}
func (o Object) Len() int
func (o Object) Object(key string) Object
func (o Object) String(key string) string
func (o Object) Time(key string) time.Time
func (o Object) ToArray(key string) (Array, error)
func (o Object) ToBool(key string) (bool, error)
func (o Object) ToFloat32(key string) (float32, error)
func (o Object) ToFloat64(key string) (float64, error)
func (o Object) ToInt(key string) (int, error)
func (o Object) ToInt32(key string) (int32, error)
func (o Object) ToInt64(key string) (int64, error)
func (o Object) ToInterfaceAt(key string) (interface{}, error)
func (o Object) ToObject(key string) (Object, error)
func (o Object) ToString(key string) (string, error)
func (o Object) ToTimeAt(key string) (time.Time, error)
func (o Object) Value(key string) interface{}
```

