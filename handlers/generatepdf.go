package handlers

import (
	"CS372-Project/utils"
	"database/sql"
	"log"
	"net/http"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

func GeneratePdfHandler(w http.ResponseWriter, r *http.Request) {
	database, err := sql.Open("sqlite3", "./cs372-app.db")

	if err != nil {
		panic(err)
	}

	defer database.Close()
	if r.Method == "GET" {

		urlVehicleArgs, _ := r.URL.Query()["vehicleId"]
		urlCustomerArgs, _ := r.URL.Query()["customerId"]

		//Args is of type array; want just string at index 0
		urlVehicleId := urlVehicleArgs[0]
		urlCustomerId := urlCustomerArgs[0]

		vehicleInt, _ := strconv.Atoi(urlVehicleId)
		customerInt, _ := strconv.Atoi(urlCustomerId)

		urlVehicle := utils.SelectVehicleById(vehicleInt, database)
		urlCustomer := utils.SelectCustomerById(customerInt, database)

		pdfPath, err := utils.CreatePdf(urlVehicle, urlCustomer)

		if err != nil {
			panic(err)
		}

		http.ServeFile(w, r, string(pdfPath))
	}
}
