package entities

import (
	"gorm.io/gorm"
	"time"
)

type GameCategory struct {
	CategoryID uint `gorm:"primaryKey"`
	GameID     uint `gorm:"primaryKey"`
	CreatedAt  time.Time
	DeletedAt  gorm.DeletedAt
}
