package es

import (
	"context"
	"net/http"
	"sync"
	"time"

	elastic "github.com/olivere/elastic/v7"
	"github.com/yuexclusive/utils/log"
)

var client *elastic.Client
var clientLock sync.Mutex

func InitClient(config *Config) error {
	if client != nil {
		return nil
	}
	clientLock.Lock()
	defer clientLock.Unlock()
	if client == nil {
		ctx, cancel := context.WithTimeout(context.Background(), DefaultQueryTimeOut*time.Second)
		defer cancel()
		options := []elastic.ClientOptionFunc{
			elastic.SetURL(config.Addr),
			elastic.SetSniff(false),
		}
		if config.Auth != nil {
			options = append(options, elastic.SetBasicAuth(config.Auth.UserName, config.Auth.Password))
		}
		var err error
		client, err = elastic.NewClient(options...)

		if err != nil {
			return err
		}

		res, code, err := client.Ping(config.Addr).Do(ctx)

		if err != nil {
			return err
		}

		if code == http.StatusOK {
			log.Info("connected to es cluster successed", "cluster name", res.ClusterName, "number", res.Version.Number)
		} else {
			log.Warn("connected to es cluster failed", "cluster name", res.ClusterName, "number", res.Version.Number, "code", code)
		}
	}

	return nil
}

func Client() *elastic.Client {
	return client
}
