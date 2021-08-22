package handlers

import (
	"github.com/gin-gonic/gin"
)

func DefaultHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hey Go URL Shortener !",
	})
}
