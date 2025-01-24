package entities

import "time"

type Tag struct {
	ID        uint       `gorm:"primaryKey;autoIncrement"`
	TagName   string     `gorm:"unique;not null"`
	CreatedAt *time.Time `gorm:"autoCreateTime"`
	UpdatedAt *time.Time `gorm:"autoUpdateTime"`
}

type GameTag struct {
	GameID uint
	TagID  uint
}
