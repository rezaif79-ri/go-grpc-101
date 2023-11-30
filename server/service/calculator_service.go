package service

import (
	"context"

	"github.com/rezaif79-ri/go-grpc-101/model/calculator"
)

type CalculatorService struct {
	calculator.UnimplementedCalculatorServiceServer
}

func NewCalculatorServices() calculator.CalculatorServiceServer {
	return &CalculatorService{
		UnimplementedCalculatorServiceServer: calculator.UnimplementedCalculatorServiceServer{},
	}
}

func (cs *CalculatorService) Addition(ctx context.Context, req *calculator.AdditionRequest) (*calculator.AdditionResponse, error) {
	res := req.NumberA + req.NumberB
	return &calculator.AdditionResponse{Result: res}, nil
}
func (cs *CalculatorService) Substraction(ctx context.Context, req *calculator.SubstractionRequest) (*calculator.SubstractionResponse, error) {
	res := req.NumberA - req.NumberB
	return &calculator.SubstractionResponse{Result: res}, nil
}
func (cs *CalculatorService) Multiplication(ctx context.Context, req *calculator.MultiplicationRequest) (*calculator.MultiplicationResponse, error) {
	res := req.NumberA * req.NumberB
	return &calculator.MultiplicationResponse{Result: res}, nil
}
func (cs *CalculatorService) Division(ctx context.Context, req *calculator.DivisionRequest) (*calculator.DivisionResponse, error) {
	res := req.NumberA / req.NumberB
	return &calculator.DivisionResponse{Result: res}, nil
}
