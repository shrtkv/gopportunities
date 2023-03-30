package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// gin nrouter start with the default configs
	r := gin.Default()
	//Initialize Routes
	initializeRoutes(r)
	// listen and serve on 0.0.0.0:8080
	r.Run()
}
