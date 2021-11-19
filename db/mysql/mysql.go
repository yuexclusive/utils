package mysql

import (
	"github.com/yuexclusive/utils/db"
	"gorm.io/driver/mysql"
)

func Init(dsn string, config *db.Config) {
	db.Init(mysql.Open(dsn), config)
}
