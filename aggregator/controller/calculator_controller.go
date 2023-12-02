package controller

import (
	"net/http"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/rezaif79-ri/go-grpc-101/aggregator/domain"
	"github.com/rezaif79-ri/go-grpc-101/aggregator/util"
)

type calculatorController struct {
	calculatorService domain.CalculatorService
	goValidator       *util.GoValidator
}

func NewCalculatorController(validator util.GoValidator, cs *domain.CalculatorService) domain.CalculatorController {
	return &calculatorController{
		calculatorService: *cs,
		goValidator:       &validator,
	}
}

// Calculate implements domain.CalculatorController.
func (cc *calculatorController) Calculate(c *fiber.Ctx) error {
	type CalculateRequest struct {
		NumberA  float32 `json:"number_a" validate:"required"`
		NumberB  float32 `json:"number_b" validate:"required"`
		Operator string  `json:"operator" validate:"required,arithmetic_operators_enum"`
	}
	var body CalculateRequest
	err := c.BodyParser(&body)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
			"data":    nil,
		})
	}
	if errs := cc.goValidator.ValidateStruct(body); len(errs) > 0 {
		return c.Status(http.StatusBadRequest).JSON(map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": "Validation error",
			"data":    errs,
		})
	}

	var result float32
	if body.Operator == "+" {
		res, sErr := cc.calculatorService.CalcAddition(body.NumberA, body.NumberB)
		result = res
		err = sErr
	} else if body.Operator == "-" {
		res, sErr := cc.calculatorService.CalcSubstraction(body.NumberA, body.NumberB)
		result = res
		err = sErr
	} else if body.Operator == "*" {
		res, sErr := cc.calculatorService.CalcMultiplication(body.NumberA, body.NumberB)
		result = res
		err = sErr
	} else if body.Operator == "/" {
		res, sErr := cc.calculatorService.CalcDivision(body.NumberA, body.NumberB)
		result = res
		err = sErr
	}

	if err != nil {
		return c.Status(http.StatusConflict).JSON(map[string]interface{}{
			"status":  http.StatusConflict,
			"message": err.Error(),
			"data":    nil,
		})
	}

	return c.Status(http.StatusOK).JSON(map[string]interface{}{
		"status":  http.StatusOK,
		"message": "OK",
		"data": map[string]interface{}{
			"result": result,
		},
	})
}
