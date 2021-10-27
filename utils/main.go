package main

import (
	"CS372-Project/handlers"
	"CS372-Project/utils"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
)


//small change
func main() {

	log.Println("D: Starting - Running CS372 App on port 8080")

	database, err := sql.Open("sqlite3", "./cs372-app.db")
	if err != nil {
		panic(err)
	}

	defer database.Close()

	custName := "customers"
	custDef := "(id INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT, address TEXT, city TEXT, state TEXT, zip INTEGER)"
	utils.CreateTable(custName, custDef, database)

	vehName := "vehicles"
	vehDef := "(id INTEGER PRIMARY KEY, vin TEXT, year INTEGER, make TEXT, model TEXT, purchase_price FLOAT, date_of_sale TEXT, customer_id INTEGER)"
	utils.CreateTable(vehName, vehDef, database)

	//http requests are routed to the endpoint with a multiplexer
	mux := http.NewServeMux()

	hh := http.HandlerFunc(handlers.HomeHandler)
	aah := http.HandlerFunc(handlers.AddAutoHandler)
	ach := http.HandlerFunc(handlers.AddCustomerHandler)

	mux.Handle("/", hh)
	mux.Handle("/add-auto", aah)
	mux.Handle("/add-customer", ach)

	//start the webserver
	http.ListenAndServe(":8080", mux)

}
