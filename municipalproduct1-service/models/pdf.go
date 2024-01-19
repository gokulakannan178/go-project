package models

type PdfConfiguration struct {
	Orentation string
}

type PDFData struct {
	Data       map[string]interface{}
	RefData    map[string]interface{}
	RefDataStr map[string]string
	Config     ProductConfiguration
}

func (p *PDFData) Inc(i int) int {
	return i + 1
}

type PDFDataV2 struct {
	ArrData    []PDFDataV2Arr
	RefDataStr map[string]string
	Config     ProductConfiguration
}

type PDFDataV2Arr struct {
	Data    map[string]interface{}
	RefData map[string]interface{}
	Config  ProductConfiguration
}
