package entities

import (
	"time"
)

type Game struct {
	ID            uint   `gorm:"primaryKey;autoIncrement"`
	GameTitle     string `gorm:"not null"`
	Description   string
	Developer     string
	ReleaseDate   *time.Time
	ThumbnailURL  string
	Technology    string
	Rating        float64
	HoverVideoUrl string
	GameURL       string     `gorm:"not null"`
	PlayCount     int        `gorm:"default:0"`
	Category      []Category `gorm:"many2many:game_categories;"`
	CreatedAt     *time.Time `gorm:"autoCreateTime"`
	UpdatedAt     *time.Time `gorm:"autoUpdateTime"`
}
