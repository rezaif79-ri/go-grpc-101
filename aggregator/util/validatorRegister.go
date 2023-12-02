package util

import "github.com/go-playground/validator/v10"

func arithmaticOperatorEnum(fl validator.FieldLevel) bool {
	return fl.Field().String() == "+" ||
		fl.Field().String() == "-" ||
		fl.Field().String() == "/" ||
		fl.Field().String() == "*"
}
