package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	msa "github.com/jgfranco17/algorithm-api/core/pkg/array/maxsubarray"
	ts "github.com/jgfranco17/algorithm-api/core/pkg/array/twosum"
	pal "github.com/jgfranco17/algorithm-api/core/pkg/palindrome"
	"github.com/jgfranco17/algorithm-api/core/pkg/utils"
)

func TwoSumHandler(c *gin.Context) {
	target, _ := strconv.Atoi(c.Param("target"))
	numList := c.QueryArray("nums")
	numbers, parseErr := utils.StringToIntArray(numList)
	if parseErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid number format"})
		return
	}

	found, indices := ts.TwoSum(numbers, target)
	c.JSON(http.StatusOK, gin.H{
		"found":  found,
		"result": indices,
	})
}

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
