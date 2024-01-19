package services

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"text/template"

	"github.com/jung-kurt/gofpdf"
)

func generatePDFV2(data []byte) (*gofpdf.Fpdf, error) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)
	pdf.Cell(0, 10, string(data))

	return pdf, nil
}

func savePDFToFile(pdf *gofpdf.Fpdf, filePath string) error {
	err := pdf.OutputFileAndClose(filePath)
	if err != nil {
		return err
	}

	fmt.Println("PDF saved to:", filePath)
	return nil
}
func (s *Service) UploadPayslip(data []byte, Path string, filename string) error {
	pdf, err := generatePDFV2(data)
	if err != nil {
		fmt.Println("Error generating PDF:", err)
		return err
	}
	filePath := Path + filename
	fmt.Println("filePath====>", filePath)
	// Save the PDF to the file
	err = savePDFToFile(pdf, filePath)
	if err != nil {
		fmt.Println("Error saving PDF:", err)
		return err
	}
	return nil
}

type PageData2 struct {
	Title   string
	Content string
}

func (s *Service) generatePDFFromTemplateV2(data interface{}, outputFilePath string, templatePath string) (string, error) {
	tmpl, err := template.ParseFiles(templatePath)
	var htmlStr strings.Builder
	err = tmpl.Execute(&htmlStr, data)
	if err != nil {
		return "", err
	}

	// Write the HTML content to a temporary file
	tempFile, err := os.CreateTemp("", "html-template-*.html")
	if err != nil {
		return "", err
	}
	defer os.Remove(tempFile.Name())
	defer tempFile.Close()
	tempFile.WriteString(htmlStr.String())

	// Execute the wkhtmltopdf command to convert HTML to PDF
	cmd := exec.Command("wkhtmltopdf", tempFile.Name(), outputFilePath)
	err = cmd.Run()
	if err != nil {
		return "", err
	}

	fmt.Println("PDF saved to:", outputFilePath)
	return outputFilePath, nil
}
