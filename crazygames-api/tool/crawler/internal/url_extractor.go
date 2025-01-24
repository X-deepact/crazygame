package crawler

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type URLInfo struct {
	URL          string
	HoverVideo   string
	ThumbnailURL string
}

type URLExtractor struct {
	ctx context.Context
}

func NewURLExtractor(ctx context.Context) *URLExtractor {
	return &URLExtractor{ctx: ctx}
}

func (e *URLExtractor) ExtractGameURLs() ([]URLInfo, error) {
	log.Println("Starting URL extraction...")
	startTime := time.Now()

	// Create a timeout context for the extraction
	timeoutCtx, cancel := context.WithTimeout(e.ctx, 15*time.Minute)
	defer cancel()

	var gameInfos []URLInfo

	// Extract links from current page
	var pageLinks []string
	err := chromedp.Run(timeoutCtx,
		chromedp.ActionFunc(func(ctx context.Context) error {
			log.Println("Finding game thumbnail links...")
			return chromedp.Evaluate(`Array.from(document.querySelectorAll('a[href*="/game/"]')).map(a => a.href)`, &pageLinks).Do(ctx)
		}),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get game links: %w", err)
	}

	if len(pageLinks) > 0 {
		log.Printf("Found %d game links. First few: %v", len(pageLinks), pageLinks[:min(3, len(pageLinks))])
	} else {
		log.Printf("Found 0 game links. HTML may have changed - check selector")
		return nil, nil
	}

	// Use a channel to collect hover video results
	type videoResult struct {
		link         string
		video        string
		err          error
		thumbnailURL string
	}
	videoChan := make(chan videoResult, len(pageLinks))

	// Process hover videos concurrently
	for _, link := range pageLinks {
		go func(link string) {
			var video string
			var thumbnailURL string
			var err error
			for attempt := 1; attempt <= 3; attempt++ {
				video, thumbnailURL, err = e.processGameLink(timeoutCtx, link, attempt)
				log.Printf("Processing attempt %d for %s - thumbnail URL: %s", attempt, link, thumbnailURL)
				if err == nil {
					break
				}
				if attempt < 3 {
					time.Sleep(5 * time.Second)
				}
			}
			videoChan <- videoResult{link: link, video: video, thumbnailURL: thumbnailURL, err: err}
		}(link)
	}

	// Collect results from goroutines
	for range pageLinks {
		result := <-videoChan
		if result.err != nil {
			log.Printf("Failed to process %s: %v", result.link, result.err)
			continue
		}

		gameInfos = append(gameInfos, URLInfo{
			URL:          result.link,
			HoverVideo:   result.video,
			ThumbnailURL: result.thumbnailURL,
		})
	}

	log.Printf("Successfully extracted %d URLs in %v", len(gameInfos), time.Since(startTime))
	return gameInfos, nil
}

func (e *URLExtractor) processGameLink(ctx context.Context, link string, attempt int) (string, string, error) {
	log.Printf("Processing game link: %s (attempt %d)", link, attempt)
	timeout := time.Duration(attempt) * 30 * time.Second
	hoverCtx, hoverCancel := context.WithTimeout(ctx, timeout)
	defer hoverCancel()

	// Process the link
	var videoSources []string
	var thumbnailURL string
	err := chromedp.Run(hoverCtx,
		chromedp.Sleep(1*time.Second),
		chromedp.ScrollIntoView(`a[href="`+link+`"]`, chromedp.ByQuery),
		chromedp.Sleep(1*time.Second),
		chromedp.Focus(`a[href="`+link+`"]`, chromedp.ByQuery),
		chromedp.Evaluate(`
			var element = document.querySelector('a[href="`+link+`"]');
			if (element) {
				element.classList.remove('hover');
				var event = new MouseEvent('mouseover', {
					bubbles: true,
					cancelable: true,
					view: window
				});
				element.dispatchEvent(event);
				element.classList.add('hover');
			}
			true;
		`, nil),
		chromedp.Sleep(3*time.Second),
		chromedp.Evaluate(`
			(() => {
				const thumbnailImg = document.querySelector('a[href="`+link+`"] img.GameThumb_gameThumbImage__isqyS');
				const videoSources = Array.from(document.querySelectorAll('a[href="`+link+`"] .GameThumb_gameThumbVideo__pfwm video source'));
				
				if (videoSources.length > 0) {
					const videoSrc = videoSources[1] ? videoSources[1].src : videoSources[0].src;
					if (videoSrc) {
						// Extract the game slug from the video URL
						const matches = videoSrc.match(/\/([^\/]+)-landscape-/);
						if (matches && matches[1]) {
							const gameSlug = matches[1];
							return "https://imgs.crazygames.com/" + gameSlug + "_16x9/" + gameSlug + "_16x9-cover?auto=format%2Ccompress&q=90&cs=strip&w=273&fit=crop";
						}
					}
					return null;
				}
				
				if (thumbnailImg) {
					console.log('Thumbnail component:', thumbnailImg.outerHTML);
					return thumbnailImg.getAttribute('src');
				}
				return null;
			})()
		`, &thumbnailURL),
		chromedp.Evaluate(`
			(() => {
				const videos = Array.from(document.querySelectorAll('a[href="`+link+`"] .GameThumb_gameThumbVideo__pfwm video source'));
				if (videos.length === 0) {
					const altVideos = Array.from(document.querySelectorAll('a[href="`+link+`"] video source'));
					return altVideos.map(src => src.src);
				}
				return videos.map(src => src.src);
			})()
		`, &videoSources),
	)

	if err != nil {
		return "", "", fmt.Errorf("failed to process link: %w", err)
	}

	if thumbnailURL != "" {
		log.Printf("Found thumbnail URL for %s: %s", link, thumbnailURL)
	} else {
		log.Printf("No thumbnail URL found for %s", link)
	}

	if len(videoSources) > 0 {
		log.Printf("Found hover video URL for %s: %s", link, strings.Join(videoSources, ", "))
	} else {
		log.Printf("No hover video found for %s", link)
	}

	return strings.Join(videoSources, ", "), thumbnailURL, nil
}
