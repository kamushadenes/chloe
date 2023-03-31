package errors

import (
	"errors"
)

var ErrActionFailed = errors.New("action failed")
var ErrCompletionFailed = errors.New("completion failed")
var ErrInvalidAction = errors.New("invalid action")
var ErrSavingMessage = errors.New("failed to save message")
var ErrForgettingUser = errors.New("failed to forget user")

func Wrap(errs ...error) error {
	return errors.Join(errs...)
}
