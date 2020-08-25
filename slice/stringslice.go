// Code generated from template using genfunc.go; DO NOT EDIT.
package slice

func StringContains(items []string, item string) bool {
	return StringIndex(items, item) >= 0
}

func StringIndex(items []string, item string) int {
	for i, v := range items {
		if v == item {
			return i
		}
	}
	return -1
}

func StringReverse(items []string) {
	l := len(items)
	if l <= 1 {
		return
	}
	for i := 0; i < l/2; i++ {
		items[i], items[l-i-1] = items[l-i-1], items[i]
	}
}
