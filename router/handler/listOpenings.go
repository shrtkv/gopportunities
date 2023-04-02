package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shrtkv/gopportunities/schemas"
)

// ListOpeningsHandler ...
func ListOpeningsHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}
	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing openings")
	}
	sendSuccess(ctx, "list-openings", openings)
}
