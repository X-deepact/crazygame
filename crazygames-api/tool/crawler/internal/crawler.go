package crawler

import (
	"context"
	"fmt"
	"log"
	"math"
	"strings"
	"time"

	"crazygames.io/tool/crawler/internal/domain"
	"crazygames.io/tool/crawler/internal/repository"

	"github.com/chromedp/chromedp"
)

type URLConfig struct {
	URL      string
	Paginate bool
}

type Crawler struct {
	ctx        context.Context
	browserCtx context.Context
	games      []domain.GameData
	startURLs  []URLConfig
}

func (c *Crawler) AddGame(game domain.GameData) {
	// Append the game to the in-memory list
	c.games = append(c.games, game)
	log.Printf("Stored game data for %s with hover video: %s", game.Name, game.HoverVideo)

	// Write the game data to the CSV file immediately
	if err := repository.AppendGameToCSV(game); err != nil {
		log.Printf("Failed to write game data to CSV: %v", err)
	}
}

func NewCrawler(ctx context.Context, startURLs []URLConfig) *Crawler {
	// Create allocator context with options
	allocatorOpts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36"),
		chromedp.Flag("headless", true),
		chromedp.Flag("disable-gpu", true),
		chromedp.Flag("no-sandbox", true),
		chromedp.Flag("disable-dev-shm-usage", true), // Prevents memory issues
		chromedp.Flag("remote-debugging-port", "9222"),
	)

	// Create browser context with proper initialization
	allocatorCtx, cancelAllocator := chromedp.NewExecAllocator(ctx, allocatorOpts...)
	browserCtx, cancelBrowser := chromedp.NewContext(allocatorCtx)

	// Verify browser context
	if err := chromedp.Run(browserCtx, chromedp.Navigate("about:blank")); err != nil {
		cancelAllocator()
		cancelBrowser()
		log.Fatalf("Failed to initialize browser: %v", err)
	}

	return &Crawler{
		ctx:        ctx,
		browserCtx: browserCtx,
		startURLs:  startURLs,
	}
}

func (c *Crawler) Close() {
	// No need to save games here since they're written incrementally
	log.Println("Crawler shutdown complete")
}

func (c *Crawler) BrowserContext() context.Context {
	return c.browserCtx
}

func (c *Crawler) Context() context.Context {
	return c.ctx
}

func (c *Crawler) NewURLExtractor(ctx context.Context) *URLExtractor {
	return NewURLExtractor(ctx)
}

func (c *Crawler) NewGameExtractor(ctx context.Context) *GameExtractor {
	return NewGameExtractor(ctx)
}

func (c *Crawler) NewTagExtractor(ctx context.Context) *TagExtractor {
	return NewTagExtractor(ctx)
}

func (c *Crawler) CrawlHomepage() error {
	maxRetries := 5
	baseRetryDelay := 2 * time.Second

	log.Printf("Starting homepage crawl for %d URLs...", len(c.startURLs))

	for _, config := range c.startURLs {
		select {
		case <-c.ctx.Done():
			log.Printf("Context canceled: stopping crawl")
			return nil // Return nil to allow clean exit
		default:
			// Continue crawling
		}

		if config.Paginate {
			// Handle paginated URLs
			pageNum := 1
			for {
				pageURL := config.URL
				if pageNum > 1 {
					pageURL = fmt.Sprintf("%s/%d", config.URL, pageNum)
				}

				log.Printf("Crawling paginated URL: %s", pageURL)
				var html string
				var err error

				for i := 0; i < maxRetries; i++ {
					log.Printf("Attempt %d/%d: Crawling %s", i+1, maxRetries, pageURL)

					attemptCtx, cancel := context.WithTimeout(c.browserCtx, 2*time.Minute)
					defer cancel()

					startTime := time.Now()
					err = chromedp.Run(attemptCtx,
						chromedp.Navigate(pageURL),
						chromedp.WaitVisible("body", chromedp.ByQuery),
						chromedp.OuterHTML("html", &html),
					)
					duration := time.Since(startTime)

					if err == nil {
						log.Printf("Successfully crawled %s in %v", pageURL, duration)

						// Extract and process game URLs from this page
						urlExtractor := c.NewURLExtractor(c.browserCtx)
						gameInfos, err := urlExtractor.ExtractGameURLs()

						// If no games found on this page, try one more time with a longer wait
						if len(gameInfos) == 0 {
							time.Sleep(5 * time.Second)
							gameInfos, err = urlExtractor.ExtractGameURLs()
							if len(gameInfos) == 0 {
								log.Printf("No games found on page %d after retry. Last page appears to be %d", pageNum, pageNum-1)
								break
							}
						}
						if err != nil {
							log.Printf("Failed to extract game URLs from %s: %v", pageURL, err)
							continue
						}

						// Process each game synchronously
						gameExtractor := c.NewGameExtractor(c.browserCtx)
						for _, info := range gameInfos {
							log.Printf("Crawling game: %s", info.URL)

							game, err := gameExtractor.ExtractGameInfo(info.URL, info.HoverVideo, info.ThumbnailURL)
							if err != nil {
								log.Printf("Failed to crawl game page %s: %v", info.URL, err)
								continue
							}
							c.AddGame(game)
							log.Printf("Successfully extracted game: %s", game.Name)

							// Add small delay between games to reduce resource contention
							time.Sleep(1 * time.Second)
						}

						break
					}

					log.Printf("Attempt %d failed after %v: %v", i+1, duration, err)
					if i < maxRetries-1 {
						retryDelay := time.Duration(math.Pow(2, float64(i))) * baseRetryDelay
						log.Printf("Retrying in %v...", retryDelay)
						time.Sleep(retryDelay)
					}
				}

				if err != nil {
					log.Printf("Failed to crawl %s after %d attempts: %v", pageURL, maxRetries, err)
					break
				}

				// Exit pagination loop if "GAME OVER" page was detected
				if strings.Contains(html, `<div class="css-14yxz7x">`) &&
					strings.Contains(html, `Oops, you've reached a dead end!`) {
					break
				}

				pageNum++
			}
		} else {
			// Handle non-paginated URL
			log.Printf("Crawling URL: %s", config.URL)
			var html string
			var err error

			for i := 0; i < maxRetries; i++ {
				log.Printf("Attempt %d/%d: Crawling %s", i+1, maxRetries, config.URL)

				attemptCtx, cancel := context.WithTimeout(c.browserCtx, 30*time.Second)
				defer cancel()

				startTime := time.Now()
				err = chromedp.Run(attemptCtx,
					chromedp.Navigate(config.URL),
					chromedp.WaitVisible("body", chromedp.ByQuery),
					chromedp.OuterHTML("html", &html),
				)
				duration := time.Since(startTime)

				if err == nil {
					log.Printf("Successfully crawled %s in %v", config.URL, duration)
					break
				}

				log.Printf("Attempt %d failed after %v: %v", i+1, duration, err)
				if i < maxRetries-1 {
					retryDelay := time.Duration(math.Pow(2, float64(i))) * baseRetryDelay
					log.Printf("Retrying in %v (attempt %d/%d)...", retryDelay, i+1, maxRetries)
					time.Sleep(retryDelay)
				}
			}

			if err != nil {
				return fmt.Errorf("failed to crawl %s after %d attempts: %w", config.URL, maxRetries, err)
			}
		}
	}

	log.Println("Finished homepage crawl")

	// Extract tags after crawling games
	tagExtractor := c.NewTagExtractor(c.browserCtx)
	tagsURL := "https://www.crazygames.com/tags"
	tagGroups, err := tagExtractor.ExtractTags(tagsURL)
	if err != nil {
		log.Printf("Failed to extract tags: %v", err)
	} else {
		log.Printf("Successfully extracted %d tag groups", len(tagGroups))
		// TODO: Store or process tag data
	}

	return nil
}
