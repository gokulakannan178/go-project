package services

import (
	"context"
	"errors"
	"fmt"
	"log"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"go.mongodb.org/mongo-driver/mongo"
)

// SavePropertyFixedArv : ""
func (s *Service) SavePropertyFixedArv(ctx *models.Context, propertyfixedarv *models.PropertyFixedArv) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	propertyfixedarv.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONPROPERTYFIXEDARV)
	// propertyfixedarv.Status = constants.PROPERTYFIXEDARVSTATUSACTIVE
	propertyfixedarv.Status = constants.PROPERTYFIXEDARVSTATUSINIT
	t := time.Now()
	propertyfixedarv.Created.On = &t
	//PropertyFixedArv.Created.By = constants.SYSTEM
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		FyID, err := s.Daos.GetSinglePropertyFixedArvWithFyID(ctx, propertyfixedarv.FyID, propertyfixedarv.PropertyID)
		if err != nil {
			return nil
		}
		fmt.Println("fyid=================>", FyID)
		if FyID == nil {
			t := time.Now()
			propertyfixedarvLog := new(models.PropertyFixedArvLog)
			propertyfixedarvLog.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONPROPERTYFIXEDARVLOG)
			propertyfixedarvLog.Status = constants.PROPERTYFIXEDARVLOGSTATUSACTIVE
			propertyfixedarvLog.FyID = propertyfixedarv.FyID
			propertyfixedarvLog.PropertyID = propertyfixedarv.PropertyID
			propertyfixedarvLog.ARV = propertyfixedarv.ARV
			propertyfixedarvLog.Tax = propertyfixedarv.Tax
			propertyfixedarvLog.Total = propertyfixedarv.Total
			propertyfixedarvLog.Created.On = &t
			propertyfixedarvLog.Created.By = propertyfixedarv.Created.By
			propertyfixedarvLog.Created.ByType = propertyfixedarv.Created.ByType
			dberr := s.Daos.SavePropertyFixedArvLog(ctx, propertyfixedarvLog)
			if dberr != nil {
				return dberr
			}
			dberr = s.Daos.SavePropertyFixedArv(ctx, propertyfixedarv)
			if dberr != nil {
				return dberr
			}

		} else {
			return errors.New("FyId Already available, please change FyId")
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

// SavePropertyFixedArv : ""
func (s *Service) UpsertPropertyFixedArv(ctx *models.Context, propertyfixedarv *models.PropertyFixedArv) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	t := time.Now()

	financialYear := new(models.FinancialYearFilter)
	financialYear.DateRange = new(models.DateRange)
	financialYear.DateRange.From = propertyfixedarv.FyFrom
	financialYear.DateRange.To = propertyfixedarv.FyTo
	//PropertyFixedArv.Created.By = constants.SYSTEM
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		resFYs, err1 := s.FilterFinancialYear(ctx, financialYear, nil)
		if err1 != nil {
			return err1
		}
		for _, v := range resFYs {
			propertyfixedarv.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONPROPERTYFIXEDARV)
			propertyfixedarv.Status = constants.PROPERTYFIXEDARVSTATUSINIT
			// propertyfixedarv.Status = constants.PROPERTYFIXEDARVSTATUSACTIVE
			propertyfixedarv.Created.On = &t
			propertyfixedarv.PropertyID = propertyfixedarv.PropertyID
			propertyfixedarv.FyID = v.FinancialYear.UniqueID
			propertyfixedarv.FyFrom = propertyfixedarv.FyFrom
			propertyfixedarv.FyTo = propertyfixedarv.FyTo
			propertyfixedarv.Requester.On = &t

			err2 := s.Daos.UpsertPropertyFixedArv(ctx, propertyfixedarv)
			if err2 != nil {
				return err2
			}
		}
		propertyfixedarvLog := new(models.PropertyFixedArvLog)
		propertyfixedarvLog.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONPROPERTYFIXEDARVLOG)
		propertyfixedarvLog.Status = constants.PROPERTYFIXEDARVLOGSTATUSACTIVE
		propertyfixedarvLog.FyID = propertyfixedarv.FyID
		propertyfixedarvLog.PropertyID = propertyfixedarv.PropertyID
		propertyfixedarvLog.ARV = propertyfixedarv.ARV
		propertyfixedarvLog.Tax = propertyfixedarv.Tax
		propertyfixedarvLog.Total = propertyfixedarv.Total
		propertyfixedarvLog.Created.On = &t
		propertyfixedarvLog.Created.By = propertyfixedarv.Created.By
		propertyfixedarvLog.Created.ByType = propertyfixedarv.Created.ByType
		dberr := s.Daos.SavePropertyFixedArvLog(ctx, propertyfixedarvLog)
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

//GetSinglePropertyFixedArv :""
func (s *Service) GetSinglePropertyFixedArv(ctx *models.Context, UniqueID string) (*models.RefPropertyFixedArv, error) {
	propertyfixedarv, err := s.Daos.GetSinglePropertyFixedArv(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return propertyfixedarv, nil
}

// UpdatePropertyFixedArv : ""
func (s *Service) UpdatePropertyFixedArv(ctx *models.Context, propertyfixedarv *models.PropertyFixedArv) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		t := time.Now()
		FyID, err := s.Daos.GetSinglePropertyFixedArv(ctx, propertyfixedarv.UniqueID)
		if err != nil {
			return nil
		}
		if FyID != nil {

			propertyfixedarvLog := new(models.PropertyFixedArvLog)
			propertyfixedarvLog.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONPROPERTYFIXEDARVLOG)
			propertyfixedarv.Status = constants.PROPERTYFIXEDARVSTATUSINIT
			// propertyfixedarv.Status = constants.PROPERTYFIXEDARVSTATUSACTIVE
			propertyfixedarvLog.Status = constants.PROPERTYFIXEDARVLOGSTATUSACTIVE
			propertyfixedarvLog.FyID = FyID.FyID
			propertyfixedarvLog.PropertyID = FyID.PropertyID
			propertyfixedarvLog.ARV = FyID.ARV
			propertyfixedarvLog.Tax = FyID.Tax
			propertyfixedarvLog.Total = FyID.Total
			propertyfixedarvLog.Created.On = FyID.Created.On
			propertyfixedarvLog.Created.By = FyID.Created.By
			propertyfixedarvLog.Created.ByType = FyID.Created.ByType
			dberr := s.Daos.SavePropertyFixedArvLog(ctx, propertyfixedarvLog)
			if dberr != nil {
				return dberr
			}
		}
		propertyfixedarv.Requester.On = &t

		err = s.Daos.UpdatePropertyFixedArv(ctx, propertyfixedarv)
		if err != nil {
			return err
		}
		if err := ctx.Session.CommitTransaction(sc); err != nil {
			return errors.New("Not able to commit - " + err.Error())
		}
		propertyfixedarvLog := new(models.PropertyFixedArvLog)
		propertyfixedarvLog.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONPROPERTYFIXEDARVLOG)
		propertyfixedarvLog.Status = constants.PROPERTYFIXEDARVLOGSTATUSACTIVE
		propertyfixedarvLog.FyID = propertyfixedarv.FyID
		propertyfixedarvLog.PropertyID = propertyfixedarv.PropertyID
		propertyfixedarvLog.ARV = propertyfixedarv.ARV
		propertyfixedarvLog.Tax = propertyfixedarv.Tax
		propertyfixedarvLog.Total = propertyfixedarv.Total
		propertyfixedarvLog.Created.On = &t
		propertyfixedarvLog.Created.By = propertyfixedarv.Created.By
		propertyfixedarvLog.Created.ByType = propertyfixedarv.Created.ByType
		dberr := s.Daos.SavePropertyFixedArvLog(ctx, propertyfixedarvLog)
		if dberr != nil {
			return dberr
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

// EnablePropertyFixedArv : ""
func (s *Service) EnablePropertyFixedArv(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnablePropertyFixedArv(ctx, UniqueID)
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

//DisablePropertyFixedArv : ""
func (s *Service) DisablePropertyFixedArv(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisablePropertyFixedArv(ctx, UniqueID)
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

//DeletePropertyFixedArv : ""
func (s *Service) DeletePropertyFixedArv(ctx *models.Context, uniqueId string, by string, byType string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		t := time.Now()
		FyID, err := s.Daos.GetSinglePropertyFixedArv(ctx, uniqueId)
		if err != nil {
			return nil
		}
		propertyfixedarvLog := new(models.PropertyFixedArvLog)
		propertyfixedarvLog.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONPROPERTYFIXEDARVLOG)
		propertyfixedarvLog.Status = constants.PROPERTYFIXEDARVLOGSTATUSACTIVE
		propertyfixedarvLog.FyID = FyID.FyID
		propertyfixedarvLog.PropertyID = FyID.PropertyID
		propertyfixedarvLog.ARV = FyID.ARV
		propertyfixedarvLog.Tax = FyID.Tax
		propertyfixedarvLog.Total = FyID.Total
		propertyfixedarvLog.Created.On = &t
		propertyfixedarvLog.Created.By = by
		propertyfixedarvLog.Created.ByType = byType
		dberr := s.Daos.SavePropertyFixedArvLog(ctx, propertyfixedarvLog)
		if dberr != nil {
			return dberr
		}

		err = s.Daos.DeletePropertyFixedArv(ctx, uniqueId)
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

// FilterPropertyFixedArv : ""
func (s *Service) FilterPropertyFixedArv(ctx *models.Context, filter *models.PropertyFixedArvFilter, pagination *models.Pagination) ([]models.RefPropertyFixedArv, error) {
	return s.Daos.FilterPropertyFixedArv(ctx, filter, pagination)
}

// FilterPropertyFixedArvExcel :""
func (s *Service) FilterPropertyFixedArvExcel(ctx *models.Context, filter *models.PropertyFixedArvFilter, pagination *models.Pagination) (*excelize.File, error) {
	res, err := s.FilterPropertyFixedArv(ctx, filter, pagination)
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
	sheet1 := "Property Fixed ARV Report"
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
	// excel.MergeCell(sheet1, "A8", "G8")

	style3, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#B6DDE8"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
	if err != nil {
		fmt.Println(err)
	}
	style1, err := excel.NewStyle(`{"fill":{"type":"pattern","color":["#FFDC6D"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`)
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
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "Property Fixed ARV Report")
	} else {
		excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "C", rowNo), fmt.Sprintf("%v%v", "C", rowNo), style3)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "Property Fixed ARV Report")
	}
	rowNo++
	rowNo++

	reportFromMsg := "Property Fixed ARV Report"
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style3)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), reportFromMsg)
	rowNo++

	t := time.Now()
	toDate := t.Format("02-January-2006")
	reportFromMsg3 := "Report Generated on" + " " + toDate
	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "A", rowNo), style3)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), reportFromMsg3)
	rowNo++

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "G", rowNo), style1)
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), "S.No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), "Property No")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), "FY")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), "Created By")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), "ARV")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), "Tax")
	excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), "Total")
	rowNo++

	for k, v := range res {
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "A", rowNo), k+1)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "B", rowNo), v.PropertyID)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "C", rowNo), func() string {
			if v.Ref.FinancialYear.Name != "" {
				return v.Ref.FinancialYear.Name
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "D", rowNo), func() string {
			if v.Created.By != "" {
				return v.Created.By
			}
			return "NA"
		}())
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "E", rowNo), v.ARV)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "F", rowNo), v.Tax)
		excel.SetCellValue(sheet1, fmt.Sprintf("%v%v", "G", rowNo), v.Total)

		rowNo++

	}

	excel.SetCellStyle(sheet1, fmt.Sprintf("%v%v", "A", rowNo), fmt.Sprintf("%v%v", "G", rowNo), style1)

	return excel, nil

}

// RejectPropertyFixedArv : ""
func (s *Service) RejectPropertyFixedArv(ctx *models.Context, req *models.RejectPropertyFixedArv) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.RejectPropertyFixedArv(ctx, req)
		if err != nil {
			return nil
		}
		if err = ctx.Session.CommitTransaction(sc); err != nil {
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

// AcceptPropertyFixedArv : ""
func (s *Service) AcceptPropertyFixedArv(ctx *models.Context, req *models.AcceptPropertyFixedArv) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	client := s.Daos.GetDBV3(context.TODO())
	defer client.Disconnect(context.TODO())

	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		err := s.Daos.AcceptPropertyFixedArv(ctx, req)
		if err != nil {
			return errors.New("Error in upating in Property Fixed Arv" + err.Error())
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

// AcceptMultiplePropertyFixedArv : ""
func (s *Service) AcceptMultiplePropertyFixedArv(ctx *models.Context, req *models.AcceptMultiplePropertyFixedArv) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	client := s.Daos.GetDBV3(context.TODO())
	defer client.Disconnect(context.TODO())

	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		for _, v := range req.UniqueID {
			accept := new(models.AcceptPropertyFixedArv)
			accept.UserName = req.UserName
			accept.UserType = req.UserType
			accept.Remark = req.Remark
			accept.UniqueID = v
			err := s.Daos.AcceptPropertyFixedArv(ctx, accept)
			if err != nil {
				return errors.New("Error in upating in Property Fixed Arv" + err.Error())
			}
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
