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
func (s *Service) SaveVehicleLog(ctx *models.Context, vehiclelog *models.VehicleLog) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	vehiclelog.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONVEHICLELOG)
	vehiclelog.Status = constants.VEHICLELOGSTATUSACTIVE
	//user.Password = "#nature32" //Default Password
	t := time.Now()
	vehiclelog.StartDate = &t
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 user.created")
	vehiclelog.Created = created
	log.Println("b4 user.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveVehicleLog(ctx, vehiclelog)
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

//GetSingleVehicleLog :""
func (s *Service) GetSingleVehicleLog(ctx *models.Context, UniqueID string) (*models.RefVehicleLog, error) {
	user, err := s.Daos.GetSingleVehicleLog(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

//UpdateVehicleLog : ""
func (s *Service) UpdateVehicleLog(ctx *models.Context, vehiclelog *models.VehicleLog) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateVehicleLog(ctx, vehiclelog)
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

//EnableVehicleLog : ""
func (s *Service) EnableVehicleLog(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableVehicleLog(ctx, UniqueID)
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

//DisableVehicleLog : ""
func (s *Service) DisableVehicleLog(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableVehicleLog(ctx, UniqueID)
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

//DeleteVehicleLog : ""
func (s *Service) DeleteVehicleLog(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteVehicleLog(ctx, UniqueID)
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

//VehicleLogFilter:""
func (s *Service) VehicleLogFilter(ctx *models.Context, filter *models.VehicleLogFilter, pagination *models.Pagination) (user []models.VehicleLog, err error) {
	return s.Daos.VehicleLogFilter(ctx, filter, pagination)

}
