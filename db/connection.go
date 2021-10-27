package db

import (
	"EStore/utils"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

var database *gorm.DB

// ConnectToDB процесс соединение с БД
func ConnectToDB() *gorm.DB {
	settingParams := utils.AppSettings.PostgresParams

	connString := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		settingParams.Server, settingParams.Port,
		settingParams.User, settingParams.DataBase,
		settingParams.Password)

	db, err := gorm.Open("postgres", connString)

	if err != nil {
		log.Fatal("Couldn't connect to database", err.Error())
	}

	// enabling gorm log mode, used for debugging
	db.LogMode(true)

	db.SingularTable(true)

	return db
}

// StartDbConnection создаёт соединение с БД
func StartDbConnection() {
	database = ConnectToDB()
}

// GetDBConn функция которая получает соединения к БД глобально
func GetDBConn() *gorm.DB {
	return database
}
