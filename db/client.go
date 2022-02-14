package db

import (
	"fmt"
	"sync"
	"time"

	log "logger"

	"gorm.io/gorm"
)

const (
	_defaultClientName    = "default_client"
	_defaultWatchDuration = time.Second * 5
)

type DB struct {
	*gorm.DB
}

func toDB(db *gorm.DB) *DB {
	return &DB{DB: db}
}

type Client struct {
	*DB
	Closed bool
}

func Open(name string, dialector gorm.Dialector, opts ...gorm.Option) error {
	db, err := gorm.Open(dialector, opts...)

	if err != nil {
		return err
	}

	if name == "" {
		name = _defaultClientName
	}

	mapClient[name] = &Client{DB: toDB(db), Closed: false}
	return nil
}

var mapClient = make(map[string]*Client)

func GetDefaultClient() *Client {
	if v, ok := mapClient[_defaultClientName]; !ok {
		panic(fmt.Sprintf("please init client %s first", _defaultClientName))
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
			if err := Open(name, v.Dialector, v.Config); err != nil {
				log.Error("recover db failed", "name", name, "err", err)
			} else {
				log.Error("recover db successful", "name", name)
			}
		}
	}
}
