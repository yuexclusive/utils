package main

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/yuexclusive/utils/web"
)

func main() {
	engine := gin.Default()

	engine.Use(web.Auth())

	engine.GET("/hello", func(c *gin.Context) {
		fmt.Println(c.Request.Header)
		test_context(c)
	})

	engine.Run(":8081")
}

func test_context(ctx context.Context) {
	fmt.Println(ctx.Value("claims"))
	fmt.Println(ctx.Value("claims"))
	fmt.Println(ctx.Value("claims"))
}
