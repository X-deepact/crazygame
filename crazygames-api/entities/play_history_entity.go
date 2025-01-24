package entities

import "time"

type PlayHistory struct {
	ID         uint `gorm:"primaryKey;autoIncrement"`
	UserID     uint
	GameID     uint
	DatePlayed *time.Time `gorm:"autoCreateTime"`
	CreatedAt  *time.Time `gorm:"autoCreateTime"`
	UpdatedAt  *time.Time `gorm:"autoUpdateTime"`
}
