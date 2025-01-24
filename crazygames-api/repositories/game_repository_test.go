package repositories

import (
	"strconv"
	"testing"

	"crazygames.io/entities"
	"crazygames.io/handler/request"
)

func createGame(game *entities.Game, categoryId string) (*entities.Game, error) {
	err := gameRepository.Create(game, categoryId)
	if err != nil {
		return nil, err
	}

	return game, nil
}

func TestGameRepository_Create_Success(t *testing.T) {
	db.Exec("DELETE FROM games")
	db.Exec("DELETE FROM categories")

	category := &entities.Category{CategoryName: "Action"}
	categoryRepository.Create(category)

	game := &entities.Game{GameTitle: "Test Game", GameURL: "http://testgame.com"}
	err := gameRepository.Create(game, strconv.Itoa(int(category.ID)))
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if game.ID == 0 {
		t.Errorf("expected game to have an ID, got 0")
	}

	if len(game.Category) != 1 || game.Category[0].ID != category.ID {
		t.Errorf("expected game to be associated with category, got %v", game.Category)
	}
}

func TestGameRepository_GetByID_Success(t *testing.T) {
	db.Exec("DELETE FROM games")
	db.Exec("DELETE FROM categories")

	category := &entities.Category{CategoryName: "Adventure"}
	categoryRepository.Create(category)

	game := &entities.Game{GameTitle: "Adventure Game", GameURL: "http://adventuregame.com"}
	gameRepository.Create(game, strconv.Itoa(int(category.ID)))

	fetchedGame, err := gameRepository.GetByID(game.ID)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if fetchedGame.ID != game.ID {
		t.Errorf("expected game ID %d, got %d", game.ID, fetchedGame.ID)
	}
}

func TestGameRepository_Update_Success(t *testing.T) {
	db.Exec("DELETE FROM games")
	db.Exec("DELETE FROM categories")

	category1 := &entities.Category{CategoryName: "Old Category"}
	category2 := &entities.Category{CategoryName: "New Category"}
	categoryRepository.Create(category1)
	categoryRepository.Create(category2)

	game := &entities.Game{GameTitle: "Old Game", GameURL: "http://oldgame.com"}
	gameRepository.Create(game, strconv.Itoa(int(category1.ID)))

	game.GameTitle = "Updated Game"
	updatedGame, err := gameRepository.Update(game, strconv.Itoa(int(category2.ID)))
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if updatedGame.GameTitle != "Updated Game" {
		t.Errorf("expected updated game title, got %s", updatedGame.GameTitle)
	}

	if len(updatedGame.Category) != 1 || updatedGame.Category[0].ID != category2.ID {
		t.Errorf("expected game to be associated with new category, got %v", updatedGame.Category)
	}
}

func TestGameRepository_Delete_Success(t *testing.T) {
	db.Exec("DELETE FROM games")
	db.Exec("DELETE FROM categories")

	category := &entities.Category{CategoryName: "Delete Test"}
	err := categoryRepository.Create(category)
	if err != nil {
		t.Fatalf("failed to create category: %v", err)
	}

	game := &entities.Game{GameTitle: "Game to Delete", GameURL: "http://deletegame.com"}
	err = gameRepository.Create(game, strconv.Itoa(int(category.ID)))
	if err != nil {
		t.Fatalf("failed to create game: %v", err)
	}

	// Fetch the game to ensure it exists
	fetchedGame, err := gameRepository.GetByID(game.ID)
	if err != nil {
		t.Fatalf("failed to fetch game before deletion: %v", err)
	}
	t.Logf("fetched game: %+v", fetchedGame)

	// Delete the game
	err = gameRepository.Delete(game.ID)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	// Try to fetch the game again to ensure it is deleted
	_, err = gameRepository.GetByID(game.ID)
	if err == nil {
		t.Errorf("expected error for deleted game, got nil")
	}
}

func TestGameRepository_ListByCategory_Success(t *testing.T) {
	db.Exec("DELETE FROM games")
	db.Exec("DELETE FROM categories")

	category := &entities.Category{CategoryName: "Category Testing"}
	err := categoryRepository.Create(category)
	if err != nil {
		t.Fatalf("failed to create category: %v", err)
	}

	game1 := &entities.Game{GameTitle: "Game 1", GameURL: "http://game1.com"}
	game2 := &entities.Game{GameTitle: "Game 2", GameURL: "http://game2.com"}
	err = gameRepository.Create(game1, strconv.Itoa(int(category.ID)))
	if err != nil {
		t.Fatalf("failed to create game1: %v", err)
	}
	err = gameRepository.Create(game2, strconv.Itoa(int(category.ID)))
	if err != nil {
		t.Fatalf("failed to create game2: %v", err)
	}

	games, err := gameRepository.ListByCategory(category.ID)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if len(games) != 2 {
		t.Errorf("expected 2 games, got %d", len(games))
	}
}

func TestGameRepository_GetAll_Success(t *testing.T) {
	db.Exec("SET FOREIGN_KEY_CHECKS = 0")
	db.Exec("TRUNCATE TABLE reviews")
	db.Exec("TRUNCATE TABLE favorites")
	db.Exec("TRUNCATE TABLE game_categories")
	db.Exec("TRUNCATE TABLE games")
	db.Exec("TRUNCATE TABLE categories")
	db.Exec("SET FOREIGN_KEY_CHECKS = 1")

	category := &entities.Category{CategoryName: "All Games Test"}
	err := categoryRepository.Create(category)
	if err != nil {
		t.Fatalf("failed to create category: %v", err)
	}

	game1 := &entities.Game{GameTitle: "Game 1", GameURL: "http://game1.com"}
	game2 := &entities.Game{GameTitle: "Game 2", GameURL: "http://game2.com"}
	err = gameRepository.Create(game1, strconv.Itoa(int(category.ID)))
	if err != nil {
		t.Fatalf("failed to create game1: %v", err)
	}
	err = gameRepository.Create(game2, strconv.Itoa(int(category.ID)))
	if err != nil {
		t.Fatalf("failed to create game2: %v", err)
	}

	var query = request.GamesRequestQuery{
		PageNumber: 1,
		PageSize:   10,
		Search:     "",
	}
	games, _, err := gameRepository.GetAll(query)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if len(games) != 2 {
		t.Errorf("expected 2 games, got %d", len(games))
	}
}

func TestGameRepository_Create_InvalidCategoryID(t *testing.T) {
	db.Exec("SET FOREIGN_KEY_CHECKS = 0")
	db.Exec("TRUNCATE TABLE reviews")
	db.Exec("TRUNCATE TABLE favorites")
	db.Exec("TRUNCATE TABLE game_categories")
	db.Exec("TRUNCATE TABLE games")
	db.Exec("TRUNCATE TABLE categories")
	db.Exec("SET FOREIGN_KEY_CHECKS = 1")

	game := &entities.Game{GameTitle: "Invalid Category Game", GameURL: "http://invalidcategory.com"}
	err := gameRepository.Create(game, "invalid_id")
	if err == nil {
		t.Errorf("expected error for invalid category ID, got nil")
	}
}

func TestGameRepository_Create_CategoryNotFound(t *testing.T) {
	db.Exec("SET FOREIGN_KEY_CHECKS = 0")
	db.Exec("TRUNCATE TABLE reviews")
	db.Exec("TRUNCATE TABLE favorites")
	db.Exec("TRUNCATE TABLE game_categories")
	db.Exec("TRUNCATE TABLE games")
	db.Exec("TRUNCATE TABLE categories")
	db.Exec("SET FOREIGN_KEY_CHECKS = 1")

	game := &entities.Game{GameTitle: "Nonexistent Category Game", GameURL: "http://nonexistentcategory.com"}
	err := gameRepository.Create(game, "9999") // Nonexistent category ID
	if err == nil {
		t.Errorf("expected error for nonexistent category, got nil")
	}
}

func TestGameRepository_Update_InvalidCategoryID(t *testing.T) {
	db.Exec("SET FOREIGN_KEY_CHECKS = 0")
	db.Exec("TRUNCATE TABLE reviews")
	db.Exec("TRUNCATE TABLE favorites")
	db.Exec("TRUNCATE TABLE game_categories")
	db.Exec("TRUNCATE TABLE games")
	db.Exec("TRUNCATE TABLE categories")
	db.Exec("SET FOREIGN_KEY_CHECKS = 1")

	category := &entities.Category{CategoryName: "Valid Category"}
	categoryRepository.Create(category)

	game := &entities.Game{GameTitle: "Game to Update", GameURL: "http://updategame.com"}
	gameRepository.Create(game, strconv.Itoa(int(category.ID)))

	game.GameTitle = "Updated Game"
	_, err := gameRepository.Update(game, "invalid_id")
	if err == nil {
		t.Errorf("expected error for invalid category ID, got nil")
	}
}

func TestGameRepository_Update_CategoryNotFound(t *testing.T) {
	db.Exec("SET FOREIGN_KEY_CHECKS = 0")
	db.Exec("TRUNCATE TABLE reviews")
	db.Exec("TRUNCATE TABLE favorites")
	db.Exec("TRUNCATE TABLE game_categories")
	db.Exec("TRUNCATE TABLE games")
	db.Exec("TRUNCATE TABLE categories")
	db.Exec("SET FOREIGN_KEY_CHECKS = 1")

	category := &entities.Category{CategoryName: "Valid Category"}
	categoryRepository.Create(category)

	game := &entities.Game{GameTitle: "Game to Update", GameURL: "http://updategame.com"}
	gameRepository.Create(game, strconv.Itoa(int(category.ID)))

	game.GameTitle = "Updated Game"
	_, err := gameRepository.Update(game, "9999") // Nonexistent category ID
	if err == nil {
		t.Errorf("expected error for nonexistent category, got nil")
	}
}

func TestGameRepository_Delete_GameNotFound(t *testing.T) {
	db.Exec("SET FOREIGN_KEY_CHECKS = 0")
	db.Exec("TRUNCATE TABLE reviews")
	db.Exec("TRUNCATE TABLE favorites")
	db.Exec("TRUNCATE TABLE game_categories")
	db.Exec("TRUNCATE TABLE games")
	db.Exec("TRUNCATE TABLE categories")
	db.Exec("SET FOREIGN_KEY_CHECKS = 1")

	err := gameRepository.Delete(9999) // Nonexistent game ID
	if err == nil {
		t.Errorf("expected error for nonexistent game, got nil")
	}
}

func TestGameRepository_GetByID_GameNotFound(t *testing.T) {
	db.Exec("SET FOREIGN_KEY_CHECKS = 0")
	db.Exec("TRUNCATE TABLE reviews")
	db.Exec("TRUNCATE TABLE favorites")
	db.Exec("TRUNCATE TABLE game_categories")
	db.Exec("TRUNCATE TABLE games")
	db.Exec("TRUNCATE TABLE categories")
	db.Exec("SET FOREIGN_KEY_CHECKS = 1")
	_, err := gameRepository.GetByID(9999) // Nonexistent game ID
	if err == nil {
		t.Errorf("expected error for nonexistent game, got nil")
	}
}

func TestGameRepository_ListByCategory_NoGames(t *testing.T) {
	db.Exec("SET FOREIGN_KEY_CHECKS = 0")
	db.Exec("TRUNCATE TABLE reviews")
	db.Exec("TRUNCATE TABLE favorites")
	db.Exec("TRUNCATE TABLE game_categories")
	db.Exec("TRUNCATE TABLE games")
	db.Exec("TRUNCATE TABLE categories")
	db.Exec("SET FOREIGN_KEY_CHECKS = 1")

	category := &entities.Category{CategoryName: "Empty Category"}
	categoryRepository.Create(category)

	games, err := gameRepository.ListByCategory(category.ID)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if len(games) != 0 {
		t.Errorf("expected 0 games, got %d", len(games))
	}
}

func TestGameRepository_Create_CategoryDeleted(t *testing.T) {
	db.Exec("SET FOREIGN_KEY_CHECKS = 0")
	db.Exec("TRUNCATE TABLE reviews")
	db.Exec("TRUNCATE TABLE favorites")
	db.Exec("TRUNCATE TABLE game_categories")
	db.Exec("TRUNCATE TABLE games")
	db.Exec("TRUNCATE TABLE categories")
	db.Exec("SET FOREIGN_KEY_CHECKS = 1")

	category := &entities.Category{CategoryName: "Temporary Category"}
	categoryRepository.Create(category)

	// Delete the category
	categoryRepository.Delete(category.ID)

	game := &entities.Game{GameTitle: "Game with Deleted Category", GameURL: "http://deletedcategory.com"}
	err := gameRepository.Create(game, strconv.Itoa(int(category.ID)))
	if err == nil {
		t.Errorf("expected error for deleted category, got nil")
	}
}

func TestGameRepository_GetAll_Preloading(t *testing.T) {
	db.Exec("SET FOREIGN_KEY_CHECKS = 0")
	db.Exec("TRUNCATE TABLE reviews")
	db.Exec("TRUNCATE TABLE favorites")
	db.Exec("TRUNCATE TABLE game_categories")
	db.Exec("TRUNCATE TABLE games")
	db.Exec("TRUNCATE TABLE categories")
	db.Exec("SET FOREIGN_KEY_CHECKS = 1")

	category := &entities.Category{CategoryName: "Preload Test"}
	categoryRepository.Create(category)

	game := &entities.Game{GameTitle: "Game with Category", GameURL: "http://preloadtest.com"}
	gameRepository.Create(game, strconv.Itoa(int(category.ID)))

	var query = request.GamesRequestQuery{
		PageNumber: 1,
		PageSize:   10,
		Search:     "",
	}
	games, _, err := gameRepository.GetAll(query)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if len(games) != 1 {
		t.Errorf("expected 1 game, got %d", len(games))
	}

	if len(games[0].Category) != 1 || games[0].Category[0].ID != category.ID {
		t.Errorf("expected game to be associated with category, got %v", games[0].Category)
	}
}

func TestGameRepository_ListByCategory_CategoryNotFound(t *testing.T) {
	db.Exec("SET FOREIGN_KEY_CHECKS = 0")
	db.Exec("TRUNCATE TABLE reviews")
	db.Exec("TRUNCATE TABLE favorites")
	db.Exec("TRUNCATE TABLE game_categories")
	db.Exec("TRUNCATE TABLE games")
	db.Exec("TRUNCATE TABLE categories")
	db.Exec("SET FOREIGN_KEY_CHECKS = 1")

	games, err := gameRepository.ListByCategory(9999) // Nonexistent category ID
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if len(games) != 0 {
		t.Errorf("expected 0 games, got %d", len(games))
	}
}

func TestGameRepository_GetByCategoryID_Success(t *testing.T) {
	db.Exec("SET FOREIGN_KEY_CHECKS = 0")
	db.Exec("TRUNCATE TABLE reviews")
	db.Exec("TRUNCATE TABLE favorites")
	db.Exec("TRUNCATE TABLE game_categories")
	db.Exec("TRUNCATE TABLE games")
	db.Exec("TRUNCATE TABLE categories")
	db.Exec("SET FOREIGN_KEY_CHECKS = 1")

	// Create a category
	category := &entities.Category{CategoryName: "Test Category"}
	err := categoryRepository.Create(category)
	if err != nil {
		t.Fatalf("failed to create category: %v", err)
	}

	// Create games associated with the category
	game1 := &entities.Game{GameTitle: "Game 1", GameURL: "http://game1.com"}
	game2 := &entities.Game{GameTitle: "Game 2", GameURL: "http://game2.com"}
	err = gameRepository.Create(game1, strconv.Itoa(int(category.ID)))
	if err != nil {
		t.Fatalf("failed to create game1: %v", err)
	}
	err = gameRepository.Create(game2, strconv.Itoa(int(category.ID)))
	if err != nil {
		t.Fatalf("failed to create game2: %v", err)
	}

	// Fetch games by category ID
	games, err := gameRepository.GetByCategoryID(category.ID)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	// Verify the number of games returned
	if len(games) != 2 {
		t.Errorf("expected 2 games, got %d", len(games))
	}

	// Verify the games' details
	if games[0].GameTitle != "Game 1" && games[1].GameTitle != "Game 2" {
		t.Errorf("expected games to match, got %+v", games)
	}
}

func TestGameRepository_GetByCategoryID_CategoryNotFound(t *testing.T) {
	db.Exec("SET FOREIGN_KEY_CHECKS = 0")
	db.Exec("TRUNCATE TABLE reviews")
	db.Exec("TRUNCATE TABLE favorites")
	db.Exec("TRUNCATE TABLE game_categories")
	db.Exec("TRUNCATE TABLE games")
	db.Exec("TRUNCATE TABLE categories")
	db.Exec("SET FOREIGN_KEY_CHECKS = 1")

	// Attempt to fetch games for a nonexistent category
	_, err := gameRepository.GetByCategoryID(9999) // Nonexistent category ID
	if err == nil {
		t.Errorf("expected error for nonexistent category, got nil")
	}
}

func TestGameRepository_GetByCategoryID_NoGames(t *testing.T) {
	db.Exec("SET FOREIGN_KEY_CHECKS = 0")
	db.Exec("TRUNCATE TABLE reviews")
	db.Exec("TRUNCATE TABLE favorites")
	db.Exec("TRUNCATE TABLE game_categories")
	db.Exec("TRUNCATE TABLE games")
	db.Exec("TRUNCATE TABLE categories")
	db.Exec("SET FOREIGN_KEY_CHECKS = 1")

	// Create a category with no games
	category := &entities.Category{CategoryName: "Empty Category"}
	err := categoryRepository.Create(category)
	if err != nil {
		t.Fatalf("failed to create category: %v", err)
	}

	// Fetch games by category ID
	games, err := gameRepository.GetByCategoryID(category.ID)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	// Verify that the returned slice is empty
	if len(games) != 0 {
		t.Errorf("expected 0 games, got %d", len(games))
	}
}
