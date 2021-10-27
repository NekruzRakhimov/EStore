package repository

import (
	"EStore/db"
	"EStore/models"
	"fmt"
	"github.com/pkg/errors"
)

// Получение корзины пользователя
func GetUserCart(userId int) (goods []models.Good, err error) {
	sqlQuery := `SELECT g.id, g.name, g.description, rg.count
				FROM goods g, reserved_goods rg
				WHERE g.id = good_id AND rg.user_id = ? ORDER BY id`
	if err := db.GetDBConn().Raw(sqlQuery, userId).Scan(&goods).Error; err != nil {
		return nil, err
	}

	return goods, nil
}

// Есть ли такая запись в таблице reserved_goods
func IsInReservedGoods(reservedGood models.ReservedGood) (bool, models.ReservedGood) {
	db.GetDBConn().Table("reserved_goods").Where("user_id = ? AND good_id = ?",
		reservedGood.UserId, reservedGood.GoodId).Scan(&reservedGood)
	if reservedGood.Id > 0 {
		return true, reservedGood
	} else {
		return false, models.ReservedGood{}
	}
}

// Увеличение количества товара в корзине пользователя
func IncrementCountOfReservedGood(currentCount int, reservedGood models.ReservedGood) error {
	if err := db.GetDBConn().Table("reserved_goods").First(&reservedGood).Error; err != nil {
		return err
	}

	good, err := GetGoodById(reservedGood.GoodId)
	if err != nil {
		return err
	}

	reservedGood.Count += currentCount
	if good.Count < currentCount {
		return errors.New(fmt.Sprintf("У нас нет столько продуктов. Вы просите [%d], а у нас на складе [%d]", currentCount, good.Count))
	} else {
		tx := db.GetDBConn().Begin()

		good.Count -= currentCount
		if err := tx.Table("goods").Save(&good).Error; err != nil {
			tx.Rollback()
			return err
		}

		if err := tx.Table("reserved_goods").Save(&reservedGood).Error; err != nil {
			tx.Rollback()
			return err
		}

		return tx.Commit().Error
	}
}

// Добавление товара в корзину пользователя
func ReserveGood(reservedGood models.ReservedGood) error {
	good, err := GetGoodById(reservedGood.GoodId)
	if err != nil {
		return err
	}

	if good.Count < reservedGood.Count {
		return errors.New(fmt.Sprintf("У нас нет столько продуктов. Вы просите [%d], а у нас на складе [%d]", reservedGood.Count, good.Count))
	} else {
		tx := db.GetDBConn().Begin()

		good.Count -= reservedGood.Count
		if err := tx.Table("goods").Save(&good).Error; err != nil {
			tx.Rollback()
			return err
		}

		if err := tx.Table("reserved_goods").Save(&reservedGood).Error; err != nil {
			tx.Rollback()
			return err
		}

		return tx.Commit().Error
	}
}

// Удаление товара из корзины пользователя
func DeleteGoodFromCart(reservedGood models.ReservedGood) error {
	return  db.GetDBConn().Table("reserved_goods").Where("good_id = ? AND user_id = ?",
		reservedGood.GoodId, reservedGood.UserId).Delete(&reservedGood).Error
}