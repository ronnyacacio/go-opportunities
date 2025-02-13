package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListOpeningsHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "GET openings",
	})
}

func CreateOpeningsHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "POST openings",
	})
}
