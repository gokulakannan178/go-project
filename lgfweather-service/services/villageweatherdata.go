package services

import (
	"context"
	"errors"
	"fmt"
	"lgfweather-service/app"
	"lgfweather-service/constants"
	"lgfweather-service/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//SaveVillageWeatherData :""
func (s *Service) SaveVillageWeatherData(ctx *models.Context, villageweatherdata *models.VillageWeatherData) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	//organisation.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONORGANISATION)

	villageweatherdata.Status = constants.VILLAGEWEATHERDATASTATUSACTIVE
	villageweatherdata.ActiveStatus = true
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 Villageweatherdata.created")
	villageweatherdata.Created = &created
	log.Println("b4 Villageweatherdata.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveVillageWeatherData(ctx, villageweatherdata)
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

//UpdateVillageWeatherData : ""
func (s *Service) UpdateVillageWeatherData(ctx *models.Context, villageweatherdata *models.VillageWeatherData) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateVillageWeatherData(ctx, villageweatherdata)
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

//EnableVillageWeatherData : ""
func (s *Service) EnableVillageWeatherData(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableVillageWeatherData(ctx, UniqueID)
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

//DisableVillageWeatherData : ""
func (s *Service) DisableVillageWeatherData(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableVillageWeatherData(ctx, UniqueID)
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

//DeleteVillageWeatherData : ""
func (s *Service) DeleteVillageWeatherData(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteVillageWeatherData(ctx, UniqueID)
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

//GetSingleVaccine :""
func (s *Service) GetSingleVillageWeatherData(ctx *models.Context, UniqueID string) (*models.RefVillageWeatherData, error) {
	villageweatherdata, err := s.Daos.GetSingleVillageWeatherData(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return villageweatherdata, nil
}

//FilterVillageWeatherData :""
func (s *Service) FilterVillageWeatherData(ctx *models.Context, villageweatherdatafilter *models.VillageWeatherDataFilter, pagination *models.Pagination) (Villageweatherdata []models.RefVillageWeatherData, err error) {
	return s.Daos.FilterVillageWeatherData(ctx, villageweatherdatafilter, pagination)
}
func (s *Service) SaveVillageWeatherDataWithOpenWebsite(ctx *models.Context, lat string, lon string) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	//organisation.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONORGANISATION)

	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		weatherDataMaster, err := s.GetWeatherData(ctx, lat, lon)
		if err != nil {
			return errors.New("weather data not found")
		}
		for _, v := range weatherDataMaster.Daily {
			villageweatherdata := new(models.VillageWeatherData)
			villageweatherdata.Status = constants.DISTRICWEATHERDATASTATUSACTIVE
			villageweatherdata.ActiveStatus = true
			t := time.Now()
			created := models.Created{}
			created.On = &t
			created.By = constants.SYSTEM
			villageweatherdata.CreatedDate = &t
			villageweatherdata.Created = &created
			villageweatherdata.WeatherData = v
			villageweatherdata.Date = time.Unix(int64(v.Dt), 0)

			dberr := s.Daos.SaveVillageWeatherData(ctx, villageweatherdata)
			if dberr != nil {
				if err1 := ctx.Session.AbortTransaction(sc); err1 != nil {
					log.Println("err in abort")
					return errors.New("Transaction Aborted with error" + err1.Error())
				}
				log.Println("err in abort out")
				return errors.New("Transaction Aborted - " + dberr.Error())
			}
		}

		return nil

	}); err != nil {
		return err
	}
	return nil
}
func (s *Service) SaveVillageWeatherDataWithVillage(ctx *models.Context, village *models.Village) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	var lat string
	var long string
	//organisation.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONORGANISATION)
	if len(village.Location.Coordinates) > 0 {
		lat = fmt.Sprintf("%v", village.Location.Coordinates[1])
		long = fmt.Sprintf("%v", village.Location.Coordinates[0])
	} else {
		log.Println("pls add a location latitude and longitude---" + village.Name + "")
		log.Println("pls add a location latitude and longitude---" + village.Name + "")

	}
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		if lat != "" && long != "" {
			weatherDataMaster, err := s.GetWeatherData(ctx, lat, long)
			if err != nil {
				return errors.New("weather data not found")
			}
			for k, v := range weatherDataMaster.Daily {
				villageweatherdata := new(models.VillageWeatherData)
				villageweatherdata.Status = constants.DISTRICWEATHERDATASTATUSACTIVE
				villageweatherdata.ActiveStatus = true
				//Villageweatherdata.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONVillageWEATHERDATA)
				t := time.Now()
				created := models.Created{}
				created.On = &t
				//	t.Month().String()
				created.By = constants.SYSTEM
				villageweatherdata.CreatedDate = &t
				villageweatherdata.Created = &created
				villageweatherdata.WeatherData = weatherDataMaster.Daily[k]
				villageweatherdata.Date = time.Unix(int64(v.Dt), 0)
				villageweatherdata.Village = village.ID
				villageweatherdata.Name = village.Name
				villageweatherdata.UniqueID = fmt.Sprintf("%v_%v_%v", villageweatherdata.Date.Day(), villageweatherdata.Date.Month().String(), villageweatherdata.Date.Year())
				//	Villageweatherdata.WeatherData.Temp.Min
				dberr := s.Daos.SaveVillageWeatherDataWithUpsert(ctx, villageweatherdata)
				if dberr != nil {

					return errors.New("Db Error" + dberr.Error())
				}
			}
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
func (s *Service) GetSingleVillageWeatherDataWithCurrentDate(ctx *models.Context, UniqueID string) (*models.RefVillageWeatherData, error) {
	villageweatherdata, err := s.Daos.GetSingleVillageWeatherDataWithCurrentDate(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return villageweatherdata, nil
}
func (s *Service) SaveVillageWeatherDataCron() {
	c := context.TODO()
	ctx := app.GetApp(c, s.Daos)
	defer ctx.Client.Disconnect(c)
	//Dissemination := new(models.Dissemination)
	Villages, err := s.Daos.GetActiveVillage(ctx)
	if err != nil {
		log.Println("dissemination not found" + err.Error())
	}
	for _, v := range Villages {
		err := s.SaveVillageWeatherDataWithVillage(ctx, &v)
		if err != nil {
			log.Println("not save Weather data this Village---" + v.Name + "" + err.Error())
			continue
		}
	}
}
