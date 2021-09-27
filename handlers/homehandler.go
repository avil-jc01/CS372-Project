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

	//building the html template
	templatePath := "templates/home.gohtml"

	t, err := template.ParseFiles(templatePath)

	if err != nil {
		log.Printf("E: Error processing template")
	}

	//validate session - if authenticated user, display logged-in homepage
	t.Execute(w, nil)
	
}
