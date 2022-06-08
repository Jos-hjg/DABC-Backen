package models

import (
	"gopkg.in/go-playground/validator.v9"
)

type ErrorType map[string]map[string]string

func ValidateError(errs validator.ValidationErrors, errType ErrorType) map[string]string {
	errMap := map[string]string{}
	for _, err := range errs {
		errMap[err.Field()] = errType[err.Field()][err.Tag()]
	}
	return errMap
}

