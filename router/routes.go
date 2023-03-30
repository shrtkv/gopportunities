package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		v1.GET("/opening", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"msg": "Get Opening"})
		})
		v1.POST("/opening", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"msg": "Create Opening"})
		})
		v1.DELETE("/opening", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"msg": "Delete Opening"})
		})
		v1.PUT("/opening", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"msg": "Update Opening"})
		})
		v1.GET("/openings", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"msg": "Get Openings"})
		})
	}
}
