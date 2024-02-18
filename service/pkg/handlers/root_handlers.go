package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to the Algorithm API page!",
	})
}

func AboutHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"author": "Chino Franco",
		"title":  "Algorithm API",
		"repo":   "https://github.com/jgfranco17/algorithm-api",
	})
}

func NotFoundHandler(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"error": "Endpoint not found",
	})
}
