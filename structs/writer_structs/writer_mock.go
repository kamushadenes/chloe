package writer_structs

import (
	"net/http"

	"github.com/kamushadenes/chloe/structs/response_object_structs"
)

type MockWriter struct {
	objs             []*response_object_structs.ResponseObject
	HTTPHeader       http.Header
	HTTPStatusCode   int
	preWriteCallback func()
}

func NewMockWriter() *MockWriter {
	return &MockWriter{
		HTTPHeader:     http.Header{},
		HTTPStatusCode: 200,
	}
}

func (w *MockWriter) WriteObject(obj *response_object_structs.ResponseObject) error {
	w.objs = append(w.objs, obj)

	return nil
}

func (w *MockWriter) Write(b []byte) (int, error) {
	if w.preWriteCallback != nil {
		w.preWriteCallback()
	}

	if len(w.objs) == 0 {
		w.objs = append(w.objs, &response_object_structs.ResponseObject{})
	}

	w.objs[0].Data = append(w.objs[0].Data, b...)

	return len(b), nil
}

func (w *MockWriter) Close() error {
	return nil
}

func (w *MockWriter) Flush() {}

func (w *MockWriter) Header() http.Header {
	return w.HTTPHeader
}

func (w *MockWriter) WriteHeader(int) {
	w.HTTPStatusCode = 200
}

func (w *MockWriter) GetObjects() []*response_object_structs.ResponseObject {
	return w.objs
}
func (w *MockWriter) SetPreWriteCallback(fn func()) {
	w.preWriteCallback = fn
}
