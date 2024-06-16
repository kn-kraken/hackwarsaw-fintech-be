package main

import (
	"flag"
	"log/slog"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kn-kraken/hackwarsaw-fintech/lib/api"
	"github.com/kn-kraken/hackwarsaw-fintech/lib/db"
)

var (
	dbhost     string
	dbport     int
	dbname     string
	dbuser     string
	dbpassword string
)

func main() {
	flag.StringVar(&dbhost, "dbhost", "localhost", "db host")
	flag.IntVar(&dbport, "dbport", 5432, "db port")
	flag.StringVar(&dbname, "dbname", "gis", "db name")
	flag.StringVar(&dbuser, "dbuser", "gisuser", "db user")
	flag.StringVar(&dbpassword, "dbpassword", "gispassword", "db password")
	flag.Parse()

	db, err := db.New(
		dbhost,
		dbport,
		dbname,
		dbuser,
		dbpassword,
	)
	if err != nil {
		slog.Error("creating database", "error", err)
		os.Exit(1)
	}
	defer db.Close()

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	api.AddRoutes(r, db)
	// Listen and Server in 0.0.0.0:8080
	r.Run("0.0.0.0:8080")
}
