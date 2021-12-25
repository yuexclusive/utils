package postgres

import (
	"github.com/yuexclusive/utils/db"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Init(dsn string, name string, opts ...gorm.Option) {
	db.Open(name, sqlite.Open(dsn), opts...)
}

func InitDefault(dsn string, opts ...gorm.Option) {
	db.Open("", sqlite.Open(dsn), opts...)
}
