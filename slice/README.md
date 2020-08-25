# slice

> provide some common methods for slice

- TypeContains 
- TypeIndex
- TypeReverse  

### Simple Example

```go
items = []interface{}{1, 2, "3", 4}
fmt.Println(slice.Index(items, "3")
```

### Doc

```go
func Contains(items []interface{}, item interface{}) bool
func Float32Contains(items []float32, item float32) bool
func Float32Index(items []float32, item float32) int
func Float32Reverse(items []float32)
func Float64Contains(items []float64, item float64) bool
func Float64Index(items []float64, item float64) int
func Float64Reverse(items []float64)
func Index(items []interface{}, item interface{}) int
func Int64Contains(items []int64, item int64) bool
func Int64Index(items []int64, item int64) int
func Int64Reverse(items []int64)
func IntContains(items []int, item int) bool
func IntIndex(items []int, item int) int
func IntReverse(items []int)
func Reverse(items []interface{})
func StringContains(items []string, item string) bool
func StringIndex(items []string, item string) int
func StringReverse(items []string)
```

