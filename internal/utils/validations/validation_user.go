package validations

import (
	"github.com/doffy007/user-login-register/internal/request"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func ValidateUserRegisterParams(param request.ParamUserRegister) error {
	rules := []*validation.FieldRules{
		// Username
		validation.Field(&param.Username, validation.Required, validation.Length(2, 99), ValidateHumanName()),

		// Email
		validation.Field(&param.Email, validation.Length(3, 50), ValidateOptionalEmail()),

		// Password
		validation.Field(&param.Password, validation.Required, validation.Length(8, 0)),

		// Confirm Password
		validation.Field(&param.ConfirmPassword, validation.When(param.ConfirmPassword != param.Password, validation.Empty.Error("password and confirm password doesn't match"))),
	}

	err := validation.ValidateStruct(&param, rules...)
	ve, ok := err.(validation.Errors)
	if !ok {
		ve = make(validation.Errors)
	}

	if len(ve) == 0 {
		return nil
	}

	return ve
}
