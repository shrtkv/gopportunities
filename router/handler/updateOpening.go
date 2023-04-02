package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// UpdateOpeningHandler ...
func UpdateOpeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"msg": "Update Opening"})
}
