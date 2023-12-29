// main.go
package main

import (
	"github.com/jgfranco17/algorithm-api/service/pkg/router"
)

func main() {
	server := router.CreateServer(8080)
	server.Run()
}
