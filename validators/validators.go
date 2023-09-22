package validators

import (
	"github.com/go-playground/validator/v10"
	"strings"
)

func ValidateCoolTitle(field validator.FieldLevel) bool {
	s := field.Field().String()
	return strings.Contains(strings.ToLower(s), "cool")
}
