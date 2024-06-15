package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kn-kraken/hackwarsaw-fintech/lib/api"
)

func main() {
	r := gin.Default()

	api.AddRoutes(r)
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
