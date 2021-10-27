package controller

import (
	"EStore/pkg/repository"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Получение списка всех товаров
func GetAllGoods(c *gin.Context) {
	goods, err := repository.GetAllGoods()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"goods": goods})
}
