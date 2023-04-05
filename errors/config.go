package errors

import (
	"fmt"
)

var ErrInvalidEnv = fmt.Errorf("invalid environment variable")
var ErrMissingEnv = fmt.Errorf("missing environment variable")
