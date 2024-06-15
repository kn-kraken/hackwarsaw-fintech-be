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

  dom, err := client.ListRealEstates(0, 100)
	if err != nil {
		slog.Error("listing real estates", "error", err)
	}

  println(dom.Text())
}
