package services

import (
	"context"
	"fmt"
	"municipalproduct1-service/app"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func (s *Service) CalcualteTotalCollectionForAllPayments() error {
	var ctx *models.Context
	c := context.TODO()
	ctx = app.GetApp(c, s.Daos)
	defer ctx.Client.Disconnect(c)
	return s.Daos.CalcualteTotalCollectionForAllPayments(ctx)
}

//DashboardTotalCollectionChart : ""
func (s *Service) DashboardTotalCollectionChart(ctx *models.Context, filter *models.DashboardTotalCollectionChartFilter) ([]models.DashboardTotalCollectionChart, error) {
	return s.Daos.DashboardTotalCollectionChart(ctx, filter)
}

func (s *Service) DashboardTotalCollectionOverview(ctx *models.Context, filter *models.DashboardTotalCollectionOverviewFilter) (models.DashboardTotalCollectionOverview, error) {
	return s.Daos.DashboardTotalCollectionOverview(ctx, filter)
}

func (s *Service) DashboardDayWiseCollectionChart(ctx *models.Context, filter *models.DashboardDayWiseCollectionChartFilter) (models.DashboardDayWiseCollectionChart, error) {

	data, err := s.Daos.DashboardDayWiseCollectionChart(ctx, filter)
	if err != nil {
		return models.DashboardDayWiseCollectionChart{}, err
	}
	for k := range data.Records {

		data.Records[k].ID.DayStr = fmt.Sprintf("%v-%v-%v", data.Records[k].ID.Day, filter.StartDate.Month(), filter.StartDate.Year())
		if ctx.ProductConfig.LocationID == "Munger" {
			if data.Records[k].ID.DayStr == "14-June-2023" {
				data.Records[k].CurrentCollection = data.Records[k].CurrentCollection - 3817620
				data.Records[k].TotalCollection = data.Records[k].TotalCollection - 3817620
				data.Records[k].TotalTax = data.Records[k].TotalTax - 3817620
			}
		}
	}
	return data, nil
}

func (s *Service) WardWiseCollectionReport(ctx *models.Context, filter *models.WardWiseCollectionReportFilter, pagination *models.Pagination) ([]models.WardWiseCollectionReport, error) {
	return s.Daos.WardWiseCollectionReport(ctx, filter, pagination)
}

func (s *Service) WardWiseCollectionReportExcel(ctx *models.Context, filter *models.WardWiseCollectionReportFilter, pagination *models.Pagination) (*excelize.File, error) {
	data, err := s.WardWiseCollectionReport(ctx, filter, pagination)
	if err != nil {
		return nil, err
	}

	resPD, err1 := s.Daos.GetSingleDefaultProductConfiguration(ctx)
	if err1 != nil {
		return nil, err1
	}

	// create an excel file
	excel := excelize.NewFile()
	sheet1 := "Ward Wise Collection"
	rowNo := 1
	index := excel.NewSheet(sheet1)
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
	excel.MergeCell(sheet1, "A7", "E7")

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
	style3, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#B6DDE8"],"pattern":1},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style1)

	if resPD.LocationID != "Bhagalpur" {
		documentUrl := s.Shared.Config.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.FILEURL)
		if err := excel.AddPicture(sheet1, fmt.Sprintf("%v%v", "A", rowNo), documentUrl+"municipal/logo.png", `{"x_scale": 0.115, "y_scale": 0.0935}`); err != nil {
			fmt.Println(err)
		}
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), s.Shared.Config.GetString(s.Shared.GetCmdArg(constants.ENV)+"."+constants.LOCATIONNAME))
	} else {
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), s.Shared.Config.GetString(s.Shared.GetCmdArg(constants.ENV)+"."+constants.LOCATIONNAME))

	}
	rowNo++
	rowNo++
	rowNo++

	if resPD.LocationID == "Bhagalpur" {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "Ward Wise Collection")
	} else {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Ward Wise Collection")
	}
	rowNo++
	rowNo++
	reportFromMsg := "Report"
	if filter != nil {
		fmt.Println(filter.StartDate, filter.EndDate)
		if filter.StartDate != nil && filter.EndDate == nil {
			reportFromMsg = reportFromMsg + " on " + fmt.Sprintf("%v-%v-%v", filter.StartDate.Day(), filter.StartDate.Month(), filter.StartDate.Year())
		}
		if filter.StartDate != nil && filter.EndDate != nil {
			reportFromMsg = reportFromMsg + " From " + fmt.Sprintf("%v-%v-%v", filter.StartDate.Day(), filter.StartDate.Month(), filter.StartDate.Year()) + " To " + fmt.Sprintf("%v-%v-%v", filter.EndDate.Day(), filter.EndDate.Month(), filter.EndDate.Year())
		}
		if filter.StartDate == nil && filter.EndDate == nil {
			fmt.Println("everything is nil")
		}

	}
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), reportFromMsg)

	rowNo++

	t3 := time.Now()
	toDate := t3.Format("02-January-2006")
	reportFromMsg2 := "Report Generated on" + " " + toDate
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), reportFromMsg2)
	rowNo++

	totalCollectionsIndex := rowNo
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style3)
	excel.MergeCell(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "E", rowNo))

	rowNo++
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "E", rowNo), style2)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "S.No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "Ward Name")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Total Propeties")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), "Property Count")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), "Amount Collected")

	rowNo++
	var totalPayments float64
	var totalProperties float64

	for i, v := range data {
		totalPayments = totalPayments + v.PropertyPayments.Payments
		totalProperties = totalProperties + float64(v.Properties.Count)

		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), i+1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v.Name)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), v.Properties.Count)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), v.PropertyPayments.Properties)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), v.PropertyPayments.Payments)
		totalPayments = totalPayments + v.PropertyPayments.Payments
		rowNo++
	}
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style2)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "Total")

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "B", rowNo), fmt.Sprintf("%v%v", "B", rowNo), style4)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "D", rowNo), fmt.Sprintf("%v%v", "D", rowNo), style4)

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style4)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%.0f", totalProperties))

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "E", rowNo), fmt.Sprintf("%v%v", "E", rowNo), style4)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), fmt.Sprintf("%.2f", totalPayments))

	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", totalCollectionsIndex), fmt.Sprintf("Total Collection - %.2f", totalPayments))

	return excel, nil
}

// TCCollectionSummaryReport : ""
func (s *Service) TCCollectionSummaryReport(ctx *models.Context, filter *models.TCCollectionSummaryFilter) ([]models.TCCollectionSummaryReport, error) {
	return s.Daos.TCCollectionSummaryReport(ctx, filter)
	// res, err := s.TCCollectionSummaryReport(ctx, filter)
	// if err != nil {
	// 	return nil, err
	// }
	// var totalConsumer int64
	// var totalAmount float64
	// if len(res) > 0 {
	// 	for _, v := range res {
	// 		totalConsumer = totalConsumer + v.PropertyPayments.PropertyCount
	// 		totalAmount = totalAmount + v.PropertyPayments.Payments
	// 	}
	// }
	// fmt.Println("totalConsumer =====> ", totalConsumer)
	// fmt.Println("totalAmount =====> ", totalAmount)
	// // tcCollection := new(models.TCCollectionSummaryReportV2)
	// if len(res) > 0 {
	// 	for _, v := range res {
	// 		v.TotalConsumer = totalConsumer
	// 		v.TotalAmount = totalAmount
	// 	}

	// }
	// return res, nil
}

// TCCollectionSummaryReportV2 : ""
func (s *Service) TCCollectionSummaryReportV2(ctx *models.Context, filter *models.TCCollectionSummaryFilter) (*models.TCCollectionSummaryReportV2, error) {
	res, err := s.TCCollectionSummaryReport(ctx, filter)
	if err != nil {
		return nil, err
	}
	var totalConsumer int64
	var totalAmount float64
	if len(res) > 0 {
		for _, v := range res {
			totalConsumer = totalConsumer + v.PropertyPayments.PropertyCount
			totalAmount = totalAmount + v.PropertyPayments.Payments
		}
	}
	fmt.Println("totalConsumer =====> ", totalConsumer)
	fmt.Println("totalAmount =====> ", totalAmount)
	tcCollection := new(models.TCCollectionSummaryReportV2)
	tcCollection.TotalConsumer = totalConsumer
	tcCollection.TotalAmount = totalAmount
	tcCollection.TCCollectionSummaryReport = res
	return tcCollection, nil
}

// TCCollectionSummaryReportExcel : ""
func (s *Service) TCCollectionSummaryReportExcel(ctx *models.Context, filter *models.TCCollectionSummaryFilter) (*excelize.File, error) {
	res, err := s.TCCollectionSummaryReport(ctx, filter)
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
	sheet1 := "Property Payment Collection"
	rowNo := 1
	index := excel.NewSheet(sheet1)
	excel.SetActiveSheet(index)
	if resPD.LocationID == "Bhagalpur" {
		excel.MergeCell(sheet1, "A1", "G3")
		excel.MergeCell(sheet1, "A4", "G5")
	} else {
		excel.MergeCell(sheet1, "A1", "B5")
		excel.MergeCell(sheet1, "C1", "G3")
		excel.MergeCell(sheet1, "C4", "G5")
	}
	excel.MergeCell(sheet1, "A6", "G6")
	excel.MergeCell(sheet1, "A7", "G7")
	excel.MergeCell(sheet1, "A8", "G8")
	excel.MergeCell(sheet1, "A9", "G9")
	style3, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#B6DDE8"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	// style1, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#FFDC6D"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	style2, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#B6DDE8"],"pattern":1},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	style1, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#FFDC6D"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	style4, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#FFDC6D"],"pattern":1},"alignment":{"horizontal":"right","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style3)

	if resPD.LocationID != "Bhagalpur" {
		documentUrl := s.Shared.Config.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.FILEURL)
		if err := excel.AddPicture(sheet1, fmt.Sprintf("%v%v", "A", rowNo), documentUrl+"municipal/logo.png", `{"x_scale": 0.115, "y_scale": 0.0935}`); err != nil {
			fmt.Println(err)
		}
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), s.Shared.Config.GetString(s.Shared.GetCmdArg(constants.ENV)+"."+constants.LOCATIONNAME))
	} else {
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), s.Shared.Config.GetString(s.Shared.GetCmdArg(constants.ENV)+"."+constants.LOCATIONNAME))
	}
	rowNo++
	rowNo++
	rowNo++

	if resPD.LocationID == "Bhagalpur" {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "Team Wise Collection")
	} else {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Team Wise Collection")
	}
	rowNo++
	rowNo++

	reportFromMsg := "Report"

	if filter != nil {
		if filter.DateRange != nil {
			fmt.Println(filter.DateRange.From, filter.DateRange.To)
			if filter.DateRange.From != nil && filter.DateRange.To == nil {
				reportFromMsg = reportFromMsg + " on " + fmt.Sprintf("%v-%v-%v", filter.DateRange.From.Day(), filter.DateRange.From.Month(), filter.DateRange.From.Year())
			}
			if filter.DateRange.From != nil && filter.DateRange.To != nil {
				reportFromMsg = reportFromMsg + " From " + fmt.Sprintf("%v-%v-%v", filter.DateRange.From.Day(), filter.DateRange.From.Month(), filter.DateRange.From.Year()) + " To " + fmt.Sprintf("%v-%v-%v", filter.DateRange.To.Day(), filter.DateRange.To.Month(), filter.DateRange.To.Year())
			}
			if filter.DateRange.From == nil && filter.DateRange.To == nil {
				fmt.Println("everything is nil")
			}

		}

	}
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style3)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), reportFromMsg)
	rowNo++

	t3 := time.Now()
	toDate := t3.Format("02-January-2006")
	reportFromMsg2 := "Report Generated on" + " " + toDate
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style3)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), reportFromMsg2)
	rowNo++

	var totalAmountRow, totalConsumerRow int
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style2)
	totalConsumerRow = rowNo
	rowNo++

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style2)
	totalAmountRow = rowNo
	rowNo++

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "G", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "S.No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "Collector Name")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Reports to")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), "User Type")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), "Mobile No")
	// excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), "Total Holding")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), "Total Properties")
	// excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), "Last Transaction Date")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), "Total Collections")

	var totalAmount float64
	var totalConsumer int64
	var totalPayments int
	// var totalAmountRow1 int
	totalPayments = len(res)
	for i, v := range res {
		totalPayments++
		totalConsumer = totalConsumer + v.PropertyPayments.PropertyCount
		totalAmount = totalAmount + v.PropertyPayments.Payments

		rowNo++

		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), i+1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v.User.Name)
		// excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), v.Manager.Name)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), func() string {
			if v.Manager.Name != "" {
				return v.Manager.Name
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), func() string {
			if v.UserType.Name != "" {
				return v.UserType.Name
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), v.User.Mobile)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), v.PropertyPayments.PropertyCount)
		// if v.PropertyPayments.RecentTransaction != nil {
		// 	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), v.PropertyPayments.RecentTransaction.Format("2006-01-02"))
		// } else {
		// 	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), "NA")

		// }
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), v.PropertyPayments.Payments)

	}
	rowNo++
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "Total")

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "B", rowNo), fmt.Sprintf("%v%v", "B", rowNo), style4)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style4)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "D", rowNo), fmt.Sprintf("%v%v", "D", rowNo), style4)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "E", rowNo), fmt.Sprintf("%v%v", "E", rowNo), style4)

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "F", rowNo), fmt.Sprintf("%v%v", "F", rowNo), style4)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), fmt.Sprintf("%v", totalConsumer))

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "G", rowNo), fmt.Sprintf("%v%v", "G", rowNo), style4)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), fmt.Sprintf("%.2f", totalAmount))

	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", totalConsumerRow), fmt.Sprintf("Total Consumer - %v", totalConsumer))
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", totalAmountRow), fmt.Sprintf("Total Collection - %.2f", totalAmount))
	return excel, nil
}

//TCCollectionSummaryReportPdf
/*func (s *Service) TCCollectionSummaryReportPdf(ctx *models.Context, filter *models.TCCollectionSummaryFilter) ([]byte, error) {
	res, err := s.TCCollectionSummaryReport(ctx, filter)
	if err != nil {
		return nil, err
	}
	m := make(map[string]interface{})
	m2 := make(map[string]interface{})
	m["demand"] = demand

	state, err := s.Daos.GetSingleState(ctx, demand.Address.StateCode)
	if state != nil {
		m2["state"] = &state.State
	}
	fmt.Println(err)
	district, err := s.Daos.GetSingleDistrict(ctx, demand.Address.DistrictCode)
	if district != nil {
		m2["district"] = &district.District
	}
	fmt.Println(err)
	village, err := s.Daos.GetSingleVillage(ctx, demand.Address.VillageCode)
	if village != nil {
		m2["village"] = &village.Village
	}
	fmt.Println(err)
	zone, err := s.Daos.GetSingleZone(ctx, demand.Address.ZoneCode)
	if zone != nil {
		m2["zone"] = &zone.Zone
	}
	fmt.Println(err)
	ward, err := s.Daos.GetSingleWard(ctx, demand.Address.WardCode)
	if ward != nil {
		m2["ward"] = &ward.Ward
	}
	fmt.Println(err)
	fy, err := s.Daos.GetCurrentFinancialYear(ctx)
	if fy != nil {
		m2["cfy"] = &fy.FinancialYear
	}

	m2["currentDate"] = time.Now()
	m["extraRef"] = m2
	r := NewRequestPdf("")
	templatePathStart := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.TEMPLATELOC)
	//html template path
	templatePath := templatePathStart + "Demand_Page.html"

	//path for download pdf
	// t := time.Now()
	// outputPath := fmt.Sprintf("storage/SampleTemplate%v.pdf", t.Unix())
	if err := r.ParseTemplate(templatePath, m); err == nil {
		ok, data, err := r.GeneratePDFAsFile()

		fmt.Println(ok, "pdf generated successfully")
		return data, err
	} else {
		return nil, errors.New("Error in parcing template - " + err.Error())
	}
	return nil, nil
	//  create an excel file

}*/

// TeamWiseCollectionReport : ""
// func (s *Service) TeamWiseCollectionReport(ctx *models.Context, filter *models.WardWiseCollectionReportFilter) (*excelize.File, error) {
// 	// create an excel file
// 	excel := excelize.NewFile()
// 	sheet1 := "Team Wise Collection"
// 	rowNo := 1
// 	index := excel.NewSheet(sheet1)
// 	excel.SetActiveSheet(index)
// 	excel.MergeCell(sheet1, "A1", "B5")
// 	excel.MergeCell(sheet1, "C1", "E3")
// 	excel.MergeCell(sheet1, "C4", "E5")
// 	excel.MergeCell(sheet1, "A6", "E6")
// 	excel.MergeCell(sheet1, "A8", "E8")

// 	style, err := excel.NewStyle(`{"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style)
// 	documentUrl := s.Shared.Config.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.FILEURL)
// 	if err := excel.AddPicture(sheet1, fmt.Sprintf("%v%v", "A", rowNo), documentUrl+"municipal/logo.png", `{"x_scale": 0.115, "y_scale": 0.0935}`); err != nil {
// 		fmt.Println(err)
// 	}
// 	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), s.Shared.Config.GetString(s.Shared.GetCmdArg(constants.ENV)+"."+constants.LOCATIONNAME))
// 	rowNo++
// 	rowNo++
// 	rowNo++

// 	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style)
// 	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Team Wise Collection")
// 	rowNo++
// 	rowNo++

// 	// excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "Team Wise Collection from __________ to __________")
// 	reportFromMsg := "Report"

// 	if filter != nil {
// 		if filter.StartDate != nil {
// 			fmt.Println(filter.StartDate, filter.EndDate)
// 			if filter.StartDate != nil && filter.EndDate == nil {
// 				reportFromMsg = reportFromMsg + " on " + fmt.Sprintf("%v-%v-%v", filter.StartDate.Day(), filter.StartDate.Month(), filter.StartDate.Year())
// 			}
// 			if filter.StartDate != nil && filter.EndDate != nil {
// 				reportFromMsg = reportFromMsg + " From " + fmt.Sprintf("%v-%v-%v", filter.StartDate.Day(), filter.StartDate.Month(), filter.StartDate.Year()) + " To " + fmt.Sprintf("%v-%v-%v", filter.EndDate.Day(), filter.EndDate.Month(), filter.EndDate.Year())
// 			}
// 			if filter.StartDate == nil && filter.EndDate == nil {
// 				fmt.Println("everything is nil")
// 			}

// 		}

// 	}
// 	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style)
// 	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), reportFromMsg)
// 	rowNo++

// 	totalCollectionsIndex := rowNo
// 	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "E", rowNo), style)
// 	excel.MergeCell(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "E", rowNo))
// 	rowNo++

// 	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "E", rowNo), style)
// 	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "Total Consumer")
// 	rowNo++

// 	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "E", rowNo), style)
// 	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "S.No")
// 	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "Collector Name")
// 	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Total Propeties")
// 	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), "User Type")
// 	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), "Total Collection")
// 	data, err := s.WardWiseCollectionReport(ctx, filter)
// 	if err != nil {
// 		return nil, err
// 	}
// 	rowNo++
// 	var totalPayments float64
// 	for i, v := range data {
// 		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), i+1)
// 		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "collectorName")
// 		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), v.Properties.Count)
// 		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), "usertype")
// 		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), v.PropertyPayments.Payments)
// 		totalPayments = totalPayments + v.PropertyPayments.Payments
// 		rowNo++
// 	}
// 	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", totalCollectionsIndex), fmt.Sprintf("Total Collection - %v", totalPayments))

// 	return excel, nil
// }

// PropertyIDWiseReport : ""
func (s *Service) PropertyIDWiseReport(ctx *models.Context, filter *models.PropertyFilter, pagination *models.Pagination) (*excelize.File, error) {
	res, err := s.FilterProperty(ctx, filter, pagination)
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
	sheet1 := "Property"
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
	excel.MergeCell(sheet1, "A8", "O8")

	// style, err := excel.NewStyle(`{"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	style1, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#FFDC6D"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	style2, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#B6DDE8"],"pattern":1},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	style3, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#B6DDE8"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	style4, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#FFDC6D"],"pattern":1},"alignment":{"horizontal":"right","vertical":"center"},"font":{"bold":true}}`)
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
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), s.Shared.Config.GetString(s.Shared.GetCmdArg(constants.ENV)+"."+constants.LOCATIONNAME))
	} else {
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), s.Shared.Config.GetString(s.Shared.GetCmdArg(constants.ENV)+"."+constants.LOCATIONNAME))
	}
	rowNo++
	rowNo++
	rowNo++

	if resPD.LocationID == "Bhagalpur" {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "Property ID Wise Report")
	} else {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Property ID Wise Report")
	}
	rowNo++
	rowNo++
	var totalAmountRow, totalDemandRow, totalBalanceRow int
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style2)
	totalDemandRow = rowNo
	rowNo++

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style2)
	totalAmountRow = rowNo
	rowNo++
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style2)
	totalBalanceRow = rowNo
	rowNo++

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "O", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "S.No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "Ward No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Property ID")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), "Owner Name")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), "Guardian Name")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), "Property Address")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), "Mobile No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "H", rowNo), "Usage Type")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "I", rowNo), "Total Demand")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "J", rowNo), "From Year Demand")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "K", rowNo), "To Year Demand")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "L", rowNo), "Total Collection")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "M", rowNo), "From Year Collection")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "N", rowNo), "To Year Collection")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "O", rowNo), "Balance")

	fmt.Println("'res length==>'", len(res))
	var totalAmount, totalDemand, totalBalance float64
	var totalPayments int
	totalPayments = len(res)

	for i, v := range res {
		totalPayments++
		totalDemand = totalDemand + v.Demand.TotalTax + v.Collection.TotalTax
		totalAmount = totalAmount + v.Collection.TotalTax
		totalBalance = totalBalance + v.Demand.TotalTax

		rowNo++
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), i+1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), i+1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), v.UniqueID)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), func() string {
			if len(v.Ref.PropertyOwner) > 0 {
				return v.Ref.PropertyOwner[0].Name
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), func() string {
			if len(v.Ref.PropertyOwner) > 0 {
				return v.Ref.PropertyOwner[0].FatherRpanRhusband
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), v.Address.AL1+" "+v.Address.Al2)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), func() string {
			if len(v.Ref.PropertyOwner) > 0 {
				return v.Ref.PropertyOwner[0].Mobile
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "H", rowNo), func() string {
			if v.Ref.PropertyType != nil {
				return v.Ref.PropertyType.Name
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "I", rowNo), v.Demand.TotalTax+v.Collection.TotalTax)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "J", rowNo), v.Demand.FromYear.Name)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "K", rowNo), v.Demand.ToYear.Name)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "L", rowNo), v.Collection.TotalTax)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "M", rowNo), v.Collection.FromYear.Name)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "N", rowNo), v.Collection.ToYear.Name)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "O", rowNo), v.Demand.TotalTax)

	}
	rowNo++
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "Total")

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "B", rowNo), fmt.Sprintf("%v%v", "B", rowNo), style4)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style4)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "D", rowNo), fmt.Sprintf("%v%v", "D", rowNo), style4)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "E", rowNo), fmt.Sprintf("%v%v", "E", rowNo), style4)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "F", rowNo), fmt.Sprintf("%v%v", "F", rowNo), style4)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "G", rowNo), fmt.Sprintf("%v%v", "G", rowNo), style4)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "H", rowNo), fmt.Sprintf("%v%v", "H", rowNo), style4)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "J", rowNo), fmt.Sprintf("%v%v", "J", rowNo), style4)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "K", rowNo), fmt.Sprintf("%v%v", "K", rowNo), style4)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "M", rowNo), fmt.Sprintf("%v%v", "M", rowNo), style4)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "N", rowNo), fmt.Sprintf("%v%v", "N", rowNo), style4)

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "I", rowNo), fmt.Sprintf("%v%v", "I", rowNo), style4)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "I", rowNo), fmt.Sprintf("%.2f", totalDemand))

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "L", rowNo), fmt.Sprintf("%v%v", "L", rowNo), style4)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "L", rowNo), fmt.Sprintf("%.2f", totalAmount))

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "O", rowNo), fmt.Sprintf("%v%v", "O", rowNo), style4)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "O", rowNo), fmt.Sprintf("%.2f", totalBalance))

	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", totalDemandRow), fmt.Sprintf("Total Demand - %.2f", totalDemand))
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", totalAmountRow), fmt.Sprintf("Total Collection - %.2f", totalAmount))
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", totalBalanceRow), fmt.Sprintf("Total Balance - %.2f", totalBalance))

	return excel, nil
}

// ZoneAndWardWiseReportExcel : ""
func (s *Service) ZoneAndWardWiseReportExcel(ctx *models.Context, filter *models.ZoneAndWardWiseReportFilter, pagination *models.Pagination) (*excelize.File, error) {
	res, err := s.ZoneAndWardWiseReport(ctx, filter, pagination)
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
	sheet1 := "Zone And Ward Wise Collection"
	rowNo := 1
	index := excel.NewSheet(sheet1)
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
	excel.MergeCell(sheet1, "A7", "D7")
	excel.MergeCell(sheet1, "A8", "D8")
	excel.MergeCell(sheet1, "A9", "D9")

	style1, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#FFDC6D"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	style3, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#B6DDE8"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	style2, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#B6DDE8"],"pattern":1},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	style4, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#FFDC6D"],"pattern":1},"alignment":{"horizontal":"right","vertical":"center"},"font":{"bold":true}}`)
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
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), s.Shared.Config.GetString(s.Shared.GetCmdArg(constants.ENV)+"."+constants.LOCATIONNAME))
	} else {
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), s.Shared.Config.GetString(s.Shared.GetCmdArg(constants.ENV)+"."+constants.LOCATIONNAME))
	}
	rowNo++
	rowNo++
	rowNo++

	if resPD.LocationID == "Bhagalpur" {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "Zone and Ward Wise Collection")
	} else {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Zone and Ward Wise Collection")
	}
	rowNo++
	rowNo++

	reportFromMsg := "Report"

	if filter != nil {
		if filter.DateRange != nil {
			fmt.Println(filter.DateRange.From, filter.DateRange.To)
			if filter.DateRange.From != nil && filter.DateRange.To == nil {
				reportFromMsg = reportFromMsg + " on " + fmt.Sprintf("%v-%v-%v", filter.DateRange.From.Day(), filter.DateRange.From.Month(), filter.DateRange.From.Year())
			}
			if filter.DateRange.From != nil && filter.DateRange.To != nil {
				reportFromMsg = reportFromMsg + " From " + fmt.Sprintf("%v-%v-%v", filter.DateRange.From.Day(), filter.DateRange.From.Month(), filter.DateRange.From.Year()) + " To " + fmt.Sprintf("%v-%v-%v", filter.DateRange.To.Day(), filter.DateRange.To.Month(), filter.DateRange.To.Year())
			}
			if filter.DateRange.From == nil && filter.DateRange.To == nil {
				fmt.Println("everything is nil")
			}

		}

	}
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style3)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), reportFromMsg)
	rowNo++

	t := time.Now()
	toDate := t.Format("02-January-2006")
	reportFromMsg2 := "Report Generated on" + " " + toDate
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style3)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), reportFromMsg2)
	rowNo++

	// Total Property
	var totalAmountRow, totalPropertyRow int
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style2)
	totalPropertyRow = rowNo
	rowNo++

	// Total Collection
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style2)
	totalAmountRow = rowNo
	rowNo++

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "D", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "Zone No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "Ward No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Total Property")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), "Total Collection")

	fmt.Println("'res length==>'", len(res))
	var totalAmount float64
	var totalProperty float64
	for _, v := range res {
		rowNo++

		currentRowNoStart := rowNo
		currentRowNoEnd := currentRowNoStart + len(v.Wards) - 1
		excel.MergeCell(sheet1, fmt.Sprintf("%v%v", "A", currentRowNoStart), fmt.Sprintf("%v%v", "A", currentRowNoEnd))
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", currentRowNoStart), fmt.Sprintf("%v%v", "A", currentRowNoEnd), style1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), v.Zone.Name)

		for _, v2 := range v.Wards {
			totalAmount = totalAmount + v2.Payments.TotalCollections
			totalProperty = totalProperty + v2.Payments.TotalProperties

			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v2.Ward.Name)
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), v2.Payments.TotalProperties)
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), v2.Payments.TotalCollections)
			rowNo++
		}
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style4)

		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "B", rowNo), fmt.Sprintf("%v%v", "B", rowNo), style1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "Total")
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style4)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%.0f", totalProperty))
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "D", rowNo), fmt.Sprintf("%v%v", "D", rowNo), style4)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), fmt.Sprintf("%.2f", totalAmount))
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", totalPropertyRow), fmt.Sprintf("Total Property - %.0f", totalProperty))
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", totalAmountRow), fmt.Sprintf("Total Collection - %.2f", totalAmount))

	}

	return excel, nil
}

// SavePropertyCollectionForAll : ""
func (s *Service) SavePropertyCollectionForAll(IDs []string) error {

	fmt.Println(IDs)

	for _, v := range IDs {
		c := context.TODO()
		ctx := app.GetApp(c, s.Daos)
		defer ctx.Client.Disconnect(c)
		err := s.PropertyUpdateCollection(ctx, v)
		fmt.Println(v, err)
	}
	return nil
}

// FilterPropertyMonthWiseCollectionReport : ""
func (s *Service) FilterPropertyMonthWiseCollectionReport(ctx *models.Context, filter *models.PropertyMonthWiseCollectionReportFilter) ([]models.PropertyMonthWiseCollectionReport, error) {
	return s.Daos.FilterPropertyMonthWiseCollectionReport(ctx, filter)
}

// FilterPropertyMonthWiseCollectionReportExcel : ""
func (s *Service) FilterPropertyMonthWiseCollectionReportExcel(ctx *models.Context, filter *models.PropertyMonthWiseCollectionReportFilter) (*excelize.File, error) {
	fmt.Println("hi I am here")
	res, err := s.FilterPropertyMonthWiseCollectionReport(ctx, filter)
	if err != nil {
		return nil, err
	}
	fmt.Println("'res length==>'", len(res))

	//  create an excel file
	excel := excelize.NewFile()
	sheet1 := "Property Month Wise Collection"
	rowNo := 1
	index := excel.NewSheet(sheet1)
	excel.SetActiveSheet(index)
	excel.MergeCell(sheet1, "A1", "B5")
	excel.MergeCell(sheet1, "C1", "D3")
	excel.MergeCell(sheet1, "C4", "D5")
	excel.MergeCell(sheet1, "A6", "D6")
	excel.MergeCell(sheet1, "A7", "D7")
	excel.MergeCell(sheet1, "A8", "D8")

	style1, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#FFDC6D"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	style3, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#B6DDE8"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	style2, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#B6DDE8"],"pattern":1},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
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
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Property Month Wise Collection")
	rowNo++
	rowNo++

	reportFromMsg := "Report"

	if filter != nil {
		// if filter.DateRange != nil {
		// 	fmt.Println(filter.DateRange.From, filter.DateRange.To)
		// 	if filter.DateRange.From != nil && filter.DateRange.To == nil {
		// 		reportFromMsg = reportFromMsg + " on " + fmt.Sprintf("%v-%v-%v", filter.DateRange.From.Day(), filter.DateRange.From.Month(), filter.DateRange.From.Year())
		// 	}
		// 	if filter.DateRange.From != nil && filter.DateRange.To != nil {
		// 		reportFromMsg = reportFromMsg + " From " + fmt.Sprintf("%v-%v-%v", filter.DateRange.From.Day(), filter.DateRange.From.Month(), filter.DateRange.From.Year()) + " To " + fmt.Sprintf("%v-%v-%v", filter.DateRange.To.Day(), filter.DateRange.To.Month(), filter.DateRange.To.Year())
		// 	}
		// 	if filter.DateRange.From == nil && filter.DateRange.To == nil {
		// 		fmt.Println("everything is nil")
		// 	}
		// }

	}
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style3)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), reportFromMsg)
	rowNo++

	// Total Property
	var totalAmountRow, totalPropertyRow int
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style2)
	totalPropertyRow = rowNo
	rowNo++

	// Total Collection
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style2)
	totalAmountRow = rowNo
	rowNo++

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "D", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "Financial Years")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "Month")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Total Current Amount")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), "total Arrear Amount")

	fmt.Println("'res length==>'", len(res))
	var totalCurrentAmount float64
	var totalArrearAmount float64
	for _, v := range res {
		rowNo++

		currentRowNoStart := rowNo
		currentRowNoEnd := currentRowNoStart + len(v.Records) - 1
		excel.MergeCell(sheet1, fmt.Sprintf("%v%v", "A", currentRowNoStart), fmt.Sprintf("%v%v", "A", currentRowNoEnd))
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", currentRowNoStart), fmt.Sprintf("%v%v", "A", currentRowNoEnd), style1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), v.FinancialYear.Name)
		fmt.Println("FY Name", v.PropertyPayment.FinancialYear.Name)
		fmt.Println("FY Name", v.FinancialYear.Name)
		for _, v2 := range v.Records {
			totalCurrentAmount = totalCurrentAmount + v2.CurrentAmount
			totalArrearAmount = totalArrearAmount + v2.ArrearAmount
			// totalProperty = totalProperty + v2.Payments.TotalProperties

			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v2.ID.Month)
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), v2.CurrentAmount)
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), v2.ArrearAmount)
			rowNo++
		}
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style4)

		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "B", rowNo), fmt.Sprintf("%v%v", "B", rowNo), style1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "Total")
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style4)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%.0f", totalCurrentAmount))
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "D", rowNo), fmt.Sprintf("%v%v", "D", rowNo), style4)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), fmt.Sprintf("%.2f", totalArrearAmount))
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", totalPropertyRow), fmt.Sprintf("Total Current Amount - %.0f", totalCurrentAmount))
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", totalAmountRow), fmt.Sprintf("Total Arrear Amount - %.2f", totalArrearAmount))

	}

	return excel, nil
}

// FilterPropertyWiseCollectionReport : ""
func (s *Service) FilterPropertyWiseCollectionReport(ctx *models.Context, filter *models.PropertyWiseCollectionReportFilter) ([]models.PropertyWiseCollectionReport, error) {
	return s.Daos.FilterPropertyWiseCollectionReport(ctx, filter)
}

// FilterPropertyWiseCollectionReportExcel
func (s *Service) FilterPropertyWiseCollectionReportExcel(ctx *models.Context, filter *models.PropertyWiseCollectionReportFilter) (*excelize.File, error) {
	data, err := s.FilterPropertyWiseCollectionReport(ctx, filter)
	if err != nil {
		return nil, err
	}

	resPD, err1 := s.Daos.GetSingleDefaultProductConfiguration(ctx)
	if err1 != nil {
		return nil, err1
	}

	excel := excelize.NewFile()
	sheet1 := "Property Collection Report"
	index := excel.NewSheet(sheet1)
	rowNo := 1

	excel.SetActiveSheet(index)
	if resPD.LocationID == "Bhagalpur" {
		excel.MergeCell(sheet1, "A1", "L3")
		excel.MergeCell(sheet1, "A4", "L5")
	} else {
		excel.MergeCell(sheet1, "A1", "B5")
		excel.MergeCell(sheet1, "C1", "L3")
		excel.MergeCell(sheet1, "C4", "L5")
	}
	excel.MergeCell(sheet1, "A6", "L6")
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

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "L", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "S.No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "Ward No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Property No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), "Owner Name")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), "Guardian Name")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), "Mobile No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), "Property Address")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "H", rowNo), "From Year")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "I", rowNo), "To Year")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "J", rowNo), "Amount")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "K", rowNo), "Penalty")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "L", rowNo), "Total Amount")
	rowNo++
	var totalAmount float64

	for i, v := range data {

		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), i+1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v.Ref.Address.Ward.Name)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), v.UniqueID)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), func() string {
			if len(v.Ref.PropertyOwner) > 0 {
				return v.Ref.PropertyOwner[0].Name
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), func() string {
			if len(v.Ref.PropertyOwner) > 0 {
				return v.Ref.PropertyOwner[0].FatherRpanRhusband
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), func() string {
			if len(v.Ref.PropertyOwner) > 0 {
				return v.Ref.PropertyOwner[0].Mobile
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), v.Address.AL1+" "+v.Address.Al2)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "H", rowNo), v.Ref.Payments.FyFrom)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "I", rowNo), v.Ref.Payments.FyTo)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "J", rowNo), v.Ref.Payments.TotalTax)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "K", rowNo), v.Ref.Payments.TotalPenalty)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "L", rowNo), v.Ref.Payments.TotalCollection)
		totalAmount = totalAmount + v.Ref.Payments.TotalCollection
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
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "K", rowNo), fmt.Sprintf("%v%v", "K", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "L", rowNo), fmt.Sprintf("%v%v", "L", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "Total")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "L", rowNo), fmt.Sprintf(" %.2f", totalAmount))

	return excel, nil

}

// FilterPropertyWiseDemandReportExcel
func (s *Service) FilterPropertyWiseDemandReportExcel(ctx *models.Context, filter *models.PropertyFilter) (*excelize.File, error) {
	data, err := s.PropertyOverallDemandReport(ctx, filter, nil)
	if err != nil {
		return nil, err
	}

	resPD, err1 := s.Daos.GetSingleDefaultProductConfiguration(ctx)
	if err1 != nil {
		return nil, err1
	}

	excel := excelize.NewFile()
	sheet1 := "Property Demand Report"
	index := excel.NewSheet(sheet1)
	rowNo := 1

	excel.SetActiveSheet(index)
	if resPD.LocationID == "Bhagalpur" {
		excel.MergeCell(sheet1, "A1", "K3")
		excel.MergeCell(sheet1, "A4", "K5")
	} else {
		excel.MergeCell(sheet1, "A1", "B5")
		excel.MergeCell(sheet1, "C1", "K3")
		excel.MergeCell(sheet1, "C4", "K5")
	}
	excel.MergeCell(sheet1, "A6", "K6")
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

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "K", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "S.No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "Ward No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Property No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), "Owner Name")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), "Guardian Name")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), "Mobile No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), "Property Address")
	// excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "H", rowNo), "From Year")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "H", rowNo), "To Year")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "I", rowNo), "Amount")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "J", rowNo), "Penalty")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "K", rowNo), "Total Amount")
	rowNo++
	var totalAmount float64

	for i, v := range data {

		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), i+1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v.Ref.Address.Ward.Name)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), v.UniqueID)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), func() string {
			if len(v.Ref.PropertyOwner) > 0 {
				return v.Ref.PropertyOwner[0].Name
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), func() string {
			if len(v.Ref.PropertyOwner) > 0 {
				return v.Ref.PropertyOwner[0].FatherRpanRhusband
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), func() string {
			if len(v.Ref.PropertyOwner) > 0 {
				return v.Ref.PropertyOwner[0].Mobile
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), v.Address.AL1+" "+v.Address.Al2)
		// excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "H", rowNo), v.DOA.Format("2006-January-02"))
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "H", rowNo), t.Format("2006-January-02"))
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "I", rowNo), (v.Ref.Demand.Total.VacantLandTax + v.Ref.Demand.Total.Tax))
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "J", rowNo), v.Ref.Demand.Total.Penalty)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "K", rowNo), (v.Ref.Demand.Total.VacantLandTax + v.Ref.Demand.Total.Tax + v.Ref.Demand.Total.Penalty))
		totalAmount = totalAmount + (v.Ref.Demand.Total.VacantLandTax + v.Ref.Demand.Total.Tax + v.Ref.Demand.Total.Penalty)
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
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "K", rowNo), fmt.Sprintf("%v%v", "K", rowNo), style1)
	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "L", rowNo), fmt.Sprintf("%v%v", "L", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "Total")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "K", rowNo), fmt.Sprintf(" %.2f", totalAmount))

	return excel, nil

}
