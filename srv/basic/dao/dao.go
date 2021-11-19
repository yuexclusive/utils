package dao

import "github.com/yuexclusive/utils/db"

type IDao interface{}

type Dao struct {
	*db.Dao
}

func NewDao() IDao {
	return &Dao{Dao: db.NewDao("test")}
}
