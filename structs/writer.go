package structs

import "strings"

type ResponseObjectType string

const (
	Text  ResponseObjectType = "text"
	Image ResponseObjectType = "image"
	Audio ResponseObjectType = "audio"
)

type ResponseObject struct {
	Name    string             `json:"name"`
	Payload []byte             `json:"payload"`
	Type    ResponseObjectType `json:"type"`
	Result  bool               `json:"result"`
}

type ChloeWriter interface {
	WriteObject(object ResponseObject) error
	Write([]byte) (int, error)
	Close() error
	Flush()
}

func NewResponseObject(objectType ResponseObjectType) ResponseObject {
	return ResponseObject{
		Type: objectType,
	}
}

func (ro *ResponseObject) SetName(name string) {
	ro.Name = strings.TrimSpace(name)
}

func (ro *ResponseObject) SetPayload(payload []byte) {
	ro.Payload = payload
}

func (ro *ResponseObject) Success() {
	ro.Result = true
}

func (ro *ResponseObject) Fail() {
	ro.Result = false
}
