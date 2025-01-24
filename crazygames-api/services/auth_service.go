package services

import (
	"errors"
	"log"
	"time"

	"crazygames.io/config"
	"crazygames.io/entities"
	"crazygames.io/handler/request"
	"crazygames.io/repositories"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo repositories.UserRepositoryInterface
	secret   string // JWT secret
}

type AuthServiceInterface interface {
	Login(request *request.LoginRequest) (string, error)
	Register(request *request.RegisterRequest) (*entities.User, error)
}

func NewAuthService(userRepo repositories.UserRepositoryInterface) *AuthService {
	return &AuthService{
		userRepo: userRepo,
		secret:   config.AppConfig.JWTSecret,
	}
}

func (s *AuthService) Login(request *request.LoginRequest) (string, error) {
	user, err := s.userRepo.GetByEmail(request.Email)
	log.Println("User===> ", user)
	log.Println("err===> ", err)
	if err != nil || user == nil {
		return "", errors.New("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		return "", errors.New("incorrect password")
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"role":    user.Role,
		"exp":     time.Now().Add(time.Hour * 72).Unix(), // Token expires in 72 hours
	})

	tokenString, err := token.SignedString([]byte(s.secret))
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return tokenString, nil
}

func (s *AuthService) Register(request *request.RegisterRequest) (*entities.User, error) {
	existingUser, _ := s.userRepo.GetByEmail(request.Email)
	if existingUser != nil {
		return nil, errors.New("email already registered")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	user := &entities.User{
		Username: request.Username,
		Email:    request.Email,
		Password: string(hashedPassword),
		Role:     entities.RolePlayer, // Default role
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}
