package cli

import (
	"bufio"
	"net/http"
	"os"

	"github.com/kamushadenes/chloe/structs/response_object_structs"
)

type CLIWriter struct {
	w        *bufio.Writer
	callback func()
}

func NewCLIWriter() *CLIWriter {
	return &CLIWriter{
		w: bufio.NewWriter(os.Stdout),
	}
}

func (w *CLIWriter) Write(p []byte) (n int, err error) {
	if w.callback != nil {
		w.callback()
	}
	return w.w.Write(p)
}

func (w *CLIWriter) WriteObject(obj *response_object_structs.ResponseObject) error {
	_, err := w.Write(obj.Data)

	return err
}

func (w *CLIWriter) Flush() {
	_ = w.w.Flush()
}

func (w *CLIWriter) Close() error {
	return nil
}

func (w *CLIWriter) WriteHeader(int)     {}
func (w *CLIWriter) Header() http.Header { return http.Header{} }
func (w *CLIWriter) SetPreWriteCallback(fn func()) {
	w.callback = fn
}

func (w *CLIWriter) GetObjects() []*response_object_structs.ResponseObject {
	return nil
}
