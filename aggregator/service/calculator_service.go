package service

import (
	"context"
	"time"

	"github.com/rezaif79-ri/go-grpc-101/aggregator/domain"
	"github.com/rezaif79-ri/go-grpc-101/aggregator/model/calculator"
)

type CalculatorService struct {
	Client calculator.CalculatorServiceClient
}

func NewCalculatorService(calculatorClient calculator.CalculatorServiceClient) domain.CalculatorService {
	return &CalculatorService{
		Client: calculatorClient,
	}
}

// CalcAddition implements domain.CalculatorService.
func (cs *CalculatorService) CalcAddition(a float32, b float32) (float32, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := cs.Client.Addition(ctx, &calculator.AdditionRequest{
		NumberA: a,
		NumberB: b,
	})
	if err != nil {
		return res.GetResult(), err
	}
	return res.GetResult(), nil
}

// CalcDivision implements domain.CalculatorService.
func (cs *CalculatorService) CalcDivision(a float32, b float32) (float32, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := cs.Client.Division(ctx, &calculator.DivisionRequest{
		NumberA: a,
		NumberB: b,
	})
	if err != nil {
		return res.GetResult(), err
	}
	return res.GetResult(), nil
}

// CalcMultiplication implements domain.CalculatorService.
func (cs *CalculatorService) CalcMultiplication(a float32, b float32) (float32, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := cs.Client.Multiplication(ctx, &calculator.MultiplicationRequest{
		NumberA: a,
		NumberB: b,
	})
	if err != nil {
		return res.GetResult(), err
	}
	return res.GetResult(), nil
}

// CalcSubstraction implements domain.CalculatorService.
func (cs *CalculatorService) CalcSubstraction(a float32, b float32) (float32, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := cs.Client.Substraction(ctx, &calculator.SubstractionRequest{
		NumberA: a,
		NumberB: b,
	})
	if err != nil {
		return res.GetResult(), err
	}
	return res.GetResult(), nil
}
