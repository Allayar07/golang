package handler

import "github.com/gin-gonic/gin"

type MessageError struct {
	Message string `json:"message"`
}

func ErrorMessage(c *gin.Context, statusCode int, message string) {
	c.AbortWithStatusJSON(statusCode, MessageError{
		Message: message,
	})
}
