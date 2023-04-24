package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rsingla/game_server/helper"
	"github.com/rsingla/game_server/model"
	"github.com/rsingla/game_server/repository"
	"github.com/rsingla/game_server/request"
)

type TagsRequest interface {
	FindAll(*gin.Context)
	FindById(id int) (*request.TagsReq, error)
	FindBySlug(slug string) (*request.TagsReq, error)
	Save(*gin.Context)
	Update(*gin.Context) (*request.TagsReq, error)
	Delete(tags *request.TagsReq) error
}

func FindAll(c *gin.Context) {

	tagsImpl := repository.TagsImpl{Db: helper.ConnectDB()}

	tags, err := tagsImpl.FindAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"data": "Error while fetching tags"})
	}

	c.JSON(http.StatusOK, gin.H{"data": tags})

}

func Save(c *gin.Context) {

	tagsImpl := repository.TagsImpl{Db: helper.ConnectDB()}

	var tags request.TagsReq

	if err := c.ShouldBindJSON(&tags); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Convert tags to model.Tags
	tagsModel := model.Tags{Name: tags.Name, Slug: tags.Slug, Description: tags.Description, Status: tags.Status, Type: tags.Type, Category: tags.Category}

	modelTags, err := tagsImpl.Save(&tagsModel)

	//Convert model.Tags to request.TagsReq
	tags = request.TagsReq{Name: modelTags.Name, Slug: modelTags.Slug, Description: modelTags.Description, Status: modelTags.Status, Type: modelTags.Type, Category: modelTags.Category}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"data": "Error while saving tags"})
	}

	c.JSON(http.StatusOK, gin.H{"data": tags})

}
