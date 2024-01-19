package services

import (
	"errors"
	"log"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// SaveMobileTowerRegistrationRateMaster: ""
func (s *Service) SaveMobileTowerRegistrationRateMaster(ctx *models.Context, mobile *models.MobileTowerRegistrationRateMaster) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	mobile.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONMOBILETOWERREGISTRATIONRATEMASTER)
	t := time.Now()
	mobile.Created.On = &t
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveMobileTowerRegistrationRateMaster(ctx, mobile)
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

//GetSingleMobileTowerRegistrationRateMaster :""
func (s *Service) GetSingleMobileTowerRegistrationRateMaster(ctx *models.Context, Name string) (*models.MobileTowerRegistrationRateMaster, error) {
	tower, err := s.Daos.GetSingleMobileTowerRegistrationRateMaster(ctx, Name)
	if err != nil {
		return nil, err
	}
	return tower, nil
}

// UpdateMobileTowerRegistrationRateMaster: ""
func (s *Service) UpdateMobileTowerRegistrationRateMaster(ctx *models.Context, mobile *models.MobileTowerRegistrationRateMaster) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateMobileTowerRegistrationRateMaster(ctx, mobile)
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

// EnableMobileTowerRegistrationRateMaster: ""
func (s *Service) EnableMobileTowerRegistrationRateMaster(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableMobileTowerRegistrationRateMaster(ctx, UniqueID)
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

//DisableMobileTowerRegistrationRateMaster : ""
func (s *Service) DisableMobileTowerRegistrationRateMaster(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableMobileTowerRegistrationRateMaster(ctx, UniqueID)
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

// DeleteMobileTowerRegistrationRateMaster : ""
func (s *Service) DeleteMobileTowerRegistrationRateMaster(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteMobileTowerRegistrationRateMaster(ctx, UniqueID)
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

// FilterMobileTowerRegistrationRateMaster : ""
func (s *Service) FilterMobileTowerRegistrationRateMaster(ctx *models.Context, filter *models.MobileTowerRegistrationRateMasterFilter, pagination *models.Pagination) ([]models.RefMobileTowerRegistrationRateMaster, error) {
	return s.Daos.FilterMobileTowerRegistrationRateMaster(ctx, filter, pagination)

}
