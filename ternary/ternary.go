//go:generate go run genternary.go

package ternary

func Interface(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}
