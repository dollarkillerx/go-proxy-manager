package utils

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func RPCAuth(token string, mt map[string]string) func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		if mt == nil {
			mt = map[string]string{}
		}
		mt["authorization"] = token
		md := metadata.New(mt)
		ctx = metadata.NewOutgoingContext(ctx, md)
		err := invoker(ctx, method, req, reply, cc, opts...)
		return err
	}
}

//func interceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
//	start := time.Now()
//	md := metadata.New(map[string]string{
//		"authorization": "1234",
//	})
//	ctx = metadata.NewOutgoingContext(ctx, md)
//	err := invoker(ctx, method, req, reply, cc, opts...)
//	fmt.Printf("耗时:%s\n", time.Since(start))
//	return err
//}
