package services

import (
	"errors"
	"fmt"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func (s *Service) WardWiseShoprentReport(ctx *models.Context, filter *models.WardWiseShoprentReportFilter) ([]models.WardWiseShoprentReport, error) {
	return s.Daos.WardWiseShoprentReport(ctx, filter)
}

// ShopRentOverallDemandReportJSON : ""
func (s *Service) ShopRentOverallDemandReportJSON(ctx *models.Context, filter *models.ShopRentFilter, p *models.Pagination) (*models.ShopRentOverallDemandReport, error) {
	shoprents, err := s.Daos.FilterShopRent(ctx, filter, p)
	if err != nil {
		return nil, err
	}

	report := new(models.ShopRentOverallDemandReport)
	cfy, err := s.Daos.GetCurrentFinancialYear(ctx)
	if err != nil {
		return nil, errors.New("Error in getting current financial year " + err.Error())
	}
	report.ShopRents = shoprents
	if cfy != nil {
		report.CFY = *cfy
	}
	return report, nil
}

// ShopRentOverallDemandReportExcel : ""
func (s *Service) ShopRentOverallDemandReportExcel(ctx *models.Context, filter *models.ShopRentFilter, pagination *models.Pagination) (*excelize.File, error) {
	res, err := s.FilterShopRent(ctx, filter, pagination)
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
	sheet1 := "ShopRentOverallDemand"
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
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "ShopRent Overall Demand")
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
			if v.DateFrom != nil && cfy != nil && cfy.To != nil {
				return fmt.Sprintf("%v - %v", v.DateFrom.Year(), cfy.To.Year())
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

// /ShopRentOverallDemandReportPDF : ""
func (s *Service) ShopRentOverallDemandReportPDF(ctx *models.Context, filter *models.ShopRentFilter, p *models.Pagination) ([]byte, error) {
	properties, err := s.Daos.FilterShopRent(ctx, filter, p)
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
	templatePath := templatePathStart + "shoprent_demand_report.html"
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

func (s *Service) DashboardDayWiseShoprentCollectionChart(ctx *models.Context, filter *models.DashboardDayWiseShoprentCollectionChartFilter) (models.DashboardDayWiseShoprentCollectionChart, error) {
	return s.Daos.DashboardDayWiseShoprentCollectionChart(ctx, filter)
}

func (s *Service) DashboardDayWiseShoprentCollectionChartPDF(ctx *models.Context, filter *models.DashboardDayWiseShoprentCollectionChartFilter) ([]byte, error) {
	data, err := s.DashboardDayWiseShoprentCollectionChart(ctx, filter)
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
	templatePath := templatePathStart + "shoprent_daywise_collection_report.html"
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

func (s *Service) DashboardDayWiseShoprentCollectionChartExcel(ctx *models.Context, filter *models.DashboardDayWiseShoprentCollectionChartFilter) (*excelize.File, error) {
	data, err := s.DashboardDayWiseShoprentCollectionChart(ctx, filter)
	if err != nil {
		return nil, err
	}
	excel := excelize.NewFile()
	sheet1 := "Shop Rent Day Wise Collection"
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
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "Shop Count")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Amount")
	rowNo++
	var totalAmount float64
	for _, v := range data.Records {
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), v.ID)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v.ShopRentCount)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), v.Amount)
		totalAmount = totalAmount + v.Amount
		rowNo++

	}
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "D", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "Total")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), totalAmount)

	return excel, nil

}

// DayWiseShoprentDemandChart : ""
func (s *Service) DayWiseShoprentDemandChart(ctx *models.Context, filter *models.DayWiseShoprentDemandChartFilter) (*models.DayWiseShoprentDemandChart, error) {

	res, err := s.Daos.DayWiseShoprentDemandChart(ctx, filter)
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

// DayWiseShoprentDemandChartPDF : ""
func (s *Service) DayWiseShoprentDemandChartPDF(ctx *models.Context, filter *models.DayWiseShoprentDemandChartFilter) ([]byte, error) {
	properties, err := s.DayWiseShoprentDemandChart(ctx, filter)
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
	templatePath := templatePathStart + "shoprent_daywise_demand_report.html"
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

// DayWiseShopRentDemandReportExcel: ""
func (s *Service) DayWiseShopRentDemandReportExcel(ctx *models.Context, filter *models.DayWiseShoprentDemandChartFilter) (*excelize.File, error) {
	data, err := s.DayWiseShoprentDemandChart(ctx, filter)
	if err != nil {
		return nil, err
	}
	excel := excelize.NewFile()
	sheet1 := "Shop Rent Day Wise Demand"
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
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "Shop Count")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Amount")
	rowNo++
	var totalAmount float64
	if data != nil {
		for _, v := range data.Records {
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), v.ID)
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v.ShopRentCount)
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

// FilterWardDayWiseShopRentCollectionReport : ""
func (s *Service) FilterWardDayWiseShopRentCollectionReport(ctx *models.Context, filter *models.WardDayWiseShopRentCollectionReportFilter, pagination *models.Pagination) ([]models.WardDayWiseShopRentCollectionReport, error) {
	return s.Daos.FilterWardDayWiseShopRentCollectionReport(ctx, filter, pagination)
}

// FilterWardDayWiseShopRentCollectionReportExcel: ""
func (s *Service) FilterWardDayWiseShopRentCollectionReportExcel(ctx *models.Context, filter *models.WardDayWiseShopRentCollectionReportFilter, pagination *models.Pagination) (*excelize.File, error) {
	data, err := s.FilterWardDayWiseShopRentCollectionReport(ctx, filter, pagination)
	if err != nil {
		return nil, err
	}
	excel := excelize.NewFile()
	sheet1 := "Shop Rent Ward Wise Collection"
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
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Paied ShopRents")
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

// FilterWardDayWiseShopRentCollectionReportPDF : ""
func (s *Service) FilterWardDayWiseShopRentCollectionReportPDF(ctx *models.Context, filter *models.WardDayWiseShopRentCollectionReportFilter, pagination *models.Pagination) ([]byte, error) {
	data, err := s.FilterWardDayWiseShopRentCollectionReport(ctx, filter, pagination)
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
	templatePath := templatePathStart + "shoprent_wardwise_collection_report.html"
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

// FilterWardMonthWiseShopRentCollectionReport : ""
func (s *Service) FilterWardMonthWiseShopRentCollectionReport(ctx *models.Context, filter *models.WardMonthWiseShopRentCollectionReportFilter, pagination *models.Pagination) ([]models.WardMonthWiseShopRentCollectionReport, error) {
	return s.Daos.FilterWardMonthWiseShopRentCollectionReport(ctx, filter, pagination)
}

// FilterWardMonthWiseShopRentCollectionReport: ""
func (s *Service) FilterWardMonthWiseShopRentCollectionReportExcel(ctx *models.Context, filter *models.WardMonthWiseShopRentCollectionReportFilter, pagination *models.Pagination) (*excelize.File, error) {
	data, err := s.FilterWardMonthWiseShopRentCollectionReport(ctx, filter, pagination)
	if err != nil {
		return nil, err
	}
	excel := excelize.NewFile()
	sheet1 := "Shop Rent Ward Wise Collection"
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
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Paied ShopRents")
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

// FilterWardDayWiseShopRentCollectionReportPDF : ""
func (s *Service) FilterWardMonthWiseShopRentCollectionReportPDF(ctx *models.Context, filter *models.WardMonthWiseShopRentCollectionReportFilter, pagination *models.Pagination) ([]byte, error) {
	data, err := s.FilterWardMonthWiseShopRentCollectionReport(ctx, filter, pagination)
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
	templatePath := templatePathStart + "shoprent_wardwisemonth_collection_report.html"
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

// FilterWardDayWiseShopRentDemandReport : ""
func (s *Service) FilterWardDayWiseShopRentDemandReport(ctx *models.Context, filter *models.WardDayWiseShopRentDemandReportFilter, pagination *models.Pagination) ([]models.WardDayWiseShopRentDemandReport, error) {
	return s.Daos.FilterWardDayWiseShopRentDemandReport(ctx, filter, pagination)
}

// DayWiseShopRentDemandReportExcel: ""
func (s *Service) FilterWardDayWiseShopRentDemandReportExcel(ctx *models.Context, filter *models.WardDayWiseShopRentDemandReportFilter, pagination *models.Pagination) (*excelize.File, error) {
	data, err := s.FilterWardDayWiseShopRentDemandReport(ctx, filter, pagination)
	if err != nil {
		return nil, err
	}
	excel := excelize.NewFile()
	sheet1 := "Shop Rent Ward Wise Demand"
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
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "No of Shops")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Amount")
	rowNo++
	var totalAmount float64
	for i, v := range data {
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), i+1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v.Report.ShopRents)
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

// FilterWardDayWiseShopRentDemandReport : ""
func (s *Service) FilterWardDayWiseShopRentDemandReportPDF(ctx *models.Context, filter *models.WardDayWiseShopRentDemandReportFilter, pagination *models.Pagination) ([]byte, error) {
	data, err := s.FilterWardDayWiseShopRentDemandReport(ctx, filter, pagination)
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
	templatePath := templatePathStart + "shoprent_wardwise_demand_report.html"
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

// FilterWardMonthWiseShopRentDemandReport : ""
func (s *Service) FilterWardMonthWiseShopRentDemandReport(ctx *models.Context, filter *models.WardMonthWiseShopRentDemandReportFilter, pagination *models.Pagination) ([]models.WardMonthWiseShopRentDemandReport, error) {
	return s.Daos.FilterWardMonthWiseShopRentDemandReport(ctx, filter, pagination)
}

// DayWiseShopRentDemandReportExcel: ""
func (s *Service) FilterWardMonthWiseShopRentDemandReportExcel(ctx *models.Context, filter *models.WardMonthWiseShopRentDemandReportFilter, pagination *models.Pagination) (*excelize.File, error) {
	data, err := s.FilterWardMonthWiseShopRentDemandReport(ctx, filter, pagination)
	if err != nil {
		return nil, err
	}
	excel := excelize.NewFile()
	sheet1 := "Shop Rent Ward Wise Demand"
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
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "No of Shops")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Amount")
	rowNo++
	var totalAmount float64
	for i, v := range data {
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), i+1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v.Report.ShopRents)
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

// FilterWardMonthWiseShopRentDemandReportPDF : ""
func (s *Service) FilterWardMonthWiseShopRentDemandReportPDF(ctx *models.Context, filter *models.WardMonthWiseShopRentDemandReportFilter, pagination *models.Pagination) ([]byte, error) {
	data, err := s.FilterWardMonthWiseShopRentDemandReport(ctx, filter, pagination)
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
	templatePath := templatePathStart + "shoprent_wardwisemonth_demand_report.html"
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

// FilterTeamDayWiseShopRentCollectionReport : ""
func (s *Service) FilterTeamDayWiseShopRentCollectionReport(ctx *models.Context, filter *models.TeamDayWiseShopRentCollectionReportFilter, pagination *models.Pagination) ([]models.TeamDayWiseShopRentCollectionReport, error) {
	return s.Daos.FilterTeamDayWiseShopRentCollectionReport(ctx, filter, pagination)
}

// FilterTeamDayWiseShopRentCollectionReportExcel: ""
func (s *Service) FilterTeamDayWiseShopRentCollectionReportExcel(ctx *models.Context, filter *models.TeamDayWiseShopRentCollectionReportFilter, pagination *models.Pagination) (*excelize.File, error) {
	data, err := s.FilterTeamDayWiseShopRentCollectionReport(ctx, filter, pagination)
	if err != nil {
		return nil, err
	}
	excel := excelize.NewFile()
	sheet1 := "Shop Rent Team Wise Collection"
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
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Paied Shops")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), "No of Payments")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), "Collection Amount")
	rowNo++
	var totalAmount float64
	for i, v := range data {
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), i+1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v.User.Name)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), v.Report.TotalNoShopRents)
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

// FilterTeamDayWiseShopRentCollectionReportPDF : ""
func (s *Service) FilterTeamDayWiseShopRentCollectionReportPDF(ctx *models.Context, filter *models.TeamDayWiseShopRentCollectionReportFilter, pagination *models.Pagination) ([]byte, error) {
	data, err := s.FilterTeamDayWiseShopRentCollectionReport(ctx, filter, pagination)
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
	templatePath := templatePathStart + "shoprent_teamwise_collection_report.html"
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

// FilterTeamMonthWiseShopRentCollectionReport : ""
func (s *Service) FilterTeamMonthWiseShopRentCollectionReport(ctx *models.Context, filter *models.TeamMonthWiseShopRentCollectionReportFilter, pagination *models.Pagination) ([]models.TeamMonthWiseShopRentCollectionReport, error) {
	return s.Daos.FilterTeamMonthWiseShopRentCollectionReport(ctx, filter, pagination)
}

// FilterTeamMonthWiseShopRentCollectionReportExcel: ""
func (s *Service) FilterTeamMonthWiseShopRentCollectionReportExcel(ctx *models.Context, filter *models.TeamMonthWiseShopRentCollectionReportFilter, pagination *models.Pagination) (*excelize.File, error) {
	data, err := s.FilterTeamMonthWiseShopRentCollectionReport(ctx, filter, pagination)
	if err != nil {
		return nil, err
	}
	excel := excelize.NewFile()
	sheet1 := "Shop Rent Team Wise Collection"
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
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Paied Shops")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), "No of Payments")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), "Collection Amount")
	rowNo++
	var totalAmount float64
	for i, v := range data {
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), i+1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v.User.Name)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), v.Report.TotalNoShopRents)
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

// FilterTeamMonthWiseShopRentCollectionReportPDF : ""
func (s *Service) FilterTeamMonthWiseShopRentCollectionReportPDF(ctx *models.Context, filter *models.TeamMonthWiseShopRentCollectionReportFilter, pagination *models.Pagination) ([]byte, error) {
	data, err := s.FilterTeamMonthWiseShopRentCollectionReport(ctx, filter, pagination)
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
	templatePath := templatePathStart + "shoprent_teamwisemonth_collection_report.html"
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

// FilterWardYearWiseShopRentCollectionReport : ""
func (s *Service) FilterWardYearWiseShopRentCollectionReport(ctx *models.Context, filter *models.WardYearWiseShopRentCollectionReportFilter, pagination *models.Pagination) ([]models.WardYearWiseShopRentCollectionReport, error) {
	return s.Daos.FilterWardYearWiseShopRentCollectionReport(ctx, filter, pagination)
}

// FilterWardYearWiseShopRentCollectionReportExcel: ""
func (s *Service) FilterWardYearWiseShopRentCollectionReportExcel(ctx *models.Context, filter *models.WardYearWiseShopRentCollectionReportFilter, pagination *models.Pagination) (*excelize.File, error) {
	data, err := s.FilterWardYearWiseShopRentCollectionReport(ctx, filter, pagination)
	if err != nil {
		return nil, err
	}
	excel := excelize.NewFile()
	sheet1 := "Shop Rent Ward Wise Collection"
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
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Paied Shops")
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

// FilterWardYearWiseShopRentCollectionReportPDF : ""
func (s *Service) FilterWardYearWiseShopRentCollectionReportPDF(ctx *models.Context, filter *models.WardYearWiseShopRentCollectionReportFilter, pagination *models.Pagination) ([]byte, error) {
	data, err := s.FilterWardYearWiseShopRentCollectionReport(ctx, filter, pagination)
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
	templatePath := templatePathStart + "shoprent_wardwiseyear_collection_report.html"
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

// FilterWardYearWiseShopRentDemandReport : ""
func (s *Service) FilterWardYearWiseShopRentDemandReport(ctx *models.Context, filter *models.WardYearWiseShopRentDemandReportFilter, pagination *models.Pagination) ([]models.WardYearWiseShopRentDemandReport, error) {
	return s.Daos.FilterWardYearWiseShopRentDemandReport(ctx, filter, pagination)
}

// FilterWardYearWiseShopRentDemandReportExcel: ""
func (s *Service) FilterWardYearWiseShopRentDemandReportExcel(ctx *models.Context, filter *models.WardYearWiseShopRentDemandReportFilter, pagination *models.Pagination) (*excelize.File, error) {
	data, err := s.FilterWardYearWiseShopRentDemandReport(ctx, filter, pagination)
	if err != nil {
		return nil, err
	}
	excel := excelize.NewFile()
	sheet1 := "Shop Rent Ward Wise Demand"
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
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "No of Shops")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Amount")
	rowNo++
	var totalAmount float64
	for i, v := range data {
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), i+1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v.Report.ShopRents)
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

// FilterWardYearWiseShopRentDemandReportPDF : ""
func (s *Service) FilterWardYearWiseShopRentDemandReportPDF(ctx *models.Context, filter *models.WardYearWiseShopRentDemandReportFilter, pagination *models.Pagination) ([]byte, error) {
	data, err := s.FilterWardYearWiseShopRentDemandReport(ctx, filter, pagination)
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
	templatePath := templatePathStart + "shoprent_wardwiseyear_demand_report.html"
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

// FilterTeamYearWiseShopRentCollectionReport : ""
func (s *Service) FilterTeamYearWiseShopRentCollectionReport(ctx *models.Context, filter *models.TeamYearWiseShopRentCollectionReportFilter, pagination *models.Pagination) ([]models.TeamYearWiseShopRentCollectionReport, error) {
	return s.Daos.FilterTeamYearWiseShopRentCollectionReport(ctx, filter, pagination)
}

// FilterTeamYearWiseShopRentCollectionReportExcel: ""
func (s *Service) FilterTeamYearWiseShopRentCollectionReportExcel(ctx *models.Context, filter *models.TeamYearWiseShopRentCollectionReportFilter, pagination *models.Pagination) (*excelize.File, error) {
	data, err := s.FilterTeamYearWiseShopRentCollectionReport(ctx, filter, pagination)
	if err != nil {
		return nil, err
	}
	excel := excelize.NewFile()
	sheet1 := "Shop Rent Team Wise Collection"
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
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Paied Shops")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), "No of Payments")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), "Collection Amount")
	rowNo++
	var totalAmount float64
	for i, v := range data {
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), i+1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v.User.Name)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), v.Report.TotalNoShopRents)
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

// FilterTeamYearWiseShopRentCollectionReportPDF : ""
func (s *Service) FilterTeamYearWiseShopRentCollectionReportPDF(ctx *models.Context, filter *models.TeamYearWiseShopRentCollectionReportFilter, pagination *models.Pagination) ([]byte, error) {
	data, err := s.FilterTeamYearWiseShopRentCollectionReport(ctx, filter, pagination)
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
	templatePath := templatePathStart + "shoprent_teamwiseyear_collection_report.html"
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
