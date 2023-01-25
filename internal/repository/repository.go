package repository

import (
	"github.com/jmoiron/sqlx"
)

type AdminRepo interface {
}

type Repository struct {
	AdminRepo
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
