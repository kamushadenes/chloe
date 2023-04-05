package errors

import errs "errors"

func Wrap(errl ...error) error {
	if errl[len(errl)-1] == nil {
		return nil
	}

	return errs.Join(errl...)
}

func Is(err, target error) bool {
	return errs.Is(err, target)
}
