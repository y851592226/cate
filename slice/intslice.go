// Code generated from template using genfunc.go; DO NOT EDIT.
package slice

func IntContains(items []int, item int) bool {
	return IntIndex(items, item) >= 0
}

func IntIndex(items []int, item int) int {
	for i, v := range items {
		if v == item {
			return i
		}
	}
	return -1
}

func IntReverse(items []int) {
	l := len(items)
	if l <= 1 {
		return
	}
	for i := 0; i < l/2; i++ {
		items[i], items[l-i-1] = items[l-i-1], items[i]
	}
}
