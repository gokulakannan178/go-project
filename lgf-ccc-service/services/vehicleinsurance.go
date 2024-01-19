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
func (s *Service) SaveVehicleInsurance(ctx *models.Context, vehicleinsurance *models.VehicleInsurance) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	vehicleinsurance.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONVEHICLEINSURANCE)
	vehicleinsurance.Status = constants.VEHICLEINSURANCESTATUSACTIVE
	//user.Password = "#nature32" //Default Password
	t := time.Now()
	vehicleinsurance.Date = &t
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 user.created")
	vehicleinsurance.Created = created
	log.Println("b4 user.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveVehicleInsurance(ctx, vehicleinsurance)
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

//GetSingleVehicleInsurance :""
func (s *Service) GetSingleVehicleInsurance(ctx *models.Context, UniqueID string) (*models.RefVehicleInsurance, error) {
	vehicleinsurance, err := s.Daos.GetSingleVehicleInsurance(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return vehicleinsurance, nil
}

//UpdateVehicleInsurance : ""
func (s *Service) UpdateVehicleInsurance(ctx *models.Context, vehicleinsurance *models.VehicleInsurance) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateVehicleInsurance(ctx, vehicleinsurance)
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

//EnableVehicleInsurance : ""
func (s *Service) EnableVehicleInsurance(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableVehicleInsurance(ctx, UniqueID)
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

//DisableVehicleInsurance : ""
func (s *Service) DisableVehicleInsurance(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableVehicleInsurance(ctx, UniqueID)
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

//DeleteVehicleInsurance : ""
func (s *Service) DeleteVehicleInsurance(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteVehicleInsurance(ctx, UniqueID)
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

//VehicleInsuranceFilter:""
func (s *Service) VehicleInsuranceFilter(ctx *models.Context, filter *models.VehicleInsuranceFilter, pagination *models.Pagination) (user []models.VehicleInsurance, err error) {
	return s.Daos.VehicleInsuranceFilter(ctx, filter, pagination)

}
