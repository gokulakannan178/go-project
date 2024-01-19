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

//SaveGramPanchayatWeatherData :""
func (s *Service) SaveGramPanchayatWeatherData(ctx *models.Context, gramPanchayatweatherdata *models.GramPanchayatWeatherData) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	//organisation.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONORGANISATION)

	gramPanchayatweatherdata.Status = constants.GRAMPANCHAYATWEATHERDATASTATUSACTIVE
	gramPanchayatweatherdata.ActiveStatus = true
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 GramPanchayatweatherdata.created")
	gramPanchayatweatherdata.Created = &created
	log.Println("b4 GramPanchayatweatherdata.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveGramPanchayatWeatherData(ctx, gramPanchayatweatherdata)
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

//UpdateGramPanchayatWeatherData : ""
func (s *Service) UpdateGramPanchayatWeatherData(ctx *models.Context, gramPanchayatweatherdata *models.GramPanchayatWeatherData) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateGramPanchayatWeatherData(ctx, gramPanchayatweatherdata)
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

//EnableGramPanchayatWeatherData : ""
func (s *Service) EnableGramPanchayatWeatherData(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableGramPanchayatWeatherData(ctx, UniqueID)
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

//DisableGramPanchayatWeatherData : ""
func (s *Service) DisableGramPanchayatWeatherData(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableGramPanchayatWeatherData(ctx, UniqueID)
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

//DeleteGramPanchayatWeatherData : ""
func (s *Service) DeleteGramPanchayatWeatherData(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteGramPanchayatWeatherData(ctx, UniqueID)
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
func (s *Service) GetSingleGramPanchayatWeatherData(ctx *models.Context, UniqueID string) (*models.RefGramPanchayatWeatherData, error) {
	gramPanchayatweatherdata, err := s.Daos.GetSingleGramPanchayatWeatherData(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return gramPanchayatweatherdata, nil
}

//FilterGramPanchayatWeatherData :""
func (s *Service) FilterGramPanchayatWeatherData(ctx *models.Context, gramPanchayatweatherdatafilter *models.GramPanchayatWeatherDataFilter, pagination *models.Pagination) (GramPanchayatweatherdata []models.RefGramPanchayatWeatherData, err error) {
	return s.Daos.FilterGramPanchayatWeatherData(ctx, gramPanchayatweatherdatafilter, pagination)
}
func (s *Service) SaveGramPanchayatWeatherDataWithOpenWebsite(ctx *models.Context, lat string, lon string) error {
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
			GramPanchayatweatherdata := new(models.GramPanchayatWeatherData)
			GramPanchayatweatherdata.Status = constants.DISTRICWEATHERDATASTATUSACTIVE
			GramPanchayatweatherdata.ActiveStatus = true
			t := time.Now()
			created := models.Created{}
			created.On = &t
			created.By = constants.SYSTEM
			GramPanchayatweatherdata.CreatedDate = &t
			GramPanchayatweatherdata.Created = &created
			GramPanchayatweatherdata.WeatherData = v
			GramPanchayatweatherdata.Date = time.Unix(int64(v.Dt), 0)

			dberr := s.Daos.SaveGramPanchayatWeatherData(ctx, GramPanchayatweatherdata)
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
func (s *Service) SaveGramPanchayatWeatherDataWithGramPanchayat(ctx *models.Context, gramPanchayat *models.GramPanchayat) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	var lat string
	var long string
	//organisation.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONORGANISATION)
	if len(gramPanchayat.Location.Coordinates) > 0 {
		lat = fmt.Sprintf("%v", gramPanchayat.Location.Coordinates[1])
		long = fmt.Sprintf("%v", gramPanchayat.Location.Coordinates[0])
	} else {
		log.Println("pls add a location latitude and longitude---" + gramPanchayat.Name + "")
		log.Println("pls add a location latitude and longitude---" + gramPanchayat.Name + "")

	}
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		if lat != "" && long != "" {
			weatherDataMaster, err := s.GetWeatherData(ctx, lat, long)
			if err != nil {
				return errors.New("weather data not found")
			}
			for k, v := range weatherDataMaster.Daily {
				GramPanchayatweatherdata := new(models.GramPanchayatWeatherData)
				GramPanchayatweatherdata.Status = constants.DISTRICWEATHERDATASTATUSACTIVE
				GramPanchayatweatherdata.ActiveStatus = true
				//GramPanchayatweatherdata.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONGramPanchayatWEATHERDATA)
				t := time.Now()
				created := models.Created{}
				created.On = &t
				//	t.Month().String()
				created.By = constants.SYSTEM
				GramPanchayatweatherdata.CreatedDate = &t
				GramPanchayatweatherdata.Created = &created
				GramPanchayatweatherdata.WeatherData = weatherDataMaster.Daily[k]
				GramPanchayatweatherdata.Date = time.Unix(int64(v.Dt), 0)
				GramPanchayatweatherdata.GramPanchayat = gramPanchayat.ID
				GramPanchayatweatherdata.Name = gramPanchayat.Name
				GramPanchayatweatherdata.UniqueID = fmt.Sprintf("%v_%v_%v", GramPanchayatweatherdata.Date.Day(), GramPanchayatweatherdata.Date.Month().String(), GramPanchayatweatherdata.Date.Year())
				//	GramPanchayatweatherdata.WeatherData.Temp.Min
				dberr := s.Daos.SaveGramPanchayatWeatherDataWithUpsert(ctx, GramPanchayatweatherdata)
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
func (s *Service) GetSingleGramPanchayatWeatherDataWithCurrentDate(ctx *models.Context, UniqueID string) (*models.RefGramPanchayatWeatherData, error) {
	gramPanchayatweatherdata, err := s.Daos.GetSingleGramPanchayatWeatherDataWithCurrentDate(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return gramPanchayatweatherdata, nil
}
func (s *Service) SaveGramPanchayatWeatherDataCron() {
	c := context.TODO()
	ctx := app.GetApp(c, s.Daos)
	defer ctx.Client.Disconnect(c)
	//Dissemination := new(models.Dissemination)
	gramPanchayats, err := s.Daos.GetActiveGramPanchayat(ctx)
	if err != nil {
		log.Println("dissemination not found" + err.Error())
	}
	for _, v := range gramPanchayats {
		err := s.SaveGramPanchayatWeatherDataWithGramPanchayat(ctx, &v)
		if err != nil {
			log.Println("not save Weather data this GramPanchayat---" + v.Name + "" + err.Error())
			continue
		}
	}
}
