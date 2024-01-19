package services

import (
	"errors"
	"hrms-services/constants"
	"hrms-services/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// SaveOffboardingCheckListMaster : ""
func (s *Service) SaveOffboardingCheckListMaster(ctx *models.Context, offboardingchecklistmaster *models.OffboardingCheckListMaster) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	offboardingchecklistmaster.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONOFFBOARDINGCHECKLISTMASTER)
	offboardingchecklistmaster.Status = constants.OFFBOARDINGCHECKLISTMASTERSTATUSACTIVE
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 OffboardingCheckListMaster.created")
	offboardingchecklistmaster.Created = &created
	log.Println("b4 OffboardingCheckListMaster.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveOffboardingCheckListMaster(ctx, offboardingchecklistmaster)
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

// GetSingleOffboardingCheckListMaster : ""
func (s *Service) GetSingleOffboardingCheckListMaster(ctx *models.Context, UniqueID string) (*models.RefOffboardingCheckListMaster, error) {
	offboardingchecklistmaster, err := s.Daos.GetSingleOffboardingCheckListMaster(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return offboardingchecklistmaster, nil
}

//UpdateOffboardingCheckListMaster : ""
func (s *Service) UpdateOffboardingCheckListMaster(ctx *models.Context, offboardingchecklistmaster *models.OffboardingCheckListMaster) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateOffboardingCheckListMaster(ctx, offboardingchecklistmaster)
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

// EnableOffboardingCheckListMaster : ""
func (s *Service) EnableOffboardingCheckListMaster(ctx *models.Context, uniqueID string) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.EnableOffboardingCheckListMaster(ctx, uniqueID)
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

// DisableOffboardingCheckListMaster : ""
func (s *Service) DisableOffboardingCheckListMaster(ctx *models.Context, uniqueID string) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		debrr := s.Daos.DisableOffboardingCheckListMaster(ctx, uniqueID)
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

//DeleteOffboardingCheckListMaster : ""
func (s *Service) DeleteOffboardingCheckListMaster(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteOffboardingCheckListMaster(ctx, UniqueID)
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

// FilterOffboardingCheckListMaster : ""
func (s *Service) FilterOffboardingCheckListMaster(ctx *models.Context, offboardingchecklistmaster *models.FilterOffboardingCheckListMaster, pagination *models.Pagination) (offboardingchecklistmasters []models.RefOffboardingCheckListMaster, err error) {
	err = s.OffboardingCheckListMasterDataAccess(ctx, offboardingchecklistmaster)
	if err != nil {
		return nil, err
	}
	return s.Daos.FilterOffboardingCheckListMaster(ctx, offboardingchecklistmaster, pagination)
}
func (s *Service) OffboardingCheckListMasterDataAccess(ctx *models.Context, filter *models.FilterOffboardingCheckListMaster) (err error) {
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
