package services

import (
	"errors"
	"log"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)


// SaveLeaseRentShopSubCategory : ""
func (s *Service) SaveLeaseRentShopSubCategory(ctx *models.Context, shopsubcategory *models.LeaseRentShopSubCategory) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	shopsubcategory.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONLEASERENTSHOPSUBCATEGORY)
	shopsubcategory.Status = constants.LEASERENTSHOPSUBCATEGORYSTATUSACTIVE
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 shopsubcategory.created")
	shopsubcategory.Created = created
	log.Println("b4 shopsubcategory.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveLeaseRentShopSubCategory(ctx, shopsubcategory)
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

//GetSingleLeaseRentShopSubCategory :""
func (s *Service) GetSingleLeaseRentShopSubCategory(ctx *models.Context, UniqueID string) (*models.LeaseRentShopSubCategory, error) {
	shopsubcategory, err := s.Daos.GetSingleLeaseRentShopSubCategory(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return shopsubcategory, nil
}

// UpdateLeaseRentShopSubCategory : ""
func (s *Service) UpdateLeaseRentShopSubCategory(ctx *models.Context, shopsubcategory *models.LeaseRentShopSubCategory) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateLeaseRentShopSubCategory(ctx, shopsubcategory)
		if err != nil {
			return err
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

// EnableLeaseRentShopSubCategory : ""
func (s *Service) EnableLeaseRentShopSubCategory(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableLeaseRentShopSubCategory(ctx, UniqueID)
		if err != nil {
			return err
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

//DisableLeaseRentShopSubCategory : ""
func (s *Service) DisableLeaseRentShopSubCategory(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableLeaseRentShopSubCategory(ctx, UniqueID)
		if err != nil {
			return err
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

// DeleteLeaseRentShopSubCategory : ""
func (s *Service) DeleteLeaseRentShopSubCategory(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteLeaseRentShopSubCategory(ctx, UniqueID)
		if err != nil {
			return err
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

// FilterLeaseRentShopSubCategory : ""
func (s *Service) FilterLeaseRentShopSubCategory(ctx *models.Context, shopsubcategoryfilter *models.LeaseRentShopSubCategoryFilter, pagination *models.Pagination) ([]models.LeaseRentShopSubCategory, error) {
	return s.Daos.FilterLeaseRentShopSubCategory(ctx, shopsubcategoryfilter, pagination)

}
