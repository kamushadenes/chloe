package cli

import (
	"io"
	"net/http"
	"os"

	"github.com/kamushadenes/chloe/structs/response_object_structs"
)

type FileWriter struct {
	Path             string
	f                io.WriteCloser
	preWriteCallback func()
}

func NewFileWriter(path string) *FileWriter {
	w := &FileWriter{
		Path: path,
	}

	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}

	w.f = f

	return w
}

func (w *FileWriter) Write(p []byte) (n int, err error) {
	if w.preWriteCallback != nil {
		w.preWriteCallback()
	}

	return w.f.Write(p)
}

func (w *FileWriter) WriteObject(obj *response_object_structs.ResponseObject) error {
	_, err := w.Write(obj.Data)

	return err
}

func (w *FileWriter) Flush() {}

func (w *FileWriter) Close() error {
	return w.f.Close()
}

func (w *FileWriter) WriteHeader(statusCode int) {}
func (w *FileWriter) Header() http.Header        { return http.Header{} }
func (w *FileWriter) SetPreWriteCallback(fn func()) {
	w.preWriteCallback = fn
}

func (w *FileWriter) GetObjects() []*response_object_structs.ResponseObject {
	return nil
}
