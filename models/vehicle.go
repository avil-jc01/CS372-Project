package models

type Vehicle struct {
	VIN           string
	year          int
	make          string
	model         string
	purchasePrice int
	dateOfSale    date
	customer      customer
}
