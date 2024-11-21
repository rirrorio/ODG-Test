package models

import "gorm.io/gorm"

type Brand struct {
	gorm.Model
	Name     string    `gorm:"unique;not null"`    // Brand Name
	Vouchers []Voucher `gorm:"foreignKey:BrandID"` // Defining one-to-many relations with vouchers
}

func (Brand) TableName() string {
	return "Brands"
}
