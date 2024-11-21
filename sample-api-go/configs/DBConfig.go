package configs

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

//initialize database connection

func ConnectDB() {
	//db connection string, ADJUST IF NEEDED

	dsn := "host=localhost user=postgres password=postgres dbname=go-sample-db port=5432 sslmode=disable"

	//connect to database

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("DB CONNECTION ERROR : %v", err)
	}
	log.Println("DB CONNECTED SUCCESFULLY")
}
