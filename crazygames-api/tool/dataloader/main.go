package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"crazygames.io/config"
	"crazygames.io/services"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// CSVLoaderConfig defines the configuration for loading data from a CSV
type CSVLoaderConfig struct {
	TableName string
	Columns   []ColumnMapping
	ChunkSize int // Number of records to process in each batch
}

// ColumnMapping represents how to map CSV columns to database columns
type ColumnMapping struct {
	To       string                                       // Target column name
	From     []string                                     // Source column names
	Mutate   func(values []string) interface{}            // Optional transformation function
	Validate func(value interface{}) (interface{}, error) // Optional validation function
}

func main() {
	csvFile := flag.String("csv", "", "Path to CSV file")
	flag.Parse()

	if *csvFile == "" {
		log.Fatal("CSV file path is required")
	}

	// Initialize database connection
	config.LoadConfig()
	db := config.ConnectDatabase()

	// Initialize MinIO service
	minioClient := config.ConnectMinIO()
	minioService := services.NewMinIOService(minioClient)

	// Disable SQL logging
	db = db.Session(&gorm.Session{Logger: db.Logger.LogMode(logger.Silent)})

	// Process CSV file with predefined configs
	configs := getGameConfigs(minioService)
	if err := processCSV(*csvFile, configs, db); err != nil {
		log.Fatalf("Error processing CSV: %v", err)
	}
}

func getGameConfigs(minioService *services.MinIOService) CSVLoaderConfig {
	return CSVLoaderConfig{
		TableName: "games",
		ChunkSize: 500, // Process 500 records at a time
		Columns: []ColumnMapping{
			{
				To:   "game_title",
				From: []string{"Name"},
				Mutate: func(values []string) interface{} {
					return strings.Title(strings.ToLower(values[0]))
				},
			},
			{
				To:   "release_date",
				From: []string{"ReleaseDate"},
				Mutate: func(values []string) interface{} {
					// Parse "Month YYYY" format into time.Time
					dateStr := values[0]
					parsedTime, err := time.Parse("January 2006", dateStr)
					if err != nil {
						log.Printf("Warning: Invalid date format '%s', using current time", dateStr)
						return time.Now()
					}
					return parsedTime
				},
			},
			{
				To:   "thumbnail_url",
				From: []string{"ThumbnailURL"},
				Mutate: func(values []string) interface{} {
					// Skip if empty URL
					if values[0] == "" {
						return ""
					}

					// Generate unique filename
					fileName := fmt.Sprintf("thumbnails/%d%s", time.Now().UnixNano(), filepath.Ext(values[0]))

					// Download and upload to MinIO
					imageUrl, err := minioService.UploadFromURL(values[0], fileName)
					if err != nil {
						log.Printf("Error processing thumbnail: %v", err)
						return "" // Return empty string if upload fails
					}
					return imageUrl
				},
			},
			{
				To:   "description",
				From: []string{"Description"},
			},
			{
				To:   "developer",
				From: []string{"Developer"},
			},
			{
				To:   "game_url",
				From: []string{"Name"},
				Mutate: func(values []string) interface{} {
					// Convert to lowercase and replace spaces with hyphens
					formattedTitle := strings.ToLower(values[0])
					formattedTitle = strings.ReplaceAll(formattedTitle, " ", "-")

					// Remove special characters
					formattedTitle = strings.Map(func(r rune) rune {
						if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == '-' {
							return r
						}
						return -1
					}, formattedTitle)

					// Build URL using DOMAIN_URL from .env
					return fmt.Sprintf("%s/game/%s", os.Getenv("DOMAIN_URL"), formattedTitle)
				},
			},
		},
	}
}

func processCSV(csvPath string, config CSVLoaderConfig, db *gorm.DB) error {
	file, err := os.Open(csvPath)
	if err != nil {
		return fmt.Errorf("error opening CSV file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// Read headers
	headers, err := reader.Read()
	if err != nil {
		return fmt.Errorf("error reading CSV headers: %v", err)
	}

	// Create header mapping
	headerMap := make(map[string]int)
	for i, header := range headers {
		headerMap[strings.TrimSpace(header)] = i
	}

	// Initialize progress tracking
	startTime := time.Now()
	savedCount := 0
	var chunk []map[string]interface{}

	// Get total records for progress tracking
	file.Seek(0, 0)
	totalRecords := 0
	reader = csv.NewReader(file)
	reader.Read() // Skip header
	for {
		_, err := reader.Read()
		if err != nil {
			break
		}
		totalRecords++
	}
	file.Seek(0, 0)
	reader = csv.NewReader(file)
	reader.Read() // Skip header again

	log.Printf("Starting data load... Total records to process: %d", totalRecords)

	for {
		record, err := reader.Read()
		if err != nil {
			// Process remaining records in last chunk
			if len(chunk) > 0 {
				if err := insertChunk(db, config.TableName, chunk); err != nil {
					return err
				}
				savedCount += len(chunk)
			}
			break // End of file
		}

		// Create dynamic map for row data
		rowData := make(map[string]interface{})

		for _, mapping := range config.Columns {
			// Collect source column values
			var sourceValues []string
			for _, field := range mapping.From {
				colIndex, exists := headerMap[field]
				if !exists {
					return fmt.Errorf("column %s not found in CSV", field)
				}
				sourceValues = append(sourceValues, record[colIndex])
			}

			// Apply mutation if provided
			var processedValue interface{} = sourceValues[0]
			if mapping.Mutate != nil {
				processedValue = mapping.Mutate(sourceValues)
			}

			// Optional validation
			if mapping.Validate != nil {
				var err error
				processedValue, err = mapping.Validate(processedValue)
				if err != nil {
					log.Printf("Validation error for %s: %v", mapping.To, err)
					continue
				}
			}

			rowData[mapping.To] = processedValue
		}

		chunk = append(chunk, rowData)

		// Process chunk when it reaches the configured size
		if len(chunk) >= config.ChunkSize {
			if err := insertChunk(db, config.TableName, chunk); err != nil {
				return err
			}
			savedCount += len(chunk)
			chunk = nil // Reset chunk
			log.Printf("Processed %d/%d records (%.1f%%)", savedCount, totalRecords, float64(savedCount)/float64(totalRecords)*100)
		}
	}

	log.Printf("Completed! Total records: %d", savedCount)
	log.Printf("Time elapsed: %s", time.Since(startTime).Round(time.Second))
	return nil
}

func insertChunk(db *gorm.DB, tableName string, chunk []map[string]interface{}) error {
	result := db.Table(tableName).Create(chunk)
	if result.Error != nil {
		return fmt.Errorf("error inserting chunk: %v", result.Error)
	}
	return nil
}
