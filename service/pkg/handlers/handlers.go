package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	fib "github.com/jgfranco17/algorithm-api/core/pkg/fibonacci"
)

func HomeHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Welcome to the home page!",
	})
}

func FibonacciHandler(c *gin.Context) {
	num, _ := strconv.Atoi(c.Param("number"))
	sequence := fib.Fibonacci(num)

	c.JSON(http.StatusOK, sequence)
}
