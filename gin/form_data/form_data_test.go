package form_data

import (
	"fmt"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

func ExampleFoo() {
	engine := gin.Default()

	engine.GET("/getb", GetDataB)
	engine.GET("/getc", GetDataC)
	engine.GET("/getd", GetDataD)

	var slice []*httptest.ResponseRecorder

	for i := 0; i < 3; i++ {
		slice = append(slice, httptest.NewRecorder())
	}

	engine.ServeHTTP(slice[0], httptest.NewRequest("GET", "/getb?field_a=hello&field_b=world", nil))
	engine.ServeHTTP(slice[1], httptest.NewRequest("GET", "/getc?field_a=hello&field_c=world", nil))
	engine.ServeHTTP(slice[2], httptest.NewRequest("GET", "/getd?field_x=hello&field_d=world", nil))

	//Output:
	//{"a":{"FieldA":"hello"},"b":"world"}
	//{"a":{"FieldA":"hello"},"c":"world"}
	//{"d":"world","x":{"FieldX":"hello"}}
	fmt.Println(slice[0].Body.String())
	fmt.Println(slice[1].Body.String())
	fmt.Println(slice[2].Body.String())
}
