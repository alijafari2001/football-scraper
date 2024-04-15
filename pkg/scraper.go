package pkg

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func Scrape(url string) (*goquery.Document, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("http error %d: %s", res.StatusCode, res.Status)
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	return doc, err
}
