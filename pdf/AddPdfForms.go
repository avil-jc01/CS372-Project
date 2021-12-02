package pdf

import (
	"log"
	"github.com/signintech/gopdf"
)

func CreateBosPdf(v models.Vehicle, c models.Customer) (string, error) {

	var pdfPath string

	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})  
	pdf.AddPage()

	err := pdf.AddTTFFont("ubuntu", "./example-font.ttf")

	if err != nil {
		log.Print(err.Error())
		return (nil, err)
	}

	pdf.Image("./44273.jpg", 0, 0, &gopdf.Rect{W: 595, H: 842})

	err = pdf.SetFont("loma", "", 14)
	if err != nil {
		log.Print(err.Error())
		return
	}
	pdf.SetX(250) //move current location
	pdf.SetY(200)
	pdf.Cell(nil, v.VIN) //print text

	pdf.WritePdf("44273_test1.pdf")

}
