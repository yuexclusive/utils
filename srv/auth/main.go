package main

import (
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/yuexclusive/utils/config"
	"github.com/yuexclusive/utils/db"
	"github.com/yuexclusive/utils/logger"
	"github.com/yuexclusive/utils/rpc"
	"github.com/yuexclusive/utils/srv/auth/handler"
	"github.com/yuexclusive/utils/srv/auth/proto/auth"
	"google.golang.org/grpc"

	_ "github.com/yuexclusive/utils/db/postgres"
)

var l = logger.Single()

var sugar = l.Sugar()

func main() {
	// tracer, closer, err := trace.Tracer()

	// if err != nil {
	// 	sugar.Fatal(err)
	// }

	// defer closer.Close()

	err := db.InitConnection("test", config.MustGet().ConnStr, db.DialectPostgres, db.ConnectionOptionLogMode(true), db.ConnectionOptionSingularTable(true))

	if err != nil {
		sugar.Fatalf("connect db failed: ", err)
	}

	s, err := rpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_zap.UnaryServerInterceptor(l),
			// grpc_opentracing.UnaryServerInterceptor(grpc_opentracing.WithTracer(tracer)),
		)),
	)

	if err != nil {
		sugar.Fatal(err)
	}

	auth.RegisterAuthServer(s.Server, new(handler.Handler))

	sugar.Fatal(s.Serve())
}
