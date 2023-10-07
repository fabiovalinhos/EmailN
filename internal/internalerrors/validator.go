package internalerrors

import (
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidateStruct(obj interface{}) error {
	validate := validator.New()
	err := validate.Struct(obj)
	if err == nil {
		return nil
	}
	validationErrors := err.(validator.ValidationErrors)
	validationError := validationErrors[0]

	fields := strings.ToLower(validationError.StructField())
	switch validationError.Tag() {
	case "required":
		return errors.New(fields + " is required")
	case "max":
		return errors.New(fields + " is required with max " + validationError.Param())
	case "min":
		return errors.New(fields + " is required with min " + validationError.Param())
	case "email":
		return errors.New(fields + " is invalid")
	}

	return nil
}
