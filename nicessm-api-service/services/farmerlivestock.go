package services

import (
	"errors"
	"log"
	"nicessm-api-service/constants"
	"nicessm-api-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//SaveFarmerLiveStock :""
func (s *Service) SaveFarmerLiveStock(ctx *models.Context, farmerLiveStock *models.FarmerLiveStock) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	// FarmerLiveStock.Name = s.Daos.GetUniqueID(ctx, constants.COLLECTIONFarmerLiveStock)
	farmerLiveStock.Status = constants.FARMERLIVESTOCKSTATUSACTIVE
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 FarmerLiveStock.created")
	farmerLiveStock.Created = &created
	log.Println("b4 FarmerLiveStock.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveFarmerLiveStock(ctx, farmerLiveStock)
		if dberr != nil {

			return errors.New("Db Error" + dberr.Error())
		}
		err := s.Daos.GetFarmerLiveStockCount(ctx, farmerLiveStock.Farmer.Hex())
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

//UpdateFarmerLiveStock : ""
func (s *Service) UpdateFarmerLiveStock(ctx *models.Context, farmerLiveStock *models.FarmerLiveStock) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateFarmerLiveStock(ctx, farmerLiveStock)
		if err != nil {

			return errors.New("Db Error" + err.Error())
		}
		farmer, err := s.Daos.GetSingleFarmerLiveStock(ctx, farmerLiveStock.ID.Hex())
		if err != nil {
			return err
		}
		err = s.Daos.GetFarmerLiveStockCount(ctx, farmer.Farmer.Hex())
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

//EnableFarmerLiveStock : ""
func (s *Service) EnableFarmerLiveStock(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableFarmerLiveStock(ctx, UniqueID)
		if err != nil {
			return err
		}
		farmer, err := s.Daos.GetSingleFarmerLiveStock(ctx, UniqueID)
		if err != nil {
			return err
		}
		err = s.Daos.GetFarmerLiveStockCount(ctx, farmer.Farmer.Hex())
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

//DisableFarmerLiveStock : ""
func (s *Service) DisableFarmerLiveStock(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableFarmerLiveStock(ctx, UniqueID)
		if err != nil {
			return err
		}
		farmer, err := s.Daos.GetSingleFarmerLiveStock(ctx, UniqueID)
		if err != nil {
			return err
		}
		err = s.Daos.GetFarmerLiveStockCount(ctx, farmer.Farmer.Hex())
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

//DeleteFarmerLiveStock : ""
func (s *Service) DeleteFarmerLiveStock(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteFarmerLiveStock(ctx, UniqueID)
		if err != nil {
			return err
		}
		farmer, err := s.Daos.GetSingleFarmerLiveStock(ctx, UniqueID)
		if err != nil {
			return err
		}
		err = s.Daos.GetFarmerLiveStockCount(ctx, farmer.Farmer.Hex())
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

//GetSingleFarmerLiveStock :""
func (s *Service) GetSingleFarmerLiveStock(ctx *models.Context, UniqueID string) (*models.RefFarmerLiveStock, error) {
	farmerLiveStock, err := s.Daos.GetSingleFarmerLiveStock(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return farmerLiveStock, nil
}

//FilterFarmerLiveStock :""
func (s *Service) FilterFarmerLiveStock(ctx *models.Context, farmerLiveStockfilter *models.FarmerLiveStockFilter, pagination *models.Pagination) (farmerLiveStock []models.RefFarmerLiveStock, err error) {
	return s.Daos.FilterFarmerLiveStock(ctx, farmerLiveStockfilter, pagination)
}
