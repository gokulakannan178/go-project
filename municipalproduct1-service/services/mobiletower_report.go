package services

import (
	"errors"
	"fmt"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
)

// DashboardDayWiseMobileTowerCollectionChart : ""
func (s *Service) DashboardDayWiseMobileTowerCollectionChart(ctx *models.Context, filter *models.DashboardDayWiseMobileTowerCollectionChartFilter) (models.DashboardDayWiseMobileTowerCollectionChart, error) {
	return s.Daos.DashboardDayWiseMobileTowerCollectionChart(ctx, filter)
}

func (s *Service) DashboardDayWiseMobileTowerCollectionChartExcel(ctx *models.Context, filter *models.DashboardDayWiseMobileTowerCollectionChartFilter) (*excelize.File, error) {
	data, err := s.DashboardDayWiseMobileTowerCollectionChart(ctx, filter)
	if err != nil {
		return nil, err
	}
	excel := excelize.NewFile()
	sheet1 := "Mobile Tower Day Wise Collection"
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
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "Mobile Tower Count")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Amount")
	rowNo++
	var totalAmount float64
	for _, v := range data.Records {
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), v.ID)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v.MobileTowerCount)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), v.Amount)
		totalAmount = totalAmount + v.Amount
		rowNo++

	}
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "D", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "Total")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), totalAmount)

	return excel, nil

}

func (s *Service) DashboardDayWiseMobileTowerCollectionChartPDF(ctx *models.Context, filter *models.DashboardDayWiseMobileTowerCollectionChartFilter) ([]byte, error) {
	data, err := s.DashboardDayWiseMobileTowerCollectionChart(ctx, filter)
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
	templatePath := templatePathStart + "mobiletower_daywise_collection_report.html"
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

// /MobileTowerOverallDemandReportPDF : ""
func (s *Service) MobileTowerOverallDemandReportPDF(ctx *models.Context, filter *models.PropertyMobileTowerFilter, p *models.Pagination) ([]byte, error) {
	properties, err := s.Daos.FilterPropertyMobileTower(ctx, filter, p)
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
	templatePath := templatePathStart + "mobiletower_demand_report.html"
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

// MobileTowerOverallDemandReportExcel : ""
func (s *Service) MobileTowerOverallDemandReportExcel(ctx *models.Context, filter *models.PropertyMobileTowerFilter, pagination *models.Pagination) (*excelize.File, error) {
	res, err := s.FilterPropertyMobileTower(ctx, filter, pagination)
	// res, err := s.MobileTowerOverallDemandReportJSON(ctx, filter, pagination)

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
	sheet1 := "MobileTowerOverallDemand"
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
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Mobile Tower Overall Demand")
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
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), v.Demand.Total.Total)
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

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "E", rowNo), fmt.Sprintf("%v%v", "F", rowNo), style1)

	return excel, nil
}

// /PropertyOverallDemandReportPDF : ""
func (s *Service) MobileTowerOverallDemandReportJSON(ctx *models.Context, filter *models.PropertyMobileTowerFilter, p *models.Pagination) (*models.MobileTowerOverallDemandReport, error) {
	mobileTowers, err := s.Daos.FilterPropertyMobileTower(ctx, filter, p)
	if err != nil {
		return nil, err
	}

	report := new(models.MobileTowerOverallDemandReport)
	cfy, err := s.Daos.GetCurrentFinancialYear(ctx)
	if err != nil {
		return nil, errors.New("Error in getting current financial year " + err.Error())
	}
	report.MobileTowers = mobileTowers
	if cfy != nil {
		report.CFY = *cfy
	}
	return report, nil
}

// DayWiseMobileTowerDemandChart : ""
func (s *Service) DayWiseMobileTowerDemandChart(ctx *models.Context, filter *models.DayWiseMobileTowerDemandChartFilter) (*models.DayWiseMobileTowerDemandChart, error) {
	res, err := s.Daos.DayWiseMobileTowerDemandChart(ctx, filter)
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

// DayWiseMobileTowerDemandReportPDF : ""
func (s *Service) DayWiseMobileTowerDemandReportPDF(ctx *models.Context, filter *models.DayWiseMobileTowerDemandChartFilter) ([]byte, error) {
	properties, err := s.DayWiseMobileTowerDemandChart(ctx, filter)
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
	templatePath := templatePathStart + "mobiletower_daywise_demand_report.html"
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

// DayWiseMobileTowerDemandChartPDF : ""
func (s *Service) DayWiseMobileTowerDemandChartPDF(ctx *models.Context, filter *models.DayWiseMobileTowerDemandChartFilter) ([]byte, error) {
	data, err := s.DayWiseMobileTowerDemandChart(ctx, filter)
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
	templatePath := templatePathStart + "mobiletower_daywise_demand_report.html"
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

// DayWiseMobileTowerDemandReportExcel: ""
func (s *Service) DayWiseMobileTowerDemandReportExcel(ctx *models.Context, filter *models.DayWiseMobileTowerDemandChartFilter) (*excelize.File, error) {
	data, err := s.DayWiseMobileTowerDemandChart(ctx, filter)
	if err != nil {
		return nil, err
	}
	excel := excelize.NewFile()
	sheet1 := "Mobile Tower Day Wise Demand"
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
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "Mobile Tower Count")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Amount")
	rowNo++
	var totalAmount float64
	if data != nil {
		for _, v := range data.Records {
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), v.ID)
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v.MobileTowerCount)
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

// FilterWardDayWiseMobileTowerDemandReport : ""
func (s *Service) FilterWardDayWiseMobileTowerDemandReport(ctx *models.Context, filter *models.WardDayWiseMobileTowerDemandReportFilter, pagination *models.Pagination) ([]models.WardDayWiseMobileTowerDemandReport, error) {
	return s.Daos.FilterWardDayWiseMobileTowerDemandReport(ctx, filter, pagination)
}

// DayWiseMobileTowerDemandReportExcel: ""
func (s *Service) FilterWardDayWiseMobileTowerDemandReportExcel(ctx *models.Context, filter *models.WardDayWiseMobileTowerDemandReportFilter, pagination *models.Pagination) (*excelize.File, error) {
	data, err := s.FilterWardDayWiseMobileTowerDemandReport(ctx, filter, pagination)
	if err != nil {
		return nil, err
	}
	excel := excelize.NewFile()
	sheet1 := "Mobile Tower Ward Wise Demand"
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
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "No of Mobile Towers")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Amount")
	rowNo++
	var totalAmount float64
	for i, v := range data {
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), i+1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v.Report.MobileTowers)
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

// FilterWardDayWiseMobileTowerDemandReport : ""
func (s *Service) FilterWardDayWiseMobileTowerDemandReportPDF(ctx *models.Context, filter *models.WardDayWiseMobileTowerDemandReportFilter, pagination *models.Pagination) ([]byte, error) {
	data, err := s.FilterWardDayWiseMobileTowerDemandReport(ctx, filter, pagination)
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
	templatePath := templatePathStart + "mobiletower_wardwise_demand_report.html"
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

// FilterWardDayWiseMobileTowerCollectionReport : ""
func (s *Service) FilterWardDayWiseMobileTowerCollectionReport(ctx *models.Context, filter *models.WardDayWiseMobileTowerCollectionReportFilter, pagination *models.Pagination) ([]models.WardDayWiseMobileTowerCollectionReport, error) {
	return s.Daos.FilterWardDayWiseMobileTowerCollectionReport(ctx, filter, pagination)
}

// FilterWardDayWiseMobileTowerCollectionReportExcel: ""
func (s *Service) FilterWardDayWiseMobileTowerCollectionReportExcel(ctx *models.Context, filter *models.WardDayWiseMobileTowerCollectionReportFilter, pagination *models.Pagination) (*excelize.File, error) {
	data, err := s.FilterWardDayWiseMobileTowerCollectionReport(ctx, filter, pagination)
	if err != nil {
		return nil, err
	}
	excel := excelize.NewFile()
	sheet1 := "Mobile Tower Ward Wise Collection"
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
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Paied Mobile Towers")
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

// FilterWardDayWiseMobileTowerCollectionReportPDF : ""
func (s *Service) FilterWardDayWiseMobileTowerCollectionReportPDF(ctx *models.Context, filter *models.WardDayWiseMobileTowerCollectionReportFilter, pagination *models.Pagination) ([]byte, error) {
	data, err := s.FilterWardDayWiseMobileTowerCollectionReport(ctx, filter, pagination)
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
	templatePath := templatePathStart + "mobiletower_wardwise_collection_report.html"
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

// FilterWardMonthWiseMobileTowerCollectionReport : ""
func (s *Service) FilterWardMonthWiseMobileTowerCollectionReport(ctx *models.Context, filter *models.WardMonthWiseMobileTowerCollectionReportFilter, pagination *models.Pagination) ([]models.WardMonthWiseMobileTowerCollectionReport, error) {
	return s.Daos.FilterWardMonthWiseMobileTowerCollectionReport(ctx, filter, pagination)
}

// FilterWardMonthWiseMobileTowerCollectionReport: ""
func (s *Service) FilterWardMonthWiseMobileTowerCollectionReportExcel(ctx *models.Context, filter *models.WardMonthWiseMobileTowerCollectionReportFilter, pagination *models.Pagination) (*excelize.File, error) {
	data, err := s.FilterWardMonthWiseMobileTowerCollectionReport(ctx, filter, pagination)
	if err != nil {
		return nil, err
	}
	excel := excelize.NewFile()
	sheet1 := "Mobile Tower Ward Wise Collection"
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
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Paied mobile Towers")
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

// FilterWardDayWiseMobileTowerCollectionReportPDF : ""
func (s *Service) FilterWardMonthWiseMobileTowerCollectionReportPDF(ctx *models.Context, filter *models.WardMonthWiseMobileTowerCollectionReportFilter, pagination *models.Pagination) ([]byte, error) {
	data, err := s.FilterWardMonthWiseMobileTowerCollectionReport(ctx, filter, pagination)
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
	templatePath := templatePathStart + "mobiletower_wardwisemonth_collection_report.html"
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

// FilterWardMonthWiseMobileTowerDemandReport : ""
func (s *Service) FilterWardMonthWiseMobileTowerDemandReport(ctx *models.Context, filter *models.WardMonthWiseMobileTowerDemandReportFilter, pagination *models.Pagination) ([]models.WardMonthWiseMobileTowerDemandReport, error) {
	return s.Daos.FilterWardMonthWiseMobileTowerDemandReport(ctx, filter, pagination)
}

// DayWiseMobileTowerDemandReportExcel: ""
func (s *Service) FilterWardMonthWiseMobileTowerDemandReportExcel(ctx *models.Context, filter *models.WardMonthWiseMobileTowerDemandReportFilter, pagination *models.Pagination) (*excelize.File, error) {
	data, err := s.FilterWardMonthWiseMobileTowerDemandReport(ctx, filter, pagination)
	if err != nil {
		return nil, err
	}
	excel := excelize.NewFile()
	sheet1 := "Mobile Tower Ward Wise Demand"
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
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "No of Mobile Towers")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Amount")
	rowNo++
	var totalAmount float64
	for i, v := range data {
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), i+1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v.Report.MobileTowers)
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

// FilterWardMonthWiseMobileTowerDemandReportPDF : ""
func (s *Service) FilterWardMonthWiseMobileTowerDemandReportPDF(ctx *models.Context, filter *models.WardMonthWiseMobileTowerDemandReportFilter, pagination *models.Pagination) ([]byte, error) {
	data, err := s.FilterWardMonthWiseMobileTowerDemandReport(ctx, filter, pagination)
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
	templatePath := templatePathStart + "mobiletower_wardwisemonth_demand_report.html"
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

// FilterTeamDayWiseMobileTowerCollectionReport : ""
func (s *Service) FilterTeamDayWiseMobileTowerCollectionReport(ctx *models.Context, filter *models.TeamDayWiseMobileTowerCollectionReportFilter, pagination *models.Pagination) ([]models.TeamDayWiseMobileTowerCollectionReport, error) {
	return s.Daos.FilterTeamDayWiseMobileTowerCollectionReport(ctx, filter, pagination)
}

// FilterTeamDayWiseMobileTowerCollectionReportExcel: ""
func (s *Service) FilterTeamDayWiseMobileTowerCollectionReportExcel(ctx *models.Context, filter *models.TeamDayWiseMobileTowerCollectionReportFilter, pagination *models.Pagination) (*excelize.File, error) {
	data, err := s.FilterTeamDayWiseMobileTowerCollectionReport(ctx, filter, pagination)
	if err != nil {
		return nil, err
	}
	excel := excelize.NewFile()
	sheet1 := "Mobile Tower Team Wise Collection"
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
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Paied Mobile Towers")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), "No of Payments")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), "Collection Amount")
	rowNo++
	var totalAmount float64
	for i, v := range data {
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), i+1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v.User.Name)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), v.Report.TotalNoMobileTowers)
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

// FilterTeamDayWiseMobileTowerCollectionReportPDF : ""
func (s *Service) FilterTeamDayWiseMobileTowerCollectionReportPDF(ctx *models.Context, filter *models.TeamDayWiseMobileTowerCollectionReportFilter, pagination *models.Pagination) ([]byte, error) {
	data, err := s.FilterTeamDayWiseMobileTowerCollectionReport(ctx, filter, pagination)
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
	templatePath := templatePathStart + "mobiletower_teamwise_collection_report.html"
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

// FilterTeamMonthWiseMobileTowerCollectionReport : ""
func (s *Service) FilterTeamMonthWiseMobileTowerCollectionReport(ctx *models.Context, filter *models.TeamMonthWiseMobileTowerCollectionReportFilter, pagination *models.Pagination) ([]models.TeamMonthWiseMobileTowerCollectionReport, error) {
	return s.Daos.FilterTeamMonthWiseMobileTowerCollectionReport(ctx, filter, pagination)
}

// FilterTeamMonthWiseMobileTowerCollectionReportExcel: ""
func (s *Service) FilterTeamMonthWiseMobileTowerCollectionReportExcel(ctx *models.Context, filter *models.TeamMonthWiseMobileTowerCollectionReportFilter, pagination *models.Pagination) (*excelize.File, error) {
	data, err := s.FilterTeamMonthWiseMobileTowerCollectionReport(ctx, filter, pagination)
	if err != nil {
		return nil, err
	}
	excel := excelize.NewFile()
	sheet1 := "Mobile Tower Team Wise Collection"
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
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Paied Mobile Towers")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), "No of Payments")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), "Collection Amount")
	rowNo++
	var totalAmount float64
	for i, v := range data {
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), i+1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v.User.Name)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), v.Report.TotalNoMobileTowers)
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

// FilterTeamMonthWiseMobileTowerCollectionReportPDF : ""
func (s *Service) FilterTeamMonthWiseMobileTowerCollectionReportPDF(ctx *models.Context, filter *models.TeamMonthWiseMobileTowerCollectionReportFilter, pagination *models.Pagination) ([]byte, error) {
	data, err := s.FilterTeamMonthWiseMobileTowerCollectionReport(ctx, filter, pagination)
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
	templatePath := templatePathStart + "mobiletower_teamwisemonth_collection_report.html"
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

// FilterWardYearWiseMobileTowerCollectionReport : ""
func (s *Service) FilterWardYearWiseMobileTowerCollectionReport(ctx *models.Context, filter *models.WardYearWiseMobileTowerCollectionReportFilter, pagination *models.Pagination) ([]models.WardYearWiseMobileTowerCollectionReport, error) {
	return s.Daos.FilterWardYearWiseMobileTowerCollectionReport(ctx, filter, pagination)
}

// FilterWardYearWiseMobileTowerCollectionReportExcel: ""
func (s *Service) FilterWardYearWiseMobileTowerCollectionReportExcel(ctx *models.Context, filter *models.WardYearWiseMobileTowerCollectionReportFilter, pagination *models.Pagination) (*excelize.File, error) {
	data, err := s.FilterWardYearWiseMobileTowerCollectionReport(ctx, filter, pagination)
	if err != nil {
		return nil, err
	}
	excel := excelize.NewFile()
	sheet1 := "Mobile Tower Ward Wise Collection"
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
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Paied Mobile Towers")
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

// FilterWardYearWiseMobileTowerCollectionReportPDF : ""
func (s *Service) FilterWardYearWiseMobileTowerCollectionReportPDF(ctx *models.Context, filter *models.WardYearWiseMobileTowerCollectionReportFilter, pagination *models.Pagination) ([]byte, error) {
	data, err := s.FilterWardYearWiseMobileTowerCollectionReport(ctx, filter, pagination)
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
	templatePath := templatePathStart + "mobiletower_wardwiseyear_collection_report.html"
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

// FilterWardYearWiseMobileTowerDemandReport : ""
func (s *Service) FilterWardYearWiseMobileTowerDemandReport(ctx *models.Context, filter *models.WardYearWiseMobileTowerDemandReportFilter, pagination *models.Pagination) ([]models.WardYearWiseMobileTowerDemandReport, error) {
	return s.Daos.FilterWardYearWiseMobileTowerDemandReport(ctx, filter, pagination)
}

// FilterWardYearWiseMobileTowerDemandReportExcel: ""
func (s *Service) FilterWardYearWiseMobileTowerDemandReportExcel(ctx *models.Context, filter *models.WardYearWiseMobileTowerDemandReportFilter, pagination *models.Pagination) (*excelize.File, error) {
	data, err := s.FilterWardYearWiseMobileTowerDemandReport(ctx, filter, pagination)
	if err != nil {
		return nil, err
	}
	excel := excelize.NewFile()
	sheet1 := "Mobile Tower Ward Wise Demand"
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
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "No of Mobile Towers")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Amount")
	rowNo++
	var totalAmount float64
	for i, v := range data {
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), i+1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v.Report.MobileTowers)
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

// FilterWardYearWiseMobileTowerDemandReportPDF : ""
func (s *Service) FilterWardYearWiseMobileTowerDemandReportPDF(ctx *models.Context, filter *models.WardYearWiseMobileTowerDemandReportFilter, pagination *models.Pagination) ([]byte, error) {
	data, err := s.FilterWardYearWiseMobileTowerDemandReport(ctx, filter, pagination)
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
	templatePath := templatePathStart + "mobiletower_wardwiseyear_demand_report.html"
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

// FilterTeamYearWiseMobileTowerCollectionReport : ""
func (s *Service) FilterTeamYearWiseMobileTowerCollectionReport(ctx *models.Context, filter *models.TeamYearWiseMobileTowerCollectionReportFilter, pagination *models.Pagination) ([]models.TeamYearWiseMobileTowerCollectionReport, error) {
	return s.Daos.FilterTeamYearWiseMobileTowerCollectionReport(ctx, filter, pagination)
}

// FilterTeamYearWiseMobileTowerCollectionReportExcel: ""
func (s *Service) FilterTeamYearWiseMobileTowerCollectionReportExcel(ctx *models.Context, filter *models.TeamYearWiseMobileTowerCollectionReportFilter, pagination *models.Pagination) (*excelize.File, error) {
	data, err := s.FilterTeamYearWiseMobileTowerCollectionReport(ctx, filter, pagination)
	if err != nil {
		return nil, err
	}
	excel := excelize.NewFile()
	sheet1 := "Mobile Tower Team Wise Collection"
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
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Paied Mobile Towers")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), "No of Payments")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), "Collection Amount")
	rowNo++
	var totalAmount float64
	for i, v := range data {
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), i+1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v.User.Name)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), v.Report.TotalNoMobileTowers)
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

// FilterTeamYearWiseMobileTowerCollectionReportPDF : ""
func (s *Service) FilterTeamYearWiseMobileTowerCollectionReportPDF(ctx *models.Context, filter *models.TeamYearWiseMobileTowerCollectionReportFilter, pagination *models.Pagination) ([]byte, error) {
	data, err := s.FilterTeamYearWiseMobileTowerCollectionReport(ctx, filter, pagination)
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
	templatePath := templatePathStart + "mobiletower_teamwiseyear_collection_report.html"
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
