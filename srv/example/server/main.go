package main

import (
	"context"
	"fmt"
	"log"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/yuexclusive/utils/logger"
	"github.com/yuexclusive/utils/rpc/middleware/auth"
	"github.com/yuexclusive/utils/srv/example/proto/hello"
	"google.golang.org/grpc"

	"github.com/yuexclusive/utils/rpc"
)

type handler struct{}

func (h *handler) Send(ctx context.Context, req *hello.Request) (*hello.Response, error) {
	// panic("not implemented") // TODO: Implement
	fmt.Println("call from client")

	return &hello.Response{Res: fmt.Sprintf("hello %s", req.Name)}, nil
}

var l = logger.Single()

func main() {
	server, err := rpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_auth.StreamServerInterceptor(auth.AuthFunc),
			grpc_zap.StreamServerInterceptor(l),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_auth.UnaryServerInterceptor(auth.AuthFunc),
			grpc_zap.UnaryServerInterceptor(l),
			grpc_recovery.UnaryServerInterceptor(),
		)),
	)
	if err != nil {
		log.Fatal(err)
	}

	hello.RegisterHelloServer(server.Server, new(handler))

	if err := server.Serve(); err != nil {
		log.Fatal(err)
	}
}
