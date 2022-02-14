package main

import (
	"net/http"

	"github.com/yuexclusive/utils/log"

	"github.com/gin-gonic/gin"
	"github.com/yuexclusive/utils/config"
	"github.com/yuexclusive/utils/redis"
)

type Config struct {
	config.Config
}

type Result struct {
	Name string `json:"name"`
}

func main() {

	cfg := config.Init[Config](config.TOML, "./config.toml")

	engine := gin.Default()

	log.GinUseZap(engine)

	if r := redis.InitClient(&redis.Config{Addr: cfg.GetConfig().Redis.Address}); r != nil {
		log.Panic(r.Error())
	}

	engine.GET("/data", func(c *gin.Context) {
		res := redis.GetClient("").Get("name")
		if err := res.Err(); err != nil {
			c.JSON(http.StatusOK, Result{Name: "no data"})
		} else {
			c.JSON(http.StatusOK, Result{Name: res.Val()})
		}
	})

	engine.Run(":8080")
}
