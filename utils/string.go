package utils

import (
	"math/rand"
	"strings"
)

func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func PickRandomString(strs ...string) string {
	return strs[rand.Intn(len(strs))]
}

func Truncate(s string, n int) string {
	s = strings.Join(strings.Fields(s), " ")

	if len(s) > n {
		return s[:n]
	}
	return s
}
