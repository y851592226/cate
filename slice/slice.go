//go:generate go run gentypeslice.go

package slice

import (
	"reflect"
)

func Contains(items []interface{}, item interface{}) bool {
	return Index(items, item) >= 0
}

func Index(items []interface{}, item interface{}) int {
	for i, v := range items {
		if reflect.DeepEqual(v, item) {
			return i
		}
	}
	return -1
}

func Reverse(items []interface{}) {
	l := len(items)
	if l <= 1 {
		return
	}
	for i := 0; i < l/2; i++ {
		items[i], items[l-i-1] = items[l-i-1], items[i]
	}
}
