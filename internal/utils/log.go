package utils

import "strings"

func SafeJSON(s string) string {
	return strings.ReplaceAll(s, "\"", "'")
}
