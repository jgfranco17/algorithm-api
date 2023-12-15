package router

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jgfranco17/algorithm-api/service/pkg/handlers"
)

func CreateRouter() *gin.Engine {
	router := gin.Default()

	// Define routes
	router.GET("/", handlers.HomeHandler)
	router.GET("/algorithms/endpoint/:id", func(c *gin.Context) {
		// Get the value of the "id" parameter from the path
		id := c.Param("id")

		// Process the input and create a response
		response := fmt.Sprintf("Received ID: %s", id)

		// Send the response
		c.JSON(http.StatusOK, gin.H{"message": response})
	})

	return router
}
