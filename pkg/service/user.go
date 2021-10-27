package service

import (
	"EStore/models"
	"EStore/pkg/repository"
)

// Получение корзины пользователя
func GetUserCart(userId int) (goods []models.Good, err error) {
	return repository.GetUserCart(userId)
}

// Добавление товара в корзину пользователя
func AddToCart(reservedGood models.ReservedGood) error {
	currentCount := reservedGood.Count
	exists, newReservedGood := repository.IsInReservedGoods(reservedGood)
	if exists {
		if err := repository.IncrementCountOfReservedGood(currentCount, newReservedGood); err != nil {
			return err
		}
	} else {
		if err := repository.ReserveGood(reservedGood); err != nil {
			return err
		}
	}

	return nil
}

// Удаление товара из корзины пользователя
func DeleteGoodFromCart(reservedGood models.ReservedGood) error {
	return repository.DeleteGoodFromCart(reservedGood)
}