package repositories

import (
	"errors"
	"strconv"

	"crazygames.io/entities"
	"crazygames.io/handler/request"
	"crazygames.io/repositories/scopes"
	"gorm.io/gorm"
)

type GameRepository struct {
	db *gorm.DB
}

type GameRepositoryInterface interface {
	Create(game *entities.Game, categoryID string) error
	GetAll(query request.GamesRequestQuery) ([]entities.Game, int64, error)
	GetByID(id uint) (*entities.Game, error)
	GetByCategoryID(id uint) ([]entities.Game, error)
	Update(game *entities.Game, categoryID string) (*entities.Game, error)
	Delete(id uint) error
	ListByCategory(categoryId uint) ([]entities.Game, error)
}

func NewGameRepository(db *gorm.DB) *GameRepository {
	return &GameRepository{db: db}
}

func (r *GameRepository) Create(game *entities.Game, categoryId string) error {
	categoryIDUint, err := strconv.ParseUint(categoryId, 10, 32)
	if err != nil {
		return errors.New("invalid category ID")
	}
	var category entities.Category
	err = r.db.Find(&category, uint(categoryIDUint)).Error
	if err != nil {
		return errors.New("category not found")
	}

	// Start a transaction
	tx := r.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	// Create the game
	if err = tx.Create(game).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Associate the game with the category
	if err = tx.Model(game).Association("Category").Append(&category); err != nil {
		tx.Rollback()
		return err
	}

	// Commit the transaction
	if err = tx.Commit().Error; err != nil {
		return err
	}

	game.Category = []entities.Category{category}
	return nil
}

func (r *GameRepository) GetAll(queryParams request.GamesRequestQuery) ([]entities.Game, int64, error) {
	var games []entities.Game
	var total int64

	query := r.db.Model(&entities.Game{}).Scopes(scopes.FilterByGameTitle(queryParams.Search)).Preload("Category")

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.Offset(queryParams.PageSize * (queryParams.PageNumber - 1)).
		Limit(queryParams.PageSize).Find(&games).Error

	return games, total, err
}

func (r *GameRepository) GetByID(id uint) (*entities.Game, error) {
	var game entities.Game
	err := r.db.Preload("Category").First(&game, id).Error
	if err != nil {
		return nil, err
	}
	return &game, nil
}

func (r *GameRepository) GetByCategoryID(id uint) ([]entities.Game, error) {
	var category entities.Category
	err := r.db.Preload("Game.Category").First(&category, id).Error
	if err != nil {
		return nil, err
	}
	return category.Game, nil
}

func (r *GameRepository) Update(game *entities.Game, categoryID string) (*entities.Game, error) {
	var category entities.Category
	if categoryID != "" {
		categoryIDUint, err := strconv.ParseUint(categoryID, 10, 32)
		if err != nil {
			return nil, errors.New("invalid category ID")
		}
		err = r.db.Find(&category, uint(categoryIDUint)).Error
		if err != nil {
			return nil, errors.New("category not found")
		}
	}

	// Start a transaction
	tx := r.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	// Create the game
	if err := tx.Save(game).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Associate the game with the category
	if err := tx.Model(game).Association("Category").Replace(&category); err != nil {
		tx.Rollback()
		return nil, err
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		return nil, err
	}
	game.Category = []entities.Category{category}
	return game, nil
}

func (r *GameRepository) Delete(id uint) error {
	var loadedGame entities.Game
	err := r.db.Preload("Category").First(&loadedGame, id).Error

	if err != nil {
		return errors.New("game not found")
	}
	// Start a transaction
	tx := r.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if err := tx.Model(&loadedGame).Association("Category").Clear(); err != nil {
		tx.Rollback()
		return err
	}

	// Create the game
	if err := tx.Delete(&loadedGame).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

func (r *GameRepository) ListByCategory(categoryId uint) ([]entities.Game, error) {
	var games []entities.Game
	err := r.db.Joins("JOIN game_categories ON game_categories.game_id = games.id").
		Where("game_categories.category_id = ?", categoryId).
		Preload("Category").
		Find(&games).Error
	if err != nil {
		return nil, err
	}
	return games, nil
}
