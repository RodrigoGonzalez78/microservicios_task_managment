package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DSN = "host=172.17.0.2 user=postgres password=12345678 dbname=taskMicro port=5432"
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
