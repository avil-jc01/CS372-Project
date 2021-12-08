package main

import (
	"CS372-Project/handlers"
	"CS372-Project/utils"
	"database/sql"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

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

	mux := http.NewServeMux()

	fileserver := http.FileServer(http.Dir("./static/"))

	hh := http.HandlerFunc(handlers.HomeHandler)
	aah := http.HandlerFunc(handlers.AddAutoHandler)
	ach := http.HandlerFunc(handlers.AddCustomerHandler)
	vah := http.HandlerFunc(handlers.ViewAutosHandler)
	gph := http.HandlerFunc(handlers.GeneratePdfHandler)
	dah := http.HandlerFunc(handlers.DeleteAutoHandler)

	mux.Handle("/", hh)
	//allows us to reference as src="/static/"
	mux.Handle("/static/", http.StripPrefix("/static", fileserver))
	mux.Handle("/add-auto", aah)
	mux.Handle("/add-customer", ach)
	mux.Handle("/view-autos", vah)
	mux.Handle("/generate-pdf", gph)
	mux.Handle("/delete-auto", dah)

	http.ListenAndServe(":8080", mux)

}
