package sqlStore

import (
	store "api/internal/app/stores"
	"database/sql"

	_ "github.com/lib/pq"
)

type SQLstore struct {
	db *sql.DB
	userRepository *UserRepository
}

func NewDB(db *sql.DB) *SQLstore {
	return &SQLstore{
		db: db,
	}
}

func (s *SQLstore) Users() store.UsersRepository {
	if s.userRepository != nil {
		return s.userRepository
	}
	s.userRepository = &UserRepository{
		store: s,
	}

	return s.userRepository
}
