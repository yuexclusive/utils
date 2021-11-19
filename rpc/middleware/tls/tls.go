package tls

import (
	"crypto/tls"

	"github.com/yuexclusive/utils/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func Option() (grpc.ServerOption, error) {
	cfg := config.MustGet()
	cert, err := tls.LoadX509KeyPair(cfg.TLS.CertFile, cfg.TLS.KeyFile)
	if err != nil {
		return nil, err
	}

	return grpc.Creds(credentials.NewServerTLSFromCert(&cert)), nil
}
