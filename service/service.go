package service

import (
	"github.com/gin-gonic/gin"
	"github.com/rsingla/game_server/request"
)

type TagsRequest interface {
	FindAll(*gin.Context)
	FindById(id int) (*request.TagsReq, error)
	FindBySlug(slug string) (*request.TagsReq, error)
	Save(*gin.Context)
	Update(*gin.Context) (*request.TagsReq, error)
	Delete(id int) error
}
