package errors

import (
	"fmt"
)

var ErrProceed = fmt.Errorf("proceed")

var ErrNotImplemented = fmt.Errorf("not implemented")

var ErrMock = fmt.Errorf("mock error")
