package service

import (
	"errors"
	"lgf-ccc-service/constants"
	"lgf-ccc-service/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// SaveMySurvey : ""
func (s *Service) SaveMySurvey(ctx *models.Context, mySurvey *models.MySurvey) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	mySurvey.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONMYSURVEY)
	mySurvey.Status = constants.MYSURVEYSTATUSACTIVE
	t := time.Now()
	created := models.CreatedV2{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 MySurvey.created")
	mySurvey.Created = &created
	log.Println("b4 MySurvey.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveMySurvey(ctx, mySurvey)
		if dberr != nil {
			return dberr
		}

		// for _, v := range MySurvey.LeaveMasterId {
		// 	policyrule := new(models.PolicyRule)
		// 	leavemaster, err := s.Daos.GetSingleLeaveMasterWithActive(ctx, v, constants.MySurveySTATUSACTIVE)
		// 	if err != nil {
		// 		fmt.Println(err)
		// 		return err
		// 	}
		// 	if leavemaster == nil {
		// 		return errors.New("Leave master id is not available" + err.Error())
		// 	}
		// 	fmt.Println("leavemaster", leavemaster)
		// 	policyrule.MySurveyID = MySurvey.UniqueID
		// 	policyrule.LeaveMasterID = v
		// 	policyrule.Name = MySurvey.Name
		// 	//policyrule.Value = leavemaster.Value
		// 	err = s.SavePolicyRuleWithoutTransaction(ctx, policyrule)
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

// GetSingleMySurvey : ""
func (s *Service) GetSingleMySurvey(ctx *models.Context, UniqueID string) (*models.RefMySurvey, error) {
	MySurvey, err := s.Daos.GetSingleMySurvey(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return MySurvey, nil
}

//UpdateMySurvey : ""
func (s *Service) UpdateMySurvey(ctx *models.Context, mySurvey *models.MySurvey) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		// if len(MySurvey.LeaveMasterId) > 0 {
		// 	err := s.Daos.PolicyRuleRemoveNotPresentValue(ctx, MySurvey.UniqueID, MySurvey.LeaveMasterId)
		// 	if err != nil {
		// 		return err
		// 	}
		// }
		// err := s.Daos.PolicyRuleUpsert(ctx, MySurvey.UniqueID, MySurvey.LeaveMasterId, MySurvey.Name)
		// if err != nil {
		// 	return err
		// }

		//fmt.Println("error==>", err)

		err := s.Daos.UpdateMySurvey(ctx, mySurvey)
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

// EnableMySurvey : ""
func (s *Service) EnableMySurvey(ctx *models.Context, uniqueID string) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.EnableMySurvey(ctx, uniqueID)
		if dberr != nil {
			return dberr
		}
		if err := sc.CommitTransaction(sc); err != nil {
			return errors.New("Not able to commit - " + err.Error())
		}
		return nil
	}); err != nil {
		if err1 := ctx.Session.AbortTransaction(ctx.CTX); err1 != nil {
			log.Println("err in abort")
			return errors.New("Transaction Aborted with error" + err1.Error())
		}
		return err
	}

	return nil
}

// DisableMySurvey : ""
func (s *Service) DisableMySurvey(ctx *models.Context, uniqueID string) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		debrr := s.Daos.DisableMySurvey(ctx, uniqueID)
		if debrr != nil {
			return debrr
		}
		if err := sc.CommitTransaction(sc); err != nil {
			return errors.New("Not able to commit - " + err.Error())
		}
		return nil
	}); err != nil {
		if err1 := ctx.Session.AbortTransaction(ctx.CTX); err1 != nil {
			log.Println("err in abort")
			return errors.New("Transaction Abort with error" + err1.Error())
		}
		return err
	}
	return nil
}

//DeleteMySurvey : ""
func (s *Service) DeleteMySurvey(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteMySurvey(ctx, UniqueID)
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

// FilterMySurvey : ""
func (s *Service) FilterMySurvey(ctx *models.Context, mySurvey *models.FilterMySurvey, pagination *models.Pagination) (MySurveys []models.RefMySurvey, err error) {
	return s.Daos.FilterMySurvey(ctx, mySurvey, pagination)
}

func (s *Service) UpdateCitizenProperty(ctx *models.Context, mySurvey *models.MySurvey) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		// if len(MySurvey.LeaveMasterId) > 0 {
		// 	err := s.Daos.PolicyRuleRemoveNotPresentValue(ctx, MySurvey.UniqueID, MySurvey.LeaveMasterId)
		// 	if err != nil {
		// 		return err
		// 	}
		// }
		// err := s.Daos.PolicyRuleUpsert(ctx, MySurvey.UniqueID, MySurvey.LeaveMasterId, MySurvey.Name)
		// if err != nil {
		// 	return err
		// }

		//fmt.Println("error==>", err)

		err := s.Daos.UpdateCitizenProperty(ctx, mySurvey)
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
