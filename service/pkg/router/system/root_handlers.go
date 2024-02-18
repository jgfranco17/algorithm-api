package system

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func HomeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to the Algorithm API page!",
	})
}

func AboutHandler(c *gin.Context) {
	// Send the parsed JSON data as a response
	c.JSON(http.StatusOK, gin.H{
		"name":       "Algorithm API",
		"author":     "Joaquin Franco",
		"repository": "https://github.com/jgfranco17/algorithm-api",
		"license":    "MIT",
		"languages":  []string{"Go"},
		"version":    "0.0.1",
		"algorithms": map[string][]string{
			"array": {
				"MaxSubArray",
				"TwoSum",
			},
			"sequence": {
				"LongestCommonSubsequence",
				"Fibonacci",
			},
			"palindrome": {
				"palindrome",
			},
		},
	})
}

func HealthCheckHandler(c *gin.Context) {
	c.JSON(http.StatusOK, HealthStatus{
		Timestamp: time.Now().Format("Mon Jan 2 15:04:05 MST 2006"),
		Status:    "healthy",
	})
}

func NotFoundHandler(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"error": "Endpoint not found",
	})
}
