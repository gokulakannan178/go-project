package service

import (
	"errors"
	"lgf-ccc-service/constants"
	"lgf-ccc-service/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// SaveCitizenGrevience : ""
func (s *Service) SaveCitizenGrevience(ctx *models.Context, CitizenGrevience *models.CitizenGrevience) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	CitizenGrevience.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONCITIZENGREVIENCE)
	CitizenGrevience.Status = constants.CITIZENGREVIENCESTATUSACTIVE
	t := time.Now()
	created := models.CreatedV2{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 CitizenGrevience.created")
	CitizenGrevience.Created = &created
	log.Println("b4 CitizenGrevience.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveCitizenGrevience(ctx, CitizenGrevience)
		if dberr != nil {
			return dberr
		}

		// for _, v := range CitizenGrevience.LeaveMasterId {
		// 	policyrule := new(models.PolicyRule)
		// 	leavemaster, err := s.Daos.GetSingleLeaveMasterWithActive(ctx, v, constants.CitizenGrevienceSTATUSACTIVE)
		// 	if err != nil {
		// 		fmt.Println(err)
		// 		return err
		// 	}
		// 	if leavemaster == nil {
		// 		return errors.New("Leave master id is not available" + err.Error())
		// 	}
		// 	fmt.Println("leavemaster", leavemaster)
		// 	policyrule.CitizenGrevienceID = CitizenGrevience.UniqueID
		// 	policyrule.LeaveMasterID = v
		// 	policyrule.Name = CitizenGrevience.Name
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

// GetSingleCitizenGrevience : ""
func (s *Service) GetSingleCitizenGrevience(ctx *models.Context, UniqueID string) (*models.RefCitizenGrevience, error) {
	CitizenGrevience, err := s.Daos.GetSingleCitizenGrevience(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return CitizenGrevience, nil
}

// UpdateCitizenGrevience : ""
func (s *Service) UpdateCitizenGrevience(ctx *models.Context, CitizenGrevience *models.CitizenGrevience) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		// if len(CitizenGrevience.LeaveMasterId) > 0 {
		// 	err := s.Daos.PolicyRuleRemoveNotPresentValue(ctx, CitizenGrevience.UniqueID, CitizenGrevience.LeaveMasterId)
		// 	if err != nil {
		// 		return err
		// 	}
		// }
		// err := s.Daos.PolicyRuleUpsert(ctx, CitizenGrevience.UniqueID, CitizenGrevience.LeaveMasterId, CitizenGrevience.Name)
		// if err != nil {
		// 	return err
		// }

		//fmt.Println("error==>", err)

		err := s.Daos.UpdateCitizenGrevience(ctx, CitizenGrevience)
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

// EnableCitizenGrevience : ""
func (s *Service) EnableCitizenGrevience(ctx *models.Context, uniqueID string) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.EnableCitizenGrevience(ctx, uniqueID)
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

// DisableCitizenGrevience : ""
func (s *Service) DisableCitizenGrevience(ctx *models.Context, uniqueID string) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		debrr := s.Daos.DisableCitizenGrevience(ctx, uniqueID)
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

// DeleteCitizenGrevience : ""
func (s *Service) DeleteCitizenGrevience(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteCitizenGrevience(ctx, UniqueID)
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

// FilterCitizenGrevience : ""
func (s *Service) FilterCitizenGrevience(ctx *models.Context, CitizenGrevience *models.FilterCitizenGrevience, pagination *models.Pagination) (CitizenGreviences []models.RefCitizenGrevience, err error) {
	return s.Daos.FilterCitizenGrevience(ctx, CitizenGrevience, pagination)
}
