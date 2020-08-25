package binding

import (
	"encoding/xml"
)

type xmlBinding struct{}

func (xmlBinding) Name() string {
	return "xml"
}

func (xmlBinding) BindBody(body []byte, obj interface{}) error {
	return xml.Unmarshal(body, obj)
}
