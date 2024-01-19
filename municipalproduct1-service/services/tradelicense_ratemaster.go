package services

import (
	"errors"
	"log"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// SaveTradeLicenseRateMaster : ""
func (s *Service) SaveTradeLicenseRateMaster(ctx *models.Context, rateMaster *models.TradeLicenseRateMaster) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	rateMaster.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONTRADELICENSERATEMASTER)
	rateMaster.Status = constants.TRADELICENSERATEMASTERSTATUSACTIVE
	t := time.Now()
	rateMaster.Created.On = &t
	rateMaster.Created.By = constants.SYSTEM
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveTradeLicenseRateMaster(ctx, rateMaster)
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

//GetSingleTradeLicenseRateMaster :""
func (s *Service) GetSingleTradeLicenseRateMaster(ctx *models.Context, UniqueID string) (*models.RefTradeLicenseRateMaster, error) {
	tower, err := s.Daos.GetSingleTradeLicenseRateMaster(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return tower, nil
}

// UpdateTradeLicenseRateMaster : ""
func (s *Service) UpdateTradeLicenseRateMaster(ctx *models.Context, rateMaster *models.TradeLicenseRateMaster) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	// t := time.Now()
	// rateMaster.Created.On = &t
	// rateMaster.Created.By = constants.SYSTEM
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateTradeLicenseRateMaster(ctx, rateMaster)
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

// EnableTradeLicenseRateMaster : ""
func (s *Service) EnableTradeLicenseRateMaster(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableTradeLicenseRateMaster(ctx, UniqueID)
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

//DisableTradeLicenseRateMaster : ""
func (s *Service) DisableTradeLicenseRateMaster(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableTradeLicenseRateMaster(ctx, UniqueID)
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

//DeleteTradeLicenseRateMaster : ""
func (s *Service) DeleteTradeLicenseRateMaster(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteTradeLicenseRateMaster(ctx, UniqueID)
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

// FilterTradeLicenseRateMaster : ""
func (s *Service) FilterTradeLicenseRateMaster(ctx *models.Context, filter *models.TradeLicenseRateMasterFilter, pagination *models.Pagination) ([]models.RefTradeLicenseRateMaster, error) {
	return s.Daos.FilterTradeLicenseRateMaster(ctx, filter, pagination)

}
