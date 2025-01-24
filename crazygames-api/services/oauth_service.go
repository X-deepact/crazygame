package services

import (
	"context"
	"errors"

	"crazygames.io/config"
	"crazygames.io/entities"
	"crazygames.io/repositories"
	"golang.org/x/oauth2"
)

type OAuthServiceInterface interface {
	GetGoogleAuthURL(state string) string
	ExchangeCodeForToken(ctx context.Context, code string) (*oauth2.Token, error)
	HandleGoogleUser(userInfo map[string]interface{}) error
}

type OAuthService struct {
	userRepo repositories.UserRepositoryInterface
}

func NewOAuthService(userRepo repositories.UserRepositoryInterface) *OAuthService {
	return &OAuthService{
		userRepo: userRepo,
	}
}

func (s *OAuthService) GetGoogleAuthURL(state string) string {
	authURL := config.GoogleOauthConfig.AuthCodeURL(state, oauth2.AccessTypeOffline)
	return authURL
}

func (s *OAuthService) ExchangeCodeForToken(ctx context.Context, code string) (*oauth2.Token, error) {
	token, err := config.GoogleOauthConfig.Exchange(ctx, code)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (s *OAuthService) HandleGoogleUser(userInfo map[string]interface{}) error {
	email, ok := userInfo["email"].(string)
	if !ok || email == "" {
		return errors.New("invalid user info: email not found")
	}

	// Check if the user already exists
	existingUser, _ := s.userRepo.GetByUsername(email)
	if existingUser != nil {
		return nil // User already exists, no further action required
	}

	// Create a new user
	user := &entities.User{
		Username: email,
		Email:    email,
		Role:     entities.RolePlayer,
		Password: "",
	}

	return s.userRepo.Create(user)
}
