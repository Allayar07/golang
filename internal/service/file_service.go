package service

import (
	"github.com/sunshineplan/imgconv"
	"io"
	"log"
	"mime/multipart"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const (
	Directory   = "uploads/"
	WebpUploads = "webpUploads/"
)

type FileService struct{}

func NewFileService() *FileService {
	return &FileService{}
}

func (s *FileService) Upload(obj multipart.File, num int64, fileName string) error {
	localFile, err := os.Create(Directory + fileName)
	if err != nil {
		if err := os.Mkdir(Directory, 0750); err != nil {
			return err
		}

		localFileErr, err := os.Create(Directory + fileName)
		if err != nil {
			return err
		}

		_, err = io.CopyN(localFileErr, obj, num)
		if err != nil {
			return err
		}

		if err := localFileErr.Close(); err != nil {
			return err
		}

		return err
	}

	_, err = io.CopyN(localFile, obj, num)
	if err != nil {
		return err
	}

	defer localFile.Close()

	return nil
}

func (s *FileService) Remove(fileName string) error {
	if err := os.Remove(Directory + fileName); err != nil {
		return err
	}

	return nil
}

func (s *FileService) ResizeWebp(imgPath string, width, height int, percent float64, writer io.Writer) error {
	img, err := imgconv.Open(imgPath)
	if err != nil {
		return err
	}

	mark := imgconv.Resize(img, &imgconv.ResizeOption{Width: width, Height: height, Percent: percent})

	if err := imgconv.Write(writer, mark, &imgconv.FormatOption{}); err != nil {
		return err
	}

	return nil
}

func (s *FileService) ResizeIMG(imgPath, fileName, format string, width, height int, percent float64, writer io.Writer) error {

	img, err := imgconv.Open(imgPath)
	if err != nil {
		return err
	}

	mark := imgconv.Resize(img, &imgconv.ResizeOption{Width: width, Height: height, Percent: percent})

	formatOfIMG, err := FormatOfImage(format, imgPath, fileName)
	if err != nil {
		return err
	}

	if width != 0 || height != 0 || percent != 0 {
		if err := imgconv.Write(writer, mark, &formatOfIMG); err != nil {
			return err
		}
	}

	if err := imgconv.Write(writer, img, &formatOfIMG); err != nil {
		return err
	}

	return nil
}

// FormatOfImage is format of images
func FormatOfImage(format, imgPath, fileName string) (imgconv.FormatOption, error) {
	var formatOfIMG imgconv.FormatOption

	switch format {
	case "webp":
		file := filepath.Ext(fileName)
		fileName = strings.TrimSuffix(fileName, file)
		cmd := exec.Command("ffmpeg", "-i", imgPath, "-c:v", "libwebp", WebpUploads+fileName+".webp")
		if err := cmd.Start(); err != nil {
			return imgconv.FormatOption{}, err
		}

		log.Println(cmd)

	case "pdf":
		formatOfIMG = imgconv.FormatOption{Format: imgconv.PDF}
	case "png":
		formatOfIMG = imgconv.FormatOption{Format: imgconv.PNG}
	case "jpg":
		formatOfIMG = imgconv.FormatOption{Format: imgconv.JPEG}
	case "gif":
		formatOfIMG = imgconv.FormatOption{Format: imgconv.GIF}
	case "tiff":
		formatOfIMG = imgconv.FormatOption{Format: imgconv.TIFF}
	case "bmp":
		formatOfIMG = imgconv.FormatOption{Format: imgconv.BMP}
	default:
		formatOfIMG = imgconv.FormatOption{}
	}

	return formatOfIMG, nil
}
