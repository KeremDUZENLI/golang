package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

type EtfInfo struct {
	Title              string
	Replication        string
	Earnings           string
	TotalExpenseRatio  string
	TrackingDifference string
	FundSize           string
}

func main() {
	isins := []string{"IE00B1XNHC34", "IE00B4L5Y983", "LU1838002480"}
	etfInfo := EtfInfo{}
	etfInfos := make([]EtfInfo, 0, 1)

	c := colly.NewCollector(colly.AllowedDomains("www.trackingdifferences.com", "trackingdifferences.com"))

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Accept-Language", "en-US;q=0.9")
		str := fmt.Sprintf("Visiting %s", r.URL)
		fmt.Println(str)
	})

	c.OnHTML("h1.page-title", func(h *colly.HTMLElement) {
		etfInfo.Title = h.Text
	})

	c.OnHTML("div.descfloat p.desc", func(h *colly.HTMLElement) {
		selection := h.DOM

		childNodes := selection.Children().Nodes

		if len(childNodes) == 3 {
			description := cleanDesc(selection.Find("span.desctitle").Text())
			value := selection.FindNodes(childNodes[2]).Text()

			switch description {
			case "Replication":
				etfInfo.Replication = value
			case "Earnings":
				etfInfo.Earnings = value
			case "TER":
				etfInfo.TotalExpenseRatio = value
			case "TD":
				etfInfo.TrackingDifference = value
			case "Fund size":
				etfInfo.FundSize = value
			}
		}
	})

	c.OnScraped(func(r *colly.Response) {
		etfInfos = append(etfInfos, etfInfo)
		etfInfo = EtfInfo{}
	})

	c.OnError(func(r *colly.Response, e error) {
		fmt.Printf("Error while scraping: %s\n", e.Error())
	})

	for _, isin := range isins {
		c.Visit(scrapeUrl(isin))
	}

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", " ")
	enc.Encode(etfInfos)
}

func cleanDesc(s string) string {
	return strings.TrimSpace(s)
}

func scrapeUrl(isin string) string {
	return "https://www.trackingdifferences.com/ETF/ISIN/" + isin
}
