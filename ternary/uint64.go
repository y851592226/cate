// Code generated from template using genfunc.go; DO NOT EDIT.
package ternary

func Uint64(condition bool, trueVal, falseVal uint64) uint64 {
	if condition {
		return trueVal
	}
	return falseVal
}
