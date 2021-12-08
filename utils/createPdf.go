package utils

import (
	"CS372-Project/models"
	"github.com/signintech/gopdf"
	"log"
	"strconv"
)

func CreatePdf(v models.Vehicle, c models.Customer) (string, error) {

	var pdfPath string

	pdf := gopdf.GoPdf{}

	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	pdf.AddPage()

	err := pdf.AddTTFFont("ubuntu", "/pdf/example-font.ttf")

	if err != nil {
		log.Print(err.Error())
		return "", err
	}

	pdf.Image("/pdf/44273.jpg", 0, 0, &gopdf.Rect{W: 595, H: 842})

	err = pdf.SetFont("ubuntu", "", 14)
	if err != nil {
		log.Print(err.Error())
		return "", err
	}
  
	pdf.SetX(48)
	pdf.SetY(155)
	spacedVin := SpacedVin(v.VIN)
	pdf.Cell(nil, spacedVin)

	pdf.SetX(65)
	pdf.SetY(198)
	yearStr := strconv.Itoa(v.Year)
	pdf.Cell(nil, yearStr)

	pdf.SetX(250)
	pdf.SetY(198)
	pdf.Cell(nil, v.Make)

	pdf.SetX(415)
	pdf.SetY(198)
	pdf.Cell(nil, v.Model)

	pdf.SetX(105)
	pdf.SetY(330)
	priceStr := strconv.Itoa(v.PurchasePrice)
	pdf.Cell(nil, priceStr)

	pdf.SetX(400)
	pdf.SetY(330)
	pdf.Cell(nil, v.DateOfSale)

	pdf.SetX(50)
	pdf.SetY(678)
	buyerName := c.LastName + ", " + c.FirstName
	pdf.Cell(nil, buyerName)

	pdf.AddPage()

	pdf.Image("/pdf/cfr.jpg", 0, 0, &gopdf.Rect{W: 595, H: 842})

	pdf.SetX(80)
	pdf.SetY(145)
	pdf.Cell(nil, v.Make)

	pdf.SetX(180)
	pdf.SetY(145)
	pdf.Cell(nil, v.Model)

	pdf.SetX(290)
	pdf.SetY(145)
	pdf.Cell(nil, yearStr)

	pdf.SetX(375)
	pdf.SetY(145)
	pdf.Cell(nil, v.VIN)

	pdfPath = "/pdfout/" + c.LastName + "_" + v.VIN + ".pdf"

	pdf.WritePdf(pdfPath)

	return pdfPath, nil

}
