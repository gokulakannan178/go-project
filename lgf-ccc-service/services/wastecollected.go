package service

import (
	"errors"

	"lgf-ccc-service/constants"
	"lgf-ccc-service/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//SaveWasteCollected :""
func (s *Service) SaveWasteCollected(ctx *models.Context, WasteCollected *models.WasteCollected) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	WasteCollected.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONWASTECOLLECTED)
	WasteCollected.Status = constants.WASTECOLLECTEDSTATUSACTIVE
	t := time.Now()
	WasteCollected.Date = &t
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 WasteCollected.created")
	WasteCollected.Created = created
	log.Println("b4 WasteCollected.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveWasteCollected(ctx, WasteCollected)
		if dberr != nil {
			return dberr
		}
		// for k, v := range WasteCollected.WasteCollectedPropertysId {
		// 	//WasteCollectedpropertys := new(models.WasteCollectedPropertys)
		// 	// WasteCollectedTypeProperty, err := s.Daos.GetSingleWasteCollectedTypePropertysWithActive(ctx, v, constants.WasteCollectedTYPEPROPERTYSSTATUSACTIVE)
		// 	// if err != nil {
		// 	// 	fmt.Println(err)
		// 	// 	return err
		// 	// }

		// 	//fmt.Println("onboardingchecklistmaster=======", WasteCollectedTypeProperty)

		// 	WasteCollected.WasteCollectedPropertysId[k].WasteCollectedTypeID = WasteCollected.WasteCollectedTypeId
		// 	WasteCollected.WasteCollectedPropertysId[k].WasteCollectedID = WasteCollected.UniqueID
		// 	WasteCollected.WasteCollectedPropertysId[k].Name = WasteCollected.Name
		// 	WasteCollected.WasteCollectedPropertysId[k].WasteCollectedPropertyId = v.WasteCollectedPropertyId
		// 	WasteCollected.WasteCollectedPropertysId[k].OrganisationID = WasteCollected.OrganisationID
		// 	WasteCollected.WasteCollectedPropertysId[k].Description = WasteCollected.Description
		// 	err := s.SaveWasteCollectedPropertysWithoutTransaction(ctx, &WasteCollected.WasteCollectedPropertysId[k])
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

//GetSingleWasteCollected :""
func (s *Service) GetSingleWasteCollected(ctx *models.Context, UniqueID string) (*models.RefWasteCollected, error) {
	WasteCollected, err := s.Daos.GetSingleWasteCollected(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return WasteCollected, nil
}

// GetSingleWasteCollectedUsingEmpID : ""
// func (s *Service) GetSingleWasteCollectedUsingEmpID(ctx *models.Context, UniqueID string) (*models.RefWasteCollected, error) {
// 	WasteCollected, err := s.Daos.GetSingleWasteCollectedUsingEmpID(ctx, UniqueID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return WasteCollected, nil
// }

//UpdateWasteCollected : ""
func (s *Service) UpdateWasteCollected(ctx *models.Context, WasteCollected *models.WasteCollected) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		//var uniqueIds []string
		// for k, v := range WasteCollected.WasteCollectedPropertysId {
		// 	fmt.Println("WasteCollectedPropertysId===>", v.UniqueID)
		// 	if v.UniqueID == "" {
		// 		WasteCollected.WasteCollectedPropertysId[k].UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONWasteCollectedPROPERTYS)

		// 	}
		// 	uniqueIds = append(uniqueIds, v.UniqueID)

		// }

		err := s.Daos.UpdateWasteCollected(ctx, WasteCollected)
		if err := ctx.Session.CommitTransaction(sc); err != nil {
			return errors.New("Not able to commit - " + err.Error())
		}
		// err = s.Daos.WasteCollectedPropertysRemoveNotPresentValue(ctx, WasteCollected.UniqueID, uniqueIds)
		if err != nil {
			return err
		}
		// err = s.Daos.WasteCollectedPropertysUpsert(ctx, WasteCollected)
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

//EnableWasteCollected : ""
func (s *Service) EnableWasteCollected(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableWasteCollected(ctx, UniqueID)
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

//DisableWasteCollected : ""
func (s *Service) DisableWasteCollected(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableWasteCollected(ctx, UniqueID)
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

//EnableWasteCollected : ""
func (s *Service) WasteCollectedCompleted(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.WasteCollectedCompleted(ctx, UniqueID)
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

//DisableWasteCollected : ""
func (s *Service) WasteCollectedPending(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.WasteCollectedPending(ctx, UniqueID)
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

//DeleteWasteCollected : ""
func (s *Service) DeleteWasteCollected(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteWasteCollected(ctx, UniqueID)
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

//FilterWasteCollected :""
func (s *Service) FilterWasteCollected(ctx *models.Context, WasteCollectedFilter *models.FilterWasteCollected, pagination *models.Pagination) ([]models.RefWasteCollected, error) {
	defer ctx.Session.EndSession(ctx.CTX)
	return s.Daos.FilterWasteCollected(ctx, WasteCollectedFilter, pagination)

}

// func (s *Service) WasteCollectedAssign(ctx *models.Context, WasteCollected *models.WasteCollectedAssign) error {
// 	log.Println("transaction start")
// 	//Start Transaction
// 	if err := ctx.Session.StartTransaction(); err != nil {
// 		return err
// 	}
// 	defer ctx.Session.EndSession(ctx.CTX)
// 	t := time.Now()
// 	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
// 		refWasteCollected, err := s.Daos.GetSingleWasteCollectedUsingUniqueId(ctx, WasteCollected.WasteCollectedId)
// 		if err != nil {
// 			return errors.New("error in getting the WasteCollectedlog- " + err.Error())
// 		}
// 		fmt.Println("WasteCollectedid=============", refWasteCollected)
// 		if refWasteCollected != nil {
// 			refWasteCollected.EmployeeId = WasteCollected.EmployeeId
// 			//refWasteCollected.Status = constants.WASTECOLLECTEDASSIGNSTATUS
// 			err = s.Daos.UpdateWasteCollected(ctx, &refWasteCollected.WasteCollected)
// 			if err != nil {
// 				return errors.New("error in updating the WasteCollectedlog" + err.Error())
// 			}
// 		}

// 		//Employee
// 		refWasteCollectedLog, err := s.Daos.GetSingleWasteCollectedLogUsingEmpID(ctx, WasteCollected.WasteCollectedId)
// 		if err != nil {
// 			return errors.New("error in getting the WasteCollectedlog- " + err.Error())
// 		}

// 		if refWasteCollectedLog != nil {
// 			refWasteCollectedLog.Status = constants.WasteCollectedREVOKESTATUS
// 			refWasteCollectedLog.EndDate = &t
// 			err = s.Daos.UpdateWasteCollectedLog(ctx, &refWasteCollectedLog.WasteCollectedLog)
// 			if err != nil {
// 				return errors.New("error in updating the WasteCollectedlog" + err.Error())
// 			}
// 		}

// 		WasteCollectedLog := new(models.WasteCollectedLog)
// 		//WasteCollectedLog.Name = WasteCollected.Name
// 		WasteCollectedLog.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONWasteCollectedLOG)
// 		//WasteCollectedLog.OrganisationID = WasteCollected.OrganisationID
// 		WasteCollectedLog.EmployeeId = WasteCollected.EmployeeId
// 		WasteCollectedLog.Action.UserID = WasteCollected.AssignId
// 		WasteCollectedLog.WasteCollectedId = WasteCollected.WasteCollectedId
// 		WasteCollectedLog.Action.Date = &t
// 		WasteCollectedLog.Status = constants.WasteCollectedASSIGNSTATUS
// 		WasteCollectedLog.Remark = WasteCollected.Remark
// 		//WasteCollectedLog.IsLog = constants.WasteCollectedASSIGNSTATUSYES
// 		// WasteCollectedLog.WasteCollectedId = refWasteCollectedLog.WasteCollectedId
// 		WasteCollectedLog.StartDate = &t
// 		err = s.Daos.SaveWasteCollectedLog(ctx, WasteCollectedLog)
// 		if err != nil {
// 			return err

// 		}
// 		// if refWasteCollectedLog == nil {

// 		// }
// 		dberr := s.Daos.WasteCollectedAssign(ctx, WasteCollected)
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

// func (s *Service) RevokeWasteCollected(ctx *models.Context, WasteCollected *models.WasteCollected) error {
// 	if err := ctx.Session.StartTransaction(); err != nil {
// 		return err
// 	}
// 	defer ctx.Session.EndSession(ctx.CTX)
// 	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

// 		err := s.Daos.RevokeWasteCollected(ctx, WasteCollected)
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
