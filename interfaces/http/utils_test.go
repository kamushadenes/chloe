package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_parseFromRequest(t *testing.T) {
	type testStruct struct {
		Field string `json:"field"`
	}

	testCases := []struct {
		name     string
		body     interface{}
		expected testStruct
		err      error
	}{
		{
			name: "valid request",
			body: testStruct{Field: "value"},
			expected: testStruct{
				Field: "value",
			},
			err: nil,
		},
		{
			name:     "invalid request",
			body:     "",
			expected: testStruct{},
			err:      errors.New(""),
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			reqBody, _ := json.Marshal(tt.body)
			req, err := http.NewRequest(http.MethodPost, "/", bytes.NewReader(reqBody))
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			w := httptest.NewRecorder()
			_, _ = w.Write(reqBody)

			err = parseFromRequest(req, &tt.expected)
			if (err == nil) != (tt.err == nil) {
				t.Fatalf("unexpected error: got %v, want %v", err, tt.err)
			}

			if err == nil {
				resBody, err := io.ReadAll(w.Body)
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}

				var got testStruct
				if err := json.Unmarshal(resBody, &got); err != nil {
					t.Fatalf("unexpected error: %v", err)
				}

				if got != tt.expected {
					t.Fatalf("unexpected response: got %v, want %v", got, tt.expected)
				}
			}

		})
	}
}
