package main

import (
	"flag"
	"log"
	"log/slog"

	"github.com/kelvins/geocoder"
	"github.com/kn-kraken/hackwarsaw-fintech/lib/db"
	"github.com/kn-kraken/hackwarsaw-fintech/lib/mapaum"
)

var (
	dbhost          string
	dbport          int
	dbname          string
	dbuser          string
	dbpassword      string
	geocodingApiKey string
)

func main() {
	flag.StringVar(&dbhost, "dbhost", "localhost", "db host")
	flag.IntVar(&dbport, "dbport", 5432, "db port")
	flag.StringVar(&dbname, "dbname", "gis", "db name")
	flag.StringVar(&dbuser, "dbuser", "gisuser", "db user")
	flag.StringVar(&dbpassword, "dbpassword", "gispassword", "db password")
	flag.StringVar(&geocodingApiKey, "geocoding-apikey", "", "Google's Geocoding API key")
	flag.Parse()

	if geocodingApiKey == "" {
		log.Fatal("required parameter -geocoding-apikey not set")
	}
	geocoder.ApiKey = geocodingApiKey

	db, err := db.New(
		dbhost,
		dbport,
		dbname,
		dbuser,
		dbpassword,
	)
	if err != nil {
		slog.Error("creating db", err)
	}

	client, err := mapaum.New()
	if err != nil {
		slog.Error("creating scrapper", err)
	}

	channel, err := client.ListRealEstates()
	if err != nil {
		slog.Error("listing real estates", "error", err)
	}

	for realEstate := range channel {
		err = db.AddRealEstate(&realEstate)
		if err != nil {
			slog.Error("saving real estate", "error", err)
		}
	}
}
