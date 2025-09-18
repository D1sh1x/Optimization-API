package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AddRequest struct {
	A float64 `json:"a"`
	B float64 `json:"b"`
}

type AddResponse struct {
	Result float64 `json:"result"`
}

func AddHandler(c *gin.Context) {
	var req AddRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := req.A + req.B

	c.JSON(http.StatusOK, AddResponse{Result: result})
}
