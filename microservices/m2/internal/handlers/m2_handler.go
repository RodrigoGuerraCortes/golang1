package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func M2Handler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"service": "M2", "status": "OK"})
}
