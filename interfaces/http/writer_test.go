package http

import (
	"github.com/kamushadenes/chloe/react/utils"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestNewHTTPResponseWriteCloser(t *testing.T) {
	w := httptest.NewRecorder()
	rwc := utils.NewHTTPResponseWriteCloser(w)

	if rwc == nil {
		t.Error("failed to initialize HTTPResponseWriteCloser")
	}
}

func TestHTTPResponseWriteCloserWrite(t *testing.T) {
	w := httptest.NewRecorder()
	rwc := utils.NewHTTPResponseWriteCloser(w)

	n, err := rwc.Write([]byte("Hello, World!"))

	if err != nil {
		t.Errorf("error while writing to Writer: %v", err)
	}

	if n != len("Hello, World!") {
		t.Errorf("invalid number of bytes written: expected %d, got %d", len("Hello, World!"), n)
	}

	if w.Body.String() != "Hello, World!" {
		t.Errorf("invalid response body: expected %s, got %s", "Hello, World!", w.Body.String())
	}
}

func TestHTTPResponseWriteCloserClose(t *testing.T) {
	w := httptest.NewRecorder()
	rwc := utils.NewHTTPResponseWriteCloser(w)

	ch := make(chan bool)

	rwc.CloseCh = ch

	go rwc.Close()

	time.Sleep(100 * time.Millisecond)

	select {
	case <-rwc.CloseCh:
	default:
		t.Errorf("CloseCh not closed")
	}
}

func TestHTTPResponseWriteCloserHeader(t *testing.T) {
	w := httptest.NewRecorder()
	rwc := utils.NewHTTPResponseWriteCloser(w)

	rwc.Header().Set("Content-Type", "application/json")

	if w.Header().Get("Content-Type") != "application/json" {
		t.Errorf("invalid response header: expected %s, got %s", "application/json", w.Header().Get("Content-Type"))
	}
}

func TestHTTPResponseWriteCloserWriteHeader(t *testing.T) {
	w := httptest.NewRecorder()
	rwc := utils.NewHTTPResponseWriteCloser(w)

	rwc.WriteHeader(http.StatusOK)

	if w.Code != http.StatusOK {
		t.Errorf("invalid response status code: expected %d, got %d", http.StatusOK, w.Code)
	}
}

func TestHTTPResponseWriteCloserFlush(t *testing.T) {
	w := httptest.NewRecorder()
	rwc := utils.NewHTTPResponseWriteCloser(w)

	rwc.Flush()
}
