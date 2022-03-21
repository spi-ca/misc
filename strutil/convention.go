package strutil

import (
	"regexp"
	"strings"
)

var (
	camelCase = regexp.MustCompile("(^[^A-Z0-9]*|[A-Z0-9]*)([A-Z0-9][^A-Z]+|$)")
)

// CamelToUnderscore method converts a naming convention from CamelCase to under_score.
func CamelToUnderscore(s string) string {
	var a []string
	for _, sub := range camelCase.FindAllStringSubmatch(s, -1) {
		if sub[1] != "" {
			a = append(a, sub[1])
		}
		if sub[2] != "" {
			a = append(a, sub[2])
		}
	}
	return strings.ToLower(strings.Join(a, "_"))
}
