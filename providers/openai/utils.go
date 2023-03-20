package openai

import (
	"io"
	"net/http"
)

func cloneHeader(resp *http.Response, w io.Writer, key string) {
	writeHeader(w, key, resp.Header.Get(key))
}

func writeHeader(w io.Writer, key string, value string) {
	switch ww := w.(type) {
	case http.ResponseWriter:
		ww.Header().Set(key, value)
	}
}

func writeStatusCode(w io.Writer, statusCode int) {
	switch ww := w.(type) {
	case http.ResponseWriter:
		ww.WriteHeader(statusCode)
	}
}

func flush(w io.Writer) {
	switch ww := w.(type) {
	case http.Flusher:
		ww.Flush()
	}
}
