package services

import (
	"errors"
	"log"
	"nicessm-api-service/constants"
	"nicessm-api-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//SaveDistrictWeatherAlertMaster :""
func (s *Service) SaveDistrictWeatherAlertMaster(ctx *models.Context, DistrictWeatherAlertMaster *models.DistrictWeatherAlertMaster) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	//organisation.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONORGANISATION)

	DistrictWeatherAlertMaster.Status = constants.DISTRICTWEATHERALERTMASTERSTATUSACTIVE
	DistrictWeatherAlertMaster.ActiveStatus = true
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 DistrictWeatherAlertMaster.created")
	DistrictWeatherAlertMaster.Created = &created
	log.Println("b4 DistrictWeatherAlertMaster.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveDistrictWeatherAlertMasterWithUpsert(ctx, DistrictWeatherAlertMaster)
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
func (s *Service) UpdateDistrictWeatherAlertMasterUpsertwithMin(ctx *models.Context, DistrictWeatherAlertMaster *models.DistrictWeatherAlertMaster) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateDistrictWeatherAlertMasterUpsertwithMin(ctx, DistrictWeatherAlertMaster)
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
func (s *Service) UpdateDistrictWeatherAlertMasterUpsertwithMax(ctx *models.Context, DistrictWeatherAlertMaster *models.DistrictWeatherAlertMaster) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateDistrictWeatherAlertMasterUpsertwithMax(ctx, DistrictWeatherAlertMaster)
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

//UpdateDistrictWeatherAlertMaster : ""
func (s *Service) UpdateDistrictWeatherAlertMaster(ctx *models.Context, DistrictWeatherAlertMaster *models.DistrictWeatherAlertMaster) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateDistrictWeatherAlertMaster(ctx, DistrictWeatherAlertMaster)
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

//EnableDistrictWeatherAlertMaster : ""
func (s *Service) EnableDistrictWeatherAlertMaster(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableDistrictWeatherAlertMaster(ctx, UniqueID)
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

//DisableDistrictWeatherAlertMaster : ""
func (s *Service) DisableDistrictWeatherAlertMaster(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableDistrictWeatherAlertMaster(ctx, UniqueID)
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

//DeleteDistrictWeatherAlertMaster : ""
func (s *Service) DeleteDistrictWeatherAlertMaster(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteDistrictWeatherAlertMaster(ctx, UniqueID)
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

//GetSingleDistrictWeatherAlertMaster :""
func (s *Service) GetSingleDistrictWeatherAlertMaster(ctx *models.Context, UniqueID string) (*models.RefDistrictWeatherAlertMaster, error) {
	DistrictWeatherAlertMaster, err := s.Daos.GetSingleDistrictWeatherAlertMaster(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return DistrictWeatherAlertMaster, nil
}

//FilterDistrictWeatherAlertMaster :""
func (s *Service) FilterDistrictWeatherAlertMaster(ctx *models.Context, DistrictWeatherAlertMasterfilter *models.DistrictWeatherAlertMasterFilter, pagination *models.Pagination) (DistrictWeatherAlertMaster []models.RefDistrictWeatherAlertMaster, err error) {
	return s.Daos.FilterDistrictWeatherAlertMaster(ctx, DistrictWeatherAlertMasterfilter, pagination)
}
func (s *Service) GetDistrictWeatherAlertMaster(ctx *models.Context, WeatherAlertMasterfilter *models.DistrictWeatherAlertMasterFilter) (DistrictWeatherAlertMaster []models.GetWeatherAlertMaster, err error) {
	return s.Daos.GetDistrictWeatherAlertMaster(ctx, WeatherAlertMasterfilter)
}
