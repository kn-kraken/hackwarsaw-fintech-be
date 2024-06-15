package main

import (
	"log"

	"github.com/kn-kraken/hackwarsaw-fintech/lib/mapaum"
)

func main() {
	client, err := mapaum.New()
	if err != nil {
		log.Fatal(err)
	}

  dom, err := client.ListRealEstates(0, 100)
	if err != nil {
		log.Fatal(err)
	}

  println(dom.Text())
}
