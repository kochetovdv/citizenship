package siteparser

import (
	"citizenship/internal/order"
	"citizenship/pkg/webservices"
	"fmt"
	"log"
	"path/filepath"
	"regexp"
	"strings"
)

// SiteParser is an structure for downloading orders from a website
type SiteParser struct {
	url  string
}

// NewSiteParser creates a new SiteParser
func NewSiteParser(url string) *SiteParser {
	s := new(SiteParser)
	s.url = url
	return s
}

// Parse parses the orders from the website
func (s *SiteParser) Parse(url string) *order.Orders {
	response, err := s.connect()
	if err != nil {
		log.Printf("Error with connecting: %s", err)
	}

	listOfOrders := s.extractData(response)

	return listOfOrders
}

// connect connects to the website and returns the response body
func (s *SiteParser) connect() ([]byte, error) {
	body, err:=webservices.GetResponse(s.url)
	if err != nil {
		return nil, fmt.Errorf("error with connecting: %s", err)
	}
	return body, nil
}

// extractData extracts the orders from the response
func (d *SiteParser) extractData(body []byte) *order.Orders {
	orders := order.NewOrders()

	html := string(body)
	htmlLines := strings.Split(html, "\n")
	var htmlLi []string

	for _, line := range htmlLines {
		if strings.Contains(line, "<li>") {
			htmlLi = append(htmlLi, line)
		}
	}

	for _, line := range htmlLi {
		dateRegExp := regexp.MustCompile(`<strong>([^<]+)</strong>`)

		date := dateRegExp.FindString(line)
		date = strings.Replace(date, "<strong>", "", -1)
		date = strings.Replace(date, "</strong>", "", -1)

		linksRegExp := regexp.MustCompile(`<a href="([^"]+)">([^<]+)</a>`)
		links := linksRegExp.FindAllStringSubmatch(line, -1)

		for _, link := range links {
			filename := filepath.Base(link[1])
			order := order.NewOrder(
				date,
				filename,
				link[1],
				link[2],
			)
			orders.Add(order)
		}
	}
	return orders
}
