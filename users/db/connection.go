package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DSN = "host=192.168.0.168 port=5435 user=postgres password=12345678 dbname=usuarios sslmode=disable"

var DB *gorm.DB

func DBConnection() {

	var dbError error
	DB, dbError = gorm.Open(postgres.Open(DSN), &gorm.Config{})

	if dbError != nil {
		log.Fatal(dbError)
	} else {
		log.Println("Base de datos conectada!!")
	}

}
