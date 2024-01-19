package service

import (
	"errors"
	"lgf-ccc-service/constants"
	"lgf-ccc-service/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//SaveUser :""
func (s *Service) SaveFuelHistory(ctx *models.Context, fuelhistory *models.FuelHistory) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	fuelhistory.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONFUELHISTORY)
	fuelhistory.Status = constants.FUELHISTORYSTATUSACTIVE
	//user.Password = "#nature32" //Default Password
	t := time.Now()
	fuelhistory.Date = &t
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 user.created")
	fuelhistory.Created = created
	log.Println("b4 user.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveFuelHistory(ctx, fuelhistory)
		if dberr != nil {

			log.Println("err in abort out")
			return errors.New("Transaction Aborted - " + dberr.Error())
		}
		return nil

	}); err != nil {
		if err1 := ctx.Session.AbortTransaction(ctx.CTX); err1 != nil {
			log.Println("err in abort")
			return errors.New("Transaction Aborted with error" + err1.Error())
		}
		return err
	}

	return nil
}

//GetSingleFuelHistory :""
func (s *Service) GetSingleFuelHistory(ctx *models.Context, UniqueID string) (*models.RefFuelHistory, error) {
	fuelhistory, err := s.Daos.GetSingleFuelHistory(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return fuelhistory, nil
}

//UpdateFuelHistory : ""
func (s *Service) UpdateFuelHistory(ctx *models.Context, fuelhistory *models.FuelHistory) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateFuelHistory(ctx, fuelhistory)
		if err != nil {
			log.Println("err in abort out")
			return errors.New("Transaction Aborted - " + err.Error())
		}
		return nil

	}); err != nil {
		if err = ctx.Session.AbortTransaction(ctx.CTX); err != nil {
			return errors.New("Transaction Aborted with error" + err.Error())

			//return errors.New("Transaction Aborted - " + err.Error())
		}
		return err
	}
	return nil
}

//EnableFuelHistory : ""
func (s *Service) EnableFuelHistory(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableFuelHistory(ctx, UniqueID)
		if err != nil {

			return errors.New("Transaction Aborted - " + err.Error())
		}
		return nil

	}); err != nil {
		if err = ctx.Session.AbortTransaction(ctx.CTX); err != nil {
			return errors.New("Transaction Aborted with error" + err.Error())
		}

		return err
	}
	return nil
}

//DisableFuelHistory : ""
func (s *Service) DisableFuelHistory(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableFuelHistory(ctx, UniqueID)
		if err != nil {

			return errors.New("Transaction Aborted - " + err.Error())
		}
		return nil

	}); err != nil {
		if err = ctx.Session.AbortTransaction(ctx.CTX); err != nil {
			return errors.New("Transaction Aborted with error" + err.Error())
		}
		return err
	}
	return nil
}

//DeleteFuelHistory : ""
func (s *Service) DeleteFuelHistory(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteFuelHistory(ctx, UniqueID)
		if err != nil {

			return errors.New("Transaction Aborted - " + err.Error())
		}
		return nil

	}); err != nil {
		if err = ctx.Session.AbortTransaction(ctx.CTX); err != nil {
			return errors.New("Transaction Aborted with error" + err.Error())
		}
		return err
	}
	return nil
}

//FuelHistoryFilter:""
func (s *Service) FuelHistoryFilter(ctx *models.Context, filter *models.FuelHistoryFilter, pagination *models.Pagination) (user []models.FuelHistory, err error) {
	return s.Daos.FuelHistoryFilter(ctx, filter, pagination)

}
