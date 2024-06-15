package main

import (
	"log/slog"

	"github.com/kn-kraken/hackwarsaw-fintech/lib/mapaum"
)

func main() {
	client, err := mapaum.New()
	if err != nil {
		slog.Error("creating scrapper", err)
	}
	_ = client

	_, err = client.ListRealEstates()
	if err != nil {
		slog.Error("listing real estates", "error", err)
	}

	// _ = dom
	// println(dom.Text())
}
