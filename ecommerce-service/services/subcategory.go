package services

import (
	"errors"
	"log"
	"time"

	"ecommerce-service/constants"
	"ecommerce-service/models"

	"go.mongodb.org/mongo-driver/mongo"
)

// SaveSubCategory : ""
func (s *Service) SaveSubCategory(ctx *models.Context, subCategory *models.SubCategory) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	subCategory.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONSUBCATEGORY)
	subCategory.Status = constants.SUBCATEGORYSTATUSACTIVE
	t := time.Now()

	created := new(models.CreatedV2)
	created.On = &t
	created.By = constants.SYSTEM
	subCategory.Created = created
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveSubCategory(ctx, subCategory)
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

//GetSingleSubCategory :""
func (s *Service) GetSingleSubCategory(ctx *models.Context, UniqueID string) (*models.RefSubCategory, error) {
	subCategory, err := s.Daos.GetSingleSubCategory(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return subCategory, nil
}

// UpdateSubCategory : ""
func (s *Service) UpdateSubCategory(ctx *models.Context, subCategory *models.SubCategory) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateSubCategory(ctx, subCategory)
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

// EnableSubCategory : ""
func (s *Service) EnableSubCategory(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableSubCategory(ctx, UniqueID)
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

//DisableSubCategory : ""
func (s *Service) DisableSubCategory(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableSubCategory(ctx, UniqueID)
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

//DeleteSubCategory : ""
func (s *Service) DeleteSubCategory(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteSubCategory(ctx, UniqueID)
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

// FilterSubCategory : ""
func (s *Service) FilterSubCategory(ctx *models.Context, filter *models.SubCategoryFilter, pagination *models.Pagination) ([]models.RefSubCategory, error) {
	return s.Daos.FilterSubCategory(ctx, filter, pagination)

}
