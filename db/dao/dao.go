package dao

import "github.com/yuexclusive/utils/db"

// Dao Dao
type Dao struct {
	ConnName string
	DB       *db.DB
}

// NewDao NewDao
func NewDao(connName string) *Dao {
	return nil
}

func (d *Dao) Test() {

}
