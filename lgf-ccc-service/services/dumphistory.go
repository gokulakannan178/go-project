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

//SaveDumpHistory :""
func (s *Service) SaveDumpHistory(ctx *models.Context, DumpHistory *models.DumpHistory) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	DumpHistory.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONDUMPHISTORY)
	DumpHistory.Status = constants.DUMPHISTORYSTATUSACTIVE
	t := time.Now()
	DumpHistory.Date = &t
	DumpHistory.Time = &t
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 DumpHistory.created")
	DumpHistory.Created = created
	log.Println("b4 DumpHistory.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		vehicle, err := s.Daos.GetSingleVechile(ctx, DumpHistory.VehicleID.Id)
		if err != nil {
			return err
		}
		if &vehicle.Driver != nil {
			driver, err := s.Daos.GetSingleDriver(ctx, vehicle.Driver.ID)
			if err != nil {
				return err
			}

			fmt.Println("property=======>", vehicle)
			DumpHistory.VehicleID.Name = vehicle.VechileName
			DumpHistory.Driver = vehicle.Driver
			if driver != nil {
				DumpHistory.MinUser = driver.Manager
			}
		}
		dberr := s.Daos.SaveDumpHistory(ctx, DumpHistory)
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

//GetSingleDumpHistory :""
func (s *Service) GetSingleDumpHistory(ctx *models.Context, UniqueID string) (*models.RefDumpHistory, error) {
	DumpHistory, err := s.Daos.GetSingleDumpHistory(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return DumpHistory, nil
}

// GetSingleDumpHistoryUsingEmpID : ""
// func (s *Service) GetSingleDumpHistoryUsingEmpID(ctx *models.Context, UniqueID string) (*models.RefDumpHistory, error) {
// 	DumpHistory, err := s.Daos.GetSingleDumpHistoryUsingEmpID(ctx, UniqueID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return DumpHistory, nil
// }

//UpdateDumpHistory : ""
func (s *Service) UpdateDumpHistory(ctx *models.Context, DumpHistory *models.DumpHistory) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		//var uniqueIds []string
		// for k, v := range DumpHistory.DumpHistoryPropertysId {
		// 	fmt.Println("DumpHistoryPropertysId===>", v.UniqueID)
		// 	if v.UniqueID == "" {
		// 		DumpHistory.DumpHistoryPropertysId[k].UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONDumpHistoryPROPERTYS)

		// 	}
		// 	uniqueIds = append(uniqueIds, v.UniqueID)

		// }

		err := s.Daos.UpdateDumpHistory(ctx, DumpHistory)
		if err := ctx.Session.CommitTransaction(sc); err != nil {
			return errors.New("Not able to commit - " + err.Error())
		}
		// err = s.Daos.DumpHistoryPropertysRemoveNotPresentValue(ctx, DumpHistory.UniqueID, uniqueIds)
		if err != nil {
			return err
		}
		// err = s.Daos.DumpHistoryPropertysUpsert(ctx, DumpHistory)
		// if err != nil {
		// 	return err
		// }
		return nil
	}); err != nil {
		if err = ctx.Session.AbortTransaction(ctx.CTX); err != nil {
			return errors.New("Transaction Aborted with error" + err.Error())
		}
		return errors.New("Transaction Aborted - " + err.Error())
	}
	return nil
}

//EnableDumpHistory : ""
func (s *Service) EnableDumpHistory(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableDumpHistory(ctx, UniqueID)
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

//DisableDumpHistory : ""
func (s *Service) DisableDumpHistory(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableDumpHistory(ctx, UniqueID)
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

//DeleteDumpHistory : ""
func (s *Service) DeleteDumpHistory(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteDumpHistory(ctx, UniqueID)
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

//FilterDumpHistory :""
func (s *Service) FilterDumpHistory(ctx *models.Context, DumpHistoryFilter *models.FilterDumpHistory, pagination *models.Pagination) ([]models.RefDumpHistory, error) {
	defer ctx.Session.EndSession(ctx.CTX)
	return s.Daos.FilterDumpHistory(ctx, DumpHistoryFilter, pagination)

}
func (s *Service) GetQuantityByManagerId(ctx *models.Context, DumpHistoryFilter *models.FilterDumpHistory) ([]models.GetQuantity, error) {
	defer ctx.Session.EndSession(ctx.CTX)
	return s.Daos.GetQuantityByManagerId(ctx, DumpHistoryFilter)

}

func (s *Service) DateWiseDumpHistory(ctx *models.Context, DumpHistoryFilter *models.DayWiseDumpHistory) ([]models.GetQuantity, error) {
	defer ctx.Session.EndSession(ctx.CTX)
	return s.Daos.DateWiseDumpHistory(ctx, DumpHistoryFilter)

}
