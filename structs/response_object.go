package structs

import (
	"io"
	"net/http"
	"strings"
)

type ResponseObjectType string

const (
	Text  ResponseObjectType = "text"
	Image ResponseObjectType = "image"
	Audio ResponseObjectType = "audio"
)

type ResponseObject struct {
	Name           string             `json:"name"`
	Data           []byte             `json:"payload"`
	Type           ResponseObjectType `json:"type"`
	Result         bool               `json:"result"`
	readIdx        int64
	HTTPHeader     http.Header
	HTTPStatusCode int
}

func NewResponseObject(objectType ResponseObjectType) *ResponseObject {
	return &ResponseObject{
		Type:           objectType,
		HTTPStatusCode: http.StatusOK,
		HTTPHeader:     http.Header{},
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

func (ro *ResponseObject) GetStorableContent() string {
	switch ro.Type {
	case Image:
		return "[generated image]"
	case Audio:
		return "[generated audio]"
	case Text:
		return string(ro.Data)
	}

	return ""
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

func (ro *ResponseObject) WriteHeader(statusCode int) {
	ro.HTTPStatusCode = statusCode
}
func (ro *ResponseObject) Header() http.Header {
	return ro.HTTPHeader
}
