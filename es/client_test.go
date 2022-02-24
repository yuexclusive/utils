package es

import (
	"context"
	"fmt"
	"testing"
)

type Employee struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestClient(t *testing.T) {
	InitClient(&Config{Addr: "http://localhost:9200"})

	// _, err := GetClient().Index().Index("employee").BodyJson(&Employee{Name: "jj", Age: 35}).Do(context.Background())

	// if err != nil {
	// 	panic(err)
	// }

	// panic("test panic")

	ok, err := GetClient().Exists().Index("employee").Id("1").Do(context.Background())

	// if err != nil {
	// 	panic(err)
	// }

	if ok {
		deleteRes, err := GetClient().Delete().Index("employee").Id("1").Do(context.Background())
		if err != nil {
			panic(err)
		}

		fmt.Println(deleteRes.Result)

	}

	res, err := GetClient().Index().Index("employee").Id("1").BodyJson(&Employee{Name: "jj", Age: 35}).Do(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Println(res)

	res2, err := GetClient().Search("employee").Do(context.Background())

	if err != nil {
		panic(err)
	}

	for _, v := range res2.Hits.Hits {
		fmt.Println(string(v.Source))
	}

}
