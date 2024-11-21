package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"sample-api-go/models"
	"sample-api-go/services"
)

func CreateBrand(writer http.ResponseWriter, request *http.Request) {

	//check httpMethod
	if request.Method != http.MethodPost {
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var brand models.Brand
	//decoding request body
	error := json.NewDecoder(request.Body).Decode(&brand)

	if error != nil {
		http.Error(writer, error.Error(), http.StatusBadRequest)
		return
	}

	//call service layer to handle further processing and DB operations
	createdBrand, error := services.CreateBrand(&brand)
	if error != nil {
		log.Printf("Error creating brand: %v", error)

		http.Error(writer, error.Error(), http.StatusInternalServerError)
		return
	}

	//creating a success messag
	writer.Header().Set("Content-Type", "aplication/json")

	successMessage := struct {
		Message string `json:"message"`
	}{
		Message: "Success creating brand " + createdBrand.Name,
	}

	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(successMessage)
}
