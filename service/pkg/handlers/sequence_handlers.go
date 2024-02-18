package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	fib "github.com/jgfranco17/algorithm-api/core/pkg/fibonacci"
)

func FibonacciHandler() func(c *gin.Context) error {
	return func(c *gin.Context) error {
		num, _ := strconv.Atoi(c.Param("number"))
		sequence, err := fib.Fibonacci(num)
		if err != nil {
			return err
		}

		c.JSON(http.StatusOK, gin.H{
			"count":    num,
			"sequence": sequence,
		})
		return nil
	}
}
