package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type accountCreationRequest struct {
	Username string `json:"username" binding:"required"`
	Currency string `json:"currency" binding:"required,oneof=USD EUR"`
}

func (server *Server) createAccount(c *gin.Context) {
	var json accountCreationRequest
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"AA": json.Username})
	return
}
