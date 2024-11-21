package services

import (
	"log"
	"sample-api-go/configs"
	"sample-api-go/models"
)

func CreateBrand(brand *models.Brand) (*models.Brand, error) {
	//save to db
	result := configs.DB.Create(&brand)

	if result.Error != nil {
		log.Println("Error Creating Brand: ", result.Error)
		return nil, result.Error
	}

	// return the created brand
	return brand, nil
}
