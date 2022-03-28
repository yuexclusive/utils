package uri

import (
	"fmt"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

type Person struct {
	ID      string `uri:"id" binding:"required,uuid"`
	Name    string `uri:"name" binding:"required"`
	Version string `uri:"version"`
}

func Example() {
	route := gin.Default()

	route.GET("/:name/:id/*version", func(c *gin.Context) {
		var person Person
		if err := c.ShouldBindUri(&person); err != nil {
			c.JSON(400, gin.H{"msg": err})
			return
		}
		c.JSON(200, gin.H{"name": person.Name, "uuid": person.ID, "version": person.Version})
	})

	w := httptest.NewRecorder()

	r := httptest.NewRequest("GET", "/thinkerou/987fbc97-4bed-5078-9f07-9141ba07c9f3/", nil)

	route.ServeHTTP(w, r)

	w2 := httptest.NewRecorder()

	r2 := httptest.NewRequest("GET", "/thinkerou/not-uuid/", nil)

	route.ServeHTTP(w2, r2)

	w3 := httptest.NewRecorder()

	r3 := httptest.NewRequest("GET", "/thinkerou/987fbc97-4bed-5078-9f07-9141ba07c9f3/1.1.1", nil)

	route.ServeHTTP(w3, r3)

	//Output:
	//{"name":"thinkerou","uuid":"987fbc97-4bed-5078-9f07-9141ba07c9f3","version":"/"}
	//{"msg":[{}]}
	//{"name":"thinkerou","uuid":"987fbc97-4bed-5078-9f07-9141ba07c9f3","version":"/1.1.1"}
	fmt.Println(w.Body.String())
	fmt.Println(w2.Body.String())
	fmt.Println(w3.Body.String())
}
