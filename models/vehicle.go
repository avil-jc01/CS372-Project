package models

type Vehicle struct {
	VehicleId     int
	VIN           string
	Year          int
	Make          string
	Model         string
	PurchasePrice int
	DateOfSale    string
	CustomerId    int
}
