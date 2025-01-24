package repositories

import (
	"log"
	"os"
	"testing"

	"crazygames.io/entities"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	db                           *gorm.DB
	userRepository               *UserRepository
	adsRepo                      *adsRepository
	categoryRepository           *CategoryRepository
	gameRepository               *GameRepository
	passwordResetTokenRepository *PasswordResetTokenRepository
)

func TestMain(m *testing.M) {
	var err error
	dsn := "test_user:test_password@tcp(localhost:3307)/test_db?charset=utf8mb4&parseTime=True&loc=Local"

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(
		&entities.User{},
		&entities.Category{},
		&entities.Game{},
		&entities.Review{},
		&entities.Favorite{},
		&entities.Ads{},
		&entities.PasswordResetToken{},
	)
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	// Initialize repositories
	userRepository = NewUserRepository(db)
	categoryRepository = NewCategoryRepository(db)
	adsRepo = NewAdsRepository(db)
	gameRepository = NewGameRepository(db)
	passwordResetTokenRepository = NewPasswordResetTokenRepository(db)

	// run the tests
	code := m.Run()

	sqlDB, err := db.DB()
	if err != nil {
		log.Printf("failed to get sql.DB: %v", err)
	} else {
		sqlDB.Close()
	}

	os.Exit(code)
}
