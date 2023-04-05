package errors

import (
	"fmt"
)

var ErrGetUser = fmt.Errorf("failed to get user")
var ErrUserNotFound = fmt.Errorf("user not found")
var ErrCreateUser = fmt.Errorf("failed to create user")
var ErrDeleteUser = fmt.Errorf("failed to delete user")
var ErrUpdateUser = fmt.Errorf("failed to update user")
var ErrSaveUser = fmt.Errorf("failed to save user")
var ErrForgetUser = fmt.Errorf("failed to forget user")
var ErrMergeUsers = fmt.Errorf("failed to merge users")
