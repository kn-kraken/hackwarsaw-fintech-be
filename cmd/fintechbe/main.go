package main

import (
	"fmt"
	"log"
	"net/url"

	"gopkg.in/headzoo/surf.v1"
)

func main() {
	bow := surf.NewBrowser()
	err := bow.Open("https://mapa.um.warszawa.pl/mapaApp1/faces/oferty/ofertySprzedaz.xhtml?lang=pl")
	if err != nil {
		panic(err)
	}

	var viewStateNode = bow.Dom().Find("input[name=javax\\.faces\\.ViewState]")

	viewState, exists := viewStateNode.First().Attr("value")
	if !exists {
		panic("viewstate not exists")
	}

	form := url.Values{}
	form.Add("javax.faces.partial.ajax", "true")
	form.Add("javax.faces.source", "form1:dataTable1")
	form.Add("javax.faces.partial.execute", "form1:dataTable1")
	form.Add("javax.faces.partial.render", "form1:dataTable1")
	form.Add("form1:dataTable1", "form1:dataTable1")
	form.Add("form1:dataTable1_pagination", "true")
	form.Add("form1:dataTable1_first", "0")
	form.Add("form1:dataTable1_rows", "10")
	form.Add("form1:dataTable1_encodeFeature", "true")
	form.Add("form1", "form1")
	form.Add("form1:dataTable:j_idt11:filter", "")
	form.Add("form1:dataTable:j_idt15:filter", "")
	form.Add("form1:dataTable:j_idt17:filter", "")
	form.Add("form1:dataTable:j_idt21:filter", "")
	form.Add("form1:dataTable:j_idt23:filter", "")
	form.Add("form1:dataTable1_selection", "8712")
	form.Add("form1:dataTable1_scrollState", "0,0")
	form.Add("javax.faces.ViewState", viewState)
	err = bow.PostForm("https://mapa.um.warszawa.pl/mapaApp1/faces/oferty/ofertySprzedaz.xhtml", form)
	if err != nil {
		log.Fatal(err)
	}

	// Outputs: "The Go Programming Language"
	// fmt.Println(bow.Body())

	err = bow.Open("https://mapa.um.warszawa.pl/mapaApp1/faces/oferty/ofertySprzedazSzczegoly.xhtml?myid=8692&city=true")
	if err != nil {
		panic(err)
	}
	fmt.Println(bow.Body())

	// options := &types.Options{
	// 	Headless:     true,
	// 	MaxDepth:     1,             // Maximum depth to crawl
	// 	FieldScope:   "rdn",         // Crawling Scope Field
	// 	BodyReadSize: math.MaxInt,   // Maximum response size to read
	// 	Timeout:      10,            // Timeout is the time to wait for request in seconds
	// 	Concurrency:  10,            // Concurrency is the number of concurrent crawling goroutines
	// 	Parallelism:  10,            // Parallelism is the number of urls processing goroutines
	// 	Delay:        0,             // Delay is the delay between each crawl requests in seconds
	// 	RateLimit:    150,           // Maximum requests to send per second
	// 	Strategy:     "depth-first", // Visit strategy (depth-first, breadth-first)
	// 	OnResult: func(result output.Result) { // Callback function to execute for result
	// 		gologger.Info().Msg(result.Request.URL)
	// 	},
	// }
	// crawlerOptions, err := types.NewCrawlerOptions(options)
	// if err != nil {
	// 	gologger.Fatal().Msg(err.Error())
	// }
	// defer crawlerOptions.Close()

	// crawler, err := standard.New(crawlerOptions)
	// if err != nil {
	// 	gologger.Fatal().Msg(err.Error())
	// }
	// defer crawler.Close()

	// var input = "https://mapa.um.warszawa.pl/mapaApp1/mapa?service=nieruchomosci&L=pl"
	// session, err := crawler.NewCrawlSessionWithURL(input)
	// if err != nil {
	// 	gologger.error().Msgf("Could not crawl %s: %s", input, err.Error())
	// }

	// err = crawler.Crawl(input)
	// if err != nil {
	// 	gologger.Warning().Msgf("Could not crawl %s: %s", input, err.Error())
	// }
	// client := &mapaum.Client{}
	// err := client.Init()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// _, err = client.ListRealEstates(1)
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
