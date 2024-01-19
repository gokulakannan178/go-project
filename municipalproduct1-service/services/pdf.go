package services

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

//pdf requestpdf struct
type RequestPdf struct {
	body       string
	Orentation string
	Margin     struct {
		IsSet  bool
		Top    uint
		Bottom uint
		Left   uint
		Right  uint
	}
}

//new request to pdf function
func NewRequestPdf(body string) *RequestPdf {
	return &RequestPdf{
		body: body,
	}
}

func NewRequestPdfV2(body, Orentiation string) *RequestPdf {
	return &RequestPdf{
		Orentation: Orentiation,
		body:       body,
	}
}
func (r *RequestPdf) SetMargin(top, right, bottom, left *uint) {

	r.Margin.IsSet = true
	if top != nil {
		r.Margin.Top = *top
	}
	if right != nil {
		r.Margin.Right = *right
	}

	if bottom != nil {
		r.Margin.Bottom = *bottom
	}
	if left != nil {
		r.Margin.Left = *left
	}
}

//parsing template function
func (r *RequestPdf) ParseTemplate(templateFileName string, data interface{}) error {
	fmt.Println("parse template started")

	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	fmt.Println("execute template started")

	if err = t.Execute(buf, data); err != nil {
		fmt.Println("Template Execute error - " + err.Error())
		return err
	}
	fmt.Println("execute finished template started")

	r.body = buf.String()
	fmt.Println("read body finished")
	return nil
}

//parsing template function
func (r *RequestPdf) ParseTemplatev2(templateFileName string, fm template.FuncMap, data interface{}) error {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return err
	}
	t = t.Funcs(fm)
	// t := template.Must(template.New("samplepdf.html").Funcs(fm).Parse(templateFileName))
	// t := template.New("demo").Funcs(fm)

	buf := new(bytes.Buffer)
	if err := t.Execute(buf, data); err != nil {
		return err
	}
	r.body = buf.String()
	return nil
}

//generate pdf function
func (r *RequestPdf) GeneratePDF(pdfPath string) (bool, error) {
	t := time.Now().Unix()

	cloneTemplateURL := "cloneTemplate/" + strconv.FormatInt(int64(t), 10) + ".html"
	// write whole the body
	err1 := ioutil.WriteFile(cloneTemplateURL, []byte(r.body), 0644)
	if err1 != nil {
		panic(err1)
	}

	f, err := os.Open("cloneTemplate/" + strconv.FormatInt(int64(t), 10) + ".html")

	if err != nil {
		log.Fatal(err)
	}

	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Fatal(err)
	}

	pdfg.AddPage(wkhtmltopdf.NewPageReader(f))

	pdfg.PageSize.Set(wkhtmltopdf.PageSizeA4)

	pdfg.Dpi.Set(300)

	err = pdfg.Create()
	if err != nil {
		log.Fatal(err)
	}

	err = pdfg.WriteFile(pdfPath)
	if err != nil {
		log.Fatal(err)
	}

	if f != nil {
		f.Close()
	}
	err = os.Remove(cloneTemplateURL)
	fmt.Println(err)
	return true, nil
}

//generate pdf function
func (r *RequestPdf) GeneratePDFAsFile() (bool, []byte, error) {
	t := time.Now().Unix()

	cloneTemplateURL := "cloneTemplate/" + strconv.FormatInt(int64(t), 10) + ".html"

	// write whole the body
	err1 := ioutil.WriteFile(cloneTemplateURL, []byte(r.body), 0644)
	if err1 != nil {
		fmt.Println("error in write file - " + err1.Error())
		panic(err1)
	}

	f, err := os.Open("cloneTemplate/" + strconv.FormatInt(int64(t), 10) + ".html")

	if err != nil {
		log.Fatal(err)
	}
	// defer f.Close()
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Fatal(err)
	}
	// po := wkhtmltopdf.NewPageOptions()
	// po.DefaultHeader = true
	pageReader := wkhtmltopdf.NewPageReader(f)

	// pageReader.PageOptions.CustomHeader=
	pdfg.AddPage(pageReader)
	pdfg.PageSize.Set(wkhtmltopdf.PageSizeA4)
	if r.Margin.IsSet {
		pdfg.MarginBottom.Set(r.Margin.Bottom)
		pdfg.MarginLeft.Set(r.Margin.Left)
		pdfg.MarginRight.Set(r.Margin.Right)
		pdfg.MarginTop.Set(r.Margin.Top)
	}

	if r.Orentation == "" {
		pdfg.Orientation.Set(wkhtmltopdf.OrientationPortrait)

	} else {
		pdfg.Orientation.Set(r.Orentation)

	}

	pdfg.Dpi.Set(300)

	err = pdfg.Create()
	if err != nil {
		log.Fatal(err)
	}
	if f != nil {
		f.Close()
	}
	err = os.Remove(cloneTemplateURL)
	fmt.Println(err)
	return true, pdfg.Bytes(), nil
}
