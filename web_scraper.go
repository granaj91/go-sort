package main

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func getHackathonEvents(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	// Convert HTML into goquery document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	var events []string
	doc.Find(".event-name").Each(func(i int, s *goquery.Selection) {
		events = append(events, s.Text())
	})
	return events, nil
}
