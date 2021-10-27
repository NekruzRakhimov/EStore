package controller

import (
	"EStore/models"
	"EStore/pkg/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// Получение списка корзин всех пользователей
func GetAllCarts(c *gin.Context) {
	carts, err := service.GetAllCarts()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"carts": carts})
}

// Создание продукта
func CreateGood(c *gin.Context) {
	var good models.Good
	if err := c.BindJSON(&good); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	goodId, err := service.CreateGood(good)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"good_id": goodId})
}

// Изменение количества продукта
func СrementGoodCount(c *gin.Context) {
	// 1. Читаем тело запроса
	var good models.Good
	if err := c.BindJSON(&good); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// 2. Определяем тип опрезации {прибавить или отнять}
	operation := c.Param("type_of_operation")
	if operation == "inc" {
		if err := service.IncrementGoodCount(good); err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	} else if operation == "dec" {
		if err := service.DecrementGoodCount(good); err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	} else {
		newErrorResponse(c, http.StatusInternalServerError, "тип операции не найден")
		return
	}

	c.JSON(http.StatusOK, gin.H{"reason": "количество продукта успешно обновлено"})
}

// Удаление продукта из БД
func DeleteGood(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := service.DeleteGood(id); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"reason": fmt.Sprintf("товар с id = %d успешно удалён", id)})
}