package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) AdminHandler(c *gin.Context) {
	id, ok := c.Get("user_id")
	if !ok {
		ErrorMessage(c, http.StatusBadRequest, "error getting user context")
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"user Id": id,
	})

}

func (h *Handler) signIn(c *gin.Context) {
	token, err := h.service.Admin.GenerateToken()

	if err != nil {
		ErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
