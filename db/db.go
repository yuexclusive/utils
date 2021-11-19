package db

import (
	"fmt"
	"sync"
	"time"

	"github.com/yuexclusive/utils/log"
	"gorm.io/gorm"
)

const (
	DefaultClientName = "default_client"
)

// DB DB
type DB struct {
	*gorm.DB
}

// toDB toDB transfer *gorm.DB to *DB
func toDB(gdb *gorm.DB) *DB {
	return &DB{gdb}
}

func open(dialector gorm.Dialector, cfg *Config) error {
	name := DefaultClientName
	if cfg != nil {
		name = cfg.Name
	}

	if client, ok := mapClient[name]; ok {
		if !client.Closed {
			return nil
		}
	}

	var (
		gdb *gorm.DB
		err error
	)
	if cfg != nil {
		gdb, err = gorm.Open(dialector, cfg)
	} else {
		gdb, err = gorm.Open(dialector)
	}

	if err != nil {
		return err
	}

	if cfg != nil {
		if cfg.MaxIdleConns != nil || cfg.MaxOpenConns != nil {
			sqldb, err := gdb.DB()

			if err != nil {
				return err
			}

			if cfg.MaxIdleConns != nil {
				sqldb.SetMaxIdleConns(*cfg.MaxIdleConns)
			}

			if cfg.MaxOpenConns != nil {
				sqldb.SetMaxOpenConns(*cfg.MaxOpenConns)
			}
		}
	}

	if cfg == nil {
		cfg = toConfig(gdb.Config)
		cfg.Name = name
	}

	mapClient[name] = &Client{DB: toDB(gdb), Config: cfg, Closed: false}
	return nil
}

type Client struct {
	*DB
	Config *Config
	Closed bool
}

var mapClient = make(map[string]*Client)

func Init(dialector gorm.Dialector, cfgs ...*Config) {
	var cfg *Config

	for _, v := range cfgs {
		if v != nil {
			cfg = v
		}
	}
	if err := open(dialector, cfg); err != nil {
		log.Fatal("init db failed", "err", err)
	}
}

func GetDefaultClient() *Client {
	if v, ok := mapClient[DefaultClientName]; !ok {
		panic(fmt.Sprintf("please init client %s first", DefaultClientName))
	} else {
		return v
	}
}

func GetClient(name string) *Client {
	if v, ok := mapClient[name]; !ok {
		panic(fmt.Sprintf("please init client %s first", name))
	} else {
		return v
	}
}

func init() {
	go watch()
}

const (
	_defaultWatchDuration = time.Second * 5
)

var watchMutex sync.Mutex

func watch() {
	ticker := time.NewTicker(_defaultWatchDuration)
	for range ticker.C {
		ping()
		recover()
	}
}

func ping() {
	watchMutex.Lock()
	defer watchMutex.Unlock()
	for name, v := range mapClient {
		if !v.Closed {
			if sqldb, err := v.DB.DB.DB(); err != nil {
				v.Closed = true
				log.Error("get sql db failed", "name", name, "err", err)
			} else {
				if err := sqldb.Ping(); err != nil {
					v.Closed = true
					log.Error("ping db failed", "name", name, "err", err)
				}
			}
		}
	}

}

func recover() {
	watchMutex.Lock()
	defer watchMutex.Unlock()
	for name, v := range mapClient {
		if v.Closed {
			if err := open(v.Dialector, v.Config); err != nil {
				log.Error("recover db failed", "name", name, "err", err)
			} else {
				log.Error("recover db successful", "name", name)
			}
		}
	}
}
