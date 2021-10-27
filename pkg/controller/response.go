package controller

import (
	"github.com/gin-gonic/gin"
	"log"
)

// Структура для ошибок
type errorResponse struct {
	Message string `json:"message"`
}

// Функция для вывода ошибок
func newErrorResponse(c *gin.Context, statusCode int, message string) {
	log.Printf(message)
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}
