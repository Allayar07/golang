package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) AccessPage(c *gin.Context) {
	token := c.GetHeader("Authorization")

	if token == "" {
		ErrorMessage(c, http.StatusUnauthorized, "empty auth")
		return
	}

	userId, err := h.service.Admin.ParseToken(token)
	if err != nil {
		ErrorMessage(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set("user_id", userId)
}
