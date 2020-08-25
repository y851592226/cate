package convert

import (
	"strconv"
	"time"
)

func ToInt(i interface{}) (int, error) {
	i64, err := ToInt64(i)
	return int(i64), err
}

func ToInt32(i interface{}) (int32, error) {
	i64, err := ToInt64(i)
	return int32(i64), err
}

func ToInt64(i interface{}) (int64, error) {
	switch v := i.(type) {
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case uint8:
		return int64(v), nil
	case uint16:
		return int64(v), nil
	case uint32:
		return int64(v), nil
	case uint64:
		return int64(v), nil
	case int8:
		return int64(v), nil
	case int16:
		return int64(v), nil
	case int32:
		return int64(v), nil
	case int64:
		return v, nil
	case float32:
		return int64(v), nil
	case float64:
		return int64(v), nil
	case int:
		return int64(v), nil
	case uint:
		return int64(v), nil
	case string:
		return strconv.ParseInt(v, 10, 64)
	case []byte:
		return strconv.ParseInt(string(v), 10, 64)
	default:
		return 0, unSupportTypeError(i)
	}
}

func ToFloat32(i interface{}) (float32, error) {
	f64, err := ToFloat64(i)
	return float32(f64), err
}

func ToFloat64(i interface{}) (float64, error) {
	switch v := i.(type) {
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case uint8:
		return float64(v), nil
	case uint16:
		return float64(v), nil
	case uint32:
		return float64(v), nil
	case uint64:
		return float64(v), nil
	case int8:
		return float64(v), nil
	case int16:
		return float64(v), nil
	case int32:
		return float64(v), nil
	case int64:
		return float64(v), nil
	case float32:
		return float64(v), nil
	case float64:
		return v, nil
	case int:
		return float64(v), nil
	case uint:
		return float64(v), nil
	case []byte:
		return strconv.ParseFloat(string(v), 64)
	case string:
		return strconv.ParseFloat(v, 64)
	default:
		return 0, unSupportTypeError(i)
	}
}

func ToString(i interface{}) (string, error) {
	switch v := i.(type) {
	case bool:
		if v {
			return "true", nil
		}
		return "false", nil
	case []byte:
		return string(v), nil
	case string:
		return v, nil
	case float32:
		return strconv.FormatFloat(float64(v), 'f', -1, 32), nil
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64), nil
	case uint8:
		return strconv.FormatInt(int64(v), 10), nil
	case uint16:
		return strconv.FormatInt(int64(v), 10), nil
	case uint32:
		return strconv.FormatInt(int64(v), 10), nil
	case uint64:
		return strconv.FormatInt(int64(v), 10), nil
	case int8:
		return strconv.FormatInt(int64(v), 10), nil
	case int16:
		return strconv.FormatInt(int64(v), 10), nil
	case int32:
		return strconv.FormatInt(int64(v), 10), nil
	case int64:
		return strconv.FormatInt(v, 10), nil
	case int:
		return strconv.FormatInt(int64(v), 10), nil
	case uint:
		return strconv.FormatInt(int64(v), 10), nil
	case time.Time:
		return v.Format(DefaultTimeFormat), nil
	case *time.Time:
		return v.Format(DefaultTimeFormat), nil
	case error:
		return v.Error(), nil
	default:
		type stringer interface {
			String() string
		}
		if v, ok := i.(stringer); ok {
			return v.String(), nil
		}
		return "", unSupportTypeError(i)
	}
}

func ToBool(i interface{}) (bool, error) {
	switch v := i.(type) {
	case bool:
		return v, nil
	case string:
		switch v {
		case "true", "True", "1":
			return true, nil
		case "false", "False", "0":
			return false, nil
		default:
			return false, unSupportTypeError(i)
		}
	case []byte:
		return ToBool(string(v))
	case uint8:
		return v != 0, nil
	case uint16:
		return v != 0, nil
	case uint32:
		return v != 0, nil
	case uint64:
		return v != 0, nil
	case int8:
		return v != 0, nil
	case int16:
		return v != 0, nil
	case int32:
		return v != 0, nil
	case int64:
		return v != 0, nil
	case float32:
		return v != 0, nil
	case float64:
		return v != 0, nil
	case int:
		return v != 0, nil
	case uint:
		return v != 0, nil
	default:
		return false, unSupportTypeError(i)
	}
}

func ToByteSlice(i interface{}) ([]byte, error) {
	switch v := i.(type) {
	case string:
		return []byte(v), nil
	case []byte:
		return v, nil
	default:
		return nil, unSupportTypeError(i)
	}
}

func ToDuration(i interface{}) (time.Duration, error) {
	switch v := i.(type) {
	case time.Duration:
		return v, nil
	case string:
		return time.ParseDuration(v)
	case int:
		return time.Duration(int64(v)), nil
	case int64:
		return time.Duration(v), nil
	default:
		return 0, unSupportTypeError(i)
	}
}

func ToTime(i interface{}) (t time.Time, err error) {
	switch v := i.(type) {
	case time.Time:
		return v, nil
	case string:
		for _, format := range []string{BaseTimeFormat,
			time.RFC3339, time.RFC3339Nano, DefaultTimeFormat} {
			t, err = time.ParseInLocation(format, v, time.Local)
			if err == nil {
				return t, nil
			}
		}
		return t, unSupportTimeFormatError(v)
	case int:
		t = time.Unix(int64(v), 0)
		return t, nil
	case int64:
		t = time.Unix(v/1e9, v%1e9)
		return t, nil
	default:
		return t, unSupportTypeError(i)
	}
}
