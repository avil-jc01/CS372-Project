package handlers

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"html/template"
	"log"
	"net/http"
)


func AddAutoHandler(w http.ResponseWriter, r *http.Request) {

	database, err := sql.Open("sqlite3", "./cs372-app.db")

	if err != nil {
		panic(err)
	}

	defer database.Close()

	templatePath := "templates/addauto.gohtml"

	t, err := template.ParseFiles(templatePath)

	if err != nil {
		log.Printf("E: Error processing template")
	}

	t.Execute(w, nil)
}
