// main.go
package main

import (
	"context"
	"flag"

	"github.com/gin-gonic/gin"
	"github.com/jgfranco17/algorithm-api/service/pkg/logging"
	"github.com/jgfranco17/algorithm-api/service/pkg/router"
)

var (
	port = flag.Int("port", 8080, "Port to listen on")
	mode = flag.String("mode", "dev", "Server run mode")
)

func main() {
	flag.Parse()
	ctx := context.WithValue(context.Background(), "section", "Main")
	log := logging.GetLogger(ctx)
	switch *mode {
	case "dev", "prod":
		log.Infof("Running API %s server on port %d", *mode, *port)
	default:
		log.Fatalf("Invalid mode provided, '%s' is not recognized.", *mode)
	}

	if *mode == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}
	server, serverCreateErr := router.CreateServer(0, *port)
	if serverCreateErr != nil {
		log.Fatalf("Failed to create server: %v", serverCreateErr)
	}
	server.Run()
}
