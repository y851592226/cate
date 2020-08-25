package json

import (
	"encoding/json"

	jsoniter "github.com/json-iterator/go"
)

var (
	Jsoniter      = jsoniter.ConfigCompatibleWithStandardLibrary
	Marshal       = Jsoniter.Marshal
	Unmarshal     = Jsoniter.Unmarshal
	MarshalIndent = Jsoniter.MarshalIndent
	NewDecoder    = Jsoniter.NewDecoder
	NewEncoder    = Jsoniter.NewEncoder
)

func MarshalDef(v interface{}, def []byte) []byte {
	data, err := Marshal(v)
	if err != nil {
		return def
	}
	return data
}

func MarshalString(v interface{}) string {
	data, _ := Marshal(v) //nolint
	return string(data)
}

func MarshalStringDef(v interface{}, def string) string {
	data, err := Marshal(v)
	if err != nil {
		return def
	}
	return string(data)
}

func MarshalIndentString(v interface{}) string {
	data, _ := json.MarshalIndent(v, "", "    ") //nolint
	return string(data)
}

func MarshalIndentStringDef(v interface{}, def string) string {
	data, err := MarshalIndent(v, "", "    ")
	if err != nil {
		return def
	}
	return string(data)
}
