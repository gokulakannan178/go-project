package services

import (
	"errors"
	"log"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//SaveShopRentRateMaster :""
func (s *Service) SaveShopRentRateMaster(ctx *models.Context, ratemaster *models.ShopRentRateMaster) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	ratemaster.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONSHOPRENTRATEMASTER)
	ratemaster.Status = constants.SHOPRENTRATEMASTERSTATUSACTIVE
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 ratemaster.created")
	ratemaster.Created = &created
	log.Println("b4 ratemaster.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveShopRentRateMaster(ctx, ratemaster)
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

//GetSingleShopRentRateMaster :""
func (s *Service) GetSingleShopRentRateMaster(ctx *models.Context, UniqueID string) (*models.RefShopRentRateMaster, error) {
	ratemaster, err := s.Daos.GetSingleShopRentRateMaster(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return ratemaster, nil
}

//UpdateShopRentRateMaster : ""
func (s *Service) UpdateShopRentRateMaster(ctx *models.Context, ratemaster *models.ShopRentRateMaster) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateShopRentRateMaster(ctx, ratemaster)
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

//EnableShopRentRateMaster : ""
func (s *Service) EnableShopRentRateMaster(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableShopRentRateMaster(ctx, UniqueID)
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

//DisableShopRentRateMaster : ""
func (s *Service) DisableShopRentRateMaster(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableShopRentRateMaster(ctx, UniqueID)
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

//DeleteShopRentRateMaster : ""
func (s *Service) DeleteShopRentRateMaster(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteShopRentRateMaster(ctx, UniqueID)
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

//FilterShopRentRateMaster :""
func (s *Service) FilterShopRentRateMaster(ctx *models.Context, ratemasterfilter *models.ShopRentRateMasterFilter, pagination *models.Pagination) (ratemaster []models.RefShopRentRateMaster, err error) {
	defer ctx.Session.EndSession(ctx.CTX)
	return s.Daos.FilterShopRentRateMaster(ctx, ratemasterfilter, pagination)
}
