package web

import (
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func isValidGameURL(url string) bool {
	// Basic validation for game URLs
	return strings.Contains(url, "/game/") &&
		(strings.HasPrefix(url, "http://") ||
			strings.HasPrefix(url, "https://"))
}

var baseURL string

func SetBaseURL(url string) {
	baseURL = url
}

func GetBaseURL() string {
	return baseURL
}

func normalizeURL(url string) string {
	// If URL already has protocol, return as-is
	if strings.HasPrefix(url, "http") {
		return url
	}
	
	// Prepend base URL
	return baseURL + url
}

func ExtractGameURLs(html string) ([]string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return nil, fmt.Errorf("failed to parse HTML: %w", err)
	}

	var urls []string
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists {
			normalized := normalizeURL(href)
			if isValidGameURL(normalized) {
				urls = append(urls, normalized)
			}
		}
	})

	if len(urls) == 0 {
		return nil, fmt.Errorf("no valid game URLs found")
	}

	log.Printf("Found %d valid game URLs", len(urls))
	return urls, nil
}
