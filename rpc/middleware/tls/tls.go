package tls

import (
	"crypto/tls"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func Option(certFile, keyFile string) (grpc.ServerOption, error) {
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return nil, err
	}

	return grpc.Creds(credentials.NewServerTLSFromCert(&cert)), nil
}
