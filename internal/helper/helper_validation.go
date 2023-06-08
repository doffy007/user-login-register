package helper

import validation "github.com/go-ozzo/ozzo-validation"

type MessageErrorValidation struct {
	Field    string   `json:"field"`
	Messages []string `json:"messages"`
}

func FormatError(err error) []MessageErrorValidation {
	var errors []MessageErrorValidation

	if err == nil {
		return nil
	}

	switch ev := err.(type) {
	case validation.Errors:
		return formatOzzoValidationError(ev)
	default:
		return append(errors, MessageErrorValidation{
			Field:    "-",
			Messages: []string{err.Error()},
		})
	}
}

func formatOzzoValidationError(err validation.Errors) []MessageErrorValidation {
	var errors []MessageErrorValidation

	for k, e := range err {
		errors = append(errors, MessageErrorValidation{
			Field:    k,
			Messages: []string{e.Error()},
		})
	}

	return errors
}
