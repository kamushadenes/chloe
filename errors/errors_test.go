package errors

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWrap(t *testing.T) {
	err := Wrap(ErrCompletionFailed, ErrInvalidAction, ErrSaveMessage, ErrForgetUser)

	assert.Error(t, err)

	assert.Equal(t, "completion failed\ninvalid action\nfailed to save message\nfailed to forget user", err.Error())
}
