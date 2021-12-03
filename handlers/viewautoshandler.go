package handlers

import (
	"CS372-Project/utils"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"html/template"
	"log"
	"net/http"
)

func ViewAutosHandler(w http.ResponseWriter, r *http.Request) {

	database, err := sql.Open("sqlite3", "./cs372-app.db")

	if err != nil {
		panic(err)
	}

	defer database.Close()

	if r.Method == "GET" {
		log.Printf("D: HTTP GET on /view-autos")

		htmlHeadPath := "/templates/head.gohtml"
		navBarPath := "/templates/navbar.gohtml"
		templatePath := "/templates/view-autos.gohtml"

		t, err := template.ParseFiles(templatePath, htmlHeadPath, navBarPath)
		if err != nil {
			log.Printf("E: Error processing template")
		}

		dbVehicles := utils.SelectAllAutos(database)

		t.Execute(w, &dbVehicles)
	}
}
