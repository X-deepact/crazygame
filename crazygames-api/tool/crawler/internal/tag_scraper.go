package crawler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"crazygames.io/tool/crawler/internal/domain"
	"crazygames.io/tool/crawler/internal/repository"
	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
)

type TagScraper struct {
	ctx context.Context
}

type TagExtractor struct {
	ctx context.Context
}

func NewTagExtractor(ctx context.Context) *TagExtractor {
	return &TagExtractor{ctx: ctx}
}

func (e *TagExtractor) ExtractTags(url string) ([]domain.TagGroup, error) {
	log.Printf("Starting tag extraction from: %s", url)
	startTime := time.Now()

	// Load the page and extract HTML
	var html string
	err := chromedp.Run(e.ctx,
		chromedp.Navigate(url),
		chromedp.WaitVisible("body", chromedp.ByQuery),
		chromedp.OuterHTML("html", &html),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to load page: %w", err)
	}

	// Parse the HTML to extract tag groups and tags
	tagGroups, err := parseTagGroups(html)
	if err != nil {
		return nil, fmt.Errorf("failed to parse tag groups: %w", err)
	}

	// Save the tags data to a CSV file
	if err := repository.WriteTagsCSV(tagGroups); err != nil {
		return nil, fmt.Errorf("failed to save tags data to CSV: %w", err)
	}

	duration := time.Since(startTime)
	log.Printf("Successfully extracted %d tag groups in %v", len(tagGroups), duration)
	return tagGroups, nil
}

func NewTagScraper(ctx context.Context) *TagScraper {
	return &TagScraper{ctx: ctx}
}

func (s *TagScraper) ScrapeTags(url string) error {
	log.Printf("Starting tag scraping from: %s", url)
	startTime := time.Now()

	// Load the page and extract HTML
	var html string
	err := chromedp.Run(s.ctx,
		chromedp.Navigate(url),
		chromedp.WaitVisible("body", chromedp.ByQuery),
		chromedp.OuterHTML("html", &html),
	)
	if err != nil {
		return fmt.Errorf("failed to load page: %w", err)
	}

	// Parse the HTML to extract tag groups and tags
	tagGroups, err := parseTagGroups(html)
	if err != nil {
		return fmt.Errorf("failed to parse tag groups: %w", err)
	}

	// Save the tags data to a JSON file
	if err := saveTagsToFile(tagGroups); err != nil {
		return fmt.Errorf("failed to save tags data: %w", err)
	}

	duration := time.Since(startTime)
	log.Printf("Successfully scraped %d tag groups in %v", len(tagGroups), duration)
	return nil
}

func parseTagGroups(html string) ([]domain.TagGroup, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return nil, fmt.Errorf("failed to parse HTML: %w", err)
	}

	var groups []domain.TagGroup

	// Find all group headings
	doc.Find("h2.css-9oxgqm").Each(func(i int, s *goquery.Selection) {
		group := domain.TagGroup{
			Group: strings.TrimSpace(s.Text()),
		}

		// Find tags within this group
		s.NextUntil("h2.css-9oxgqm").Find("div.css-wy93c2").Each(func(i int, tagDiv *goquery.Selection) {
			name := tagDiv.Find("p").Text()
			count := tagDiv.Find("span").Text()
			url, _ := tagDiv.Parent().Attr("href")

			group.Tags = append(group.Tags, domain.TagData{
				Name:  strings.TrimSpace(name),
				Count: strings.TrimSpace(count),
				URL:   strings.TrimSpace(url),
			})
		})

		groups = append(groups, group)
	})

	return groups, nil
}

func saveTagsToFile(tagGroups []domain.TagGroup) error {
	const fileName = "tags.json"

	// Convert tag groups to JSON
	jsonData, err := json.MarshalIndent(tagGroups, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal tags data: %w", err)
	}

	// Write JSON to file
	if err := os.WriteFile(fileName, jsonData, 0644); err != nil {
		return fmt.Errorf("failed to write tags data to file: %w", err)
	}

	log.Printf("Saved tags data to %s", fileName)
	return nil
}
