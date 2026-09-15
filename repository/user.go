package repository

import (
	"database/sql"
	"errors"
	"merchandise-service/domain"
	"merchandise-service/postgres"
)

type PostgresUserRepository struct {
	db *sql.DB
}

func NewPostgresUserRepository(db *sql.DB) *PostgresUserRepository {
	return &PostgresUserRepository{
		db: db,
	}
}

func (postgresUserRepository PostgresUserRepository) Check(user *domain.User) error {
	var username, password string
	if err := postgres.DB.QueryRow(getUsername, user.GetUsername()).Scan(&username); err != nil {
		if err == sql.ErrNoRows {
			return errors.New("invalid username")
		}

		return err
	}

	if err := postgres.DB.QueryRow(getPassword, user.GetPassword()).Scan(&password); err != nil {
		if err == sql.ErrNoRows {
			return errors.New("invalid password")
		}

		return err
	}

	return nil
}
