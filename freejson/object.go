package freejson

import (
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/pkg/errors"
	"github.com/y851592226/cate/convert"
	"github.com/y851592226/cate/encoding/json"
)

type ObjectFilter func(object Object, key string, value interface{}) bool
type ObjectProcessor func(object Object, key string, value interface{}) error
type Object map[string]interface{}

func (o *Object) Scan(src interface{}) error {
	b, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("unsupported type %T", src)
	}
	if len(b) == 0 {
		return nil
	}
	return json.Unmarshal(b, o)
}

func (o Object) Value() (driver.Value, error) {
	data, err := json.Marshal(o)
	return string(data), err
}

func (o Object) Len() int {
	return len(o)
}

func (o Object) Print() {
	fmt.Println(json.MarshalString(o))
}

func (o Object) PrettyPrint() {
	fmt.Println(json.MarshalIndentString(o))
}

func (o Object) Filter(f ObjectFilter) Object {
	r := Object{}
	for k, v := range o {
		if f(o, k, v) {
			r[k] = v
		}
	}
	return r
}

func (o Object) Each(f ObjectProcessor) error {
	for k, v := range o {
		err := f(o, k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

func (o Object) Exists(key string) bool {
	if o == nil {
		return false
	}
	_, ok := o[key]
	return ok
}

func (o Object) Get(key string) interface{} {
	if o == nil {
		return nil
	}
	return o[key]
}

func (o Object) Bool(key string) bool {
	return o[key].(bool)
}

func (o Object) Int(key string) int {
	return o[key].(int)
}

func (o Object) Int32(key string) int32 {
	return o[key].(int32)
}

func (o Object) Int64(key string) int64 {
	return o[key].(int64)
}

func (o Object) String(key string) string {
	return o[key].(string)
}

func (o Object) Float32(key string) float32 {
	return o[key].(float32)
}

func (o Object) Float64(key string) float64 {
	return o[key].(float64)
}

func (o Object) Object(key string) Object {
	switch v := o[key].(type) {
	case Object:
		return v
	case map[string]interface{}:
		return Object(v)
	}
	return o[key].(Object)
}

func (o Object) Array(key string) Array {
	switch v := o[key].(type) {
	case Array:
		return v
	case []interface{}:
		return Array(v)
	}
	return o[key].(Array)
}

func (o Object) Time(key string) time.Time {
	value := o[key]
	t, ok := value.(*time.Time)
	if ok {
		return *t
	}
	return value.(time.Time)
}

func (o Object) Interface(key string) interface{} {
	return o[key]
}

func (o Object) AsBool(key string, def ...bool) bool {
	if o.Exists(key) {
		return convert.AsBoolDefault(o[key], append(def, false)[0])
	}
	return append(def, false)[0]
}

func (o Object) AsInt(key string, def ...int) int {
	if o.Exists(key) {
		return convert.AsIntDefault(o[key], append(def, 0)[0])
	}
	return append(def, 0)[0]
}

func (o Object) AsInt32(key string, def ...int32) int32 {
	if o.Exists(key) {
		return convert.AsInt32Default(o[key], append(def, 0)[0])
	}
	return append(def, 0)[0]
}

func (o Object) AsInt64(key string, def ...int64) int64 {
	if o.Exists(key) {
		return convert.AsInt64Default(o[key], append(def, 0)[0])
	}
	return append(def, 0)[0]
}

func (o Object) AsString(key string, def ...string) string {
	if o.Exists(key) {
		return convert.AsStringDefault(o[key], append(def, "")[0])
	}
	return append(def, "")[0]
}

func (o Object) AsFloat32(key string, def ...float32) float32 {
	if o.Exists(key) {
		return convert.AsFloat32Default(o[key], append(def, 0)[0])
	}
	return append(def, 0)[0]
}

func (o Object) AsFloat64(key string, def ...float64) float64 {
	if o.Exists(key) {
		return convert.AsFloat64Default(o[key], append(def, 0)[0])
	}
	return append(def, 0)[0]
}

func (o Object) AsObject(key string, def ...Object) Object {
	if o.Exists(key) {
		return AsObjectDefault(o[key], append(def, nil)[0])
	}
	return append(def, nil)[0]
}

func (o Object) AsArray(key string, def ...Array) Array {
	if o.Exists(key) {
		return AsArrayDefault(o[key], append(def, nil)[0])
	}
	return append(def, nil)[0]
}

func (o Object) AsTime(key string, def ...time.Time) time.Time {
	if o.Exists(key) {
		return convert.AsTimeDefault(o[key], append(def, time.Time{})[0])
	}
	return append(def, time.Time{})[0]
}

func (o Object) AsInterface(key string, def ...interface{}) interface{} {
	if o.Exists(key) {
		return o[key]
	}
	return append(def, nil)[0]
}

func (o Object) AsStringSlice(key string, def ...[]string) []string {
	if o.Exists(key) {
		result := []string{}
		for _, v := range o.AsArray(key) {
			result = append(result, convert.AsString(v))
		}
		return result
	}
	return append(def, nil)[0]
}

func (o Object) ToBool(key string) (bool, error) {
	if o.Exists(key) {
		return convert.ToBool(o[key])
	}
	return false, errors.Errorf("object don't have key %s", key)
}

func (o Object) ToInt(key string) (int, error) {
	if o.Exists(key) {
		return convert.ToInt(o[key])
	}
	return 0, errors.Errorf("object don't have key %s", key)
}

func (o Object) ToInt32(key string) (int32, error) {
	if o.Exists(key) {
		return convert.ToInt32(o[key])
	}
	return 0, errors.Errorf("object don't have key %s", key)
}

func (o Object) ToInt64(key string) (int64, error) {
	if o.Exists(key) {
		return convert.ToInt64(o[key])
	}
	return 0, errors.Errorf("object don't have key %s", key)
}

func (o Object) ToString(key string) (string, error) {
	if o.Exists(key) {
		return convert.ToString(o[key])
	}
	return "", errors.Errorf("object don't have key %s", key)
}

func (o Object) ToFloat32(key string) (float32, error) {
	if o.Exists(key) {
		return convert.ToFloat32(o[key])
	}
	return 0, errors.Errorf("object don't have key %s", key)
}

func (o Object) ToFloat64(key string) (float64, error) {
	if o.Exists(key) {
		return convert.ToFloat64(o[key])
	}
	return 0, errors.Errorf("object don't have key %s", key)
}

func (o Object) ToObject(key string) (Object, error) {
	if o.Exists(key) {
		return ToObject(o[key])
	}
	return nil, errors.Errorf("object don't have key %s", key)
}

func (o Object) ToArray(key string) (Array, error) {
	if o.Exists(key) {
		return ToArray(o[key])
	}
	return nil, errors.Errorf("object don't have key %s", key)
}

func (o Object) ToTime(key string) (time.Time, error) {
	if o.Exists(key) {
		return convert.ToTime(o[key])
	}
	return time.Time{}, errors.Errorf("object don't have key %s", key)
}

func (o Object) ToInterface(key string) (interface{}, error) {
	if o.Exists(key) {
		return o[key], nil
	}
	return nil, errors.Errorf("object don't have key %s", key)
}

func (o Object) Bind(key string, v interface{}) error {
	if o.Exists(key) {
		data, err := json.Marshal(o[key])
		if err != nil {
			return err
		}
		return json.Unmarshal(data, v)
	}
	return errors.Errorf("object don't have key %s", key)
}
