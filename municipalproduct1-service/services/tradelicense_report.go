package services

import (
	"errors"
	"fmt"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
)

// DashboardDayWiseTradelicenseCollectionChart : ""
func (s *Service) DashboardDayWiseTradelicenseCollectionChart(ctx *models.Context, filter *models.DashboardDayWiseTradeLicenseCollectionChartFilter) (models.DashboardDayWiseTradeLicenseCollectionChart, error) {
	return s.Daos.DashboardDayWiseTradelicenseCollectionChart(ctx, filter)
}

// DashboardDayWiseTradeLicenseCollectionChartExcel: ""
func (s *Service) DashboardDayWiseTradeLicenseCollectionChartExcel(ctx *models.Context, filter *models.DashboardDayWiseTradeLicenseCollectionChartFilter) (*excelize.File, error) {
	data, err := s.DashboardDayWiseTradelicenseCollectionChart(ctx, filter)
	if err != nil {
		return nil, err
	}
	excel := excelize.NewFile()
	sheet1 := "Trade License Day Wise Collection"
	index := excel.NewSheet(sheet1)
	rowNo := 1

	excel.SetActiveSheet(index)
	excel.MergeCell(sheet1, "A1", "B5")
	excel.MergeCell(sheet1, "C1", "D3")
	excel.MergeCell(sheet1, "C4", "D5")
	excel.MergeCell(sheet1, "A6", "D6")
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

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "D", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "Date")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "License Count")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Amount")
	rowNo++
	var totalAmount float64

	for _, v := range data.Records {
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), v.ID)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v.TradeLicenseCount)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), v.Amount)
		totalAmount = totalAmount + v.Amount
		rowNo++

	}

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "D", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "Total")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), totalAmount)

	return excel, nil

}

// DashboardDayWiseTradeLicenseCollectionChartPDF : ""
func (s *Service) DashboardDayWiseTradeLicenseCollectionChartPDF(ctx *models.Context, filter *models.DashboardDayWiseTradeLicenseCollectionChartFilter) ([]byte, error) {
	data, err := s.DashboardDayWiseTradelicenseCollectionChart(ctx, filter)
	if err != nil {
		return nil, err
	}

	m := make(map[string]interface{})
	m2 := make(map[string]interface{})
	m["demand"] = data
	m2["currentDate"] = time.Now()

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
	templatePath := templatePathStart + "tradelicense_daywise_collection_report.html"
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

// DayWiseTradeLicenseDemandChart : ""
func (s *Service) DayWiseTradeLicenseDemandChart(ctx *models.Context, filter *models.DayWiseTradeLicenseDemandChartFilter) (*models.DayWiseTradeLicenseDemandChart, error) {
	res, err := s.Daos.DayWiseTradeLicenseDemandChart(ctx, filter)
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

// DayWiseTradeLicenseDemandReportPDF : ""
func (s *Service) DayWiseTradeLicenseDemandReportPDF(ctx *models.Context, filter *models.DayWiseTradeLicenseDemandChartFilter) ([]byte, error) {
	properties, err := s.DayWiseTradeLicenseDemandChart(ctx, filter)
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
	templatePath := templatePathStart + "tradelicense_daywise_demand_report.html"
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

// DayWiseTradeLicenseDemandReportExcel: ""
func (s *Service) DayWiseTradeLicenseDemandReportExcel(ctx *models.Context, filter *models.DayWiseTradeLicenseDemandChartFilter) (*excelize.File, error) {
	data, err := s.DayWiseTradeLicenseDemandChart(ctx, filter)
	if err != nil {
		return nil, err
	}
	excel := excelize.NewFile()
	sheet1 := "Trade License Day Wise Demand"
	index := excel.NewSheet(sheet1)
	rowNo := 1

	excel.SetActiveSheet(index)
	excel.MergeCell(sheet1, "A1", "B5")
	excel.MergeCell(sheet1, "C1", "C3")
	excel.MergeCell(sheet1, "C4", "C5")
	excel.MergeCell(sheet1, "A6", "C6")
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

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "Date")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "Trade License Count")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Amount")
	rowNo++
	var totalAmount float64
	if data != nil {

		for _, v := range data.Records {
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), v.ID)
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v.TradeLicenseCount)
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

// TradeLicenseOverallDemandReportJSON : ""
func (s *Service) TradeLicenseOverallDemandReportJSON(ctx *models.Context, filter *models.TradeLicenseFilter, p *models.Pagination) (*models.TradeLicenseOverallDemandReport, error) {
	tradelicenses, err := s.Daos.FilterTradeLicense(ctx, filter, p)
	if err != nil {
		return nil, err
	}

	report := new(models.TradeLicenseOverallDemandReport)
	cfy, err := s.Daos.GetCurrentFinancialYear(ctx)
	if err != nil {
		return nil, errors.New("Error in getting current financial year " + err.Error())
	}
	report.TradeLicenses = tradelicenses
	if cfy != nil {
		report.CFY = *cfy
	}
	return report, nil
}

// TradeLicenseOverallDemandReportExcel : ""
func (s *Service) TradeLicenseOverallDemandReportExcel(ctx *models.Context, filter *models.TradeLicenseFilter, pagination *models.Pagination) (*excelize.File, error) {
	res, err := s.FilterTradeLicense(ctx, filter, pagination)
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
	sheet1 := "TradeLicenseOverallDemand"
	rowNo := 1
	index := excel.NewSheet(sheet1)
	excel.SetActiveSheet(index)
	excel.MergeCell(sheet1, "A1", "B5")
	excel.MergeCell(sheet1, "C1", "F3")
	excel.MergeCell(sheet1, "C4", "F5")
	excel.MergeCell(sheet1, "A6", "F6")

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
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "TradeLicense Overall Demand")
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
			if v.LicenseDate != nil && cfy != nil && v.LicenseExpiryDate != nil {
				return fmt.Sprintf("%v - %v", v.LicenseDate.Year(), v.LicenseExpiryDate.Year())
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), fmt.Sprintf("%.2f", v.Demand.Total.Total))
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), v.OwnerName)

	}
	var totalTax float64
	for _, v1 := range res {
		totalTax = totalTax + v1.Demand.Total.Total
	}

	rowNo++
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "Total")

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "B", rowNo), fmt.Sprintf("%v%v", "B", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "D", rowNo), fmt.Sprintf("%v%v", "D", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "E", rowNo), fmt.Sprintf("%v%v", "E", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), totalTax)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), fmt.Sprintf("%.2f", totalTax))

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "F", rowNo), fmt.Sprintf("%v%v", "F", rowNo), style1)

	return excel, nil
}

// /TradeLicenseOverallDemandReportPDF : ""
func (s *Service) TradeLicenseOverallDemandReportPDF(ctx *models.Context, filter *models.TradeLicenseFilter, p *models.Pagination) ([]byte, error) {
	properties, err := s.Daos.FilterTradeLicense(ctx, filter, p)
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
	m2["inc"] = func(a int) int {
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
	templatePath := templatePathStart + "tradelicense_demand_report.html"
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

// FilterWardDayWiseTradeLicenseCollectionReport : ""
func (s *Service) FilterWardDayWiseTradeLicenseCollectionReport(ctx *models.Context, filter *models.WardDayWiseTradeLicenseCollectionReportFilter, pagination *models.Pagination) ([]models.WardDayWiseTradeLicenseCollectionReport, error) {
	return s.Daos.FilterWardDayWiseTradeLicenseCollectionReport(ctx, filter, pagination)
}

// FilterWardDayWiseTradeLicenseCollectionReportExcel: ""
func (s *Service) FilterWardDayWiseTradeLicenseCollectionReportExcel(ctx *models.Context, filter *models.WardDayWiseTradeLicenseCollectionReportFilter, pagination *models.Pagination) (*excelize.File, error) {
	data, err := s.FilterWardDayWiseTradeLicenseCollectionReport(ctx, filter, pagination)
	if err != nil {
		return nil, err
	}
	excel := excelize.NewFile()
	sheet1 := "Trade License Ward Wise Collection"
	index := excel.NewSheet(sheet1)
	rowNo := 1

	excel.SetActiveSheet(index)
	excel.MergeCell(sheet1, "A1", "B5")
	excel.MergeCell(sheet1, "C1", "E3")
	excel.MergeCell(sheet1, "C4", "E5")
	excel.MergeCell(sheet1, "A6", "E6")
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

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "E", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "S.No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "Ward")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Paied Trade Licenses")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), "No of Payments")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), "Collection Amount")
	rowNo++
	var totalAmount float64
	for i, v := range data {
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), i+1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v.Ward.Name)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), v.Report.TotalNoProperties)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), v.Report.TotalNoPayments)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), v.Report.TotalCollections)
		totalAmount = totalAmount + v.Report.TotalCollections
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

// FilterWardDayWiseTradeLicenseCollectionReportPDF : ""
func (s *Service) FilterWardDayWiseTradeLicenseCollectionReportPDF(ctx *models.Context, filter *models.WardDayWiseTradeLicenseCollectionReportFilter, pagination *models.Pagination) ([]byte, error) {
	data, err := s.FilterWardDayWiseTradeLicenseCollectionReport(ctx, filter, pagination)
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
	templatePath := templatePathStart + "tradelicense_wardwise_collection_report.html"
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

// FilterWardMonthWiseTradeLicenseCollectionReport : ""
func (s *Service) FilterWardMonthWiseTradeLicenseCollectionReport(ctx *models.Context, filter *models.WardMonthWiseTradeLicenseCollectionReportFilter, pagination *models.Pagination) ([]models.WardMonthWiseTradeLicenseCollectionReport, error) {
	return s.Daos.FilterWardMonthWiseTradeLicenseCollectionReport(ctx, filter, pagination)
}

// FilterWardMonthWiseTradeLicenseCollectionReport: ""
func (s *Service) FilterWardMonthWiseTradeLicenseCollectionReportExcel(ctx *models.Context, filter *models.WardMonthWiseTradeLicenseCollectionReportFilter, pagination *models.Pagination) (*excelize.File, error) {
	data, err := s.FilterWardMonthWiseTradeLicenseCollectionReport(ctx, filter, pagination)
	if err != nil {
		return nil, err
	}
	excel := excelize.NewFile()
	sheet1 := "Trade License Ward Wise Collection"
	index := excel.NewSheet(sheet1)
	rowNo := 1

	excel.SetActiveSheet(index)
	excel.MergeCell(sheet1, "A1", "B5")
	excel.MergeCell(sheet1, "C1", "E3")
	excel.MergeCell(sheet1, "C4", "E5")
	excel.MergeCell(sheet1, "A6", "E6")
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

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "E", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "S.No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "Ward")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Paied Trade Licenses")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), "No of Payments")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), "Collection Amount")
	rowNo++
	var totalAmount float64
	for i, v := range data {
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), i+1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v.Ward.Name)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), v.Report.TotalNoProperties)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), v.Report.TotalNoPayments)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), v.Report.TotalCollections)
		totalAmount = totalAmount + v.Report.TotalCollections
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

// FilterWardDayWiseTradeLicenseCollectionReportPDF : ""
func (s *Service) FilterWardMonthWiseTradeLicenseCollectionReportPDF(ctx *models.Context, filter *models.WardMonthWiseTradeLicenseCollectionReportFilter, pagination *models.Pagination) ([]byte, error) {
	data, err := s.FilterWardMonthWiseTradeLicenseCollectionReport(ctx, filter, pagination)
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
	templatePath := templatePathStart + "tradelicense_wardwisemonth_collection_report.html"
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

// FilterWardDayWiseTradeLicenseDemandReport : ""
func (s *Service) FilterWardDayWiseTradeLicenseDemandReport(ctx *models.Context, filter *models.WardDayWiseTradeLicenseDemandReportFilter, pagination *models.Pagination) ([]models.WardDayWiseTradeLicenseDemandReport, error) {
	return s.Daos.FilterWardDayWiseTradeLicenseDemandReport(ctx, filter, pagination)
}

// DayWiseTradeLicenseDemandReportExcel: ""
func (s *Service) FilterWardDayWiseTradeLicenseDemandReportExcel(ctx *models.Context, filter *models.WardDayWiseTradeLicenseDemandReportFilter, pagination *models.Pagination) (*excelize.File, error) {
	data, err := s.FilterWardDayWiseTradeLicenseDemandReport(ctx, filter, pagination)
	if err != nil {
		return nil, err
	}
	excel := excelize.NewFile()
	sheet1 := "Trade License Ward Wise Demand"
	index := excel.NewSheet(sheet1)
	rowNo := 1

	excel.SetActiveSheet(index)
	excel.MergeCell(sheet1, "A1", "B5")
	excel.MergeCell(sheet1, "C1", "C3")
	excel.MergeCell(sheet1, "C4", "C5")
	excel.MergeCell(sheet1, "A6", "C6")

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

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "S.No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "No of Trade Licenses")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Amount")
	rowNo++
	var totalAmount float64
	for i, v := range data {
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), i+1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v.Report.TradeLicenses)
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

// FilterWardDayWiseTradeLicenseDemandReport : ""
func (s *Service) FilterWardDayWiseTradeLicenseDemandReportPDF(ctx *models.Context, filter *models.WardDayWiseTradeLicenseDemandReportFilter, pagination *models.Pagination) ([]byte, error) {
	data, err := s.FilterWardDayWiseTradeLicenseDemandReport(ctx, filter, pagination)
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
	templatePath := templatePathStart + "tradelicense_wardwise_demand_report.html"
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

// FilterWardMonthWiseTradeLicenseDemandReport : ""
func (s *Service) FilterWardMonthWiseTradeLicenseDemandReport(ctx *models.Context, filter *models.WardMonthWiseTradeLicenseDemandReportFilter, pagination *models.Pagination) ([]models.WardMonthWiseTradeLicenseDemandReport, error) {
	return s.Daos.FilterWardMonthWiseTradeLicenseDemandReport(ctx, filter, pagination)
}

// DayWiseTradeLicenseDemandReportExcel: ""
func (s *Service) FilterWardMonthWiseTradeLicenseDemandReportExcel(ctx *models.Context, filter *models.WardMonthWiseTradeLicenseDemandReportFilter, pagination *models.Pagination) (*excelize.File, error) {
	data, err := s.FilterWardMonthWiseTradeLicenseDemandReport(ctx, filter, pagination)
	if err != nil {
		return nil, err
	}
	excel := excelize.NewFile()
	sheet1 := "Trade License Ward Wise Demand"
	index := excel.NewSheet(sheet1)
	rowNo := 1

	excel.SetActiveSheet(index)
	excel.MergeCell(sheet1, "A1", "B5")
	excel.MergeCell(sheet1, "C1", "C3")
	excel.MergeCell(sheet1, "C4", "C5")
	excel.MergeCell(sheet1, "A6", "C6")
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

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "S.No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "No of Trade Licenses")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Amount")
	rowNo++
	var totalAmount float64
	for i, v := range data {
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), i+1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v.Report.TradeLicenses)
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

// FilterWardMonthWiseTradeLicenseDemandReportPDF : ""
func (s *Service) FilterWardMonthWiseTradeLicenseDemandReportPDF(ctx *models.Context, filter *models.WardMonthWiseTradeLicenseDemandReportFilter, pagination *models.Pagination) ([]byte, error) {
	data, err := s.FilterWardMonthWiseTradeLicenseDemandReport(ctx, filter, pagination)
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
	templatePath := templatePathStart + "tradelicense_wardwisemonth_demand_report.html"
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

// FilterTeamDayWiseTradeLicenseCollectionReport : ""
func (s *Service) FilterTeamDayWiseTradeLicenseCollectionReport(ctx *models.Context, filter *models.TeamDayWiseTradeLicenseCollectionReportFilter, pagination *models.Pagination) ([]models.TeamDayWiseTradeLicenseCollectionReport, error) {
	return s.Daos.FilterTeamDayWiseTradeLicenseCollectionReport(ctx, filter, pagination)
}

// FilterTeamDayWiseTradeLicenseCollectionReportExcel: ""
func (s *Service) FilterTeamDayWiseTradeLicenseCollectionReportExcel(ctx *models.Context, filter *models.TeamDayWiseTradeLicenseCollectionReportFilter, pagination *models.Pagination) (*excelize.File, error) {
	data, err := s.FilterTeamDayWiseTradeLicenseCollectionReport(ctx, filter, pagination)
	if err != nil {
		return nil, err
	}
	excel := excelize.NewFile()
	sheet1 := "Trade License Team Wise Collection"
	index := excel.NewSheet(sheet1)
	rowNo := 1

	excel.SetActiveSheet(index)
	excel.MergeCell(sheet1, "A1", "B5")
	excel.MergeCell(sheet1, "C1", "E3")
	excel.MergeCell(sheet1, "C4", "E5")
	excel.MergeCell(sheet1, "A6", "E6")
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

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "E", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "S.No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "User")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Paied Trade Licenses")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), "No of Payments")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), "Collection Amount")
	rowNo++
	var totalAmount float64
	for i, v := range data {
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), i+1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v.User.Name)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), v.Report.TotalNoTradeLicenses)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), v.Report.TotalNoPayments)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), v.Report.TotalCollections)
		totalAmount = totalAmount + v.Report.TotalCollections
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

// FilterTeamDayWiseTradeLicenseCollectionReportPDF : ""
func (s *Service) FilterTeamDayWiseTradeLicenseCollectionReportPDF(ctx *models.Context, filter *models.TeamDayWiseTradeLicenseCollectionReportFilter, pagination *models.Pagination) ([]byte, error) {
	data, err := s.FilterTeamDayWiseTradeLicenseCollectionReport(ctx, filter, pagination)
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
	templatePath := templatePathStart + "tradelicense_teamwise_collection_report.html"
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

// FilterTeamMonthWiseTradeLicenseCollectionReport : ""
func (s *Service) FilterTeamMonthWiseTradeLicenseCollectionReport(ctx *models.Context, filter *models.TeamMonthWiseTradeLicenseCollectionReportFilter, pagination *models.Pagination) ([]models.TeamMonthWiseTradeLicenseCollectionReport, error) {
	return s.Daos.FilterTeamMonthWiseTradeLicenseCollectionReport(ctx, filter, pagination)
}

// FilterTeamMonthWiseTradeLicenseCollectionReportExcel: ""
func (s *Service) FilterTeamMonthWiseTradeLicenseCollectionReportExcel(ctx *models.Context, filter *models.TeamMonthWiseTradeLicenseCollectionReportFilter, pagination *models.Pagination) (*excelize.File, error) {
	data, err := s.FilterTeamMonthWiseTradeLicenseCollectionReport(ctx, filter, pagination)
	if err != nil {
		return nil, err
	}
	excel := excelize.NewFile()
	sheet1 := "Trade License Team Wise Collection"
	index := excel.NewSheet(sheet1)
	rowNo := 1

	excel.SetActiveSheet(index)
	excel.MergeCell(sheet1, "A1", "B5")
	excel.MergeCell(sheet1, "C1", "E3")
	excel.MergeCell(sheet1, "C4", "E5")
	excel.MergeCell(sheet1, "A6", "E6")
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

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "E", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "S.No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "User")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Paied Trade Licenses")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), "No of Payments")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), "Collection Amount")
	rowNo++
	var totalAmount float64
	for i, v := range data {
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), i+1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v.User.Name)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), v.Report.TotalNoTradeLicenses)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), v.Report.TotalNoPayments)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), v.Report.TotalCollections)
		totalAmount = totalAmount + v.Report.TotalCollections
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

// FilterTeamMonthWiseTradeLicenseCollectionReportPDF : ""
func (s *Service) FilterTeamMonthWiseTradeLicenseCollectionReportPDF(ctx *models.Context, filter *models.TeamMonthWiseTradeLicenseCollectionReportFilter, pagination *models.Pagination) ([]byte, error) {
	data, err := s.FilterTeamMonthWiseTradeLicenseCollectionReport(ctx, filter, pagination)
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
	templatePath := templatePathStart + "tradelicense_teamwisemonth_collection_report.html"
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

// FilterWardYearWiseTradeLicenseCollectionReport : ""
func (s *Service) FilterWardYearWiseTradeLicenseCollectionReport(ctx *models.Context, filter *models.WardYearWiseTradeLicenseCollectionReportFilter, pagination *models.Pagination) ([]models.WardYearWiseTradeLicenseCollectionReport, error) {
	return s.Daos.FilterWardYearWiseTradeLicenseCollectionReport(ctx, filter, pagination)
}

// FilterWardYearWiseTradeLicenseCollectionReportExcel: ""
func (s *Service) FilterWardYearWiseTradeLicenseCollectionReportExcel(ctx *models.Context, filter *models.WardYearWiseTradeLicenseCollectionReportFilter, pagination *models.Pagination) (*excelize.File, error) {
	data, err := s.FilterWardYearWiseTradeLicenseCollectionReport(ctx, filter, pagination)
	if err != nil {
		return nil, err
	}
	excel := excelize.NewFile()
	sheet1 := "Trade License Ward Wise Collection"
	index := excel.NewSheet(sheet1)
	rowNo := 1

	excel.SetActiveSheet(index)
	excel.MergeCell(sheet1, "A1", "B5")
	excel.MergeCell(sheet1, "C1", "E3")
	excel.MergeCell(sheet1, "C4", "E5")
	excel.MergeCell(sheet1, "A6", "E6")
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

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "E", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "S.No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "Ward")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Paied Trade Licenses")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), "No of Payments")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), "Collection Amount")
	rowNo++
	var totalAmount float64
	for i, v := range data {
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), i+1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v.Ward.Name)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), v.Report.TotalNoProperties)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), v.Report.TotalNoPayments)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), v.Report.TotalCollections)
		totalAmount = totalAmount + v.Report.TotalCollections
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

// FilterWardYearWiseTradeLicenseCollectionReportPDF : ""
func (s *Service) FilterWardYearWiseTradeLicenseCollectionReportPDF(ctx *models.Context, filter *models.WardYearWiseTradeLicenseCollectionReportFilter, pagination *models.Pagination) ([]byte, error) {
	data, err := s.FilterWardYearWiseTradeLicenseCollectionReport(ctx, filter, pagination)
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
	templatePath := templatePathStart + "tradelicense_wardwiseyear_collection_report.html"
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

// FilterWardYearWiseTradeLicenseDemandReport : ""
func (s *Service) FilterWardYearWiseTradeLicenseDemandReport(ctx *models.Context, filter *models.WardYearWiseTradeLicenseDemandReportFilter, pagination *models.Pagination) ([]models.WardYearWiseTradeLicenseDemandReport, error) {
	return s.Daos.FilterWardYearWiseTradeLicenseDemandReport(ctx, filter, pagination)
}

// FilterWardYearWiseTradeLicenseDemandReportExcel: ""
func (s *Service) FilterWardYearWiseTradeLicenseDemandReportExcel(ctx *models.Context, filter *models.WardYearWiseTradeLicenseDemandReportFilter, pagination *models.Pagination) (*excelize.File, error) {
	data, err := s.FilterWardYearWiseTradeLicenseDemandReport(ctx, filter, pagination)
	if err != nil {
		return nil, err
	}
	excel := excelize.NewFile()
	sheet1 := "Trade License Ward Wise Demand"
	index := excel.NewSheet(sheet1)
	rowNo := 1

	excel.SetActiveSheet(index)
	excel.MergeCell(sheet1, "A1", "B5")
	excel.MergeCell(sheet1, "C1", "C3")
	excel.MergeCell(sheet1, "C4", "C5")
	excel.MergeCell(sheet1, "A6", "C6")
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

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "S.No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "No of Trade Licenses")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Amount")
	rowNo++
	var totalAmount float64
	for i, v := range data {
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), i+1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v.Report.TradeLicenses)
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

// FilterWardYearWiseTradeLicenseDemandReportPDF : ""
func (s *Service) FilterWardYearWiseTradeLicenseDemandReportPDF(ctx *models.Context, filter *models.WardYearWiseTradeLicenseDemandReportFilter, pagination *models.Pagination) ([]byte, error) {
	data, err := s.FilterWardYearWiseTradeLicenseDemandReport(ctx, filter, pagination)
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
	templatePath := templatePathStart + "tradelicense_wardwiseyear_demand_report.html"
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

// FilterTeamYearWiseTradeLicenseCollectionReport : ""
func (s *Service) FilterTeamYearWiseTradeLicenseCollectionReport(ctx *models.Context, filter *models.TeamYearWiseTradeLicenseCollectionReportFilter, pagination *models.Pagination) ([]models.TeamYearWiseTradeLicenseCollectionReport, error) {
	return s.Daos.FilterTeamYearWiseTradeLicenseCollectionReport(ctx, filter, pagination)
}

// FilterTeamYearWiseTradeLicenseCollectionReportExcel: ""
func (s *Service) FilterTeamYearWiseTradeLicenseCollectionReportExcel(ctx *models.Context, filter *models.TeamYearWiseTradeLicenseCollectionReportFilter, pagination *models.Pagination) (*excelize.File, error) {
	data, err := s.FilterTeamYearWiseTradeLicenseCollectionReport(ctx, filter, pagination)
	if err != nil {
		return nil, err
	}
	excel := excelize.NewFile()
	sheet1 := "Trade License Team Wise Collection"
	index := excel.NewSheet(sheet1)
	rowNo := 1

	excel.SetActiveSheet(index)
	excel.MergeCell(sheet1, "A1", "B5")
	excel.MergeCell(sheet1, "C1", "E3")
	excel.MergeCell(sheet1, "C4", "E5")
	excel.MergeCell(sheet1, "A6", "E6")
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

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "E", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "S.No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "User")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Paied TradeLicenses")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), "No of Payments")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), "Collection Amount")
	rowNo++
	var totalAmount float64
	for i, v := range data {
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), i+1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v.User.Name)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), v.Report.TotalNoTradeLicenses)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), v.Report.TotalNoPayments)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), v.Report.TotalCollections)
		totalAmount = totalAmount + v.Report.TotalCollections
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

// FilterTeamYearWiseTradeLicenseCollectionReportPDF : ""
func (s *Service) FilterTeamYearWiseTradeLicenseCollectionReportPDF(ctx *models.Context, filter *models.TeamYearWiseTradeLicenseCollectionReportFilter, pagination *models.Pagination) ([]byte, error) {
	data, err := s.FilterTeamYearWiseTradeLicenseCollectionReport(ctx, filter, pagination)
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
	templatePath := templatePathStart + "tradelicense_teamwiseyear_collection_report.html"
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
