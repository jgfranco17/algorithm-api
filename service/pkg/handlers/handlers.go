package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	fib "github.com/jgfranco17/algorithm-api/core/pkg/fibonacci"
	pal "github.com/jgfranco17/algorithm-api/core/pkg/palindrome"
	ts "github.com/jgfranco17/algorithm-api/core/pkg/twosum"
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

func FibonacciHandler(c *gin.Context) {
	num, _ := strconv.Atoi(c.Param("number"))
	sequence := fib.Fibonacci(num)

	c.JSON(http.StatusOK, sequence)
}

func TwoSumHandler(c *gin.Context) {
	// Get the "numbers" parameter from the URL
	numbersStr := c.Param("numbers")
	target, _ := strconv.Atoi(c.Param("target"))

	// Split the string by dashes
	numberStrings := strings.Split(numbersStr, "-")

	// Convert each string to an integer
	var numbers []int
	for _, numStr := range numberStrings {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid number format"})
			return
		}
		numbers = append(numbers, num)
	}

	result := ts.TwoSum(numbers, target)
	c.JSON(http.StatusOK, result)
}

func PalindromeHandler(c *gin.Context) {
	word := c.Param("word")
	check := pal.PalindromeCheck(word)

	c.JSON(http.StatusOK, check)
}
