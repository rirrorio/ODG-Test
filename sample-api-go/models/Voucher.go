package models

import "gorm.io/gorm"

type Voucher struct {
	gorm.Model
	Name         string        `gorm:"not null"`
	CostInPoints int           `gorm:"not null"`
	BrandID      uint          `gorm:"not null"`
	Brand        Brand         `gorm:"foreignKey:BrandID;references:ID"` // defining Many-to-One relation with Brand
	Transactions []Transaction `gorm:"foreignKey:VoucherID"`             // defining One-to-Many relation with Transaction
}

func (Voucher) TableName() string {
	return "Vouchers"
}
