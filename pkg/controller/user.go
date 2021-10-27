package controller

import (
	"EStore/models"
	"EStore/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// Получение корзины пользователя
func GetUserCart(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	goods, err := service.GetUserCart(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"goods": goods})
}

// Добавление товара в корзину пользователя
func AddToCart(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	reservedGood := models.ReservedGood{UserId: userId}
	if err := c.BindJSON(&reservedGood); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := service.AddToCart(reservedGood); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"reason": "товар успешно добавлен в корзину"})
}

// Удаление товара из корзины пользователя
func DeleteGoodFromCart(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	goodIdStr := c.Param("good_id")
	goodId, err := strconv.Atoi(goodIdStr)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	reservedGood := models.ReservedGood{GoodId: goodId, UserId: userId}
	if err := service.DeleteGoodFromCart(reservedGood); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"reason": "товар успешно удалён из вашей корзины"})
}