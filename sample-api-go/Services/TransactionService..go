package services

import (
	"errors"
	"log"
	"sample-api-go/configs"
	"sample-api-go/models"
)

func MakeRedemption(transaction *models.Transaction, voucherIDs []uint) ([]models.Transaction, int, error) {
	var totalPoints int
	var transactions []models.Transaction

	// Fetch all vouchers at once
	var vouchers []models.Voucher
	result := configs.DB.Where("id IN ?", voucherIDs).Find(&vouchers)
	if result.Error != nil {
		log.Println("Error fetching vouchers:", result.Error)
		return nil, 0, result.Error
	}

	// Create a map for fast lookup by voucher ID
	voucherMap := make(map[uint]models.Voucher)
	for _, voucher := range vouchers {
		voucherMap[voucher.ID] = voucher
	}

	// Iterate through voucherIDs and build the transactions
	for _, voucherID := range voucherIDs {
		voucher, found := voucherMap[voucherID]
		if !found {
			// If any voucher is not found, cancel all operations
			log.Printf("Voucher with ID %d not found. Cancelling redemption.\n", voucherID)
			return nil, 0, errors.New("voucher with ID " + string(voucherID) + " not found")
		}

		// Increment total cost in points
		totalPoints += voucher.CostInPoints

		// Create the transaction (no DB save yet)
		newTransaction := models.Transaction{
			CustomerID: transaction.CustomerID,
			VoucherID:  voucherID,
		}
		transactions = append(transactions, newTransaction)
	}

	// save all the transactions at once
	if len(transactions) > 0 {
		result := configs.DB.Create(&transactions)
		if result.Error != nil {
			log.Println("Error saving transactions:", result.Error)
			return nil, 0, result.Error
		}
	}

	return transactions, totalPoints, nil

}

func GetTransactionDetail(transactionID uint) (*models.Transaction, error) {
	var transaction models.Transaction
	result := configs.DB.First(&transaction, transactionID)
	if result.Error != nil {
		log.Println("Error fetching transaction:", result.Error)
		return nil, result.Error
	}
	return &transaction, nil
}
