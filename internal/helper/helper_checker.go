package helper

import (
	"strings"

	"github.com/go-ozzo/ozzo-validation/v4/is"
)

func CheckEmail(s string) bool {
	if strings.TrimSpace(s) == "" {
		return true
	}

	return is.Email.Validate(s) == nil
}
