package entities

import "time"

type Favorite struct {
	ID        uint `gorm:"primaryKey;autoIncrement"`
	UserID    uint
	GameID    uint
	CreatedAt *time.Time `gorm:"autoCreateTime"`
	UpdatedAt *time.Time `gorm:"autoUpdateTime"`
}
