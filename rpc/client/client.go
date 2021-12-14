package client

import (
	"fmt"
	"io"

	"github.com/yuexclusive/utils/log"

	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/yuexclusive/utils/config"
	"github.com/yuexclusive/utils/registry"
	"github.com/yuexclusive/utils/rpc/middleware/trace"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"
)

type Config struct {
	config.Config `mapstructure:"config"`
}

func Dial(name string, token string, dialOptions ...grpc.DialOption) (io.Closer, *grpc.ClientConn, error) {
	var opts []grpc.DialOption
	if token != "" {
		perRPC := oauth.NewOauthAccess(fetchToken(token))
		opts = append(opts, grpc.WithPerRPCCredentials(perRPC))
	}

	creds, err := credentials.NewClientTLSFromFile(cfg.TLS.CACertFile, cfg.TLS.ServerNameOverride)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to load credentials: %v", err)
	}
	opts = append(opts, grpc.WithTransportCredentials(creds))

	tracer, closer, err := trace.Tracer()

	if err != nil {
		log.Fatal(err.Error())
	}

	// defer closer.Close()

	// opts = append(opts, grpc.WithUnaryInterceptor(grpc_prometheus.UnaryClientInterceptor))

	opts = append(opts, grpc.WithUnaryInterceptor(grpc_opentracing.UnaryClientInterceptor(
		grpc_opentracing.WithTracer(tracer),
	)))

	opts = append(opts, grpc.WithBlock())

	opts = append(opts, dialOptions...)

	dis := registry.NewDiscovery(cfg.ETCD.Address, name)

	address, err := dis.Get(name)
	if err != nil {
		return nil, nil, err
	}

	conn, err := grpc.Dial(address, opts...)
	return closer, conn, err
}

func fetchToken(token string) *oauth2.Token {
	return &oauth2.Token{
		AccessToken: token,
	}
}
