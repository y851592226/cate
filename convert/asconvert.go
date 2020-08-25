package convert

import (
	"time"
)

func AsInt(i interface{}) int {
	return int(AsInt64(i))
}

func AsInt32(i interface{}) int32 {
	return int32(AsInt64(i))
}

func AsInt64(i interface{}) int64 {
	i64, _ := ToInt64(i) // nolint
	return i64
}

func AsFloat32(i interface{}) float32 {
	return float32(AsFloat64(i))
}

func AsFloat64(i interface{}) float64 {
	f64, _ := ToFloat64(i) // nolint
	return f64
}

func AsString(i interface{}) string {
	s, _ := ToString(i) // nolint
	return s
}

func AsBool(i interface{}) bool {
	b, _ := ToBool(i) // nolint
	return b
}

func AsByteSlice(i interface{}) []byte {
	bytes, _ := ToByteSlice(i) // nolint
	return bytes
}

func AsDuration(i interface{}) time.Duration {
	d, _ := ToDuration(i) // nolint
	return d
}

func AsTime(i interface{}) time.Time {
	t, _ := ToTime(i) // nolint
	return t
}
