package services

import (
	"errors"
	"log"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// SaveSolidWasteUserChargeCategory : ""
func (s *Service) SaveSolidWasteUserChargeCategory(ctx *models.Context, solidwasteuserchargecategory *models.SolidWasteUserChargeCategory) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	solidwasteuserchargecategory.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONSOLIDWASTEUSERCHARGECATEGORY)
	solidwasteuserchargecategory.Status = constants.SOLIDWASTEUSERCHARGERATESTATUSACTIVE
	t := time.Now()
	Created := new(models.CreatedV2)
	Created.On = &t
	// SolidWasteUserChargeCategory.Created.By = constants.SYSTEM
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveSolidWasteUserChargeCategory(ctx, solidwasteuserchargecategory)
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

//GetSingleSolidWasteUserChargeCategory :""
func (s *Service) GetSingleSolidWasteUserChargeCategory(ctx *models.Context, UniqueID string) (*models.RefSolidWasteUserChargeCategory, error) {
	tower, err := s.Daos.GetSingleSolidWasteUserChargeCategory(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return tower, nil
}

// UpdateSolidWasteUserChargeCategory : ""
func (s *Service) UpdateSolidWasteUserChargeCategory(ctx *models.Context, SolidWasteUserChargeCategory *models.SolidWasteUserChargeCategory) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateSolidWasteUserChargeCategory(ctx, SolidWasteUserChargeCategory)
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

// EnableSolidWasteUserChargeCategory : ""
func (s *Service) EnableSolidWasteUserChargeCategory(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableSolidWasteUserChargeCategory(ctx, UniqueID)
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

//DisableSolidWasteUserChargeCategory : ""
func (s *Service) DisableSolidWasteUserChargeCategory(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableSolidWasteUserChargeCategory(ctx, UniqueID)
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

//DeleteSolidWasteUserChargeCategory : ""
func (s *Service) DeleteSolidWasteUserChargeCategory(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteSolidWasteUserChargeCategory(ctx, UniqueID)
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

// FilterSolidWasteUserChargeCategory : ""
func (s *Service) FilterSolidWasteUserChargeCategory(ctx *models.Context, filter *models.SolidWasteUserChargeCategoryFilter, pagination *models.Pagination) ([]models.RefSolidWasteUserChargeCategory, error) {
	return s.Daos.FilterSolidWasteUserChargeCategory(ctx, filter, pagination)

}
