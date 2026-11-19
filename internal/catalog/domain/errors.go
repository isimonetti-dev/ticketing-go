package domain

import "errors"

var (
	ErrRequired = errors.New("required")
	ErrTooLong  = errors.New("too long")
)

type ValidationError struct {
	Field string
	Err   error
}

func (e *ValidationError) Error() string {
	return e.Field + ": " + e.Err.Error()
}

func (e *ValidationError) Unwrap() error {
	return e.Err
}
