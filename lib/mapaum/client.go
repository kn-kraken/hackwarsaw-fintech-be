package mapaum

import (
	"fmt"
	"log"
	"net/url"

	"github.com/headzoo/surf"
	"github.com/headzoo/surf/browser"
)

type Client struct {
	bow       *browser.Browser
	viewState string
}

func New() (*Client, error) {
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

	client := &Client{bow: bow, viewState: viewState}
	return client, nil
}

func (r *Client) ListRealEstates(page int, pageSize int) ([]RealEstate, error) {
	form := url.Values{}
	form.Add("javax.faces.partial.ajax", "true")
	form.Add("javax.faces.source", "form1:dataTable1")
	form.Add("javax.faces.partial.execute", "form1:dataTable1")
	form.Add("javax.faces.partial.render", "form1:dataTable1")
	form.Add("form1:dataTable1", "form1:dataTable1")
	form.Add("form1:dataTable1_pagination", "true")
	form.Add("form1:dataTable1_first", fmt.Sprint(page*pageSize))
	form.Add("form1:dataTable1_rows", fmt.Sprint(pageSize))
	form.Add("form1:dataTable1_encodeFeature", "true")
	form.Add("form1", "form1")
	form.Add("form1:dataTable:j_idt11:filter", "")
	form.Add("form1:dataTable:j_idt15:filter", "")
	form.Add("form1:dataTable:j_idt17:filter", "")
	form.Add("form1:dataTable:j_idt21:filter", "")
	form.Add("form1:dataTable:j_idt23:filter", "")
	form.Add("form1:dataTable1_selection", "8712")
	form.Add("form1:dataTable1_scrollState", "0,0")
	form.Add("javax.faces.ViewState", r.viewState)

	err := r.bow.PostForm("https://mapa.um.warszawa.pl/mapaApp1/faces/oferty/ofertySprzedaz.xhtml", form)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(r.bow.Body())

	return []RealEstate{}, nil
}
