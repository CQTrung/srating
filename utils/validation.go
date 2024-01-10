package utils

import (
	"fmt"
	"strings"
	"unicode"

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

func ValidatePhone(phone string) (string, bool) {
	allowedRunes := func(r rune) rune {
		if unicode.IsNumber(r) || r == '+' {
			return r
		}
		if r == '.' {
			return ' '
		}
		return -1
	}
	sanitize := strings.Map(allowedRunes, phone)
	sanitize = strings.TrimSpace(sanitize)
	sanitize = strings.ReplaceAll(sanitize, " ", "")
	if sanitize[0] != '+' && sanitize[0] != '0' {
		return "", false
	}
	if len(sanitize) != 10 && len(sanitize) != 12 {
		return "", false
	}
	fmt.Println(sanitize)
	return sanitize, true
}
