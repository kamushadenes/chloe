package utils

import (
	"bytes"
	"encoding/json"
	"regexp"
)

var jsonPtn = regexp.MustCompile(`\{(?:[^{}]|(\{(?:[^{}]|())*}))*}`)

func FindJSON(s string) string {
	if s == "" {
		return ""
	}

	var m map[string]interface{}
	var buf bytes.Buffer

	err := json.Unmarshal([]byte(s), &m)
	if err == nil {
		err = json.Compact(&buf, []byte(s))
		if err == nil {
			return buf.String()
		}
	}

	matches := jsonPtn.FindAllString(s, -1)
	if len(matches) == 0 {
		return ""
	}

	err = json.Unmarshal([]byte(matches[0]), &m)
	if err == nil {
		err = json.Compact(&buf, []byte(matches[0]))
		if err == nil {
			return buf.String()
		}
	}

	return ""
}
