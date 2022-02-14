package redis

import (
	"fmt"
	"testing"
)

func TestClient(t *testing.T) {
	InitClient(&Config{Addr: "localhost:6379"})

	client := GetClient("")

	res, _, err := client.Scan(0, "*", 1).Result()

	if err != nil {
		panic(err)
	}

	fmt.Println(res)

}
