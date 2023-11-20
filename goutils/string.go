package goutils

import (
	"sort"
	"strings"
	"unicode"
)

// TrimSpaceAndQuote returns a string with all leading and trailing spaces and quotes removed
func TrimSpaceAndQuote(s string) string {
	return strings.TrimFunc(s, func(r rune) bool {
		return unicode.IsSpace(r) || r == '"' || r == '\''
	})
}

// HidePassword returns a string that hides the password as a *
func HidePassword(s string, password ...string) string {
	sort.Slice(password, func(i, j int) bool {
		return len(password[i]) > len(password[j])
	})
	for _, v := range password {
		if len(v) > 0 {
			s = strings.ReplaceAll(s,
				v,
				strings.Repeat("*", len(v)),
			)
		}
	}
	return s
}
