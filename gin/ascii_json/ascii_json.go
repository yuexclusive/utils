package ascii_json

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Foo(r *gin.Engine) {
	r.GET("/someJSON", func(c *gin.Context) {
		data := map[string]interface{}{
			"lang": "GO语言",
			"tag":  "<br>",
		}

		c.AsciiJSON(http.StatusOK, data)
	})
}
