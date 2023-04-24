package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rsingla/game_server/request"
	"github.com/rsingla/game_server/service"
)

func main() {

	server := gin.Default()
	server.SetTrustedProxies([]string{"192.168.1.1", "localhost", "127.0.0.1"})

	server.GET("/", welcome)

	server.POST("/tags", Save)
	server.DELETE("/tags/:id", DeleteById)

	server.Run(":8000")
}

func Save(c *gin.Context) {
	var tags request.TagsReq

	if err := c.ShouldBindJSON(&tags); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tagsService := service.TagsService{}
	tagModel, err := tagsService.Save(&tags)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tagModel)
}

func DeleteById(c *gin.Context) {

	id := c.Param("id")

	log.Println("id", id)

	if id == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Id is required"})
		return
	}

	i, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.Println("id Value : ", i)

	tagsRepository := service.TagsService{}
	tagsSvc := service.TagsService{TagsRepository: tagsRepository.TagsRepository}

	tagsSvc.Delete(i)
	c.JSON(http.StatusOK, gin.H{"data": "Tag deleted successfully"})
}

func welcome(c *gin.Context) {
	c.JSON(http.StatusOK, "Welcome to Game server")
}
