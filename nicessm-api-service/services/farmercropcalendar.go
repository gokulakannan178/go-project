package services

import (
	"errors"
	"log"
	"nicessm-api-service/constants"
	"nicessm-api-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//SaveFarmerCropCalendar :""
func (s *Service) SaveFarmerCropCalendar(ctx *models.Context, farmerCropCalendar *models.FarmerCropCalendar) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	// FarmerCropCalendar.Name = s.Daos.GetUniqueID(ctx, constants.COLLECTIONFarmerCropCalendar)
	farmerCropCalendar.Status = constants.FARMERCROPCALENDARSTATUSACTIVE
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 FarmerCropCalendar.created")
	farmerCropCalendar.Created = &created
	log.Println("b4 FarmerCropCalendar.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveFarmerCropCalendar(ctx, farmerCropCalendar)
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

//UpdateFarmerCropCalendar : ""
func (s *Service) UpdateFarmerCropCalendar(ctx *models.Context, farmerCropCalendar *models.FarmerCropCalendar) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateFarmerCropCalendar(ctx, farmerCropCalendar)
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

//EnableFarmerCropCalendar : ""
func (s *Service) EnableFarmerCropCalendar(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableFarmerCropCalendar(ctx, UniqueID)
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

//DisableFarmerCropCalendar : ""
func (s *Service) DisableFarmerCropCalendar(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableFarmerCropCalendar(ctx, UniqueID)
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

//DeleteFarmerCropCalendar : ""
func (s *Service) DeleteFarmerCropCalendar(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteFarmerCropCalendar(ctx, UniqueID)
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

//GetSingleFarmerCropCalendar :""
func (s *Service) GetSingleFarmerCropCalendar(ctx *models.Context, UniqueID string) (*models.RefFarmerCropCalendar, error) {
	farmerCropCalendar, err := s.Daos.GetSingleFarmerCropCalendar(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return farmerCropCalendar, nil
}

//FilterFarmerCropCalendar :""
func (s *Service) FilterFarmerCropCalendar(ctx *models.Context, farmerCropCalendarfilter *models.FarmerCropCalendarFilter, pagination *models.Pagination) (FarmerCropCalendar []models.RefFarmerCropCalendar, err error) {
	return s.Daos.FilterFarmerCropCalendar(ctx, farmerCropCalendarfilter, pagination)
}
