package convert

import (
	"time"
)

func AsIntDefault(i interface{}, def int) int {
	i64, err := ToInt64(i)
	if err != nil {
		return def
	}
	return int(i64)
}

func AsInt32Default(i interface{}, def int32) int32 {
	i64, err := ToInt64(i)
	if err != nil {
		return def
	}
	return int32(i64)
}

func AsInt64Default(i interface{}, def int64) int64 {
	i64, err := ToInt64(i)
	if err != nil {
		return def
	}
	return i64
}

func AsFloat32Default(i interface{}, def float32) float32 {
	f64, err := ToFloat64(i)
	if err != nil {
		return def
	}
	return float32(f64)
}

func AsFloat64Default(i interface{}, def float64) float64 {
	f64, err := ToFloat64(i)
	if err != nil {
		return def
	}
	return f64
}

func AsStringDefault(i interface{}, def string) string {
	s, err := ToString(i)
	if err != nil {
		return def
	}
	return s
}

func AsBoolDefault(i interface{}, def bool) bool {
	b, err := ToBool(i)
	if err != nil {
		return def
	}
	return b
}

func AsByteSliceDefault(i interface{}, def []byte) []byte {
	bytes, err := ToByteSlice(i)
	if err != nil {
		return def
	}
	return bytes
}

func AsDurationDefault(i interface{}, def time.Duration) time.Duration {
	d, err := ToDuration(i)
	if err != nil {
		return def
	}
	return d
}

func AsTimeDefault(i interface{}, def time.Time) time.Time {
	t, err := ToTime(i)
	if err != nil {
		return def
	}
	return t
}
