package domain

import "github.com/gofiber/fiber/v2"

type CalculatorController interface {
	Calculate(c *fiber.Ctx) error
}

type CalculatorService interface {
	CalcAddition(a float32, b float32) (res float32, err error)
	CalcSubstraction(a float32, b float32) (res float32, err error)
	CalcMultiplication(a float32, b float32) (res float32, err error)
	CalcDivision(a float32, b float32) (res float32, err error)
}
