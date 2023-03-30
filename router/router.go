package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// gin router start with the default configs
	r := gin.Default()
	//defines a route
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// listen and serve on 0.0.0.0:8080
	r.Run()
}
