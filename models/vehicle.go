package models

type Vehicle struct {
	VIN           string
	Year          int
	Make          string
	Model         string
	PurchasePrice float32
	DateOfSale    string
	Customer      Customer
}
