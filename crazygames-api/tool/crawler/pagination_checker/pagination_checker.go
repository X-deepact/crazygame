package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	crawler "crazygames.io/tool/crawler/internal"
	"crazygames.io/tool/crawler/internal/web"
	"github.com/chromedp/chromedp"
)

func main() {
	// Parse command line flags
	url := flag.String("url", "", "URL to check pagination")
	maxPages := flag.Int("max-pages", 10, "Maximum number of pages to check")
	flag.Parse()

	if *url == "" {
		log.Fatal("URL is required. Use -url flag")
	}

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Minute)
	defer cancel()

	// Handle interrupt signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigChan
		log.Println("Received interrupt signal. Shutting down gracefully...")
		cancel()
		os.Exit(0)
	}()

	// Set base URL for the web package
	web.SetBaseURL(*url)

	// Initialize crawler with single URL config
	startURLs := []crawler.URLConfig{
		{
			URL:      *url,
			Paginate: true,
		},
	}

	// Initialize crawler
	c := crawler.NewCrawler(ctx, startURLs)
	defer c.Close()

	log.Printf("Starting pagination check for: %s", *url)
	log.Printf("Will check up to %d pages", *maxPages)

	// Initialize counters
	pageCount := 1
	totalGames := 0

	for {
		if pageCount > *maxPages {
			log.Printf("Reached maximum configured pages: %d", *maxPages)
			break
		}

		currentURL := *url
		if pageCount > 1 {
			currentURL = fmt.Sprintf("%s/%d", *url, pageCount)
		}

		log.Printf("Checking page %d: %s", pageCount, currentURL)

		// Create a timeout context for this page
		pageCtx, pageCancel := context.WithTimeout(c.BrowserContext(), 30*time.Second)
		defer pageCancel()

		// Navigate to the current page and wait for content to load
		err := chromedp.Run(pageCtx,
			chromedp.Navigate(currentURL),
			// Wait for any game link to appear
			chromedp.WaitVisible(`a[href*="/game/"]`, chromedp.ByQuery),
		)
		if err != nil {
			log.Printf("Failed to load page %d: %v", pageCount, err)
			break
		}

		// Find all game links on the page with timeout
		var gameLinks []string
		err = chromedp.Run(pageCtx,
			chromedp.Evaluate(`Array.from(document.querySelectorAll('a[href*="/game/"]')).map(a => a.href)`, &gameLinks),
		)
		if err != nil {
			log.Printf("Failed to extract game links from page %d: %v", pageCount, err)
			break
		}

		// Remove duplicates from gameLinks
		seen := make(map[string]bool)
		uniqueLinks := []string{}
		for _, link := range gameLinks {
			if !seen[link] {
				seen[link] = true
				uniqueLinks = append(uniqueLinks, link)
			}
		}

		if len(uniqueLinks) == 0 {
			// Try one more time with a longer wait
			chromedp.Run(pageCtx,
				chromedp.Sleep(5*time.Second),
				chromedp.Evaluate(`Array.from(document.querySelectorAll('a[href*="/game/"]')).map(a => a.href)`, &gameLinks),
			)

			// Remove duplicates again
			seen = make(map[string]bool)
			uniqueLinks = []string{}
			for _, link := range gameLinks {
				if !seen[link] {
					seen[link] = true
					uniqueLinks = append(uniqueLinks, link)
				}
			}

			if len(uniqueLinks) == 0 {
				log.Printf("No games found on page %d after retry. Last page appears to be %d", pageCount, pageCount-1)
				break
			}
		}

		totalGames += len(uniqueLinks)
		log.Printf("Page %d: Found %d unique games (Total: %d)", pageCount, len(uniqueLinks), totalGames)

		pageCount++
	}

	log.Printf("Pagination check complete:")
	log.Printf("Total pages: %d", pageCount-1)
	log.Printf("Total games: %d", totalGames)
	log.Printf("Average games per page: %.2f", float64(totalGames)/float64(pageCount-1))
}
