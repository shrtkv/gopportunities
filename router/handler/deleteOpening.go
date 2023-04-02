package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// DeleteOpeningHandler is...
func DeleteOpeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"msg": "Delete Opening"})
}
