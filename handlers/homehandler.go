package handlers

import (
	_ "github.com/mattn/go-sqlite3"
	"html/template"
	"log"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	htmlHeadPath := "/templates/head.gohtml"
	navBarPath := "/templates/navbar.gohtml"
	templatePath := "/templates/home.gohtml"

	t, err := template.ParseFiles(templatePath, htmlHeadPath, navBarPath)
	if err != nil {
		log.Printf("E: Error processing template")
	}

	t.Execute(w, nil)

}
