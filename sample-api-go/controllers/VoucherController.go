package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"sample-api-go/models"
	"sample-api-go/services"
	"strconv"
	"strings"
)

func CreateVoucher(writer http.ResponseWriter, request *http.Request) {
	//check httpMethod
	if request.Method != http.MethodPost {
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var voucher models.Voucher

	//decode json payload
	error := json.NewDecoder(request.Body).Decode(&voucher)

	if error != nil {
		http.Error(writer, error.Error(), http.StatusBadRequest)
		return
	}

	//validate required fields
	var errorMessage string

	switch true {
	case strings.TrimSpace(voucher.Name) == "":
		errorMessage = "Voucher Name cannot be empty"
	case voucher.CostInPoints <= 0:
		errorMessage = "CostInPoints must be greater than zero"
	case voucher.BrandID == 0:
		errorMessage = "BrandID is required"
	}

	if errorMessage != "" {
		http.Error(writer, errorMessage, http.StatusBadRequest)
		return
	}

	//call service layer to handle further processing and DB operations
	createdVoucher, error := services.CreateVoucher(&voucher)
	if error != nil {
		log.Printf("Error creating voucher: %v", error)

		http.Error(writer, error.Error(), http.StatusInternalServerError)
		return
	}

	//creating a success message
	writer.Header().Set("Content-Type", "application/json")

	successMessage := struct {
		Message string `json:"message"`
	}{
		Message: "Success creating voucher " + createdVoucher.Name,
	}

	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(successMessage)
}

func GetVoucher(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	voucherID, _ := strconv.Atoi(request.URL.Query().Get("id"))

	voucher, err := services.GetVoucher(uint(voucherID))
	if err != nil {
		http.Error(writer, "Voucher not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(writer).Encode(voucher)
}

func GetVouchersByBrand(writer http.ResponseWriter, request *http.Request) {
	//check httpMethod
	if request.Method != http.MethodGet {
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	brandID, _ := strconv.Atoi(request.URL.Query().Get("id"))

	vouchers, err := services.GetVouchersByBrandID(uint(brandID))

	if err != nil {
		http.Error(writer, "Failed to fetch vouchers by brand: ", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(writer).Encode(vouchers)
}
