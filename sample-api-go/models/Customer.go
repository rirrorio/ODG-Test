package models

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	Name         string        `gorm:"not null"`
	Email        string        `gorm:"unique;not null"`
	Transactions []Transaction `gorm:"foreignKey:CustomerID"` //defining One-To-Many relationship with Transaction
}

func (Customer) TableName() string {
	return "Customers"
}
