package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/yuexclusive/utils/config"
	"github.com/yuexclusive/utils/log"
)

type Config struct {
	config.Config
}

func main() {
	engine := gin.Default()

	log.GinUseZap(engine)

	engine.GET("/metrics", func(c *gin.Context) {
		h := promhttp.Handler()
		h.ServeHTTP(c.Writer, c.Request)
	})

	engine.GET("/foo", func(c *gin.Context) {
		name, err := os.Hostname()

		log.Info("foo msg by hand for test", "name", "jiaojiao")

		if err != nil {
			panic(err)
		}
		c.String(http.StatusOK, "%s %s", "foo", name)
	})

	engine.GET("/bar", func(c *gin.Context) {
		name, err := os.Hostname()

		if err != nil {
			panic(err)
		}
		c.String(http.StatusOK, "%s %s", "bar", name)
	})

	engine.GET("/name", func(c *gin.Context) {

		res := config.Init[Config](config.TOML, "./config.toml")

		cc := res.GetConfig()

		c.String(http.StatusOK, "%s %s", "name", cc.Name)
	})

	engine.GET("/name2", func(c *gin.Context) {
		res, err := http.Get("http://cache/data")

		if err != nil {
			c.Error(err)
		}

		bytes := make([]byte, 1<<10)

		n, _ := res.Body.Read(bytes)

		defer res.Body.Close()

		var data Data
		if err := json.Unmarshal(bytes[:n], &data); err != nil {
			c.Error(err)
		}

		c.String(http.StatusOK, data.Name)
	})

	engine.GET("/name3", func(c *gin.Context) {
		os.Exit(1)
	})

	engine.GET("/error", func(c *gin.Context) {
		c.String(http.StatusInternalServerError, "test error")
	})

	if err := engine.Run(":8080"); err != nil {
		log.Panic(err.Error())
	}
}

type Data struct {
	Name string `json:"name"`
}
