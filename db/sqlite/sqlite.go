package sqlite

import (
	"github.com/yuexclusive/utils/db"
	"gorm.io/driver/sqlite"
)

// Init dsn := "sqlserver://gorm:LoremIpsum86@localhost:9930?database=gorm"
func Init(dsn string, config *db.Config) {
	db.Init(sqlite.Open(dsn), config)
}
