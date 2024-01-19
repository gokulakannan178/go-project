package services

import (
	"ecommerce-service/constants"
	"errors"
	"log"

	"ecommerce-service/models"

	"go.mongodb.org/mongo-driver/mongo"
)

// SaveProductVariantType : ""
func (s *Service) SaveProductVariantType(ctx *models.Context, block *models.ProductVariantType) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	block.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONPRODUCTVARIANTTYPE)
	block.Status = constants.PRODUCTVARIANTTYPESTATUSACTIVE
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveProductVariantType(ctx, block)
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

//GetSingleProductVariantType :""
func (s *Service) GetSingleProductVariantType(ctx *models.Context, UniqueID string) (*models.RefProductVariantType, error) {
	tower, err := s.Daos.GetSingleProductVariantType(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return tower, nil
}

// UpdateProductVariantType : ""
func (s *Service) UpdateProductVariantType(ctx *models.Context, crop *models.ProductVariantType) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateProductVariantType(ctx, crop)
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

// EnableProductVariantType : ""
func (s *Service) EnableProductVariantType(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableProductVariantType(ctx, UniqueID)
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

//DisableProductVariantType : ""
func (s *Service) DisableProductVariantType(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableProductVariantType(ctx, UniqueID)
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

//DeleteProductVariantType : ""
func (s *Service) DeleteProductVariantType(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteProductVariantType(ctx, UniqueID)
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

// FilterProductVariantType : ""
func (s *Service) FilterProductVariantType(ctx *models.Context, filter *models.ProductVariantTypeFilter, pagination *models.Pagination) ([]models.RefProductVariantType, error) {
	return s.Daos.FilterProductVariantType(ctx, filter, pagination)

}
