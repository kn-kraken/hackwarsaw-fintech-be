package mapaum

import (
	"context"
	"fmt"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

const (
	URL = "https://mapa.um.warszawa.pl/mapaApp1/faces/oferty/ofertyWynajem.xhtml?lang=pl"
)

type Client struct {
	ctx    context.Context
	cancel context.CancelFunc
}

func New() (*Client, error) {
	ctx, cancel := chromedp.NewContext(context.Background())

	err := chromedp.Run(
		ctx,
		chromedp.Navigate(URL),
		chromedp.WaitReady("input[name=javax\\.faces\\.ViewState]"),
	)
	if err != nil {
		return nil, err
	}

	client := &Client{ctx, cancel}
	return client, nil
}

func (r *Client) Close() {
	r.cancel()
}

func (r *Client) ListRealEstates() (*goquery.Selection, error) {
	var currentPage string

	err := chromedp.Run(
		r.ctx,
		chromedp.WaitVisible(".ui-paginator-page.ui-state-active"),
		chromedp.Text(".ui-paginator-page.ui-state-active", &currentPage),
	)
	if err != nil {
		return nil, fmt.Errorf("reading page: %w", err)
	}

	if currentPage != "1" {
		err = chromedp.Run(
			r.ctx,
			chromedp.Click(".ui-paginator-first"),
			chromedp.Value(".ui-paginator-page.ui-state-active", &currentPage),
		)
		if err != nil {
			return nil, fmt.Errorf("skipping to first page: %w", err)
		}
	}

	var streets []*cdp.Node
	var occurenceTypes []*cdp.Node
	var destinations []*cdp.Node
	var areas []*cdp.Node
	var initialPrices []*cdp.Node
	var districts []*cdp.Node
	err = chromedp.Run(
		r.ctx,
    chromedp.Nodes("tr.ui-widget-content>td:nth-child(1)", &streets),
    chromedp.Nodes("tr.ui-widget-content>td:nth-child(2)", &occurenceTypes),
    chromedp.Nodes("tr.ui-widget-content>td:nth-child(3) li.ui-datalist-item", &destinations),
    chromedp.Nodes("tr.ui-widget-content>td:nth-child(4)", &areas),
    chromedp.Nodes("tr.ui-widget-content>td:nth-child(5)", &initialPrices),
    chromedp.Nodes("tr.ui-widget-content>td:nth-child(6)", &districts),
	)
	if err != nil {
		return nil, fmt.Errorf("reading table rows: %w", err)
	}

  println(streets[0].Children[0].NodeValue)
  println(occurenceTypes[0].Children[0].NodeValue)
  println(destinations[0].Children[0].Children[0].NodeValue)
  println(areas[0].Children[0].NodeValue)
  println(initialPrices[0].Children[0].NodeValue)
  println(districts[0].Children[0].NodeValue)

	return nil, nil
}
