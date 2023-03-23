package utils

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWriteHeader(t *testing.T) {
	key := "key1"
	value := "value1"
	rr := httptest.NewRecorder()
	WriteHeader(rr, key, value)
	h := rr.Header()
	got := h.Get(key)
	if got != value {
		t.Errorf("WriteHeader() got %q, want %q", got, value)
	}
}

func TestWriteStatusCode(t *testing.T) {
	statusCode := http.StatusOK
	rr := httptest.NewRecorder()
	WriteStatusCode(rr, statusCode)
	got := rr.Code
	if got != statusCode {
		t.Errorf("WriteStatusCode() got %d, want %d", got, statusCode)
	}
}

func TestFlush(t *testing.T) {
	rr := httptest.NewRecorder()
	Flush(rr)
	if !rr.Flushed {
		t.Errorf("Flush() did not Flush the response")
	}
}
