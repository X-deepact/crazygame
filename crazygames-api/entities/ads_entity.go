package entities

import (
	"time"

	"gorm.io/gorm"
)

type Ads struct {
	ID        uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	ImageUrl  string         `gorm:"not null;check:image_url <> ''" json:"image_url"`
	Position  uint           `gorm:"not null;check:position > 0" json:"position"`
	GameId    uint           `gorm:"not null;check:game_id > 0" json:"game_id"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	// Relations
	Game *Game `gorm:"foreignKey:GameId" json:"game"`
}
