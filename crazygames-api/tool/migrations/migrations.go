package main

import (
	"log"

	"crazygames.io/config"
	"crazygames.io/entities"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMigrations() []*gormigrate.Migration {
	return []*gormigrate.Migration{
		{
			ID: "20250112_create_users_table",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&entities.User{})
			},
		},
		{
			ID: "20250112_create_categories_table",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&entities.Category{})
			},
		},
		{
			ID: "20250112_create_games_table",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&entities.Game{})
			},
		},
		{
			ID: "20250112_create_reviews_table",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&entities.Review{})
			},
		},
		{
			ID: "20250112_create_favorites_table",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&entities.Favorite{})
			},
		},
		{
			ID: "20250112_create_tags_table",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&entities.Tag{})
			},
		},
		{
			ID: "20250112_create_game_tags_table",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&entities.GameTag{})
			},
		},
		{
			ID: "20250112_create_play_history_table",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&entities.PlayHistory{})
			},
		},
		{
			ID: "20250115_create_ads_table",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&entities.Ads{})
			},
		},
		{
			ID: "20250115_game_categories_table",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&entities.GameCategory{})
			},
		},
		{
			ID: "20250112_password_reset_token_table",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&entities.PasswordResetToken{})
			},
		},
	}
}

func main() {
	config.LoadConfig()
	db := config.ConnectDatabase()
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: db.ConnPool,
	}), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to MySQL database: %v", err)
	}
	m := gormigrate.New(gormDB, gormigrate.DefaultOptions, InitMigrations())

	if err = m.Migrate(); err != nil {
		log.Fatalf("could not migrate: %v", err)
	}

	log.Println("Migrations ran successfully")
}
