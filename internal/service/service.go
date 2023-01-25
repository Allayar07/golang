package service

import (
	"file_work/internal/repository"
	"io"
	"mime/multipart"
)

type File interface {
	Remove(fileName string) error
	Upload(obj multipart.File, num int64, fileName string) error
	ResizeWebp(imgPath string, width, height int, percent float64, writer io.Writer) error
	ResizeIMG(imgPath, fileName, format string, width, height int, percent float64, writer io.Writer) error
}

type Service struct {
	File
}

func NewService(repos *repository.Repository) *Service {
	return &Service{

		File: NewFileService(),
	}
}
