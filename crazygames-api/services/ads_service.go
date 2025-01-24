package services

import (
	"context"
	"os"

	"crazygames.io/handler/request"

	"crazygames.io/entities"
	"crazygames.io/repositories"
	"crazygames.io/utils"
	"github.com/minio/minio-go/v7"
)

type adsService struct {
	adsRepo     repositories.AdsRepositoryInterface
	minioClient *minio.Client
}

type AdsServiceInterface interface {
	Create(request *request.AdsRequestCreate) (*entities.Ads, error)
	GetAll(query request.AdsRequestQuery) ([]entities.Ads, int64, error)
	GetByID(id uint) (*entities.Ads, error)
	Update(request *request.AdsRequestUpdate, id uint) (*entities.Ads, error)
	Delete(id uint) error
}

func NewAdsService(adsRepo repositories.AdsRepositoryInterface, minioClient *minio.Client) *adsService {
	return &adsService{adsRepo: adsRepo, minioClient: minioClient}
}

func (a *adsService) Create(request *request.AdsRequestCreate) (*entities.Ads, error) {
	// Upload Image to MinIO
	imageUrl, err := utils.UploadFileToMinio(a.minioClient, request.Image, os.Getenv("MINIO_BUCKET_NAME"))
	if err != nil {
		return nil, err
	}

	ads := &entities.Ads{
		ImageUrl: imageUrl,
		GameId:   request.GameId,
		Position: request.Position,
	}

	err = a.adsRepo.Create(context.Background(), ads)
	if err != nil {
		return nil, err
	}

	return ads, nil
}

func (a *adsService) GetAll(query request.AdsRequestQuery) ([]entities.Ads, int64, error) {
	ads, total, err := a.adsRepo.GetAll(context.Background(), query)
	if err != nil {
		return []entities.Ads{}, 0, err
	}
	return ads, total, nil
}

func (a *adsService) GetByID(id uint) (*entities.Ads, error) {
	ads, err := a.adsRepo.GetById(context.Background(), id)
	if err != nil {
		return nil, err
	}
	return ads, nil
}

func (a *adsService) Update(request *request.AdsRequestUpdate, id uint) (*entities.Ads, error) {
	ads, err := a.adsRepo.GetById(context.Background(), id)
	if err != nil {
		return nil, err
	}

	imageUrl := ads.ImageUrl
	if request.Image != nil {
		imageUrl, err = utils.UploadFileToMinio(a.minioClient, request.Image, os.Getenv("MINIO_BUCKET_NAME"))
		if err != nil {
			return nil, err
		}
	}

	ads.ImageUrl = imageUrl
	ads.Position = request.Position
	ads.GameId = request.GameId

	return a.adsRepo.Update(context.Background(), ads)
}

func (a *adsService) Delete(id uint) error {
	err := a.adsRepo.Delete(context.Background(), id)
	if err != nil {
		return err
	}
	return nil
}
