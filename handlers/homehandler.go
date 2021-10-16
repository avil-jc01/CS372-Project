package handlers

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"html/template"
	"log"
	"net/http"
)


func HomeHandler(w http.ResponseWriter, r *http.Request) {

	database, err := sql.Open("sqlite3", "./budget-app.db")
	if err != nil {
		panic(err)
	}

	defer database.Close()

	//building the html template - provide html template path
	//we add html head and navbar templates here
	htmlHeadPath := "/templates/head.gohtml"
	navBarPath := "/templates/navbar.gohtml"
	templatePath := "/templates/home.gohtml"

	// add all the templates to the parsefiles arguments
	t, err := template.ParseFiles(templatePath,htmlHeadPath,navBarPath)
	if err != nil {
		log.Printf("E: Error processing template")
	}

	//show the templated web page to the user
	// we can pass golang structs/objects to the template
	// in this case we pass nil since we don't have one to pass
	t.Execute(w, nil)
	
}
