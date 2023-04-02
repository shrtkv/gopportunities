package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ShowOpeningHandler ...
func ShowOpeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"msg": "Get Opening"})
}
