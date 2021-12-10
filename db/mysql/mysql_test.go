package mysql

import (
	"testing"
)

func TestInit(t *testing.T) {
	Init("test:123@tcp(127.0.0.1:3306)/evolve?charset=utf8mb4&parseTime=True&loc=Local")

	// select {}
}
