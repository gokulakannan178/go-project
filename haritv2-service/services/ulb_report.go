package services

import (
	"fmt"
	"haritv2-service/models"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func (s *Service) ULBMasterReportV2Json(ctx *models.Context, filter *models.ULBMasterReportV2Filter) ([]models.RefULBMasterReportV2, error) {
	return s.Daos.ULBMasterReportV2(ctx, filter)
}

func (s *Service) ULBMasterReportV2Excel(ctx *models.Context, filter *models.ULBMasterReportV2Filter) (*excelize.File, error) {
	res, err := s.ULBMasterReportV2Json(ctx, filter)
	if err != nil {
		return nil, err
	}
	//  create an excel file
	excel := excelize.NewFile()
	sheet1 := "ULBs"
	rowNo := 1
	index := excel.NewSheet(sheet1)
	excel.SetActiveSheet(index)

	// excel.MergeCell(sheet1, "I1", "M1")
	// excel.MergeCell(sheet1, "N1", "R1")
	// excel.MergeCell(sheet1, "S1", "W1")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "S.No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "ULB Code")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "ULB Name")
	var columnNo1 byte = 'D'
	var columnNo2 byte = 'F'
	for _, v := range filter.Months {
		tempRow := rowNo
		excel.MergeCell(sheet1, fmt.Sprintf("%v%v", string(columnNo1), tempRow), fmt.Sprintf("%v%v", string(columnNo1+4), tempRow))

		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", string(columnNo1), tempRow), fmt.Sprintf("%v-%v", v, filter.Year))
		tempRow++
		excel.MergeCell(sheet1, fmt.Sprintf("%v%v", string(columnNo2), tempRow), fmt.Sprintf("%v%v", string(columnNo2+2), tempRow))
		// D =068
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", string(columnNo1), tempRow), "Compost Generated")
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", string(columnNo1+1), tempRow), "Self")
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", string(columnNo1+2), tempRow), "Sold")
		tempRow++
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", string(columnNo1), tempRow), "in MT")
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", string(columnNo1+1), tempRow), "in MT")
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", string(columnNo1+2), tempRow), "in MT")
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", string(columnNo1+3), tempRow), "Customers")
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", string(columnNo1+4), tempRow), "Revenue Generated")
		columnNo1 = columnNo1 + 5
		columnNo2 = columnNo2 + 3
	}
	rowNo++
	rowNo++
	rowNo++
	for k, v := range res {
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), k+1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v.ULBCode)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), v.Name)
		var columnNo1 byte = 'D'
		// var columnNo2 byte = 'F'
		for _, v2 := range v.Months {
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", string(columnNo1), rowNo), v2.CompostGenerated.Quantity)
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", string(columnNo1+1), rowNo), v2.Sale.Self.Quantity)
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", string(columnNo1+2), rowNo), v2.Sale.Customer.Quantity)
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", string(columnNo1+3), rowNo), v2.Sale.Customer.CustomerCount)
			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", string(columnNo1+4), rowNo), v2.Sale.Customer.Amount)
			columnNo1 = columnNo1 + 5
		}
		rowNo++

	}

	return excel, nil
}
