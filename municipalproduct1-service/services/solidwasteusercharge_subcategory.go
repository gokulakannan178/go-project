package services

import (
	"errors"
	"log"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// SaveSolidWasteUserChargeSubCategory : ""
func (s *Service) SaveSolidWasteUserChargeSubCategory(ctx *models.Context, solidwasteuserchargesubcategory *models.SolidWasteUserChargeSubCategory) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	solidwasteuserchargesubcategory.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONSOLIDWASTEUSERCHARGESUBCATEGORY)
	solidwasteuserchargesubcategory.Status = constants.SOLIDWASTEUSERCHARGESUBCATEGORYSTATUSACTIVE
	t := time.Now()
	Created := new(models.CreatedV2)
	Created.On = &t
	// SolidWasteUserChargeSubCategory.Created.By = constants.SYSTEM
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveSolidWasteUserChargeSubCategory(ctx, solidwasteuserchargesubcategory)
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

//GetSingleSolidWasteUserChargeSubCategory :""
func (s *Service) GetSingleSolidWasteUserChargeSubCategory(ctx *models.Context, UniqueID string) (*models.RefSolidWasteUserChargeSubCategory, error) {
	tower, err := s.Daos.GetSingleSolidWasteUserChargeSubCategory(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return tower, nil
}

// UpdateSolidWasteUserChargeSubCategory : ""
func (s *Service) UpdateSolidWasteUserChargeSubCategory(ctx *models.Context, solidwasteuserchargesubcategory *models.SolidWasteUserChargeSubCategory) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateSolidWasteUserChargeSubCategory(ctx, solidwasteuserchargesubcategory)
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

// EnableSolidWasteUserChargeSubCategory : ""
func (s *Service) EnableSolidWasteUserChargeSubCategory(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableSolidWasteUserChargeSubCategory(ctx, UniqueID)
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

//DisableSolidWasteUserChargeSubCategory : ""
func (s *Service) DisableSolidWasteUserChargeSubCategory(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableSolidWasteUserChargeSubCategory(ctx, UniqueID)
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

//DeleteSolidWasteUserChargeSubCategory : ""
func (s *Service) DeleteSolidWasteUserChargeSubCategory(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteSolidWasteUserChargeSubCategory(ctx, UniqueID)
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

// FilterSolidWasteUserChargeSubCategory : ""
func (s *Service) FilterSolidWasteUserChargeSubCategory(ctx *models.Context, filter *models.SolidWasteUserChargeSubCategoryFilter, pagination *models.Pagination) ([]models.RefSolidWasteUserChargeSubCategory, error) {
	return s.Daos.FilterSolidWasteUserChargeSubCategory(ctx, filter, pagination)

}
