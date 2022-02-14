package mysql

import (
	"github.com/yuexclusive/utils/db"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Init(dsn string, name string, opts ...gorm.Option) {
	db.Open(name, mysql.Open(dsn), opts...)
}

func InitDefault(dsn string, opts ...gorm.Option) {
	db.Open("", mysql.Open(dsn), opts...)
}
