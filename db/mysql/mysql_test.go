package mysql

import (
	"fmt"
	"testing"

	"github.com/yuexclusive/utils/db"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type Role struct {
	gorm.Model
	Name        string
	Description string
}

type UserRoleMap struct {
	ID     uint
	RoleID uint
	UserID uint
	Role   Role
}

func TestInit(t *testing.T) {
	InitDefault("test:123456@tcp(127.0.0.1:3306)/evolve?charset=utf8mb4&parseTime=True&loc=Local", &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
		Logger:         logger.Default.LogMode(logger.Info),
	})

	var data []UserRoleMap

	db.GetDefaultClient().Joins("Role").Find(&data)

	fmt.Println(data)
}
