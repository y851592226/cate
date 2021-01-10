package rand

import (
	"math/rand"
	"time"
)

var (
	Seed        = rand.Seed
	Int63       = rand.Int63
	Uint32      = rand.Uint32
	Uint64      = rand.Uint64
	Int31       = rand.Int31
	Int         = rand.Int
	Int63n      = rand.Int63n
	Int31n      = rand.Int31n
	Intn        = rand.Intn
	Float64     = rand.Float64
	Float32     = rand.Float32
	Perm        = rand.Perm
	Shuffle     = rand.Shuffle
	Read        = rand.Read
	NormFloat64 = rand.NormFloat64
	ExpFloat64  = rand.ExpFloat64
)

const (
	LowerLetters        = "abcdefghijklmnopqrstuvwxyz"
	UpperLetters        = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Numbers             = "0123456789"
	NumbersLowerLetters = Numbers + LowerLetters
	NumbersUpperLetters = Numbers + UpperLetters
	NumbersLetters      = Numbers + LowerLetters + UpperLetters
)

func init() {
	InitSeed()
}

func InitSeed() {
	rand.Seed(time.Now().UnixNano())
}

// [low,high)
func BetweenInt(low, high int) int {
	return low + rand.Intn(high-low)
}

// [low,high)
func BetweenInt32(low, high int32) int32 {
	return low + rand.Int31n(high-low)
}

// [low,high)
func BetweenInt64(low, high int64) int64 {
	return low + rand.Int63n(high-low)
}

// [low,high)
func BetweenFloat32(low, high float32) float32 {
	return low + (high-low)*rand.Float32()
}

// [low,high)
func BetweenFloat64(low, high float64) float64 {
	return low + (high-low)*rand.Float64()
}

func String(s string, n int) string {
	r := make([]byte, n)
	for i := 0; i < n; i++ {
		r[i] = s[rand.Intn(len(s))]
	}
	return string(r)
}

func choiseN(l, n int) []int {
	if l < 100 || l/n < 5 {
		order := rand.Perm(l)
		return order[:n]
	}
	order := make([]int, n)
	exists := map[int]struct{}{}
	count := 0
	for count < n {
		i := rand.Intn(l)
		if _, ok := exists[i]; !ok {
			exists[i] = struct{}{}
			order[count] = i
			count++
		}
	}
	return order

}

func ChoiceOneInt(l []int) int {
	return l[rand.Intn(len(l))]
}

func ChoiceNInt(l []int, n int) []int {
	order := choiseN(len(l), n)
	r := make([]int, n)
	for i := range r {
		r[i] = l[order[i]]
	}
	return r
}

func ChoiceOneInt32(l []int32) int32 {
	return l[rand.Intn(len(l))]
}

func ChoiceNInt32(l []int32, n int) []int32 {
	order := choiseN(len(l), n)
	r := make([]int32, n)
	for i := range r {
		r[i] = l[order[i]]
	}
	return r
}

func ChoiceOneInt64(l []int64) int64 {
	return l[rand.Intn(len(l))]
}

func ChoiceNInt64(l []int64, n int) []int64 {
	order := choiseN(len(l), n)
	r := make([]int64, n)
	for i := range r {
		r[i] = l[order[i]]
	}
	return r
}

func ChoiceOneFloat32(l []float32) float32 {
	return l[rand.Intn(len(l))]
}

func ChoiceNFloat32(l []float32, n int) []float32 {
	order := choiseN(len(l), n)
	r := make([]float32, n)
	for i := range r {
		r[i] = l[order[i]]
	}
	return r
}

func ChoiceOneFloat64(l []float64) float64 {
	return l[rand.Intn(len(l))]
}

func ChoiceNFloat64(l []float64, n int) []float64 {
	order := choiseN(len(l), n)
	r := make([]float64, n)
	for i := range r {
		r[i] = l[order[i]]
	}
	return r
}

func ChoiceOneInterface(l []interface{}) interface{} {
	return l[rand.Intn(len(l))]
}

func ChoiceNInterface(l []interface{}, n int) []interface{} {
	order := choiseN(len(l), n)
	r := make([]interface{}, n)
	for i := range r {
		r[i] = l[order[i]]
	}
	return r
}

func Rand100(n int) bool {
	return rand.Intn(100) < n
}

func Rand1000(n int) bool {
	return rand.Intn(1000) < n
}

func Rand10000(n int) bool {
	return rand.Intn(10000) < n
}
