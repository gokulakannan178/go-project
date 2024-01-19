package services

import (
	"errors"
	"log"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// SaveMobileTowerRegistrationTax : ""
func (s *Service) SaveMobileTowerRegistrationTax(ctx *models.Context, mtrt *models.MobileTowerRegistrationTax) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	mtrt.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONMOBILETOWERREGISTRATIONTAX)
	mtrt.Status = constants.MOBILETOWERREGISTRATIONTAXSTATUSACTIVE
	t := time.Now()
	created := new(models.CreatedV2)
	mtrt.Created.On = &t
	mtrt.Created.By = constants.SYSTEM
	mtrt.Created = *created
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveMobileTowerRegistrationTax(ctx, mtrt)
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

//GetSingleMobileTowerRegistrationTax :""
func (s *Service) GetSingleMobileTowerRegistrationTax(ctx *models.Context, UniqueID string) (*models.RefMobileTowerRegistrationTax, error) {
	tower, err := s.Daos.GetSingleMobileTowerRegistrationTax(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return tower, nil
}

//GetSingleDefaultMobileTowerregistration :""
func (s *Service) GetSingleDefaultMobileTowerRegistrationTax(ctx *models.Context) (*models.RefMobileTowerRegistrationTax, error) {
	tower, err := s.Daos.GetSingleDefaultMobileTowerRegistrationTax(ctx)
	if err != nil {
		return nil, err
	}
	return tower, nil
}

// UpdateMobileTowerRegistrationTax : ""
func (s *Service) UpdateMobileTowerRegistrationTax(ctx *models.Context, mtrt *models.MobileTowerRegistrationTax) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateMobileTowerRegistrationTax(ctx, mtrt)
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

// EnableMobileTowerRegistrationTax : ""
func (s *Service) EnableMobileTowerRegistrationTax(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableMobileTowerRegistrationTax(ctx, UniqueID)
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

//DisableMobileTowerRegistrationTax : ""
func (s *Service) DisableMobileTowerRegistrationTax(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableMobileTowerRegistrationTax(ctx, UniqueID)
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

//DeleteMobileTowerRegistrationTax : ""
func (s *Service) DeleteMobileTowerRegistrationTax(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteMobileTowerRegistrationTax(ctx, UniqueID)
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

// FilterMobileTowerRegistrationTax : ""
func (s *Service) FilterMobileTowerRegistrationTax(ctx *models.Context, filter *models.MobileTowerRegistrationTaxFilter, pagination *models.Pagination) ([]models.RefMobileTowerRegistrationTax, error) {
	return s.Daos.FilterMobileTowerRegistrationTax(ctx, filter, pagination)

}
