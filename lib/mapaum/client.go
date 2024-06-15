package mapaum

import (
	"context"
	"fmt"
	"log/slog"
	"strconv"
	"strings"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	"github.com/kn-kraken/hackwarsaw-fintech/lib/utils"
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

func (r *Client) ListRealEstates() (chan (RealEstate), error) {
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
      chromedp.WaitVisible(".ui-paginator-page.ui-state-active:nth-child(1)"),
		)
		if err != nil {
			return nil, fmt.Errorf("skipping to first page: %w", err)
		}
	}

	channel := make(chan (RealEstate), 1)

	impl := func() {
		for {

      err = chromedp.Run(
        r.ctx,
        chromedp.Click(".ui-paginator-first"),
        chromedp.WaitVisible(".ui-paginator-page.ui-state-active:nth-child(1)"),
      )
      if err != nil {
        slog.Error("waiting for page to load", "error", err)
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
				slog.Error("reading table rows", "error", err)
			}

			for i := range streets {
				var err error
				var realEstate RealEstate
				realEstate.Address = streets[i].Children[0].NodeValue
				realEstate.OccuanceType = occurenceTypes[i].Children[0].NodeValue

				for _, destination := range destinations[i].Children[0].Children {
					realEstate.Destinations = append(realEstate.Destinations, destination.NodeValue)
				}
				realEstate.Area, err = parseFloat(areas[i].Children[0].NodeValue)
				if err != nil {
					slog.Error("parsing area", err)
					continue
				}
				realEstate.InitialPrice, err = parseFloat(initialPrices[i].Children[0].NodeValue)
				if err != nil {
					slog.Error("parsing initial price", err)
					continue
				}
				realEstate.District = districts[i].Children[0].NodeValue
				channel <- realEstate
			}

			var lastClass string
			var ok bool
			err = chromedp.Run(
				r.ctx,
				chromedp.WaitReady(".ui-paginator-last"),
				chromedp.AttributeValue(".ui-paginator-last", "class", &lastClass, &ok),
			)
			if err != nil {
				slog.Error("checking for last page", err)
			} else if !ok {
				slog.Error("checking for last page: not ok")
			}

			lastClasses := strings.Split(lastClass, " ")
			isDisabled := utils.Any(lastClasses, func(cls string) bool {
				return cls == "ui-state-disabled"
			})

			if isDisabled {
				close(channel)
				break
			}

			err = chromedp.Run(
				r.ctx,
				chromedp.Click(".ui-paginator-next"),
			)
		}
	}

	go impl()

	return channel, nil
}

func parseFloat(text string) (float32, error) {
	text = strings.ReplaceAll(text, ",", ".")
	result, err := strconv.ParseFloat(text, 32)
	if err != nil {
		return 0, err
	}
	return float32(result), nil
}
