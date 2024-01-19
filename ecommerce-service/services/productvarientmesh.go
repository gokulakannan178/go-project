package services

import (
	"ecommerce-service/constants"
	"errors"
	"log"

	"ecommerce-service/models"

	"go.mongodb.org/mongo-driver/mongo"
)

// SaveProductVariantMesh : ""
func (s *Service) SaveProductVariantMesh(ctx *models.Context, ProductVariantMesh *models.ProductVariantMesh) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	ProductVariantMesh.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONPRODUCTVARIANT)
	ProductVariantMesh.Status = constants.PRODUCTVARIANTSTATUSACTIVE
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveProductVariantMesh(ctx, ProductVariantMesh)
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

//GetSingleProductVariantMesh :""
func (s *Service) GetSingleProductVariantMesh(ctx *models.Context, UniqueID string) (*models.RefProductVariantMesh, error) {
	tower, err := s.Daos.GetSingleProductVariantMesh(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return tower, nil
}

// UpdateProductVariantMesh : ""
func (s *Service) UpdateProductVariantMesh(ctx *models.Context, crop *models.ProductVariantMesh) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateProductVariantMesh(ctx, crop)
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

// EnableProductVariantMesh : ""
func (s *Service) EnableProductVariantMesh(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableProductVariantMesh(ctx, UniqueID)
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

//DisableProductVariantMesh : ""
func (s *Service) DisableProductVariantMesh(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableProductVariantMesh(ctx, UniqueID)
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

//DeleteProductVariantMesh : ""
func (s *Service) DeleteProductVariantMesh(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteProductVariantMesh(ctx, UniqueID)
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

// FilterProductVariantMesh : ""
func (s *Service) FilterProductVariantMesh(ctx *models.Context, filter *models.ProductVariantMeshFilter, pagination *models.Pagination) ([]models.RefProductVariantMesh, error) {
	return s.Daos.FilterProductVariantMesh(ctx, filter, pagination)

}
