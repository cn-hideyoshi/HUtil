package errorx

import (
	"errors"
	"fmt"
)

type Error struct {
	code    int
	message string
	err     error
}

func (e *Error) Error() string {
	if e.err == nil {
		return e.message
	}
	return fmt.Sprintf("%s: %v", e.message, e.err)
}

func (e *Error) Unwrap() error {
	return e.err
}

func (e *Error) Code() int {
	return e.code
}

func New(code int, msg string) *Error {
	return &Error{
		code:    code,
		message: msg,
	}
}

func Wrap(err error, code int, msg string) *Error {
	if err == nil {
		return nil
	}
	return &Error{
		code:    code,
		message: msg,
		err:     err,
	}
}

func Code(err error) int {
	if err == nil {
		return CodeOK
	}

	var e *Error
	if errors.As(err, &e) {
		return e.code
	}
	return CodeUnknown
}

func Is(err error, code int) bool {
	return Code(err) == code
}
