package freejson

import (
	"encoding/json"
	"reflect"
)

func ToArray(i interface{}) (Array, error) {
	switch v := i.(type) {
	case Array:
		return v, nil
	case []interface{}:
		return Array(v), nil
	}
	var a Array
	v := reflect.ValueOf(i)
	switch v.Kind() {
	case reflect.Array, reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			a = append(a, v.Index(i).Interface())
		}
		return a, nil
	}

	var data []byte
	switch t := i.(type) {
	case []byte:
		data = t
	case string:
		data = []byte(t)
	}
	err := json.Unmarshal(data, &a)
	if err != nil {
		return nil, err
	}
	return a, nil
}

func ToObject(i interface{}) (Object, error) {
	switch v := i.(type) {
	case Object:
		return v, nil
	case map[string]interface{}:
		return Object(v), nil
	}
	o := Object{}
	t := reflect.TypeOf(i)
	switch t.Kind() {
	case reflect.Map:
		if t.Key().Kind() != reflect.String {
			break
		}
		iter := reflect.ValueOf(i).MapRange()
		for iter.Next() {
			k := iter.Key()
			v := iter.Value()
			o[k.String()] = v.Interface()
		}
		return o, nil
	}

	var data []byte
	var err error
	switch t := i.(type) {
	case []byte:
		data = t
	case string:
		data = []byte(t)
	default:
		data, err = json.Marshal(i)
		if err != nil {
			return nil, err
		}
	}
	err = json.Unmarshal(data, &o)
	if err != nil {
		return nil, err
	}
	return o, nil

}

func AsArray(i interface{}) Array {
	a, _ := ToArray(i) //nolint
	return a
}

func AsObject(i interface{}) Object {
	o, _ := ToObject(i) //nolint
	return o
}

func AsArrayDefault(i interface{}, def Array) Array {
	a, err := ToArray(i)
	if err != nil {
		return def
	}
	return a
}

func AsObjectDefault(i interface{}, def Object) Object {
	o, err := ToObject(i)
	if err != nil {
		return def
	}
	return o
}
