package query_string

import (
	"fmt"
	"net/http/httptest"
	"time"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Name     string    `form:"name"`
	Address  string    `form:"address"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02/15:04:05"`
}

func startPage(c *gin.Context) {
	var person Person
	// If `GET`, only `Form` binding engine (`query`) used.
	// If `POST`, first checks the `content-type` for `JSON` or `XML`, then uses `Form` (`form-data`).
	// See more at https://github.com/gin-gonic/gin/blob/master/binding/binding.go#L48
	if err := c.Bind(&person); err != nil {
		panic(err)
	}

	c.JSON(200, person)
}
func Example() {
	engine := gin.Default()
	engine.GET("/test", startPage)
	w := httptest.NewRecorder()

	r := httptest.NewRequest("GET", "/test?name=appleboy&address=xyz&birthday=1992-03-15/00:13:11", nil)

	engine.ServeHTTP(w, r)

	//Output:
	//{"Name":"appleboy","Address":"xyz","Birthday":"1992-03-15T00:13:11+08:00"}
	fmt.Println(w.Body.String())
}
