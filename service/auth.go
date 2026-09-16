package service

import (
	"merchandise-service/domain"

	"github.com/golang-jwt/jwt/v5"
	_ "github.com/lib/pq"
)

type AuthService struct {
	repo UserRepository
}

func NewAuthService(repo UserRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (authService AuthService) Auth(user *domain.User) (string, error) {
	authService.repo.Check(user)
	err := authService.repo.Check(user)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.GetUsername(),
	})
	tokenString, err := token.SignedString([]byte("token_secret"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
