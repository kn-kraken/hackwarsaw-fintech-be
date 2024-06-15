package main

import (
	"log/slog"

	"github.com/kn-kraken/hackwarsaw-fintech/lib/db"
	"github.com/kn-kraken/hackwarsaw-fintech/lib/mapaum"
)

func main() {
  db, err := db.New()
	if err != nil {
		slog.Error("creating db", err)
	}

	client, err := mapaum.New()
	if err != nil {
		slog.Error("creating scrapper", err)
	}
	_ = client

  channel, err := client.ListRealEstates()
	if err != nil {
		slog.Error("listing real estates", "error", err)
	}

  for realEstate := range channel {
    err = db.AddRealEstate(&realEstate)  
    if err != nil {
      slog.Error("saving real estate", "error", err)
    }
    // fmt.Printf("%#v", realEstate)
  }

	// _ = dom
	// println(dom.Text())
}
