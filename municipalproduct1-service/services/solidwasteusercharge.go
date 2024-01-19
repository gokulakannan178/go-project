package services

import (
	"errors"
	"log"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"

	"go.mongodb.org/mongo-driver/mongo"
)

// SaveSolidWasteUserCharge : ""
func (s *Service) SaveSolidWasteUserCharge(ctx *models.Context, solidwasteusercharge *models.SolidWasteUserCharge) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	solidwasteusercharge.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONSOLIDWASTEUSERCHARGE)
	solidwasteusercharge.Status = constants.SOLIDWASTEUSERCHARGESTATUSACTIVE
	//t := time.Now()
	// Created = new(models.CreatedV2)
	// Created.On = &t
	// SolidWasteUserCharge.Created.By = constants.SYSTEM
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveSolidWasteUserCharge(ctx, solidwasteusercharge)
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

//GetSingleSolidWasteUserCharge :""
func (s *Service) GetSingleSolidWasteUserCharge(ctx *models.Context, UniqueID string) (*models.RefSolidWasteUserCharge, error) {
	tower, err := s.Daos.GetSingleSolidWasteUserCharge(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return tower, nil
}

// UpdateSolidWasteUserCharge : ""
func (s *Service) UpdateSolidWasteUserCharge(ctx *models.Context, solidwasteusercharge *models.SolidWasteUserCharge) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateSolidWasteUserCharge(ctx, solidwasteusercharge)
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

// EnableSolidWasteUserCharge : ""
func (s *Service) EnableSolidWasteUserCharge(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableSolidWasteUserCharge(ctx, UniqueID)
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

//DisableSolidWasteUserCharge : ""
func (s *Service) DisableSolidWasteUserCharge(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableSolidWasteUserCharge(ctx, UniqueID)
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

//DeleteSolidWasteUserCharge : ""
func (s *Service) DeleteSolidWasteUserCharge(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteSolidWasteUserCharge(ctx, UniqueID)
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

// FilterSolidWasteUserCharge : ""
func (s *Service) FilterSolidWasteUserCharge(ctx *models.Context, filter *models.SolidWasteUserChargeFilter, pagination *models.Pagination) ([]models.RefSolidWasteUserCharge, error) {
	return s.Daos.FilterSolidWasteUserCharge(ctx, filter, pagination)

}
