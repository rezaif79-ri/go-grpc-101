package util

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

const arithmetic_operators_enum = "arithmetic_operators_enum"

type (
	ErrorResponse struct {
		Error       bool
		FailedField string
		Message     string
		Tag         string
		Value       interface{}
	}
	GoValidator struct {
		validator *validator.Validate
	}
)

var goValidatorErrDict = map[string]string{
	arithmetic_operators_enum: "Only +, -, /, * math operators allowed",
}

func getGoValidatorRegisteredErr(tag string) (msg string, found bool) {
	msg = ""
	found = false

	if errMsg, ok := goValidatorErrDict[tag]; ok {
		msg = errMsg
		found = true
	}
	return msg, found
}

func NewGoValidator() *GoValidator {
	var validate = validator.New()

	validate.RegisterValidation(arithmetic_operators_enum, arithmeticOperatorEnum)

	return &GoValidator{
		validator: validate,
	}
}

func (gv *GoValidator) ValidateStruct(data interface{}) []ErrorResponse {
	validationErrors := []ErrorResponse{}

	errs := gv.validator.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			var res ErrorResponse

			res.FailedField = err.Field() // Export struct field name
			res.Tag = err.Tag()           // Export struct tag
			res.Value = err.Value()       // Export field value
			res.Error = true

			if msg, found := getGoValidatorRegisteredErr(res.Tag); found {
				res.Message = fmt.Sprintf(
					"[%s]: '%v' | Error '%s' - %s",
					res.FailedField,
					res.Value,
					res.Tag,
					msg,
				)
			} else {
				res.Message = fmt.Sprintf(
					"[%s]: '%v' | Needs to implement '%s'",
					res.FailedField,
					res.Value,
					res.Tag,
				)
			}

			validationErrors = append(validationErrors, res)
		}
	}

	return validationErrors
}
