package handlers

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func AddCustomerHandler(w http.ResponseWriter, r *http.Request) {
	database, err := sql.Open("sqlite3", "./budget-app.db")
	if err != nil {
		panic(err)
	}

	defer database.Close()
	htmlHeadPath := "/templates/head.gohtml"
	navBarPath := "/templates/navbar.gohtml"
	customerPath := "/templates/addcustomer.gohtml"

	t, err := template.ParseFiles(customerPath, htmlHeadPath, navBarPath)
	if err != nil {
		log.Printf("E: Error processing template")
	}
	t.Execute(w, nil)
}
