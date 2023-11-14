package goutils

import (
	"reflect"
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

// ToString returns a string converted directly by type without serialization.
// If it cannot be converted, a string of zero value is returned.
func ToString(v any) string {
	val := reflect.ValueOf(v)
	for val.Kind() == reflect.Ptr {
		if val.IsNil() {
			return ""
		}
		val = val.Elem()
	}
	switch s := val.Interface().(type) {
	case []byte:
		return BytesToString(s)
	case string:
		return s
	}
	return ""
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
