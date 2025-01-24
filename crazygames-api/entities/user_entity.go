package entities

import (
	"time"
)

const (
	RoleAdmin  = "admin"
	RolePlayer = "player"
)

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Username  string    `json:"username" gorm:"unique;not null;check:username <> ''"`
	Password  string    `json:"-" gorm:"not null;check:password <> ''"`
	Email     string    `json:"email" gorm:"unique;not null;check:email <> ''"`
	Role      string    `json:"role" gorm:"not null;default:player"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
