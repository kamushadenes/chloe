package http

import (
	"encoding/json"
	"io"
	"net/http"
)

func parseFromRequest(r *http.Request, to interface{}) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(body, &to); err != nil {
		return err
	}

	return nil
}
