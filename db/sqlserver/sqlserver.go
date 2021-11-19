package mysql

import (
	"github.com/yuexclusive/utils/db"
	"gorm.io/driver/sqlserver"
)

// Open dsn := "sqlserver://gorm:LoremIpsum86@localhost:9930?database=gorm"
func Init(dsn string, config *db.Config) {
	db.Init(sqlserver.Open(dsn), config)
}
