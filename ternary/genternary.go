// +build ignore

package main

import (
	"fmt"
	"os"
	"strings"
)

var template = `// Code generated from template using genfunc.go; DO NOT EDIT.
package ternary

func %[1]s(condition bool, trueVal, falseVal %[2]s) %[2]s {
	if condition {
		return trueVal
	}
	return falseVal
}
`

func main() {
	for _, t := range []string{"Bool", "Float32", "Float64", "Int", "Uint", "String", "Interface",
		"Uint8", "Uint16", "Uint32", "Uint64",
		"Int8", "Int16", "Int32", "Int64"} {
		filename := fmt.Sprintf("%s.go", strings.ToLower(t))
		file, err := os.Create(filename)
		if err != nil {
			panic(err)
		}
		data := fmt.Sprintf(template, t, strings.ToLower(t))
		file.WriteString(data)
		file.Close()
	}
}
