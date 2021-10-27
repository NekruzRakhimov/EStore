package repository

import (
	"EStore/db"
	"EStore/models"
	"fmt"
)

// Получение списка всех товаров
func GetAllGoods() (goods []models.Good, err error) {
	sqlQuery := "SELECT * FROM goods"
	if err := db.GetDBConn().Raw(sqlQuery).Order("id").Scan(&goods).Error; err != nil {
		return nil, err
	}

	return goods, nil
}

// Проверка, есть ли в БД тип пользователя с таким именем
func TypeExists(name string) (int, bool) {
	var (
		count int
		userType models.UserType
	)
	sqlQuery := "SELECT * FROM user_types WHERE name = $1"

	db.GetDBConn().Raw(sqlQuery, name).Scan(&userType).Count(&count)
	if count > 0 {
		return userType.Id, true
	} else {
		return 0, false
	}
}

// Создание пользователя
func CreateUser(user models.User) (int, error) {
	if err := db.GetDBConn().Table("users").Omit("type").Create(&user).Error; err != nil {
		return 0, err
	}

	return user.Id, nil
}

// Получение пользователя
func GetUser(username, password string) (user models.User, err error) {
	sqlQuery := "SELECT * FROM users WHERE username = $1 AND password = $2"
	if err := db.GetDBConn().Raw(sqlQuery, username, password).Scan(&user).Error; err != nil {
		return models.User{}, err
	}

	return user, nil
}

// Получение товара по id
func GetGoodById(goodId int) (models.Good, error) {
	good := models.Good{Id: goodId}

	fmt.Println(">>>>>Repository] good: ", goodId)
	if err := db.GetDBConn().Table("goods").First(&good, good.Id).Error; err != nil {
		return models.Good{}, err
	}

	return good, nil
}