package utils

import (
	"github.com/go-playground/validator/v10"
)

// NewValidator func for create a new validator for model fields.
func NewValidator() *validator.Validate {
	// Create a new validator for a Book model.
	validate := validator.New()

	// Custom validation for uuid.UUID fields.
	// _ = validate.RegisterValidation("uuid", func(fl validator.FieldLevel) bool {
	// 	field := fl.Field().String()
	// 	if _, err := uuid.Parse(field); err != nil {
	// 		return true
	// 	}
	// 	return false
	// })

	return validate
}

func ValidatorErrors(err error) map[string]string {
	fields := map[string]string{}

	for _, err := range err.(validator.ValidationErrors) {
		fields[err.Field()] = err.Error()
	}

	return fields
}

// func ValidatorError(message string) map[string]interface{} {
// 	return map[string]interface{}{
// 		"error":   true,
// 		"message": message,
// 	}
// }

// func ValidatorErrorStruct(err error) map[string]interface{} {

// 	fields := map[string]interface {
// 	}{
// 		"error":   true,
// 		"message": "Incorrect Parameter",
// 		"require": []string{},
// 	}
// 	// errors := []string{}

// 	for _, err := range err.(validator.ValidationErrors) {
// 		fields["require"] = append(fields["require"].([]string), err.Field())
// 		// errors = append(errors, err.Field())
// 	}

// 	// fields["fields"] = errors

// 	return fields
// }
