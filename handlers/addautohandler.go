package handlers

import (
	"CS372-Project/models"
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func AddAutoHandler(w http.ResponseWriter, r *http.Request) {

	database, err := sql.Open("sqlite3", "./cs372-app.db")

	if err != nil {
		panic(err)
	}

	defer database.Close()

	htmlHeadPath := "/templates/head.gohtml"
	navBarPath := "/templates/navbar.gohtml"
	templatePath := "/templates/addauto.gohtml"

	t, err := template.ParseFiles(templatePath, htmlHeadPath, navBarPath)
	if err != nil {
		panic(err)
	}

	if r.Method != http.MethodPost {
		t.Execute(w, nil)
		return
	}
	Year, err := strconv.ParseInt(r.FormValue("year"), 10, 32)
	if err != nil {
		panic(err)
	}
	PurchasePrice, err := strconv.ParseFloat(r.FormValue("price"), 32)
	if err != nil {
		panic(err)
	}

	const (
		layoutISO = "2006-01-02"
	)
	date := r.FormValue("date")
	DateOfSale, err := time.Parse(layoutISO, date)
	fmt.Println(DateOfSale)
	newVehicle := models.Vehicle{
		VIN:           r.FormValue("vin"),
		Year:          int(Year),
		Make:          r.FormValue("make"),
		Model:         r.FormValue("model"),
		PurchasePrice: float32(PurchasePrice),
		DateOfSale:    DateOfSale.String(),
	}
	fmt.Println(newVehicle)

	if err != nil {
		log.Printf("E: Error processing template")
	}

	t.Execute(w, nil)
}
