package errors

import (
	"fmt"
)

var ErrLoadMessages = fmt.Errorf("failed to load messages")
var ErrUpdateMessage = fmt.Errorf("failed to update message")
var ErrDeleteMessage = fmt.Errorf("failed to delete message")
var ErrSaveMessage = fmt.Errorf("failed to save message")
