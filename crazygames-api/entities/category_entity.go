package entities

import (
	"time"
)

type Category struct {
	ID           uint   `gorm:"primaryKey;autoIncrement"`
	CategoryName string `gorm:"unique;not null;check:category_name <> ''"`
	Description  string
	Icon         string
	Path         string
	IsMenu       bool
	CreatedAt    *time.Time `gorm:"autoCreateTime"`
	UpdatedAt    *time.Time `gorm:"autoUpdateTime"`
	Game         []Game     `gorm:"many2many:game_categories;"`
}
