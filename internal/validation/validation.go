package validation

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

// CompanyEmailValidator validates that the email contains "@company.co.id"
func CompanyEmailValidator(fl validator.FieldLevel) bool {
	email := fl.Field().String()
	return strings.Contains(email, "@company.co.id")
}
