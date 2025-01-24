package entities

import "time"

type Review struct {
	ID         uint `gorm:"primaryKey;autoIncrement"`
	GameID     uint
	UserID     uint
	Rating     int `gorm:"not null;check:rating >= 1 AND rating <= 5"`
	Comment    string
	DatePosted *time.Time `gorm:"autoCreateTime"`
	CreatedAt  *time.Time `gorm:"autoCreateTime"`
	UpdatedAt  *time.Time `gorm:"autoUpdateTime"`
}
