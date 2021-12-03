package handlers

import (
	"CS372-Project/models"
	"CS372-Project/utils"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func AddCustomerHandler(w http.ResponseWriter, r *http.Request) {

	database, err := sql.Open("sqlite3", "./cs372-app.db")
	if err != nil {
		panic(err)
	}

	defer database.Close()

	if r.Method == "GET" {
		log.Printf("D: HTTP GET on /add-customer")

		htmlHeadPath := "/templates/head.gohtml"
		navBarPath := "/templates/navbar.gohtml"
		customerPath := "/templates/addcustomer.gohtml"

		t, err := template.ParseFiles(customerPath, htmlHeadPath, navBarPath)
		if err != nil {
			log.Printf("E: Error processing template")
		}
		t.Execute(w, nil)
	}
	if r.Method == "POST" {
		log.Printf("D: HTTP POST on /add-customer")

		r.ParseForm()

		var firstName, lastName, address, city, state string
		var zip int

		_, has_param := r.Form["first_name"]
		if has_param == true {
			firstName = r.Form["first_name"][0]
		}

		_, has_param = r.Form["last_name"]
		if has_param == true {
			lastName = r.Form["last_name"][0]
		}

		_, has_param = r.Form["address"]
		if has_param == true {
			address = r.Form["address"][0]
		}

		_, has_param = r.Form["city"]
		if has_param == true {
			city = r.Form["city"][0]
		}

		_, has_param = r.Form["state"]
		if has_param == true {
			state = r.Form["state"][0]
		}

		_, has_param = r.Form["zip"]
		if has_param == true {
			zip_str := r.Form["zip"][0]
			zip, _ = strconv.Atoi(zip_str)
		}

		newCustomer := models.Customer{
			FirstName: firstName,
			LastName:  lastName,
			Address:   address,
			City:      city,
			State:     state,
			Zip:       zip,
		}

		err = utils.InsertCustomer(newCustomer, database)

		if err != nil {
			log.Printf("E: Error inserting customer", newCustomer)
			log.Printf(err.Error())
		}

		http.Redirect(w, r, "/", http.StatusFound)

	}
}
