package http

import (
	"github.com/kamushadenes/chloe/structs"
	"net/http"
)

type HTTPWriter struct {
	Writer  http.ResponseWriter
	CloseCh chan bool
}

func NewHTTPResponseWriteCloser(w http.ResponseWriter) *HTTPWriter {
	return &HTTPWriter{Writer: w, CloseCh: make(chan bool)}
}

func (rwc *HTTPWriter) Write(p []byte) (n int, err error) {
	return rwc.Writer.Write(p)
}

func (rwc *HTTPWriter) WriteObject(obj *structs.ResponseObject) error {
	_, err := rwc.Write(obj.Data)

	return err
}

func (rwc *HTTPWriter) Close() error {
	rwc.CloseCh <- true
	return nil
}

func (rwc *HTTPWriter) Header() http.Header {
	return rwc.Writer.Header()
}

func (rwc *HTTPWriter) WriteHeader(statusCode int) {
	rwc.Writer.WriteHeader(statusCode)
}

func (rwc *HTTPWriter) Flush() {
	rwc.Writer.(http.Flusher).Flush()
}
