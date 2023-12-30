package router

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/jgfranco17/algorithm-api/service/pkg/handlers"
	"github.com/jgfranco17/algorithm-api/service/pkg/logging"
)

type Server struct {
	Router *gin.Engine
	Port   string
}

func newAlgorithmRoute(version int, algorithm string, params string) string {
	versionNumber := fmt.Sprintf("v%d", version)
	return fmt.Sprintf("%s/algorithms/%s/%s", versionNumber, algorithm, params)
}

func CreateServer(version int, port int) (*Server, error) {
	if version < 0 {
		return nil, fmt.Errorf("Version must be greater or equal to 0.")
	}
	router := gin.Default()
	router.GET("/", handlers.HomeHandler)
	router.GET("/about", handlers.AboutHandler)
	router.GET(newAlgorithmRoute(version, "fibonacci", ":number"), handlers.FibonacciHandler)
	router.GET(newAlgorithmRoute(version, "twosum", ":numbers/:target"), handlers.TwoSumHandler)
	router.GET(newAlgorithmRoute(version, "palindrome", ":word"), handlers.FibonacciHandler)
	router.GET(newAlgorithmRoute(version, "maxsubarray", ":numbers"), handlers.MaxSubArrayHandler)
	router.NoRoute(handlers.NotFoundHandler)

	return &Server{
		Router: router,
		Port:   fmt.Sprintf(":%d", port),
	}, nil
}

func (s *Server) Run() error {
	ctx := context.WithValue(context.Background(), "section", "ServerCreate")
	log := logging.GetLogger(ctx)
	log.Infof("Starting Algorithm API server!")
	s.Router.Run(s.Port)
	return nil
}
