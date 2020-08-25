package binding

import (
	"encoding/json"
)

type jsonBinding struct{}

func (jsonBinding) Name() string {
	return "json"
}

func (jsonBinding) BindBody(body []byte, obj interface{}) error {
	return json.Unmarshal(body, obj)
}
