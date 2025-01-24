# Data Loader Tool

This tool loads game data from CSV files into the database.

## CSV Format Requirements

The CSV file must have these exact column headers:
- Name
- ReleaseDate (format: "Month YYYY")
- ThumbnailURL
- Description
- Developer
- Iframe

Example CSV row:
```
"Super Game","December 2024","https://example.com/thumb.jpg","A fun game","GameDev Co","<iframe>"
```

## Usage

1. Prepare your CSV file with the required format
2. Run the loader:
```bash
go run tool/dataloader/main.go -csv path/to/games.csv
```

## Data Processing

The loader will:
1. Process data in chunks of 2000 records at a time
2. Convert game titles to Title Case
3. Parse "Month YYYY" dates into proper datetime values
4. Download and upload thumbnails to MinIO
5. Map CSV columns to database fields:
   - Name → game_title
   - ReleaseDate → release_date
   - ThumbnailURL → thumbnail_url
   - Description → description
   - Developer → developer
   - Iframe → game_url

## Error Handling

- Invalid dates will be logged and replaced with current time
- Missing required fields will cause the row to be skipped
- Duplicate game titles will be logged but still inserted

## Example Output

```
2025/01/17 21:00:00 Saved game: Super Game (Total: 1)
2025/01/17 21:00:01 Warning: Invalid date format 'Invalid Date', using current time
2025/01/17 21:00:02 Total games saved: 42
```
