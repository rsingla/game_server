package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rsingla/game_server/request"
	"github.com/rsingla/game_server/service"
)

type TagsController struct {
	tagsService service.TagsService
}

func NewTagsController(service service.TagsService) *TagsController {
	return &TagsController{
		tagsService: service,
	}
}

func (t *TagsController) FindAll() {
	panic("implement me")
}

func (t *TagsController) Save(c *gin.Context) {
	var tags request.TagsReq

	if err := c.ShouldBindJSON(&tags); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tagModel, err := t.tagsService.Save(tags)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tagModel)
}

func (t *TagsController) DeleteById(c *gin.Context) {

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

	t.tagsService.Delete(i)

	c.JSON(http.StatusOK, gin.H{"data": "Tag deleted successfully"})
}
