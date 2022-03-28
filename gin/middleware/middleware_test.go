package middleware

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// Set example variable
		c.Set("example", "12345")

		// before request

		c.Next()

		// after request
		latency := time.Since(t)
		log.Print(latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
	}
}

func Example() {
	engine := gin.New()
	engine.Use(Logger())

	engine.GET("/test", func(c *gin.Context) {
		example := c.MustGet("example").(string)

		// it would print: "12345"
		c.String(http.StatusOK, example)
	})

	w := httptest.NewRecorder()

	r := httptest.NewRequest("GET", "/test", nil)

	engine.ServeHTTP(w, r)

	//Output:
	//12345
	fmt.Println(w.Body.String())
}
