package checkbox

import (
	"fmt"
	"net/http/httptest"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

type myForm struct {
	Colors []string `form:"colors" json:"color"`
}

func ExampleFoo() {
	engine := gin.Default()

	engine.POST("/form_handler", func(c *gin.Context) {
		var fakeForm myForm
		c.ShouldBind(&fakeForm)
		c.JSON(200, fakeForm)
	})

	w := httptest.NewRecorder()

	data := make(url.Values)

	data.Add("colors", "red")
	data.Add("colors", "green")
	data.Add("colors", "blue")

	r := httptest.NewRequest("POST", "/form_handler", strings.NewReader(data.Encode()))

	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	engine.ServeHTTP(w, r)

	//Output: {"color":["red","green","blue"]}
	fmt.Println(w.Body.String())
}
