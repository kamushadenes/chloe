package errors

import (
	"fmt"
)

var ErrOpenDatabase = fmt.Errorf("failed to open database")
var ErrMigrateDatabase = fmt.Errorf("failed to migrate database")
