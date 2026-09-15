package repository

import "merchandise-service/domain"

type UserRepository interface {
	Check(user *domain.User) error
}

type MerchRepository interface {
	Buy(merch *domain.Merch) error
}
