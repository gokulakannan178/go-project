package services

import (
	"context"
	"errors"
	"fmt"
	"math"
	"municipalproduct1-service/app"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
)

// /PropertyOverallDemandReportPDF : ""
func (s *Service) PropertyOverallDemandReportPDF(ctx *models.Context, filter *models.PropertyFilter, p *models.Pagination) ([]byte, error) {
	properties, err := s.PropertyOverallDemandReport(ctx, filter, p)
	if err != nil {
		return nil, err
	}
	m := make(map[string]interface{})
	m2 := make(map[string]interface{})
	m["demand"] = properties
	m2["currentDate"] = time.Now()
	cfy, err := s.Daos.GetCurrentFinancialYear(ctx)
	if err != nil {
		return nil, errors.New("Error in getting current financial year " + err.Error())
	}
	m2["currentFy"] = cfy
	var pdfdata models.PDFData
	pdfdata.Data = m
	pdfdata.RefData = m2
	productConfigUniqueID := "1"
	productConfig, err := s.Daos.GetSingleProductConfiguration(ctx, productConfigUniqueID)
	if err != nil {
		return nil, errors.New("Error in getting product config" + err.Error())
	}
	pdfdata.Config = productConfig.ProductConfiguration

	r := NewRequestPdf("")
	templatePathStart := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.TEMPLATELOC)
	//html template path
	templatePath := templatePathStart + "property_demand_report.html.html"
	if err := r.ParseTemplate(templatePath, pdfdata); err == nil {
		fmt.Println("start pdf generated successfully")

		ok, data, err := r.GeneratePDFAsFile()

		fmt.Println(ok, "pdf generated successfully")
		return data, err
	} else {
		fmt.Println("Error in parcing template - " + err.Error())

		return nil, errors.New("Error in parcing template - " + err.Error())
	}
	return nil, nil

}

// PropertyOverallDemandReportExcel : ""
func (s *Service) PropertyOverallDemandReportExcel(ctx *models.Context, filter *models.PropertyFilter, pagination *models.Pagination) (*excelize.File, error) {
	res, err := s.PropertyOverallDemandReport(ctx, filter, pagination)
	if err != nil {
		return nil, err
	}

	resPD, err1 := s.Daos.GetSingleDefaultProductConfiguration(ctx)
	if err1 != nil {
		return nil, err1
	}

	cfy, err := s.Daos.GetCurrentFinancialYear(ctx)
	if err != nil {
		return nil, errors.New("Error in getting current financial year " + err.Error())
	}
	fmt.Println("'res length==>'", len(res))

	//  create an excel file
	excel := excelize.NewFile()
	sheet1 := "PropertyOverallDemand"
	rowNo := 1
	index := excel.NewSheet(sheet1)
	excel.SetActiveSheet(index)
	if resPD.LocationID == "Bhagalpur" {
		excel.MergeCell(sheet1, "A1", "F3")
		excel.MergeCell(sheet1, "A4", "F5")
	} else {
		excel.MergeCell(sheet1, "A1", "B5")
		excel.MergeCell(sheet1, "C1", "F3")
		excel.MergeCell(sheet1, "C4", "F5")
	}
	excel.MergeCell(sheet1, "A6", "F6")

	style1, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#FFDC6D"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	style3, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#B6DDE8"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}

	if resPD.LocationID != "Bhagalpur" {
		documentUrl := s.Shared.Config.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.FILEURL)
		if err := excel.AddPicture(sheet1, fmt.Sprintf("%v%v", "A", rowNo), documentUrl+"municipal/logo.png", `{"x_scale": 0.115, "y_scale": 0.0935}`); err != nil {
			fmt.Println(err)
		}
	}

	if resPD.LocationID == "Bhagalpur" {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), ctx.ProductConfig.Name)
	} else {
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), s.Shared.Config.GetString(s.Shared.GetCmdArg(constants.ENV)+"."+constants.LOCATIONNAME))
	}
	rowNo++
	rowNo++
	rowNo++

	if resPD.LocationID == "Bhagalpur" {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "Property Overall Demand")
	} else {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Property Overall Demand")
	}
	rowNo++
	rowNo++
	t := time.Now()
	toDate := t.Format("02-January-2006")
	reportFromMsg := "Report Generated on" + " " + toDate
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style3)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), reportFromMsg)
	rowNo++

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "F", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "S.No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "Holding No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Ward")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), "Year")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), "Amount")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), "Owner Name")

	fmt.Println("'res length==>'", len(res))
	for i, v := range res {
		rowNo++
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), i+1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v.UniqueID)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), func() string {
			if v.Ref.Address.Ward != nil {
				return v.Ref.Address.Ward.Name
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), func() string {
			if v.DOA != nil && cfy != nil && cfy.To != nil {
				return fmt.Sprintf("%v - %v", v.DOA.Year(), cfy.To.Year())
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), v.Ref.Demand.Total.TotalTax)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), func() string {
			if len(v.Ref.PropertyOwner) > 0 {
				return v.Ref.PropertyOwner[0].Name
			}
			return "NA"
		}())

	}
	var totalTax float64
	for _, v1 := range res {
		totalTax = totalTax + v1.Ref.Demand.Total.TotalTax
	}

	rowNo++
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "Total")

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "B", rowNo), fmt.Sprintf("%v%v", "B", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "D", rowNo), fmt.Sprintf("%v%v", "D", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "E", rowNo), fmt.Sprintf("%v%v", "E", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), totalTax)

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "E", rowNo), fmt.Sprintf("%v%v", "F", rowNo), style1)

	return excel, nil
}

// /PropertyOverallDemandReportPDF : ""
func (s *Service) PropertyOverallDemandReportJSON(ctx *models.Context, filter *models.PropertyFilter, p *models.Pagination) (*models.PropertyOverallDemandReport, error) {
	properties, err := s.PropertyOverallDemandReport(ctx, filter, p)
	if err != nil {
		return nil, err
	}

	report := new(models.PropertyOverallDemandReport)
	cfy, err := s.Daos.GetCurrentFinancialYear(ctx)
	if err != nil {
		return nil, errors.New("Error in getting current financial year " + err.Error())
	}
	report.Properties = properties
	if cfy != nil {
		report.CFY = *cfy
	}
	return report, nil
}

// DayWisePropertyDemandReport : ""
func (s *Service) DayWisePropertyDemandReport(ctx *models.Context, filter *models.DayWisePropertyDemandChartFilter) (*models.DayWisePropertyDemandChart, error) {

	res, err := s.Daos.DayWisePropertyDemandReport(ctx, filter)
	if err != nil {
		return nil, err
	}
	if res != nil {
		for _, v := range res.Records {
			res.Total = res.Total + v.Amount
		}

	}
	return res, nil
}

// DayWisePropertyDemandReportExcel: ""
func (s *Service) DayWisePropertyDemandReportExcel(ctx *models.Context, filter *models.DayWisePropertyDemandChartFilter) (*excelize.File, error) {
	data, err := s.DayWisePropertyDemandReport(ctx, filter)
	if err != nil {
		return nil, err
	}
	if data == nil {
		data = new(models.DayWisePropertyDemandChart)
	}

	resPD, err1 := s.Daos.GetSingleDefaultProductConfiguration(ctx)
	if err1 != nil {
		return nil, err1
	}

	excel := excelize.NewFile()
	sheet1 := "Property Day Wise Demand"
	index := excel.NewSheet(sheet1)
	rowNo := 1

	excel.SetActiveSheet(index)
	if resPD.LocationID == "Bhagalpur" {
		excel.MergeCell(sheet1, "A1", "C3")
		excel.MergeCell(sheet1, "A4", "C5")
	} else {
		excel.MergeCell(sheet1, "A1", "B5")
		excel.MergeCell(sheet1, "C1", "C3")
		excel.MergeCell(sheet1, "C4", "C5")
	}
	excel.MergeCell(sheet1, "A6", "C6")

	style1, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#FFDC6D"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	style3, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#B6DDE8"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}

	if resPD.LocationID != "Bhagalpur" {
		documentUrl := s.Shared.Config.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.FILEURL)
		if err := excel.AddPicture(sheet1, fmt.Sprintf("%v%v", "A", rowNo), documentUrl+"municipal/logo.png", `{"x_scale": 0.115, "y_scale": 0.0935}`); err != nil {
			fmt.Println(err)
		}
	}

	if resPD.LocationID == "Bhagalpur" {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), ctx.ProductConfig.Name)
	} else {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), ctx.ProductConfig.Name)
	}
	rowNo++
	rowNo++
	rowNo++

	if resPD.LocationID == "Bhagalpur" {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), sheet1)
	} else {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), sheet1)
	}

	rowNo++
	rowNo++

	t := time.Now()
	toDate := t.Format("02-January-2006")
	reportFromMsg := "Report Generated on" + " " + toDate
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style3)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), reportFromMsg)
	rowNo++

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "Date")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "Property Count")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Amount")
	rowNo++
	var totalAmount float64
	if data != nil {
		for _, v := range data.Records {
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), v.ID)
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v.PropertyCount)
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), v.Amount)
			totalAmount = totalAmount + v.Amount
			rowNo++

		}
	}
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "B", rowNo), fmt.Sprintf("%v%v", "B", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "Total")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf(" %.2f", totalAmount))

	return excel, nil

}

// /DayWisePropertyDemandReportPDF : ""
func (s *Service) DayWisePropertyDemandReportPDF(ctx *models.Context, filter *models.DayWisePropertyDemandChartFilter) ([]byte, error) {
	properties, err := s.DayWisePropertyDemandReport(ctx, filter)
	if err != nil {
		return nil, err
	}
	if properties == nil {
		properties = new(models.DayWisePropertyDemandChart)
	}
	m := make(map[string]interface{})
	m2 := make(map[string]interface{})
	m["demand"] = properties
	m2["currentDate"] = time.Now()
	cfy, err := s.Daos.GetCurrentFinancialYear(ctx)
	if err != nil {
		return nil, errors.New("Error in getting current financial year " + err.Error())
	}
	m2["currentFy"] = cfy
	var pdfdata models.PDFData
	pdfdata.Data = m
	pdfdata.RefData = m2
	productConfigUniqueID := "1"
	productConfig, err := s.Daos.GetSingleProductConfiguration(ctx, productConfigUniqueID)
	if err != nil {
		return nil, errors.New("Error in getting product config" + err.Error())
	}
	pdfdata.Config = productConfig.ProductConfiguration

	r := NewRequestPdf("")
	templatePathStart := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.TEMPLATELOC)
	//html template path
	templatePath := templatePathStart + "property_daywise_demand_report.html"
	if err := r.ParseTemplate(templatePath, pdfdata); err == nil {
		fmt.Println("start pdf generated successfully")

		ok, data, err := r.GeneratePDFAsFile()

		fmt.Println(ok, "pdf generated successfully")
		return data, err
	} else {
		fmt.Println("Error in parcing template - " + err.Error())

		return nil, errors.New("Error in parcing template - " + err.Error())
	}
	return nil, nil

}

// DayWisePropertyCollectionReportExcel: ""
func (s *Service) DayWisePropertyCollectionReportExcel(ctx *models.Context, filter *models.DashboardTotalCollectionChartFilter) (*excelize.File, error) {
	data, err := s.DashboardTotalCollectionChart(ctx, filter)
	if err != nil {
		return nil, err
	}
	resPD, err1 := s.Daos.GetSingleDefaultProductConfiguration(ctx)
	if err1 != nil {
		return nil, err1
	}

	excel := excelize.NewFile()
	sheet1 := "Property Day Wise Collection"
	index := excel.NewSheet(sheet1)
	rowNo := 1

	excel.SetActiveSheet(index)
	if resPD.LocationID == "Bhagalpur" {
		excel.MergeCell(sheet1, "A1", "D3")
		excel.MergeCell(sheet1, "A4", "D5")
	} else {
		excel.MergeCell(sheet1, "A1", "B5")
		excel.MergeCell(sheet1, "C1", "D3")
		excel.MergeCell(sheet1, "C4", "D5")
	}

	style1, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#FFDC6D"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	style3, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#B6DDE8"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	if resPD.LocationID != "Bhagalpur" {
		documentUrl := s.Shared.Config.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.FILEURL)
		if err := excel.AddPicture(sheet1, fmt.Sprintf("%v%v", "A", rowNo), documentUrl+"municipal/logo.png", `{"x_scale": 0.115, "y_scale": 0.0935}`); err != nil {
			fmt.Println(err)
		}
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), ctx.ProductConfig.Name)
	} else {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), ctx.ProductConfig.Name)
	}

	rowNo++
	rowNo++
	rowNo++

	if resPD.LocationID == "Bhagalpur" {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), sheet1)
	} else {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), sheet1)
	}
	rowNo++
	rowNo++
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "B", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "Name")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "Amount")
	rowNo++
	var totalAmount float64
	for _, v := range data {
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), v.Name)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v.Payments.Collection)
		totalAmount = totalAmount + v.Payments.Collection
		rowNo++

	}
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "B", rowNo), fmt.Sprintf("%v%v", "B", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "Total")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), fmt.Sprintf(" %.2f", totalAmount))

	return excel, nil

}

// /DayWisePropertyCollectionReportPDF : ""
func (s *Service) DayWisePropertyCollectionReportPDF(ctx *models.Context, filter *models.DashboardTotalCollectionChartFilter) ([]byte, error) {
	properties, err := s.DashboardTotalCollectionChart(ctx, filter)
	if err != nil {
		return nil, err
	}
	m := make(map[string]interface{})
	m2 := make(map[string]interface{})
	var totalCollection float64
	for _, v := range properties {
		totalCollection = totalCollection + v.Payments.Collection
	}
	m["demand"] = properties
	m["totalCollection"] = totalCollection
	m2["currentDate"] = time.Now()
	cfy, err := s.Daos.GetCurrentFinancialYear(ctx)
	if err != nil {
		return nil, errors.New("Error in getting current financial year " + err.Error())
	}
	m2["currentFy"] = cfy
	var pdfdata models.PDFData
	pdfdata.Data = m
	pdfdata.RefData = m2
	productConfigUniqueID := "1"
	productConfig, err := s.Daos.GetSingleProductConfiguration(ctx, productConfigUniqueID)
	if err != nil {
		return nil, errors.New("Error in getting product config" + err.Error())
	}
	pdfdata.Config = productConfig.ProductConfiguration

	r := NewRequestPdf("")
	templatePathStart := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.TEMPLATELOC)
	//html template path
	templatePath := templatePathStart + "property_daywise_collection_report.html"
	if err := r.ParseTemplate(templatePath, pdfdata); err == nil {
		fmt.Println("start pdf generated successfully")

		ok, data, err := r.GeneratePDFAsFile()

		fmt.Println(ok, "pdf generated successfully")
		return data, err
	} else {
		fmt.Println("Error in parcing template - " + err.Error())

		return nil, errors.New("Error in parcing template - " + err.Error())
	}
	return nil, nil

}

// DayWisePropertyCollectionReportExcel: ""
func (s *Service) DashboardDayWiseCollectionChartExcel(ctx *models.Context, filter *models.DashboardDayWiseCollectionChartFilter) (*excelize.File, error) {
	data, err := s.DashboardDayWiseCollectionChart(ctx, filter)
	if err != nil {
		return nil, err
	}

	resPD, err1 := s.Daos.GetSingleDefaultProductConfiguration(ctx)
	if err1 != nil {
		return nil, err1
	}

	excel := excelize.NewFile()
	sheet1 := "Property Day Wise Collection"
	index := excel.NewSheet(sheet1)
	rowNo := 1

	excel.SetActiveSheet(index)
	if resPD.LocationID == "Bhagalpur" {
		excel.MergeCell(sheet1, "A1", "E3")
		excel.MergeCell(sheet1, "A4", "E5")
	} else {
		excel.MergeCell(sheet1, "A1", "B5")
		excel.MergeCell(sheet1, "C1", "E3")
		excel.MergeCell(sheet1, "C4", "E5")
	}

	excel.MergeCell(sheet1, "A6", "E6")
	style1, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#FFDC6D"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	style3, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#B6DDE8"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	if resPD.LocationID != "Bhagalpur" {
		documentUrl := s.Shared.Config.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.FILEURL)
		if err := excel.AddPicture(sheet1, fmt.Sprintf("%v%v", "A", rowNo), documentUrl+"municipal/logo.png", `{"x_scale": 0.115, "y_scale": 0.0935}`); err != nil {
			fmt.Println(err)
		}
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), ctx.ProductConfig.Name)
	} else {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), ctx.ProductConfig.Name)
	}
	rowNo++
	rowNo++
	rowNo++
	if resPD.LocationID == "Bhagalpur" {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), sheet1)
	} else {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), sheet1)
	}
	rowNo++
	rowNo++

	t := time.Now()
	toDate := t.Format("02-January-2006")
	reportFromMsg := "Report Generated on" + " " + toDate
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style3)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), reportFromMsg)
	rowNo++

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "E", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "Name")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "Arrear")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Current")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), "Penalty")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), "Amount")
	rowNo++
	var totalAmount float64

	for _, v := range data.Records {
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), v.ID.DayStr)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v.ArrearCollection)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), v.CurrentCollection)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), v.CurrentPenalty+v.ArrearPenalty)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), v.TotalTax)
		totalAmount = totalAmount + v.TotalTax
		rowNo++

	}
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "B", rowNo), fmt.Sprintf("%v%v", "B", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "D", rowNo), fmt.Sprintf("%v%v", "D", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "E", rowNo), fmt.Sprintf("%v%v", "E", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "Total")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), fmt.Sprintf(" %.2f", totalAmount))

	return excel, nil

}

// propert payment Daywise collection Report
func (s *Service) DayWiseCollectionChartExcel(ctx *models.Context, filter *models.DashboardDayWiseCollectionChartFilter) (*excelize.File, error) {
	data, err := s.DashboardDayWiseCollectionChart(ctx, filter)
	if err != nil {
		return nil, err
	}
	excel := excelize.NewFile()
	sheet1 := "Day Wise Collection Report"
	index := excel.NewSheet(sheet1)
	rowNo := 1

	excel.SetActiveSheet(index)
	excel.MergeCell(sheet1, "A1", "B5")
	excel.MergeCell(sheet1, "C1", "J3")
	excel.MergeCell(sheet1, "C4", "J5")
	excel.MergeCell(sheet1, "A6", "J6")
	style1, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#FFDC6D"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	style3, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#B6DDE8"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	documentUrl := s.Shared.Config.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.FILEURL)
	if err := excel.AddPicture(sheet1, fmt.Sprintf("%v%v", "A", rowNo), documentUrl+"municipal/logo.png", `{"x_scale": 0.115, "y_scale": 0.0935}`); err != nil {
		fmt.Println(err)
	}
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style3)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), ctx.ProductConfig.Name)

	rowNo++
	rowNo++
	rowNo++
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style3)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), sheet1)
	rowNo++
	rowNo++

	t := time.Now()
	toDate := t.Format("02-January-2006")
	reportFromMsg := "Report Generated on" + " " + toDate
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style3)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), reportFromMsg)
	rowNo++

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "J", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "S.No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "Date")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "No.Of.Holdings")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), "Arrer Amount")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), "ArrearPenalty")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), "CurrentAmount")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), "CurrentPenalty")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "H", rowNo), "RebateAmount")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "I", rowNo), "AdvanceAmount")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "J", rowNo), "TotalAmount")
	rowNo++
	var totalAmount float64

	for i, v := range data.Records {

		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), i+1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v.ID.DayStr)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), v.PropertyCount)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), v.ArrearCollection)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), v.ArrearPenalty)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), v.CurrentCollection)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), v.CurrentPenalty)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "H", rowNo), v.RebateAmount)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "I", rowNo), v.AdvanceAmount)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "J", rowNo), v.TotalCollection)
		totalAmount = totalAmount + v.TotalCollection
		rowNo++

	}
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "B", rowNo), fmt.Sprintf("%v%v", "B", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "D", rowNo), fmt.Sprintf("%v%v", "D", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "E", rowNo), fmt.Sprintf("%v%v", "E", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "F", rowNo), fmt.Sprintf("%v%v", "F", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "G", rowNo), fmt.Sprintf("%v%v", "G", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "H", rowNo), fmt.Sprintf("%v%v", "H", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "I", rowNo), fmt.Sprintf("%v%v", "I", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "J", rowNo), fmt.Sprintf("%v%v", "J", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "Total")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "J", rowNo), fmt.Sprintf(" %.2f", totalAmount))

	return excel, nil

}

// /DashboardDayWiseCollectionChartPDF : ""
func (s *Service) DashboardDayWiseCollectionChartPDF(ctx *models.Context, filter *models.DashboardDayWiseCollectionChartFilter) ([]byte, error) {
	properties, err := s.DashboardDayWiseCollectionChart(ctx, filter)
	if err != nil {
		return nil, err
	}
	m := make(map[string]interface{})
	m2 := make(map[string]interface{})
	var totalCollection float64
	for _, v := range properties.Records {
		totalCollection = totalCollection + v.TotalTax
	}
	m["demand"] = properties
	m["totalCollection"] = totalCollection
	m2["currentDate"] = time.Now()
	cfy, err := s.Daos.GetCurrentFinancialYear(ctx)
	if err != nil {
		return nil, errors.New("Error in getting current financial year " + err.Error())
	}
	m2["currentFy"] = cfy
	var pdfdata models.PDFData
	pdfdata.Data = m
	pdfdata.RefData = m2
	productConfigUniqueID := "1"
	productConfig, err := s.Daos.GetSingleProductConfiguration(ctx, productConfigUniqueID)
	if err != nil {
		return nil, errors.New("Error in getting product config" + err.Error())
	}
	pdfdata.Config = productConfig.ProductConfiguration

	r := NewRequestPdf("")
	templatePathStart := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.TEMPLATELOC)
	//html template path
	templatePath := templatePathStart + "property_daywise_collection_report.html"
	if err := r.ParseTemplate(templatePath, pdfdata); err == nil {
		fmt.Println("start pdf generated successfully")

		ok, data, err := r.GeneratePDFAsFile()

		fmt.Println(ok, "pdf generated successfully")
		return data, err
	} else {
		fmt.Println("Error in parcing template - " + err.Error())

		return nil, errors.New("Error in parcing template - " + err.Error())
	}
	return nil, nil

}

// PropertyOverallDemandReportExcelV2 : ""
func (s *Service) PropertyOverallDemandReportExcelV2(ctx *models.Context, filter *models.PropertyFilter, pagination *models.Pagination) (*excelize.File, error) {
	res, err := s.PropertyOverallDemandReport(ctx, filter, pagination)
	if err != nil {
		return nil, err
	}
	cfy, err := s.Daos.GetCurrentFinancialYear(ctx)
	if err != nil {
		return nil, errors.New("Error in getting current financial year " + err.Error())
	}
	fmt.Println("'res length==>'", len(res))

	//  create an excel file
	excel := excelize.NewFile()
	sheet1 := "PropertyOverallDemand"
	rowNo := 1
	index := excel.NewSheet(sheet1)
	excel.SetActiveSheet(index)
	excel.MergeCell(sheet1, "A1", "B5")
	excel.MergeCell(sheet1, "C1", "M3")
	excel.MergeCell(sheet1, "C4", "M5")

	style1, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#FFDC6D"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	style3, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#B6DDE8"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	// style2, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#B6DDE8"],"pattern":1},"font":{"bold":true}}`)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// style4, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#FFDC6D"],"pattern":1},"alignment":{"horizontal":"right","vertical":"center"},"font":{"bold":true}}`)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style3)
	documentUrl := s.Shared.Config.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.FILEURL)
	if err := excel.AddPicture(sheet1, fmt.Sprintf("%v%v", "A", rowNo), documentUrl+"municipal/logo.png", `{"x_scale": 0.115, "y_scale": 0.0935}`); err != nil {
		fmt.Println(err)
	}
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), s.Shared.Config.GetString(s.Shared.GetCmdArg(constants.ENV)+"."+constants.LOCATIONNAME))
	rowNo++
	rowNo++
	rowNo++

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style3)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Property Overall Demand")
	rowNo++
	rowNo++

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "M", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "S.No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "Ward No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Property No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), "Guardian Name")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), "MobileNo")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), "Owner Name")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), "Property Type")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "H", rowNo), "Address")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "I", rowNo), "Road Type")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "J", rowNo), "Due Upto Year")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "K", rowNo), "Current Demand")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "L", rowNo), "Arrear Demand")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "M", rowNo), "Total Demand")

	fmt.Println("'res length==>'", len(res))
	for i, v := range res {
		rowNo++
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), i+1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), func() string {
			if v.Ref.Address.Ward != nil {
				return v.Ref.Address.Ward.Name
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), v.UniqueID)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), func() string {
			if len(v.Owner) > 0 {
				return v.Owner[0].FatherRpanRhusband
			}
			return "NA"
		}())

		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), func() string {
			if len(v.Ref.PropertyOwner) > 0 {
				return v.Ref.PropertyOwner[0].Mobile
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), func() string {
			if len(v.Ref.PropertyOwner) > 0 {
				return v.Ref.PropertyOwner[0].Name
			}
			return "NA"
		}())
		if v.Ref.PropertyType != nil {
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), v.Ref.PropertyType.Name)
		} else {
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), "NA")

		}
		// excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "H", rowNo), func() string {
		// 	if len(v.Ref.Floors) > 0 {
		// 		return v.Ref.Floors[0].UsageType
		// 	}
		// 	return "NA"
		// }())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "H", rowNo), v.Address.AL1+v.Address.Al2)
		if v.Ref.RoadType != nil {
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "I", rowNo), v.Ref.RoadType.Name)
		} else {
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "I", rowNo), "NA")

		}
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "J", rowNo), func() string {
			if v.DOA != nil && cfy != nil && cfy.To != nil {
				return fmt.Sprintf("%v - %v", v.DOA.Year(), cfy.To.Year())
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "K", rowNo), v.Ref.Demand.Current.TotalTax)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "L", rowNo), v.Ref.Demand.Arrear.TotalTax)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "M", rowNo), v.Ref.Demand.Total.TotalTax)

	}
	var totalDemand, arrearDemand, currentDemand float64
	for _, v1 := range res {
		totalDemand = totalDemand + v1.Ref.Demand.Total.TotalTax
		arrearDemand = arrearDemand + v1.Ref.Demand.Arrear.TotalTax
		currentDemand = currentDemand + v1.Ref.Demand.Current.TotalTax
	}

	rowNo++
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "Total")

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "B", rowNo), fmt.Sprintf("%v%v", "B", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "D", rowNo), fmt.Sprintf("%v%v", "D", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "E", rowNo), fmt.Sprintf("%v%v", "E", rowNo), style1)

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "F", rowNo), fmt.Sprintf("%v%v", "F", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "G", rowNo), fmt.Sprintf("%v%v", "G", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "H", rowNo), fmt.Sprintf("%v%v", "H", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "I", rowNo), fmt.Sprintf("%v%v", "I", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "J", rowNo), fmt.Sprintf("%v%v", "J", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "K", rowNo), fmt.Sprintf("%v%v", "K", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "L", rowNo), fmt.Sprintf("%v%v", "L", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "M", rowNo), fmt.Sprintf("%v%v", "M", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "K", rowNo), currentDemand)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "L", rowNo), arrearDemand)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "M", rowNo), totalDemand)

	return excel, nil
}

// PropertyOverallDemandReportExcelV3 : "MinorChanges needed adding pending demand"
func (s *Service) PropertyOverallDemandReportExcelV3(ctx *models.Context, filter *models.PropertyFilter, pagination *models.Pagination) (*excelize.File, error) {
	res, err := s.PropertyOverallDemandReport(ctx, filter, pagination)
	if err != nil {
		return nil, err
	}
	cfy, err := s.Daos.GetCurrentFinancialYear(ctx)
	if err != nil {
		return nil, errors.New("Error in getting current financial year " + err.Error())
	}
	fmt.Println("'res length==>'", len(res))
	//  create an excel file
	excel := excelize.NewFile()
	sheet1 := "PropertyOverallDemand"
	rowNo := 1
	index := excel.NewSheet(sheet1)
	excel.SetActiveSheet(index)
	excel.MergeCell(sheet1, "A1", "B5")
	excel.MergeCell(sheet1, "C1", "N3")
	excel.MergeCell(sheet1, "C4", "N5")

	style1, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#FFDC6D"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	style3, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#B6DDE8"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	// style2, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#B6DDE8"],"pattern":1},"font":{"bold":true}}`)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// style4, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#FFDC6D"],"pattern":1},"alignment":{"horizontal":"right","vertical":"center"},"font":{"bold":true}}`)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style3)
	documentUrl := s.Shared.Config.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.FILEURL)
	if err := excel.AddPicture(sheet1, fmt.Sprintf("%v%v", "A", rowNo), documentUrl+"municipal/logo.png", `{"x_scale": 0.115, "y_scale": 0.0935}`); err != nil {
		fmt.Println(err)
	}
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), s.Shared.Config.GetString(s.Shared.GetCmdArg(constants.ENV)+"."+constants.LOCATIONNAME))
	rowNo++
	rowNo++
	rowNo++

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style3)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Property Overall Demand")
	rowNo++
	rowNo++

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "N", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "S.No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "Ward No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Property No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), "Guardian Name")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), "MobileNo")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), "Owner Name")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), "Property Type")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "H", rowNo), "Address")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "I", rowNo), "Road Type")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "J", rowNo), "Due Upto Year")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "K", rowNo), "Current Demand")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "L", rowNo), "Arrear Demand")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "M", rowNo), "Total Demand")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "N", rowNo), "Pending Demand")

	fmt.Println("'res length==>'", len(res))
	for i, v := range res {
		rowNo++
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), i+1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), func() string {
			if v.Ref.Address.Ward != nil {
				return v.Ref.Address.Ward.Name
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), v.UniqueID)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), func() string {
			if len(v.Owner) > 0 {
				return v.Owner[0].FatherRpanRhusband
			}
			return "NA"
		}())

		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), func() string {
			if len(v.Ref.PropertyOwner) > 0 {
				return v.Ref.PropertyOwner[0].Mobile
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), func() string {
			if len(v.Ref.PropertyOwner) > 0 {
				return v.Ref.PropertyOwner[0].Name
			}
			return "NA"
		}())
		if v.Ref.PropertyType != nil {
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), v.Ref.PropertyType.Name)
		} else {
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), "NA")

		}
		// excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "H", rowNo), func() string {
		// 	if len(v.Ref.Floors) > 0 {
		// 		return v.Ref.Floors[0].UsageType
		// 	}
		// 	return "NA"
		// }())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "H", rowNo), v.Address.AL1+v.Address.Al2)
		if v.Ref.RoadType != nil {
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "I", rowNo), v.Ref.RoadType.Name)
		} else {
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "I", rowNo), "NA")

		}
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "J", rowNo), func() string {
			if v.DOA != nil && cfy != nil && cfy.To != nil {
				return fmt.Sprintf("%v - %v", v.DOA.Year(), cfy.To.Year())
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "K", rowNo), v.Ref.Demand.Current.TotalTax)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "L", rowNo), v.Ref.Demand.Arrear.TotalTax)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "M", rowNo), v.Ref.Demand.Total.TotalTax)

		pendingtax0 := (v.Collection.TotalTax - (v.Ref.Demand.Total.TotalTax + v.Collection.Penalty))
		pendingtax0 = math.Floor(pendingtax0*-100) / 100
		//val := fmt.Sprintf("Value %v,======= %v - (%v + %v)", pendingtax0, v.Collection.TotalTax, v.Ref.Demand.Total.TotalTax, v.Collection.Penalty)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "N", rowNo), pendingtax0)
		//excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "O", rowNo), val)

	}
	var totalDemand, arrearDemand, currentDemand, totalPending float64
	for _, v1 := range res {
		pendingtax := (v1.Collection.TotalTax - (v1.Ref.Demand.Total.TotalTax + v1.Collection.Penalty))
		totalDemand = totalDemand + v1.Ref.Demand.Total.TotalTax
		arrearDemand = arrearDemand + v1.Ref.Demand.Arrear.TotalTax
		currentDemand = currentDemand + v1.Ref.Demand.Current.TotalTax
		totalPending = totalPending + pendingtax
	}

	rowNo++
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "Total")

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "B", rowNo), fmt.Sprintf("%v%v", "B", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "D", rowNo), fmt.Sprintf("%v%v", "D", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "E", rowNo), fmt.Sprintf("%v%v", "E", rowNo), style1)

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "F", rowNo), fmt.Sprintf("%v%v", "F", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "G", rowNo), fmt.Sprintf("%v%v", "G", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "H", rowNo), fmt.Sprintf("%v%v", "H", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "I", rowNo), fmt.Sprintf("%v%v", "I", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "J", rowNo), fmt.Sprintf("%v%v", "J", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "K", rowNo), fmt.Sprintf("%v%v", "K", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "L", rowNo), fmt.Sprintf("%v%v", "L", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "M", rowNo), fmt.Sprintf("%v%v", "M", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "N", rowNo), fmt.Sprintf("%v%v", "N", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "K", rowNo), currentDemand)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "L", rowNo), arrearDemand)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "M", rowNo), totalDemand)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "N", rowNo), totalPending)

	return excel, nil
}

//PropertyOverallDemandReport : ""
func (s *Service) PropertyOverallDemandReport(ctx *models.Context, propertyfilter *models.PropertyFilter, pagination *models.Pagination) (property []models.RefProperty, err error) {
	defer ctx.Session.EndSession(ctx.CTX)

	return s.Daos.PropertyOverallDemandReport(ctx, propertyfilter, pagination)
}

// FilterWardDayWisePropertyCollectionReport : ""
func (s *Service) FilterWardDayWisePropertyCollectionReport(ctx *models.Context, filter *models.WardDayWisePropertyCollectionReportFilter, pagination *models.Pagination) ([]models.WardDayWisePropertyCollectionReport, error) {
	return s.Daos.FilterWardDayWisePropertyCollectionReport(ctx, filter, pagination)
}

// FilterWardDayWisePropertyCollectionReportExcel: ""
func (s *Service) FilterWardDayWisePropertyCollectionReportExcel(ctx *models.Context, filter *models.WardDayWisePropertyCollectionReportFilter, pagination *models.Pagination) (*excelize.File, error) {
	data, err := s.FilterWardDayWisePropertyCollectionReport(ctx, filter, pagination)
	if err != nil {
		return nil, err
	}

	resPD, err1 := s.Daos.GetSingleDefaultProductConfiguration(ctx)
	if err1 != nil {
		return nil, err1
	}

	excel := excelize.NewFile()
	sheet1 := "Property Ward Wise Collection"
	index := excel.NewSheet(sheet1)
	rowNo := 1

	excel.SetActiveSheet(index)

	if resPD.LocationID == "Bhagalpur" {
		excel.MergeCell(sheet1, "A1", "E3")
		excel.MergeCell(sheet1, "A4", "E5")
	} else {
		excel.MergeCell(sheet1, "A1", "B5")
		excel.MergeCell(sheet1, "C1", "E3")
		excel.MergeCell(sheet1, "C4", "E5")
	}

	excel.MergeCell(sheet1, "A6", "E6")
	style1, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#FFDC6D"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	style3, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#B6DDE8"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}

	if resPD.LocationID != "Bhagalpur" {
		documentUrl := s.Shared.Config.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.FILEURL)
		if err := excel.AddPicture(sheet1, fmt.Sprintf("%v%v", "A", rowNo), documentUrl+"municipal/logo.png", `{"x_scale": 0.115, "y_scale": 0.0935}`); err != nil {
			fmt.Println(err)
		}
	}

	if resPD.LocationID == "Bhagalpur" {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), ctx.ProductConfig.Name)
	} else {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), ctx.ProductConfig.Name)
	}
	rowNo++
	rowNo++
	rowNo++

	if resPD.LocationID == "Bhagalpur" {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), sheet1)
	} else {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), sheet1)
	}
	rowNo++
	rowNo++

	t := time.Now()
	toDate := t.Format("02-January-2006")
	reportFromMsg := "Report Generated on" + " " + toDate
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style3)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), reportFromMsg)
	rowNo++

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "E", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "S.No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "Ward")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Paied Properties")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), "No of Payments")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), "Collection Amount")
	rowNo++
	var totalAmount, totalProperties, totalPayments float64
	for i, v := range data {
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), i+1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v.Ward.Name)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), v.Report.TotalNoProperties)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), v.Report.TotalNoPayments)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), v.Report.TotalCollections)
		totalProperties = totalProperties + v.Report.TotalNoProperties
		totalPayments = totalPayments + v.Report.TotalNoPayments
		totalAmount = totalAmount + v.Report.TotalCollections
		rowNo++

	}
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "B", rowNo), fmt.Sprintf("%v%v", "B", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "D", rowNo), fmt.Sprintf("%v%v", "D", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "E", rowNo), fmt.Sprintf("%v%v", "E", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "Total")

	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf(" %.2f", totalProperties))
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), fmt.Sprintf(" %.2f", totalPayments))
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), fmt.Sprintf(" %.2f", totalAmount))

	return excel, nil

}

// FilterWardDayWisePropertyCollectionReportPDF : ""
func (s *Service) FilterWardDayWisePropertyCollectionReportPDF(ctx *models.Context, filter *models.WardDayWisePropertyCollectionReportFilter, pagination *models.Pagination) ([]byte, error) {
	data, err := s.FilterWardDayWisePropertyCollectionReport(ctx, filter, pagination)
	if err != nil {
		return nil, err
	}
	var totalCollection float64
	for _, v := range data {
		totalCollection = totalCollection + v.Report.TotalCollections
	}
	m := make(map[string]interface{})
	m2 := make(map[string]interface{})
	m3 := make(map[string]interface{})
	m["demand"] = data
	m["totalCollection"] = totalCollection
	m2["currentDate"] = time.Now()
	m3["inc"] = func(a int) int {
		return a + 1
	}
	var pdfdata models.PDFData
	pdfdata.Data = m
	pdfdata.RefData = m2
	productConfigUniqueID := "1"
	productConfig, err := s.Daos.GetSingleProductConfiguration(ctx, productConfigUniqueID)
	if err != nil {
		return nil, errors.New("Error in getting product config" + err.Error())
	}
	pdfdata.Config = productConfig.ProductConfiguration

	r := NewRequestPdf("")
	templatePathStart := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.TEMPLATELOC)
	//html template path
	templatePath := templatePathStart + "property_wardwise_collection_report.html"
	if err := r.ParseTemplate(templatePath, pdfdata); err == nil {
		fmt.Println("start pdf generated successfully")

		ok, data, err := r.GeneratePDFAsFile()

		fmt.Println(ok, "pdf generated successfully")
		return data, err
	} else {
		fmt.Println("Error in parcing template - " + err.Error())

		return nil, errors.New("Error in parcing template - " + err.Error())
	}

}

// FilterWardMonthWisePropertyCollectionReport : ""
func (s *Service) FilterWardMonthWisePropertyCollectionReport(ctx *models.Context, filter *models.WardMonthWisePropertyCollectionReportFilter, pagination *models.Pagination) ([]models.WardMonthWisePropertyCollectionReport, error) {
	return s.Daos.FilterWardMonthWisePropertyCollectionReport(ctx, filter, pagination)
}

// FilterWardMonthWisePropertyCollectionReport: ""
func (s *Service) FilterWardMonthWisePropertyCollectionReportExcel(ctx *models.Context, filter *models.WardMonthWisePropertyCollectionReportFilter, pagination *models.Pagination) (*excelize.File, error) {
	data, err := s.FilterWardMonthWisePropertyCollectionReport(ctx, filter, pagination)
	if err != nil {
		return nil, err
	}

	resPD, err1 := s.Daos.GetSingleDefaultProductConfiguration(ctx)
	if err1 != nil {
		return nil, err1
	}

	excel := excelize.NewFile()
	sheet1 := "Property Ward Wise Collection"
	index := excel.NewSheet(sheet1)
	rowNo := 1

	excel.SetActiveSheet(index)
	if resPD.LocationID == "Bhagalpur" {
		excel.MergeCell(sheet1, "A1", "E3")
		excel.MergeCell(sheet1, "A4", "E5")
	} else {
		excel.MergeCell(sheet1, "A1", "B5")
		excel.MergeCell(sheet1, "C1", "E3")
		excel.MergeCell(sheet1, "C4", "E5")
	}

	excel.MergeCell(sheet1, "A6", "E6")
	style1, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#FFDC6D"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	style3, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#B6DDE8"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}

	if resPD.LocationID != "Bhagalpur" {
		documentUrl := s.Shared.Config.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.FILEURL)
		if err := excel.AddPicture(sheet1, fmt.Sprintf("%v%v", "A", rowNo), documentUrl+"municipal/logo.png", `{"x_scale": 0.115, "y_scale": 0.0935}`); err != nil {
			fmt.Println(err)
		}
	}

	if resPD.LocationID == "Bhagalpur" {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), ctx.ProductConfig.Name)
	} else {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), ctx.ProductConfig.Name)
	}

	rowNo++
	rowNo++
	rowNo++

	if resPD.LocationID == "Bhagalpur" {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), sheet1)
	} else {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), sheet1)
	}
	rowNo++
	rowNo++

	t := time.Now()
	toDate := t.Format("02-January-2006")
	reportFromMsg := "Report Generated on" + " " + toDate
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style3)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), reportFromMsg)
	rowNo++

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "E", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "S.No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "Ward")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Paied Properties")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), "No of Payments")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), "Collection Amount")
	rowNo++
	var totalAmount, totalProperties, totalPayments float64
	for i, v := range data {
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), i+1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v.Ward.Name)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), v.Report.TotalNoProperties)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), v.Report.TotalNoPayments)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), v.Report.TotalCollections)
		totalProperties = totalProperties + v.Report.TotalNoProperties
		totalPayments = totalPayments + v.Report.TotalNoPayments
		totalAmount = totalAmount + v.Report.TotalCollections
		rowNo++

	}
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "B", rowNo), fmt.Sprintf("%v%v", "B", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "D", rowNo), fmt.Sprintf("%v%v", "D", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "E", rowNo), fmt.Sprintf("%v%v", "E", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "Total")

	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf(" %.2f", totalProperties))
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), fmt.Sprintf(" %.2f", totalPayments))
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), fmt.Sprintf(" %.2f", totalAmount))

	return excel, nil

}

// FilterWardDayWisePropertyCollectionReportPDF : ""
func (s *Service) FilterWardMonthWisePropertyCollectionReportPDF(ctx *models.Context, filter *models.WardMonthWisePropertyCollectionReportFilter, pagination *models.Pagination) ([]byte, error) {
	data, err := s.FilterWardMonthWisePropertyCollectionReport(ctx, filter, pagination)
	if err != nil {
		return nil, err
	}
	var totalCollection float64
	for _, v := range data {
		totalCollection = totalCollection + v.Report.TotalCollections
	}
	m := make(map[string]interface{})
	m2 := make(map[string]interface{})
	m3 := make(map[string]interface{})
	m["demand"] = data
	m["totalCollection"] = totalCollection
	m2["currentDate"] = time.Now()
	m3["inc"] = func(a int) int {
		return a + 1
	}
	var pdfdata models.PDFData
	pdfdata.Data = m
	pdfdata.RefData = m2
	productConfigUniqueID := "1"
	productConfig, err := s.Daos.GetSingleProductConfiguration(ctx, productConfigUniqueID)
	if err != nil {
		return nil, errors.New("Error in getting product config" + err.Error())
	}
	pdfdata.Config = productConfig.ProductConfiguration

	r := NewRequestPdf("")
	templatePathStart := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.TEMPLATELOC)
	//html template path
	templatePath := templatePathStart + "property_wardwisemonth_collection_report.html"
	if err := r.ParseTemplate(templatePath, pdfdata); err == nil {
		fmt.Println("start pdf generated successfully")

		ok, data, err := r.GeneratePDFAsFile()

		fmt.Println(ok, "pdf generated successfully")
		return data, err
	} else {
		fmt.Println("Error in parcing template - " + err.Error())

		return nil, errors.New("Error in parcing template - " + err.Error())
	}

}

// FilterWardDayWisePropertyDemandReport : ""
func (s *Service) FilterWardDayWisePropertyDemandReport(ctx *models.Context, filter *models.WardDayWisePropertyDemandReportFilter, pagination *models.Pagination) ([]models.WardDayWisePropertyDemandReport, error) {
	return s.Daos.FilterWardDayWisePropertyDemandReport(ctx, filter, pagination)
}

// FilterWardDayWisePropertyDemandReportExcel: ""
func (s *Service) FilterWardDayWisePropertyDemandReportExcel(ctx *models.Context, filter *models.WardDayWisePropertyDemandReportFilter, pagination *models.Pagination) (*excelize.File, error) {
	data, err := s.FilterWardDayWisePropertyDemandReport(ctx, filter, pagination)
	if err != nil {
		return nil, err
	}

	resPD, err1 := s.Daos.GetSingleDefaultProductConfiguration(ctx)
	if err1 != nil {
		return nil, err1
	}

	excel := excelize.NewFile()
	sheet1 := "Property Ward Wise Demand"
	index := excel.NewSheet(sheet1)
	rowNo := 1

	excel.SetActiveSheet(index)
	if resPD.LocationID == "Bhagalpur" {
		excel.MergeCell(sheet1, "A1", "C3")
		excel.MergeCell(sheet1, "A4", "C5")
	} else {
		excel.MergeCell(sheet1, "A1", "B5")
		excel.MergeCell(sheet1, "C1", "C3")
		excel.MergeCell(sheet1, "C4", "C5")
	}

	excel.MergeCell(sheet1, "A6", "C6")
	style1, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#FFDC6D"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	style3, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#B6DDE8"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}

	if resPD.LocationID != "Bhagalpur" {
		documentUrl := s.Shared.Config.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.FILEURL)
		if err := excel.AddPicture(sheet1, fmt.Sprintf("%v%v", "A", rowNo), documentUrl+"municipal/logo.png", `{"x_scale": 0.115, "y_scale": 0.0935}`); err != nil {
			fmt.Println(err)
		}
	}

	if resPD.LocationID == "Bhagalpur" {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), ctx.ProductConfig.Name)
	} else {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), ctx.ProductConfig.Name)
	}

	rowNo++
	rowNo++
	rowNo++

	if resPD.LocationID == "Bhagalpur" {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), sheet1)
	} else {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), sheet1)
	}
	rowNo++
	rowNo++

	t := time.Now()
	toDate := t.Format("02-January-2006")
	reportFromMsg := "Report Generated on" + " " + toDate
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style3)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), reportFromMsg)
	rowNo++

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "S.No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "No of Properties")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Amount")
	rowNo++
	var totalAmount float64
	for i, v := range data {
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), i+1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v.Report.TotalNoProperties)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), v.Report.TotalDemand)
		totalAmount = totalAmount + v.Report.TotalDemand
		rowNo++

	}
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "B", rowNo), fmt.Sprintf("%v%v", "B", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "Total")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf(" %.2f", totalAmount))

	return excel, nil

}

// FilterWardDayWisePropertyDemandReportPDF : ""
func (s *Service) FilterWardDayWisePropertyDemandReportPDF(ctx *models.Context, filter *models.WardDayWisePropertyDemandReportFilter, pagination *models.Pagination) ([]byte, error) {
	data, err := s.FilterWardDayWisePropertyDemandReport(ctx, filter, pagination)
	if err != nil {
		return nil, err
	}
	var totalCollection float64
	for _, v := range data {
		totalCollection = totalCollection + v.Report.TotalDemand
	}
	m := make(map[string]interface{})
	m2 := make(map[string]interface{})
	m3 := make(map[string]interface{})
	m["demand"] = data
	m["totalCollection"] = totalCollection
	m2["currentDate"] = time.Now()
	m3["inc"] = func(a int) int {
		return a + 1
	}
	var pdfdata models.PDFData
	pdfdata.Data = m
	pdfdata.RefData = m2
	productConfigUniqueID := "1"
	productConfig, err := s.Daos.GetSingleProductConfiguration(ctx, productConfigUniqueID)
	if err != nil {
		return nil, errors.New("Error in getting product config" + err.Error())
	}
	pdfdata.Config = productConfig.ProductConfiguration

	r := NewRequestPdf("")
	templatePathStart := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.TEMPLATELOC)
	//html template path
	templatePath := templatePathStart + "property_wardwise_demand_report.html"
	if err := r.ParseTemplate(templatePath, pdfdata); err == nil {
		fmt.Println("start pdf generated successfully")

		ok, data, err := r.GeneratePDFAsFile()

		fmt.Println(ok, "pdf generated successfully")
		return data, err
	} else {
		fmt.Println("Error in parcing template - " + err.Error())

		return nil, errors.New("Error in parcing template - " + err.Error())
	}

}

// FilterWardMonthWisePropertyDemandReport : ""
func (s *Service) FilterWardMonthWisePropertyDemandReport(ctx *models.Context, filter *models.WardMonthWisePropertyDemandReportFilter, pagination *models.Pagination) ([]models.WardMonthWisePropertyDemandReport, error) {
	return s.Daos.FilterWardMonthWisePropertyDemandReport(ctx, filter, pagination)
}

// FilterWardMonthWisePropertyDemandReportExcel: ""
func (s *Service) FilterWardMonthWisePropertyDemandReportExcel(ctx *models.Context, filter *models.WardMonthWisePropertyDemandReportFilter, pagination *models.Pagination) (*excelize.File, error) {
	data, err := s.FilterWardMonthWisePropertyDemandReport(ctx, filter, pagination)
	if err != nil {
		return nil, err
	}

	resPD, err1 := s.Daos.GetSingleDefaultProductConfiguration(ctx)
	if err1 != nil {
		return nil, err1
	}

	excel := excelize.NewFile()
	sheet1 := "Property Ward Wise Demand"
	index := excel.NewSheet(sheet1)
	rowNo := 1

	excel.SetActiveSheet(index)
	if resPD.LocationID == "Bhagalpur" {
		excel.MergeCell(sheet1, "A1", "C3")
		excel.MergeCell(sheet1, "A4", "C5")
	} else {
		excel.MergeCell(sheet1, "A1", "B5")
		excel.MergeCell(sheet1, "C1", "C3")
		excel.MergeCell(sheet1, "C4", "C5")
	}

	excel.MergeCell(sheet1, "A6", "C6")
	style1, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#FFDC6D"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	style3, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#B6DDE8"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}

	if resPD.LocationID != "Bhagalpur" {
		documentUrl := s.Shared.Config.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.FILEURL)
		if err := excel.AddPicture(sheet1, fmt.Sprintf("%v%v", "A", rowNo), documentUrl+"municipal/logo.png", `{"x_scale": 0.115, "y_scale": 0.0935}`); err != nil {
			fmt.Println(err)
		}
	}

	if resPD.LocationID == "Bhagalpur" {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), ctx.ProductConfig.Name)
	} else {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), ctx.ProductConfig.Name)
	}

	rowNo++
	rowNo++
	rowNo++

	if resPD.LocationID == "Bhagalpur" {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), sheet1)
	} else {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), sheet1)
	}
	rowNo++
	rowNo++

	t := time.Now()
	toDate := t.Format("02-January-2006")
	reportFromMsg := "Report Generated on" + " " + toDate
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style3)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), reportFromMsg)
	rowNo++

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "S.No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "No of Properties")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Amount")
	rowNo++
	var totalAmount float64
	for i, v := range data {
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), i+1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v.Report.TotalNoProperties)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), v.Report.TotalDemand)
		totalAmount = totalAmount + v.Report.TotalDemand
		rowNo++

	}
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "B", rowNo), fmt.Sprintf("%v%v", "B", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "Total")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf(" %.2f", totalAmount))

	return excel, nil

}

// FilterWardMonthWisePropertyDemandReportPDF : ""
func (s *Service) FilterWardMonthWisePropertyDemandReportPDF(ctx *models.Context, filter *models.WardMonthWisePropertyDemandReportFilter, pagination *models.Pagination) ([]byte, error) {
	data, err := s.FilterWardMonthWisePropertyDemandReport(ctx, filter, pagination)
	if err != nil {
		return nil, err
	}
	var totalCollection float64
	for _, v := range data {
		totalCollection = totalCollection + v.Report.TotalDemand
	}
	m := make(map[string]interface{})
	m2 := make(map[string]interface{})
	m3 := make(map[string]interface{})
	m["demand"] = data
	m["totalCollection"] = totalCollection
	m2["currentDate"] = time.Now()
	m3["inc"] = func(a int) int {
		return a + 1
	}
	var pdfdata models.PDFData
	pdfdata.Data = m
	pdfdata.RefData = m2
	productConfigUniqueID := "1"
	productConfig, err := s.Daos.GetSingleProductConfiguration(ctx, productConfigUniqueID)
	if err != nil {
		return nil, errors.New("Error in getting product config" + err.Error())
	}
	pdfdata.Config = productConfig.ProductConfiguration

	r := NewRequestPdf("")
	templatePathStart := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.TEMPLATELOC)
	//html template path
	templatePath := templatePathStart + "property_wardwisemonth_demand_report.html"
	if err := r.ParseTemplate(templatePath, pdfdata); err == nil {
		fmt.Println("start pdf generated successfully")

		ok, data, err := r.GeneratePDFAsFile()

		fmt.Println(ok, "pdf generated successfully")
		return data, err
	} else {
		fmt.Println("Error in parcing template - " + err.Error())

		return nil, errors.New("Error in parcing template - " + err.Error())
	}

}

// FilterTeamDayWisePropertyCollectionReport : ""
func (s *Service) FilterTeamDayWisePropertyCollectionReport(ctx *models.Context, filter *models.TeamDayWisePropertyCollectionReportFilter, pagination *models.Pagination) ([]models.TeamDayWisePropertyCollectionReport, error) {
	return s.Daos.FilterTeamDayWisePropertyCollectionReport(ctx, filter, pagination)
}

// FilterTeamDayWisePropertyCollectionReportExcel: ""
func (s *Service) FilterTeamDayWisePropertyCollectionReportExcel(ctx *models.Context, filter *models.TeamDayWisePropertyCollectionReportFilter, pagination *models.Pagination) (*excelize.File, error) {
	data, err := s.FilterTeamDayWisePropertyCollectionReport(ctx, filter, pagination)
	if err != nil {
		return nil, err
	}

	resPD, err1 := s.Daos.GetSingleDefaultProductConfiguration(ctx)
	if err1 != nil {
		return nil, err1
	}

	excel := excelize.NewFile()
	sheet1 := "Property Team Wise Collection"
	index := excel.NewSheet(sheet1)
	rowNo := 1

	excel.SetActiveSheet(index)
	if resPD.LocationID == "Bhagalpur" {
		excel.MergeCell(sheet1, "A1", "F3")
		excel.MergeCell(sheet1, "A4", "F5")
	} else {
		excel.MergeCell(sheet1, "A1", "B5")
		excel.MergeCell(sheet1, "C1", "F3")
		excel.MergeCell(sheet1, "C4", "F5")
	}

	excel.MergeCell(sheet1, "A6", "F6")
	style1, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#FFDC6D"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	style3, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#B6DDE8"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}

	if resPD.LocationID != "Bhagalpur" {
		documentUrl := s.Shared.Config.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.FILEURL)
		if err := excel.AddPicture(sheet1, fmt.Sprintf("%v%v", "A", rowNo), documentUrl+"municipal/logo.png", `{"x_scale": 0.115, "y_scale": 0.0935}`); err != nil {
			fmt.Println(err)
		}
	}

	if resPD.LocationID == "Bhagalpur" {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), ctx.ProductConfig.Name)
	} else {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), ctx.ProductConfig.Name)
	}

	rowNo++
	rowNo++
	rowNo++

	if resPD.LocationID == "Bhagalpur" {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), sheet1)
	} else {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), sheet1)
	}
	rowNo++
	rowNo++

	t := time.Now()
	toDate := t.Format("02-January-2006")
	reportFromMsg := "Report Generated on" + " " + toDate
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style3)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), reportFromMsg)
	rowNo++

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "F", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "S.No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "User")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "UserType")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), "Paied Properties")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), "No of Payments")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), "Collection Amount")
	rowNo++
	var totalAmount float64
	var totalProperties float64
	var totalPayments float64
	for i, v := range data {
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), i+1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v.User.Name)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), v.User.Type)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), v.Report.TotalNoProperties)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), v.Report.TotalNoPayments)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), v.Report.TotalCollections)
		totalAmount = totalAmount + v.Report.TotalCollections
		totalProperties = totalProperties + v.Report.TotalNoProperties
		totalPayments = totalPayments + v.Report.TotalNoPayments
		rowNo++

	}
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "B", rowNo), fmt.Sprintf("%v%v", "B", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "D", rowNo), fmt.Sprintf("%v%v", "D", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "E", rowNo), fmt.Sprintf("%v%v", "E", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "F", rowNo), fmt.Sprintf("%v%v", "F", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "Total")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), fmt.Sprintf(" %.2f", totalProperties))
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), fmt.Sprintf(" %.2f", totalPayments))
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), fmt.Sprintf(" %.2f", totalAmount))

	return excel, nil

}

// FilterTeamDayWisePropertyCollectionReportPDF : ""
func (s *Service) FilterTeamDayWisePropertyCollectionReportPDF(ctx *models.Context, filter *models.TeamDayWisePropertyCollectionReportFilter, pagination *models.Pagination) ([]byte, error) {
	data, err := s.FilterTeamDayWisePropertyCollectionReport(ctx, filter, pagination)
	if err != nil {
		return nil, err
	}
	var totalCollection float64
	for _, v := range data {
		totalCollection = totalCollection + v.Report.TotalCollections
	}
	m := make(map[string]interface{})
	m2 := make(map[string]interface{})
	m3 := make(map[string]interface{})
	m["demand"] = data
	m["totalCollection"] = totalCollection
	m2["currentDate"] = time.Now()
	m3["inc"] = func(a int) int {
		return a + 1
	}
	var pdfdata models.PDFData
	pdfdata.Data = m
	pdfdata.RefData = m2
	productConfigUniqueID := "1"
	productConfig, err := s.Daos.GetSingleProductConfiguration(ctx, productConfigUniqueID)
	if err != nil {
		return nil, errors.New("Error in getting product config" + err.Error())
	}
	pdfdata.Config = productConfig.ProductConfiguration

	r := NewRequestPdf("")
	templatePathStart := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.TEMPLATELOC)
	//html template path
	templatePath := templatePathStart + "property_teamwise_collection_report.html"
	if err := r.ParseTemplate(templatePath, pdfdata); err == nil {
		fmt.Println("start pdf generated successfully")

		ok, data, err := r.GeneratePDFAsFile()

		fmt.Println(ok, "pdf generated successfully")
		return data, err
	} else {
		fmt.Println("Error in parcing template - " + err.Error())

		return nil, errors.New("Error in parcing template - " + err.Error())
	}

}

// FilterTeamMonthWisePropertyCollectionReport : ""
func (s *Service) FilterTeamMonthWisePropertyCollectionReport(ctx *models.Context, filter *models.TeamMonthWisePropertyCollectionReportFilter, pagination *models.Pagination) ([]models.TeamMonthWisePropertyCollectionReport, error) {
	return s.Daos.FilterTeamMonthWisePropertyCollectionReport(ctx, filter, pagination)
}

// FilterTeamMonthWisePropertyCollectionReportExcel: ""
func (s *Service) FilterTeamMonthWisePropertyCollectionReportExcel(ctx *models.Context, filter *models.TeamMonthWisePropertyCollectionReportFilter, pagination *models.Pagination) (*excelize.File, error) {
	data, err := s.FilterTeamMonthWisePropertyCollectionReport(ctx, filter, pagination)
	if err != nil {
		return nil, err
	}

	resPD, err1 := s.Daos.GetSingleDefaultProductConfiguration(ctx)
	if err1 != nil {
		return nil, err1
	}

	excel := excelize.NewFile()
	sheet1 := "Property Team Wise Collection"
	index := excel.NewSheet(sheet1)
	rowNo := 1

	excel.SetActiveSheet(index)
	if resPD.LocationID == "Bhagalpur" {
		excel.MergeCell(sheet1, "A1", "F3")
		excel.MergeCell(sheet1, "A4", "F5")
	} else {
		excel.MergeCell(sheet1, "A1", "B5")
		excel.MergeCell(sheet1, "C1", "F3")
		excel.MergeCell(sheet1, "C4", "F5")
	}

	excel.MergeCell(sheet1, "A6", "F6")
	style1, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#FFDC6D"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	style3, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#B6DDE8"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}

	if resPD.LocationID != "Bhagalpur" {
		documentUrl := s.Shared.Config.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.FILEURL)
		if err := excel.AddPicture(sheet1, fmt.Sprintf("%v%v", "A", rowNo), documentUrl+"municipal/logo.png", `{"x_scale": 0.115, "y_scale": 0.0935}`); err != nil {
			fmt.Println(err)
		}
	}

	if resPD.LocationID == "Bhagalpur" {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), ctx.ProductConfig.Name)
	} else {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), ctx.ProductConfig.Name)
	}

	rowNo++
	rowNo++
	rowNo++

	if resPD.LocationID == "Bhagalpur" {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), sheet1)
	} else {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), sheet1)
	}
	rowNo++
	rowNo++

	t := time.Now()
	toDate := t.Format("02-January-2006")
	reportFromMsg := "Report Generated on" + " " + toDate
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style3)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), reportFromMsg)
	rowNo++

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "F", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "S.No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "User")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "UserType")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), "Paied Properties")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), "No of Payments")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), "Collection Amount")
	rowNo++
	var totalAmount, totalProperties, totalPayments float64
	for i, v := range data {
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), i+1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v.User.Name)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), v.User.Type)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), v.Report.TotalNoProperties)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), v.Report.TotalNoPayments)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), v.Report.TotalCollections)
		totalAmount = totalAmount + v.Report.TotalCollections
		totalProperties = totalProperties + v.Report.TotalNoProperties
		totalPayments = totalPayments + v.Report.TotalNoPayments
		rowNo++

	}
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "B", rowNo), fmt.Sprintf("%v%v", "B", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "D", rowNo), fmt.Sprintf("%v%v", "D", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "E", rowNo), fmt.Sprintf("%v%v", "E", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "F", rowNo), fmt.Sprintf("%v%v", "F", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "Total")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), fmt.Sprintf(" %.2f", totalProperties))
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), fmt.Sprintf(" %.2f", totalPayments))
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), fmt.Sprintf(" %.2f", totalAmount))

	return excel, nil

}

// FilterTeamMonthWisePropertyCollectionReportPDF : ""
func (s *Service) FilterTeamMonthWisePropertyCollectionReportPDF(ctx *models.Context, filter *models.TeamMonthWisePropertyCollectionReportFilter, pagination *models.Pagination) ([]byte, error) {
	data, err := s.FilterTeamMonthWisePropertyCollectionReport(ctx, filter, pagination)
	if err != nil {
		return nil, err
	}
	var totalCollection float64
	for _, v := range data {
		totalCollection = totalCollection + v.Report.TotalCollections
	}
	m := make(map[string]interface{})
	m2 := make(map[string]interface{})
	m3 := make(map[string]interface{})
	m["demand"] = data
	m["totalCollection"] = totalCollection
	m2["currentDate"] = time.Now()
	m3["inc"] = func(a int) int {
		return a + 1
	}
	var pdfdata models.PDFData
	pdfdata.Data = m
	pdfdata.RefData = m2
	productConfigUniqueID := "1"
	productConfig, err := s.Daos.GetSingleProductConfiguration(ctx, productConfigUniqueID)
	if err != nil {
		return nil, errors.New("Error in getting product config" + err.Error())
	}
	pdfdata.Config = productConfig.ProductConfiguration

	r := NewRequestPdf("")
	templatePathStart := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.TEMPLATELOC)
	//html template path
	templatePath := templatePathStart + "property_teamwisemonth_collection_report.html"
	if err := r.ParseTemplate(templatePath, pdfdata); err == nil {
		fmt.Println("start pdf generated successfully")

		ok, data, err := r.GeneratePDFAsFile()

		fmt.Println(ok, "pdf generated successfully")
		return data, err
	} else {
		fmt.Println("Error in parcing template - " + err.Error())

		return nil, errors.New("Error in parcing template - " + err.Error())
	}

}

// FilterWardYearWisePropertyCollectionReport : ""
func (s *Service) FilterWardYearWisePropertyCollectionReport(ctx *models.Context, filter *models.WardYearWisePropertyCollectionReportFilter, pagination *models.Pagination) ([]models.WardYearWisePropertyCollectionReport, error) {
	return s.Daos.FilterWardYearWisePropertyCollectionReport(ctx, filter, pagination)
}

// FilterWardYearWisePropertyCollectionReportExcel: ""
func (s *Service) FilterWardYearWisePropertyCollectionReportExcel(ctx *models.Context, filter *models.WardYearWisePropertyCollectionReportFilter, pagination *models.Pagination) (*excelize.File, error) {
	data, err := s.FilterWardYearWisePropertyCollectionReport(ctx, filter, pagination)
	if err != nil {
		return nil, err
	}

	resPD, err1 := s.Daos.GetSingleDefaultProductConfiguration(ctx)
	if err1 != nil {
		return nil, err1
	}

	excel := excelize.NewFile()
	sheet1 := "Property Ward Wise Collection"
	index := excel.NewSheet(sheet1)
	rowNo := 1

	excel.SetActiveSheet(index)
	if resPD.LocationID == "Bhagalpur" {
		excel.MergeCell(sheet1, "A1", "E3")
		excel.MergeCell(sheet1, "A4", "E5")
	} else {
		excel.MergeCell(sheet1, "A1", "B5")
		excel.MergeCell(sheet1, "C1", "E3")
		excel.MergeCell(sheet1, "C4", "E5")
	}

	excel.MergeCell(sheet1, "A6", "E6")
	style1, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#FFDC6D"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	style3, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#B6DDE8"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}

	if resPD.LocationID != "Bhagalpur" {
		documentUrl := s.Shared.Config.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.FILEURL)
		if err := excel.AddPicture(sheet1, fmt.Sprintf("%v%v", "A", rowNo), documentUrl+"municipal/logo.png", `{"x_scale": 0.115, "y_scale": 0.0935}`); err != nil {
			fmt.Println(err)
		}
	}

	if resPD.LocationID == "Bhagalpur" {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), ctx.ProductConfig.Name)
	} else {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), ctx.ProductConfig.Name)
	}
	rowNo++
	rowNo++
	rowNo++

	if resPD.LocationID == "Bhagalpur" {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), sheet1)
	} else {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), sheet1)
	}
	rowNo++
	rowNo++

	t := time.Now()
	toDate := t.Format("02-January-2006")
	reportFromMsg := "Report Generated on" + " " + toDate
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style3)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), reportFromMsg)
	rowNo++

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "E", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "S.No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "Ward")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Paied Properties")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), "No of Payments")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), "Collection Amount")
	rowNo++
	var totalAmount, totalPayments, totalProperties float64
	for i, v := range data {
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), i+1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v.Ward.Name)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), v.Report.TotalNoProperties)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), v.Report.TotalNoPayments)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), v.Report.TotalCollections)
		totalProperties = totalProperties + v.Report.TotalNoProperties
		totalPayments = totalPayments + v.Report.TotalNoPayments
		totalAmount = totalAmount + v.Report.TotalCollections
		rowNo++

	}
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "B", rowNo), fmt.Sprintf("%v%v", "B", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "D", rowNo), fmt.Sprintf("%v%v", "D", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "E", rowNo), fmt.Sprintf("%v%v", "E", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "Total")

	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf(" %.2f", totalProperties))
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), fmt.Sprintf(" %.2f", totalPayments))
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), fmt.Sprintf(" %.2f", totalAmount))

	return excel, nil

}

// FilterWardYearWisePropertyCollectionReportPDF : ""
func (s *Service) FilterWardYearWisePropertyCollectionReportPDF(ctx *models.Context, filter *models.WardYearWisePropertyCollectionReportFilter, pagination *models.Pagination) ([]byte, error) {
	data, err := s.FilterWardYearWisePropertyCollectionReport(ctx, filter, pagination)
	if err != nil {
		return nil, err
	}
	var totalCollection float64
	for _, v := range data {
		totalCollection = totalCollection + v.Report.TotalCollections
	}
	m := make(map[string]interface{})
	m2 := make(map[string]interface{})
	m3 := make(map[string]interface{})
	m["demand"] = data
	m["totalCollection"] = totalCollection
	m2["currentDate"] = time.Now()
	m3["inc"] = func(a int) int {
		return a + 1
	}
	var pdfdata models.PDFData
	pdfdata.Data = m
	pdfdata.RefData = m2
	productConfigUniqueID := "1"
	productConfig, err := s.Daos.GetSingleProductConfiguration(ctx, productConfigUniqueID)
	if err != nil {
		return nil, errors.New("Error in getting product config" + err.Error())
	}
	pdfdata.Config = productConfig.ProductConfiguration

	r := NewRequestPdf("")
	templatePathStart := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.TEMPLATELOC)
	//html template path
	templatePath := templatePathStart + "property_wardwiseyear_collection_report.html"
	if err := r.ParseTemplate(templatePath, pdfdata); err == nil {
		fmt.Println("start pdf generated successfully")

		ok, data, err := r.GeneratePDFAsFile()

		fmt.Println(ok, "pdf generated successfully")
		return data, err
	} else {
		fmt.Println("Error in parcing template - " + err.Error())

		return nil, errors.New("Error in parcing template - " + err.Error())
	}

}

// FilterWardYearWisePropertyDemandReport : ""
func (s *Service) FilterWardYearWisePropertyDemandReport(ctx *models.Context, filter *models.WardYearWisePropertyDemandReportFilter, pagination *models.Pagination) ([]models.WardYearWisePropertyDemandReport, error) {
	return s.Daos.FilterWardYearWisePropertyDemandReport(ctx, filter, pagination)
}

// FilterWardYearWisePropertyDemandReportExcel: ""
func (s *Service) FilterWardYearWisePropertyDemandReportExcel(ctx *models.Context, filter *models.WardYearWisePropertyDemandReportFilter, pagination *models.Pagination) (*excelize.File, error) {
	data, err := s.FilterWardYearWisePropertyDemandReport(ctx, filter, pagination)
	if err != nil {
		return nil, err
	}

	resPD, err1 := s.Daos.GetSingleDefaultProductConfiguration(ctx)
	if err1 != nil {
		return nil, err1
	}

	excel := excelize.NewFile()
	sheet1 := "Property Ward Wise Demand"
	index := excel.NewSheet(sheet1)
	rowNo := 1

	excel.SetActiveSheet(index)
	if resPD.LocationID == "Bhagalpur" {
		excel.MergeCell(sheet1, "A1", "C3")
		excel.MergeCell(sheet1, "A4", "C5")
	} else {
		excel.MergeCell(sheet1, "A1", "B5")
		excel.MergeCell(sheet1, "C1", "C3")
		excel.MergeCell(sheet1, "C4", "C5")
	}

	excel.MergeCell(sheet1, "A6", "C6")
	style1, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#FFDC6D"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	style3, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#B6DDE8"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}

	if resPD.LocationID != "Bhagalpur" {
		documentUrl := s.Shared.Config.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.FILEURL)
		if err := excel.AddPicture(sheet1, fmt.Sprintf("%v%v", "A", rowNo), documentUrl+"municipal/logo.png", `{"x_scale": 0.115, "y_scale": 0.0935}`); err != nil {
			fmt.Println(err)
		}
	}

	if resPD.LocationID == "Bhagalpur" {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), ctx.ProductConfig.Name)
	} else {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), ctx.ProductConfig.Name)
	}

	rowNo++
	rowNo++
	rowNo++

	if resPD.LocationID == "Bhagalpur" {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), sheet1)
	} else {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), sheet1)
	}
	rowNo++
	rowNo++

	t := time.Now()
	toDate := t.Format("02-January-2006")
	reportFromMsg := "Report Generated on" + " " + toDate
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style3)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), reportFromMsg)
	rowNo++

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "S.No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "No of Properties")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Amount")
	rowNo++
	var totalAmount float64
	for i, v := range data {
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), i+1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v.Report.TotalNoProperties)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), v.Report.TotalDemand)
		totalAmount = totalAmount + v.Report.TotalDemand
		rowNo++

	}
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "B", rowNo), fmt.Sprintf("%v%v", "B", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "Total")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf(" %.2f", totalAmount))

	return excel, nil

}

// FilterWardYearWisePropertyDemandReportPDF : ""
func (s *Service) FilterWardYearWisePropertyDemandReportPDF(ctx *models.Context, filter *models.WardYearWisePropertyDemandReportFilter, pagination *models.Pagination) ([]byte, error) {
	data, err := s.FilterWardYearWisePropertyDemandReport(ctx, filter, pagination)
	if err != nil {
		return nil, err
	}
	var totalCollection float64
	for _, v := range data {
		totalCollection = totalCollection + v.Report.TotalDemand
	}
	m := make(map[string]interface{})
	m2 := make(map[string]interface{})
	m3 := make(map[string]interface{})
	m["demand"] = data
	m["totalCollection"] = totalCollection
	m2["currentDate"] = time.Now()
	m3["inc"] = func(a int) int {
		return a + 1
	}
	var pdfdata models.PDFData
	pdfdata.Data = m
	pdfdata.RefData = m2
	productConfigUniqueID := "1"
	productConfig, err := s.Daos.GetSingleProductConfiguration(ctx, productConfigUniqueID)
	if err != nil {
		return nil, errors.New("Error in getting product config" + err.Error())
	}
	pdfdata.Config = productConfig.ProductConfiguration

	r := NewRequestPdf("")
	templatePathStart := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.TEMPLATELOC)
	//html template path
	templatePath := templatePathStart + "property_wardwiseyear_demand_report.html"
	if err := r.ParseTemplate(templatePath, pdfdata); err == nil {
		fmt.Println("start pdf generated successfully")

		ok, data, err := r.GeneratePDFAsFile()

		fmt.Println(ok, "pdf generated successfully")
		return data, err
	} else {
		fmt.Println("Error in parcing template - " + err.Error())

		return nil, errors.New("Error in parcing template - " + err.Error())
	}

}

// FilterTeamYearWisePropertyCollectionReport : ""
func (s *Service) FilterTeamYearWisePropertyCollectionReport(ctx *models.Context, filter *models.TeamYearWisePropertyCollectionReportFilter, pagination *models.Pagination) ([]models.TeamYearWisePropertyCollectionReport, error) {
	return s.Daos.FilterTeamYearWisePropertyCollectionReport(ctx, filter, pagination)
}

// FilterTeamYearWisePropertyCollectionReportExcel: ""
func (s *Service) FilterTeamYearWisePropertyCollectionReportExcel(ctx *models.Context, filter *models.TeamYearWisePropertyCollectionReportFilter, pagination *models.Pagination) (*excelize.File, error) {
	data, err := s.FilterTeamYearWisePropertyCollectionReport(ctx, filter, pagination)
	if err != nil {
		return nil, err
	}

	resPD, err1 := s.Daos.GetSingleDefaultProductConfiguration(ctx)
	if err1 != nil {
		return nil, err1
	}

	excel := excelize.NewFile()
	sheet1 := "Property Team Wise Collection"
	index := excel.NewSheet(sheet1)
	rowNo := 1

	excel.SetActiveSheet(index)
	if resPD.LocationID == "Bhagalpur" {
		excel.MergeCell(sheet1, "A1", "F3")
		excel.MergeCell(sheet1, "A4", "F5")
	} else {
		excel.MergeCell(sheet1, "A1", "B5")
		excel.MergeCell(sheet1, "C1", "F3")
		excel.MergeCell(sheet1, "C4", "F5")
	}

	excel.MergeCell(sheet1, "A6", "F6")
	style1, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#FFDC6D"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	style3, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#B6DDE8"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}

	if resPD.LocationID != "Bhagalpur" {
		documentUrl := s.Shared.Config.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.FILEURL)
		if err := excel.AddPicture(sheet1, fmt.Sprintf("%v%v", "A", rowNo), documentUrl+"municipal/logo.png", `{"x_scale": 0.115, "y_scale": 0.0935}`); err != nil {
			fmt.Println(err)
		}
	}

	if resPD.LocationID == "Bhagalpur" {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), ctx.ProductConfig.Name)
	} else {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), ctx.ProductConfig.Name)
	}

	rowNo++
	rowNo++
	rowNo++

	if resPD.LocationID == "Bhagalpur" {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), sheet1)
	} else {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), sheet1)
	}
	rowNo++
	rowNo++

	t := time.Now()
	toDate := t.Format("02-January-2006")
	reportFromMsg := "Report Generated on" + " " + toDate
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style3)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), reportFromMsg)
	rowNo++

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "F", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "S.No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "User")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "UserType")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), "Paied Properties")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), "No of Payments")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), "Collection Amount")
	rowNo++
	var totalAmount, totalProperties, totalPayments float64
	for i, v := range data {
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), i+1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v.User.Name)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), v.User.Type)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), v.Report.TotalNoProperties)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), v.Report.TotalNoPayments)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), v.Report.TotalCollections)
		totalAmount = totalAmount + v.Report.TotalCollections
		totalProperties = totalProperties + v.Report.TotalNoProperties
		totalPayments = totalPayments + v.Report.TotalNoPayments
		rowNo++

	}
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "B", rowNo), fmt.Sprintf("%v%v", "B", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "D", rowNo), fmt.Sprintf("%v%v", "D", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "E", rowNo), fmt.Sprintf("%v%v", "E", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "F", rowNo), fmt.Sprintf("%v%v", "F", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "Total")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), fmt.Sprintf(" %.2f", totalProperties))
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), fmt.Sprintf(" %.2f", totalPayments))
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), fmt.Sprintf(" %.2f", totalAmount))

	return excel, nil

}

// FilterTeamYearWisePropertyCollectionReportPDF : ""
func (s *Service) FilterTeamYearWisePropertyCollectionReportPDF(ctx *models.Context, filter *models.TeamYearWisePropertyCollectionReportFilter, pagination *models.Pagination) ([]byte, error) {
	data, err := s.FilterTeamYearWisePropertyCollectionReport(ctx, filter, pagination)
	if err != nil {
		return nil, err
	}
	var totalCollection float64
	for _, v := range data {
		totalCollection = totalCollection + v.Report.TotalCollections
	}
	m := make(map[string]interface{})
	m2 := make(map[string]interface{})
	m3 := make(map[string]interface{})
	m["demand"] = data
	m["totalCollection"] = totalCollection
	m2["currentDate"] = time.Now()
	m3["inc"] = func(a int) int {
		return a + 1
	}
	var pdfdata models.PDFData
	pdfdata.Data = m
	pdfdata.RefData = m2
	productConfigUniqueID := "1"
	productConfig, err := s.Daos.GetSingleProductConfiguration(ctx, productConfigUniqueID)
	if err != nil {
		return nil, errors.New("Error in getting product config" + err.Error())
	}
	pdfdata.Config = productConfig.ProductConfiguration

	r := NewRequestPdf("")
	templatePathStart := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.TEMPLATELOC)
	//html template path
	templatePath := templatePathStart + "property_teamwiseyear_collection_report.html"
	if err := r.ParseTemplate(templatePath, pdfdata); err == nil {
		fmt.Println("start pdf generated successfully")

		ok, data, err := r.GeneratePDFAsFile()

		fmt.Println(ok, "pdf generated successfully")
		return data, err
	} else {
		fmt.Println("Error in parcing template - " + err.Error())

		return nil, errors.New("Error in parcing template - " + err.Error())
	}

}

// FilterYearWisePropertyDemandReport : ""
func (s *Service) FilterYearWisePropertyDemandReport(ctx *models.Context, filter *models.YearWisePropertyDemandReportFilter, pagination *models.Pagination) ([]models.YearWisePropertyDemandReport, error) {
	return s.Daos.FilterYearWisePropertyDemandReport(ctx, filter, pagination)
}

// FilterYearWisePropertyDemandReportExcel: ""
func (s *Service) FilterYearWisePropertyDemandReportExcel(ctx *models.Context, filter *models.YearWisePropertyDemandReportFilter, pagination *models.Pagination) (*excelize.File, error) {
	data, err := s.FilterYearWisePropertyDemandReport(ctx, filter, pagination)
	if err != nil {
		return nil, err
	}

	resPD, err1 := s.Daos.GetSingleDefaultProductConfiguration(ctx)
	if err1 != nil {
		return nil, err1
	}

	excel := excelize.NewFile()
	sheet1 := "Property Year Wise Demand"
	index := excel.NewSheet(sheet1)
	rowNo := 1

	excel.SetActiveSheet(index)
	if resPD.LocationID == "Bhagalpur" {
		excel.MergeCell(sheet1, "A1", "D3")
		excel.MergeCell(sheet1, "A4", "D5")
	} else {
		excel.MergeCell(sheet1, "A1", "B5")
		excel.MergeCell(sheet1, "C1", "D3")
		excel.MergeCell(sheet1, "C4", "D5")
	}
	excel.MergeCell(sheet1, "A6", "D6")

	style1, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#FFDC6D"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	style3, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#B6DDE8"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}

	if resPD.LocationID != "Bhagalpur" {
		documentUrl := s.Shared.Config.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.FILEURL)
		if err := excel.AddPicture(sheet1, fmt.Sprintf("%v%v", "A", rowNo), documentUrl+"municipal/logo.png", `{"x_scale": 0.115, "y_scale": 0.0935}`); err != nil {
			fmt.Println(err)
		}
	}

	if resPD.LocationID == "Bhagalpur" {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), ctx.ProductConfig.Name)
	} else {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), ctx.ProductConfig.Name)
	}

	rowNo++
	rowNo++
	rowNo++

	if resPD.LocationID == "Bhagalpur" {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), sheet1)
	} else {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), sheet1)
	}
	rowNo++
	rowNo++

	t := time.Now()
	toDate := t.Format("02-January-2006")
	reportFromMsg := "Report Generated on" + " " + toDate
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style3)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), reportFromMsg)
	rowNo++

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "D", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "S.No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "Month")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "No of Properties")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), "Demand")
	rowNo++
	var totalAmount float64
	for i, v := range data {
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), i+1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v.Month)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), v.NoOfProperties)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), v.Demand)
		totalAmount = totalAmount + v.Demand
		rowNo++

	}
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "B", rowNo), fmt.Sprintf("%v%v", "B", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "D", rowNo), fmt.Sprintf("%v%v", "D", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "Total")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), fmt.Sprintf(" %.2f", totalAmount))

	return excel, nil

}

// /DayWisePropertyDemandReportPDF : ""
// func (s *Service) FilterYearWisePropertyDemandReportPDF(ctx *models.Context, filter *models.YearWisePropertyDemandReportFilter, pagination *models.Pagination) ([]byte, error) {
// 	properties, err := s.FilterYearWisePropertyDemandReport(ctx, filter, pagination)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if properties == nil {
// 		properties = new(models.YearWisePropertyDemandReportFilter)
// 	}
// 	m := make(map[string]interface{})
// 	m2 := make(map[string]interface{})
// 	m["demand"] = properties
// 	m2["currentDate"] = time.Now()
// 	cfy, err := s.Daos.GetCurrentFinancialYear(ctx)
// 	if err != nil {
// 		return nil, errors.New("Error in getting current financial year " + err.Error())
// 	}
// 	m2["currentFy"] = cfy
// 	var pdfdata models.PDFData
// 	pdfdata.Data = m
// 	pdfdata.RefData = m2
// 	productConfigUniqueID := "1"
// 	productConfig, err := s.Daos.GetSingleProductConfiguration(ctx, productConfigUniqueID)
// 	if err != nil {
// 		return nil, errors.New("Error in getting product config" + err.Error())
// 	}
// 	pdfdata.Config = productConfig.ProductConfiguration

// 	r := NewRequestPdf("")
// 	templatePathStart := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.TEMPLATELOC)
// 	//html template path
// 	templatePath := templatePathStart + "property_daywise_demand_report.html"
// 	if err := r.ParseTemplate(templatePath, pdfdata); err == nil {
// 		fmt.Println("start pdf generated successfully")

// 		ok, data, err := r.GeneratePDFAsFile()

// 		fmt.Println(ok, "pdf generated successfully")
// 		return data, err
// 	} else {
// 		fmt.Println("Error in parcing template - " + err.Error())

// 		return nil, errors.New("Error in parcing template - " + err.Error())
// 	}
// 	return nil, nil

// }

// FilterYearWisePropertyCollectionReport : ""
func (s *Service) FilterYearWisePropertyCollectionReport(ctx *models.Context, filter *models.YearWisePropertyCollectionReportFilter) ([]models.YearWisePropertyCollectionReport, error) {
	return s.Daos.FilterYearWisePropertyCollectionReport(ctx, filter)
}

// FilterYearWisePropertyDemandReportExcel: ""
func (s *Service) FilterYearWisePropertyCollectionReportExcel(ctx *models.Context, filter *models.YearWisePropertyCollectionReportFilter) (*excelize.File, error) {
	data, err := s.FilterYearWisePropertyCollectionReport(ctx, filter)
	if err != nil {
		return nil, err
	}

	resPD, err1 := s.Daos.GetSingleDefaultProductConfiguration(ctx)
	if err1 != nil {
		return nil, err1
	}

	excel := excelize.NewFile()
	sheet1 := "Property Year Wise Collection"
	index := excel.NewSheet(sheet1)
	rowNo := 1

	excel.SetActiveSheet(index)
	if resPD.LocationID == "Bhagalpur" {
		excel.MergeCell(sheet1, "A1", "D3")
		excel.MergeCell(sheet1, "A4", "D5")
	} else {
		excel.MergeCell(sheet1, "A1", "B5")
		excel.MergeCell(sheet1, "C1", "D3")
		excel.MergeCell(sheet1, "C4", "D5")
	}
	excel.MergeCell(sheet1, "A6", "D6")
	style1, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#FFDC6D"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	style3, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#B6DDE8"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}

	if resPD.LocationID != "Bhagalpur" {
		documentUrl := s.Shared.Config.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.FILEURL)
		if err := excel.AddPicture(sheet1, fmt.Sprintf("%v%v", "A", rowNo), documentUrl+"municipal/logo.png", `{"x_scale": 0.115, "y_scale": 0.0935}`); err != nil {
			fmt.Println(err)
		}
	}

	if resPD.LocationID == "Bhagalpur" {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), ctx.ProductConfig.Name)
	} else {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), ctx.ProductConfig.Name)
	}

	rowNo++
	rowNo++
	rowNo++

	if resPD.LocationID == "Bhagalpur" {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), sheet1)
	} else {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), sheet1)
	}
	rowNo++
	rowNo++

	t := time.Now()
	toDate := t.Format("02-January-2006")
	reportFromMsg := "Report Generated on" + " " + toDate
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style3)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), reportFromMsg)
	rowNo++

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "D", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "S.No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "Month")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "No of Properties")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), "Collection")
	rowNo++
	var totalAmount float64
	for i, v := range data {
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), i+1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v.Month)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), v.PropertyCount)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), v.TotalTax)
		totalAmount = totalAmount + v.TotalTax
		rowNo++

	}
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "B", rowNo), fmt.Sprintf("%v%v", "B", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "D", rowNo), fmt.Sprintf("%v%v", "D", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "Total")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), fmt.Sprintf(" %.2f", totalAmount))

	return excel, nil

}

// FilterYearWisePropertyDemandReportExcel: ""
func (s *Service) FilterWardWisePropertyDemandAndCollectionReportExcel(ctx *models.Context, filter *models.YearWisePropertyCollectionReportFilter) (*excelize.File, error) {
	data, err := s.FilterWardWisePropertyDemandAndCollectionReport(ctx, filter)
	if err != nil {
		return nil, err
	}

	resPD, err1 := s.Daos.GetSingleDefaultProductConfiguration(ctx)
	if err1 != nil {
		return nil, err1
	}

	excel := excelize.NewFile()
	sheet1 := "Property Ward Wise Demand And Collection"
	index := excel.NewSheet(sheet1)
	rowNo := 1

	excel.SetActiveSheet(index)
	if resPD.LocationID == "Bhagalpur" {
		excel.MergeCell(sheet1, "A1", "M3")
		excel.MergeCell(sheet1, "A4", "M5")
	} else {
		excel.MergeCell(sheet1, "A1", "B5")
		excel.MergeCell(sheet1, "C1", "M3")
		excel.MergeCell(sheet1, "C4", "M5")
	}

	style1, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#FFDC6D"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	style3, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#B6DDE8"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}

	if resPD.LocationID != "Bhagalpur" {
		documentUrl := s.Shared.Config.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.FILEURL)
		if err := excel.AddPicture(sheet1, fmt.Sprintf("%v%v", "A", rowNo), documentUrl+"municipal/logo.png", `{"x_scale": 0.115, "y_scale": 0.0935}`); err != nil {
			fmt.Println(err)
		}
	}

	if resPD.LocationID == "Bhagalpur" {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), ctx.ProductConfig.Name)
	} else {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), ctx.ProductConfig.Name)
	}

	rowNo++
	rowNo++
	rowNo++

	if resPD.LocationID == "Bhagalpur" {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), sheet1)
	} else {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), sheet1)
	}

	rowNo++
	rowNo++
	excel.MergeCell(sheet1, "B6", "D6")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "Property Details")
	excel.MergeCell(sheet1, "G6", "K6")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "H", rowNo), "Collection During the period")
	rowNo++
	//excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "M", rowNo), style1)
	excel.MergeCell(sheet1, "A6", "A7")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "Ward No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "Total")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Paid")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), "Unpaid")
	excel.MergeCell(sheet1, "E6", "E7")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), "OutStanding Demand At the beginning of the period")
	excel.MergeCell(sheet1, "F6", "F7")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), "Demand of the Current Period")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), "Arrear")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "H", rowNo), "Current")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "I", rowNo), "Penalty")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "J", rowNo), "Rebate")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "K", rowNo), "Advance")
	excel.MergeCell(sheet1, "L6", "L7")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "L", rowNo), "TotalCollection")
	excel.MergeCell(sheet1, "M6", "M7")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "M", rowNo), "TotalOutstandingDemand")
	rowNo++
	var totalAmount float64
	for i, v := range data.Report {
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), i+1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v.TotalAmount)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), v.PaidAmount)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), v.TotalAmount-v.PaidAmount)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), v.OutstandingDemand)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), v.CurrentDemand)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), v.ArrearCollection)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "H", rowNo), v.CurrentCollection)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "I", rowNo), v.Penalty)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "J", rowNo), v.Rebate)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "K", rowNo), v.AdvanceAmount)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "L", rowNo), v.TotalCollection)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "M", rowNo), v.TotalOutstandingDemand)
		totalAmount = totalAmount + v.TotalOutstandingDemand
		rowNo++

	}
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "B", rowNo), fmt.Sprintf("%v%v", "B", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "D", rowNo), fmt.Sprintf("%v%v", "D", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "Total")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), fmt.Sprintf(" %.2f", totalAmount))

	return excel, nil

}

func (s *Service) FilterWardWisePropertyDemandAndCollectionReport(ctx *models.Context, filter *models.YearWisePropertyCollectionReportFilter) (models.WardWisePropertyDemandAndCollectionReport, error) {
	return s.Daos.FilterWardWisePropertyDemandAndCollectionReport(ctx, filter)
}

// UserWisePropertyCollectionReport : ""
func (s *Service) UserWisePropertyCollectionReport(ctx *models.Context, filter *models.UserWisePropertyCollectionReportFilter) ([]models.UserWisePropertyCollectionReport, error) {
	return s.Daos.UserWisePropertyCollectionReport(ctx, filter)
}

// UserWisePropertyCollectionReportExcel : ""
func (s *Service) UserWisePropertyCollectionReportExcel(ctx *models.Context, filter *models.UserWisePropertyCollectionReportFilter) (*excelize.File, error) {
	data, err := s.UserWisePropertyCollectionReport(ctx, filter)
	if err != nil {
		return nil, err
	}

	resPD, err1 := s.Daos.GetSingleDefaultProductConfiguration(ctx)
	if err1 != nil {
		return nil, err1
	}

	excel := excelize.NewFile()
	sheet1 := "User Wise Property Collection Report"
	rowNo := 1
	index := excel.NewSheet(sheet1)
	excel.SetActiveSheet(index)
	if resPD.LocationID == "Bhagalpur" {
		excel.MergeCell(sheet1, "A1", "AF3")
		excel.MergeCell(sheet1, "A4", "AF5")
	} else {
		excel.MergeCell(sheet1, "A1", "B5")
		excel.MergeCell(sheet1, "C1", "AF3")
		excel.MergeCell(sheet1, "C4", "AF5")
	}

	excel.MergeCell(sheet1, "A6", "AF6")
	style3, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#B6DDE8"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	style1, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#FFDC6D"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	// style4, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#FFDC6D"],"pattern":1},"alignment":{"horizontal":"right","vertical":"center"},"font":{"bold":true}}`)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// style2, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#B6DDE8"],"pattern":1},"font":{"bold":true}}`)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style3)

	if resPD.LocationID != "Bhagalpur" {
		documentUrl := s.Shared.Config.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.FILEURL)
		if err := excel.AddPicture(sheet1, fmt.Sprintf("%v%v", "A", rowNo), documentUrl+"municipal/logo.png", `{"x_scale": 0.115, "y_scale": 0.0935}`); err != nil {
			fmt.Println(err)
		}
	}

	if resPD.LocationID == "Bhagalpur" {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), ctx.ProductConfig.Name)
	} else {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), ctx.ProductConfig.Name)
	}
	rowNo++
	rowNo++
	rowNo++

	if resPD.LocationID == "Bhagalpur" {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "User Wise Property Collection")
	} else {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "User Wise Property Collection")
	}
	rowNo++
	rowNo++
	t := time.Now()
	toDate := t.Format("02-January-2006")

	reportFromMsg := "Report Generated on" + " " + toDate

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style3)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), reportFromMsg)
	rowNo++

	if len(data) > 0 {
		columnNo1 := 'A'
		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", columnNo1, rowNo), "S.No")
		columnNo1++
		excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", columnNo1, rowNo), "User Name")
		columnNo1++
		var ch1 rune
		var ch2 rune
		ch1 = 'A'
		ch2 = 'A'
		column := ""
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "AF", rowNo), style1)
		for _, v := range data[0].Payments {

			if columnNo1 <= 'Z' {
				column = fmt.Sprintf("%c", columnNo1)
				day := fmt.Sprintf("%v", v.ID)
				excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", column, rowNo), day)
			} else {
				CH1 := fmt.Sprintf("%c", ch1)
				CH2 := fmt.Sprintf("%c", ch2)
				fmt.Println("ch1===>", CH1)
				fmt.Println("ch2===>", CH2)
				column = fmt.Sprintf("%c%c", ch1, ch2)
				day := fmt.Sprintf("%v", v.ID)
				excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", column, rowNo), day)
			}
			if columnNo1 > 'Z' {
				ch2++
			}
			columnNo1++
		}
		rowNo++
		for i, v2 := range data {
			columnNo2 := 'A'
			excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", columnNo2, rowNo), i+1)
			columnNo2++

			excel.SetCellValue(sheet1, fmt.Sprintf("%c%v", columnNo2, rowNo), v2.Name)
			columnNo2++

			ch3 := 'A'
			ch4 := 'A'
			column2 := ""
			for _, v3 := range v2.Payments {
				if columnNo2 <= 'Z' {
					column2 = fmt.Sprintf("%c", columnNo2)
					collection := fmt.Sprintf("%v", v3.TotalCollection)
					excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", column2, rowNo), collection)
				} else {
					column2 = fmt.Sprintf("%c%c", ch3, ch4)
					collection := fmt.Sprintf("%v", v3.TotalCollection)
					excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", column2, rowNo), collection)

				}
				if columnNo2 > 'Z' {
					ch4++
				}
				columnNo2++

			}
			rowNo++
		}
	}
	return excel, nil
}

// GetBasicPropertyDetailsPDF : ""
func (s *Service) GetBasicPropertyDetailsPDF(ctx *models.Context, PropertyID string) ([]byte, error) {
	data, err := s.GetSingleProperty(ctx, PropertyID)
	if err != nil {
		return nil, err
	}

	m := make(map[string]interface{})
	m2 := make(map[string]interface{})
	m["demand"] = data
	m2["currentDate"] = time.Now()
	m2["IncOwner"] = func(a int) int {
		return a + 1
	}
	m2["IncFloor"] = func(a int) int {
		return a + 1
	}

	var pdfdata models.PDFData
	pdfdata.Data = m
	pdfdata.RefData = m2
	productConfigUniqueID := "1"
	productConfig, err := s.Daos.GetSingleProductConfiguration(ctx, productConfigUniqueID)
	if err != nil {
		return nil, errors.New("Error in getting product config" + err.Error())
	}
	pdfdata.Config = productConfig.ProductConfiguration

	r := NewRequestPdf("")
	templatePathStart := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.TEMPLATELOC)
	//html template path
	templatePath := templatePathStart + "propertydetails.html"
	if err := r.ParseTemplate(templatePath, pdfdata); err == nil {
		fmt.Println("start pdf generated successfully")

		ok, data, err := r.GeneratePDFAsFile()

		fmt.Println(ok, "pdf generated successfully")
		return data, err
	} else {
		fmt.Println("Error in parcing template - " + err.Error())

		return nil, errors.New("Error in parcing template - " + err.Error())
	}

}

// FilterPropertyArrearAndCurrentCollectionReportJSON : ""
func (s *Service) FilterPropertyArrearAndCurrentCollectionReportJSON(ctx *models.Context, filter *models.PropertyArrearAndCurrentCollectionFilter) ([]models.PropertyArrearAndCurrentCollectionReport, error) {

	defer ctx.Session.EndSession(ctx.CTX)
	return s.Daos.FilterPropertyArrearAndCurrentCollectionReportJSON(ctx, filter)
}

// FilterPropertyArrearAndCurrentCollectionReportExcel : ""
func (s *Service) FilterPropertyArrearAndCurrentCollectionReportExcel(ctx *models.Context, filter *models.PropertyArrearAndCurrentCollectionFilter) (*excelize.File, error) {
	res, err := s.FilterPropertyArrearAndCurrentCollectionReportJSON(ctx, filter)
	if err != nil {
		return nil, err
	}

	resPD, err1 := s.Daos.GetSingleDefaultProductConfiguration(ctx)
	if err1 != nil {
		return nil, err1
	}

	excel := excelize.NewFile()
	sheet1 := "Holding Tax Arrear and Current Collection Report"
	index := excel.NewSheet(sheet1)
	rowNo := 1

	excel.SetActiveSheet(index)
	if resPD.LocationID == "Bhagalpur" {
		excel.MergeCell(sheet1, "A1", "J3")
		excel.MergeCell(sheet1, "A4", "J5")
	} else {
		excel.MergeCell(sheet1, "A1", "B5")
		excel.MergeCell(sheet1, "C1", "J3")
		excel.MergeCell(sheet1, "C4", "J5")
	}
	excel.MergeCell(sheet1, "A6", "J6")
	excel.MergeCell(sheet1, "A7", "J7")
	excel.MergeCell(sheet1, "A8", "A9")
	excel.MergeCell(sheet1, "B8", "B9")
	excel.MergeCell(sheet1, "C8", "D8")
	excel.MergeCell(sheet1, "E8", "F8")
	excel.MergeCell(sheet1, "G8", "G9")
	excel.MergeCell(sheet1, "H8", "H9")
	excel.MergeCell(sheet1, "I8", "I9")
	excel.MergeCell(sheet1, "J8", "J9")
	style1, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#FFDC6D"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	style3, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#B6DDE8"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}

	if resPD.LocationID != "Bhagalpur" {
		documentUrl := s.Shared.Config.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.FILEURL)
		if err := excel.AddPicture(sheet1, fmt.Sprintf("%v%v", "A", rowNo), documentUrl+"municipal/logo.png", `{"x_scale": 0.115, "y_scale": 0.0935}`); err != nil {
			fmt.Println(err)
		}
	}

	if resPD.LocationID == "Bhagalpur" {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), ctx.ProductConfig.Name)
	} else {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), ctx.ProductConfig.Name)
	}
	rowNo++
	rowNo++
	rowNo++

	if resPD.LocationID == "Bhagalpur" {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), sheet1)
	} else {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), sheet1)
	}
	rowNo++
	rowNo++

	t := time.Now()
	toDate := t.Format("02-January-2006")

	reportFromMsg2 := "Report"
	if filter != nil {
		fmt.Println(filter.DateRange.From, filter.DateRange.To)
		if filter.DateRange.From != nil && filter.DateRange.To == nil {
			reportFromMsg2 = reportFromMsg2 + " on " + fmt.Sprintf("%v-%v-%v", filter.DateRange.From.Day(), filter.DateRange.From.Month(), filter.DateRange.From.Year())
		}
		if filter.DateRange.From != nil && filter.DateRange.To != nil {
			reportFromMsg2 = reportFromMsg2 + " From " + fmt.Sprintf("%v-%v-%v", filter.DateRange.From.Day(), filter.DateRange.From.Month(), filter.DateRange.From.Year()) + " To " + fmt.Sprintf("%v-%v-%v", filter.DateRange.To.Day(), filter.DateRange.To.Month(), filter.DateRange.To.Year())
		}
		if filter.DateRange.From == nil && filter.DateRange.To == nil {
			fmt.Println("everything is nil")
		}

	}
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style3)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), reportFromMsg2)
	rowNo++

	reportFromMsg := "This Report is Generated on" + " " + toDate
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style3)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), reportFromMsg)
	rowNo++

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "J", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v", "Ward No"))
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), fmt.Sprintf("%v", "No Of Holding Tax"))
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), fmt.Sprintf("%v", "Holding Tax"))
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), fmt.Sprintf("%v", "Penalty"))
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), fmt.Sprintf("%v", "Rebate"))
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "H", rowNo), fmt.Sprintf("%v", "Form Fees"))
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "I", rowNo), fmt.Sprintf("%v", "Other Demand"))
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "J", rowNo), fmt.Sprintf("%v", "Total"))
	rowNo++

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "I", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v", "Current"))
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), fmt.Sprintf("%v", "Arrears"))
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), fmt.Sprintf("%v", "Current"))
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), fmt.Sprintf("%v", "Arrears"))
	rowNo++
	var totalAmount float64

	for _, v := range res {
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), v.Ward.Name)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v.Payments.TotalProperties)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), v.Payments.CurrentTax-v.Payments.CurrentAlreadyPaid)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), v.Payments.ArrearTax-v.Payments.ArrearAlreadyPaid)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), v.Payments.CurrentPenalty)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), v.Payments.ArrearPenalty)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), v.Payments.CurrentRebate+v.Payments.ArrearRebate)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "H", rowNo), v.Payments.FormFee)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "I", rowNo), v.Payments.OtherDemand)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "J", rowNo), v.Payments.TotalTax+v.Payments.FormFee)
		totalAmount = totalAmount + (v.Payments.TotalTax + v.Payments.FormFee)
		rowNo++

	}
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "B", rowNo), fmt.Sprintf("%v%v", "B", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "D", rowNo), fmt.Sprintf("%v%v", "D", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "E", rowNo), fmt.Sprintf("%v%v", "E", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "F", rowNo), fmt.Sprintf("%v%v", "F", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "G", rowNo), fmt.Sprintf("%v%v", "G", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "H", rowNo), fmt.Sprintf("%v%v", "H", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "I", rowNo), fmt.Sprintf("%v%v", "I", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "J", rowNo), fmt.Sprintf("%v%v", "I", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "Total")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "J", rowNo), fmt.Sprintf(" %.2f", totalAmount))

	return excel, nil

}

// FilterCounterReportV2JSON : ""
func (s *Service) FilterCounterReportV2JSON(ctx *models.Context, filter *models.PropertyPaymentFilter, pagination *models.Pagination) ([]models.RefCounterReportV2, error) {
	defer ctx.Session.EndSession(ctx.CTX)
	return s.Daos.FilterCounterReportV2JSON(ctx, filter, pagination)
}

// CounterReportV2 : ""
func (s *Service) FilterCounterReportV2Excel(ctx *models.Context, filter *models.PropertyPaymentFilter, pagination *models.Pagination) (*excelize.File, error) {
	res, err := s.FilterCounterReportV2JSON(ctx, filter, pagination)
	if err != nil {
		return nil, err
	}
	fmt.Println("'res length==>'", len(res))

	resPD, err1 := s.Daos.GetSingleDefaultProductConfiguration(ctx)
	if err1 != nil {
		return nil, err1
	}

	//  create an excel file
	excel := excelize.NewFile()
	sheet1 := "Counter Report"
	rowNo := 1
	index := excel.NewSheet(sheet1)
	excel.SetActiveSheet(index)
	if resPD.LocationID == "Bhagalpur" {
		excel.MergeCell(sheet1, "A1", "AB3")
		excel.MergeCell(sheet1, "A4", "AB5")
	} else {
		excel.MergeCell(sheet1, "A1", "B5")
		excel.MergeCell(sheet1, "C1", "AB3")
		excel.MergeCell(sheet1, "C4", "AB5")
	}
	excel.MergeCell(sheet1, "A6", "AB6")
	excel.MergeCell(sheet1, "A7", "AB7")
	excel.MergeCell(sheet1, "A8", "AB8")

	style3, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#B6DDE8"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	style1, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#FFDC6D"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style3)

	if resPD.LocationID != "Bhagalpur" {
		documentUrl := s.Shared.Config.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.FILEURL)
		if err := excel.AddPicture(sheet1, fmt.Sprintf("%v%v", "A", rowNo), documentUrl+"municipal/logo.png", `{"x_scale": 0.115, "y_scale": 0.0935}`); err != nil {
			fmt.Println(err)
		}
	}

	if resPD.LocationID == "Bhagalpur" {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), ctx.ProductConfig.Name)
	} else {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), ctx.ProductConfig.Name)
	}

	rowNo++
	rowNo++
	rowNo++

	if resPD.LocationID == "Bhagalpur" {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "Counter Report")
	} else {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Counter Report")
	}
	rowNo++
	rowNo++

	reportFromMsg := "Counter Collection Report"
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style3)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), reportFromMsg)
	rowNo++

	reportFromMsg2 := "Report"
	// if filter != nil {
	// 	fmt.Println(filter.DateRange.From, filter.DateRange.To)
	// 	if filter.DateRange.From != nil && filter.DateRange.To == nil {
	// 		reportFromMsg2 = reportFromMsg2 + " on " + fmt.Sprintf("%v-%v-%v", filter.DateRange.From.Day(), filter.DateRange.From.Month(), filter.DateRange.From.Year())
	// 	}
	// 	if filter.DateRange.From != nil && filter.DateRange.To != nil {
	// 		reportFromMsg2 = reportFromMsg2 + " From " + fmt.Sprintf("%v-%v-%v", filter.DateRange.From.Day(), filter.DateRange.From.Month(), filter.DateRange.From.Year()) + " To " + fmt.Sprintf("%v-%v-%v", filter.DateRange.To.Day(), filter.DateRange.To.Month(), filter.DateRange.To.Year())
	// 	}
	// 	if filter.DateRange.From == nil && filter.DateRange.To == nil {
	// 		fmt.Println("everything is nil")
	// 	}

	// }
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style3)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), reportFromMsg2)
	rowNo++

	t := time.Now()
	toDate := t.Format("02-January-2006")
	reportFromMsg3 := "Report Generated on" + " " + toDate
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style3)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), reportFromMsg3)
	rowNo++

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "AB", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "S.No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "Holding No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Application No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), "Old Holding No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), "Ward")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), "Owner Name")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), "Address")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "H", rowNo), "Mobile")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "I", rowNo), "From Year")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "J", rowNo), "Upto Year")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "K", rowNo), "Txn Date")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "L", rowNo), "Receipt No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "M", rowNo), "Payment Mode")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "N", rowNo), "Cheque / DD No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "O", rowNo), "Bank")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "P", rowNo), "Branch Name")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "Q", rowNo), "Received Amount")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "R", rowNo), "Current Holding Tax")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "S", rowNo), "Arrear Holding Tax")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "T", rowNo), "Current Penalty")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "U", rowNo), "Arrear Penalty")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "V", rowNo), "From Fee")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "W", rowNo), "Discount")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "X", rowNo), "Bounced")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "Y", rowNo), "Total Amount")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "Z", rowNo), "Assessment By")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "AA", rowNo), "Activated By")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "AB", rowNo), "Collected By")
	rowNo++

	var totalAmount float64

	for k, v := range res {
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), k+1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v.Basic.Property.UniqueID)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), v.Basic.Property.ApplicationNo)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), func() string {
			if v.Basic.Property.OldHoldingNumber != "" {
				return v.Basic.Property.OldHoldingNumber
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), v.Ref.Ward.Name)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), v.Ref.Owner.Name)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), v.Address.AL1+v.Address.Al2)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "H", rowNo), v.Ref.Owner.Mobile)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "I", rowNo), func() string {
			if len(v.PaymentFYs.FYs) > 0 {
				return v.PaymentFYs.FYs[0].FY.Name
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "J", rowNo), func() string {
			if len(v.PaymentFYs.FYs) > 0 {
				if len(v.PaymentFYs.FYs) == 1 {
					return v.PaymentFYs.FYs[0].FY.Name
				}
				if len(v.PaymentFYs.FYs) > 1 {
					return v.PaymentFYs.FYs[len(v.PaymentFYs.FYs)-1].FY.Name

				}
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "K", rowNo), v.PropertyPayment.CompletionDate.Format("2006-01-02"))
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "L", rowNo), v.PropertyPayment.ReciptNo)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "M", rowNo), func() string {
			if v.PropertyPayment.Details != nil {
				switch v.PropertyPayment.Details.MOP.Mode {
				case constants.MOPCASH:
					return "Cash"
				case constants.MOPCHEQUE:
					return "Cheque"
				case constants.MOPDD:
					return "DD"
				case constants.MOPNETBANKING:
					return "Online"
				default:
					return "Invalid"
				}
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "N", rowNo), func() string {
			if v.PropertyPayment.Details != nil {
				if v.PropertyPayment.Details.MOP.Mode == constants.MOPCHEQUE {
					if v.PropertyPayment.Details.MOP.Cheque != nil {
						return v.PropertyPayment.Details.MOP.Cheque.No
					}
				} else if v.PropertyPayment.Details.MOP.Mode == constants.MOPDD {
					if v.PropertyPayment.Details.MOP.DD != nil {
						return v.PropertyPayment.Details.MOP.DD.No
					}
				} else {
					return "NA"
				}
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "O", rowNo), func() string {
			if v.PropertyPayment.Details != nil {
				if v.PropertyPayment.Details.MOP.Mode == constants.MOPCHEQUE {
					if v.PropertyPayment.Details.MOP.Cheque != nil {
						return v.PropertyPayment.Details.MOP.Cheque.Bank
					}
				} else if v.PropertyPayment.Details.MOP.Mode == constants.MOPDD {
					if v.PropertyPayment.Details.MOP.DD != nil {
						return v.PropertyPayment.Details.MOP.DD.Bank
					}
				} else {
					return "NA"
				}
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "P", rowNo), func() string {
			if v.PropertyPayment.Details != nil {
				if v.PropertyPayment.Details.MOP.Mode == constants.MOPCHEQUE {
					if v.PropertyPayment.Details.MOP.Cheque != nil {
						return v.PropertyPayment.Details.MOP.Cheque.Branch
					}
				} else if v.PropertyPayment.Details.MOP.Mode == constants.MOPDD {
					if v.PropertyPayment.Details.MOP.DD != nil {
						return v.PropertyPayment.Details.MOP.DD.Branch
					}
				} else {
					return "NA"
				}
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "Q", rowNo), v.PaymentFYs.TotalTax+func() float64 {
			if v.PropertyPayment.Demand != nil {
				return v.PropertyPayment.Demand.FormFee
			}
			return 0
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "R", rowNo), v.PaymentFYs.CurrentTax-v.PaymentFYs.CurrentAlreadyPaid)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "S", rowNo), v.PaymentFYs.ArrearTax-v.PaymentFYs.ArrearAlreadyPaid)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "T", rowNo), v.PaymentFYs.CurrentPenalty)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "U", rowNo), v.PaymentFYs.ArrearPenalty)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "V", rowNo), func() float64 {
			if v.PropertyPayment.Demand != nil {
				return v.PropertyPayment.Demand.FormFee
			}
			return 0
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "W", rowNo), v.PaymentFYs.ArrearRebate+v.PaymentFYs.CurrentRebate)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "X", rowNo), "0")
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "Y", rowNo), func() float64 {
			if v.PropertyPayment.Details != nil {
				return v.PropertyPayment.Details.Amount
			}
			return 0
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "Z", rowNo), func() string {
			if v.Ref.Creator.Name != "" {
				return v.Ref.Creator.Name
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "AA", rowNo), func() string {
			if v.Ref.Activator.Name != "" {
				return v.Ref.Activator.Name
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "AB", rowNo), func() string {
			if v.Ref.Collector.Name != "" {
				return v.Ref.Collector.Name
			}
			return "NA"
		}())
		totalAmount = totalAmount + v.PropertyPayment.Details.Amount

		rowNo++

	}

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "AB", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "Total")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "Y", rowNo), fmt.Sprintf(" %.2f", totalAmount))

	return excel, nil

}

// FilterUserWisePropertyCollectionReport : ""
func (s *Service) FilterUserWisePropertyCollectionReport(ctx *models.Context, filter *models.UserWisePropertyCollectionFilter, pagination *models.Pagination) (*models.RefUserWisePropertyCollectionWithTotal, error) {
	res, err := s.Daos.FilterUserWisePropertyCollectionReport(ctx, filter, pagination)
	if err != nil {
		return nil, err
	}

	var totalCash, totalCheque, totalNB, totalCashTransactions, totalChequeTransactions, totalNBTransactions float64
	var totalPropertyCash, totalPropertyCheque, totalPropertyNB, totalPropertyCashTransactions, totalPropertyChequeTransactions, totalPropertyNBTransactions float64
	var totalTLCash, totalTLCheque, totalTLNB, totalTLCashTransactions, totalTLChequeTransactions, totalTLNBTransactions float64
	var totalSRCash, totalSRCheque, totalSRNB, totalSRCashTransactions, totalSRChequeTransactions, totalSRNBTransactions float64
	var totalMTCash, totalMTCheque, totalMTNB, totalMTCashTransactions, totalMTChequeTransactions, totalMTNBTransactions float64

	if len(res) > 0 {

		for i, v := range res {
			propertyPayments := []models.PropertyPayment{}
			tlPayments := []models.TradeLicensePayments{}
			srPayments := []models.ShopRentPayments{}
			mtPayments := []models.MobileTowerPayments{}

			propertyPayments = append(propertyPayments, v.PropertyCash.PropertyPayments[:]...)
			propertyPayments = append(propertyPayments, v.PropertyCheque.PropertyPayments[:]...)
			propertyPayments = append(propertyPayments, v.PropertyNB.PropertyPayments[:]...)

			tlPayments = append(tlPayments, v.TLCash.TLPayments[:]...)
			tlPayments = append(tlPayments, v.TLCheque.TLPayments[:]...)
			tlPayments = append(tlPayments, v.TLNB.TLPayments[:]...)

			srPayments = append(srPayments, v.SRCash.SRPayments[:]...)
			srPayments = append(srPayments, v.SRCheque.SRPayments[:]...)
			srPayments = append(srPayments, v.SRNB.SRPayments[:]...)

			mtPayments = append(mtPayments, v.MTCash.MTPayments[:]...)
			mtPayments = append(mtPayments, v.MTCheque.MTPayments[:]...)
			mtPayments = append(mtPayments, v.MTNB.MTPayments[:]...)

			res[i].PropertyPayments = propertyPayments
			res[i].TLPayments = tlPayments
			res[i].SRPayments = srPayments
			res[i].MTPayments = mtPayments

			totalCash = totalCash + v.PropertyCash.TotalAmount + v.TLCash.TotalAmount + v.SRCash.TotalAmount + v.MTCash.TotalAmount
			totalCheque = totalCheque + v.PropertyCheque.TotalAmount + v.TLCheque.TotalAmount + v.SRCheque.TotalAmount + v.MTCheque.TotalAmount
			totalNB = totalNB + v.PropertyNB.TotalAmount + v.TLNB.TotalAmount + v.SRNB.TotalAmount + v.MTNB.TotalAmount
			totalCashTransactions = totalCashTransactions + v.PropertyCash.NoOfPayments + v.TLCash.NoOfPayments + v.SRCash.NoOfPayments + v.MTCash.NoOfPayments
			totalChequeTransactions = totalChequeTransactions + v.PropertyCheque.NoOfPayments + v.TLCheque.NoOfPayments + v.SRCheque.NoOfPayments + v.MTCheque.NoOfPayments
			totalNBTransactions = totalNBTransactions + v.PropertyNB.NoOfPayments + v.TLNB.NoOfPayments + v.SRNB.NoOfPayments + v.MTNB.NoOfPayments
			totalPropertyCash = totalPropertyCash + v.PropertyCash.TotalAmount
			totalPropertyCheque = totalPropertyCheque + v.PropertyCheque.TotalAmount
			totalPropertyNB = totalPropertyNB + v.PropertyNB.TotalAmount
			totalPropertyCashTransactions = totalPropertyCashTransactions + v.PropertyCash.NoOfPayments
			totalPropertyChequeTransactions = totalPropertyChequeTransactions + v.PropertyCheque.NoOfPayments
			totalPropertyNBTransactions = totalPropertyNBTransactions + v.PropertyNB.NoOfPayments
			totalTLCash = totalTLCash + v.TLCash.TotalAmount
			totalTLCheque = totalTLCheque + v.TLCheque.TotalAmount
			totalTLNB = totalTLNB + v.TLNB.TotalAmount
			totalTLCashTransactions = totalTLCashTransactions + v.TLCash.NoOfPayments
			totalTLChequeTransactions = totalTLChequeTransactions + v.TLCheque.NoOfPayments
			totalTLNBTransactions = totalTLNBTransactions + v.TLNB.NoOfPayments
			totalSRCash = totalSRCash + v.SRCash.TotalAmount
			totalSRCheque = totalSRCheque + v.SRCheque.TotalAmount
			totalSRNB = totalSRNB + v.SRNB.TotalAmount
			totalSRCashTransactions = totalSRCashTransactions + v.SRCash.NoOfPayments
			totalSRChequeTransactions = totalSRChequeTransactions + v.SRCheque.NoOfPayments
			totalSRNBTransactions = totalSRNBTransactions + v.SRNB.NoOfPayments
			totalMTCash = totalMTCash + v.MTCash.TotalAmount
			totalMTCheque = totalMTCheque + v.MTCheque.TotalAmount
			totalMTNB = totalMTNB + v.MTNB.TotalAmount
			totalMTCashTransactions = totalMTCashTransactions + v.MTCash.NoOfPayments
			totalMTChequeTransactions = totalMTChequeTransactions + v.MTCheque.NoOfPayments
			totalMTNBTransactions = totalMTNBTransactions + v.MTNB.NoOfPayments
		}
	}

	refTotal := new(models.RefUserWisePropertyCollectionWithTotal)
	refTotal.Collection = res

	refTotal.Total.TotalAmount = totalCash + totalCheque + totalNB
	refTotal.Total.TotalPayment = totalCashTransactions + totalChequeTransactions + totalNBTransactions
	refTotal.Total.TotalCashAmount = totalCash
	refTotal.Total.TotalChequeAmount = totalCheque
	refTotal.Total.TotalNBAmount = totalNB
	refTotal.Total.TotalCashPayment = totalCashTransactions
	refTotal.Total.TotalChequePayment = totalChequeTransactions
	refTotal.Total.TotalNBPayment = totalNBTransactions

	refTotal.Property.TotalAmount = totalPropertyCash + totalPropertyCheque + totalPropertyNB
	refTotal.Property.TotalPayment = totalPropertyCashTransactions + totalPropertyChequeTransactions + totalPropertyNBTransactions
	refTotal.Property.TotalCashAmount = totalPropertyCash
	refTotal.Property.TotalChequeAmount = totalPropertyCheque
	refTotal.Property.TotalNBAmount = totalPropertyNB
	refTotal.Property.TotalCashPayment = totalPropertyCashTransactions
	refTotal.Property.TotalChequePayment = totalPropertyChequeTransactions
	refTotal.Property.TotalNBPayment = totalPropertyNBTransactions

	refTotal.TradeLicense.TotalAmount = totalTLCash + totalTLCheque + totalTLNB
	refTotal.TradeLicense.TotalPayment = totalTLCashTransactions + totalTLChequeTransactions + totalTLNBTransactions
	refTotal.TradeLicense.TotalCashAmount = totalTLCash
	refTotal.TradeLicense.TotalChequeAmount = totalTLCheque
	refTotal.TradeLicense.TotalNBAmount = totalTLNB
	refTotal.TradeLicense.TotalCashPayment = totalTLCashTransactions
	refTotal.TradeLicense.TotalChequePayment = totalTLChequeTransactions
	refTotal.TradeLicense.TotalNBPayment = totalTLNBTransactions

	refTotal.ShopRent.TotalAmount = totalSRCash + totalSRCheque + totalSRNB
	refTotal.ShopRent.TotalPayment = totalSRCashTransactions + totalSRChequeTransactions + totalSRNBTransactions
	refTotal.ShopRent.TotalCashAmount = totalSRCash
	refTotal.ShopRent.TotalChequeAmount = totalSRCheque
	refTotal.ShopRent.TotalNBAmount = totalSRNB
	refTotal.ShopRent.TotalCashPayment = totalSRCashTransactions
	refTotal.ShopRent.TotalChequePayment = totalSRChequeTransactions
	refTotal.ShopRent.TotalNBPayment = totalSRNBTransactions

	refTotal.MobileTower.TotalAmount = totalMTCash + totalMTCheque + totalMTNB
	refTotal.MobileTower.TotalPayment = totalMTCashTransactions + totalMTChequeTransactions + totalMTNBTransactions
	refTotal.MobileTower.TotalCashAmount = totalMTCash
	refTotal.MobileTower.TotalChequeAmount = totalMTCheque
	refTotal.MobileTower.TotalNBAmount = totalMTNB
	refTotal.MobileTower.TotalCashPayment = totalMTCashTransactions
	refTotal.MobileTower.TotalChequePayment = totalMTChequeTransactions
	refTotal.MobileTower.TotalNBPayment = totalMTNBTransactions
	return refTotal, nil
}

// FilterHoldingWiseCollectionReportJSON : ""
func (s *Service) FilterHoldingWiseCollectionReportJSON(ctx *models.Context, filter *models.PropertyPaymentFilter, pagination *models.Pagination) ([]models.RefHoldingWiseCollectionReport, error) {
	defer ctx.Session.EndSession(ctx.CTX)
	return s.Daos.FilterHoldingWiseCollectionReportJSON(ctx, filter, pagination)
}

// FilterHoldingWiseCollectionReportExcel : ""
func (s *Service) FilterHoldingWiseCollectionReportExcel(ctx *models.Context, filter *models.PropertyPaymentFilter, pagination *models.Pagination) (*excelize.File, error) {
	res, err := s.FilterHoldingWiseCollectionReportJSON(ctx, filter, pagination)
	if err != nil {
		return nil, err
	}
	fmt.Println("'res length==>'", len(res))

	resPD, err1 := s.Daos.GetSingleDefaultProductConfiguration(ctx)
	if err1 != nil {
		return nil, err1
	}

	//  create an excel file
	excel := excelize.NewFile()
	sheet1 := "Holding Wise Collection Report"
	rowNo := 1
	index := excel.NewSheet(sheet1)
	excel.SetActiveSheet(index)
	if resPD.LocationID == "Bhagalpur" {
		excel.MergeCell(sheet1, "A1", "O3")
		excel.MergeCell(sheet1, "A4", "O5")
	} else {
		excel.MergeCell(sheet1, "A1", "B5")
		excel.MergeCell(sheet1, "C1", "O3")
		excel.MergeCell(sheet1, "C4", "O5")
	}
	excel.MergeCell(sheet1, "A6", "O6")
	excel.MergeCell(sheet1, "A7", "O7")
	style3, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#B6DDE8"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	style1, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#FFDC6D"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style3)

	if resPD.LocationID != "Bhagalpur" {
		documentUrl := s.Shared.Config.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.FILEURL)
		if err := excel.AddPicture(sheet1, fmt.Sprintf("%v%v", "A", rowNo), documentUrl+"municipal/logo.png", `{"x_scale": 0.115, "y_scale": 0.0935}`); err != nil {
			fmt.Println(err)
		}
	}

	if resPD.LocationID == "Bhagalpur" {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), ctx.ProductConfig.Name)
	} else {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), ctx.ProductConfig.Name)
	}

	rowNo++
	rowNo++
	rowNo++

	if resPD.LocationID == "Bhagalpur" {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), sheet1)
	} else {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), sheet1)
	}
	rowNo++
	rowNo++

	reportFromMsg2 := "Report"
	if filter != nil {
		fmt.Println(filter.DateRange.From, filter.DateRange.To)
		if filter.DateRange.From != nil && filter.DateRange.To == nil {
			reportFromMsg2 = reportFromMsg2 + " on " + fmt.Sprintf("%v-%v-%v", filter.DateRange.From.Day(), filter.DateRange.From.Month(), filter.DateRange.From.Year())
		}
		if filter.DateRange.From != nil && filter.DateRange.To != nil {
			reportFromMsg2 = reportFromMsg2 + " From " + fmt.Sprintf("%v-%v-%v", filter.DateRange.From.Day(), filter.DateRange.From.Month(), filter.DateRange.From.Year()) + " To " + fmt.Sprintf("%v-%v-%v", filter.DateRange.To.Day(), filter.DateRange.To.Month(), filter.DateRange.To.Year())
		}
		if filter.DateRange.From == nil && filter.DateRange.To == nil {
			fmt.Println("everything is nil")
		}

	}
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "O", rowNo), style3)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), reportFromMsg2)
	rowNo++

	t := time.Now()
	toDate := t.Format("02-January-2006")
	reportFromMsg3 := "Report Generated on" + " " + toDate
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "O", rowNo), style3)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), reportFromMsg3)
	rowNo++

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "O", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "S.No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "Ward No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "SAF No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), "Holding No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), "Owner Name")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), "Father/Husband Name/PAN")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), "Mobile No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "H", rowNo), "Road Type")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "I", rowNo), "Property Type")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "J", rowNo), "Address")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "K", rowNo), "Arrear collection")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "L", rowNo), "Current Collection")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "M", rowNo), "Penalty")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "N", rowNo), "Rebate")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "O", rowNo), "Total Collection")
	rowNo++

	var totalAmount float64

	for k, v := range res {
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), k+1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), func() string {
			if v.Ref.Ward.Name != "" {
				return v.Ref.Ward.Name
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), func() string {
			if v.Basic.Property.ApplicationNo != "" {
				return v.Basic.Property.ApplicationNo
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), func() string {
			if v.Basic.Property.UniqueID != "" {
				return v.Basic.Property.UniqueID
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), func() string {
			if v.Ref.Owner.Name != "" {
				return v.Ref.Owner.Name
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), func() string {
			if v.Ref.Owner.FatherRpanRhusband != "" {
				return v.Ref.Owner.FatherRpanRhusband
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), func() string {
			if v.Ref.Owner.Mobile != "" {
				return v.Ref.Owner.Mobile
			}
			return "NA"
		}())

		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "H", rowNo), func() string {
			if v.Ref.RoadType.Name != "" {
				return v.Ref.RoadType.Name
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "I", rowNo), func() string {
			if v.Ref.PropertyType.Name != "" {
				return v.Ref.PropertyType.Name
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "J", rowNo), func() string {
			if v.Basic.Property.Address.AL1 != "" {
				return v.Basic.Property.Address.AL1 + v.Basic.Property.Address.Al2
			}
			return "NA"
		}())
		// excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "K", rowNo), v.Payment.ArrearTax-v.Payment.ArrearAlreadyPaid)
		// excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "L", rowNo), v.Payment.CurrentTax-v.Payment.CurrentAlreadyPaid)
		// excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "M", rowNo), v.Payment.ArrearPenalty+v.Payment.CurrentPenalty)
		// excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "N", rowNo), v.Payment.ArrearRebate+v.Payment.CurrentRebate)
		// excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "O", rowNo), v.Payment.TotalAmount)
		// totalAmount = totalAmount + v.Payment.TotalAmount
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "K", rowNo), v.ArrearTax-v.ArrearAlreadyPaid)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "L", rowNo), v.CurrentTax-v.CurrentAlreadyPaid)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "M", rowNo), v.ArrearPenalty+v.CurrentPenalty)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "N", rowNo), v.ArrearRebate+v.CurrentRebate)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "O", rowNo), v.TotalAmount)
		totalAmount = totalAmount + v.TotalAmount

		rowNo++

	}

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "O", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "Total")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "O", rowNo), fmt.Sprintf(" %.2f", totalAmount))

	return excel, nil

}

// PropertyWiseDemandandCollectionV2JSON : ""
func (s *Service) PropertyWiseDemandandCollectionV2JSON(ctx *models.Context, propertyfilter *models.PropertyFilter, pagination *models.Pagination) (property []models.ResPropertyWiseDemandandCollectionV2Report, err error) {
	defer ctx.Session.EndSession(ctx.CTX)
	return s.Daos.PropertyWiseDemandandCollectionV2JSON(ctx, propertyfilter, pagination)

}

// PropertyWiseDemandandCollectionExcel: ""
func (s *Service) PropertyWiseDemandandCollectionV2Excel(ctx *models.Context, filter *models.PropertyFilter, pagination *models.Pagination) (*excelize.File, error) {
	res, err := s.PropertyWiseDemandandCollectionV2JSON(ctx, filter, pagination)
	if err != nil {
		return nil, err
	}
	fmt.Println("'res length==>'", len(res))
	resPD, err1 := s.Daos.GetSingleDefaultProductConfiguration(ctx)
	if err1 != nil {
		return nil, err1
	}
	resFYFrom, err := s.Daos.GetSingleFinancialYearUsingDateV2(ctx, filter.AppliedRange.From)
	if err != nil {
		return nil, err
	}
	resFYTo, err := s.Daos.GetSingleFinancialYearUsingDateV2(ctx, filter.AppliedRange.To)
	if err != nil {
		return nil, err
	}
	//  create an excel file
	excel := excelize.NewFile()
	sheet1 := "Property Wise Demand and Collection"
	rowNo := 1
	index := excel.NewSheet(sheet1)
	excel.SetActiveSheet(index)
	if resPD.LocationID == "Bhagalpur" {
		excel.MergeCell(sheet1, "A1", "V3")
		excel.MergeCell(sheet1, "A4", "V5")
		excel.MergeCell(sheet1, "A6", "V6")
		excel.MergeCell(sheet1, "A7", "V7")
	} else {
		excel.MergeCell(sheet1, "A1", "B5")
		excel.MergeCell(sheet1, "C1", "V3")
		excel.MergeCell(sheet1, "C4", "V5")
		excel.MergeCell(sheet1, "A6", "V6")
		excel.MergeCell(sheet1, "A7", "V7")
	}

	style1, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#B6DDE8"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	style2, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#FFDC6D"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	style4, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#FFDC6D"],"pattern":1},"alignment":{"horizontal":"right","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style1)

	if resPD.LocationID != "Bhagalpur" {
		documentUrl := s.Shared.Config.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.FILEURL)
		if err := excel.AddPicture(sheet1, fmt.Sprintf("%v%v", "A", rowNo), documentUrl+"municipal/logo.png", `{"x_scale": 0.115, "y_scale": 0.0935}`); err != nil {
			fmt.Println(err)
		}
	}

	if resPD.LocationID == "Bhagalpur" {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), ctx.ProductConfig.Name)
	} else {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), ctx.ProductConfig.Name)
	}

	rowNo++
	rowNo++
	rowNo++
	if resPD.LocationID == "Bhagalpur" {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "Property Wise Demand And Collection")
	} else {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Property Wise Demand And Collection")
	}
	rowNo++
	rowNo++

	t := time.Now()
	toDate := t.Format("02-January-2006")

	reportFromMsg2 := "Report"
	if filter != nil {
		if filter.AppliedRange != nil {
			if filter.AppliedRange.From != nil && filter.AppliedRange.To == nil {
				reportFromMsg2 = reportFromMsg2 + " on " + fmt.Sprintf("%v-%v-%v", filter.AppliedRange.From.Day(), filter.AppliedRange.From.Month(), filter.AppliedRange.From.Year())
			}
			if filter.AppliedRange.From != nil && filter.AppliedRange.To != nil {
				reportFromMsg2 = reportFromMsg2 + " From " + fmt.Sprintf("%v-%v-%v", filter.AppliedRange.From.Day(), filter.AppliedRange.From.Month(), filter.AppliedRange.From.Year()) + " To " + fmt.Sprintf("%v-%v-%v", filter.AppliedRange.To.Day(), filter.AppliedRange.To.Month(), filter.AppliedRange.To.Year())
			}
			if filter.AppliedRange.From == nil && filter.AppliedRange.To == nil {
				fmt.Println("everything is nil")
			}
		}
	}
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), reportFromMsg2)
	rowNo++

	reportFromMsg := "Report Generated on" + " " + toDate
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), reportFromMsg)
	rowNo++

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "V", rowNo), style2)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "S.No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "Ward No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Property No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), "Owner Name")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), "Mobile No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), "Property Type")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), "Address")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "H", rowNo), "Road Type")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "I", rowNo), "From Year")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "J", rowNo), "To Year")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "K", rowNo), "Current Demand")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "L", rowNo), "Arrear Demand")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "M", rowNo), "Penalty")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "N", rowNo), "Total Demand")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "O", rowNo), "Current Collection")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "P", rowNo), "Arrear Collection")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "Q", rowNo), "Penalty")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "R", rowNo), "Total Collection")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "S", rowNo), "Current Pending Demand")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "T", rowNo), "Arrear Pending Demand")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "U", rowNo), "Penalty")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "V", rowNo), "Total Pending Demand")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "W", rowNo), "Current Outstanding Demand")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "X", rowNo), "Total Outstanding Demand")

	fmt.Println("'res length==>'", len(res))
	var arrearDemand, currentDemand, totalDemand float64
	var arrearCollection, currentCollection, totalCollection float64
	var arrearOutstandingDemand, currentOutstandingDemand, totalOutstandingDemand float64

	for i, v := range res {

		// v.Ref.Demand.Total.TotalTax = v.Ref.Demand.Total.TotalTax - v.Ref.Demand.Total.Ecess
		// v.Ref.Collections.TotalTax = v.Ref.Collections.TotalTax + v.Ref.PropertyPayments.FormFee
		// v.Ref.Demand.Current.TotalTax = v.Ref.Demand.Current.TotalTax + v.Ref.PropertyPayments.Rebate

		// arrearDemand = arrearDemand + v.Ref.Demand.Arrear.TotalTax
		// currentDemand = currentDemand + v.Ref.Demand.Current.TotalTax
		// totalDemand = totalDemand + v.Ref.Demand.Total.TotalTax
		// arrearCollection = arrearCollection + v.Ref.Collections.ArrearTax
		// currentCollection = currentCollection + v.Ref.Collections.CurrentTax
		// totalCollection = totalCollection + v.Ref.Collections.TotalTax
		// arrearOutstandingDemand = arrearOutstandingDemand + (v.Ref.Demand.Arrear.TotalTax - v.Ref.Collections.ArrearTax - v.Ref.Collections.ArrearPenalty)
		// currentOutstandingDemand = currentOutstandingDemand + (v.Ref.Demand.Current.TotalTax - v.Ref.Collections.CurrentTax - v.Ref.Collections.CurrentPenalty)
		// totalOutstandingDemand = totalOutstandingDemand + (v.Ref.Demand.Total.TotalTax - v.Ref.Collections.TotalTax)

		rowNo++
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), i+1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), func() string {
			if v.Ref.Address.Ward != nil {
				return v.Ref.Address.Ward.Name
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), v.Property.UniqueID)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), func() string {
			if v.Ref.PropertyOwner != nil {
				return v.Ref.PropertyOwner.Name
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), func() string {
			if v.Ref.PropertyOwner != nil {
				return v.Ref.PropertyOwner.Mobile
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), func() string {
			if v.Ref.PropertyType != nil {
				return v.Ref.PropertyType.Name
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), v.Property.Address.AL1+" , "+v.Property.Address.Al2)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "H", rowNo), func() string {
			if v.Ref.RoadType != nil {
				return v.Ref.RoadType.Name
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "I", rowNo), func() string {
			if resFYFrom != nil {
				return resFYFrom.Name
			}
			return v.Property.DOA.Format("02-January-2006")
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "J", rowNo), func() string {
			if resFYTo != nil {
				return resFYTo.Name
			}
			return t.Format("02-January-2006")
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "K", rowNo), v.Ref.Demand.Current.TotalTax)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "L", rowNo), v.Ref.Demand.Arrear.TotalTax)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "M", rowNo), v.Ref.Demand.Total.Penalty)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "N", rowNo), v.Ref.Demand.Total.TotalTax)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "O", rowNo), v.Ref.Payments.PaymentFys.CurrentTax)
		// excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "O", rowNo), (v.Ref.Payments.PaymentFys.ArrearPenalty + v.Ref.Payments.PaymentFys.CurrentPenalty))
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "P", rowNo), v.Ref.Payments.PaymentFys.ArrearTax)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "Q", rowNo), (v.Ref.Payments.PaymentFys.ArrearPenalty + v.Ref.Payments.PaymentFys.CurrentPenalty))
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "R", rowNo), v.Ref.Payments.PaymentFys.TotalTax)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "S", rowNo), (v.Ref.Demand.Current.TotalTax - v.Ref.Payments.PaymentFys.CurrentTax - v.Ref.Payments.PaymentFys.CurrentPenalty))
		// excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "T", rowNo), (v.Ref.Demand.Arrear.TotalTax - v.Ref.Payments.PaymentFys.ArrearTax - v.Ref.Payments.PaymentFys.ArrearPenalty))
		// excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "U", rowNo), v.Ref.Collections.TotalTax)
		// excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "V", rowNo), (v.Ref.Demand.Arrear.TotalTax - v.Ref.Collections.ArrearTax - v.Ref.Collections.ArrearPenalty))

	}
	rowNo++
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style2)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "Total")

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "B", rowNo), fmt.Sprintf("%v%v", "B", rowNo), style4)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style4)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "D", rowNo), fmt.Sprintf("%v%v", "D", rowNo), style4)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "E", rowNo), fmt.Sprintf("%v%v", "E", rowNo), style4)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "F", rowNo), fmt.Sprintf("%v%v", "F", rowNo), style4)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "G", rowNo), fmt.Sprintf("%v%v", "G", rowNo), style4)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "H", rowNo), fmt.Sprintf("%v%v", "H", rowNo), style4)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "I", rowNo), fmt.Sprintf("%v%v", "I", rowNo), style4)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "J", rowNo), fmt.Sprintf("%v%v", "J", rowNo), style4)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "K", rowNo), fmt.Sprintf("%v%v", "K", rowNo), style4)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "R", rowNo), fmt.Sprintf("%v%v", "R", rowNo), style4)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "S", rowNo), fmt.Sprintf("%v%v", "S", rowNo), style4)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "T", rowNo), fmt.Sprintf("%v%v", "T", rowNo), style4)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "U", rowNo), fmt.Sprintf("%v%v", "U", rowNo), style4)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "V", rowNo), fmt.Sprintf("%v%v", "V", rowNo), style4)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "W", rowNo), fmt.Sprintf("%v%v", "W", rowNo), style4)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "X", rowNo), fmt.Sprintf("%v%v", "X", rowNo), style4)

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "L", rowNo), fmt.Sprintf("%v%v", "L", rowNo), style4)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "L", rowNo), fmt.Sprintf("%.2f", arrearDemand))
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "M", rowNo), fmt.Sprintf("%v%v", "M", rowNo), style4)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "M", rowNo), fmt.Sprintf("%.2f", currentDemand))
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "Q", rowNo), fmt.Sprintf("%v%v", "Q", rowNo), style4)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "Q", rowNo), fmt.Sprintf("%.2f", totalDemand))
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "R", rowNo), fmt.Sprintf("%v%v", "R", rowNo), style4)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "R", rowNo), fmt.Sprintf("%.2f", arrearCollection))
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "S", rowNo), fmt.Sprintf("%v%v", "S", rowNo), style4)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "S", rowNo), fmt.Sprintf("%.2f", currentCollection))
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "U", rowNo), fmt.Sprintf("%v%v", "U", rowNo), style4)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "U", rowNo), fmt.Sprintf("%.2f", totalCollection))
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "V", rowNo), fmt.Sprintf("%v%v", "V", rowNo), style4)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "V", rowNo), fmt.Sprintf("%.2f", arrearOutstandingDemand))
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "W", rowNo), fmt.Sprintf("%v%v", "W", rowNo), style4)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "W", rowNo), fmt.Sprintf("%.2f", currentOutstandingDemand))
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "X", rowNo), fmt.Sprintf("%v%v", "X", rowNo), style4)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "X", rowNo), fmt.Sprintf("%.2f", totalOutstandingDemand))
	return excel, nil
}

// PropertyDemandAndCollectionReportJSON : ""
func (s *Service) PropertyDemandAndCollectionReportJSON(ctx *models.Context, filter *models.PropertyFilter, pagination *models.Pagination) (property []models.RefPropertyDemandAndCollectionReport, err error) {
	defer ctx.Session.EndSession(ctx.CTX)
	return s.Daos.PropertyDemandAndCollectionReportJSON(ctx, filter, pagination)
}

// PropertyDemandAndCollectionReportExcel: ""
func (s *Service) PropertyDemandAndCollectionReportExcel(ctx *models.Context, filter *models.PropertyFilter, pagination *models.Pagination) (*excelize.File, error) {
	res, err := s.PropertyDemandAndCollectionReportJSON(ctx, filter, pagination)
	if err != nil {
		return nil, err
	}
	fmt.Println("'res length==>'", len(res))
	resPD, err1 := s.Daos.GetSingleDefaultProductConfiguration(ctx)
	if err1 != nil {
		return nil, err1
	}
	resFYs, err := s.Daos.GetSingleFinancialYearUsingDateV2(ctx, filter.Date)
	if err != nil {
		return nil, errors.New("error in getting current financial year" + err.Error())
	}

	//  create an excel file
	excel := excelize.NewFile()
	sheet1 := "Property Wise Demand and Collection"
	rowNo := 1
	index := excel.NewSheet(sheet1)
	excel.SetActiveSheet(index)
	if resPD.LocationID == "Bhagalpur" {
		excel.MergeCell(sheet1, "A1", "M3")
		excel.MergeCell(sheet1, "A4", "M5")
		excel.MergeCell(sheet1, "A6", "M6")
		excel.MergeCell(sheet1, "A7", "A8")
		excel.MergeCell(sheet1, "B7", "B8")
		excel.MergeCell(sheet1, "C7", "C8")
		excel.MergeCell(sheet1, "D7", "F7")
		excel.MergeCell(sheet1, "G7", "M7")
	} else {
		excel.MergeCell(sheet1, "A1", "B5")
		excel.MergeCell(sheet1, "C1", "M3")
		excel.MergeCell(sheet1, "C4", "M5")
		excel.MergeCell(sheet1, "A6", "M6")
		excel.MergeCell(sheet1, "A7", "A8")
		excel.MergeCell(sheet1, "B7", "B8")
		excel.MergeCell(sheet1, "C7", "C8")
		excel.MergeCell(sheet1, "D7", "F7")
		excel.MergeCell(sheet1, "G7", "M7")

	}

	style1, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#B6DDE8"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	style2, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#FFDC6D"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style1)

	if resPD.LocationID != "Bhagalpur" {
		documentUrl := s.Shared.Config.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.FILEURL)
		if err := excel.AddPicture(sheet1, fmt.Sprintf("%v%v", "A", rowNo), documentUrl+"municipal/logo.png", `{"x_scale": 0.115, "y_scale": 0.0935}`); err != nil {
			fmt.Println(err)
		}
	}

	if resPD.LocationID == "Bhagalpur" {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), ctx.ProductConfig.Name)
	} else {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), ctx.ProductConfig.Name)
	}

	rowNo++
	rowNo++
	rowNo++
	if resPD.LocationID == "Bhagalpur" {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "Property Wise Demand And Collection")
	} else {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Property Wise Demand And Collection")
	}
	rowNo++
	rowNo++

	t := time.Now()
	toDate := t.Format("02-January-2006")

	reportFromMsg := "Report Generated on" + " " + toDate
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), reportFromMsg)
	rowNo++

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "M", rowNo), style2)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "S.No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "Property No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Property Type")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), "Overall Demand")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), fmt.Sprintf("%v%v%v", resFYs.Name, " ", "Collections"))
	rowNo++
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "M", rowNo), style2)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), "Arrear")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), "Current")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), "Total")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), "Arrear")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "H", rowNo), "Current")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "I", rowNo), "Penalty")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "J", rowNo), "Rebate")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "K", rowNo), "Form Fee")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "L", rowNo), "Boring Charge")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "M", rowNo), "Other Demand")

	fmt.Println("'res length==>'", len(res))
	// var arrearDemand, currentDemand, totalDemand float64
	// var arrearCollection, currentCollection, totalCollection float64
	// var arrearOutstandingDemand, currentOutstandingDemand, totalOutstandingDemand float64

	for i, v := range res {

		rowNo++
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), i+1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v.UniqueID)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), func() string {
			if v.PropertyType != nil {
				return v.PropertyType.Name
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), v.Demand.Actual.Arrear.TotalTax)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), v.Demand.Actual.Current.TotalTax)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), v.Demand.Actual.Total.TotalTax)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), v.Collections.ArrearTax)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "H", rowNo), v.Collections.CurrentTax)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "I", rowNo), v.Collections.ArrearPenalty+v.Collections.CurrentPenalty)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "J", rowNo), v.Collections.ArrearRebate+v.Collections.CurrentRebate)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "K", rowNo), v.Collections.FormFee)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "L", rowNo), v.Collections.BoreCharge)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "M", rowNo), v.Collections.ArrearOtherDemand+v.Collections.CurrentOtherDemand)

	}
	// rowNo++
	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "K", rowNo), style2)
	// excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "Total")

	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "L", rowNo), fmt.Sprintf("%v%v", "L", rowNo), style4)
	// excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "L", rowNo), fmt.Sprintf("%.2f", arrearDemand))

	return excel, nil
}

func (s *Service) FilterPaymentCOllection(ctx *models.Context, filter *models.PropertyPaymentFilter, pagination *models.Pagination) ([]models.RefPropertyCollectionReport, error) {
	defer ctx.Session.EndSession(ctx.CTX)
	return s.Daos.FilterPaymentCOllection(ctx, filter, pagination)
}

// CounterReport : ""
func (s *Service) CounterReportV2(ctx *models.Context, filter *models.PropertyPaymentFilter, pagination *models.Pagination) (*excelize.File, error) {
	res, err := s.FilterPaymentCOllection(ctx, filter, pagination)
	if err != nil {
		return nil, err
	}
	fmt.Println("'res length==>'", len(res))

	resPD, err1 := s.Daos.GetSingleDefaultProductConfiguration(ctx)
	if err1 != nil {
		return nil, err1
	}

	//  create an excel file
	excel := excelize.NewFile()
	sheet1 := "Counter Report"
	rowNo := 1
	index := excel.NewSheet(sheet1)
	excel.SetActiveSheet(index)
	if resPD.LocationID == "Bhagalpur" {
		excel.MergeCell(sheet1, "A1", "AB3")
		excel.MergeCell(sheet1, "A4", "AB5")
	} else {
		excel.MergeCell(sheet1, "A1", "B5")
		excel.MergeCell(sheet1, "C1", "AB3")
		excel.MergeCell(sheet1, "C4", "AB5")
	}
	excel.MergeCell(sheet1, "A6", "AB6")
	excel.MergeCell(sheet1, "A7", "AB7")
	excel.MergeCell(sheet1, "A8", "AB8")

	style3, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#B6DDE8"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	style1, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#FFDC6D"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style3)

	if resPD.LocationID != "Bhagalpur" {
		documentUrl := s.Shared.Config.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.FILEURL)
		if err := excel.AddPicture(sheet1, fmt.Sprintf("%v%v", "A", rowNo), documentUrl+"municipal/logo.png", `{"x_scale": 0.115, "y_scale": 0.0935}`); err != nil {
			fmt.Println(err)
		}
	}

	if resPD.LocationID == "Bhagalpur" {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), ctx.ProductConfig.Name)
	} else {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), ctx.ProductConfig.Name)
	}

	rowNo++
	rowNo++
	rowNo++

	if resPD.LocationID == "Bhagalpur" {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "Counter Report")
	} else {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Counter Report")
	}
	rowNo++
	rowNo++

	reportFromMsg := "Counter Collection Report"
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style3)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), reportFromMsg)
	rowNo++

	reportFromMsg2 := "Report"
	// if filter != nil {
	// 	fmt.Println(filter.DateRange.From, filter.DateRange.To)
	// 	if filter.DateRange.From != nil && filter.DateRange.To == nil {
	// 		reportFromMsg2 = reportFromMsg2 + " on " + fmt.Sprintf("%v-%v-%v", filter.DateRange.From.Day(), filter.DateRange.From.Month(), filter.DateRange.From.Year())
	// 	}
	// 	if filter.DateRange.From != nil && filter.DateRange.To != nil {
	// 		reportFromMsg2 = reportFromMsg2 + " From " + fmt.Sprintf("%v-%v-%v", filter.DateRange.From.Day(), filter.DateRange.From.Month(), filter.DateRange.From.Year()) + " To " + fmt.Sprintf("%v-%v-%v", filter.DateRange.To.Day(), filter.DateRange.To.Month(), filter.DateRange.To.Year())
	// 	}
	// 	if filter.DateRange.From == nil && filter.DateRange.To == nil {
	// 		fmt.Println("everything is nil")
	// 	}

	// }
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style3)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), reportFromMsg2)
	rowNo++

	t := time.Now()
	toDate := t.Format("02-January-2006")
	reportFromMsg3 := "Report Generated on" + " " + toDate
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style3)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), reportFromMsg3)
	rowNo++

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "AB", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "S.No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "Holding No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Application No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), "Old Holding No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), "Ward")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), "Owner Name")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), "Address")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "H", rowNo), "Mobile")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "I", rowNo), "From Year")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "J", rowNo), "Upto Year")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "K", rowNo), "Txn Date")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "L", rowNo), "Receipt No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "M", rowNo), "Payment Mode")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "N", rowNo), "Cheque / DD No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "O", rowNo), "Bank")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "P", rowNo), "Branch Name")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "Q", rowNo), "Received Amount")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "R", rowNo), "Current Holding Tax")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "S", rowNo), "Arrear Holding Tax")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "T", rowNo), "Current Penalty")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "U", rowNo), "Arrear Penalty")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "V", rowNo), "From Fee")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "W", rowNo), "Discount")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "X", rowNo), "Bore Charge")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "Y", rowNo), "Other Demand")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "Z", rowNo), "Already Payed")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "AA", rowNo), "Total Amount")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "AB", rowNo), "Assessment By")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "AC", rowNo), "Activated By")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "AD", rowNo), "Collected By")
	rowNo++

	var totalAmount float64

	for k, v := range res {
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), k+1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v.PropertyID)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), v.Property.ApplicationNo)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), func() string {
			if v.Property.OldHoldingNumber != "" {
				return v.Property.OldHoldingNumber
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), v.Property.Address.WardCode)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), v.Owner.Name)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), v.Address.AL1+v.Address.Al2)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "H", rowNo), v.Owner.Mobile)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "I", rowNo), func() string {
			if len(v.PaymentFYs.FYs) > 0 {
				return v.PaymentFYs.FYs[0].FY.Name
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "J", rowNo), func() string {
			if len(v.PaymentFYs.FYs) > 0 {
				if len(v.PaymentFYs.FYs) == 1 {
					return v.PaymentFYs.FYs[0].FY.Name
				}
				if len(v.PaymentFYs.FYs) > 1 {
					return v.PaymentFYs.FYs[len(v.PaymentFYs.FYs)-1].FY.Name

				}
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "K", rowNo), v.CompletionDate.Format("2006-01-02"))
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "L", rowNo), v.ReciptNo)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "M", rowNo), func() string {
			if v.Details != nil {
				switch v.Details.MOP.Mode {
				case constants.MOPCASH:
					return "Cash"
				case constants.MOPCHEQUE:
					return "Cheque"
				case constants.MOPDD:
					return "DD"
				case constants.MOPNETBANKING:
					return "Online"
				default:
					return "Invalid"
				}
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "N", rowNo), func() string {
			if v.Details != nil {
				if v.Details.MOP.Mode == constants.MOPCHEQUE {
					if v.Details.MOP.Cheque != nil {
						return v.Details.MOP.Cheque.No
					}
				} else if v.Details.MOP.Mode == constants.MOPDD {
					if v.Details.MOP.DD != nil {
						return v.Details.MOP.DD.No
					}
				} else {
					return "NA"
				}
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "O", rowNo), func() string {
			if v.Details != nil {
				if v.Details.MOP.Mode == constants.MOPCHEQUE {
					if v.Details.MOP.Cheque != nil {
						return v.Details.MOP.Cheque.Bank
					}
				} else if v.Details.MOP.Mode == constants.MOPDD {
					if v.Details.MOP.DD != nil {
						return v.Details.MOP.DD.Bank
					}
				} else {
					return "NA"
				}
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "P", rowNo), func() string {
			if v.Details != nil {
				if v.Details.MOP.Mode == constants.MOPCHEQUE {
					if v.Details.MOP.Cheque != nil {
						return v.Details.MOP.Cheque.Branch
					}
				} else if v.Details.MOP.Mode == constants.MOPDD {
					if v.Details.MOP.DD != nil {
						return v.Details.MOP.DD.Branch
					}
				} else {
					return "NA"
				}
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "Q", rowNo), v.PaymentFYs.TotalTax+func() float64 {
			if v.Demand != nil {
				return v.Demand.FormFee + v.Demand.BoreCharge
			}
			return 0
		}())
		if v.ReciptNo == "11717" {
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "Q", rowNo), func() string {
				if v.Details != nil {

					return fmt.Sprintf("  %v, ", v.PaymentFYs.TotalTax+v.Demand.FormFee+v.Demand.BoreCharge)
				}
				return ""
			}())
		}
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "R", rowNo), v.PaymentFYs.CurrentTax-v.PaymentFYs.CurrentAlreadyPaid)
		if v.ReciptNo == "11717" {
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "R", rowNo), func() string {
				if v.Details != nil {

					return fmt.Sprintf("  %v, ", v.PaymentFYs.CurrentTax-v.PaymentFYs.CurrentAlreadyPaid)
				}
				return ""
			}())
		}

		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "S", rowNo), v.PaymentFYs.ArrearTax-v.PaymentFYs.ArrearAlreadyPaid)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "T", rowNo), v.PaymentFYs.CurrentPenalty)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "U", rowNo), v.PaymentFYs.ArrearPenalty)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "V", rowNo), func() float64 {
			if v.Demand != nil {
				return v.Demand.FormFee
			}
			return 0
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "W", rowNo), v.PaymentFYs.ArrearRebate+v.PaymentFYs.CurrentRebate)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "X", rowNo), v.Demand.BoreCharge)
		if v.ReciptNo == "11717" {
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "X", rowNo), func() string {
				if v.Details != nil {

					return fmt.Sprintf("  %v, ", v.Demand.BoreCharge)
				}
				return ""
			}())
		}

		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "Y", rowNo), v.Demand.OtherDemand)
		if v.ReciptNo == "11717" {
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "Y", rowNo), func() string {
				if v.Details != nil {

					return fmt.Sprintf("  %v, ", v.Demand.OtherDemand)
				}
				return ""
			}())
		}
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "Z", rowNo), v.Demand.PreviousCollection.Amount)

		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "AA", rowNo), func() float64 {
			if v.Details != nil {

				return v.Details.Amount
			}
			return 0
		}())

		if v.ReciptNo == "11717" {
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "AA", rowNo), func() string {
				if v.Details != nil {

					return fmt.Sprintf("  %v, ", v.Details.Amount)
				}
				return ""
			}())
		}
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "AB", rowNo), func() string {
			if v.Creator.Name != "" {
				return v.Creator.Name
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "AC", rowNo), func() string {
			if v.Activator.Name != "" {
				return v.Activator.Name
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "AD", rowNo), func() string {
			if v.Collector.Name != "" {
				return v.Collector.Name
			}
			return "NA"
		}())
		if v.ReciptNo != "11717" {
			totalAmount = totalAmount + v.Details.Amount
		}
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "AE", rowNo), v.Status)

		rowNo++

	}

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "AB", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "Total")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "AA", rowNo), fmt.Sprintf(" %.2f", totalAmount))

	return excel, nil

}

func (s *Service) FilterPaymentSummery() {
	c := context.TODO()
	ctx := app.GetApp(c, s.Daos)
	defer ctx.Client.Disconnect(c)
	data, err := s.Daos.FilterPaymentSummary(ctx)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	c2 := context.TODO()

	for _, v := range data {
		ctx2 := app.GetApp(c2, s.Daos)
		// defer ctx2.Client.Disconnect(c)
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Got panic - ")
				fmt.Println(r)
				time.Sleep(10000)
				ctx2 = app.GetApp(c2, s.Daos)
			}
		}()
		fmt.Println(v)
		fmt.Println(s.Daos.UpdatePaymentSummary(ctx2, v))
		ctx2.Client.Disconnect(c2)
	}

	fmt.Println(len(data))

}

func (s *Service) FilterHoldingWiseCollectionReportJSONV2(ctx *models.Context, filter *models.PropertyPaymentFilter, pagination *models.Pagination) ([]models.RefHoldingWiseCollectionReportV2, error) {
	defer ctx.Session.EndSession(ctx.CTX)
	return s.Daos.FilterHoldingWiseCollectionReportJSONV2(ctx, filter, pagination)
}

func (s *Service) FilterHoldingWiseCollectionReportExcelV2(ctx *models.Context, filter *models.PropertyPaymentFilter, pagination *models.Pagination) (*excelize.File, error) {
	res, err := s.FilterHoldingWiseCollectionReportJSONV2(ctx, filter, pagination)
	if err != nil {
		return nil, err
	}
	fmt.Println("'res length==>'", len(res))

	resPD, err1 := s.Daos.GetSingleDefaultProductConfiguration(ctx)
	if err1 != nil {
		return nil, err1
	}

	//  create an excel file
	excel := excelize.NewFile()
	sheet1 := "Holding Wise Collection Report"
	rowNo := 1
	index := excel.NewSheet(sheet1)
	excel.SetActiveSheet(index)
	if resPD.LocationID == "Bhagalpur" {
		excel.MergeCell(sheet1, "A1", "O3")
		excel.MergeCell(sheet1, "A4", "O5")
	} else {
		excel.MergeCell(sheet1, "A1", "B5")
		excel.MergeCell(sheet1, "C1", "O3")
		excel.MergeCell(sheet1, "C4", "O5")
	}
	excel.MergeCell(sheet1, "A6", "O6")
	excel.MergeCell(sheet1, "A7", "O7")
	style3, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#B6DDE8"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	style1, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#FFDC6D"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style3)

	if resPD.LocationID != "Bhagalpur" {
		documentUrl := s.Shared.Config.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.FILEURL)
		if err := excel.AddPicture(sheet1, fmt.Sprintf("%v%v", "A", rowNo), documentUrl+"municipal/logo.png", `{"x_scale": 0.115, "y_scale": 0.0935}`); err != nil {
			fmt.Println(err)
		}
	}

	if resPD.LocationID == "Bhagalpur" {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), ctx.ProductConfig.Name)
	} else {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), ctx.ProductConfig.Name)
	}

	rowNo++
	rowNo++
	rowNo++

	if resPD.LocationID == "Bhagalpur" {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), sheet1)
	} else {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), sheet1)
	}
	rowNo++
	rowNo++

	reportFromMsg2 := "Report"
	if filter != nil {
		fmt.Println(filter.DateRange.From, filter.DateRange.To)
		if filter.DateRange.From != nil && filter.DateRange.To == nil {
			reportFromMsg2 = reportFromMsg2 + " on " + fmt.Sprintf("%v-%v-%v", filter.DateRange.From.Day(), filter.DateRange.From.Month(), filter.DateRange.From.Year())
		}
		if filter.DateRange.From != nil && filter.DateRange.To != nil {
			reportFromMsg2 = reportFromMsg2 + " From " + fmt.Sprintf("%v-%v-%v", filter.DateRange.From.Day(), filter.DateRange.From.Month(), filter.DateRange.From.Year()) + " To " + fmt.Sprintf("%v-%v-%v", filter.DateRange.To.Day(), filter.DateRange.To.Month(), filter.DateRange.To.Year())
		}
		if filter.DateRange.From == nil && filter.DateRange.To == nil {
			fmt.Println("everything is nil")
		}

	}
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "O", rowNo), style3)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), reportFromMsg2)
	rowNo++

	t := time.Now()
	toDate := t.Format("02-January-2006")
	reportFromMsg3 := "Report Generated on" + " " + toDate
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "O", rowNo), style3)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), reportFromMsg3)
	rowNo++

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "O", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "S.No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "Ward No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "SAF No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), "Holding No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), "Owner Name")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), "Father/Husband Name/PAN")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), "Mobile No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "H", rowNo), "Road Type")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "I", rowNo), "Property Type")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "J", rowNo), "Address")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "K", rowNo), "Arrear collection")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "L", rowNo), "Current Collection")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "M", rowNo), "Penalty")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "N", rowNo), "Rebate")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "O", rowNo), "Total Collection")
	rowNo++

	var totalAmount float64

	for k, v := range res {
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), k+1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), func() string {
			if v.Ward.Name != "" {
				return v.Ward.Name
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), func() string {
			if v.Basic.ApplicationNo != "" {
				return v.Basic.ApplicationNo
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), func() string {
			if v.Basic.OldHoldingNumber != "" {
				return v.Basic.OldHoldingNumber
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), func() string {
			if v.Owner.Name != "" {
				return v.Owner.Name
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), func() string {
			if v.Owner.FatherRpanRhusband != "" {
				return v.Owner.FatherRpanRhusband
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), func() string {
			if v.Owner.Mobile != "" {
				return v.Owner.Mobile
			}
			return "NA"
		}())

		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "H", rowNo), func() string {
			if v.RoadType.Name != "" {
				return v.RoadType.Name
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "I", rowNo), func() string {
			if v.PropertyType.Name != "" {
				return v.PropertyType.Name
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "J", rowNo), func() string {
			if v.Basic.Address.AL1 != "" {
				return v.Basic.Address.AL1 + v.Basic.Address.Al2
			}
			return "NA"
		}())
		// excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "K", rowNo), v.Payment.ArrearTax-v.Payment.ArrearAlreadyPaid)
		// excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "L", rowNo), v.Payment.CurrentTax-v.Payment.CurrentAlreadyPaid)
		// excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "M", rowNo), v.Payment.ArrearPenalty+v.Payment.CurrentPenalty)
		// excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "N", rowNo), v.Payment.ArrearRebate+v.Payment.CurrentRebate)
		// excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "O", rowNo), v.Payment.TotalAmount)
		// totalAmount = totalAmount + v.Payment.TotalAmount
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "K", rowNo), v.Payment.ArrearTax-v.Payment.ArrearAlreadyPaid)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "L", rowNo), v.Payment.CurrentTax-v.Payment.CurrentAlreadyPaid)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "M", rowNo), v.Payment.ArrearPenalty+v.Payment.CurrentPenalty)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "N", rowNo), v.Payment.ArrearRebate+v.Payment.CurrentRebate)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "O", rowNo), v.Payment.TotalAmount)
		totalAmount = totalAmount + v.Payment.TotalAmount

		rowNo++

	}

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "O", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "Total")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "O", rowNo), fmt.Sprintf(" %.2f", totalAmount))

	return excel, nil

}

func (s *Service) PropertyWiseDemandandCollectionExcelV2(ctx *models.Context, propertyfilter *models.PropertyFilter, pagination *models.Pagination) (property []models.RefPropertyV2, err error) {
	defer ctx.Session.EndSession(ctx.CTX)

	return s.Daos.PropertyWiseDemandandCollectionExcelV2(ctx, propertyfilter, pagination)

}

func (s *Service) PropertyWiseDemandandCollectionExcelV3(ctx *models.Context, filter *models.PropertyFilter, pagination *models.Pagination) (*excelize.File, error) {
	res, err := s.PropertyWiseDemandandCollectionExcelV2(ctx, filter, pagination)
	if err != nil {
		return nil, err
	}
	fmt.Println("'res length==>'", len(res))
	resPD, err1 := s.Daos.GetSingleDefaultProductConfiguration(ctx)
	if err1 != nil {
		return nil, err1
	}

	//  create an excel file
	excel := excelize.NewFile()
	sheet1 := "Property Wise Demand and Collection"
	rowNo := 1
	index := excel.NewSheet(sheet1)
	excel.SetActiveSheet(index)
	if resPD.LocationID == "Bhagalpur" {
		excel.MergeCell(sheet1, "A1", "X3")
		excel.MergeCell(sheet1, "A4", "X5")
		excel.MergeCell(sheet1, "A6", "X6")
		excel.MergeCell(sheet1, "A7", "X7")
	} else {
		excel.MergeCell(sheet1, "A1", "B5")
		excel.MergeCell(sheet1, "C1", "X3")
		excel.MergeCell(sheet1, "C4", "X5")
		excel.MergeCell(sheet1, "A6", "X6")
		excel.MergeCell(sheet1, "A7", "X7")
	}

	style1, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#B6DDE8"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	style2, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#FFDC6D"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	style4, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#FFDC6D"],"pattern":1},"alignment":{"horizontal":"right","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style1)

	if resPD.LocationID != "Bhagalpur" {
		documentUrl := s.Shared.Config.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.FILEURL)
		if err := excel.AddPicture(sheet1, fmt.Sprintf("%v%v", "A", rowNo), documentUrl+"municipal/logo.png", `{"x_scale": 0.115, "y_scale": 0.0935}`); err != nil {
			fmt.Println(err)
		}
	}

	if resPD.LocationID == "Bhagalpur" {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), ctx.ProductConfig.Name)
	} else {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), ctx.ProductConfig.Name)
	}

	rowNo++
	rowNo++
	rowNo++
	if resPD.LocationID == "Bhagalpur" {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "Property Wise Demand And Collection")
	} else {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Property Wise Demand And Collection")
	}
	rowNo++
	rowNo++

	t := time.Now()
	toDate := t.Format("02-January-2006")

	reportFromMsg2 := "Report"
	if filter != nil {
		if filter.AppliedRange != nil {
			if filter.AppliedRange.From != nil && filter.AppliedRange.To == nil {
				reportFromMsg2 = reportFromMsg2 + " on " + fmt.Sprintf("%v-%v-%v", filter.AppliedRange.From.Day(), filter.AppliedRange.From.Month(), filter.AppliedRange.From.Year())
			}
			if filter.AppliedRange.From != nil && filter.AppliedRange.To != nil {
				reportFromMsg2 = reportFromMsg2 + " From " + fmt.Sprintf("%v-%v-%v", filter.AppliedRange.From.Day(), filter.AppliedRange.From.Month(), filter.AppliedRange.From.Year()) + " To " + fmt.Sprintf("%v-%v-%v", filter.AppliedRange.To.Day(), filter.AppliedRange.To.Month(), filter.AppliedRange.To.Year())
			}
			if filter.AppliedRange.From == nil && filter.AppliedRange.To == nil {
				fmt.Println("everything is nil")
			}
		}
	}
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), reportFromMsg2)
	rowNo++

	reportFromMsg := "Report Generated on" + " " + toDate
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), reportFromMsg)
	rowNo++

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "X", rowNo), style2)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "S.No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "Ward No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "SAF No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), "Holding No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), "Owner Name")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), "Father/Husband Name/PAN")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), "Mobile No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "H", rowNo), "Road Type")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "I", rowNo), "Property Type")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "J", rowNo), "Address")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "K", rowNo), "Plot Area")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "L", rowNo), "Arrear Demand")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "M", rowNo), "Current Demand")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "N", rowNo), "Form Fee")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "O", rowNo), "Penalty")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "P", rowNo), "Advance")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "Q", rowNo), "Total Demand")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "R", rowNo), "Arrear Collection")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "S", rowNo), "Current Collection")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "T", rowNo), "Rebate")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "U", rowNo), "Total Collection")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "V", rowNo), "Arrear Outstanding Demand")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "W", rowNo), "Current Outstanding Demand")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "X", rowNo), "Total Outstanding Demand")
	fmt.Println("'res length==>'", len(res))
	var arrearDemand, currentDemand, totalDemand float64
	var arrearCollection, currentCollection, totalCollection float64
	var arrearOutstandingDemand, currentOutstandingDemand, totalOutstandingDemand float64

	for i, v := range res {

		v.Ref.Demand.Total.TotalTax = v.Ref.Demand.Total.TotalTax - v.Ref.Demand.Total.Ecess
		v.Ref.Collections.TotalTax = v.Ref.Collections.TotalTax + v.Ref.PropertyPayments.FormFee
		v.Ref.Demand.Current.TotalTax = v.Ref.Demand.Current.TotalTax + v.Ref.PropertyPayments.Rebate

		arrearDemand = arrearDemand + v.Ref.Demand.Arrear.TotalTax
		currentDemand = currentDemand + v.Ref.Demand.Current.TotalTax
		totalDemand = totalDemand + v.Ref.Demand.Total.TotalTax
		arrearCollection = arrearCollection + v.Ref.Collections.ArrearTax
		currentCollection = currentCollection + v.Ref.Collections.CurrentTax
		totalCollection = totalCollection + v.Ref.Collections.TotalTax
		arrearOutstandingDemand = arrearOutstandingDemand + (v.Ref.Demand.Arrear.TotalTax - v.Ref.Collections.ArrearTax - v.Ref.Collections.ArrearPenalty)
		currentOutstandingDemand = currentOutstandingDemand + (v.Ref.Demand.Current.TotalTax - v.Ref.Collections.CurrentTax - v.Ref.Collections.CurrentPenalty)
		totalOutstandingDemand = totalOutstandingDemand + (v.Ref.Demand.Total.TotalTax - v.Ref.Collections.TotalTax)

		rowNo++
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), i+1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), func() string {
			if v.Ref.Address.Ward.Name != "" {
				return v.Ref.Address.Ward.Name
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), v.ApplicationNo)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), v.UniqueID)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), func() string {
			if v.Ref.PropertyOwner != nil {
				if len(v.Ref.PropertyOwner) > 0 {
					return v.Ref.PropertyOwner[0].Name
				}
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), func() string {
			if v.Ref.PropertyOwner != nil {
				if len(v.Ref.PropertyOwner) > 0 {
					return v.Ref.PropertyOwner[0].FatherRpanRhusband
				}
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), func() string {
			if v.Ref.PropertyOwner != nil {
				if len(v.Ref.PropertyOwner) > 0 {
					return v.Ref.PropertyOwner[0].Mobile
				}
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "H", rowNo), func() string {
			if v.Ref.RoadType.Name != "" {
				return v.Ref.RoadType.Name
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "I", rowNo), func() string {
			if v.Ref.PropertyType.Name != "" {
				return v.Ref.PropertyType.Name
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "J", rowNo), v.Address.AL1+v.Address.Al2)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "K", rowNo), v.AreaOfPlot)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "L", rowNo), v.Ref.Demand.Arrear.TotalTax)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "M", rowNo), v.Ref.Demand.Current.TotalTax)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "N", rowNo), v.Ref.PropertyPayments.FormFee)

		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "O", rowNo), (v.Ref.Collections.ArrearPenalty + v.Ref.Collections.CurrentPenalty))
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "P", rowNo), v.Advance)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "Q", rowNo), v.Ref.Demand.Total.TotalTax)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "R", rowNo), v.Ref.Collections.ArrearTax)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "S", rowNo), v.Ref.Collections.CurrentTax)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "T", rowNo), v.Ref.PropertyPayments.Rebate)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "U", rowNo), v.Ref.Collections.TotalTax)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "V", rowNo), (v.Ref.Demand.Arrear.TotalTax - v.Ref.Collections.ArrearTax - v.Ref.Collections.ArrearPenalty))
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "W", rowNo), (v.Ref.Demand.Current.TotalTax - v.Ref.Collections.CurrentTax - v.Ref.Collections.CurrentPenalty))
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "X", rowNo), (v.Ref.Demand.Total.TotalTax - v.Ref.Collections.TotalTax))

	}
	rowNo++
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style2)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "Total")

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "B", rowNo), fmt.Sprintf("%v%v", "B", rowNo), style4)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style4)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "D", rowNo), fmt.Sprintf("%v%v", "D", rowNo), style4)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "E", rowNo), fmt.Sprintf("%v%v", "E", rowNo), style4)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "F", rowNo), fmt.Sprintf("%v%v", "F", rowNo), style4)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "G", rowNo), fmt.Sprintf("%v%v", "G", rowNo), style4)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "H", rowNo), fmt.Sprintf("%v%v", "H", rowNo), style4)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "I", rowNo), fmt.Sprintf("%v%v", "I", rowNo), style4)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "J", rowNo), fmt.Sprintf("%v%v", "J", rowNo), style4)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "K", rowNo), fmt.Sprintf("%v%v", "K", rowNo), style4)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "R", rowNo), fmt.Sprintf("%v%v", "R", rowNo), style4)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "S", rowNo), fmt.Sprintf("%v%v", "S", rowNo), style4)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "T", rowNo), fmt.Sprintf("%v%v", "T", rowNo), style4)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "U", rowNo), fmt.Sprintf("%v%v", "U", rowNo), style4)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "V", rowNo), fmt.Sprintf("%v%v", "V", rowNo), style4)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "W", rowNo), fmt.Sprintf("%v%v", "W", rowNo), style4)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "X", rowNo), fmt.Sprintf("%v%v", "X", rowNo), style4)

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "L", rowNo), fmt.Sprintf("%v%v", "L", rowNo), style4)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "L", rowNo), fmt.Sprintf("%.2f", arrearDemand))
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "M", rowNo), fmt.Sprintf("%v%v", "M", rowNo), style4)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "M", rowNo), fmt.Sprintf("%.2f", currentDemand))
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "Q", rowNo), fmt.Sprintf("%v%v", "Q", rowNo), style4)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "Q", rowNo), fmt.Sprintf("%.2f", totalDemand))
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "R", rowNo), fmt.Sprintf("%v%v", "R", rowNo), style4)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "R", rowNo), fmt.Sprintf("%.2f", arrearCollection))
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "S", rowNo), fmt.Sprintf("%v%v", "S", rowNo), style4)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "S", rowNo), fmt.Sprintf("%.2f", currentCollection))
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "U", rowNo), fmt.Sprintf("%v%v", "U", rowNo), style4)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "U", rowNo), fmt.Sprintf("%.2f", totalCollection))
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "V", rowNo), fmt.Sprintf("%v%v", "V", rowNo), style4)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "V", rowNo), fmt.Sprintf("%.2f", arrearOutstandingDemand))
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "W", rowNo), fmt.Sprintf("%v%v", "W", rowNo), style4)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "W", rowNo), fmt.Sprintf("%.2f", currentOutstandingDemand))
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "X", rowNo), fmt.Sprintf("%v%v", "X", rowNo), style4)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "X", rowNo), fmt.Sprintf("%.2f", totalOutstandingDemand))
	return excel, nil
}

func (s *Service) PropertyWiseDemandCollectionandBalanceReport(ctx *models.Context, propertyfilter *models.PropertyFilter, pagination *models.Pagination) (property []models.RefProperty, err error) {
	defer ctx.Session.EndSession(ctx.CTX)

	return s.Daos.PropertyWiseDemandCollectionandBalanceReport(ctx, propertyfilter, pagination)

}

func (s *Service) PropertyWiseDemandCollectionandBalanceReportExcel(ctx *models.Context, filter *models.PropertyFilter, pagination *models.Pagination) (*excelize.File, error) {
	res, err := s.Daos.PropertyWiseDemandCollectionandBalanceReport(ctx, filter, pagination)
	if err != nil {
		return nil, err
	}
	fmt.Println("'res length==>'", len(res))
	excel := excelize.NewFile()
	sheet1 := "PropertyWiseDemandcollectionandBalance"
	rowNo := 1
	index := excel.NewSheet(sheet1)
	excel.SetActiveSheet(index)
	excel.MergeCell(sheet1, "A1", "B5")
	excel.MergeCell(sheet1, "C1", "T3")
	excel.MergeCell(sheet1, "C4", "T5")
	excel.MergeCell(sheet1, "A6", "T6")
	//excel.MergeCell(sheet1, "A7", "M9")
	excel.MergeCell(sheet1, "A7", "A8")
	excel.MergeCell(sheet1, "B7", "B8")
	excel.MergeCell(sheet1, "C7", "C8")
	excel.MergeCell(sheet1, "D7", "D8")
	excel.MergeCell(sheet1, "E7", "E8")
	excel.MergeCell(sheet1, "F7", "F8")
	excel.MergeCell(sheet1, "G7", "G8")
	excel.MergeCell(sheet1, "H7", "J7")
	excel.MergeCell(sheet1, "K7", "P7")
	excel.MergeCell(sheet1, "Q7", "T7")
	// excel.MergeCell(sheet1, "H8", "J8")
	// excel.MergeCell(sheet1, "K8", "P8")
	// excel.MergeCell(sheet1, "Q8", "T8")
	// excel.MergeCell(sheet1, "J8", "K8")
	// excel.MergeCell(sheet1, "L7", "L9")
	style1, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#FFDC6D"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	style3, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#B6DDE8"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	// style2, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#B6DDE8"],"pattern":1},"font":{"bold":true}}`)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	style4, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#FFDC6D"],"pattern":1},"alignment":{"horizontal":"right","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style3)
	documentUrl := s.Shared.Config.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.FILEURL)
	if err := excel.AddPicture(sheet1, fmt.Sprintf("%v%v", "A", rowNo), documentUrl+"municipal/logo.png", `{"x_scale": 0.115, "y_scale": 0.0935}`); err != nil {
		fmt.Println(err)
	}
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), s.Shared.Config.GetString(s.Shared.GetCmdArg(constants.ENV)+"."+constants.LOCATIONNAME))
	rowNo++
	rowNo++
	rowNo++

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style3)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Property DCB Report")
	rowNo++
	rowNo++
	//
	reportFromMsg := "DCB Report"

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style3)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), reportFromMsg)
	rowNo++
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "G", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "H", rowNo), fmt.Sprintf("%v%v", "J", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "H", rowNo), "Demand")
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "K", rowNo), fmt.Sprintf("%v%v", "P", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "K", rowNo), "Collection")
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "Q", rowNo), fmt.Sprintf("%v%v", "T", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "Q", rowNo), "Balance")
	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "J", rowNo), fmt.Sprintf("%v%v", "K", rowNo), style1)
	// excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "J", rowNo), "DD")
	rowNo++
	//
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "T", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "S.No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "HoldingNo")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "ApplicationNo")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), "OldHoldingNo")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), "Ward")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), "OwnerName")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), "MobileNo")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "H", rowNo), "Arrer")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "I", rowNo), "Current")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "J", rowNo), "Total")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "K", rowNo), "Tax")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "L", rowNo), "Rebate")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "M", rowNo), "Penalty")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "N", rowNo), "Form Fee")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "O", rowNo), "Boring Charge")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "P", rowNo), "Total")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "Q", rowNo), "Tax")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "R", rowNo), "Rebate")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "S", rowNo), "Penalty")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "T", rowNo), "Total")

	fmt.Println("'res length==>'", len(res))
	// var amount float64
	// var count float64
	var totalTopayAmount float64
	var totalArrerAmount float64
	var totalCurrentAmount float64

	for i, v := range res {
		// amount = v.
		// count = v.TradeLicensePayments.Cash.Count + v.TradeLicensePayments.Cheque.Count + v.TradeLicensePayments.NetBanking.Count + v.TradeLicensePayments.DD.Count

		rowNo++
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), i+1)
		//excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), sd)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v.UniqueID)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), v.ApplicationNo)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), v.OldHoldingNumber)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), v.Ref.Address.Ward.Name)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), func() string {
			if v.Ref.PropertyOwner != nil {
				if len(v.Ref.PropertyOwner) > 0 {
					return v.Ref.PropertyOwner[0].Name
				}
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), func() string {
			if v.Ref.PropertyOwner != nil {
				if len(v.Ref.PropertyOwner) > 0 {
					return v.Ref.PropertyOwner[0].Mobile
				}
			}
			return "NA"
		}())
		// excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), v.Ref.Basic.)
		// excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), v.AreaOfPlot)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "H", rowNo), v.Sumary.Demand.Arrear.TotalTax)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "I", rowNo), v.Sumary.Demand.Current.TotalTax)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "J", rowNo), v.Sumary.Demand.Total.TotalTax)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "K", rowNo), v.Sumary.Collections.Total.TotalTax)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "L", rowNo), v.Sumary.Collections.Total.Rebate)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "M", rowNo), v.Sumary.Collections.Total.Penalty)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "N", rowNo), v.Sumary.Collections.Total.FormFee)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "O", rowNo), v.Sumary.Collections.Total.BoringCharge)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "P", rowNo), v.Sumary.Collections.Total.TotalTax)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "Q", rowNo), v.Sumary.ToPay.Total.TotalTax)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "R", rowNo), v.Sumary.ToPay.Total.Rebate)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "S", rowNo), v.Sumary.ToPay.Total.Penalty)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "T", rowNo), v.Sumary.ToPay.Total.Total)

		totalTopayAmount = totalTopayAmount + v.Sumary.ToPay.Total.TotalTax
		totalArrerAmount = totalArrerAmount + v.Sumary.Demand.Total.TotalTax
		totalCurrentAmount = totalCurrentAmount + v.Sumary.Collections.Total.TotalTax
	}

	rowNo++
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "T", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "Total")

	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "B", rowNo), fmt.Sprintf("%v%v", "B", rowNo), style1)
	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style1)
	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "D", rowNo), fmt.Sprintf("%v%v", "D", rowNo), style1)
	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "E", rowNo), fmt.Sprintf("%v%v", "E", rowNo), style1)
	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "F", rowNo), fmt.Sprintf("%v%v", "E", rowNo), style1)
	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "G", rowNo), fmt.Sprintf("%v%v", "G", rowNo), style1)
	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "H", rowNo), fmt.Sprintf("%v%v", "H", rowNo), style1)
	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "I", rowNo), fmt.Sprintf("%v%v", "I", rowNo), style1)
	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "J", rowNo), fmt.Sprintf("%v%v", "J", rowNo), style1)
	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "K", rowNo), fmt.Sprintf("%v%v", "K", rowNo), style1)
	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "L", rowNo), fmt.Sprintf("%v%v", "L", rowNo), style1)
	// //excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "M", rowNo), fmt.Sprintf("%v%v", "M", rowNo), style1)

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "J", rowNo), fmt.Sprintf("%v%v", "J", rowNo), style4)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "J", rowNo), fmt.Sprintf("%.0f", totalArrerAmount))
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "P", rowNo), fmt.Sprintf("%v%v", "P", rowNo), style4)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "P", rowNo), fmt.Sprintf("%.0f", totalCurrentAmount))
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "T", rowNo), fmt.Sprintf("%v%v", "T", rowNo), style4)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "T", rowNo), fmt.Sprintf("%.0f", totalTopayAmount))

	return excel, nil
}
