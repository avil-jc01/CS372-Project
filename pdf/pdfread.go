package main

import (
	"log"
	"github.com/signintech/gopdf"
	"io"
        "net/http"
        "os"
)

func main(){

	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})  
	pdf.AddPage()

	var err error

	// Download a Font
        fontUrl := "https://github.com/google/fonts/blob/main/ufl/ubuntu/Ubuntu-Regular.ttf?raw=true"
        if err = DownloadFile("example-font.ttf", fontUrl); err != nil {
            panic(err)
        }

	
	err = pdf.AddTTFFont("loma", "./example-font.ttf")
	if err != nil {
		log.Print(err.Error())
		return
	}


	pdf.Image("./44273.jpg", 0, 0, &gopdf.Rect{W: 595, H: 842}) //print image
	err = pdf.SetFont("loma", "", 14)
	if err != nil {
		log.Print(err.Error())
		return
	}
	pdf.SetX(250) //move current location
	pdf.SetY(200)
	pdf.Cell(nil, "This is a test text box over the form") //print text

	pdf.WritePdf("image.pdf")
}

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func DownloadFile(filepath string, url string) error {
        // Get the data
        resp, err := http.Get(url)
        if err != nil {
            return err
        }
        defer resp.Body.Close()

        // Create the file
        out, err := os.Create(filepath)
        if err != nil {
            return err
        }
        defer out.Close()

        // Write the body to file
        _, err = io.Copy(out, resp.Body)
        return err
}
