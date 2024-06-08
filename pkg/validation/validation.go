package validation

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/nurularifin27/go-util/pkg/util"
)

type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func ValidateRequest(err error) interface{} {
	var ve validator.ValidationErrors
	var out []ErrorMsg
	if errors.As(err, &ve) {
		out = make([]ErrorMsg, 0)
		for _, fe := range ve {
			out = append(out, ErrorMsg{Field: util.ConvertToSnakeCase(fe.Field()), Message: getErrorMsg(fe)})
		}
		return out
	}
	return err.Error()
}

func getErrorMsg(fe validator.FieldError) string {
	msg := "Unknown error"
	switch fe.Tag() {
	case "required":
		msg = "This field is required"
	case "email":
		msg = "This field is not a valid email"
	case "lte":
		msg = fmt.Sprintf("Should be less than %s or equal", fe.Param())
	case "gte":
		msg = fmt.Sprintf("Should be greater than %s or equal", fe.Param())
	case "lt":
		msg = fmt.Sprintf("Should be less than %s", fe.Param())
	case "gt":
		msg = fmt.Sprintf("Should be greater than %s", fe.Param())
	case "min":
		msg = fmt.Sprintf("Must be at least %s characters long", fe.Param())
	case "max":
		msg = fmt.Sprintf("cannot be longer than %s characters", fe.Param())
	case "eqfield":
		msg = fmt.Sprintf("This field is must be equal to %s", util.ConvertToSnakeCase(fe.Param()))
	}

	return msg
}
