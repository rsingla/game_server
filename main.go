package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()
	server.SetTrustedProxies([]string{"192.168.1.1", "localhost", "127.0.0.1"})

	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, "pong")
	})

	server.GET("/", welcome)

	server.Run(":8000")
}

func welcome(c *gin.Context) {
	c.JSON(http.StatusOK, "Welcome to Game server")
}
