package redis

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/go-redis/redis"

	"github.com/yuexclusive/utils/log"
)

// Client client interface
type Client struct {
	*redis.Client
	Config *Config
}

var clientMapLock sync.Mutex
var clientMap = make(map[ClientName]*Client)

// GetClient 根据name获取客户端连接
func GetClient(name ClientName) *Client {
	return clientMap[name]
}

// InitClient 初始化连接
func InitClient(config *Config) error {
	check := func(name ClientName, m map[ClientName]*Client) bool {
		if v, ok := m[config.ClientName]; ok && v != nil {
			if _, err := v.Ping().Result(); err == nil {
				return true
			}
		}
		return false
	}

	if check(config.ClientName, clientMap) {
		return nil
	}

	clientMapLock.Lock()
	defer clientMapLock.Unlock()

	// double check
	if check(config.ClientName, clientMap) {
		return nil
	}

	client, err := connect(config)
	if err != nil {
		return err
	}

	clientMap[config.ClientName] = &Client{Client: client, Config: config}

	return nil
}

// init
func init() {
	go monitoring()
}

// monitoring 重连监控,无限循环，检查redis客服端是否断开连接，如果断开重新连接
func monitoring() {
	log.Info("redis 默认开启重连监控")
	defer func() {
		if err := recover(); err != nil {
			log.Error("redis connection aliveness sniffer stopped", "error", err)
			return
		}
		log.Error("redis connection aliveness sniffer stopped")
	}()

	for {
		// 先休眠30秒
		time.Sleep(30 * time.Second)
		c := ping()
		if len(c) <= 0 {
			continue
		}

		log.Info("redis异常断开，正在尝试重连~~~~~")
		reconnect(c)
	}
}

// getExpirationConn 获取过期连接
func ping() []ClientName {
	r := make([]ClientName, 0, 1)
	clientMapLock.Lock()
	defer clientMapLock.Unlock()

	for _, c := range clientMap {
		if _, err := c.Ping().Result(); err == nil {
			continue
		}
		r = append(r, c.Config.ClientName)
	}

	return r
}

// connect 建立redis连接
func connect(config *Config) (*redis.Client, error) {
	opt := &redis.Options{
		Addr:     config.Addr,
		Password: config.Password,
		DB:       config.DB,
		PoolSize: config.PoolSize,
	}
	client := redis.NewClient(opt)

	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}

	// 设置client name(app name)，方便定位问题
	if config.AppName != "" {
		if err := client.Process(redis.NewStringCmd("client", "setname", fmt.Sprintf("%s:%s", config.ClientName, config.AppName))); err != nil {
			return nil, err
		}
	}

	return client, nil
}

// reconnect 重连
func reconnect(c []ClientName) {
	for _, v := range c {
		config := clientMap[v].Config
		err := InitClient(config)
		if err != nil {
			b, _ := json.Marshal(config)
			log.Error("reconnect redis重连失败", "config", string(b), "error", err)
		} else {
			log.Info("reconnect redis重连成功", "client name", config.ClientName)
		}
	}
}
