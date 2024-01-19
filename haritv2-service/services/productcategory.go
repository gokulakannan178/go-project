package services

import (
	"errors"
	"haritv2-service/constants"
	"haritv2-service/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//SaveProductCategory :""
func (s *Service) SaveProductCategory(ctx *models.Context, product *models.ProductCategory) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	product.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONPRODUCTCATEGORY)
	product.Status = constants.PRODUCTCATEGORYSTATUSACTIVE
	t := time.Now()
	created := models.CreatedV2{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 ProductCategory.created")
	product.Created = &created
	log.Println("b4 ProductCategory.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveProductCategory(ctx, product)
		if dberr != nil {
			if err1 := ctx.Session.AbortTransaction(sc); err1 != nil {
				log.Println("err in abort")
				return errors.New("Transaction Aborted with error" + err1.Error())
			}
			log.Println("err in abort out")
			return errors.New("Transaction Aborted - " + dberr.Error())
		}
		return nil

	}); err != nil {
		return err
	}

	return nil
}

//UpdateProductCategory : ""
func (s *Service) UpdateProductCategory(ctx *models.Context, product *models.ProductCategory) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateProductCategory(ctx, product)
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

//EnableProductCategory : ""
func (s *Service) EnableProductCategory(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableProductCategory(ctx, UniqueID)
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

//DisableProductCategory : ""
func (s *Service) DisableProductCategory(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableProductCategory(ctx, UniqueID)
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

//DeleteProductCategory : ""
func (s *Service) DeleteProductCategory(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteProductCategory(ctx, UniqueID)
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

//GetSingleProductCategory :""
func (s *Service) GetSingleProductCategory(ctx *models.Context, UniqueID string) (*models.RefProductCategory, error) {
	ProductCategory, err := s.Daos.GetSingleProductCategory(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return ProductCategory, nil
}

//FilterProductCategory :""
func (s *Service) FilterProductCategory(ctx *models.Context, filter *models.ProductCategoryFilter, pagination *models.Pagination) ([]models.RefProductCategory, error) {
	defer ctx.Session.EndSession(ctx.CTX)
	return s.Daos.FilterProductCategory(ctx, filter, pagination)

}

//GetDefaultProductCategory :""
func (s *Service) GetDefaultProductCategory(ctx *models.Context) (*models.RefProductCategory, error) {
	product, err := s.Daos.GetDefaultProductCategory(ctx)
	if err != nil {
		return nil, err
	}
	return product, nil
}
