package main

import (
	"net/http"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/yuexclusive/utils/rpc"
	"github.com/yuexclusive/utils/rpc/middleware/auth"
	"google.golang.org/grpc"

	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/yuexclusive/utils/rpc/middleware/trace"
	"github.com/yuexclusive/utils/srv/basic/handler/role_handler"
	"github.com/yuexclusive/utils/srv/basic/handler/user_handler"
	"github.com/yuexclusive/utils/srv/basic/proto/role"
	"github.com/yuexclusive/utils/srv/basic/proto/user"
)

func main() {
	tracer, closer, err := trace.Tracer()

	if err != nil {
		sugar.Fatal(err)
	}

	defer closer.Close()

	s, err := rpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_auth.StreamServerInterceptor(auth.AuthFunc),
			grpc_zap.StreamServerInterceptor(l),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_auth.UnaryServerInterceptor(auth.AuthFunc),
			grpc_zap.UnaryServerInterceptor(l),
			grpc_recovery.UnaryServerInterceptor(),
			grpc_opentracing.UnaryServerInterceptor(grpc_opentracing.WithTracer(tracer)),
			grpc_prometheus.UnaryServerInterceptor,
		)),
	)

	if err != nil {
		sugar.Fatal(err)
	}

	role.RegisterRoleServer(s.Server, new(role_handler.Handler))
	user.RegisterUserServer(s.Server, new(user_handler.Handler))

	grpc_prometheus.Register(s.Server)

	go func() {
		http.Handle("/metrics", promhttp.Handler())
	}()

	sugar.Fatal(s.Serve())
}
