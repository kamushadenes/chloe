package writer_structs

import (
	"net/http"

	"github.com/kamushadenes/chloe/structs/response_object_structs"
)

type ChloeWriter interface {
	WriteObject(*response_object_structs.ResponseObject) error
	Write([]byte) (int, error)
	Close() error
	Flush()
	Header() http.Header
	WriteHeader(int)
	SetPreWriteCallback(func())
	GetObjects() []*response_object_structs.ResponseObject
}
