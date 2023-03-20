package http

import (
	"net/http"
)

type HTTPResponseWriteCloser struct {
	Writer  http.ResponseWriter
	CloseCh chan bool
}

func NewHTTPResponseWriteCloser(w http.ResponseWriter) *HTTPResponseWriteCloser {
	return &HTTPResponseWriteCloser{Writer: w, CloseCh: make(chan bool)}
}

func (rwc *HTTPResponseWriteCloser) Write(p []byte) (n int, err error) {
	return rwc.Writer.Write(p)
}

func (rwc *HTTPResponseWriteCloser) Close() error {
	rwc.CloseCh <- true
	return nil
}

func (rwc *HTTPResponseWriteCloser) Header() http.Header {
	return rwc.Writer.Header()
}

func (rwc *HTTPResponseWriteCloser) WriteHeader(statusCode int) {
	rwc.Writer.WriteHeader(statusCode)
}

func (rwc *HTTPResponseWriteCloser) Flush() {
	rwc.Writer.(http.Flusher).Flush()
}
