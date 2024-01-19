package services

import (
	"ecommerce-service/constants"
	"errors"
	"log"

	"ecommerce-service/models"

	"go.mongodb.org/mongo-driver/mongo"
)

// SaveProductVariant : ""
func (s *Service) SaveProductVariant(ctx *models.Context, block *models.ProductVariant) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	//block.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONPRODUCTVARIANT)
	block.Status = constants.PRODUCTVARIANTSTATUSACTIVE
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveProductVariant(ctx, block)
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

//GetSingleProductVariant :""
func (s *Service) GetSingleProductVariant(ctx *models.Context, UniqueID string) (*models.RefProductVariant, error) {
	tower, err := s.Daos.GetSingleProductVariant(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return tower, nil
}

// UpdateProductVariant : ""
func (s *Service) UpdateProductVariant(ctx *models.Context, crop *models.ProductVariant) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateProductVariant(ctx, crop)
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

// EnableProductVariant : ""
func (s *Service) EnableProductVariant(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableProductVariant(ctx, UniqueID)
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

//DisableProductVariant : ""
func (s *Service) DisableProductVariant(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableProductVariant(ctx, UniqueID)
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

//DeleteProductVariant : ""
func (s *Service) DeleteProductVariant(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteProductVariant(ctx, UniqueID)
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

// FilterProductVariant : ""
func (s *Service) FilterProductVariant(ctx *models.Context, filter *models.ProductVariantFilter, pagination *models.Pagination) ([]models.RefProductVariant, error) {
	return s.Daos.FilterProductVariant(ctx, filter, pagination)

}
func (s *Service) ProductVariantRegister(ctx *models.Context, product *models.RegProductVariant) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	//block.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONPRODUCTVARIANT)
	product.Status = constants.PRODUCTVARIANTSTATUSACTIVE
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.ProductVariantRegisterWithoutTransaction(ctx, product)
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
func (s *Service) ProductVariantRegisterWithoutTransaction(ctx *models.Context, product *models.RegProductVariant) error {
	log.Println("transaction start")
	dispalyName := ""
	product.Status = constants.PRODUCTSTATUSACTIVE
	if len(product.Mesh) > 0 {
		for k, v := range product.Mesh {
			product.Mesh[k].ProductVariantID = product.UniqueID
			product.Mesh[k].ProductID = product.ProductID
			product.Mesh[k].UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONPRODUCTVARIANTMESH)
			product.Mesh[k].Status = constants.PRODUCTSTATUSACTIVE
			dispalyName = dispalyName + v.VariantTypeName + ":" + v.Value
			if k < len(product.Mesh) {
				dispalyName = dispalyName + "|"
			}
		}
		dberr := s.Daos.ProductVariantMeshRegister(ctx, product)

		if dberr != nil {

			return errors.New("Db Error" + dberr.Error())
		}
	}
	product.DispalyName = dispalyName
	dberr := s.Daos.ProductVariantRegister(ctx, product)
	if dberr != nil {

		return errors.New("Db Error" + dberr.Error())
	}

	return nil
}

// FilterProductVariant : ""
func (s *Service) GetMyInventory(ctx *models.Context, filter *models.ProductVariantInventoryFilter, pagination *models.Pagination) ([]models.RefProductVariant, error) {
	return s.Daos.GetMyInventory(ctx, filter, pagination)

}
