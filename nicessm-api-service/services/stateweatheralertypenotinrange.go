package services

import (
	"errors"
	"log"
	"nicessm-api-service/constants"
	"nicessm-api-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//SaveWeatherAlertNotInRange :""
func (s *Service) SaveWeatherAlertNotInRange(ctx *models.Context, WeatherAlertNotInRange *models.WeatherAlertNotInRange) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	//organisation.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONORGANISATION)

	WeatherAlertNotInRange.Status = constants.WEATHERALERTNOTINRANGESTATUSACTIVE
	WeatherAlertNotInRange.ActiveStatus = true
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 WeatherAlertNotInRange.created")
	WeatherAlertNotInRange.Created = &created
	log.Println("b4 WeatherAlertNotInRange.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveWeatherAlertNotInRange(ctx, WeatherAlertNotInRange)
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

//UpdateWeatherAlertNotInRange : ""
func (s *Service) UpdateWeatherAlertNotInRange(ctx *models.Context, WeatherAlertNotInRange *models.WeatherAlertNotInRange) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateWeatherAlertNotInRange(ctx, WeatherAlertNotInRange)
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

//EnableWeatherAlertNotInRange : ""
func (s *Service) EnableWeatherAlertNotInRange(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableWeatherAlertNotInRange(ctx, UniqueID)
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

//DisableWeatherAlertNotInRange : ""
func (s *Service) DisableWeatherAlertNotInRange(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableWeatherAlertNotInRange(ctx, UniqueID)
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

//DeleteWeatherAlertNotInRange : ""
func (s *Service) DeleteWeatherAlertNotInRange(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteWeatherAlertNotInRange(ctx, UniqueID)
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

//GetSingleWeatherAlertNotInRange :""
func (s *Service) GetSingleWeatherAlertNotInRange(ctx *models.Context, UniqueID string) (*models.RefWeatherAlertNotInRange, error) {
	WeatherAlertNotInRange, err := s.Daos.GetSingleWeatherAlertNotInRange(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return WeatherAlertNotInRange, nil
}

//FilterWeatherAlertNotInRange :""
func (s *Service) FilterWeatherAlertNotInRange(ctx *models.Context, WeatherAlertNotInRangefilter *models.WeatherAlertNotInRangeFilter, pagination *models.Pagination) (WeatherAlertNotInRange []models.RefWeatherAlertNotInRange, err error) {
	return s.Daos.FilterWeatherAlertNotInRange(ctx, WeatherAlertNotInRangefilter, pagination)
}
