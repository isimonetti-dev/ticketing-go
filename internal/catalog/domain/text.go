package domain

import (
	"strings"
	"unicode/utf8"
)

func requireText(field, s string, maxLen int) (string, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return "", &ValidationError{Field: field, Err: ErrRequired}
	}
	if utf8.RuneCountInString(s) > maxLen {
		return "", &ValidationError{Field: field, Err: ErrTooLong}
	}
	return s, nil
}
