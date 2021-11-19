package postgres

import (
	"github.com/yuexclusive/utils/db"
	"gorm.io/driver/postgres"
)

func Init(dsn string, cfgs ...*db.Config) {
	db.Init(postgres.Open(dsn), cfgs...)
}
