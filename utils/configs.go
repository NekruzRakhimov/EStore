package utils

import (
	"encoding/json"
	"fmt"
	"homis/models"
	"log"
	"os"
)

var (
	AppSettings models.Settings
)

// Чтение файла конфигурации
func ReadSettings() {
	fmt.Println("Starting reading settings file")
	configFile, err := os.Open("./settings-dev.json")
	if err != nil {
		log.Fatal("Couldn't open config file. Error is: ", err.Error())
	}

	defer func(configFile *os.File) {
		err := configFile.Close()
		if err != nil {
			log.Fatal("Couldn't close config file. Error is: ", err.Error())
		}
	}(configFile)

	fmt.Println("Starting decoding settings file")
	if err = json.NewDecoder(configFile).Decode(&AppSettings); err != nil {
		log.Fatal("Couldn't decode settings json file. Error is: ", err.Error())
	}

	log.Println(AppSettings)
	return
}
