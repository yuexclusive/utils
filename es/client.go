package es

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	elastic "github.com/olivere/elastic/v7"
	"github.com/yuexclusive/utils/logger"
)

var client *elastic.Client
var clientLock sync.Mutex
var log = logger.Single()

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
			log.Info(fmt.Sprintf("connected to es cluster: %s ,version: %s", res.ClusterName, res.Version.Number))
		} else {
			log.Warn(fmt.Sprintf("connected to es cluster: %s ,version: %s, return code %d", res.ClusterName, res.Version.Number, code))
		}
	}

	return nil
}

func Client() *elastic.Client {
	return client
}
