package utils

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func ParseValidation(errs error) interface{} {
	var validationError validation.Errors
	if !errors.As(errs, &validationError) {
		return errs
	}
	var res = make(map[string]string)
	validationError = errs.(validation.Errors)
	for key, valueError := range validationError {
		res[key] = valueError.Error()
	}
	return res
}
