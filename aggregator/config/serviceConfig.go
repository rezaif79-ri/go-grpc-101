package config

import (
	"google.golang.org/grpc"
)

func NewServiceGrpcConn() (*grpc.ClientConn, error) {
	//Create connection to gRPC server
	conn, err := grpc.Dial("localhost:5001", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return conn, err
}
