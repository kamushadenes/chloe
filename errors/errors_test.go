package errors

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWrap(t *testing.T) {
	err := Wrap(ErrActionFailed, ErrCompletionFailed, ErrInvalidAction, ErrSavingMessage, ErrForgettingUser)

	assert.Error(t, err)

	assert.Equal(t, "action failed\ncompletion failed\ninvalid action\nfailed to save message\nfailed to forget user", err.Error())
}
