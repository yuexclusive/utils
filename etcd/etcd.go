package etcd

import (
	"time"

	etcd "go.etcd.io/etcd/client/v3"
)

func Client(address []string) (*etcd.Client, error) {
	config := etcd.Config{
		Endpoints:   address,
		DialTimeout: 10 * time.Second,
	}
	return etcd.New(config)
}
