package services

import (
	"errors"
	"fmt"
	"log"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"go.mongodb.org/mongo-driver/mongo"
)

// SaveTradeLicenseDashboard : ""
func (s *Service) SaveTradeLicenseDashboard(ctx *models.Context, tradeLicense *models.TradeLicenseDashboard) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	tradeLicense.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONDASHBOARDTRADELICENSE)
	tradeLicense.Status = constants.DASHBOARDTRADELICENSESTATUSACTIVE
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveTradeLicenseDashboard(ctx, tradeLicense)
		if dberr != nil {
			return dberr
		}
		if err := ctx.Session.CommitTransaction(sc); err != nil {
			return errors.New("Not able to commit - " + err.Error())
		}
		return nil
	}); err != nil {
		log.Println("Transaction start aborting")
		if abortError := ctx.Session.AbortTransaction(ctx.CTX); abortError != nil {
			return errors.New("Error while aborting transaction" + abortError.Error())
		}
		log.Println("Transaction aborting completed successfully")
		return err
	}
	return nil
}

//GetSingleTradeLicenseDashboard :""
func (s *Service) GetSingleTradeLicenseDashboard(ctx *models.Context, UniqueID string) (*models.RefTradeLicenseDashboard, error) {
	tower, err := s.Daos.GetSingleTradeLicenseDashboard(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return tower, nil
}

// UpdateTradeLicenseDashboard : ""
func (s *Service) UpdateTradeLicenseDashboard(ctx *models.Context, tradeLicense *models.TradeLicenseDashboard) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateTradeLicenseDashboard(ctx, tradeLicense)
		if err != nil {
			if err = ctx.Session.AbortTransaction(sc); err != nil {
				return errors.New("Transaction Aborted with error" + err.Error())
			}
			return errors.New("Transaction Aborted - " + err.Error())
		}
		return nil

	}); err != nil {
		return err
	}
	return nil
}

// EnableTradeLicenseDashboard : ""
func (s *Service) EnableTradeLicenseDashboard(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableTradeLicenseDashboard(ctx, UniqueID)
		if err != nil {
			if err = ctx.Session.AbortTransaction(sc); err != nil {
				return errors.New("Transaction Aborted with error" + err.Error())
			}
			return errors.New("Transaction Aborted - " + err.Error())
		}
		return nil

	}); err != nil {
		return err
	}
	return nil
}

//DisableTradeLicenseDashboard : ""
func (s *Service) DisableTradeLicenseDashboard(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableTradeLicenseDashboard(ctx, UniqueID)
		if err != nil {
			if err = ctx.Session.AbortTransaction(sc); err != nil {
				return errors.New("Transaction Aborted with error" + err.Error())
			}
			return errors.New("Transaction Aborted - " + err.Error())
		}
		return nil

	}); err != nil {
		return err
	}
	return nil
}

//DeleteTradeLicenseDashboard : ""
func (s *Service) DeleteDashBoardTradeLicense(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteTradeLicenseDashboard(ctx, UniqueID)
		if err != nil {
			if err = ctx.Session.AbortTransaction(sc); err != nil {
				return errors.New("Transaction Aborted with error" + err.Error())
			}
			return errors.New("Transaction Aborted - " + err.Error())
		}
		return nil

	}); err != nil {
		return err
	}
	return nil
}

// FilterTradeLicenseDashboard : ""
func (s *Service) FilterTradeLicenseDashboard(ctx *models.Context, filter *models.TradeLicenseDashboardFilter, pagination *models.Pagination) ([]models.RefTradeLicenseDashboard, error) {
	return s.Daos.FilterTradeLicenseDashboard(ctx, filter, pagination)

}

// DashboardTradeLicenseDemandAndCollection : ""
func (s *Service) DashboardTradeLicenseDemandAndCollection(ctx *models.Context, filter *models.DashboardTradeLicenseDemandAndCollectionFilter) (*models.DashboardTradeLicenseDemandAndCollection, error) {
	//defer ctx.Session.EndSession(ctx.CTX)
	if filter != nil {
		t := time.Now()
		sdt := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
		edt := time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, t.Location())
		filter.TodayRange.From = &sdt
		filter.TodayRange.To = &edt
		tYesterday := t.AddDate(0, 0, -1)
		sdtyesterday := time.Date(tYesterday.Year(), tYesterday.Month(), tYesterday.Day(), 0, 0, 0, 0, tYesterday.Location())
		edtyesterday := time.Date(tYesterday.Year(), tYesterday.Month(), tYesterday.Day(), 59, 59, 59, 0, tYesterday.Location())
		filter.YesterdayRange.From = &sdtyesterday
		filter.YesterdayRange.To = &edtyesterday
		fmt.Println("date to =======>", filter.YesterdayRange.To)
	}
	res, err := s.Daos.DashBoardStatusWiseTradeLicenseCollectionAndChart(ctx, filter)
	if err != nil {
		return nil, err
	}
	data, err := s.Daos.DashboardTradeLicenseDemandAndCollection(ctx, filter)
	if err != nil {
		return nil, err
	}
	data.SAFCount = *res

	return data, nil
}

func (s *Service) UserwiseTradelicenseReport(ctx *models.Context, filter *models.UserFilter) ([]models.UserwiseTradeLicense, error) {
	return s.Daos.UserwiseTradelicenseReport(ctx, filter)
}
func (s *Service) UserwiseTradeLicenseReportExcel(ctx *models.Context, filter *models.UserFilter) (*excelize.File, error) {
	res, err := s.Daos.UserwiseTradelicenseReport(ctx, filter)
	if err != nil {
		return nil, err
	}
	fmt.Println("'res length==>'", len(res))
	excel := excelize.NewFile()
	sheet1 := "UserWiseChargeReport"
	rowNo := 1
	index := excel.NewSheet(sheet1)
	excel.SetActiveSheet(index)
	excel.MergeCell(sheet1, "A1", "B5")
	excel.MergeCell(sheet1, "C1", "L3")
	excel.MergeCell(sheet1, "C4", "L5")
	excel.MergeCell(sheet1, "A6", "L6")
	//excel.MergeCell(sheet1, "A7", "M9")
	excel.MergeCell(sheet1, "A7", "A9")
	excel.MergeCell(sheet1, "B7", "B9")
	excel.MergeCell(sheet1, "C7", "C9")
	excel.MergeCell(sheet1, "D7", "K7")
	//excel.MergeCell(sheet1, "E7", "L7")
	excel.MergeCell(sheet1, "D8", "E8")
	excel.MergeCell(sheet1, "F8", "G8")
	excel.MergeCell(sheet1, "H8", "I8")
	excel.MergeCell(sheet1, "J8", "K8")
	excel.MergeCell(sheet1, "L7", "L9")
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
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Usercharge Payment List")
	rowNo++
	rowNo++
	//
	reportFromMsg := "Report"
	//t := time.Now()
	// var sd string

	// if filter != nil {
	// 	if filter.DateRange != nil {
	// 		fmt.Println(filter.DateRange.From, filter.DateRange.To)
	// 		if filter.DateRange.From != nil && filter.DateRange.To == nil {
	// 			ed := time.Date(filter.DateRange.From.Year(), filter.DateRange.From.Month(), filter.DateRange.From.Day(), 23, 59, 59, 0, filter.DateRange.From.Location())
	// 			toDate := ed.Format("02-January-2006")
	// 			sd = fmt.Sprintf("%v-%v-%v", filter.DateRange.From.Day(), filter.DateRange.From.Month(), filter.DateRange.From.Year()) + " To " + toDate
	// 		}
	// 		if filter.DateRange.From != nil && filter.DateRange.To != nil {
	// 			sd = reportFromMsg + " From " + fmt.Sprintf("%v-%v-%v", filter.DateRange.From.Day(), filter.DateRange.From.Month(), filter.DateRange.From.Year()) + " To " + fmt.Sprintf("%v-%v-%v", filter.DateRange.To.Day(), filter.DateRange.To.Month(), filter.DateRange.To.Year())
	// 		}
	// 		if filter.DateRange.From == nil && filter.DateRange.To == nil {
	// 			fmt.Println("everything is nil")
	// 		}

	// 	}

	// }
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style3)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), reportFromMsg)
	rowNo++
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "M", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), "Amount Collected")
	rowNo++
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "D", rowNo), fmt.Sprintf("%v%v", "E", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), "Cash")
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "F", rowNo), fmt.Sprintf("%v%v", "G", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), "Cheque")
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "H", rowNo), fmt.Sprintf("%v%v", "I", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "H", rowNo), "NetBanking")
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "J", rowNo), fmt.Sprintf("%v%v", "K", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "J", rowNo), "DD")
	rowNo++
	//
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "M", rowNo), style1)
	// excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "D", rowNo), style1)
	//excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A7", rowNo), fmt.Sprintf("%v%v", "A9", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "S.No")
	//	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "B7", rowNo), fmt.Sprintf("%v%v", "B9", rowNo), style1)
	//excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "Date")
	//	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C7", rowNo), fmt.Sprintf("%v%v", "C9", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "UserName")
	//	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "D7", rowNo), fmt.Sprintf("%v%v", "D9", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "UserType")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), "HoldingNo")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), "Amount")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), "HoldingNo")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), "Amount")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "H", rowNo), "HoldingNo")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "I", rowNo), "Amount")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "J", rowNo), "HoldingNo")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "K", rowNo), "Amount")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "L", rowNo), "Total Count")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "M", rowNo), "Total Amount")
	//excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "K", rowNo), "Rejected By")

	fmt.Println("'res length==>'", len(res))
	var amount float64
	var count float64
	var totalAmount float64
	var totalCount float64

	for i, v := range res {
		amount = v.TradeLicensePayments.Cash.TotalAmount + v.TradeLicensePayments.Cheque.TotalAmount + v.TradeLicensePayments.NetBanking.TotalAmount + v.TradeLicensePayments.DD.TotalAmount
		count = v.TradeLicensePayments.Cash.Count + v.TradeLicensePayments.Cheque.Count + v.TradeLicensePayments.NetBanking.Count + v.TradeLicensePayments.DD.Count

		rowNo++
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), i+1)
		//excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), sd)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v.Name)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), v.Type)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), v.TradeLicensePayments.Cash.Count)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), v.TradeLicensePayments.Cash.TotalAmount)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), v.TradeLicensePayments.Cheque.Count)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), v.TradeLicensePayments.Cheque.TotalAmount)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "H", rowNo), v.TradeLicensePayments.NetBanking.Count)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "I", rowNo), v.TradeLicensePayments.NetBanking.TotalAmount)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "J", rowNo), v.TradeLicensePayments.DD.Count)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "K", rowNo), v.TradeLicensePayments.DD.TotalAmount)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "L", rowNo), count)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "M", rowNo), amount)
		totalAmount = totalAmount + amount
		totalCount = totalCount + count
	}

	rowNo++
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "Total")

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "B", rowNo), fmt.Sprintf("%v%v", "B", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "D", rowNo), fmt.Sprintf("%v%v", "D", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "E", rowNo), fmt.Sprintf("%v%v", "E", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "F", rowNo), fmt.Sprintf("%v%v", "E", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "G", rowNo), fmt.Sprintf("%v%v", "G", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "H", rowNo), fmt.Sprintf("%v%v", "H", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "I", rowNo), fmt.Sprintf("%v%v", "I", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "J", rowNo), fmt.Sprintf("%v%v", "J", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "K", rowNo), fmt.Sprintf("%v%v", "K", rowNo), style1)
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "L", rowNo), fmt.Sprintf("%v%v", "L", rowNo), style1)
	//excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "M", rowNo), fmt.Sprintf("%v%v", "M", rowNo), style1)

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "L", rowNo), fmt.Sprintf("%v%v", "L", rowNo), style4)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "L", rowNo), fmt.Sprintf("%.0f", totalCount))
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "M", rowNo), fmt.Sprintf("%v%v", "M", rowNo), style4)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "M", rowNo), fmt.Sprintf("%.0f", totalAmount))

	return excel, nil
}
