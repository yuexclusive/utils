package main

import (
	"net/http"

	"github.com/yuexclusive/utils/websocket/chat"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	hub := chat.NewHub()
	go hub.Run()
	engine.LoadHTMLGlob("*.html")
	engine.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", nil)
	})

	engine.GET("/ws", func(c *gin.Context) {
		chat.ServeWs(hub, c.Writer, c.Request)
	})

	engine.Run(":8080")

}
