package react

type BytesWriter struct {
	Bytes []byte
}

func (w *BytesWriter) Write(p []byte) (n int, err error) {
	w.Bytes = append(w.Bytes, p...)
	return len(p), nil
}

func (w *BytesWriter) Close() error {
	return nil
}
