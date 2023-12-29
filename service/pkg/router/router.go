package router

import (
	"github.com/gin-gonic/gin"

	"github.com/jgfranco17/algorithm-api/service/pkg/handlers"
)

func CreateRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", handlers.HomeHandler)
	router.GET("/algorithms/fibonacci/:number", handlers.FibonacciHandler)
	router.GET("/algorithms/twosum/:numbers/:target", handlers.TwoSumHandler)

	return router
}
