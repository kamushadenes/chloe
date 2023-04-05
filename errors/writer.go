package errors

import (
	"fmt"
)

var ErrSendMessage = fmt.Errorf("failed to send message")
var ErrCloseWriter = fmt.Errorf("failed to close writer")
