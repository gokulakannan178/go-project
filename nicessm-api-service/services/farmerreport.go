package services

import (
	"fmt"
	"log"
	"nicessm-api-service/models"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
)

//FilterFarmerReport :""
func (s *Service) FilterFarmerReport(ctx *models.Context, farmerfilter *models.FarmerReportFilter, pagination *models.Pagination) (Query []models.RefFarmer, err error) {
	defer ctx.Session.EndSession(ctx.CTX)
	return s.Daos.FilterFarmerReport(ctx, farmerfilter, pagination)
}

//FilterFarmer :""
func (s *Service) FilterFarmerReport2(ctx *models.Context, Farmerfilter *models.FarmerFilter, pagination *models.Pagination) (Farmer []models.RefFarmer, err error) {
	defer ctx.Session.EndSession(ctx.CTX)
	return s.Daos.FilterFarmerReport2(ctx, Farmerfilter, pagination)

}

//FilterFarmer :""
func (s *Service) FilterDuplicateFarmer(ctx *models.Context, Farmerfilter *models.DuplicateFarmerFilter, pagination *models.Pagination) (Farmer []models.DuplicateFarmerReport, err error) {
	defer ctx.Session.EndSession(ctx.CTX)
	return s.Daos.FilterDuplicateFarmer(ctx, Farmerfilter, pagination)

}
func (s *Service) DuplicateFarmerReportExcel(ctx *models.Context, filter *models.DuplicateFarmerFilter, pagination *models.Pagination) (*excelize.File, error) {
	t := time.Now()
	data, err := s.FilterDuplicateFarmer(ctx, filter, nil)
	if err != nil {
		return nil, err
	}
	duration := time.Since(t)
	log.Println("query Time taken ===> ", duration.Minutes(), "m")
	t = time.Now()

	excel := excelize.NewFile()
	sheet1 := "DuplicateFarmerReport"
	index := excel.NewSheet(sheet1)
	rowNo := 1

	excel.SetActiveSheet(index)
	excel.MergeCell(sheet1, "A1", "E1")
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

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "E", rowNo), style3)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), sheet1)
	rowNo++
	rowNo++
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "E", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "Name")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "FatherName")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "MobileNumber")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), "Village")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), "FarmerId")

	rowNo++

	//	var totalAmount float64
	for _, v := range data {
		for _, v2 := range v.Farmers {
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), v2.Name)
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v2.FatherName)
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), v2.MobileNumber)
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), v2.Ref.Village.Name)
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), v2.FarmerID)
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
