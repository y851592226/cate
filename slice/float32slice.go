// Code generated from template using genfunc.go; DO NOT EDIT.
package slice

func Float32Contains(items []float32, item float32) bool {
	return Float32Index(items, item) >= 0
}

func Float32Index(items []float32, item float32) int {
	for i, v := range items {
		if v == item {
			return i
		}
	}
	return -1
}

func Float32Reverse(items []float32) {
	l := len(items)
	if l <= 1 {
		return
	}
	for i := 0; i < l/2; i++ {
		items[i], items[l-i-1] = items[l-i-1], items[i]
	}
}
