package mysql

import (
	"fmt"
	"testing"

	"github.com/yuexclusive/utils/db"
	"github.com/yuexclusive/utils/db/dao"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Role struct {
	ID          int64
	Name        string
	Description string
}

func (r *Role) PK() int64 {
	return r.ID
}

func TestInit(t *testing.T) {
	InitDefault("test:123@tcp(127.0.0.1:30002)/evolve?charset=utf8mb4&parseTime=True&loc=Local", &gorm.Config{NamingStrategy: schema.NamingStrategy{SingularTable: true}})

	d := dao.NewDao[Role](db.GetDefaultClient().DB)

	d.Update(&Role{Description: "heh1"}, []dao.Where{{Query: "id=?", Args: []interface{}{1}}}, "Description")

	data := d.Find(nil)

	fmt.Println(data)
}
