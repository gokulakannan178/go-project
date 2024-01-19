package service

import (
	"errors"

	"lgf-ccc-service/constants"
	"lgf-ccc-service/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//SaveDumpSite :""
func (s *Service) SaveDumpSite(ctx *models.Context, DumpSite *models.DumpSite) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	DumpSite.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONDUMPSITE)
	DumpSite.Status = constants.DUMPSITESTATUSACTIVE
	t := time.Now()
	DumpSite.Date = &t
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 DumpSite.created")
	DumpSite.Created = created
	log.Println("b4 DumpSite.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveDumpSite(ctx, DumpSite)
		if dberr != nil {
			return dberr
		}
		// for k, v := range DumpSite.DumpSitePropertysId {
		// 	//DumpSitepropertys := new(models.DumpSitePropertys)
		// 	// DumpSiteTypeProperty, err := s.Daos.GetSingleDumpSiteTypePropertysWithActive(ctx, v, constants.DumpSiteTYPEPROPERTYSSTATUSACTIVE)
		// 	// if err != nil {
		// 	// 	fmt.Println(err)
		// 	// 	return err
		// 	// }

		// 	//fmt.Println("onboardingchecklistmaster=======", DumpSiteTypeProperty)

		// 	DumpSite.DumpSitePropertysId[k].DumpSiteTypeID = DumpSite.DumpSiteTypeId
		// 	DumpSite.DumpSitePropertysId[k].DumpSiteID = DumpSite.UniqueID
		// 	DumpSite.DumpSitePropertysId[k].Name = DumpSite.Name
		// 	DumpSite.DumpSitePropertysId[k].DumpSitePropertyId = v.DumpSitePropertyId
		// 	DumpSite.DumpSitePropertysId[k].OrganisationID = DumpSite.OrganisationID
		// 	DumpSite.DumpSitePropertysId[k].Description = DumpSite.Description
		// 	err := s.SaveDumpSitePropertysWithoutTransaction(ctx, &DumpSite.DumpSitePropertysId[k])
		// 	if err != nil {
		// 		return err
		// 	}

		// }
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

//GetSingleDumpSite :""
func (s *Service) GetSingleDumpSite(ctx *models.Context, UniqueID string) (*models.RefDumpSite, error) {
	DumpSite, err := s.Daos.GetSingleDumpSite(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return DumpSite, nil
}

// GetSingleDumpSiteUsingEmpID : ""
// func (s *Service) GetSingleDumpSiteUsingEmpID(ctx *models.Context, UniqueID string) (*models.RefDumpSite, error) {
// 	DumpSite, err := s.Daos.GetSingleDumpSiteUsingEmpID(ctx, UniqueID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return DumpSite, nil
// }

//UpdateDumpSite : ""
func (s *Service) UpdateDumpSite(ctx *models.Context, DumpSite *models.DumpSite) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		//var uniqueIds []string
		// for k, v := range DumpSite.DumpSitePropertysId {
		// 	fmt.Println("DumpSitePropertysId===>", v.UniqueID)
		// 	if v.UniqueID == "" {
		// 		DumpSite.DumpSitePropertysId[k].UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONDumpSitePROPERTYS)

		// 	}
		// 	uniqueIds = append(uniqueIds, v.UniqueID)

		// }

		err := s.Daos.UpdateDumpSite(ctx, DumpSite)
		if err := ctx.Session.CommitTransaction(sc); err != nil {
			return errors.New("Not able to commit - " + err.Error())
		}
		// err = s.Daos.DumpSitePropertysRemoveNotPresentValue(ctx, DumpSite.UniqueID, uniqueIds)
		if err != nil {
			return err
		}
		// err = s.Daos.DumpSitePropertysUpsert(ctx, DumpSite)
		// if err != nil {
		// 	return err
		// }
		return nil
	}); err != nil {
		if err = ctx.Session.AbortTransaction(ctx.CTX); err != nil {
			return errors.New("Transaction Aborted with error" + err.Error())
		}
		return errors.New("Transaction Aborted - " + err.Error())
	}
	return nil
}

//EnableDumpSite : ""
func (s *Service) EnableDumpSite(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableDumpSite(ctx, UniqueID)
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

//DisableDumpSite : ""
func (s *Service) DisableDumpSite(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableDumpSite(ctx, UniqueID)
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

//DeleteDumpSite : ""
func (s *Service) DeleteDumpSite(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteDumpSite(ctx, UniqueID)
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

//FilterDumpSite :""
func (s *Service) FilterDumpSite(ctx *models.Context, DumpSiteFilter *models.FilterDumpSite, pagination *models.Pagination) ([]models.RefDumpSite, error) {
	defer ctx.Session.EndSession(ctx.CTX)
	return s.Daos.FilterDumpSite(ctx, DumpSiteFilter, pagination)

}

// func (s *Service) DumpSiteAssign(ctx *models.Context, DumpSite *models.DumpSiteAssign) error {
// 	log.Println("transaction start")
// 	//Start Transaction
// 	if err := ctx.Session.StartTransaction(); err != nil {
// 		return err
// 	}
// 	defer ctx.Session.EndSession(ctx.CTX)
// 	t := time.Now()
// 	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
// 		refDumpSite, err := s.Daos.GetSingleDumpSiteUsingUniqueId(ctx, DumpSite.DumpSiteId)
// 		if err != nil {
// 			return errors.New("error in getting the DumpSitelog- " + err.Error())
// 		}
// 		fmt.Println("DumpSiteid=============", refDumpSite)
// 		if refDumpSite != nil {
// 			refDumpSite.EmployeeId = DumpSite.EmployeeId
// 			//refDumpSite.Status = constants.DumpSiteASSIGNSTATUS
// 			err = s.Daos.UpdateDumpSite(ctx, &refDumpSite.DumpSite)
// 			if err != nil {
// 				return errors.New("error in updating the DumpSitelog" + err.Error())
// 			}
// 		}

// 		//Employee
// 		refDumpSiteLog, err := s.Daos.GetSingleDumpSiteLogUsingEmpID(ctx, DumpSite.DumpSiteId)
// 		if err != nil {
// 			return errors.New("error in getting the DumpSitelog- " + err.Error())
// 		}

// 		if refDumpSiteLog != nil {
// 			refDumpSiteLog.Status = constants.DumpSiteREVOKESTATUS
// 			refDumpSiteLog.EndDate = &t
// 			err = s.Daos.UpdateDumpSiteLog(ctx, &refDumpSiteLog.DumpSiteLog)
// 			if err != nil {
// 				return errors.New("error in updating the DumpSitelog" + err.Error())
// 			}
// 		}

// 		DumpSiteLog := new(models.DumpSiteLog)
// 		//DumpSiteLog.Name = DumpSite.Name
// 		DumpSiteLog.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONDumpSiteLOG)
// 		//DumpSiteLog.OrganisationID = DumpSite.OrganisationID
// 		DumpSiteLog.EmployeeId = DumpSite.EmployeeId
// 		DumpSiteLog.Action.UserID = DumpSite.AssignId
// 		DumpSiteLog.DumpSiteId = DumpSite.DumpSiteId
// 		DumpSiteLog.Action.Date = &t
// 		DumpSiteLog.Status = constants.DumpSiteASSIGNSTATUS
// 		DumpSiteLog.Remark = DumpSite.Remark
// 		//DumpSiteLog.IsLog = constants.DumpSiteASSIGNSTATUSYES
// 		// DumpSiteLog.DumpSiteId = refDumpSiteLog.DumpSiteId
// 		DumpSiteLog.StartDate = &t
// 		err = s.Daos.SaveDumpSiteLog(ctx, DumpSiteLog)
// 		if err != nil {
// 			return err

// 		}
// 		// if refDumpSiteLog == nil {

// 		// }
// 		dberr := s.Daos.DumpSiteAssign(ctx, DumpSite)
// 		if dberr != nil {
// 			return dberr
// 		}
// 		if err := ctx.Session.CommitTransaction(sc); err != nil {
// 			return errors.New("Not able to commit - " + err.Error())
// 		}
// 		return nil
// 	}); err != nil {
// 		log.Println("Transaction start aborting")
// 		if abortError := ctx.Session.AbortTransaction(ctx.CTX); abortError != nil {
// 			return errors.New("Error while aborting transaction" + abortError.Error())
// 		}
// 		log.Println("Transaction aborting completed successfully")
// 		return err
// 	}

// 	return nil
// }

// func (s *Service) RevokeDumpSite(ctx *models.Context, DumpSite *models.DumpSite) error {
// 	if err := ctx.Session.StartTransaction(); err != nil {
// 		return err
// 	}
// 	defer ctx.Session.EndSession(ctx.CTX)
// 	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

// 		err := s.Daos.RevokeDumpSite(ctx, DumpSite)
// 		if err != nil {
// 			if err = ctx.Session.AbortTransaction(sc); err != nil {
// 				return errors.New("Transaction Aborted with error" + err.Error())
// 			}
// 			return errors.New("Transaction Aborted - " + err.Error())
// 		}
// 		return nil

// 	}); err != nil {
// 		return err
// 	}
// 	return nil
// }
