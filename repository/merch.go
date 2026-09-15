package repository

import "database/sql"

type PostgresMerchRepository struct {
	db *sql.DB
}

func NewPostgresMerchRepository(db *sql.DB) *PostgresMerchRepository {
	return &PostgresMerchRepository{
		db: db,
	}
}
