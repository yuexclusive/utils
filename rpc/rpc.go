package rpc

import (
	"net"

	"github.com/yuexclusive/utils/config"
	"github.com/yuexclusive/utils/registry"
	"github.com/yuexclusive/utils/rpc/middleware/tls"
	"google.golang.org/grpc"
	_ "google.golang.org/grpc/encoding/gzip"
)

type Server struct {
	*grpc.Server
}

func NewServer(serverOptions ...grpc.ServerOption) (*Server, error) {
	t, err := tls.Option()
	if err != nil {
		return nil, err
	}
	serverOptions = append(serverOptions, t)

	server := grpc.NewServer(
		serverOptions...,
	)

	res := &Server{Server: server}

	return res, nil
}

func (s *Server) Serve() error {
	cfg := config.MustGet()

	var address string
	if cfg.Host == "" {
		address = "127.0.0.1:0"
	} else {
		address = cfg.Host + ":" + cfg.Port
	}
	listener, err := net.Listen("tcp", address)

	//registry
	_, err = registry.NewService(cfg.ETCDAddress, cfg.Name, listener.Addr().String(), cfg.Lease)
	if err != nil {
		return err
	}

	return s.Server.Serve(listener)
}
