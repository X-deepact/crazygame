package repositories

import (
	"time"

	"crazygames.io/entities"
	"gorm.io/gorm"
)

type PasswordResetTokenRepository struct {
	db *gorm.DB
}

type PasswordResetTokenRepositoryInterface interface {
	Create(token *entities.PasswordResetToken) error
	Update(token *entities.PasswordResetToken) (*entities.PasswordResetToken, error)
	GetByEmail(email string) (*entities.PasswordResetToken, error)
	GetByToken(token string) (*entities.PasswordResetToken, error)
	SetResetToken(email, token string, expiresAt time.Time) error
	GetUserByResetToken(token string) (*entities.PasswordResetToken, error)
	MarkTokenAsUsed(tokenID uint) error
}

func NewPasswordResetTokenRepository(db *gorm.DB) *PasswordResetTokenRepository {
	return &PasswordResetTokenRepository{db: db}
}

func (r *PasswordResetTokenRepository) Create(token *entities.PasswordResetToken) error {
	return r.db.Create(token).Error
}

func (r *PasswordResetTokenRepository) Update(token *entities.PasswordResetToken) (*entities.PasswordResetToken, error) {
	err := r.db.Save(token).Error
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (r *PasswordResetTokenRepository) GetByEmail(email string) (*entities.PasswordResetToken, error) {
	var resetToken entities.PasswordResetToken
	err := r.db.Where("email = ?", email).First(&resetToken).Error
	if err != nil {
		return nil, err
	}
	return &resetToken, nil
}

func (r *PasswordResetTokenRepository) GetByToken(token string) (*entities.PasswordResetToken, error) {
	var resetToken entities.PasswordResetToken
	err := r.db.Where("token = ?", token).First(&resetToken).Error
	if err != nil {
		return nil, err
	}
	return &resetToken, nil
}

func (r *PasswordResetTokenRepository) SetResetToken(email, token string, expiresAt time.Time) error {
	return r.db.Model(&entities.PasswordResetToken{}).
		Where("email = ?", email).
		Updates(map[string]interface{}{
			"token":      token,
			"expired_at": expiresAt,
			"is_used":    false,
		}).Error
}

func (r *PasswordResetTokenRepository) GetUserByResetToken(token string) (*entities.PasswordResetToken, error) {
	var resetToken entities.PasswordResetToken
	err := r.db.Where("token = ? AND expired_at > ? AND is_used = ?", token, time.Now(), false).
		First(&resetToken).Error
	if err != nil {
		return nil, err
	}
	return &resetToken, nil
}

func (r *PasswordResetTokenRepository) MarkTokenAsUsed(tokenID uint) error {
	return r.db.Model(&entities.PasswordResetToken{}).
		Where("id = ?", tokenID).
		Update("is_used", true).
		Error
}
