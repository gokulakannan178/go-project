package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"lgfweather-service/app"
	"lgfweather-service/constants"
	"lgfweather-service/models"
	"log"
	"net/url"
	"strings"
	"time"

	"github.com/jlaffaye/ftp"
	"go.mongodb.org/mongo-driver/mongo"
)

//SaveStateWeatherData :""
func (s *Service) SaveStateWeatherData(ctx *models.Context, stateweatherdata *models.StateWeatherData) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	//organisation.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONORGANISATION)

	stateweatherdata.Status = constants.DISTRICWEATHERDATASTATUSACTIVE
	stateweatherdata.ActiveStatus = true
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 stateweatherdata.created")
	stateweatherdata.Created = &created
	log.Println("b4 stateweatherdata.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveStateWeatherData(ctx, stateweatherdata)
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

//UpdateStateWeatherData : ""
func (s *Service) UpdateStateWeatherData(ctx *models.Context, stateweatherdata *models.StateWeatherData) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateStateWeatherData(ctx, stateweatherdata)
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

//EnableStateWeatherData : ""
func (s *Service) EnableStateWeatherData(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableStateWeatherData(ctx, UniqueID)
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

//DisableStateWeatherData : ""
func (s *Service) DisableStateWeatherData(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableStateWeatherData(ctx, UniqueID)
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

//DeleteStateWeatherData : ""
func (s *Service) DeleteStateWeatherData(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteStateWeatherData(ctx, UniqueID)
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
func (s *Service) GetSingleStateWeatherData(ctx *models.Context, UniqueID string) (*models.RefStateWeatherData, error) {
	stateweatherdata, err := s.Daos.GetSingleStateWeatherData(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return stateweatherdata, nil
}

//FilterStateWeatherData :""
func (s *Service) FilterStateWeatherData(ctx *models.Context, stateweatherdatafilter *models.StateWeatherDataFilter, pagination *models.Pagination) (stateweatherdata []models.RefStateWeatherData, err error) {
	return s.Daos.FilterStateWeatherData(ctx, stateweatherdatafilter, pagination)
}
func (s *Service) SaveStateWeatherDataWithOpenWebsite(ctx *models.Context, lat string, lon string) error {
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
			stateweatherdata := new(models.StateWeatherData)
			stateweatherdata.Status = constants.DISTRICWEATHERDATASTATUSACTIVE
			stateweatherdata.ActiveStatus = true
			t := time.Now()
			created := models.Created{}
			created.On = &t
			created.By = constants.SYSTEM
			stateweatherdata.CreatedDate = &t
			stateweatherdata.Created = &created
			stateweatherdata.WeatherData = v
			stateweatherdata.Date = time.Unix(int64(v.Dt), 0)

			dberr := s.Daos.SaveStateWeatherData(ctx, stateweatherdata)
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
func (s *Service) SaveStateWeatherDataWithState(ctx *models.Context, state *models.State) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	var lat string
	var long string
	//organisation.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONORGANISATION)
	if len(state.Location.Coordinates) > 0 {
		lat = fmt.Sprintf("%v", state.Location.Coordinates[1])
		long = fmt.Sprintf("%v", state.Location.Coordinates[0])
	} else {
		log.Println("pls add a location latitude and longitude---" + state.Name + "")
		log.Println("pls add a location latitude and longitude---" + state.Name + "")

	}
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		if lat != "" && long != "" {
			weatherDataMaster, err := s.GetWeatherData(ctx, lat, long)
			if err != nil {
				return errors.New("weather data not found")
			}
			for k, v := range weatherDataMaster.Daily {
				stateweatherdata := new(models.StateWeatherData)
				stateweatherdata.Status = constants.DISTRICWEATHERDATASTATUSACTIVE
				stateweatherdata.ActiveStatus = true
				//stateweatherdata.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONSTATEWEATHERDATA)
				t := time.Now()
				created := models.Created{}
				created.On = &t
				//	t.Month().String()
				created.By = constants.SYSTEM
				stateweatherdata.CreatedDate = &t
				stateweatherdata.Created = &created
				stateweatherdata.WeatherData = weatherDataMaster.Daily[k]
				stateweatherdata.Date = time.Unix(int64(v.Dt), 0)
				stateweatherdata.State = state.ID
				stateweatherdata.Name = state.Name
				stateweatherdata.UniqueID = fmt.Sprintf("%v_%v_%v", stateweatherdata.Date.Day(), stateweatherdata.Date.Month().String(), stateweatherdata.Date.Year())
				//	stateweatherdata.WeatherData.Temp.Min
				dberr := s.Daos.SaveStateWeatherDataWithUpsert(ctx, stateweatherdata)
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
func (s *Service) GetSingleStateWeatherDataWithCurrentDate(ctx *models.Context, UniqueID string) (*models.RefStateWeatherData, error) {
	stateweatherdata, err := s.Daos.GetSingleStateWeatherDataWithCurrentDate(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return stateweatherdata, nil
}
func (s *Service) SaveStateWeatherDataCron() {
	c := context.TODO()
	ctx := app.GetApp(c, s.Daos)
	defer ctx.Client.Disconnect(c)
	//Dissemination := new(models.Dissemination)
	States, err := s.Daos.GetActiveState(ctx)
	if err != nil {
		log.Println("dissemination not found" + err.Error())
	}
	for _, v := range States {
		err := s.SaveStateWeatherDataWithState(ctx, &v)
		if err != nil {
			log.Println("not save Weather data this state---" + v.Name + "" + err.Error())
			continue
		}
	}
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
func (s *Service) GetWeatherDataWithImd(ctx *models.Context) ([]string, error) {
	client, err := ftp.Dial("103.215.208.49:21")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if err := client.Login("imdgfs", "imdgfs@2012"); err != nil {
		fmt.Println(err)
		return nil, err
	}
	var date string
	//var param1 string
	IMDWeatherDataFile := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.IMDWEATHERDATA_FILE)
	IMDWeatherDataUrl := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.IMDWEATHERDATA_URL)
	IMDWeatherDataLocation := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.IMDWEATHERDATA_LOCATION)
	//var t time.Time
	t := time.Now()
	month := s.Shared.GetMonthInt(t.Month().String())
	yearstr := fmt.Sprintf("%v", t.Year())
	years := strings.Split(yearstr, "20")
	year := years[1]
	if t.Day() < 10 {
		date = fmt.Sprintf("%v%v0%v", year, month, t.Day())
	} else {
		date = fmt.Sprintf("%v%v%v", year, month, t.Day())
	}
	IMDURL := fmt.Sprintf("%v%v%v%v", IMDWeatherDataUrl, IMDWeatherDataLocation, IMDWeatherDataFile, date)

	fmt.Println("URL===>", IMDURL)
	r, err := client.Retr(IMDURL)
	if err != nil {
		panic(err)
	}
	//defer r.Close()

	defer r.Close()
	body, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
		log.Println("result not found")
	}
	lines := strings.Split(string(body), "\n")
	//var words []string
	fmt.Println("lines===>", len(lines))
	return lines, nil
}
func (s *Service) GetWeatherDataWithImdWithState(ctx *models.Context, ImdName string, ImdFile string) ([]string, error) {
	client, err := ftp.Dial("103.215.208.49:21")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if err := client.Login("imdgfs", "imdgfs@2012"); err != nil {
		fmt.Println(err)
		return nil, err
	}
	var date string
	//var param1 string
	//	IMDWeatherDataFile := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.IMDWEATHERDATA_FILE)
	//IMDWeatherDataUrl := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.IMDWEATHERDATA_URL)
	//IMDWeatherDataLocation := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.IMDWEATHERDATA_LOCATION)
	//var t time.Time
	t := time.Now()
	month := s.Shared.GetMonthInt(t.Month().String())
	yearstr := fmt.Sprintf("%v", t.Year())
	years := strings.Split(yearstr, "20")
	year := years[1]
	if t.Day() < 10 {
		date = fmt.Sprintf("%v%v0%v", year, month, t.Day())
	} else {
		date = fmt.Sprintf("%v%v%v", year, month, t.Day())
	}
	date = fmt.Sprintf("%v%v0%v", year, month, 2)
	imdurl := "DIST_BLOCK_FT1534/2020/dfcst/"
	param1 := "madhya-pradesh"
	file := "/dfcst2000z"
	//IMDURL := fmt.Sprintf("%v%v/%v%v", IMDWeatherDataUrl, ImdName, ImdFile, date)
	IMDURL := fmt.Sprintf("%v%v%v%v", imdurl, param1, file, date)

	fmt.Println("URL===>", IMDURL)
	r, err := client.Retr(IMDURL)
	if err != nil {
		panic(err)
	}
	//defer r.Close()

	defer r.Close()
	body, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
		log.Println("result not found")
	}
	lines := strings.Split(string(body), "\n")
	//	var words []string

	fmt.Println("lines===>", len(lines))
	return lines, nil
}
