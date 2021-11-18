package handlers

import (
	"CS372-Project/models"
	"CS372-Project/utils"
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
	if r.Method == "GET" {
		log.Printf("D: HTTP GET on /add-auto")
		htmlHeadPath := "/templates/head.gohtml"
		navBarPath := "/templates/navbar.gohtml"
		templatePath := "/templates/addauto.gohtml"

		t, err := template.ParseFiles(templatePath, htmlHeadPath, navBarPath)
		if err != nil {
			log.Printf("E: Error processing template")
		}
		t.Execute(w, nil)
	}
	if r.Method == "POST" {
		log.Printf("D: HTTP POST on /add-customer")
		r.ParseForm()
		var err error
		var DateOfSale time.Time

		_, has_param := r.Form["date"]
		if has_param {
			const (
				layoutISO = "2006-01-02"
			)
			date := r.FormValue("date")
			DateOfSale, err = time.Parse(layoutISO, date)

			if err != nil {
				panic(err)
			}

		}
		Year, err := strconv.ParseInt(r.FormValue("year"), 10, 32)
		if err != nil {
			panic(err)
		}
		PurchasePrice, err := strconv.ParseFloat(r.FormValue("price"), 32)
		if err != nil {
			panic(err)
		}

		fmt.Println(DateOfSale)
		newVehicle := models.Vehicle{
			VIN:           r.FormValue("vin"),
			Year:          int(Year),
			Make:          r.FormValue("make"),
			Model:         r.FormValue("model"),
			PurchasePrice: int(PurchasePrice),
			DateOfSale:    DateOfSale.String(),
		}
		err = utils.InsterAuto(newVehicle, database)
		if err != nil {
			log.Printf("E: Error inserting customer", newVehicle)
			log.Printf(err.Error())
		}

		http.Redirect(w, r, "/", http.StatusFound)
	}
}
