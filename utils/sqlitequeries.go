package utils

import (
	"CS372-Project/models"
	"database/sql"
	"log"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

func SelectCustomerById(cid int, d *sql.DB) models.Customer {
	var returnCustomer models.Customer

	cidString := strconv.Itoa(cid)
	resultRow, _ := d.Query("SELECT * FROM customers WHERE id = " + cidString)

	var zip int
	var fn, ln, addr, city, st string
	for resultRow.Next() {
		resultRow.Scan(&cid, &fn, &ln, &addr, &city, &st, &zip)

		returnCustomer = models.Customer{
			CustomerId: cid,
			FirstName:  fn,
			LastName:   ln,
			Address:    addr,
			City:       city,
			State:      st,
			Zip:        zip,
		}
	}

	return returnCustomer
}

func SelectVehicleById(vid int, d *sql.DB) models.Vehicle {
	var returnVehicle models.Vehicle

	vidString := strconv.Itoa(vid)
	resultRow, _ := d.Query("SELECT * FROM vehicles WHERE id = " + vidString)

	var year, price, cid int
	var vin, make, model, date string
	for resultRow.Next() {
		resultRow.Scan(&vid, &vin, &year, &make, &model, &price, &date, &cid)

		returnVehicle = models.Vehicle{
			VehicleId:     vid,
			VIN:           vin,
			Year:          year,
			Make:          make,
			Model:         model,
			PurchasePrice: price,
			DateOfSale:    date,
			CustomerId:    cid,
		}
	}

	return returnVehicle
}

func SelectAllCustomers(d *sql.DB) []models.Customer {
	var dbCustomers []models.Customer

	resultRows, err := d.Query("SELECT * FROM customers")
	if err != nil {
		panic(err)
	}

	var cid, zip int
	var fn, ln, addr, city, st string

	for resultRows.Next() {
		resultRows.Scan(&cid, &fn, &ln, &addr, &city, &st, &zip)

		resultCustomer := models.Customer{
			CustomerId: cid,
			FirstName:  fn,
			LastName:   ln,
			Address:    addr,
			City:       city,
			State:      st,
			Zip:        zip,
		}

		dbCustomers = append(dbCustomers, resultCustomer)

	}

	return dbCustomers

}

func SelectAllAutos(d *sql.DB) []models.Vehicle {
	var dbVehicles []models.Vehicle

	resultRows, err := d.Query("SELECT * FROM vehicles")
	if err != nil {
		panic(err)
	}

	var vid, yr, price, cid int
	var vin, make, model, date string

	for resultRows.Next() {
		resultRows.Scan(&vid, &vin, &yr, &make, &model, &price, &date, &cid)
		readVehicle := models.Vehicle{
			VehicleId:     vid,
			VIN:           vin,
			Year:          yr,
			Make:          make,
			Model:         model,
			PurchasePrice: price,
			DateOfSale:    date,
			CustomerId:    cid,
		}

		dbVehicles = append(dbVehicles, readVehicle)
	}

	return dbVehicles

}

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

func InsertAuto(v models.Vehicle, d *sql.DB) error {
	statement, err := d.Prepare("INSERT INTO vehicles ( vin, year, make, model, purchase_price, date_of_sale, customer_id) VALUES ( ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	result, err := statement.Exec(v.VIN, v.Year, v.Make, v.Model, v.PurchasePrice, v.DateOfSale, v.CustomerId)
	if err != nil {
		panic(err)
	}
	res, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	log.Printf("D: Added %d to vehicles table", res)
	return nil

}

func CreateTable(name string, definition string, d *sql.DB) {
	statement, err := d.Prepare("CREATE TABLE IF NOT EXISTS " + name + " " + definition)
	if err != nil {
		panic(err)
	}

	_, err = statement.Exec()
	if err != nil {
		panic(err)
	}
	log.Printf("D: Successfully created (if not exists) table %s\n", name)
}

func DeleteAuto(VehicleId, CustomerId int, d *sql.DB) error {
	statement, err := d.Prepare("DELETE FROM vehicles WHERE id = " + strconv.Itoa(VehicleId) + " AND customer_id = " + strconv.Itoa(CustomerId))
	if err != nil {
		return err
	}
	_, err = statement.Exec()
	if err != nil {
		return err
	}
	log.Printf("D: Successfully deleted vehicle #%d\n", VehicleId)
	return nil
}
