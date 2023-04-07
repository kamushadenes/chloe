package cli

import (
	"github.com/kamushadenes/chloe/structs"
	"io"
	"net/http"
	"os"
)

type FileWriter struct {
	Path string
	f    io.WriteCloser
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
	return w.f.Write(p)
}

func (w *FileWriter) WriteObject(obj *structs.ResponseObject) error {
	_, err := w.Write(obj.Data)

	return err
}

func (w *FileWriter) Flush() {}

func (w *FileWriter) Close() error {
	return w.f.Close()
}

func (w *FileWriter) WriteHeader(statusCode int) {}
func (w *FileWriter) Header() http.Header        { return http.Header{} }
