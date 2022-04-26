package store

import (
	"api/internal/app/model"
	"database/sql"
)



type UsersRepository interface {
	Create(*model.Users)  error
	Find(int) (*model.Users, error)
	FindBYemail(string) (*model.Users, error)
	GetAll() (*sql.Rows, error)
	Update(string, string, int) (*model.Users, error)
}