package main

import (
	"net/http"
	"time"

	"github.com/yuexclusive/utils/jwt"
	"github.com/yuexclusive/utils/log"

	"github.com/gin-gonic/gin"
)

const (
	KEY = "sadhasldjkko126312jljdkhfasu0"
)

type Auth struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

func (a *Auth) Login() error {
	// todo: implememnt login
	return nil
}

func (a *Auth) Auth() (string, error) {
	if err := a.Login(); err != nil {
		return "", err
	} else {
		return jwt.GenToken(jwt.StandardClaims{Id: a.UserName, ExpiresAt: time.Now().Add(time.Hour * 24).Unix()}, KEY)
	}
}

func main() {
	engine := gin.Default()

	engine.POST("/auth", func(c *gin.Context) {
		var auth Auth
		if err := c.BindJSON(&auth); err != nil {
			c.Error(err)
			return
		}

		if token, err := auth.Auth(); err != nil {
			c.Error(err)
			return
		} else {
			c.String(http.StatusOK, token)
		}
	})

	engine.GET("/validate", func(c *gin.Context) {
		token := c.GetHeader("token")
		if claims, err := jwt.GetClaims(token, KEY); err != nil {
			c.Error(err)
			return
		} else {
			c.JSON(http.StatusOK, claims)
		}
	})

	if err := engine.Run(":8080"); err != nil {
		log.Panic(err.Error())
	}
}
