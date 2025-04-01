package utils

import (
	"github.com/go-playground/validator/v10"
)

// Validator instance
var validate = validator.New()

// ValidationError holds field errors
type ValidationError struct {
	Field string `json:"field"`
	Error string `json:"error"`
}

// ValidateStruct validates a struct and returns a formatted error response
func ValidateStruct(data interface{}) []ValidationError {
	var errors []ValidationError
	err := validate.Struct(data)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, ValidationError{
				Field: err.Field(),
				Error: err.Tag(),
			})
		}
	}
	return errors
}
