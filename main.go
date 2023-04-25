package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/go-playground/validator/v10"

	"github.com/rsingla/game_server/controller"
	"github.com/rsingla/game_server/helper"
	"github.com/rsingla/game_server/repository"
	"github.com/rsingla/game_server/service"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {

	server := gin.Default()
	server.SetTrustedProxies([]string{"192.168.1.1", "localhost", "127.0.0.1"})

	db := helper.ConnectDB()
	validate := validator.New()

	tagsRepository := repository.TagsRepoImpl(db)

	tagsService := service.NewTagsService(tagsRepository, validate)

	tagsController := controller.NewTagsController(*tagsService)

	server.GET("/", welcome)

	router := NewRouter(tagsController)

	router.Run(":8000")
}

func NewRouter(tagsController *controller.TagsController) *gin.Engine {
	router := gin.Default()
	// add swagger
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	baseRouter := router.Group("/api")
	tagsRouter := baseRouter.Group("/tags")
	//tagsRouter.GET("", tagsController.FindAll)
	//tagsRouter.GET("/:tagId", tagsController.FindById)
	tagsRouter.POST("", tagsController.Save)
	//tagsRouter.PATCH("/:tagId", tagsController.Update)
	tagsRouter.DELETE("/:id", tagsController.DeleteById)

	return router
}

func welcome(c *gin.Context) {
	c.JSON(http.StatusOK, "Welcome to Game server")
}
