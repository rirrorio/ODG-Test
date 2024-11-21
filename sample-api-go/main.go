package main

import (
	"log"
	"net/http"
	"sample-api-go/configs"
	"sample-api-go/models"
	"sample-api-go/routes"
)

func main() {
	configs.ConnectDB()

	//run AutoMigration to create or update the database schema if there is changes in class
	err := configs.DB.AutoMigrate(
		&models.Brand{},
		&models.Customer{},
		&models.Voucher{},
		&models.Transaction{},
	)

	if err != nil {
		log.Fatalf("Error migrating the database : %v", err)
	} else {
		log.Println("Database migration completed successfully.")
	}

	//register routes
	routes.SetupRoutes()

	// Start the HTTP server
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
