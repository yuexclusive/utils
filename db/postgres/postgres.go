package postgres

import (
	"github.com/yuexclusive/utils/db"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(dsn string, name string, opts ...gorm.Option) {
	db.Open(name, postgres.Open(dsn), opts...)
}

func InitDefault(dsn string, opts ...gorm.Option) {
	db.Open("", postgres.Open(dsn), opts...)
}
