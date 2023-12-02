package util

import "github.com/go-playground/validator/v10"

type (
	ErrorResponse struct {
		Error       bool
		FailedField string
		Tag         string
		Value       interface{}
	}
	GoValidator struct {
		validator *validator.Validate
	}
)

func NewGoValidator() *GoValidator {
	var validate = validator.New()
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

			validationErrors = append(validationErrors, res)
		}
	}

	return validationErrors
}
