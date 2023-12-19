package utils

import (
	"srating/x/rest"

	"github.com/go-playground/validator/v10"
)

// validate validates the given planType using a new validator instance.
// If validation fails, it logs the error and returns it.
func Validate(data interface{}) error {
	validate := validator.New()
	if err := validate.Struct(data); err != nil {
		LogError(err, "Error validating input")
		return rest.ErrValidation
	}
	return nil
}
