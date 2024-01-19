package services

import (
	"errors"
	"hrms-services/constants"
	"hrms-services/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// SaveOnboardingCheckListMaster : ""
func (s *Service) SaveOnboardingCheckListMaster(ctx *models.Context, onboardingchecklistmaster *models.OnboardingCheckListMaster) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	onboardingchecklistmaster.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONONBOARDINGCHECKLISTMASTER)
	onboardingchecklistmaster.Status = constants.ONBOARDINGCHECKLISTMASTERSTATUSACTIVE
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 OnboardingCheckListMaster.created")
	onboardingchecklistmaster.Created = &created
	log.Println("b4 OnboardingCheckListMaster.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveOnboardingCheckListMaster(ctx, onboardingchecklistmaster)
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

// GetSingleOnboardingCheckListMaster : ""
func (s *Service) GetSingleOnboardingCheckListMaster(ctx *models.Context, UniqueID string) (*models.RefOnboardingCheckListMaster, error) {
	onboardingchecklistmaster, err := s.Daos.GetSingleOnboardingCheckListMaster(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return onboardingchecklistmaster, nil
}

//UpdateOnboardingCheckListMaster : ""
func (s *Service) UpdateOnboardingCheckListMaster(ctx *models.Context, onboardingchecklistmaster *models.OnboardingCheckListMaster) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateOnboardingCheckListMaster(ctx, onboardingchecklistmaster)
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

// EnableOnboardingCheckListMaster : ""
func (s *Service) EnableOnboardingCheckListMaster(ctx *models.Context, uniqueID string) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.EnableOnboardingCheckListMaster(ctx, uniqueID)
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

// DisableOnboardingCheckListMaster : ""
func (s *Service) DisableOnboardingCheckListMaster(ctx *models.Context, uniqueID string) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		debrr := s.Daos.DisableOnboardingCheckListMaster(ctx, uniqueID)
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

//DeleteOnboardingCheckListMaster : ""
func (s *Service) DeleteOnboardingCheckListMaster(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteOnboardingCheckListMaster(ctx, UniqueID)
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

// FilterOnboardingCheckListMaster : ""
func (s *Service) FilterOnboardingCheckListMaster(ctx *models.Context, onboardingchecklistmaster *models.FilterOnboardingCheckListMaster, pagination *models.Pagination) (onboardingchecklistmasters []models.RefOnboardingCheckListMaster, err error) {
	err = s.OnboardingCheckListMasterDataAccess(ctx, onboardingchecklistmaster)
	if err != nil {
		return nil, err
	}
	return s.Daos.FilterOnboardingCheckListMaster(ctx, onboardingchecklistmaster, pagination)
}
func (s *Service) OnboardingCheckListMasterDataAccess(ctx *models.Context, filter *models.FilterOnboardingCheckListMaster) (err error) {
	if filter != nil {
		dataaccess, err := s.Daos.DataAccess(ctx, &filter.DataAccess)
		if err != nil {
			return err
		}
		if dataaccess != nil {
			if len(dataaccess.Organisation) > 0 {
				for _, v := range dataaccess.Organisation {
					filter.OrganisationID = append(filter.OrganisationID, v.UniqueID)
				}
			}

		}

	}
	return err
}
