package mapaum

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

type Client struct {
	client http.Client
	jar    http.CookieJar
}

func New() (*Client, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}
	client := &Client{
		client: http.Client{},
		jar:    jar,
	}
	return client, nil
}

func (r *Client) Init() error {
	resp, err := r.client.Get("https://mapa.um.warszawa.pl/mapaApp1/mapa?service=nieruchomosci")

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("status %v", resp.StatusCode)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	bodyString := string(bodyBytes)
	log.Println(bodyString)

	return nil
}

func (r *Client) ListRealEstates(page int) ([]RealEstate, error) {
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
	form.Add("javax.faces.ViewState", "6458804742846979572:2467599149130526145")
	resp, err := r.client.PostForm("https://mapa.um.warszawa.pl/mapaApp1/faces/oferty/ofertySprzedaz.xhtml", form)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status %v", resp.StatusCode)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	bodyString := string(bodyBytes)
	log.Println(bodyString)

	return []RealEstate{{}}, nil
}
