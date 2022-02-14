package es

import (
	"context"
	"fmt"
	"testing"
)

type Employee struct {
	Name string
	Age  int
}

func TestClient(t *testing.T) {
	InitClient(&Config{Addr: "http://localhost:9200"})

	// _, err := GetClient().Index().Index("employee").BodyJson(&Employee{Name: "jj", Age: 35}).Do(context.Background())

	// if err != nil {
	// 	panic(err)
	// }

	res, err := GetClient().Search("employee").Do(context.Background())

	if err != nil {
		panic(err)
	}

	for _, v := range res.Hits.Hits {
		fmt.Println(v.Id)
	}

}
