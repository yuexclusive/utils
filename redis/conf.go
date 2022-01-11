package redis

// Config redis配置结构
type Config struct {
	AppName    string // app name
	ClientName string // client name
	Addr       string // redis addr，例如 127.0.0.1:6379
	Password   string // redis password
	DB         int    // redis db
	PoolSize   int    // redis poolsize
}
