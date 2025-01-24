package services

import (
	"crazygames.io/handler/request"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"

	"crazygames.io/entities"
	"crazygames.io/repositories"
)

type UserService struct {
	userRepo               repositories.UserRepositoryInterface
	passwordResetTokenRepo repositories.PasswordResetTokenRepositoryInterface
}

type UserServiceInterface interface {
	Create(request *request.UserRequest) (*entities.User, error)
	GetAll(page, limit int) ([]entities.User, error)
	GetByID(id uint) (*entities.User, error)
	Update(id uint, request *request.UserUpdateRequest) (*entities.User, error)
	Delete(id uint) error
	CheckPassword(u *entities.User, password string) error
	GenerateResetToken(email string) (string, error)
	ResetPassword(newPassword string, userEmail string) error
	ValidateResetToken(token string) (*entities.PasswordResetToken, error)
	MarkTokenAsUsed(tokenID uint) error
	GetByEmail(email string) (*entities.User, error)
}

func NewUserService(userRepo repositories.UserRepositoryInterface, passwordResetTokenRepo repositories.PasswordResetTokenRepositoryInterface) *UserService {
	return &UserService{userRepo: userRepo, passwordResetTokenRepo: passwordResetTokenRepo}
}

func (s *UserService) Create(request *request.UserRequest) (*entities.User, error) {
	existingUser, _ := s.userRepo.GetByUsername(request.Username)
	if existingUser != nil {
		return nil, errors.New("username already exists")
	}

	role := request.Role
	if role == "" {
		role = entities.RolePlayer
	} else {
		if err := s.ValidateRole(role); err != nil {
			return nil, err
		}
	}

	user := &entities.User{
		Username: request.Username,
		Password: request.Password,
		Email:    request.Email,
		Role:     role,
	}

	if err := s.HashPassword(user); err != nil {
		return nil, err
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) GetAll(page, limit int) ([]entities.User, error) {
	return s.userRepo.GetAll(page, limit)
}

func (s *UserService) GetByID(id uint) (*entities.User, error) {
	return s.userRepo.GetByID(id)
}

func (s *UserService) Update(id uint, request *request.UserUpdateRequest) (*entities.User, error) {
	user, err := s.userRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if request.Password != "" {
		user.Password = request.Password
		if err := s.HashPassword(user); err != nil {
			return nil, err
		}
	}

	if request.Email != "" {
		user.Email = request.Email
	}

	if request.Username != "" {
		user.Username = request.Username
	}

	if request.Role != "" {
		if err := s.ValidateRole(request.Role); err != nil {
			return nil, err
		}
		user.Role = request.Role
	}

	return s.userRepo.Update(user)
}

func (s *UserService) Delete(id uint) error {
	return s.userRepo.Delete(id)
}

func (s *UserService) HashPassword(u *entities.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (s *UserService) CheckPassword(u *entities.User, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

func (s *UserService) ValidateRole(role string) error {
	if role != "" && role != entities.RolePlayer && role != entities.RoleAdmin {
		return errors.New("invalid role: must be either 'player' or 'admin'")
	}
	return nil
}

func (s *UserService) GenerateResetToken(email string) (string, error) {
	existingToken, _ := s.passwordResetTokenRepo.GetByEmail(email)

	generateNewToken := func() (string, error) {
		token := make([]byte, 16)
		if _, err := rand.Read(token); err != nil {
			return "", errors.New("failed to generate reset token")
		}
		tokenString := hex.EncodeToString(token)
		expiresAt := time.Now().Add(1 * time.Hour)

		if existingToken != nil {
			existingToken.Token = tokenString
			existingToken.ExpiredAt = expiresAt
			existingToken.IsUsed = false

			s.passwordResetTokenRepo.Update(existingToken)
		} else {
			resetToken := &entities.PasswordResetToken{
				Email:     email,
				Token:     tokenString,
				ExpiredAt: expiresAt,
				IsUsed:    false,
			}

			if err := s.passwordResetTokenRepo.Create(resetToken); err != nil {
				return "", err
			}
		}

		fmt.Printf("Send email to %s with reset link: http://localhost:8080/api/user/reset-password?token=%s\n", email, tokenString)
		return tokenString, nil
	}

	if existingToken != nil {
		if time.Now().After(existingToken.ExpiredAt) || existingToken.IsUsed {
			return generateNewToken()
		} else {
			return "", errors.New("email already sent")
		}
	}

	return generateNewToken()
}

func (s *UserService) ResetPassword(newPassword string, userEmail string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	return s.userRepo.UpdatePassword(userEmail, string(hashedPassword))
}

func (s *UserService) ValidateResetToken(token string) (*entities.PasswordResetToken, error) {
	resetToken, err := s.passwordResetTokenRepo.GetByToken(token)
	if err != nil {
		return nil, fmt.Errorf("invalid token")
	}
	return resetToken, nil
}

func (s *UserService) MarkTokenAsUsed(tokenID uint) error {
	return s.passwordResetTokenRepo.MarkTokenAsUsed(tokenID)
}

func (s *UserService) GetByEmail(email string) (*entities.User, error) {
	return s.userRepo.GetByEmail(email)
}
