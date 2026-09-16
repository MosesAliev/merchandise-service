package service

import "merchandise-service/domain"

type UserRepository interface {
	Check(user *domain.User) error
}
