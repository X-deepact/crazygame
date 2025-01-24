package repositories

import (
	"crazygames.io/entities"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

type UserRepositoryInterface interface {
	OauthCreate(user *entities.User) error
	Create(user *entities.User) error
	GetAll(page, limit int) ([]entities.User, error)
	GetByID(id uint) (*entities.User, error)
	GetByUsername(username string) (*entities.User, error)
	Update(user *entities.User) (*entities.User, error)
	UpdatePassword(userEmail string, hashedPassword string) error
	Delete(id uint) error
	GetByEmail(email string) (*entities.User, error)
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) OauthCreate(user *entities.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) Create(user *entities.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) GetAll(page, limit int) ([]entities.User, error) {
	var users []entities.User
	offset := (page - 1) * limit
	err := r.db.Limit(limit).Offset(offset).Find(&users).Error
	return users, err
}

func (r *UserRepository) GetByID(id uint) (*entities.User, error) {
	var user entities.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetByUsername(username string) (*entities.User, error) {
	var user entities.User
	err := r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Update(user *entities.User) (*entities.User, error) {
	err := r.db.Save(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) UpdatePassword(userEmail string, hashedPassword string) error {
	return r.db.Model(&entities.User{}).
		Where("email = ?", userEmail).
		Update("password", hashedPassword).
		Error
}

func (r *UserRepository) Delete(id uint) error {
	return r.db.Delete(&entities.User{}, id).Error
}

func (r *UserRepository) GetByEmail(email string) (*entities.User, error) {
	var user entities.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
