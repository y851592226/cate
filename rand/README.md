# rand

> Rand tools

### Simple Example

```go
r := rand.BetweenFloat64(-0.1, 0.1)
```

### Doc

```go
func BetweenFloat32(low, high float32) float32
    [low,high)

func BetweenFloat64(low, high float64) float64
    [low,high)

func BetweenInt(low, high int) int
    [low,high)

func BetweenInt32(low, high int32) int32
    [low,high)

func BetweenInt64(low, high int64) int64
    [low,high)

func ChoiceNFloat32(l []float32, n int) []float32
func ChoiceNFloat64(l []float64, n int) []float64
func ChoiceNInt(l []int, n int) []int
func ChoiceNInt32(l []int32, n int) []int32
func ChoiceNInt64(l []int64, n int) []int64
func ChoiceNInterface(l []interface{}, n int) []interface{}
func ChoiceOneFloat32(l []float32) float32
func ChoiceOneFloat64(l []float64) float64
func ChoiceOneInt(l []int) int
func ChoiceOneInt32(l []int32) int32
func ChoiceOneInt64(l []int64) int64
func ChoiceOneInterface(l []interface{}) interface{}
func InitSeed()
func String(s string, n int) string
```

