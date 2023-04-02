package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ListOpeningsHandler ...
func ListOpeningsHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"msg": "List Openings"})
}
