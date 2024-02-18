package v0

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jgfranco17/algorithm-api/service/pkg/handlers"
)

func newAlgorithmRoute(algorithm string, params string) string {
	return fmt.Sprintf("algorithms/%s/%s", algorithm, params)
}

// Adds v0 routes to the router.
func SetRoutes(route *gin.Engine) {
	v0 := route.Group("/v0")
	{
		v0.GET(newAlgorithmRoute("fibonacci", ":number"), withErrorHandling(handlers.FibonacciHandler()))
	}
}
