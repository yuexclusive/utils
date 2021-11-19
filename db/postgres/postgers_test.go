package postgres

import (
	"testing"
)

func Test_Open(t *testing.T) {
	Init("host=127.0.0.1 port=5432 user=postgres password=123456 sslmode=disable dbname=evolve")
	select {}
}
