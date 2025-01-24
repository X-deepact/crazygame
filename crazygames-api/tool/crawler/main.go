package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	crawler "crazygames.io/tool/crawler/internal"
	"crazygames.io/tool/crawler/internal/web"
)

func main() {
	// Parse command line flags
	tagsOnly := flag.Bool("tags", false, "Only scrape tags data")
	flag.Parse()

	// Create context with 48 hour timeout for long-running operations
	ctx, cancel := context.WithTimeout(context.Background(), 48*time.Hour)
	defer cancel()

	// Handle interrupt signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigChan
		log.Println("Received interrupt signal. Shutting down gracefully...")
		cancel()
		os.Exit(0) // Immediately exit the program
	}()

	// Define base URL
	baseURL := "https://www.crazygames.com"
	web.SetBaseURL(baseURL)

	if *tagsOnly {
		// Initialize crawler to get the properly configured browser context
		crawler := crawler.NewCrawler(ctx, nil) // Pass nil for startURLs since we're only scraping tags
		defer crawler.Close()

		// Run tag extractor using the crawler's browser context
		tagExtractor := crawler.NewTagExtractor(crawler.BrowserContext())
		tagsURL := "https://www.crazygames.com/tags"
		if _, err := tagExtractor.ExtractTags(tagsURL); err != nil {
			log.Fatalf("Failed to scrape tags: %v", err)
		}
		log.Println("Tag scraping complete")
		return
	}

	// Generate start URLs with configuration
	startURLs := []crawler.URLConfig{
		// {
		// 	URL:      baseURL,
		// 	Paginate: true, // Set to true if you want pagination for this URL
		// },
		// {
		// 	URL:      "https://www.crazygames.com/new",
		// 	Paginate: true,
		// },
		// {URL: "https://www.crazygames.com/originals", Paginate: true},
		// {URL: "https://www.crazygames.com/t/multiplayer", Paginate: true},
		// {URL: "https://www.crazygames.com/t/2-player", Paginate: true},
		// {URL: "https://www.crazygames.com/c/action", Paginate: true},
		// {URL: "https://www.crazygames.com/c/adventure", Paginate: true},
		// {URL: "https://www.crazygames.com/t/basketball", Paginate: true},
		// {URL: "https://www.crazygames.com/c/beauty", Paginate: true},
		// {URL: "https://www.crazygames.com/t/bike", Paginate: true},
		// {URL: "https://www.crazygames.com/t/car", Paginate: true},
		// {URL: "https://www.crazygames.com/t/card", Paginate: true},
		// {URL: "https://www.crazygames.com/c/casual", Paginate: true},
		// {URL: "https://www.crazygames.com/c/clicker", Paginate: true},
		// {URL: "https://www.crazygames.com/t/controller", Paginate: true},
		// {URL: "https://www.crazygames.com/t/dress-up", Paginate: true},
		// {URL: "https://www.crazygames.com/c/driving", Paginate: true},
		// {URL: "https://www.crazygames.com/t/escape", Paginate: true},
		// {URL: "https://www.crazygames.com/t/flash", Paginate: true},
		// {URL: "https://www.crazygames.com/t/first-person-shooter", Paginate: true},
		// {URL: "https://www.crazygames.com/t/horror", Paginate: true},
		// {URL: "https://www.crazygames.com/c/io", Paginate: true},
		// {URL: "https://www.crazygames.com/t/mahjong", Paginate: true},
		// {URL: "https://www.crazygames.com/t/minecraft", Paginate: true},
		// {URL: "https://www.crazygames.com/t/pool", Paginate: true},
		// {URL: "https://www.crazygames.com/c/puzzle", Paginate: true},
		// {URL: "https://www.crazygames.com/c/shooting", Paginate: true},
		// {URL: "https://www.crazygames.com/t/soccer", Paginate: true},
		// {URL: "https://www.crazygames.com/c/sports", Paginate: true},
		// {URL: "https://www.crazygames.com/t/stick", Paginate: true},
		{URL: "https://www.crazygames.com/t/tower-defense", Paginate: true},
	}

	// Initialize crawler
	crawler := crawler.NewCrawler(ctx, startURLs)
	defer crawler.Close()

	// Crawl homepages and process games synchronously
	if err := crawler.CrawlHomepage(); err != nil {
		log.Fatalf("Failed to crawl homepages: %v", err)
	}

	log.Println("Scraping complete")
}
