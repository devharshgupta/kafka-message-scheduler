package validator

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
)

var Validator *validator.Validate

func InitValidator() {
	v := validator.New()
	v.RegisterValidation("timestamp", CustomValidation)
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

// CustomValidation is a function that validates the timestamp format.
func CustomValidation(fl validator.FieldLevel) bool {
	scheduledAt := fl.Field().String()

	// Attempt to parse the string into a time.Time
	_, err := time.Parse(time.RFC3339, scheduledAt)
	return err == nil
}
