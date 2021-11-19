package etcd

import (
	"time"

	"github.com/yuexclusive/utils/config"

	etcd "go.etcd.io/etcd/client/v3"
)

func Client() (*etcd.Client, error) {
	config := etcd.Config{
		Endpoints:   config.MustGet().ETCDAddress,
		DialTimeout: 10 * time.Second,
	}
	return etcd.New(config)
}
