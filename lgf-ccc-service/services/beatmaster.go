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

// SaveBeatMaster : ""
func (s *Service) SaveBeatMaster(ctx *models.Context, BeatMaster *models.BeatMaster) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	BeatMaster.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONBEATMASTER)
	BeatMaster.Status = constants.BEATMASTERSTATUSACTIVE
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 BeatMaster.created")
	BeatMaster.Created = &created
	log.Println("b4 BeatMaster.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveBeatMaster(ctx, BeatMaster)
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

// GetSingleBeatMaster : ""
func (s *Service) GetSingleBeatMaster(ctx *models.Context, UniqueID string) (*models.RefBeatMaster, error) {
	BeatMaster, err := s.Daos.GetSingleBeatMaster(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return BeatMaster, nil
}

// UpdateBeatMaster : ""
func (s *Service) UpdateBeatMaster(ctx *models.Context, BeatMaster *models.BeatMaster) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateBeatMaster(ctx, BeatMaster)
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

// EnableBeatMaster : ""
func (s *Service) EnableBeatMaster(ctx *models.Context, uniqueID string) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.EnableBeatMaster(ctx, uniqueID)
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

// DisableBeatMaster : ""
func (s *Service) DisableBeatMaster(ctx *models.Context, uniqueID string) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		debrr := s.Daos.DisableBeatMaster(ctx, uniqueID)
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

// DeleteBeatMaster : ""
func (s *Service) DeleteBeatMaster(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteBeatMaster(ctx, UniqueID)
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

// FilterBeatMaster : ""
func (s *Service) FilterBeatMaster(ctx *models.Context, BeatMaster *models.FilterBeatMaster, pagination *models.Pagination) (BeatMasters []models.RefBeatMaster, err error) {
	return s.Daos.FilterBeatMaster(ctx, BeatMaster, pagination)
}

func (s *Service) VehicleAssignForBeatMaster(ctx *models.Context, assign *models.VehicleAssignBeat) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	//	t := time.Now()
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
		beat := new(models.BeatMaster)
		beat.UniqueID = assign.UniqueID
		beat.Driver.Id = driver.UniqueID
		beat.Driver.Name = driver.Name
		beat.Vehicle.Id = refvehicle.UniqueID
		beat.Vehicle.Name = refvehicle.VechileName
		dberr := s.Daos.VehicleAssignForBeatMaster(ctx, beat)
		if dberr != nil {
			return dberr
		}
		// if refvehicle != nil {
		// 	var VechileLog models.VehicleLog
		// 	VechileLog.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONVEHICLELOG)
		// 	VechileLog.Vehicle = refvehicle.Vechile
		// 	if refvehicleType != nil {
		// 		VechileLog.VechileType = &refvehicleType.VehicleType
		// 	}
		// 	VechileLog.Driver = refdriver.DriverDetails
		// 	VechileLog.StartDate = &t
		// 	VechileLog.Status = constants.VECHILESTATUSASSIGN
		// 	VechileLog.StartDate = &t
		// 	err = s.Daos.SaveVehicleLog(ctx, &VechileLog)
		// 	if err != nil {
		// 		return err

		// 	}
		// }
		// if driver != nil {
		// 	refVechileLog, err := s.Daos.GetSingleVehicleLogWithDriverID(ctx, refvehicle.Driver.ID)
		// 	if err != nil {
		// 		return errors.New("error in getting the Vechilelog- " + err.Error())
		// 	}
		// 	fmt.Println("refVechileLog==============>", refVechileLog)

		// 	if refVechileLog != nil {
		// 		err = s.Daos.RevokeVehicleLog(ctx, refVechileLog.UniqueID)
		// 		if err != nil {
		// 			return errors.New("error in updating the Vechilelog" + err.Error())
		// 		}
		// 	}
		// 	if refVechileLog == nil {
		// 		return nil
		// 	}
		// }

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
