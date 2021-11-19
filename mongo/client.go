package mongo

import (
	"context"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/yuexclusive/utils/logger"
)

var clientMapLock sync.Mutex
var clientMap = make(map[ClientName]*mongo.Client)
var configMap = make(map[ClientName]*Config)

var log = logger.Single().Sugar()

// Client 根据name获取client
// 注意：本方法直接返回client结果，是因为传入的name一定是写好的constant，而且一定已经成功初始化完成，否则程序不可能执行到这里，所以不用做任何判断
func Client(name ClientName) *mongo.Client {
	return clientMap[name]
}

// InitClient 根据配置初始化Client
func InitClient(c *Config) (*mongo.Client, error) {
	if clientMap[c.ClientName] != nil {
		return clientMap[c.ClientName], nil
	}
	var err error
	var client *mongo.Client
	clientMapLock.Lock()
	defer clientMapLock.Unlock()

	if _, exist := clientMap[c.ClientName]; !exist {
		client, err = connect(c)
		if client != nil {
			clientMap[c.ClientName] = client
			configMap[c.ClientName] = c
		}
	}

	return client, err
}

// connect
func connect(c *Config) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DefaultConnectTimeout*time.Second)
	defer cancel()

	opts := []*options.ClientOptions{
		options.Client().ApplyURI(c.Addr),
		options.Client().SetConnectTimeout(DefaultConnectTimeout * time.Second),
	}

	if c.Auth != nil {
		opts = append(opts,
			options.Client().SetAuth(options.Credential{
				AuthMechanism: "SCRAM-SHA-1",
				AuthSource:    c.Auth.AuthSource,
				Username:      c.Auth.Username,
				Password:      c.Auth.Password,
				PasswordSet:   c.Auth.PasswordSet,
			}),
		)
	}

	client, err := mongo.Connect(
		ctx,
		opts...,
	// options.Client().SetAppName(c.AppName),
	)

	if err != nil {
		return nil, err
	}
	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		return nil, err
	}
	return client, nil
}

func init() {
	go monitoring()
}

// monitoring 开启状态监控
func monitoring() {
	log.Info("mongo connection aliveness start")
	defer func() {
		if err := recover(); err != nil {
			log.Errorf("mongo connection aliveness monitoring stopped: %v", err)
			return
		}
		log.Error("mongo connection aliveness monitoring stopped")
	}()
	for {
		// 先休眠30秒
		time.Sleep(30 * time.Second)
		c := ping()
		if len(c) == 0 {
			continue
		}

		// 重连
		reconnect(c)
	}
}

// ping 获取过期连接
func ping() []ClientName {
	r := make([]ClientName, 0, 1)
	clientMapLock.Lock()
	defer clientMapLock.Unlock()
	for k, v := range clientMap {
		if v == nil {
			r = append(r, k)
			continue
		}
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()
		if err := v.Ping(ctx, readpref.Primary()); err != nil {
			log.Errorf("[%s]mongo connection closed", k)
			r = append(r, k)
			// {
			// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
			// 	defer cancel()
			// 	_ = v.Disconnect(ctx)
			// }
			// clientMap[k] = nil
		}
	}
	return r
}

// reconnect 重连
func reconnect(c []ClientName) {
	for _, v := range c {
		config := configMap[v]
		if _, err := InitClient(config); err != nil {
			log.Errorf("reconnect %s failed: %v", v, err)
		}
	}
}
