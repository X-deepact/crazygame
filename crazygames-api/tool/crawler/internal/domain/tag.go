package domain

// TagData represents a single tag with its name, count, and URL
type TagData struct {
	Name  string `json:"name"`
	Count string `json:"count"`
	URL   string `json:"url"`
}

// TagGroup represents a group of tags (e.g., "A", "B", etc.)
type TagGroup struct {
	Group string    `json:"group"`
	Tags  []TagData `json:"tags"`
}
