package services

import (
	"errors"
	"log"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// SaveTradeLicenseCategoryType : ""
func (s *Service) SaveTradeLicenseCategoryType(ctx *models.Context, categoryType *models.TradeLicenseCategoryType) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	categoryType.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONTRADELICENSECATEGORYTYPE)
	categoryType.Status = constants.TRADELICENSECATEGORYTYPESTATUSACTIVE
	t := time.Now()
	categoryType.Created.On = &t
	categoryType.Created.By = constants.SYSTEM
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveTradeLicenseCategoryType(ctx, categoryType)
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

//GetSingleTradeLicenseCategoryType :""
func (s *Service) GetSingleTradeLicenseCategoryType(ctx *models.Context, UniqueID string) (*models.RefTradeLicenseCategoryType, error) {
	tower, err := s.Daos.GetSingleTradeLicenseCategoryType(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return tower, nil
}

// UpdateTradeLicenseCategoryType : ""
func (s *Service) UpdateTradeLicenseCategoryType(ctx *models.Context, categoryType *models.TradeLicenseCategoryType) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	// t := time.Now()
	// categoryType.Created.On = &t
	// categoryType.Created.By = constants.SYSTEM
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateTradeLicenseCategoryType(ctx, categoryType)
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

// EnableTradeLicenseCategoryType : ""
func (s *Service) EnableTradeLicenseCategoryType(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableTradeLicenseCategoryType(ctx, UniqueID)
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

//DisableTradeLicenseCategoryType : ""
func (s *Service) DisableTradeLicenseCategoryType(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableTradeLicenseCategoryType(ctx, UniqueID)
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

//DeleteTradeLicenseCategoryType : ""
func (s *Service) DeleteTradeLicenseCategoryType(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteTradeLicenseCategoryType(ctx, UniqueID)
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

// FilterTradeLicenseCategoryType : ""
func (s *Service) FilterTradeLicenseCategoryType(ctx *models.Context, filter *models.TradeLicenseCategoryTypeFilter, pagination *models.Pagination) ([]models.RefTradeLicenseCategoryType, error) {
	return s.Daos.FilterTradeLicenseCategoryType(ctx, filter, pagination)

}
