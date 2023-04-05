package errors

import errs "errors"

func Wrap(errl ...error) error {
	return errs.Join(errl...)
}
