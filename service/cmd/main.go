// main.go
package main

import (
	"flag"

	"github.com/jgfranco17/algorithm-api/service/pkg/router"
)

var (
	port = flag.Int("port", 8080, "Port to listen on")
	mode = flag.String("mode", "dev", "Server run mode")
)

func main() {
	flag.Parse()
	server := router.CreateServer(0, *port)
	server.Run()
}
