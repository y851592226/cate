package binding

const (
	MIMEJSON = "application/json"
	MIMEXML  = "application/xml"
	MIMEXML2 = "text/xml"
)

type Binding interface {
	Name() string
	BindBody([]byte, interface{}) error
}

var (
	JSON = jsonBinding{}
	XML  = xmlBinding{}
)

func Default(contentType string) Binding {
	contentType = filterFlags(contentType)
	switch contentType {
	case MIMEJSON:
		return JSON
	case MIMEXML, MIMEXML2:
		return XML
	default:
		return JSON
	}
}
