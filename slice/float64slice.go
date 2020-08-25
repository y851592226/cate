// Code generated from template using genfunc.go; DO NOT EDIT.
package slice

func Float64Contains(items []float64, item float64) bool {
	return Float64Index(items, item) >= 0
}

func Float64Index(items []float64, item float64) int {
	for i, v := range items {
		if v == item {
			return i
		}
	}
	return -1
}

func Float64Reverse(items []float64) {
	l := len(items)
	if l <= 1 {
		return
	}
	for i := 0; i < l/2; i++ {
		items[i], items[l-i-1] = items[l-i-1], items[i]
	}
}
