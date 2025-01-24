package services

import (
	"fmt"
	"time"

	"crazygames.io/entities"
	"crazygames.io/handler/request"
	"crazygames.io/repositories"

	"os"

	"crazygames.io/utils"
	"github.com/minio/minio-go/v7"
)

type GameService struct {
	gameRepo    repositories.GameRepositoryInterface
	minioClient *minio.Client
}

type GameServiceInterface interface {
	Create(request *request.GameRequestCreate) (*entities.Game, error)
	GetAll(query request.GamesRequestQuery) ([]entities.Game, int64, error)
	GetByID(id uint) (*entities.Game, error)
	GetByCategoryID(id uint) ([]entities.Game, error)
	Update(id uint, request *request.GameRequestUpdate) (*entities.Game, error)
	Delete(id uint) error
	ListByCategory(categoryId uint) ([]entities.Game, error)
}

func NewGameService(gameRepo repositories.GameRepositoryInterface, minioClient *minio.Client) *GameService {
	return &GameService{gameRepo: gameRepo, minioClient: minioClient}
}

func (gs *GameService) Create(request *request.GameRequestCreate) (*entities.Game, error) {
	// Upload thumbnail to MinIO
	Thumbnail, err := utils.UploadFileToMinio(gs.minioClient, request.Thumbnail, os.Getenv("MINIO_BUCKET_NAME"))
	if err != nil {
		return nil, err
	}

	hoverVideoUrl, err := utils.UploadFileToMinio(gs.minioClient, request.HoverVideo, os.Getenv("MINIO_BUCKET_NAME"))
	if err != nil {
		return nil, err
	}

	layout := "2006-01-02"
	date, err := time.Parse(layout, request.ReleaseDate)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return nil, err
	}

	game := &entities.Game{
		GameTitle:     request.GameTitle,
		Description:   request.Description,
		Developer:     request.Developer,
		ReleaseDate:   &date,
		ThumbnailURL:  Thumbnail,
		Technology:    request.Technology,
		Rating:        request.Rating,
		HoverVideoUrl: hoverVideoUrl,
		GameURL:       request.GameURL,
		PlayCount:     request.PlayCount,
	}

	err = gs.gameRepo.Create(game, request.CategoryID)
	if err != nil {
		return nil, err
	}
	return game, nil
}

func (gs *GameService) GetAll(query request.GamesRequestQuery) ([]entities.Game, int64, error) {
	return gs.gameRepo.GetAll(query)
}

func (gs *GameService) GetByID(id uint) (*entities.Game, error) {
	return gs.gameRepo.GetByID(id)
}

func (gs *GameService) GetByCategoryID(id uint) ([]entities.Game, error) {
	return gs.gameRepo.GetByCategoryID(id)
}

func (gs *GameService) Update(id uint, request *request.GameRequestUpdate) (*entities.Game, error) {
	game, err := gs.GetByID(id)
	if err != nil {
		return nil, err
	}

	if request.Thumbnail != nil {
		thumbnail, err := utils.UploadFileToMinio(gs.minioClient, request.Thumbnail, os.Getenv("MINIO_BUCKET_NAME"))
		if err != nil {
			return nil, err
		}
		game.ThumbnailURL = thumbnail
	}

	if request.HoverVideo != nil {
		hoverVideoUrl, err := utils.UploadFileToMinio(gs.minioClient, request.HoverVideo, os.Getenv("MINIO_BUCKET_NAME"))
		if err != nil {
			return nil, err
		}
		game.HoverVideoUrl = hoverVideoUrl
	}

	if request.ReleaseDate == "" {
		layout := "2006-01-02"
		date, err := time.Parse(layout, request.ReleaseDate)
		if err != nil {
			fmt.Println("Error parsing date:", err)
			return nil, err
		}
		game.ReleaseDate = &date
	}
	if request.GameTitle != "" {
		game.GameTitle = request.GameTitle
	}
	if request.Description != "" {
		game.Description = request.Description
	}
	if request.Developer != "" {
		game.Developer = request.Developer
	}
	if request.Technology != "" {
		game.Technology = request.Technology
	}
	if request.Rating != 0 {
		game.Rating = request.Rating
	}
	if request.GameURL != "" {
		game.GameURL = request.GameURL
	}
	if request.PlayCount != 0 {
		game.PlayCount = request.PlayCount
	}

	categoryID := ""
	if request.CategoryID != "" {
		categoryID = request.CategoryID
	}

	return gs.gameRepo.Update(game, categoryID)
}

func (gs *GameService) Delete(id uint) error {
	return gs.gameRepo.Delete(id)
}

func (gs *GameService) ListByCategory(categoryId uint) ([]entities.Game, error) {
	return gs.gameRepo.ListByCategory(categoryId)
}
