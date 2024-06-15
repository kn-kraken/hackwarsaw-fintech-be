package main

import (
	"log/slog"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kn-kraken/hackwarsaw-fintech/lib/api"
	"github.com/kn-kraken/hackwarsaw-fintech/lib/db"
)

func main() {
	db, err := db.New()
	if err != nil {
		slog.Error("creating database", "error", err)
		os.Exit(1)
	}
	defer db.Close()

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowCredentials: true,
	}))
	api.AddRoutes(r, db)
	// Listen and Server in 0.0.0.0:8080
	r.Run("0.0.0.0:8080")
}
