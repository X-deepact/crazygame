package web

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

func createBrowserContext(ctx context.Context, userAgent string) (context.Context, context.CancelFunc) {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.UserAgent(userAgent),
	)
	return chromedp.NewExecAllocator(ctx, opts...)
}

type ScraperConfig struct {
	Timeout      time.Duration
	MaxRetries   int
	RetryDelay   time.Duration
	UserAgent    string
	WaitSelector string
	WaitTimeout  time.Duration
}

type Scraper struct {
	ctx    context.Context
	config ScraperConfig
}

func NewScraper(ctx context.Context, config ScraperConfig) *Scraper {
	if config.Timeout == 0 {
		config.Timeout = 10 * time.Second
	}
	if config.MaxRetries == 0 {
		config.MaxRetries = 3
	}
	if config.RetryDelay == 0 {
		config.RetryDelay = 1 * time.Second
	}
	if config.WaitSelector == "" {
		config.WaitSelector = "body"
	}
	if config.WaitTimeout == 0 {
		config.WaitTimeout = 5 * time.Second
	}

	return &Scraper{
		ctx:    ctx,
		config: config,
	}
}

func (s *Scraper) CrawlPage(url string) (string, error) {
	var htmlContent string
	var lastError error

	for attempt := 0; attempt <= s.config.MaxRetries; attempt++ {
		browserCtx, cancelBrowser := createBrowserContext(s.ctx, s.config.UserAgent)
		pageCtx, cancelPage := context.WithTimeout(browserCtx, s.config.Timeout)
		defer cancelPage()
		defer cancelBrowser()

		startTime := time.Now()
		err := chromedp.Run(pageCtx,
			chromedp.Navigate(url),
			chromedp.WaitVisible(s.config.WaitSelector, chromedp.ByQuery),
			chromedp.OuterHTML("html", &htmlContent),
		)

		if err == nil {
			duration := time.Since(startTime)
			log.Printf("Successfully crawled page %s in %v, HTML length: %d",
				url, duration, len(htmlContent))
			return htmlContent, nil
		}

		lastError = err
		log.Printf("Attempt %d/%d failed for %s: %v",
			attempt+1, s.config.MaxRetries+1, url, err)

		if attempt < s.config.MaxRetries {
			time.Sleep(s.config.RetryDelay)
		}
	}

	log.Printf("Failed to crawl page %s after %d attempts: %v",
		url, s.config.MaxRetries+1, lastError)
	return "", fmt.Errorf("failed after %d attempts: %w",
		s.config.MaxRetries+1, lastError)
}
