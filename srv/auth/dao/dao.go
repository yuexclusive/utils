package dao

import (
	"github.com/yuexclusive/utils/db"
	"github.com/yuexclusive/utils/srv/auth/model"
)

// IDao interface
type IDao interface {
	GetByID(id string) model.User
}

// Dao dao
type Dao struct {
	*db.Dao
}

// NewDao new dao
func NewDao() IDao {
	return &Dao{Dao: db.NewDao("test")}
}

// GetByID GetByID
func (d *Dao) GetByID(id string) model.User {
	var user model.User
	d.Where("name=?", id).First(&user)
	return user
}
