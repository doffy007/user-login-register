package validations

import (
	"regexp"
	"strings"

	"github.com/go-ozzo/ozzo-validation/is"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

var (
	humanNameRegex = regexp.MustCompile(`^[a-zA-Z\'’.,\s]+$`)
)

func validateHumanName(v string) bool {
	if v == "" {
		return true
	}
	return humanNameRegex.MatchString(v)
}

func ValidateHumanName() validation.StringRule {
	return validation.NewStringRuleWithError(
		validateHumanName,
		validation.NewError("validation_name", "Invalid format. This field only allow these following characters: alphabet, single quote ('), space, comma(,), and period(.)"))
}

func ValidateOptionalEmail() validation.StringRule {
	return validation.NewStringRuleWithError(
		func(s string) bool {
			if strings.TrimSpace(s) == "" {
				return true
			}

			return is.Email.Validate(s) == nil
		},
		validation.NewError("validation_email", "invalid email format"))
}
