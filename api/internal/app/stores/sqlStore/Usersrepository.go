package sqlStore

import (
	"api/internal/app/model"
	store "api/internal/app/stores"
	"database/sql"
)

type UserRepository struct {
	store *SQLstore
}

func (req *UserRepository) Create(u *model.Users) error {

	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreated(); err != nil {
		return err
	}

	return req.store.db.QueryRow(
		"INSERT INTO adamlar (name, email, encrypted_password) VALUES ($1,$2,$3) RETURNING id",
		u.Name,
		u.Email,
		u.EncryptedPassword,
	).Scan(&u.ID)
}

func (req *UserRepository) Find(id int) (*model.Users, error) {
	u := &model.Users{}

	if err := req.store.db.QueryRow(
		"SELECT id, name, email, encrypted_password FROM adamlar WHERE id = $1",
		id).Scan(
		&u.ID,
		&u.Name,
		&u.Email,
		&u.EncryptedPassword,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrorNotFoundRecord
		}
		return nil, err
	}

	return u, nil
}

func (req *UserRepository) FindBYemail(email string) (*model.Users, error) {
	u := &model.Users{}

	if err := req.store.db.QueryRow(
		"SELECT id, name, email, encrypted_password FROM adamlar WHERE email = $1",
		email).Scan(
		&u.ID,
		&u.Name,
		&u.Email,
		&u.EncryptedPassword,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrorNotFoundRecord
		}
		return nil, err
	}

	return u, nil
}

func (req *UserRepository) GetAll() (*sql.Rows, error) {

	res, err := req.store.db.Query("SELECT * FROM adamlar")
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (req *UserRepository) Update(name, email string, id int) (*model.Users, error) {
	u := &model.Users{}
	q := `UPDATE adamlar SET name = $1, email = $2 WHERE id = $3 RETURNING *`
	err := req.store.db.QueryRow(q, name, email, id).Scan(
		&u.ID,
		&u.Name,
		&u.Email,
		&u.EncryptedPassword,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrorNotFoundRecord
		}

		return nil, err
	}

	return u, nil
}
