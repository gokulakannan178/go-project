package models

type PdfConfiguration struct {
	Orentation string
}

type PDFData struct {
	Data    map[string]interface{}
	RefData map[string]interface{}
	Config  ProductConfig
}

func (p *PDFData) Inc(i int) int {
	return i + 1
}
