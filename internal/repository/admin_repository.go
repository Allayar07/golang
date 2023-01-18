package repository

import (
	"file_work/internal/model"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type AdminRepository struct {
	db *sqlx.DB
}

func NewAdminRepository(db *sqlx.DB) *AdminRepository {
	return &AdminRepository{
		db: db,
	}
}

func (r *AdminRepository) CreatAdmin(user model.Admin) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (name, email, hash_password, admin_role) values ($1, $2, $3, $4) RETURNING id", "admin")

	row := r.db.QueryRow(query, user.Name, user.Email, user.Password, user.AdminRole)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
