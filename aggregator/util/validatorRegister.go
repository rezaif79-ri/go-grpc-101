package util

import "github.com/go-playground/validator/v10"

func arithmeticOperatorEnum(fl validator.FieldLevel) bool {
	return fl.Field().String() == "+" ||
		fl.Field().String() == "-" ||
		fl.Field().String() == "/" ||
		fl.Field().String() == "*"
}
