package errors

import (
	"fmt"
)

var ErrCreateAPIKey = fmt.Errorf("failed to create API key")
var ErrGetAPIKey = fmt.Errorf("failed to get API key")
