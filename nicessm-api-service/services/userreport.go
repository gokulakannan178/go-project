package services

import (
	"fmt"
	"log"
	"nicessm-api-service/models"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
)

//FilterUserReport :""
func (s *Service) FilterUserReport(ctx *models.Context, filter *models.UserReportFilter, pagination *models.Pagination) (Query []models.RefUser, err error) {
	defer ctx.Session.EndSession(ctx.CTX)
	return s.Daos.FilterUserReport(ctx, filter, pagination)

}

//FilterDuplicateUserReport
func (s *Service) FilterDuplicateUserReport(ctx *models.Context, userfilter *models.DuplicateUserFilter, pagination *models.Pagination) (user []models.DuplicateUserReport, err error) {
	return s.Daos.FilterDuplicateUserReport(ctx, userfilter, pagination)
}
func (s *Service) UserReportExcel(ctx *models.Context, filter *models.UserReportFilter, pagination *models.Pagination) (*excelize.File, error) {
	t := time.Now()
	data, err := s.FilterUserReport(ctx, filter, nil)
	if err != nil {
		return nil, err
	}
	duration := time.Since(t)
	log.Println("query Time taken ===> ", duration.Minutes(), "m")
	t = time.Now()

	excel := excelize.NewFile()
	sheet1 := "UserQueryReport"
	index := excel.NewSheet(sheet1)
	rowNo := 1

	excel.SetActiveSheet(index)
	excel.MergeCell(sheet1, "A1", "D1")
	// excel.MergeCell(sheet1, "C1", "C3")
	// excel.MergeCell(sheet1, "C4", "C5")
	style1, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#FFDC6D"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	style3, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#B6DDE8"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}

	// documentUrl := s.Shared.Config.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.FILEURL)
	// if err := excel.AddPicture(sheet1, fmt.Sprintf("%v%v", "A", rowNo), documentUrl+"municipal/logo.png", `{"x_scale": 0.115, "y_scale": 0.0935}`); err != nil {
	// 	fmt.Println(err)
	// }
	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style3)
	// excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), ctx.ProductConfig.Name)

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "D", rowNo), style3)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), sheet1)
	rowNo++
	rowNo++
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "D", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "Name")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "UserName")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "EmailId")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), "MobileNumber")
	rowNo++

	//	var totalAmount float64
	for _, v := range data {
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), v.Name)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v.UserName)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), v.Email)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), v.Mobile)
		rowNo++
	}

	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style1)
	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "B", rowNo), fmt.Sprintf("%v%v", "B", rowNo), style1)
	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style1)
	// //	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf(" %.2f", totalAmount))
	duration = time.Since(t)
	log.Println("excel Time taken ===> ", duration.Minutes(), "m")
	return excel, nil

}
func (s *Service) DuplicateUserReportExcel(ctx *models.Context, filter *models.DuplicateUserFilter, pagination *models.Pagination) (*excelize.File, error) {
	t := time.Now()
	data, err := s.FilterDuplicateUserReport(ctx, filter, nil)
	if err != nil {
		return nil, err
	}
	duration := time.Since(t)
	log.Println("query Time taken ===> ", duration.Minutes(), "m")
	t = time.Now()

	excel := excelize.NewFile()
	sheet1 := "DuplicateUserReport"
	index := excel.NewSheet(sheet1)
	rowNo := 1

	excel.SetActiveSheet(index)
	excel.MergeCell(sheet1, "A1", "G1")
	// excel.MergeCell(sheet1, "C1", "C3")
	// excel.MergeCell(sheet1, "C4", "C5")
	style1, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#FFDC6D"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	style3, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#B6DDE8"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}

	// documentUrl := s.Shared.Config.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.FILEURL)
	// if err := excel.AddPicture(sheet1, fmt.Sprintf("%v%v", "A", rowNo), documentUrl+"municipal/logo.png", `{"x_scale": 0.115, "y_scale": 0.0935}`); err != nil {
	// 	fmt.Println(err)
	// }
	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style3)
	// excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), ctx.ProductConfig.Name)

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "G", rowNo), style3)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), sheet1)
	rowNo++
	rowNo++
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "G", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "MobileNumber")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "Name")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "UserName")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), "EmailId")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), "Role")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), "State")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), "Distric")
	rowNo++

	//	var totalAmount float64
	for _, v := range data {
		for _, v2 := range v.Users {
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), v2.Mobile)
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v2.Name)
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), v2.UserName)
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), v2.Email)
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), v2.Role)
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), v2.Ref.State.Name)
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), v2.Ref.District.Name)
			rowNo++
		}
	}

	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style1)
	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "B", rowNo), fmt.Sprintf("%v%v", "B", rowNo), style1)
	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style1)
	// //	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf(" %.2f", totalAmount))
	duration = time.Since(t)
	log.Println("excel Time taken ===> ", duration.Minutes(), "m")
	return excel, nil

}
