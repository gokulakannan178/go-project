package services

import (
	"errors"
	"fmt"
	"haritv2-service/constants"
	"haritv2-service/models"
	"log"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"

	"go.mongodb.org/mongo-driver/mongo"
)

// SaveFPO : ""
func (s *Service) SaveFPO(ctx *models.Context, fpo *models.FPO) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	fpo.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONFPO)
	fpo.Status = constants.FPOSTATUSACTIVE
	t := time.Now()
	created := new(models.CreatedV2)
	created.On = &t
	created.By = constants.SYSTEM
	fpo.Created = created
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveFPO(ctx, fpo)
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

// SaveFPOregister
func (s *Service) SaveFPORegistration(ctx *models.Context, fpo *models.FPO) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	user := new(models.User)
	fpo.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONFPO)
	fpo.Status = constants.FPOSTATUSACTIVE
	t := time.Now()
	created := new(models.CreatedV2)
	created.On = &t
	created.By = constants.SYSTEM
	fpo.Created = created
	user.Name = fpo.ChairMan
	user.CompanyID = fpo.UniqueID
	user.Mobile = fpo.Mobile
	user.Email = fpo.Email
	user.Type = constants.USERTYPE
	user.Status = constants.USERSTATUSACTIVE
	user.UserName = s.Daos.GetUniqueID(ctx, constants.FPOCOLLECTIONUSER)
	user.Address = fpo.Address
	fpoinven := new(models.FPOInventory)
	fpoinven.CompanyID = fpo.UniqueID
	fpoinven.UniqueID = s.Daos.GetUniqueID(ctx, constants.FPOINVENTORYCOLLECTION)
	fpoinven.BuyingPrice = 0
	fpoinven.Quantity = 0
	fpoinven.Sellingprice = 0

	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveFPO(ctx, fpo)
		if dberr != nil {
			return errors.New("fposaving " + dberr.Error())
		}
		dberr = s.Daos.SaveUser(ctx, user)
		if dberr != nil {
			return errors.New("usersaving " + dberr.Error())
		}
		dberr = s.Daos.SaveFPOInventory(ctx, fpoinven)
		if dberr != nil {
			return errors.New("fpoSavingInventory " + dberr.Error())
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

//GetSingleFPO :""
func (s *Service) GetSingleFPO(ctx *models.Context, UniqueID string) (*models.RefFPO, error) {
	tower, err := s.Daos.GetSingleFPO(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return tower, nil
}

// UpdateFPO : ""
func (s *Service) UpdateFPO(ctx *models.Context, fpo *models.FPO) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateFPO(ctx, fpo)
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

// EnableFPO : ""
func (s *Service) EnableFPO(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableFPO(ctx, UniqueID)
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

//DisableFPO : ""
func (s *Service) DisableFPO(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableFPO(ctx, UniqueID)
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

//DeleteFPO : ""
func (s *Service) DeleteFPO(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteFPO(ctx, UniqueID)
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

// FilterFPO : ""
func (s *Service) FilterFPO(ctx *models.Context, filter *models.FPOFilter, pagination *models.Pagination) ([]models.RefFPO, error) {
	return s.Daos.FilterFPO(ctx, filter, pagination)

}

// SaveFPOregister
func (s *Service) UpdateFPORegistration(ctx *models.Context, fpo *models.FPO) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		user := new(models.User)
		user.Name = fpo.ChairMan
		user.Mobile = fpo.Mobile
		user.Email = fpo.Email
		user.Type = constants.USERTYPE
		user.CompanyID = fpo.UniqueID
		user.Address = fpo.Address
		dberr := s.Daos.UpdateFPO(ctx, fpo)
		if dberr != nil {
			return errors.New("fposaving " + dberr.Error())
		}
		dberr = s.Daos.UpdateUserWithCompanyId(ctx, user)
		if dberr != nil {
			return errors.New("usersaving " + dberr.Error())
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

//ULBMasterReportJSON  : ""
func (s *Service) FPOMasterReportJSON(ctx *models.Context, filter *models.FPOReportFilter, pagination *models.Pagination) ([]models.FPOReport, error) {
	return s.Daos.FPOMasterReport(ctx, filter, pagination)
}

func (s *Service) FPOMasterReportExcel(ctx *models.Context, filter *models.FPOReportFilter, pagination *models.Pagination) (*excelize.File, error) {
	res, err := s.FPOMasterReportJSON(ctx, filter, pagination)
	if err != nil {
		return nil, err
	}
	fmt.Println("'res length==>'", res)

	//  create an excel file
	excel := excelize.NewFile()
	sheet1 := "FPOs"
	rowNo := 1
	index := excel.NewSheet(sheet1)
	excel.SetActiveSheet(index)
	// excel.MergeCell(sheet1, "A1", "B5")Compost Purchased For How Many ulbs till Date
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "S.No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "UinqueID")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "FPO Name")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), "District")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), "Registration Date")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), "Chairperson Name")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), "Mobile")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "H", rowNo), "Total Compost Purchased Till Date (MT)")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "I", rowNo), "Total Cost of Compost Purchased Till Date (MT)")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "J", rowNo), "Compost Purchased For How Many ulbs till Date")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "K", rowNo), "Compost Purchased in Current Month(MT)")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "L", rowNo), "Total Cost of Compost  in Current Month(INR)")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "M", rowNo), "Compost Purchased From Current Month")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "N", rowNo), "Current Pending Orders, If Any (numbers)")
	rowNo++
	if res == nil {
		// err := excel.MergeCell(sheet1, "A2", "Y2")
		// if err != nil {
		// 	fmt.Println("merge error " + err.Error())
		// }
		excel.SetCellValue(sheet1, fmt.Sprintf("%v", "A2"), "No Data")
		return excel, nil
	}
	for k, v := range res {

		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), k+1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v.UniqueID)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), v.Name)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), v.Ref.Address.District.Name)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), v.Created.On.Format("2006-January-02"))
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), v.ChairMan)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), v.Mobile)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "H", rowNo), v.CompostPurchasedTillDate.Quantity)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "I", rowNo), v.CompostPurchasedTillDate.Amount)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "J", rowNo), v.CompostPurchasedTillDate.ULbs)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "K", rowNo), v.CompostPurchasedCurrMonth.Quantity)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "L", rowNo), v.CompostPurchasedCurrMonth.Amount)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "M", rowNo), v.CompostPurchasedCurrMonth.ULbs)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "N", rowNo), v.PendingOrders.ULbs)

		rowNo++
	}
	return excel, nil
}
func (s *Service) FPOMonthReportJSON(ctx *models.Context, filter *models.FPOMothWiseeportFilter) ([]models.FPOMothWiseeport, error) {
	return s.Daos.FPOMonthReport(ctx, filter)
}
func (s *Service) FPOMonthReportExcel(ctx *models.Context, filter *models.FPOMothWiseeportFilter) (*excelize.File, error) {
	res, err := s.FPOMonthReportJSON(ctx, filter)
	if err != nil {
		return nil, err
	}
	fmt.Println("'res length==>'", res)

	//  create an excel file
	excel := excelize.NewFile()
	sheet1 := "FPOsMonthWiseReport"
	rowNo := 1
	index := excel.NewSheet(sheet1)
	excel.SetActiveSheet(index)
	// excel.MergeCell(sheet1, "A1", "B5")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "S.No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "UinqueID")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "FPO Name")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), "Chairperson Name")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), "Chairperson MobileNo")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), "Customers Count")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), "Sales Amount")
	rowNo++
	if res == nil {
		// err := excel.MergeCell(sheet1, "A2", "Y2")
		// if err != nil {
		// 	fmt.Println("merge error " + err.Error())
		// }
		excel.SetCellValue(sheet1, fmt.Sprintf("%v", "A2"), "No Data")
		return excel, nil
	}
	for k, v := range res {

		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), k+1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v.UniqueID)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), v.Name)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), v.ChairMan)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), v.Mobile)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), v.Sale.TotalCustomers)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), v.Sale.TotalsaleAmount)
		// excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), v.Ref.Inventory.Quantity)
		rowNo++
	}
	return excel, nil
}

func (s *Service) FBONearBy(ctx *models.Context, fbonb *models.FBONearBy, pagination *models.Pagination) ([]models.RefFPO, error) {

	ulbs, err := s.Daos.FBONearBy(ctx, fbonb, pagination)
	if err != nil {
		return nil, err
	}

	return ulbs, nil
}
