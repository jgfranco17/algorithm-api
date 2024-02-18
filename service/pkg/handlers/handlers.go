package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	msa "github.com/jgfranco17/algorithm-api/core/pkg/array/maxsubarray"
	pal "github.com/jgfranco17/algorithm-api/core/pkg/palindrome"
	utils "github.com/jgfranco17/algorithm-api/core/pkg/utils"
)

func PalindromeHandler(c *gin.Context) {
	word := c.Param("word")
	palindromeCheck := pal.PalindromeCheck(word)

	c.JSON(http.StatusOK, gin.H{
		"word":   word,
		"result": palindromeCheck,
	})
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
