package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"nicessm-api-service/constants"
	"nicessm-api-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//SaveStateWeatherAlertMaster :""
func (s *Service) SaveStateWeatherAlertMaster(ctx *models.Context, StateWeatherAlertMaster *models.StateWeatherAlertMaster) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	//organisation.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONORGANISATION)

	StateWeatherAlertMaster.Status = constants.STATEWEATHERALERTMASTERSTATUSACTIVE
	StateWeatherAlertMaster.ActiveStatus = true
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 StateWeatherAlertMaster.created")
	StateWeatherAlertMaster.Created = &created
	log.Println("b4 StateWeatherAlertMaster.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveStateWeatherAlertMasterWithUpsert(ctx, StateWeatherAlertMaster)
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
func (s *Service) UpdateStateWeatherAlertMasterUpsertwithMin(ctx *models.Context, StateWeatherAlertMaster *models.StateWeatherAlertMaster) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateStateWeatherAlertMasterUpsertwithMin(ctx, StateWeatherAlertMaster)
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
func (s *Service) UpdateStateWeatherAlertMasterUpsertwithMax(ctx *models.Context, StateWeatherAlertMaster *models.StateWeatherAlertMaster) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateStateWeatherAlertMasterUpsertwithMax(ctx, StateWeatherAlertMaster)
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

//UpdateStateWeatherAlertMaster : ""
func (s *Service) UpdateStateWeatherAlertMaster(ctx *models.Context, StateWeatherAlertMaster *models.StateWeatherAlertMaster) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateStateWeatherAlertMaster(ctx, StateWeatherAlertMaster)
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

//EnableStateWeatherAlertMaster : ""
func (s *Service) EnableStateWeatherAlertMaster(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableStateWeatherAlertMaster(ctx, UniqueID)
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

//DisableStateWeatherAlertMaster : ""
func (s *Service) DisableStateWeatherAlertMaster(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableStateWeatherAlertMaster(ctx, UniqueID)
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

//DeleteStateWeatherAlertMaster : ""
func (s *Service) DeleteStateWeatherAlertMaster(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteStateWeatherAlertMaster(ctx, UniqueID)
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

//GetSingleStateWeatherAlertMaster :""
func (s *Service) GetSingleStateWeatherAlertMaster(ctx *models.Context, UniqueID string) (*models.RefStateWeatherAlertMaster, error) {
	StateWeatherAlertMaster, err := s.Daos.GetSingleStateWeatherAlertMaster(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return StateWeatherAlertMaster, nil
}

//FilterStateWeatherAlertMaster :""
func (s *Service) FilterStateWeatherAlertMaster(ctx *models.Context, StateWeatherAlertMasterfilter *models.StateWeatherAlertMasterFilter, pagination *models.Pagination) (StateWeatherAlertMaster []models.RefStateWeatherAlertMaster, err error) {
	return s.Daos.FilterStateWeatherAlertMaster(ctx, StateWeatherAlertMasterfilter, pagination)
}
func (s *Service) GetStateWeatherAlertMaster(ctx *models.Context, WeatherAlertMasterfilter *models.StateWeatherAlertMasterFilter) (StateWeatherAlertMaster []models.GetWeatherAlertMaster, err error) {
	return s.Daos.GetStateWeatherAlertMaster(ctx, WeatherAlertMasterfilter)
}
func (s *Service) GetWeatherData(ctx *models.Context, latt string, long string) (*models.WeatherDataMaster, error) {
	WeatherDataUrl := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.WEATHERDATA_URL)
	WeatherDataId := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.WEATHERDATA_ID)
	fmt.Println("WeatherDataUrl : ", WeatherDataUrl)
	fmt.Println("WeatherDataId : ", WeatherDataId)

	var URL *url.URL
	URL, err := url.Parse(WeatherDataUrl)
	if err != nil {
		return nil, errors.New("url const err - " + err.Error())
	}
	parameters := url.Values{}
	parameters.Add("appid", WeatherDataId)
	parameters.Add("lon", long)
	parameters.Add("lat", latt)
	parameters.Add("units", "metric")

	URL.RawQuery = parameters.Encode()
	fmt.Println("URL : ", URL.String())
	var WeatherData models.WeatherDataMaster
	resp, err := s.Shared.Get(URL.String(), nil)
	// if err != nil {
	// 	return errors.New("api err - " + err.Error())
	// }
	// log.Println(resp)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal(body, &WeatherData); err != nil { // Parse []byte to the go struct pointer

		fmt.Println("Can not unmarshal JSON", err)
	}
	return &WeatherData, nil
}
