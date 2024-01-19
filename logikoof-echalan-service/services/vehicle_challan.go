package services

import (
	"errors"
	"fmt"
	"log"
	"logikoof-echalan-service/constants"
	"logikoof-echalan-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//SaveVehicleChallan :""
func (s *Service) SaveVehicleChallan(ctx *models.Context, vehicleChallan *models.VehicleChallan) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	vehicleChallan.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLVEHICLECHALLAN)
	vehicleChallan.Status = constants.VEHICLECHALLANSTATUSACTIVE
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 vehicleChallan.created")
	vehicleChallan.Created = created
	vehicleChallan.Payment.Status = constants.VEHICLECHALLANPAYMENTSTATUSPENDING
	log.Println("b4 vehicleChallan.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveVehicleChallan(ctx, vehicleChallan)
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
	offenceTypeName := "<Offence Type>"
	offenctType, err := s.Daos.GetSingleOffenceType(ctx, vehicleChallan.OffenceType)
	if err != nil {
		fmt.Println("Error in geting offece type - " + err.Error())
	}
	if offenctType != nil {
		offenceTypeName = offenctType.Name
	}
	layout := "January 2, 2006"
	msg := fmt.Sprintf("Mr./Mrs./Miss %v, your vehicle %v - %v have been charged with a penalty of Rs.%v for %v (%v) at %v on %v  \n From\n Delhi Traffic Police",
		vehicleChallan.Vehicle.OwnerName,
		vehicleChallan.Vehicle.RegNo,
		vehicleChallan.Vehicle.VehicleClass+"-"+vehicleChallan.Vehicle.Model,
		vehicleChallan.Pelalty,
		offenceTypeName,
		vehicleChallan.OffenceDetail,
		vehicleChallan.OffenceAt,
		vehicleChallan.OffenceDate.Format(layout),
	)
	s.SendSMS(vehicleChallan.Vehicle.Mobile, msg)
	// s.SendSMS("7299424027", msg)
	return nil
}

//UpdateVehicleChallan : ""
func (s *Service) UpdateVehicleChallan(ctx *models.Context, vehicleChallan *models.VehicleChallan) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateVehicleChallan(ctx, vehicleChallan)
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

//EnableVehicleChallan : ""
func (s *Service) EnableVehicleChallan(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableVehicleChallan(ctx, UniqueID)
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

//DisableVehicleChallan : ""
func (s *Service) DisableVehicleChallan(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableVehicleChallan(ctx, UniqueID)
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

//DeleteVehicleChallan : ""
func (s *Service) DeleteVehicleChallan(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteVehicleChallan(ctx, UniqueID)
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

//GetSingleVehicleChallan :""
func (s *Service) GetSingleVehicleChallan(ctx *models.Context, UniqueID string) (*models.RefVehicleChallan, error) {
	vehicleChallan, err := s.Daos.GetSingleVehicleChallan(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return vehicleChallan, nil
}

//FilterVehicleChallan :""
func (s *Service) FilterVehicleChallan(ctx *models.Context, vehicleChallanfilter *models.VehicleChallanFilter, pagination *models.Pagination) (vehicleChallan []models.RefVehicleChallan, err error) {
	return s.Daos.FilterVehicleChallan(ctx, vehicleChallanfilter, pagination)
}
