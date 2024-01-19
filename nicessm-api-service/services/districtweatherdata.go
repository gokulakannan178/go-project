package services

import (
	"context"
	"errors"
	"fmt"
	"log"
	"nicessm-api-service/app"
	"nicessm-api-service/constants"
	"nicessm-api-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//SaveDistrictWeatherData :""
func (s *Service) SaveDistrictWeatherData(ctx *models.Context, districtweatherdata *models.DistrictWeatherData) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	//organisation.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONORGANISATION)

	districtweatherdata.Status = constants.DISTRICWEATHERDATASTATUSACTIVE
	districtweatherdata.ActiveStatus = true
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 districtweatherdata.created")
	districtweatherdata.Created = &created
	log.Println("b4 districtweatherdata.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveDistrictWeatherData(ctx, districtweatherdata)
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

//UpdateDistrictWeatherData : ""
func (s *Service) UpdateDistrictWeatherData(ctx *models.Context, districtweatherdata *models.DistrictWeatherData) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateDistrictWeatherData(ctx, districtweatherdata)
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

//EnableDistrictWeatherData : ""
func (s *Service) EnableDistrictWeatherData(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableDistrictWeatherData(ctx, UniqueID)
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

//DisableDistrictWeatherData : ""
func (s *Service) DisableDistrictWeatherData(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableDistrictWeatherData(ctx, UniqueID)
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

//DeleteDistrictWeatherData : ""
func (s *Service) DeleteDistrictWeatherData(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteDistrictWeatherData(ctx, UniqueID)
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
func (s *Service) GetSingleDistrictWeatherData(ctx *models.Context, UniqueID string) (*models.RefDistrictWeatherData, error) {
	districtweatherdata, err := s.Daos.GetSingleDistrictWeatherData(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return districtweatherdata, nil
}

//FilterDistrictWeatherData :""
func (s *Service) FilterDistrictWeatherData(ctx *models.Context, districtweatherdatafilter *models.DistrictWeatherDataFilter, pagination *models.Pagination) (districtweatherdata []models.RefDistrictWeatherData, err error) {
	return s.Daos.FilterDistrictWeatherData(ctx, districtweatherdatafilter, pagination)
}
func (s *Service) SaveDistrictWeatherDataWithDistrict(ctx *models.Context, district *models.District) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	var lat string
	var long string
	//organisation.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONORGANISATION)
	if len(district.Location.Coordinates) > 0 {
		lat = fmt.Sprintf("%v", district.Location.Coordinates[1])
		long = fmt.Sprintf("%v", district.Location.Coordinates[0])
	} else {
		log.Println("pls add a location latitude and longitude---" + district.Name + "")
		log.Println("pls add a location latitude and longitude---" + district.Name + "")
	}
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		if lat != "" && long != "" {
			weatherDataMaster, err := s.GetWeatherData(ctx, lat, long)
			if err != nil {
				return errors.New("weather data not found")
			}
			for k, v := range weatherDataMaster.Daily {
				districtweatherdata := new(models.DistrictWeatherData)
				districtweatherdata.Status = constants.DISTRICWEATHERDATASTATUSACTIVE
				districtweatherdata.ActiveStatus = true
				//stateweatherdata.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONSTATEWEATHERDATA)
				t := time.Now()
				created := models.Created{}
				created.On = &t
				//	t.Month().String()
				created.By = constants.SYSTEM
				districtweatherdata.CreatedDate = &t
				districtweatherdata.Created = &created
				districtweatherdata.WeatherData = weatherDataMaster.Daily[k]
				districtweatherdata.Date = time.Unix(int64(v.Dt), 0)
				fmt.Println("district date==>", districtweatherdata.Date)
				districtweatherdata.District = district.ID
				districtweatherdata.Name = district.Name
				districtweatherdata.UniqueID = fmt.Sprintf("%v_%v_%v", districtweatherdata.Date.Day(), districtweatherdata.Date.Month().String(), districtweatherdata.Date.Year())
				fmt.Println("district UniqueID==>", districtweatherdata.UniqueID)
				//	stateweatherdata.WeatherData.Temp.Min
				dberr := s.Daos.SaveDistrictWeatherDataWithUpsert(ctx, districtweatherdata)
				if dberr != nil {

					return errors.New("Db Error" + dberr.Error())
				}
				err := s.Daos.DistrictWeatherAlertMaster(ctx, districtweatherdata)
				if err != nil {
					return errors.New("WeatherAlertMaster not Saved" + err.Error())
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
func (s *Service) SaveDistrictWeatherDataCron() {
	c := context.TODO()
	ctx := app.GetApp(c, s.Daos)
	defer ctx.Client.Disconnect(c)
	//Dissemination := new(models.Dissemination)
	districts, err := s.Daos.GetActiveDistrict(ctx)
	if err != nil {
		log.Println("dissemination not found" + err.Error())
	}
	for _, v := range districts {
		err := s.SaveDistrictWeatherDataWithDistrict(ctx, &v)
		if err != nil {
			log.Println("not save Weather data this state---" + v.Name + "" + err.Error())
			continue
		}
	}
}
