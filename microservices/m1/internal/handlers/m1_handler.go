package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func M1Handler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"service": "M1", "status": "OK"})
}
