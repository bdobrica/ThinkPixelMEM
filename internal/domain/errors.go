package domain

import (
	"errors"
	"fmt"
)

// ErrorCode is a stable, machine-readable failure category.
type ErrorCode string

const (
	CodeInvalid      ErrorCode = "invalid"
	CodeNotFound     ErrorCode = "not_found"
	CodeConflict     ErrorCode = "conflict"
	CodeUnauthorized ErrorCode = "unauthorized"
	CodeForbidden    ErrorCode = "forbidden"
	CodeExpired      ErrorCode = "expired"
	CodeInternal     ErrorCode = "internal"
)

// Error carries a stable code while preserving the underlying cause. Message
// is safe context for callers; sensitive values must not be placed in it.
type Error struct {
	Code    ErrorCode
	Message string
	Err     error
}

func (e *Error) Error() string {
	if e == nil {
		return "<nil>"
	}
	if e.Message != "" {
		return e.Message
	}
	if e.Err != nil {
		return e.Err.Error()
	}
	return string(e.Code)
}

func (e *Error) Unwrap() error { return e.Err }

// NewError constructs a coded error and rejects unknown codes.
func NewError(code ErrorCode, message string, cause error) error {
	if !validErrorCode(code) {
		return fmt.Errorf("unknown error code %q", code)
	}
	return &Error{Code: code, Message: message, Err: cause}
}

// ErrorCodeOf returns the nearest coded error's stable category.
func ErrorCodeOf(err error) (ErrorCode, bool) {
	var target *Error
	if !errors.As(err, &target) {
		return "", false
	}
	return target.Code, true
}

func validErrorCode(code ErrorCode) bool {
	switch code {
	case CodeInvalid, CodeNotFound, CodeConflict, CodeUnauthorized, CodeForbidden, CodeExpired, CodeInternal:
		return true
	default:
		return false
	}
}
