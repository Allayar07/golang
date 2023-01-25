package handler

import (
	"file_work/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	routes := gin.Default()

	api := routes.Group("/api")
	{
		//endpoint for upload files into uploads folder
		api.POST("/uploads", h.Download)

		//endpoint for delete files from uploads folder
		api.DELETE("/uploads/:any", h.Delete)

		//endpoint for opening files
		api.GET("/uploads/open/:any", h.OpenFile)

		//endpoint for exchange image formats. for example: convert jpg format to png format,
		//format of images taken from query params
		api.GET("/uploads/get/:any", h.ChangeIMG)
	}

	return routes

}
