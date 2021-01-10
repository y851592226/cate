// Code generated from template using genfunc.go; DO NOT EDIT.
package ternary

func Uint(condition bool, trueVal, falseVal uint) uint {
	if condition {
		return trueVal
	}
	return falseVal
}
