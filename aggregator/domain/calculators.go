package domain

import "github.com/gofiber/fiber/v2"

type CalculatorController interface {
	Calculate(c *fiber.Ctx) error
}

type CalculatorService interface {
	CalcAddition(a float64, b float64) (res float64, err error)
	CalcSubstraction(a float64, b float64) (res float64, err error)
	CalcMultiplication(a float64, b float64) (res float64, err error)
	CalcDivision(a float64, b float64) (res float64, err error)
}
