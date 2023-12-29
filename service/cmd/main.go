package main

import (
	"context"
	"flag"

	"github.com/gin-gonic/gin"
	"github.com/jgfranco17/algorithm-api/service/pkg/logging"
	"github.com/jgfranco17/algorithm-api/service/pkg/router"
)

var (
	port    = flag.Int("port", 8080, "Port to listen on")
	devMode = flag.Bool("dev", true, "Run server in debug mode")
)

func main() {
	flag.Parse()
	ctx := context.WithValue(context.Background(), "section", "Main")
	log := logging.GetLogger(ctx)
	if *devMode {
		log.Infof("Running API server on port %d in dev mode", *port)
	} else {
		log.Infof("Running API production server on port %d", *port)
		gin.SetMode(gin.ReleaseMode)
	}
	server, serverCreateErr := router.CreateServer(0, *port)
	if serverCreateErr != nil {
		log.Fatalf("Failed to create server: %v", serverCreateErr)
	}
	server.Run()
}
