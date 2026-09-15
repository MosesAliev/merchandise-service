package usecase

import (
	"merchandise-service/domain"
	"merchandise-service/postgres"
	"merchandise-service/repository"
	"merchandise-service/service"
)

func Auth(username, password string) (string, error) {
	user := domain.NewUser(username, password)
	repo := repository.NewPostgresUserRepository(postgres.DB)
	authService := service.NewAuthService(repo)
	token, err := authService.Auth(user)
	if err != nil {
		return "", err
	}

	return token, nil
}
