package services

import (
	"errors"
	"log"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"

	"go.mongodb.org/mongo-driver/mongo"
)

// SaveSolidWasteUserChargeRate : ""
func (s *Service) SaveSolidWasteUserChargeRate(ctx *models.Context, solidwasteuserchargerate *models.SolidWasteUserChargeRate) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	solidwasteuserchargerate.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONSOLIDWASTEUSERCHARGERATE)
	solidwasteuserchargerate.Status = constants.SOLIDWASTEUSERCHARGERATESTATUSACTIVE
	//t := time.Now()
	// Created = new(models.CreatedV2)
	// Created.On = &t
	// SolidWasteUserChargeRate.Created.By = constants.SYSTEM
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveSolidWasteUserChargeRate(ctx, solidwasteuserchargerate)
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

//GetSingleSolidWasteUserChargeRate :""
func (s *Service) GetSingleSolidWasteUserChargeRate(ctx *models.Context, UniqueID string) (*models.RefSolidWasteUserChargeRate, error) {
	tower, err := s.Daos.GetSingleSolidWasteUserChargeRate(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return tower, nil
}

// UpdateSolidWasteUserChargeRate : ""
func (s *Service) UpdateSolidWasteUserChargeRate(ctx *models.Context, solidwasteuserchargerate *models.SolidWasteUserChargeRate) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateSolidWasteUserChargeRate(ctx, solidwasteuserchargerate)
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

// EnableSolidWasteUserChargeRate : ""
func (s *Service) EnableSolidWasteUserChargeRate(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableSolidWasteUserChargeRate(ctx, UniqueID)
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

//DisableSolidWasteUserChargeRate : ""
func (s *Service) DisableSolidWasteUserChargeRate(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableSolidWasteUserChargeRate(ctx, UniqueID)
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

//DeleteSolidWasteUserChargeRate : ""
func (s *Service) DeleteSolidWasteUserChargeRate(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteSolidWasteUserChargeRate(ctx, UniqueID)
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

// FilterSolidWasteUserChargeRate : ""
func (s *Service) FilterSolidWasteUserChargeRate(ctx *models.Context, filter *models.SolidWasteUserChargeRateFilter, pagination *models.Pagination) ([]models.RefSolidWasteUserChargeRate, error) {
	return s.Daos.FilterSolidWasteUserChargeRate(ctx, filter, pagination)

}
