package main

import (
	"fmt"
	"log"
	"strings"

	pdf "github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

func main() {
	fmt.Println("Working")
	//Client
	pdfg := pdf.NewPDFPreparer()
	html := "<html>Hi</html>"
	pdfg.AddPage(pdf.NewPageReader(strings.NewReader(html)))
	// Create PDF document in internal buffer
	err := pdfg.Create()
	if err != nil {
		log.Fatal(err)
	}

	// Write buffer contents to file on disk
	err = pdfg.WriteFile("simplesample.pdf")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Done")
}
