package repository

import (
	"EStore/db"
	"EStore/models"
	"fmt"
	"github.com/pkg/errors"
)

// Получение id всех пользоватлей
func GetAllUsersIds() (userIds []models.User, err error) {
	sqlQuery := "SELECT id FROM users WHERE type_id = 1"
	if err := db.GetDBConn().Raw(sqlQuery).Scan(&userIds).Error; err != nil {
		return nil, err
	}

	return userIds, nil
}

// Создание нового товара
func CreateGood(good models.Good) (int, error) {
	if err := db.GetDBConn().Table("goods").Create(&good).Error; err != nil {
		return 0, err
	}

	return good.Id, nil
}

// Увеличение количесва товара в БД
func IncrementGoodCount(good models.Good) error {
	currentGoodCount := good.Count

	good, err := GetGoodById(good.Id)
	if err != nil {
		return err
	}

	good.Count += currentGoodCount
	if err := db.GetDBConn().Table("goods").Save(&good).Error; err != nil {
		return err
	}

	return nil
}

// Уменьшение количесва товара в БД
func DecrementGoodCount(good models.Good) error {
	currentGoodCount := good.Count

	good, err := GetGoodById(good.Id)
	if err != nil {
		return err
	}

	if good.Count > currentGoodCount {
		good.Count -= currentGoodCount
		if err := db.GetDBConn().Table("goods").Save(&good).Error; err != nil {
			return err
		}
	} else {
		return errors.New(fmt.Sprintf("У нас нет столько продуктов. Вы просите [%d], а у нас на складе [%d]",
			currentGoodCount, good.Count))
	}

	return nil
}

// Удаление товара из БД
func DeleteGood(id int) error {
	good := models.Good{Id: id}

	return db.GetDBConn().Table("goods").Delete(&good).Error
}