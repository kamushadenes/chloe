package utils

import (
	"bytes"
	"encoding/json"
	"regexp"
)

var jsonPtn = regexp.MustCompile(`\{(?:[^{}]|(\{(?:[^{}]|())*}))*}`)

func ExtractJSON(s string) string {
	if s == "" {
		return ""
	}

	var m map[string]interface{}
	var buf bytes.Buffer

	if err := json.Unmarshal([]byte(s), &m); err == nil {
		if err := json.Compact(&buf, []byte(s)); err == nil {
			return buf.String()
		}
	}

	matches := jsonPtn.FindAllString(s, -1)
	if len(matches) == 0 {
		return ""
	}

	if err := json.Unmarshal([]byte(matches[0]), &m); err == nil {
		if err := json.Compact(&buf, []byte(matches[0])); err == nil {
			return buf.String()
		}
	}

	return ""
}
