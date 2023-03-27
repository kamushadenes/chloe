package http

import (
	"encoding/json"
	"net/http"
)

func parseFromRequest(r *http.Request, to interface{}) error {
	return json.NewDecoder(r.Body).Decode(to)
}
