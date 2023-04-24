package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rsingla/game_server/helper"
	"github.com/rsingla/game_server/service"
)

func main() {

	helper.ConnectDB()

	server := gin.Default()
	server.SetTrustedProxies([]string{"192.168.1.1", "localhost", "127.0.0.1"})

	server.GET("/", welcome)

	server.POST("/tags", service.Save)

	server.Run(":8000")
}

func welcome(c *gin.Context) {
	c.JSON(http.StatusOK, "Welcome to Game server")
}
