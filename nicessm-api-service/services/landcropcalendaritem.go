package services

import (
	"errors"
	"log"
	"nicessm-api-service/constants"
	"nicessm-api-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//SaveLandCropCalendar :""
func (s *Service) SaveLandCropCalendar(ctx *models.Context, landCropCalendar *models.LandCropCalendar) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	// LandCropCalendar.Name = s.Daos.GetUniqueID(ctx, constants.COLLECTIONLandCropCalendar)
	landCropCalendar.Status = constants.LANDCROPCALENDARITEMSTATUSACTIVE
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 landCropCalendar.created")
	landCropCalendar.Created = &created
	log.Println("b4 landCropCalendar.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveLandCropCalendar(ctx, landCropCalendar)
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

//UpdateLandCropCalendar : ""
func (s *Service) UpdateLandCropCalendar(ctx *models.Context, landCropCalendar *models.LandCropCalendar) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateLandCropCalendar(ctx, landCropCalendar)
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

//EnableLandCropCalendar : ""
func (s *Service) EnableLandCropCalendar(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableLandCropCalendar(ctx, UniqueID)
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

//DisableLandCropCalendar : ""
func (s *Service) DisableLandCropCalendar(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableLandCropCalendar(ctx, UniqueID)
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

//DeleteLandCropCalendar : ""
func (s *Service) DeleteLandCropCalendar(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteLandCropCalendar(ctx, UniqueID)
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

//GetSingleLandCropCalendar :""
func (s *Service) GetSingleLandCropCalendar(ctx *models.Context, UniqueID string) (*models.RefLandCropCalendar, error) {
	landCropCalendar, err := s.Daos.GetSingleLandCropCalendar(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return landCropCalendar, nil
}

//FilterLandCropCalendar :""
func (s *Service) FilterLandCropCalendar(ctx *models.Context, landCropCalendarfilter *models.LandCropCalendarFilter, pagination *models.Pagination) (LandCropCalendar []models.RefLandCropCalendar, err error) {
	return s.Daos.FilterLandCropCalendar(ctx, landCropCalendarfilter, pagination)
}
