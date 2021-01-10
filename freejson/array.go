package freejson

import (
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/pkg/errors"
	"github.com/y851592226/cate/convert"
	"github.com/y851592226/cate/encoding/json"
)

type ArrayFilter func(array Array, index int, item interface{}) bool
type ArrayProcessor func(array Array, index int, item interface{}) error
type Array []interface{}

func (a *Array) Scan(src interface{}) error {
	b, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("unsupported type %T", src)
	}
	if len(b) == 0 {
		return nil
	}
	return json.Unmarshal(b, a)
}

func (a Array) Value() (driver.Value, error) {
	data, err := json.Marshal(a)
	return string(data), err
}

func (a Array) Len() int {
	return len(a)
}

func (a Array) Print() {
	fmt.Println(json.MarshalString(a))
}

func (a Array) PrettyPrint() {
	fmt.Println(json.MarshalIndentString(a))
}

func (a Array) Filter(f ArrayFilter) Array {
	r := Array{}
	for i, item := range a {
		if f(a, i, item) {
			r = append(r, item)
		}
	}
	return r
}

func (a Array) Each(f ArrayProcessor) error {
	for i, item := range a {
		err := f(a, i, item)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a Array) Exists(i int) bool {
	if i < 0 {
		return false
	}
	if len(a) > i {
		return true
	}
	return false
}

func (a Array) At(i int) interface{} {
	if a == nil {
		return nil
	}
	return a[i]
}

func (a Array) First() interface{} {
	if len(a) > 0 {
		return a[0]
	}
	return nil
}

func (a Array) Last() interface{} {
	if len(a) > 0 {
		return a[len(a)-1]
	}
	return nil
}

func (a Array) BoolAt(i int) bool {
	return a[i].(bool)
}

func (a Array) IntAt(i int) int {
	return a[i].(int)
}

func (a Array) Int32At(i int) int32 {
	return a[i].(int32)
}

func (a Array) Int64At(i int) int64 {
	return a[i].(int64)
}

func (a Array) StringAt(i int) string {
	return a[i].(string)
}

func (a Array) Float32At(i int) float32 {
	return a[i].(float32)
}

func (a Array) Float64At(i int) float64 {
	return a[i].(float64)
}

func (a Array) ObjectAt(i int) Object {
	switch v := a[i].(type) {
	case Object:
		return v
	case map[string]interface{}:
		return Object(v)
	}
	return a[i].(Object)
}

func (a Array) ArrayAt(i int) Array {
	switch v := a[i].(type) {
	case Array:
		return v
	case []interface{}:
		return Array(v)
	}
	return a[i].(Array)
}

func (a Array) TimeAt(i int) time.Time {
	item := a[i]
	t, ok := item.(*time.Time)
	if ok {
		return *t
	}
	return item.(time.Time)
}

func (a Array) InterfaceAt(i int) interface{} {
	return a[i]
}

func (a Array) AsBoolAt(i int, def ...bool) bool {
	if a.Exists(i) {
		return convert.AsBoolDefault(a[i], append(def, false)[0])
	}
	return append(def, false)[0]
}

func (a Array) AsIntAt(i int, def ...int) int {
	if a.Exists(i) {
		return convert.AsIntDefault(a[i], append(def, 0)[0])
	}
	return append(def, 0)[0]
}

func (a Array) AsInt32At(i int, def ...int32) int32 {
	if a.Exists(i) {
		return convert.AsInt32Default(a[i], append(def, 0)[0])
	}
	return append(def, 0)[0]
}

func (a Array) AsInt64At(i int, def ...int64) int64 {
	if a.Exists(i) {
		return convert.AsInt64Default(a[i], append(def, 0)[0])
	}
	return append(def, 0)[0]
}

func (a Array) AsStringAt(i int, def ...string) string {
	if a.Exists(i) {
		return convert.AsStringDefault(a[i], append(def, "")[0])
	}
	return append(def, "")[0]
}

func (a Array) AsFloat32At(i int, def ...float32) float32 {
	if a.Exists(i) {
		return convert.AsFloat32Default(a[i], append(def, 0)[0])
	}
	return append(def, 0)[0]
}

func (a Array) AsFloat64At(i int, def ...float64) float64 {
	if a.Exists(i) {
		return convert.AsFloat64Default(a[i], append(def, 0)[0])
	}
	return append(def, 0)[0]
}

func (a Array) AsObjectAt(i int, def ...Object) Object {
	if a.Exists(i) {
		return AsObjectDefault(a[i], append(def, nil)[0])
	}
	return append(def, nil)[0]
}

func (a Array) AsArrayAt(i int, def ...Array) Array {
	if a.Exists(i) {
		return AsArrayDefault(a[i], append(def, nil)[0])
	}
	return append(def, nil)[0]
}

func (a Array) AsTimeAt(i int, def ...time.Time) time.Time {
	if a.Exists(i) {
		return convert.AsTimeDefault(a[i], append(def, time.Time{})[0])
	}
	return append(def, time.Time{})[0]
}

func (a Array) AsInterfaceAt(i int, def ...interface{}) interface{} {
	if a.Exists(i) {
		return a[i]
	}
	return append(def, nil)[0]
}

func (a Array) To(v interface{}) error {
	data, err := json.Marshal(a)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}

func (a Array) ToBoolAt(i int) (bool, error) {
	if a.Exists(i) {
		return convert.ToBool(a[i])
	}
	return false, errors.Errorf("array don't have index %d", i)
}

func (a Array) ToIntAt(i int) (int, error) {
	if a.Exists(i) {
		return convert.ToInt(a[i])
	}
	return 0, errors.Errorf("array don't have index %d", i)
}

func (a Array) ToInt32At(i int) (int32, error) {
	if a.Exists(i) {
		return convert.ToInt32(a[i])
	}
	return 0, errors.Errorf("array don't have index %d", i)
}

func (a Array) ToInt64At(i int) (int64, error) {
	if a.Exists(i) {
		return convert.ToInt64(a[i])
	}
	return 0, errors.Errorf("array don't have index %d", i)
}

func (a Array) ToStringAt(i int) (string, error) {
	if a.Exists(i) {
		return convert.ToString(a[i])
	}
	return "", errors.Errorf("array don't have index %d", i)
}

func (a Array) ToFloat32At(i int) (float32, error) {
	if a.Exists(i) {
		return convert.ToFloat32(a[i])
	}
	return 0, errors.Errorf("array don't have index %d", i)
}

func (a Array) ToFloat64At(i int) (float64, error) {
	if a.Exists(i) {
		return convert.ToFloat64(a[i])
	}
	return 0, errors.Errorf("array don't have index %d", i)
}

func (a Array) ToObjectAt(i int) (Object, error) {
	if a.Exists(i) {
		return ToObject(a[i])
	}
	return nil, errors.Errorf("array don't have index %d", i)
}

func (a Array) ToArrayAt(i int) (Array, error) {
	if a.Exists(i) {
		return ToArray(a[i])
	}
	return nil, errors.Errorf("array don't have index %d", i)
}

func (a Array) ToTimeAt(i int) (time.Time, error) {
	if a.Exists(i) {
		return convert.ToTime(a[i])
	}
	return time.Time{}, errors.Errorf("array don't have index %d", i)
}

func (a Array) ToInterfaceAt(i int) (interface{}, error) {
	if a.Exists(i) {
		return a[i], nil
	}
	return nil, errors.Errorf("array don't have index %d", i)
}

func (a Array) BindAt(i int, v interface{}) error {
	if a.Exists(i) {
		data, err := json.Marshal(a[i])
		if err != nil {
			return err
		}
		return json.Unmarshal(data, v)
	}
	return errors.Errorf("array don't have index %d", i)
}

func (a Array) AsStringSlice() []string {
	result := []string{}
	for _, item := range a {
		result = append(result, convert.AsString(item))
	}
	return result
}

func (a Array) AsIntSlice() []int {
	result := []int{}
	for _, item := range a {
		result = append(result, convert.AsInt(item))
	}
	return result
}

func (a Array) AsInt64Slice() []int64 {
	result := []int64{}
	for _, item := range a {
		result = append(result, convert.AsInt64(item))
	}
	return result
}

func (a Array) AsFloat64Slice() []float64 {
	result := []float64{}
	for _, item := range a {
		result = append(result, convert.AsFloat64(item))
	}
	return result
}

func (a Array) ToStringSlice() ([]string, error) {
	result := []string{}
	for _, item := range a {
		v, err := convert.ToString(item)
		if err != nil {
			return nil, err
		}
		result = append(result, v)
	}
	return result, nil
}

func (a Array) ToIntSlice() ([]int, error) {
	result := []int{}
	for _, item := range a {
		v, err := convert.ToInt(item)
		if err != nil {
			return nil, err
		}
		result = append(result, v)
	}
	return result, nil
}

func (a Array) ToInt64Slice() ([]int64, error) {
	result := []int64{}
	for _, item := range a {
		v, err := convert.ToInt64(item)
		if err != nil {
			return nil, err
		}
		result = append(result, v)
	}
	return result, nil
}

func (a Array) ToFloat64Slice() ([]float64, error) {
	result := []float64{}
	for _, item := range a {
		v, err := convert.ToFloat64(item)
		if err != nil {
			return nil, err
		}
		result = append(result, v)
	}
	return result, nil
}
