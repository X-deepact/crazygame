package repositories

import (
	"context"

	"crazygames.io/entities"
	"crazygames.io/handler/request"
	"gorm.io/gorm"
)

type AdsRepositoryInterface interface {
	GetAll(ctx context.Context, query request.AdsRequestQuery) ([]entities.Ads, int64, error)
	GetById(ctx context.Context, id uint) (*entities.Ads, error)
	Create(ctx context.Context, ads *entities.Ads) error
	Update(ctx context.Context, ads *entities.Ads) (*entities.Ads, error)
	Delete(ctx context.Context, id uint) error
}

type adsRepository struct {
	db *gorm.DB
}

func NewAdsRepository(db *gorm.DB) *adsRepository {
	return &adsRepository{db: db}
}

func (r *adsRepository) GetById(ctx context.Context, id uint) (*entities.Ads, error) {
	var ads entities.Ads
	err := r.db.WithContext(ctx).First(&ads, id).Error
	if err != nil {
		return nil, err
	}
	return &ads, nil
}

func (r *adsRepository) GetAll(ctx context.Context, queryParams request.AdsRequestQuery) ([]entities.Ads, int64, error) {
	var ads []entities.Ads
	var total int64

	query := r.db.Model(&entities.Ads{}).WithContext(ctx).Order("created_at DESC")

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.Offset(queryParams.PageSize * (queryParams.PageNumber - 1)).
		Limit(queryParams.PageSize).Find(&ads).Error
	if err != nil {
		return nil, 0, err
	}

	return ads, total, nil
}

func (r *adsRepository) Create(ctx context.Context, ads *entities.Ads) error {
	err := r.db.WithContext(ctx).Create(&ads).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *adsRepository) Update(ctx context.Context, ads *entities.Ads) (*entities.Ads, error) {
	err := r.db.WithContext(ctx).Where("id = ?", ads.ID).Updates(&ads).Error
	if err != nil {
		return nil, err
	}
	return ads, nil
}

func (r *adsRepository) Delete(ctx context.Context, id uint) error {
	err := r.db.WithContext(ctx).Delete(&entities.Ads{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
