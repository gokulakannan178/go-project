package service

import (
	"errors"
	"lgf-ccc-service/constants"
	"lgf-ccc-service/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// SaveVehicleLocation : ""
func (s *Service) SaveVehicleLocation(ctx *models.Context, vehiclelocation *models.VehicleLocation) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	vehiclelocation.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONVEHICLELOCATION)
	vehiclelocation.Status = constants.VEHICLELOCATIONSTATUSACTIVE
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 VehicleLocation.created")
	vehiclelocation.Created = &created
	log.Println("b4 VehicleLocation.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		vehicle := new(models.VehicleLocationUpdate)
		vehicle.VehicleID = vehiclelocation.VehicleId
		vehicle.Location = vehiclelocation.Location
		dberr := s.Daos.UpdateVechileLocation(ctx, vehicle)
		if dberr != nil {
			return dberr
		}

		dberr = s.Daos.SaveVehicleLocation(ctx, vehiclelocation)
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

// GetSingleVehicleLocation : ""
func (s *Service) GetSingleVehicleLocation(ctx *models.Context, UniqueID string) (*models.RefVehicleLocation, error) {
	vehiclelocation, err := s.Daos.GetSingleVehicleLocation(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return vehiclelocation, nil
}

func (s *Service) UpdateVehicleLocation(ctx *models.Context, vehiclelocation *models.VehicleLocation) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateVehicleLocation(ctx, vehiclelocation)
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

// EnableVehicleLocation : ""
func (s *Service) EnableVehicleLocation(ctx *models.Context, uniqueID string) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.EnableVehicleLocation(ctx, uniqueID)
		if dberr != nil {
			return dberr
		}
		if err := sc.CommitTransaction(sc); err != nil {
			return errors.New("Not able to commit - " + err.Error())
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

// DisableVehicleLocation : ""
func (s *Service) DisableVehicleLocation(ctx *models.Context, uniqueID string) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		debrr := s.Daos.DisableVehicleLocation(ctx, uniqueID)
		if debrr != nil {
			return debrr
		}
		if err := sc.CommitTransaction(sc); err != nil {
			return errors.New("Not able to commit - " + err.Error())
		}
		return nil
	}); err != nil {
		if err1 := ctx.Session.AbortTransaction(ctx.CTX); err1 != nil {
			log.Println("err in abort")
			return errors.New("Transaction Abort with error" + err1.Error())
		}
		return err
	}
	return nil
}

//DeleteState : ""
func (s *Service) DeleteVehicleLocation(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteVehicleLocation(ctx, UniqueID)
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

// FilterVehicleLocation : ""
func (s *Service) FilterVehicleLocation(ctx *models.Context, Filter *models.FilterVehicleLocation, pagination *models.Pagination) (Dept []models.RefVehicleLocation, err error) {
	return s.Daos.FilterVehicleLocation(ctx, Filter, pagination)
}
