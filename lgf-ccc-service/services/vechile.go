package service

import (
	"errors"
	"fmt"

	"lgf-ccc-service/constants"
	"lgf-ccc-service/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//SaveVechile :""
func (s *Service) SaveVechile(ctx *models.Context, Vechile *models.Vechile) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	Vechile.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONVECHILE)
	Vechile.Status = constants.VECHILESTATUSACTIVE
	t := time.Now()
	Vechile.Date = &t
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 Vechile.created")
	Vechile.Created = created
	log.Println("b4 Vechile.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveVechile(ctx, Vechile)
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

//GetSingleVechile :""
func (s *Service) GetSingleVechile(ctx *models.Context, UniqueID string) (*models.RefVechile, error) {
	Vechile, err := s.Daos.GetSingleVechile(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return Vechile, nil
}

//UpdateVechile : ""
func (s *Service) UpdateVechile(ctx *models.Context, vehicle *models.Vechile) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateVechile(ctx, vehicle)
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

//EnableVechile : ""
func (s *Service) EnableVechile(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableVechile(ctx, UniqueID)
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

//DisableVechile : ""
func (s *Service) DisableVechile(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableVechile(ctx, UniqueID)
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

//DeleteVechile : ""
func (s *Service) DeleteVechile(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteVechile(ctx, UniqueID)
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

//FilterVechile :""
func (s *Service) FilterVechile(ctx *models.Context, VechileFilter *models.FilterVechile, pagination *models.Pagination) ([]models.RefVechile, error) {
	defer ctx.Session.EndSession(ctx.CTX)
	return s.Daos.FilterVechile(ctx, VechileFilter, pagination)

}

func (s *Service) VechileAssign(ctx *models.Context, assign *models.VechileAssign) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	t := time.Now()
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		refvehicle, err := s.Daos.GetSingleVechile(ctx, assign.VehicleID)
		if err != nil {
			return errors.New("error in getting the Vechilelog- " + err.Error())
		}
		driver, err := s.Daos.GetSingleDriver(ctx, refvehicle.Driver.ID)
		if err != nil {
			return errors.New("error in getting the Vechilelog- " + err.Error())
		}
		fmt.Println("driver=========", refvehicle.Driver)
		refdriver, err := s.Daos.GetSingleDriver(ctx, assign.DriverID)
		if err != nil {
			return errors.New("error in getting the Vechilelog- " + err.Error())
		}

		refvehicleType, err := s.Daos.GetSingleVehicleType(ctx, refvehicle.VechileTypeID)
		if err != nil {
			return errors.New("error in getting the Vechilelog- " + err.Error())
		}
		fmt.Println("==============>", refvehicleType)
		fmt.Println("vechile==============>", refvehicle)

		if refvehicle != nil {
			var VechileLog models.VehicleLog
			VechileLog.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONVEHICLELOG)
			VechileLog.Vehicle = refvehicle.Vechile
			if refvehicleType != nil {
				VechileLog.VechileType = &refvehicleType.VehicleType
			}
			VechileLog.Driver = refdriver.User
			VechileLog.StartDate = &t
			VechileLog.Status = constants.VECHILESTATUSASSIGN
			VechileLog.StartDate = &t
			err = s.Daos.SaveVehicleLog(ctx, &VechileLog)
			if err != nil {
				return err

			}
		}
		if driver != nil {
			refVechileLog, err := s.Daos.GetSingleVehicleLogWithDriverID(ctx, refvehicle.Driver.ID)
			if err != nil {
				return errors.New("error in getting the Vechilelog- " + err.Error())
			}
			fmt.Println("refVechileLog==============>", refVechileLog)

			if refVechileLog != nil {
				err = s.Daos.RevokeVehicleLog(ctx, refVechileLog.UniqueID)
				if err != nil {
					return errors.New("error in updating the Vechilelog" + err.Error())
				}
			}
			if refVechileLog == nil {
				return nil
			}
		}

		if err := ctx.Session.CommitTransaction(sc); err != nil {
			return errors.New("Not able to commit - " + err.Error())
		}
		dberr := s.Daos.VechileAssign(ctx, assign)
		if dberr != nil {
			return dberr
		}
		dberr1 := s.Daos.UpdateDriverWithVehicle(ctx, assign)
		if dberr1 != nil {
			return dberr1
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
