package services

import (
	"errors"
	"log"
	"nicessm-api-service/constants"
	"nicessm-api-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//SaveFarmerSoilData :""
func (s *Service) SaveFarmerSoilData(ctx *models.Context, farmerSoilData *models.FarmerSoilData) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	// FarmerSoilData.Name = s.Daos.GetUniqueID(ctx, constants.COLLECTIONFarmerSoilData)
	farmerSoilData.Status = constants.FARMERSOILDATASTATUSACTIVE
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 FarmerSoilData.created")
	farmerSoilData.Created = &created
	log.Println("b4 FarmerSoilData.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveFarmerSoilData(ctx, farmerSoilData)
		if dberr != nil {

			return errors.New("Db Error" + dberr.Error())
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

//UpdateFarmerSoilData : ""
func (s *Service) UpdateFarmerSoilData(ctx *models.Context, farmerSoilData *models.FarmerSoilData) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateFarmerSoilData(ctx, farmerSoilData)
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

//EnableFarmerSoilData : ""
func (s *Service) EnableFarmerSoilData(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableFarmerSoilData(ctx, UniqueID)
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

//DisableFarmerSoilData : ""
func (s *Service) DisableFarmerSoilData(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableFarmerSoilData(ctx, UniqueID)
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

//DeleteFarmerSoilData : ""
func (s *Service) DeleteFarmerSoilData(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteFarmerSoilData(ctx, UniqueID)
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

//GetSingleFarmerSoilData :""
func (s *Service) GetSingleFarmerSoilData(ctx *models.Context, UniqueID string) (*models.RefFarmerSoilData, error) {
	farmerSoilData, err := s.Daos.GetSingleFarmerSoilData(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return farmerSoilData, nil
}

//FilterFarmerSoilData :""
func (s *Service) FilterFarmerSoilData(ctx *models.Context, farmerSoilDatafilter *models.FarmerSoilDataFilter, pagination *models.Pagination) (farmerSoilData []models.RefFarmerSoilData, err error) {
	return s.Daos.FilterFarmerSoilData(ctx, farmerSoilDatafilter, pagination)
}
