package api

import (
	v1 "github.com/SaidovZohid/blog_db/api/v1"
	"github.com/SaidovZohid/blog_db/config"
	"github.com/SaidovZohid/blog_db/storage"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type RoutetOptions struct {
	Cfg     *config.Config
	Storage storage.StorageI
}

// @title           Swagger for blog api
// @version         2.0
// @description     This is a blog service api.
// @host      localhost:8080
// @BasePath  /v1
func New(opt *RoutetOptions) *gin.Engine {
	router := gin.Default()

	handlerV1 := v1.New(&v1.HandlerV1Options{
		Cfg:     opt.Cfg,
		Storage: opt.Storage,
	})

	apiV1 := router.Group("/v1")

	apiV1.POST("/users", handlerV1.CreateUser)
	apiV1.GET("/users/:id", handlerV1.GetUser)
	apiV1.PUT("/users/update/:id", handlerV1.UpdateUser)
	apiV1.DELETE("/users/delete/:id", handlerV1.DeleteUser)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
