package service

import (
	"file_work/internal/model"
	"file_work/internal/repository"
	"io"
	"mime/multipart"
)

type Admin interface {
	Create(user model.Admin) (int, error)
	GenerateToken() (string, error)
	ParseToken(tokenString string) (string, error)
}

type File interface {
	Remove(fileName string) error
	Upload(obj multipart.File, num int64, fileName string) error
	ResizeWebp(imgPath string, width, height int, percent float64, writer io.Writer) error
	ResizeIMG(imgPath, fileName, format string, width, height int, percent float64, writer io.Writer) error
}

type Service struct {
	Admin
	File
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Admin: NewAdminService(repos.AdminRepo),
		File:  NewFileService(),
	}
}
