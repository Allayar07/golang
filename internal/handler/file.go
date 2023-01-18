package handler

import (
	servicePKG "file_work/internal/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func (h *Handler) Download(c *gin.Context) {
	file, err := c.FormFile("File")
	if err != nil {
		ErrorMessage(c, http.StatusBadRequest, err.Error())
		return
	}

	obj, err := file.Open()
	if err != nil {
		ErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := h.service.File.Upload(obj, file.Size, file.Filename); err != nil {
		ErrorMessage(c, http.StatusInternalServerError, err.Error())
		fmt.Println(err)
		return
	}

	defer obj.Close()
	//

	c.JSON(http.StatusOK, map[string]interface{}{
		"file": "ok. downloaded into uploads",
	})
}

func (h *Handler) Delete(c *gin.Context) {
	FileName := c.Param("any")

	if err := h.service.Remove(FileName); err != nil {
		ErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *Handler) OpenFile(c *gin.Context) {
	fileName := c.Param("any")

	//servicePKG.Directory = "uploads/"
	fileBytes, err := os.ReadFile(servicePKG.Directory + fileName)
	if err != nil {
		ErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	_, err = c.Writer.Write(fileBytes)
	if err != nil {
		ErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

}

// ChangeIMG  changing IMG parameters
func (h *Handler) ChangeIMG(c *gin.Context) {
	fileName := c.Param("any")

	requestValues := c.Request.URL.Query()

	height, _ := strconv.Atoi(requestValues.Get("height"))

	width, _ := strconv.Atoi(requestValues.Get("width"))

	formatOfIMG := requestValues.Get("format")

	percent, _ := strconv.ParseFloat(requestValues.Get("percent"), 64)

	if err := h.service.File.ResizeIMG(servicePKG.Directory+fileName, fileName, formatOfIMG, width, height, percent, c.Writer); err != nil {
		ErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	if formatOfIMG == "webp" {
		file := filepath.Ext(fileName)
		fileName = strings.TrimSuffix(fileName, file)

		if err := h.service.File.ResizeWebp(servicePKG.WebpUploads+fileName+".webp", width, height, percent, c.Writer); err != nil {
			ErrorMessage(c, http.StatusBadRequest, err.Error())
			return
		}

	}
}
