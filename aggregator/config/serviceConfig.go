package config

import (
	"github.com/rezaif79-ri/go-grpc-101/aggregator/util"
	"google.golang.org/grpc"
)

func NewServiceGrpcConn() (*grpc.ClientConn, error) {
	//Create connection to gRPC server
	host := util.GetEnv("GRPC_SERVICE_HOST", "localhost:5001")
	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return conn, err
}
