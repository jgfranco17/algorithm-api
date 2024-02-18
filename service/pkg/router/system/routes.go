package system_info

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var startTime time.Time

func init() {
	startTime = time.Now()
}

type healthStatus struct {
	Status string `json:"status"`
}

func SetSystemRoutes(route *gin.Engine) {
	route.GET("/health-check", func(c *gin.Context) {
		c.JSON(http.StatusOK, healthStatus{Status: "healthy"})
	})
}
