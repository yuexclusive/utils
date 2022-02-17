package config

// Config Config
type Config struct {
	Name     string    `mapstructure:"name"` //name of app
	Host     string    `mapstructure:"host"`
	Port     string    `mapstructure:"port"`
	Log      *Log      `mapstructure:"log"`
	ETCD     *ETCD     `mapstructure:"etcd"`
	Postgres *Postgres `mapstructure:"postgres"`
	Redis    *Redis    `mapstructure:"redis"`
	Jaeger   *Jaeger   `mapstructure:"jaeger"`
	ES       *ES       `mapstructure:"es"`
	TLS      *TLS      `mapstructure:"tls"`
	Mongo    *Mongo    `mapstructure:"mongo"`
	Nats     *Nats     `mapstructure:"nats"`
	AuthHost string    `mapstructure:"authHost"`
}

// Mongo Mongo
type Mongo struct {
	Address string `mapstructure:"address"`
}

// Log Log
type Log struct {
	// Path path of log file
	Path string `mapstructure:"path"`
	// Mode "production" means production mode
	Mode string `mapstructure:"mode"`
}

// Jaeger Jaeger
type Jaeger struct {
	Address string `mapstructure:"address"`
}

// Redis Redis
type Redis struct {
	Address  string `mapstructure:"address"`
	Password string `mapstructure:"password"`
	PoolSize string `mapstructure:"poolsize"`
	Cachedb  string `mapstructure:"cachedb"`
}

// Postgre Postgre
type Postgres struct {
	Address string `mapstructure:"address"`
}

// ES ES
type ES struct {
	Address string `mapstructure:"address"`
}

// Nats Nats
type Nats struct {
	Address string `mapstructure:"address"`
}

// ETCD ETCD
type ETCD struct {
	Address []string `mapstructure:"address"`
	Lease   int      `mapstructure:"lease"`
}

// TLS TLS
type TLS struct {
	CertFile           string `mapstructure:"certFile"`
	KeyFile            string `mapstructure:"keyFile"`
	CACertFile         string `mapstructure:"caCertFile"`
	ServerNameOverride string `mapstructure:"serverNameOverride"`
}
