package postgres

import (
	"github.com/yuexclusive/utils/db"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func Init(dsn string, name string, opts ...gorm.Option) {
	db.Open(name, sqlserver.Open(dsn), opts...)
}

func InitDefault(dsn string, opts ...gorm.Option) {
	db.Open("", sqlserver.Open(dsn), opts...)
}
