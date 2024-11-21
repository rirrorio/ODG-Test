package controllers

import (
	"encoding/json"
	"net/http"
	"sample-api-go/models"
	"sample-api-go/services"
	"strconv"
)

func MakeRedemption(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		CustomerID uint   `json:"customer_id"`
		VoucherIDs []uint `json:"voucher_ids"`
	}

	// Decode the request body into the struct
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Prepare a new transaction object
	transaction := models.Transaction{
		CustomerID: requestData.CustomerID,
	}

	// Call the service to process the redemption
	transactions, totalPoints, err := services.MakeRedemption(&transaction, requestData.VoucherIDs)
	if err != nil {
		http.Error(w, "Failed to make redemption", http.StatusInternalServerError)
		return
	}

	// Respond with the transaction details and total points
	response := struct {
		Transactions []models.Transaction `json:"transactions"`
		TotalPoints  int                  `json:"total_points"`
	}{
		Transactions: transactions,
		TotalPoints:  totalPoints,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func GetTransactionDetail(w http.ResponseWriter, r *http.Request) {
	transactionID, _ := strconv.Atoi(r.URL.Query().Get("transactionId"))

	transaction, err := services.GetTransactionDetail(uint(transactionID))
	if err != nil {
		http.Error(w, "Transaction not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(transaction)
}
