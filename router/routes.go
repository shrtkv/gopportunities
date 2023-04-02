package router

import (
	"github.com/gin-gonic/gin"
	docs "github.com/shrtkv/gopportunities/docs"
	"github.com/shrtkv/gopportunities/router/handler"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(r *gin.Engine) {
	// initialize Handler
	handler.InitializeHandler()
	basePath := "api/v1"
	docs.SwaggerInfo.BasePath = basePath
	v1 := r.Group(basePath)
	{
		v1.GET("/opening", handler.ShowOpeningHandler)
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.DELETE("/opening", handler.DeleteOpeningHandler)
		v1.PUT("/opening", handler.UpdateOpeningHandler)
		v1.GET("/openings", handler.ListOpeningsHandler)
	}
	// Initialize swagger
	r.GET("swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
