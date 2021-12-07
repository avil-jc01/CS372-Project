package handlers

import (
	"CS372-Project/utils"
	"database/sql"
	"log"
	"net/http"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

func DeleteAutoHandler(w http.ResponseWriter, r *http.Request) {

	database, err := sql.Open("sqlite3", "./cs372-app.db")

	if err != nil {
		panic(err)
	}

	defer database.Close()

	if r.Method == "DELETE" {
		log.Printf("D: HTTP DELETE on /delete-auto")
		urlVehicleArgs, _ := r.URL.Query()["vehicleId"]
		urlCustomerArgs, _ := r.URL.Query()["customerId"]

		//Args is of type array; want just string at index 0
		urlVehicleId := urlVehicleArgs[0]
		urlCustomerId := urlCustomerArgs[0]

		vehicleInt, _ := strconv.Atoi(urlVehicleId)
		customerInt, _ := strconv.Atoi(urlCustomerId)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		err := utils.DeleteAuto(vehicleInt, customerInt, database)
		if err != nil {
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		} else {
			http.Redirect(w, r, "/", http.StatusFound)
		}

	}
}
