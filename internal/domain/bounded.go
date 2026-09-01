package domain

import (
	"fmt"
	"unicode/utf8"
)

// BoundedString is a validated UTF-8 string whose length is measured in
// Unicode code points, matching user-visible API string limits.
type BoundedString struct{ value string }

func NewBoundedString(value string, minRunes, maxRunes int) (BoundedString, error) {
	if minRunes < 0 || maxRunes < minRunes {
		return BoundedString{}, NewError(CodeInternal, "invalid string bounds", nil)
	}
	if !utf8.ValidString(value) {
		return BoundedString{}, NewError(CodeInvalid, "string must be valid UTF-8", nil)
	}
	n := utf8.RuneCountInString(value)
	if n < minRunes || n > maxRunes {
		return BoundedString{}, NewError(CodeInvalid, fmt.Sprintf("string length must be between %d and %d characters", minRunes, maxRunes), nil)
	}
	return BoundedString{value: value}, nil
}

func (s BoundedString) String() string               { return s.value }
func (s BoundedString) MarshalText() ([]byte, error) { return []byte(s.value), nil }
