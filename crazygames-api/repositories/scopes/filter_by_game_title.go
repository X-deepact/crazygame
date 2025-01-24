package scopes

import "gorm.io/gorm"

func FilterByGameTitle(gameTitle string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if gameTitle == "" {
			return db
		}

		return db.Where("game_title LIKE ?", "%"+gameTitle+"%")
	}
}
