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
		api.POST("/uploads", h.Download)
		api.DELETE("/uploads/:any", h.Delete)
		api.GET("/uploads/open/:any", h.OpenFile)
		api.GET("/uploads/get/:any", h.ChangeIMG)
	}

	login := routes.Group("/sign_in")
	{
		login.POST("/login", h.signIn)
	}

	auth := routes.Group("/auth", h.AccessPage)
	{
		auth.POST("/admin", h.AdminHandler)
	}

	return routes

}
