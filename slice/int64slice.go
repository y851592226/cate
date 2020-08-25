// Code generated from template using genfunc.go; DO NOT EDIT.
package slice

func Int64Contains(items []int64, item int64) bool {
	return Int64Index(items, item) >= 0
}

func Int64Index(items []int64, item int64) int {
	for i, v := range items {
		if v == item {
			return i
		}
	}
	return -1
}

func Int64Reverse(items []int64) {
	l := len(items)
	if l <= 1 {
		return
	}
	for i := 0; i < l/2; i++ {
		items[i], items[l-i-1] = items[l-i-1], items[i]
	}
}
