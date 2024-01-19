package service

import (
	"errors"

	"lgf-ccc-service/constants"
	"lgf-ccc-service/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	//"go.mongodb.org/mongo-DriverDetails/mongo"
)

// SaveDriverDetails :""
func (s *Service) SaveDriverDetails(ctx *models.Context, DriverDetails *models.DriverDetails) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	DriverDetails.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONDRIVER)
	//DriverDetails.UserName = s.Daos.GetUniqueID(ctx, constants.COLLECTIONUSER)
	DriverDetails.Status = constants.DRIVERDETAILSSTATUSACTIVE
	t := time.Now()
	DriverDetails.Date = &t
	created := models.CreatedV2{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 DriverDetails.created")
	DriverDetails.Created = created
	log.Println("b4 DriverDetails.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		user := new(models.User)
		user.UserName = s.Daos.GetUniqueID(ctx, constants.COLLECTIONUSER)
		user.Name = DriverDetails.Name
		user.Mobile = DriverDetails.Mobile
		user.Type = "Driver"
		user.Role = "Driver"
		user.DOB = DriverDetails.DOB
		user.Gender = DriverDetails.Gender
		user.Status = constants.DRIVERDETAILSSTATUSACTIVE
		user.Email = DriverDetails.Email
		user.JoiningDate = DriverDetails.DateofJoining
		user.ProfileImg = DriverDetails.ProfileImg
		dberr := s.Daos.SaveUser(ctx, user)
		if dberr != nil {
			return dberr
		}
		dberr = s.Daos.SaveDriverDetails(ctx, DriverDetails)
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

// GetSingleDriverDetails :""
func (s *Service) GetSingleDriverDetails(ctx *models.Context, UniqueID string) (*models.RefDriverDetails, error) {
	DriverDetails, err := s.Daos.GetSingleDriverDetails(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return DriverDetails, nil
}

// GetSingleDriverDetailsUsingEmpID : ""
// func (s *Service) GetSingleDriverDetailsUsingEmpID(ctx *models.Context, UniqueID string) (*models.RefDriverDetails, error) {
// 	DriverDetails, err := s.Daos.GetSingleDriverDetailsUsingEmpID(ctx, UniqueID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return DriverDetails, nil
// }

// UpdateDriverDetails : ""
func (s *Service) UpdateDriverDetails(ctx *models.Context, DriverDetails *models.DriverDetails) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		//var uniqueIds []string
		// for k, v := range DriverDetails.DriverDetailsPropertysId {
		// 	fmt.Println("DriverDetailsPropertysId===>", v.UniqueID)
		// 	if v.UniqueID == "" {
		// 		DriverDetails.DriverDetailsPropertysId[k].UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONDriverDetailsPROPERTYS)

		// 	}
		// 	uniqueIds = append(uniqueIds, v.UniqueID)

		// }

		err := s.Daos.UpdateDriverDetails(ctx, DriverDetails)
		if err := ctx.Session.CommitTransaction(sc); err != nil {
			return errors.New("Not able to commit - " + err.Error())
		}
		// err = s.Daos.DriverDetailsPropertysRemoveNotPresentValue(ctx, DriverDetails.UniqueID, uniqueIds)
		if err != nil {
			return err
		}
		// err = s.Daos.DriverDetailsPropertysUpsert(ctx, DriverDetails)
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

// EnableDriverDetails : ""
func (s *Service) EnableDriverDetails(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableDriverDetails(ctx, UniqueID)
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

// DisableDriverDetails : ""
func (s *Service) DisableDriverDetails(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableDriverDetails(ctx, UniqueID)
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

// DeleteDriverDetails : ""
func (s *Service) DeleteDriverDetails(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteDriverDetails(ctx, UniqueID)
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

// FilterDriverDetails :""
func (s *Service) FilterDriverDetails(ctx *models.Context, DriverDetailsFilter *models.FilterDriverDetails, pagination *models.Pagination) ([]models.RefDriverDetails, error) {
	defer ctx.Session.EndSession(ctx.CTX)
	return s.Daos.FilterDriverDetails(ctx, DriverDetailsFilter, pagination)

}

// func (s *Service) DriverDetailsAssign(ctx *models.Context, DriverDetails *models.DriverDetailsAssign) error {
// 	log.Println("transaction start")
// 	//Start Transaction
// 	if err := ctx.Session.StartTransaction(); err != nil {
// 		return err
// 	}
// 	defer ctx.Session.EndSession(ctx.CTX)
// 	t := time.Now()
// 	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
// 		refDriverDetails, err := s.Daos.GetSingleDriverDetailsUsingUniqueId(ctx, DriverDetails.DriverDetailsId)
// 		if err != nil {
// 			return errors.New("error in getting the DriverDetailslog- " + err.Error())
// 		}
// 		fmt.Println("DriverDetailsid=============", refDriverDetails)
// 		if refDriverDetails != nil {
// 			refDriverDetails.EmployeeId = DriverDetails.EmployeeId
// 			//refDriverDetails.Status = constants.DriverDetailsASSIGNSTATUS
// 			err = s.Daos.UpdateDriverDetails(ctx, &refDriverDetails.DriverDetails)
// 			if err != nil {
// 				return errors.New("error in updating the DriverDetailslog" + err.Error())
// 			}
// 		}

// 		//Employee
// 		refDriverDetailsLog, err := s.Daos.GetSingleDriverDetailsLogUsingEmpID(ctx, DriverDetails.DriverDetailsId)
// 		if err != nil {
// 			return errors.New("error in getting the DriverDetailslog- " + err.Error())
// 		}

// 		if refDriverDetailsLog != nil {
// 			refDriverDetailsLog.Status = constants.DriverDetailsREVOKESTATUS
// 			refDriverDetailsLog.EndDate = &t
// 			err = s.Daos.UpdateDriverDetailsLog(ctx, &refDriverDetailsLog.DriverDetailsLog)
// 			if err != nil {
// 				return errors.New("error in updating the DriverDetailslog" + err.Error())
// 			}
// 		}

// 		DriverDetailsLog := new(models.DriverDetailsLog)
// 		//DriverDetailsLog.Name = DriverDetails.Name
// 		DriverDetailsLog.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONDriverDetailsLOG)
// 		//DriverDetailsLog.OrganisationID = DriverDetails.OrganisationID
// 		DriverDetailsLog.EmployeeId = DriverDetails.EmployeeId
// 		DriverDetailsLog.Action.UserID = DriverDetails.AssignId
// 		DriverDetailsLog.DriverDetailsId = DriverDetails.DriverDetailsId
// 		DriverDetailsLog.Action.Date = &t
// 		DriverDetailsLog.Status = constants.DriverDetailsASSIGNSTATUS
// 		DriverDetailsLog.Remark = DriverDetails.Remark
// 		//DriverDetailsLog.IsLog = constants.DriverDetailsASSIGNSTATUSYES
// 		// DriverDetailsLog.DriverDetailsId = refDriverDetailsLog.DriverDetailsId
// 		DriverDetailsLog.StartDate = &t
// 		err = s.Daos.SaveDriverDetailsLog(ctx, DriverDetailsLog)
// 		if err != nil {
// 			return err

// 		}
// 		// if refDriverDetailsLog == nil {

// 		// }
// 		dberr := s.Daos.DriverDetailsAssign(ctx, DriverDetails)
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

// func (s *Service) RevokeDriverDetails(ctx *models.Context, DriverDetails *models.DriverDetails) error {
// 	if err := ctx.Session.StartTransaction(); err != nil {
// 		return err
// 	}
// 	defer ctx.Session.EndSession(ctx.CTX)
// 	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

// 		err := s.Daos.RevokeDriverDetails(ctx, DriverDetails)
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
