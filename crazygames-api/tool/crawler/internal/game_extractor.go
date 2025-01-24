package crawler

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"crazygames.io/tool/crawler/internal/domain"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
)

type GameExtractor struct {
	ctx context.Context
}

func NewGameExtractor(ctx context.Context) *GameExtractor {
	return &GameExtractor{ctx: ctx}
}

func (e *GameExtractor) ExtractGameInfo(url string, hoverVideo string, thumbnailURL string) (domain.GameData, error) {
	// Check if context is already canceled
	select {
	case <-e.ctx.Done():
		log.Printf("Context canceled before starting extraction for %s", url)
		return domain.GameData{}, fmt.Errorf("context canceled")
	default:
		// Continue with extraction
	}

	log.Printf("Starting extraction for game: %s", url)

	var gameHTML string
	var lastError error
	startTime := time.Now()

	// Retry up to 3 times with increasing timeouts
	for attempt := 1; attempt <= 3; attempt++ {
		// Check context cancellation before each attempt
		select {
		case <-e.ctx.Done():
			log.Printf("Context canceled during extraction attempt for %s", url)
			return domain.GameData{}, fmt.Errorf("context canceled")
		default:
			// Continue with attempt
		}

		timeout := time.Duration(attempt) * 60 * time.Second // Increased to 60 seconds per attempt
		pageCtx, cancel := context.WithTimeout(e.ctx, timeout)
		defer cancel()

		err := chromedp.Run(pageCtx,
			chromedp.Navigate(url),
			chromedp.WaitVisible("body", chromedp.ByQuery),
			chromedp.OuterHTML("html", &gameHTML),
		)

		if err == nil {
			log.Printf("Successfully crawled page in %v, HTML length: %d", time.Since(startTime), len(gameHTML))
			break
		}

		lastError = err
		log.Printf("Attempt %d/3 failed for %s after %v: %v", attempt, url, time.Since(startTime), err)

		if attempt < 3 {
			time.Sleep(5 * time.Second)
		}
	}

	if lastError != nil {
		// If context was canceled, return immediately
		select {
		case <-e.ctx.Done():
			log.Printf("Context canceled after failed attempts for %s", url)
			return domain.GameData{}, fmt.Errorf("context canceled")
		default:
			log.Printf("Failed to crawl page %s after %v: %v", url, time.Since(startTime), lastError)
			return domain.GameData{}, fmt.Errorf("failed to crawl page %s: %w", url, lastError)
		}
	}

	log.Printf("Successfully crawled page in %v, HTML length: %d", time.Since(startTime), len(gameHTML))

	// Parse game name from HTML
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(gameHTML))
	if err != nil {
		return domain.GameData{}, fmt.Errorf("failed to parse HTML for %s: %w", url, err)
	}

	// Create game data structure with all fields initialized
	game := domain.GameData{
		Name:           extractGameName(doc),
		URL:            url,
		HoverVideo:     hoverVideo,
		ThumbnailURL:   thumbnailURL,
		Rating:         "N/A",
		RatingVotes:    "N/A",
		Developer:      "N/A",
		ReleaseDate:    "N/A",
		LastUpdated:    "N/A",
		Technology:     "N/A",
		Platforms:      "N/A",
		Classification: "N/A",
		WikiPages:      "N/A",
		Iframe:         "N/A",
		Description:    "N/A",
		Features:       "N/A",
		Controls:       "N/A",
		FAQ:            "N/A",
		GameplayVideo:  "N/A",
	}

	// Extract all game details
	game.Rating, game.RatingVotes = extractRating(doc)
	game.Developer = extractDeveloper(doc)
	game.ReleaseDate, game.LastUpdated = extractDates(doc)
	game.Technology, game.Platforms = extractTechAndPlatforms(doc)
	game.Classification = extractClassification(doc)
	game.WikiPages = extractWikiPages(doc)
	game.Iframe = extractIframe(doc)
	game.Description = extractDescription(doc)
	game.Features = extractFeatures(doc)
	game.Controls = extractControls(doc)
	game.FAQ = extractFAQ(doc)
	game.GameplayVideo = extractGameplayVideo(doc)

	return game, nil
}

func extractGameName(doc *goquery.Document) string {
	name := doc.Find("h1").First().Text()
	return strings.TrimSpace(name)
}

func extractRating(doc *goquery.Document) (string, string) {
	// Try multiple ways to find the rating section
	var rating, votes string = "N/A", "N/A"

	// Look for the specific rating container structure
	ratingContainer := doc.Find("div.css-16rvtsf")
	if ratingContainer.Length() > 0 {
		// Extract rating from the first div with font-weight:900
		ratingDiv := ratingContainer.Find("div[style*='font-weight:900']")
		if ratingDiv.Length() > 0 {
			rating = strings.TrimSpace(ratingDiv.Text())
		}

		// Extract votes from the div with font-size:12px
		votesDiv := ratingContainer.Find("div[style*='font-size:12px']")
		if votesDiv.Length() > 0 {
			votes = strings.TrimSpace(votesDiv.Text())
			// Clean up votes text
			votes = strings.NewReplacer(
				"(", "",
				")", "",
				"votes", "",
				"ratings", "",
			).Replace(votes)
			votes = strings.TrimSpace(votes)
		}
	} else {
		log.Println("Rating section not found using any method")
	}

	// Log the full HTML of the rating section for debugging
	if ratingContainer.Length() > 0 {
		html, _ := ratingContainer.Html()
		log.Printf("Rating section HTML:\n%s", html)
	}

	log.Printf("Rating extraction result - Rating: %s, Votes: %s", rating, votes)
	return rating, votes
}

func extractDeveloper(doc *goquery.Document) string {
	developer := doc.Find(".css-exrwgm").First().Text()
	return strings.TrimSpace(developer)
}

func extractDates(doc *goquery.Document) (string, string) {
	releaseDate := doc.Find(".css-12hp3i5:contains('Released:') .css-16rvtsf").First().Text()
	lastUpdated := doc.Find(".css-12hp3i5:contains('Last Updated:') .css-16rvtsf").First().Text()
	return strings.TrimSpace(releaseDate), strings.TrimSpace(lastUpdated)
}

func extractTechAndPlatforms(doc *goquery.Document) (string, string) {
	tech := doc.Find(".css-12hp3i5:contains('Technology:') .css-16rvtsf").First().Text()
	platforms := doc.Find(".css-12hp3i5:contains('Platforms:') .css-16rvtsf").First().Text()
	return strings.TrimSpace(tech), strings.TrimSpace(platforms)
}

func extractClassification(doc *goquery.Document) string {
	var categories []string
	doc.Find(".css-ez83vb a").Each(func(i int, s *goquery.Selection) {
		categories = append(categories, s.Text())
	})
	return strings.Join(categories, " » ")
}

func extractWikiPages(doc *goquery.Document) string {
	wikiDiv := doc.Find("div.css-16rvtsf").Has("a[target='_blank']")
	if wikiDiv.Length() > 0 {
		var wikis []string
		wikiDiv.Find("a").Each(func(i int, s *goquery.Selection) {
			wikis = append(wikis, s.Text())
		})
		result := strings.Join(wikis, ", ")
		log.Printf("Wiki pages: %s", result)
		return result
	}
	log.Println("Wiki pages not found")
	return "N/A"
}

func extractIframe(doc *goquery.Document) string {
	iframe, _ := doc.Find("iframe").First().Attr("src")
	return iframe
}

func extractDescription(doc *goquery.Document) string {
	description := doc.Find(".gameDescription_first p").First().Text()
	return strings.TrimSpace(description)
}

func extractFeatures(doc *goquery.Document) string {
	features := []string{}
	doc.Find("h3:contains('Features') + ul li").Each(func(i int, s *goquery.Selection) {
		features = append(features, strings.TrimSpace(s.Text()))
	})
	log.Printf("Found %d features", len(features))
	return strings.Join(features, "\n")
}

func extractControls(doc *goquery.Document) string {
	var controls []string
	doc.Find(".css-4ydurg ul li").Each(func(i int, s *goquery.Selection) {
		controls = append(controls, strings.TrimSpace(s.Text()))
	})
	return strings.Join(controls, "\n")
}

func extractFAQ(doc *goquery.Document) string {
	faqItems := []string{}
	doc.Find("h2:contains('FAQ') + div > div").Each(func(i int, s *goquery.Selection) {
		question := s.Find("h3").Text()
		answer := s.Find("div").Text()
		faqItems = append(faqItems, fmt.Sprintf("Q: %s\nA: %s", question, answer))
	})
	log.Printf("Found %d FAQ items", len(faqItems))
	return strings.Join(faqItems, "\n\n")
}

func extractGameplayVideo(doc *goquery.Document) string {
	videoContainer := doc.Find("article.yt-lite")
	if videoContainer.Length() > 0 {
		log.Printf("Found video container")

		// Extract video title
		title, _ := videoContainer.Attr("data-title")
		log.Printf("Video title: %s", title)

		// Extract background image URL
		style, _ := videoContainer.Attr("style")
		var bgImageURL string
		if strings.Contains(style, "background-image: url(") {
			bgImageURL = strings.Split(strings.Split(style, "background-image: url(")[1], ")")[0]
			bgImageURL = strings.Trim(bgImageURL, `"`)
			log.Printf("Background image URL: %s", bgImageURL)
		}

		// Extract video ID from background image URL
		var videoID string
		if bgImageURL != "" {
			parts := strings.Split(bgImageURL, "/")
			if len(parts) > 0 {
				videoID = strings.Split(parts[len(parts)-1], "_")[0]
				log.Printf("Extracted video ID: %s", videoID)
			}
		}

		// Try to find iframe
		videoIframe := videoContainer.Find("iframe")
		if videoIframe.Length() > 0 {
			src, exists := videoIframe.Attr("src")
			if exists {
				log.Printf("Gameplay video iframe found: %s", src)
				return src
			}
			log.Println("Iframe found but no src attribute")
		} else if videoID != "" {
			// Construct iframe URL if we have video ID
			videoURL := fmt.Sprintf("https://www.youtube-nocookie.com/embed/%s", videoID)
			log.Printf("Constructed gameplay video URL: %s", videoURL)
			return videoURL
		}
		log.Println("No iframe found within video container")
	}
	log.Println("Video container not found")
	return "N/A"
}
