// +build ignore

package main

import (
	"fmt"
	"os"
	"strings"
)

var template = `// Code generated from template using genfunc.go; DO NOT EDIT.
package slice

func %[1]sContains(items []%[2]s, item %[2]s) bool {
	return %[1]sIndex(items, item) >= 0
}

func %[1]sIndex(items []%[2]s, item %[2]s) int {
	for i, v := range items {
		if v == item {
			return i
		}
	}
	return -1
}

func %[1]sReverse(items []%[2]s) {
	l := len(items)
	if l <= 1 {
		return
	}
	for i := 0; i < l/2; i++ {
		items[i], items[l-i-1] = items[l-i-1], items[i]
	}
}
`

func main() {
	for _, t := range []string{"Int", "Int64", "String", "Float32", "Float64"} {
		filename := fmt.Sprintf("%sslice.go", strings.ToLower(t))
		file, err := os.Create(filename)
		if err != nil {
			panic(err)
		}
		data := fmt.Sprintf(template, t, strings.ToLower(t))
		file.WriteString(data)
		file.Close()
	}
}
