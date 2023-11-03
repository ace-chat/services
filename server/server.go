package server

import (
	"ace/controller"
	"ace/middleware"
	"github.com/gin-gonic/gin"
)

func NewServer(mode string) *gin.Engine {
	gin.SetMode(mode)
	r := gin.Default()

	r.Use(middleware.Cors())

	api := r.Group("/api/v1")
	{
		api.POST("/login", controller.Login)
		template := api.Group("/template")
		template.Use(middleware.Auth())
		{
			template.POST("/example", controller.Example)
		}
	}

	return r
}
