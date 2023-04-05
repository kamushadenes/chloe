package structs

import (
	"io"
	"strings"
)

type ResponseObjectType string

const (
	Text  ResponseObjectType = "text"
	Image ResponseObjectType = "image"
	Audio ResponseObjectType = "audio"
)

type ResponseObject struct {
	Name    string             `json:"name"`
	Data    []byte             `json:"payload"`
	Type    ResponseObjectType `json:"type"`
	Result  bool               `json:"result"`
	readIdx int64
}

func NewResponseObject(objectType ResponseObjectType) *ResponseObject {
	return &ResponseObject{
		Type: objectType,
	}
}

func (ro *ResponseObject) SetName(name string) {
	ro.Name = strings.TrimSpace(name)
}

func (ro *ResponseObject) WriteObject(object *ResponseObject) error {
	ro.Name = object.Name
	ro.Data = object.Data
	ro.Type = object.Type
	ro.Result = object.Result

	return nil
}

func (ro *ResponseObject) Write(p []byte) (int, error) {
	ro.Data = append(ro.Data, p...)

	return len(p), nil
}

func (ro *ResponseObject) Flush() {}

func (ro *ResponseObject) Close() error {
	return nil
}

func (ro *ResponseObject) Read(p []byte) (int, error) {
	if ro.readIdx >= int64(len(ro.Data)) {
		return 0, io.EOF
	}

	n := copy(p, ro.Data[ro.readIdx:])
	ro.readIdx += int64(n)

	return n, nil
}

func (ro *ResponseObject) String() string {
	return string(ro.Data)
}

func (ro *ResponseObject) Bytes() []byte {
	return ro.Data
}

func (ro *ResponseObject) Size() int {
	return len(ro.Data)
}

func (ro *ResponseObject) Success() {
	ro.Result = true
}

func (ro *ResponseObject) Fail() {
	ro.Result = false
}

// TODO: add Store() method to ResponseObject
