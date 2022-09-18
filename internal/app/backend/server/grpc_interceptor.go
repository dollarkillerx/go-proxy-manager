package server

import (
	"context"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var (
	token = ""
)

func init() {
	token = os.Getenv("Token")
	if token == "" {
		token = "f32e4a2f-0aec-4e5b-b4ab-e88ab8a68353"
	}
}

func grpcAuth(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return resp, status.Error(codes.Unauthenticated, "authorization不存在")
	}

	var authorization string
	if val, ok := md["authorization"]; ok {
		authorization = val[0]
	}

	if token != authorization {
		return resp, status.Error(codes.Unauthenticated, "authorization不正确")
	}

	res, err := handler(ctx, req)
	return res, err
}
