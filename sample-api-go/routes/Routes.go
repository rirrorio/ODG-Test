package routes

import (
	"net/http"
	"sample-api-go/controllers"
)

func SetupRoutes() {
	// POST /brand
	http.HandleFunc("/brand", controllers.CreateBrand)

	// POST /voucher & GET /voucher
	http.HandleFunc("/voucher", func(writer http.ResponseWriter, request *http.Request) {

		if request.Method == http.MethodPost {
			controllers.CreateVoucher(writer, request)
		} else if request.Method == http.MethodGet {
			controllers.GetVoucher(writer, request)
		} else {
			http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	// GET /voucher/brand
	http.HandleFunc("/voucher/brand", controllers.GetVouchersByBrand)
	// // POST /transaction/redemption
	http.HandleFunc("/transaction/redemption", controllers.MakeRedemption)
	// // GET transaction/redemption
	// http.HandleFunc("/transaction/redemption", controllers.GetTransactionDetail)

	http.ListenAndServe(":8080", nil)
}
