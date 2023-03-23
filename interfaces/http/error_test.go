package http

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/render"
	"github.com/stretchr/testify/assert"
)

func TestErrInvalidRequest(t *testing.T) {
	err := errors.New("invalid request")
	rr := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/invalid-request", nil)

	render.Render(rr, req, ErrInvalidRequest(err))

	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Equal(t, "application/json; charset=utf-8", rr.Header().Get("Content-Type"))
	// Test response body, which should contain JSON with error data:
	// {"status":"Invalid request.","error":"invalid request"}
}

func TestErrRender(t *testing.T) {
	err := errors.New("error rendering response")
	rr := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/error-rendering", nil)

	render.Render(rr, req, ErrRender(err))

	assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
	assert.Equal(t, "application/json; charset=utf-8", rr.Header().Get("Content-Type"))
	// Test response body, which should contain JSON with error data:
	// {"status":"Error rendering response.","error":"error rendering response"}
}

func TestErrNotFound_Render(t *testing.T) {
	rr := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/not-found", nil)

	render.Render(rr, req, ErrNotFound)

	assert.Equal(t, http.StatusNotFound, rr.Code)
	assert.Equal(t, "application/json; charset=utf-8", rr.Header().Get("Content-Type"))
	// Test response body, which should contain JSON with error data:
	// {"status":"Resource not found."}
}
