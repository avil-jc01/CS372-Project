package utils

import (
	"database/sql"
	"CS372-Project/models"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func InsertCustomer(c models.Customer, d *sql.DB) error {
	statement, err := d.Prepare("INSERT INTO customers (firstname, lastname, address, city, state, zip) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}

	_, err = statement.Exec(c.FirstName, c.LastName, c.Address, c.City, c.State, c.Zip)

	if err != nil {
		panic(err)
	}
	log.Printf("D: Added %s %s to customers table", c.FirstName, c.LastName)
	return nil

}

func CreateTable(name string, definition string, d *sql.DB) {
	//create table - preparing statement for execution
	statement, err := d.Prepare("CREATE TABLE IF NOT EXISTS " + name + " " + definition)
	if err != nil {
		panic(err)
	}
	//run the statement prepared above
	_, err = statement.Exec()
	if err != nil {
		panic(err)
	}
	log.Printf("D: Successfully created (if not exists) table %s\n", name)
}
