package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	ts "github.com/jgfranco17/algorithm-api/core/pkg/array/twosum"
	utils "github.com/jgfranco17/algorithm-api/core/pkg/utils"
	error_handling "github.com/jgfranco17/algorithm-api/service/pkg/router/error_handling"
)

func TwoSumHandler(c *gin.Context) func(c *gin.Context) error {
	return func(c *gin.Context) error {
		target, _ := strconv.Atoi(c.Param("target"))
		numList := c.QueryArray("nums")
		numbers, err := utils.StringToIntArray(numList)
		if err != nil {
			error_handling.ServeError(c, http.StatusBadRequest, "Invalid number format", err)
			return err
		}

		found, indices := ts.TwoSum(numbers, target)
		c.JSON(http.StatusOK, gin.H{
			"found":  found,
			"result": indices,
		})
		return nil
	}
}
