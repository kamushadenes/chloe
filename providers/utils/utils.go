package utils

import (
	"io"
	"net/http"
)

func CloneHeader(resp *http.Response, w io.Writer, key string) {
	WriteHeader(w, key, resp.Header.Get(key))
}

func WriteHeader(w io.Writer, key string, value string) {
	switch ww := w.(type) {
	case http.ResponseWriter:
		ww.Header().Set(key, value)
	}
}

func WriteStatusCode(w io.Writer, statusCode int) {
	switch ww := w.(type) {
	case http.ResponseWriter:
		ww.WriteHeader(statusCode)
	}
}

func Flush(w io.Writer) {
	switch ww := w.(type) {
	case http.Flusher:
		ww.Flush()
	}
}
