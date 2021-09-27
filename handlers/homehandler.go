package handlers

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"html/template"
	"log"
	"budget_app/utils"
	"budget_app/models"
	"net/http"
)



func HomeHandler(w http.ResponseWriter, r *http.Request) {

	database, err := sql.Open("sqlite3", "./budget-app.db")
	if err != nil {
		panic(err)
	}

	//building the html template
	headPath := "templates/head.gohtml"
	navTemplatePath := "templates/navbar.gohtml"
	templatePath := "templates/home.gohtml"

	t, err := template.ParseFiles(templatePath, navTemplatePath, headPath)

	if err != nil {
		log.Printf("E: Error processing template")
	}

	payload := new(models.TemplatePayload)
	//validate session - if authenticated user, display logged-in homepage
	t.Execute(w, &payload)
	
	database.Close()
}
