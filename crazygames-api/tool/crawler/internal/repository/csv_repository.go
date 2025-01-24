package repository

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"

	"crazygames.io/tool/crawler/internal/domain"
)

// readExistingGames reads the existing games from a CSV file.
func readExistingGames(file *os.File) (map[string]bool, error) {
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	// Create a map to track existing games by URL.
	existingGames := make(map[string]bool)

	// Skip the header row.
	for _, record := range records[1:] {
		if len(record) > 1 {
			existingGames[strings.TrimSpace(record[1])] = true
		}
	}
	return existingGames, nil
}

// AppendGameToCSV writes a single game to the CSV file, appending if the file exists.
func AppendGameToCSV(game domain.GameData) error {
	const fileName = "games.csv"

	// Open the file in append mode (create if it doesn't exist)
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open CSV file: %w", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Check if the file is empty (write header if it is)
	fileInfo, err := file.Stat()
	if err != nil {
		return fmt.Errorf("failed to get file info: %w", err)
	}

	if fileInfo.Size() == 0 {
		header := []string{
			"Name",
			"URL",
			"Rating",
			"RatingVotes",
			"Developer",
			"ReleaseDate",
			"LastUpdated",
			"Technology",
			"Platforms",
			"Classification",
			"WikiPages",
			"Iframe",
			"Description",
			"Features",
			"Controls",
			"FAQ",
			"GameplayVideo",
			"HoverVideo",
			"ThumbnailURL",
		}
		if err := writer.Write(header); err != nil {
			return fmt.Errorf("failed to write header to CSV: %w", err)
		}
	}

	// Prepare the game record
	record := []string{
		game.Name,
		game.URL,
		game.Rating,
		game.RatingVotes,
		game.Developer,
		game.ReleaseDate,
		game.LastUpdated,
		game.Technology,
		game.Platforms,
		game.Classification,
		game.WikiPages,
		game.Iframe,
		game.Description,
		game.Features,
		game.Controls,
		game.FAQ,
		game.GameplayVideo,
		game.HoverVideo,
		game.ThumbnailURL,
	}

	// Ensure all fields have at least empty strings
	for i := range record {
		if record[i] == "" {
			record[i] = "N/A"
		}
	}

	// Write the record to the file
	if err := writer.Write(record); err != nil {
		return fmt.Errorf("failed to write game record to CSV: %w", err)
	}

	log.Printf("Successfully wrote game data for %s to CSV", game.Name)
	return nil
}

// WriteTagsCSV writes a list of tag groups to a CSV file.
func WriteTagsCSV(tagGroups []domain.TagGroup) error {
	const fileName = "tags.csv"

	// Open file in write mode
	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("failed to create tags CSV file: %w", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header
	header := []string{"Group", "Tag Name", "Tag Count", "Tag URL"}
	if err := writer.Write(header); err != nil {
		return fmt.Errorf("failed to write header to tags CSV: %w", err)
	}

	// Write data
	for _, group := range tagGroups {
		for _, tag := range group.Tags {
			record := []string{
				group.Group,
				tag.Name,
				tag.Count,
				tag.URL,
			}

			// Ensure all fields have at least empty strings
			for i := range record {
				if record[i] == "" {
					record[i] = "N/A"
				}
			}

			// Write the record to the file
			if err := writer.Write(record); err != nil {
				return fmt.Errorf("failed to write tag record to CSV: %w", err)
			}
		}
	}

	log.Printf("Saved tags data to %s", fileName)
	return nil
}
