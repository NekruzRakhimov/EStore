package service

import (
	"EStore/models"
	"EStore/pkg/repository"
)

// Получение списка корзин всех пользвоателей
func GetAllCarts() (carts []models.CartsForManager, err error) {
	userIds, err := repository.GetAllUsersIds()
	if err != nil {
		return nil, err
	}

	for _, userId := range userIds {
		userCart, err := repository.GetUserCart(userId.Id)
		if err != nil {
			return nil, err
		}

		cartForManager := models.CartsForManager{UserId: userId.Id, Cart: userCart}
		carts = append(carts, cartForManager)
	}

	return carts, nil
}

// Создание нового товара
func CreateGood(good models.Good) (int, error) {
	id, err := repository.CreateGood(good)
	if err != nil {
		return 0, err
	}

	return id, nil
}

// Увеличение количесва товара в БД
func IncrementGoodCount(good models.Good) error {
	return repository.IncrementGoodCount(good)
}

// Уменьшение количесва товара в БД
func DecrementGoodCount(good models.Good) error {
	return repository.DecrementGoodCount(good)
}

// Удаление товара из БД
func DeleteGood(id int) error {
	return repository.DeleteGood(id)
}