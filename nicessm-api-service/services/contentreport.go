package services

import (
	"fmt"
	"log"
	"nicessm-api-service/models"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
)

//FilterContentReport :""
func (s *Service) FilterContentReport(ctx *models.Context, Contentfilter *models.ContentReportFilter, pagination *models.Pagination) (Query []models.RefContent, err error) {
	defer ctx.Session.EndSession(ctx.CTX)
	return s.Daos.FilterContentReport(ctx, Contentfilter, pagination)
}

//FilterDuplicateContentReport
func (s *Service) FilterDuplicateContentReport(ctx *models.Context, contentfilter *models.ContentFilter, pagination *models.Pagination) (content []models.DuplicateContentReport, err error) {
	return s.Daos.FilterDuplicateContentReport(ctx, contentfilter, pagination)
}
func (s *Service) DuplicateContentExcel(ctx *models.Context, filter *models.ContentFilter, pagination *models.Pagination) (*excelize.File, error) {
	t := time.Now()
	data, err := s.FilterDuplicateContentReport(ctx, filter, nil)
	if err != nil {
		return nil, err
	}
	duration := time.Since(t)
	log.Println("query Time taken ===> ", duration.Minutes(), "m")
	t = time.Now()

	excel := excelize.NewFile()
	sheet1 := "DuplicateContent"
	index := excel.NewSheet(sheet1)
	rowNo := 1

	excel.SetActiveSheet(index)
	excel.MergeCell(sheet1, "A1", "F1")
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

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "F", rowNo), style3)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), sheet1)
	rowNo++
	rowNo++
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "F", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "Knowledge Domain")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "Sub Domain")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Topic")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), "Sub Topic")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), "Season")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), "RecordId")
	rowNo++

	//	var totalAmount float64
	for _, v := range data {
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), v.Ref.KnowledgeDomain.Name)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v.Ref.SubDomain.Name)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), v.Ref.Topic.Name)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), v.Ref.SubTopic.Name)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), v.Ref.Season.Name)
		for _, v2 := range v.Contents {

			excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), v2.RecordId)
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
