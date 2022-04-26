package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"golang.org/x/crypto/bcrypt"
)

type Users struct {
	ID                int    `json:"id"`
	Name              string `json:"name"`
	Email             string `json:"email"`
	Password          string `json:"password,omitempty"`
	EncryptedPassword string `json:"-"`
}




func (u *Users) Validate() error {
	return validation.ValidateStruct(u,
		validation.Field(&u.Name, validation.Required),
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Password, validation.By(RequiredIF(u.EncryptedPassword == "")), validation.Length(5, 100)))
}

func (u *Users) BeforeCreated() error {
	if len(u.Password) > 0 {
		encrypt, err := EncryptedString(u.Password)
		if err != nil {
			return err
		}
		u.EncryptedPassword = encrypt
	}

	return nil
}

func (u *Users) ParolGizle() {
	u.Password = ""
}

func (u *Users) ComparePassWord(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.EncryptedPassword), []byte(password)) == nil
}

func EncryptedString(s string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
