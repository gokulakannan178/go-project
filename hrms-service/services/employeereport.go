package services

import (
	"fmt"
	"hrms-services/models"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func (s *Service) EmployeeReport(ctx *models.Context, filter *models.FilterEmployee) (*excelize.File, error) {
	data, err := s.FilterEmployee(ctx, filter, nil)
	if err != nil {
		return nil, err
	}
	// create an excel file
	excel := excelize.NewFile()
	sheet1 := "EmployeeReport"
	rowNo := 1
	index := excel.NewSheet(sheet1)
	excel.SetActiveSheet(index)
	// excel.MergeCell(sheet1, "A1", "B5")
	// excel.MergeCell(sheet1, "C1", "E3")
	// excel.MergeCell(sheet1, "C4", "E5")
	// excel.MergeCell(sheet1, "A6", "E6")

	// style1, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#B6DDE8"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// style2, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#FFDC6D"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// style4, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#FFDC6D"],"pattern":1},"alignment":{"horizontal":"right","vertical":"center"},"font":{"bold":true}}`)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// style3, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#B6DDE8"],"pattern":1},"font":{"bold":true}}`)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style1)
	// documentUrl := s.Shared.Config.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.FILEURL)
	// if err := excel.AddPicture(sheet1, fmt.Sprintf("%v%v", "A", rowNo), documentUrl+"municipal/logo.png", `{"x_scale": 0.115, "y_scale": 0.0935}`); err != nil {
	// 	fmt.Println(err)
	// }
	// excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), s.Shared.Config.GetString(s.Shared.GetCmdArg(constants.ENV)+"."+constants.LOCATIONNAME))
	// rowNo++
	// rowNo++
	// rowNo++

	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style1)
	// excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Ward Wise Collection")
	// rowNo++
	// rowNo++
	// reportFromMsg := "Report"
	// if filter != nil {
	// 	fmt.Println(filter.StartDate, filter.EndDate)
	// 	if filter.StartDate != nil && filter.EndDate == nil {
	// 		reportFromMsg = reportFromMsg + " on " + fmt.Sprintf("%v-%v-%v", filter.StartDate.Day(), filter.StartDate.Month(), filter.StartDate.Year())
	// 	}
	// 	if filter.StartDate != nil && filter.EndDate != nil {
	// 		reportFromMsg = reportFromMsg + " From " + fmt.Sprintf("%v-%v-%v", filter.StartDate.Day(), filter.StartDate.Month(), filter.StartDate.Year()) + " To " + fmt.Sprintf("%v-%v-%v", filter.EndDate.Day(), filter.EndDate.Month(), filter.EndDate.Year())
	// 	}
	// 	if filter.StartDate == nil && filter.EndDate == nil {
	// 		fmt.Println("everything is nil")
	// 	}

	// }
	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style1)
	// excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), reportFromMsg)

	// rowNo++
	// totalCollectionsIndex := rowNo
	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style3)
	// excel.MergeCell(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "E", rowNo))

	//rowNo++
	//excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "E", rowNo), style2)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "S.No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "Name")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "DOB")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), "Gender")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), "Mobile")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), "Email")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), "Designation")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "H", rowNo), "OrganisationID")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "I", rowNo), "ManagerID")

	rowNo++
	//var totalPayments float64
	//var totalProperties float64

	for i, v := range data {
		//totalPayments = totalPayments + v.PropertyPayments.Payments
		//totalProperties = totalProperties + float64(v.Properties.Count)

		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), i+1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v.Name)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), v.DOB)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), v.Gender)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), v.Mobile)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), v.Email)
		//excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), v.Designation)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "H", rowNo), v.OrganisationID)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "I", rowNo), v.ManagerId)

		//totalPayments = totalPayments + v.PropertyPayments.Payments
		rowNo++
	}
	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style2)
	// excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "Total")

	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "B", rowNo), fmt.Sprintf("%v%v", "B", rowNo), style4)
	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "D", rowNo), fmt.Sprintf("%v%v", "D", rowNo), style4)

	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style4)
	// excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%.0f", totalProperties))

	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "E", rowNo), fmt.Sprintf("%v%v", "E", rowNo), style4)
	// excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), fmt.Sprintf("%.2f", totalPayments))

	// excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", totalCollectionsIndex), fmt.Sprintf("Total Collection - %.2f", totalPayments))

	return excel, nil
}
