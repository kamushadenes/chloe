package cli

import (
	"bufio"
	"github.com/kamushadenes/chloe/structs"
	"os"
)

type CLIWriter struct {
	w *bufio.Writer
}

func NewCLIWriter() *CLIWriter {
	return &CLIWriter{
		w: bufio.NewWriter(os.Stdout),
	}
}

func (w *CLIWriter) Write(p []byte) (n int, err error) {
	return w.w.Write(p)
}

func (w *CLIWriter) WriteObject(obj *structs.ResponseObject) error {
	_, err := w.Write(obj.Data)

	return err
}

func (w *CLIWriter) Flush() {
	_ = w.w.Flush()
}

func (w *CLIWriter) Close() error {
	return nil
}
