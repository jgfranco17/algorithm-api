package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	msa "github.com/jgfranco17/algorithm-api/core/pkg/array/maxsubarray"
	ts "github.com/jgfranco17/algorithm-api/core/pkg/array/twosum"
	fib "github.com/jgfranco17/algorithm-api/core/pkg/fibonacci"
	pal "github.com/jgfranco17/algorithm-api/core/pkg/palindrome"
	"github.com/jgfranco17/algorithm-api/core/pkg/utils"
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

func FibonacciHandler(c *gin.Context) {
	num, _ := strconv.Atoi(c.Param("number"))
	sequence := fib.Fibonacci(num)

	c.JSON(http.StatusOK, sequence)
}

func TwoSumHandler(c *gin.Context) {
	target, _ := strconv.Atoi(c.Param("target"))
	numList := c.QueryArray("nums")
	numbers, parseErr := utils.StringToIntArray(numList)
	if parseErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid number format"})
		return
	}

	result := ts.TwoSum(numbers, target)
	c.JSON(http.StatusOK, result)
}

func PalindromeHandler(c *gin.Context) {
	word := c.Param("word")
	check := pal.PalindromeCheck(word)

	c.JSON(http.StatusOK, check)
}

func MaxSubArrayHandler(c *gin.Context) {
	numList := c.QueryArray("nums")
	numbers, parseErr := utils.StringToIntArray(numList)
	if parseErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid number format"})
		return
	}
	result := msa.MaxSubArray(numbers)
	c.JSON(http.StatusOK, gin.H{
		"sequence": numbers,
		"subarray": result,
	})
}
