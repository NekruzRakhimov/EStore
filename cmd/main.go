package main

import (
	"EStore/db"
	"EStore/routes"
	"EStore/utils"
)

func main() {
	// Чтение config'ов
	utils.ReadSettings()

	// Соединение с БД
	db.StartDbConnection()

	// Запуск всех роутов
	routes.RunAllRoutes()
}
