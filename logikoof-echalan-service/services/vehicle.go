package services

import (
	"errors"
	"log"
	"logikoof-echalan-service/constants"
	"logikoof-echalan-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//SaveVehicle :""
func (s *Service) SaveVehicle(ctx *models.Context, vehicle *models.Vehicle) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	// vehicle.RegNo = s.Daos.GetUniqueID(ctx, constants.COLLVEHICLE)
	vehicle.Status = constants.VEHICLESTATUSACTIVE
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 vehicle.created")
	vehicle.Created = created
	log.Println("b4 vehicle.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveVehicle(ctx, vehicle)
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

//UpdateVehicle : ""
func (s *Service) UpdateVehicle(ctx *models.Context, vehicle *models.Vehicle) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateVehicle(ctx, vehicle)
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

//EnableVehicle : ""
func (s *Service) EnableVehicle(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableVehicle(ctx, UniqueID)
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

//DisableVehicle : ""
func (s *Service) DisableVehicle(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableVehicle(ctx, UniqueID)
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

//DeleteVehicle : ""
func (s *Service) DeleteVehicle(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteVehicle(ctx, UniqueID)
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

//GetSingleVehicle :""
func (s *Service) GetSingleVehicle(ctx *models.Context, UniqueID string) (*models.RefVehicle, error) {
	vehicle, err := s.Daos.GetSingleVehicle(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return vehicle, nil
}

//FilterVehicle :""
func (s *Service) FilterVehicle(ctx *models.Context, vehiclefilter *models.VehicleFilter, pagination *models.Pagination) (vehicle []models.RefVehicle, err error) {
	return s.Daos.FilterVehicle(ctx, vehiclefilter, pagination)
}
