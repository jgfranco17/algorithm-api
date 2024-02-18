package system

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jgfranco17/algorithm-api/service/pkg/data"
)

func HomeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to the Algorithm API page!",
	})
}

func AboutHandler(c *gin.Context) {
	// Send the parsed JSON data as a response
	c.JSON(http.StatusOK, data.AboutInfo{
		Name:       "Algorithm API",
		Author:     "Joaquin Franco",
		Repository: "https://github.com/jgfranco17/algorithm-api",
		Version:    "0.0.1",
		License:    "MIT",
		Languages:  []string{"Go"},
		Algorithms: map[string][]string{
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
	c.JSON(http.StatusOK, data.HealthStatus{
		Timestamp: time.Now().Format("Mon Jan 2 15:04:05 MST 2006"),
		Status:    "healthy",
	})
}

func NotFoundHandler(c *gin.Context) {
	c.JSON(http.StatusNotFound, data.BasicErrorInfo{
		StatusCode: http.StatusNotFound,
		Endpoint:   c.Request.URL.Path,
		Message:    "Endpoint does not exist",
	})
}
