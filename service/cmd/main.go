// main.go
package main

import (
	"github.com/jgfranco17/algorithm-api/service/pkg/router"
)

func main() {
	r := router.CreateRouter()

	port := ":8080"
	r.Run(port)
}
