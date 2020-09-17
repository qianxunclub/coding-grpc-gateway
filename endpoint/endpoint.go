package endpoint

import (
	gw "coding-grpc-gateway/grpc"
	"context"
	"flag"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

var (
	userEndpoint = flag.String("user_endpoint", "dev.coding.oa.com:20153", "endpoint of user_endpoint gRPC Service")
)

func RegisterEndpoint(ctx context.Context, mux *runtime.ServeMux, opts []grpc.DialOption) (err error) {
	error := gw.RegisterUserServiceHandlerFromEndpoint(ctx, mux, *userEndpoint, opts)
	if error != nil {
		return error
	}
	return nil
}
