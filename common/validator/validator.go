package validator

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var Validator *validator.Validate

func InitValidator() {
	v := validator.New()
	Validator = v
}

func ValidateStruct(s interface{}) map[string]string {

	var errorMessages map[string]string
	// Validate the s (struct) using the validator
	if err := Validator.Struct(s); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errorMessage := map[string]string{}

		// Iterate through validation errors and build a map of error messages
		for _, e := range validationErrors {
			errorMessage[e.Field()] = fmt.Sprintf("%s is %s", e.Field(), e.Tag())
		}

		// Return validation error messages
		return errorMessage
	}
	return errorMessages
}
