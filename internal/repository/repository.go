package repository

import (
	"file_work/internal/model"
	"github.com/jmoiron/sqlx"
)

type AdminRepo interface {
	CreatAdmin(user model.Admin) (int, error)
}

type Repository struct {
	AdminRepo
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		AdminRepo: NewAdminRepository(db),
	}
}
