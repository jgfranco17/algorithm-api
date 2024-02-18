package system

import (
	"time"

	"github.com/gin-gonic/gin"
)

var startTime time.Time

func init() {
	startTime = time.Now()
}

type HealthStatus struct {
	Timestamp string `json:"timestamp"`
	Status    string `json:"status"`
}

func SetSystemRoutes(route *gin.Engine) {
	route.GET("/", HomeHandler)
	route.GET("/about", AboutHandler)
	route.GET("/health-check", HealthCheckHandler)
}
