package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	CustomerID uint     `gorm:"not null"`
	VoucherID  uint     `gorm:"not null"`
	Customer   Customer `gorm:"foreignKey:CustomerID;references:ID"`
	Voucher    Voucher  `gorm:"foreignKey:VoucherID;references:ID"`
}

func (Transaction) TableName() string {
	return "Transactions"
}
