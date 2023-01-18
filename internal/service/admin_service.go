package service

import (
	"crypto/sha1"
	"errors"
	"file_work/internal/model"
	"file_work/internal/repository"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type AdminService struct {
	repos repository.AdminRepo
}

type TokenClaims struct {
	jwt.StandardClaims
	Userid string `json:"userid"`
}

func NewAdminService(repo repository.AdminRepo) *AdminService {
	return &AdminService{repos: repo}
}

func (s *AdminService) Create(user model.Admin) (int, error) {
	user.Password = HashFunc(user.Password)
	return s.repos.CreatAdmin(user)
}

func (s *AdminService) GenerateToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(1 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		Userid: "1",
	})

	return token.SignedString([]byte("oiqwopiepwba23342b"))
}

func (s *AdminService) ParseToken(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &TokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte("oiqwopiepwba23342b"), nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*TokenClaims)
	if !ok {
		return "", errors.New("token claim are not that type")
	}

	return claims.Userid, nil
}

func HashFunc(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte("jbksagdiwei7w89")))
}
