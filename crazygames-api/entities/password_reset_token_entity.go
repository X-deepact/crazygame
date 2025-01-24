package entities

import (
	"time"
)

type PasswordResetToken struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	Email     string    `json:"email" gorm:"unique;not null"`
	Token     string    `gorm:"size:255;not null"`
	ExpiredAt time.Time `gorm:"not null"`
	IsUsed    bool      `gorm:"default:false"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
