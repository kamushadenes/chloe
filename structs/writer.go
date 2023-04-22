package structs

import (
	"net/http"
)

type ChloeWriter interface {
	WriteObject(*ResponseObject) error
	Write([]byte) (int, error)
	Close() error
	Flush()
	Header() http.Header
	WriteHeader(int)
}
