package services

import (
	"errors"
	"log"
	"sample-api-go/configs"
	"sample-api-go/models"
)

func CreateVoucher(voucher *models.Voucher) (*models.Voucher, error) {
	// Check if the brand exists
	var brand models.Brand
	if err := configs.DB.First(&brand, voucher.BrandID).Error; err != nil {
		log.Println("Brand not found:", err)
		return nil, errors.New("brand does not exist")
	}

	// Save to database
	result := configs.DB.Create(voucher)
	if result.Error != nil {
		log.Println("Error creating voucher:", result.Error)
		return nil, result.Error
	}

	return voucher, nil
}

func GetVoucher(voucherID uint) (*models.Voucher, error) {
	var voucher models.Voucher
	result := configs.DB.First(&voucher, voucherID)

	if result.Error != nil {
		log.Println("Error fetching voucher : ", result.Error)
		return nil, result.Error
	}

	return &voucher, nil
}

func GetVouchersByBrandID(brandID uint) ([]models.Voucher, error) {
	var vouchers []models.Voucher

	result := configs.DB.Where("brand_id = ?", brandID).Find(&vouchers)

	if result.Error != nil {
		log.Println("Error fetching vouchers by brand Id : ", result.Error)
		return nil, result.Error
	}

	return vouchers, nil
}
