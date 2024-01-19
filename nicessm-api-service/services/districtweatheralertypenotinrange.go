package services

import (
	"errors"
	"log"
	"nicessm-api-service/constants"
	"nicessm-api-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//SaveDistrictWeatherAlertNotInRange :""
func (s *Service) SaveDistrictWeatherAlertNotInRange(ctx *models.Context, DistrictWeatherAlertNotInRange *models.DistrictWeatherAlertNotInRange) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	//organisation.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONORGANISATION)

	DistrictWeatherAlertNotInRange.Status = constants.DISTRICTWEATHERALERTNOTINRANGESTATUSACTIVE
	DistrictWeatherAlertNotInRange.ActiveStatus = true
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 DistrictWeatherAlertNotInRange.created")
	DistrictWeatherAlertNotInRange.Created = &created
	log.Println("b4 DistrictWeatherAlertNotInRange.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveDistrictWeatherAlertNotInRange(ctx, DistrictWeatherAlertNotInRange)
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

//UpdateDistrictWeatherAlertNotInRange : ""
func (s *Service) UpdateDistrictWeatherAlertNotInRange(ctx *models.Context, DistrictWeatherAlertNotInRange *models.DistrictWeatherAlertNotInRange) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateDistrictWeatherAlertNotInRange(ctx, DistrictWeatherAlertNotInRange)
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

//EnableDistrictWeatherAlertNotInRange : ""
func (s *Service) EnableDistrictWeatherAlertNotInRange(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableDistrictWeatherAlertNotInRange(ctx, UniqueID)
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

//DisableDistrictWeatherAlertNotInRange : ""
func (s *Service) DisableDistrictWeatherAlertNotInRange(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableDistrictWeatherAlertNotInRange(ctx, UniqueID)
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

//DeleteDistrictWeatherAlertNotInRange : ""
func (s *Service) DeleteDistrictWeatherAlertNotInRange(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteDistrictWeatherAlertNotInRange(ctx, UniqueID)
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

//GetSingleDistrictWeatherAlertNotInRange :""
func (s *Service) GetSingleDistrictWeatherAlertNotInRange(ctx *models.Context, UniqueID string) (*models.RefDistrictWeatherAlertNotInRange, error) {
	DistrictWeatherAlertNotInRange, err := s.Daos.GetSingleDistrictWeatherAlertNotInRange(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return DistrictWeatherAlertNotInRange, nil
}

//FilterDistrictWeatherAlertNotInRange :""
func (s *Service) FilterDistrictWeatherAlertNotInRange(ctx *models.Context, DistrictWeatherAlertNotInRangefilter *models.DistrictWeatherAlertNotInRangeFilter, pagination *models.Pagination) (DistrictWeatherAlertNotInRange []models.RefDistrictWeatherAlertNotInRange, err error) {
	return s.Daos.FilterDistrictWeatherAlertNotInRange(ctx, DistrictWeatherAlertNotInRangefilter, pagination)
}
