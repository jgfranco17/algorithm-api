package router

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/jgfranco17/algorithm-api/service/pkg/handlers"
	"github.com/jgfranco17/algorithm-api/service/pkg/logging"
)

type Server struct {
	Router *gin.Engine
	Port   string
}

func CreateServer(port int) *Server {
	router := gin.Default()
	router.GET("/", handlers.HomeHandler)
	router.GET("/algorithms/fibonacci/:number", handlers.FibonacciHandler)
	router.GET("/algorithms/twosum/:numbers/:target", handlers.TwoSumHandler)

	return &Server{
		Router: router,
		Port:   fmt.Sprintf(":%d", port),
	}
}

func (s *Server) Run() error {
	log := logging.GetLogger()
	log.Infof("Starting Algorithm API server!")
	s.Router.Run(s.Port)
	return nil
}
