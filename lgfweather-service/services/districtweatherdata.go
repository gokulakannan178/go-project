package services

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"lgfweather-service/app"
	"lgfweather-service/config"
	"lgfweather-service/constants"
	"lgfweather-service/models"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/jlaffaye/ftp"
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
			for _, v := range weatherDataMaster.Daily {
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
				//	districtweatherdata.WeatherData = weatherDataMaster.Daily[k]
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
				// err := s.Daos.WeatherAlertMaster(ctx, stateweatherdata)
				// if err != nil {
				// 	return errors.New("WeatherAlertMaster not Saved" + err.Error())
				// }
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
func (s *Service) SaveDistrictWeatherDataWithImd(ctx *models.Context, lines []string) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	config := config.NewConfig("districtimd", "config")
	maxLength := config.GetInt(constants.IMDDISTRICT + "." + constants.IMDmaxlength)

	// stateNo := config.GetInt(constants.IMDDISTRICT + "." + constants.IMDstate)
	districtNo := config.GetInt(constants.IMDDISTRICT + "." + constants.IMDdistrict2)
	latNo := config.GetInt(constants.IMDDISTRICT + "." + constants.IMDlat2)
	longNo := config.GetInt(constants.IMDDISTRICT + "." + constants.IMDlong2)
	altoNo := config.GetInt(constants.IMDDISTRICT + "." + constants.IMDalto2)
	pcodNo := config.GetInt(constants.IMDDISTRICT + "." + constants.IMDpcod2)
	dayNo := config.GetInt(constants.IMDDISTRICT + "." + constants.IMDday2)
	monthNo := config.GetInt(constants.IMDDISTRICT + "." + constants.IMDmonth2)
	yearNo := config.GetInt(constants.IMDDISTRICT + "." + constants.IMDyear2)
	msplNo := config.GetInt(constants.IMDDISTRICT + "." + constants.IMDmspl2)
	icidNo := config.GetInt(constants.IMDDISTRICT + "." + constants.IMDicid2)
	rainfallNo := config.GetInt(constants.IMDDISTRICT + "." + constants.IMDrainfall2)
	maxtempNo := config.GetInt(constants.IMDDISTRICT + "." + constants.IMDmaxtemp2)
	mintempNo := config.GetInt(constants.IMDDISTRICT + "." + constants.IMDmintemp2)
	maxrelhumNo := config.GetInt(constants.IMDDISTRICT + "." + constants.IMDmaxrelhum2)
	minremhumNo := config.GetInt(constants.IMDDISTRICT + "." + constants.IMDminremhum2)
	windspeedNo := config.GetInt(constants.IMDDISTRICT + "." + constants.IMDwindspeed2)
	winddirectionNo := config.GetInt(constants.IMDDISTRICT + "." + constants.IMDwinddirection2)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		if len(lines) > 0 {
			for _, v := range lines {
				words := strings.Fields(v)
				if len(words) < maxLength {

					dwd := models.DistrictWeatherData{}
					dwd.Name = words[districtNo]
					var err error
					latStr := words[longNo]
					longStr := words[latNo]

					if latStr != "" && latStr != "" {
						lat, err := strconv.ParseFloat(latStr, 64)
						if err != nil {
							log.Println("Error in geting latStr", err.Error())
							err = nil
						}
						long, err := strconv.ParseFloat(longStr, 64)
						if err != nil {
							log.Println("Error in geting latStr", err.Error())
							err = nil
						}
						dwd.Location = models.Location{
							Type:        "point",
							Coordinates: []float64{long, lat},
						}
					}

					alto := words[altoNo]
					if alto != "" {
						dwd.Alto, err = strconv.ParseFloat(alto, 64)
						if err != nil {
							log.Println("Error in geting alto", err.Error())
							err = nil
						}
					}
					dwd.Pcod = words[pcodNo]

					yearStr := words[yearNo]
					monthStr := words[monthNo]
					dayStr := words[dayNo]
					var t time.Time
					if yearStr != "" {
						year, err := strconv.ParseInt(yearStr[0:4], 10, 0)
						if err != nil {
							log.Println("Error in geting year", err.Error())
							err = nil
						}
						month, err := strconv.ParseInt(monthStr, 10, 0)
						if err != nil {
							log.Println("Error in geting year", err.Error())
							err = nil
						}
						day, err := strconv.ParseInt(dayStr, 10, 0)
						if err != nil {
							log.Println("Error in geting year", err.Error())
							err = nil
						}
						loc, _ := time.LoadLocation("Asia/Kolkata")

						t := time.Date(int(year), time.Month(int(month)), int(day), 0, 0, 0, 0, loc)
						dwd.Date = t

					}

					mspl := words[msplNo]
					if mspl != "" {
						dwd.Mspl, err = strconv.ParseFloat(mspl, 64)
						if err != nil {
							log.Println("Error in geting mspl", err.Error())
							err = nil
						}
					}
					icid := words[icidNo]
					if icid != "" {
						dwd.Icid, err = strconv.ParseFloat(icid, 64)
						if err != nil {
							log.Println("Error in geting icid", err.Error())
							err = nil
						}
					}
					rainfall := words[rainfallNo]
					if rainfall != "" {
						dwd.WeatherData.Rain, err = strconv.ParseFloat(rainfall, 64)
						if err != nil {
							log.Println("Error in geting rainfall", err.Error())
							err = nil
						}
					}
					maxtemp := words[maxtempNo]
					if maxtemp != "" {
						dwd.WeatherData.Temp.Max, err = strconv.ParseFloat(maxtemp, 64)
						if err != nil {
							log.Println("Error in geting maxtemp", err.Error())
							err = nil
						}
					}
					mintemp := words[mintempNo]
					if mintemp != "" {
						dwd.WeatherData.Temp.Min, err = strconv.ParseFloat(mintemp, 64)
						if err != nil {
							log.Println("Error in geting mintemp", err.Error())
							err = nil
						}
					}
					maxrelhum := words[maxrelhumNo]
					if maxrelhum != "" {
						dwd.WeatherData.HumidityMax, err = strconv.ParseFloat(maxrelhum, 64)
						if err != nil {
							log.Println("Error in geting maxrelhum", err.Error())
							err = nil
						}
					}
					minremhum := words[minremhumNo]
					if minremhum != "" {
						dwd.WeatherData.HumidityMin, err = strconv.ParseFloat(minremhum, 64)
						if err != nil {
							log.Println("Error in geting minremhum", err.Error())
							err = nil
						}
					}
					windspeed := words[windspeedNo]
					if windspeed != "" {
						dwd.WeatherData.Windspeed, err = strconv.ParseFloat(windspeed, 64)
						if err != nil {
							log.Println("Error in geting windspeed", err.Error())
							err = nil
						}
					}
					winddirection := words[winddirectionNo]
					if winddirection != "" {
						dwd.WeatherData.Winddeg, err = strconv.ParseFloat(winddirection, 64)
						if err != nil {
							log.Println("Error in geting winddirection", err.Error())
							err = nil
						}
					}
					dwd.UniqueID = fmt.Sprintf("%v%v%v", t.Day(), t.Month().String(), t.Year())
					dberr := s.Daos.SaveDistrictWeatherDataNameWithUpsert(ctx, &dwd)
					if dberr != nil {

						return errors.New("Db Error" + dberr.Error())
					}

				} else {
					log.Println("words not in range", len(words))
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
func (s *Service) SaveDistrictWeatherDataWithImdWithState(ctx *models.Context, lines string, state string) error {
	//log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	config := config.NewConfig("districtimd", "config")
	//maxLength := config.GetInt(constants.IMDDISTRICT + "." + constants.IMDmaxlength)
	// stateNo := config.GetInt(constants.IMDDISTRICT + "." + constants.IMDstate)
	// state1No := config.GetInt(constants.IMDDISTRICT + "." + constants.IMDstate1)
	districtNo2 := config.GetInt(constants.IMDDISTRICT + "." + constants.IMDdistrict2)
	latNo2 := config.GetInt(constants.IMDDISTRICT + "." + constants.IMDlat2)
	longNo2 := config.GetInt(constants.IMDDISTRICT + "." + constants.IMDlong2)
	altoNo2 := config.GetInt(constants.IMDDISTRICT + "." + constants.IMDalto2)
	pcodNo2 := config.GetInt(constants.IMDDISTRICT + "." + constants.IMDpcod2)
	dayNo2 := config.GetInt(constants.IMDDISTRICT + "." + constants.IMDday2)
	monthNo2 := config.GetInt(constants.IMDDISTRICT + "." + constants.IMDmonth2)
	yearNo2 := config.GetInt(constants.IMDDISTRICT + "." + constants.IMDyear2)
	msplNo2 := config.GetInt(constants.IMDDISTRICT + "." + constants.IMDmspl2)
	icidNo2 := config.GetInt(constants.IMDDISTRICT + "." + constants.IMDicid2)
	rainfallNo2 := config.GetInt(constants.IMDDISTRICT + "." + constants.IMDrainfall2)
	maxtempNo2 := config.GetInt(constants.IMDDISTRICT + "." + constants.IMDmaxtemp2)
	mintempNo2 := config.GetInt(constants.IMDDISTRICT + "." + constants.IMDmintemp2)
	maxrelhumNo2 := config.GetInt(constants.IMDDISTRICT + "." + constants.IMDmaxrelhum2)
	minremhumNo2 := config.GetInt(constants.IMDDISTRICT + "." + constants.IMDminremhum2)
	windspeedNo2 := config.GetInt(constants.IMDDISTRICT + "." + constants.IMDwindspeed2)
	winddirectionNo2 := config.GetInt(constants.IMDDISTRICT + "." + constants.IMDwinddirection2)
	districtNo1 := config.GetInt(constants.IMDDISTRICT + "." + constants.IMDdistrict1)
	latNo1 := config.GetInt(constants.IMDDISTRICT + "." + constants.IMDlat1)
	longNo1 := config.GetInt(constants.IMDDISTRICT + "." + constants.IMDlong1)
	altoNo1 := config.GetInt(constants.IMDDISTRICT + "." + constants.IMDalto1)
	pcodNo1 := config.GetInt(constants.IMDDISTRICT + "." + constants.IMDpcod1)
	dayNo1 := config.GetInt(constants.IMDDISTRICT + "." + constants.IMDday1)
	monthNo1 := config.GetInt(constants.IMDDISTRICT + "." + constants.IMDmonth1)
	yearNo1 := config.GetInt(constants.IMDDISTRICT + "." + constants.IMDyear1)
	msplNo1 := config.GetInt(constants.IMDDISTRICT + "." + constants.IMDmspl1)
	icidNo1 := config.GetInt(constants.IMDDISTRICT + "." + constants.IMDicid1)
	rainfallNo1 := config.GetInt(constants.IMDDISTRICT + "." + constants.IMDrainfall1)
	maxtempNo1 := config.GetInt(constants.IMDDISTRICT + "." + constants.IMDmaxtemp1)
	mintempNo1 := config.GetInt(constants.IMDDISTRICT + "." + constants.IMDmintemp1)
	maxrelhumNo1 := config.GetInt(constants.IMDDISTRICT + "." + constants.IMDmaxrelhum1)
	minremhumNo1 := config.GetInt(constants.IMDDISTRICT + "." + constants.IMDminremhum1)
	windspeedNo1 := config.GetInt(constants.IMDDISTRICT + "." + constants.IMDwindspeed1)
	winddirectionNo1 := config.GetInt(constants.IMDDISTRICT + "." + constants.IMDwinddirection1)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		fmt.Println("words===>", len(lines))
		if len(lines) > 0 {
			dwd := models.DistrictWeatherData{}
			var err error
			// state := lines[stateNo:state1No]
			// t := strings.TrimSpace(state)
			// dwd.State = t
			name := lines[districtNo1:districtNo2]
			districtName := strings.TrimSpace(name)
			refstate, err := s.Daos.GetSingleState(ctx, state)
			if err != nil {
				return err
			}
			dwd.State = refstate.ID
			dwd.Name = districtName
			latitude := lines[latNo1:latNo2]
			latStr := strings.TrimSpace(latitude)
			longitude := lines[longNo1:longNo2]
			longStr := strings.TrimSpace(longitude)
			if latStr != "" && latStr != "" {
				lat, err := strconv.ParseFloat(latStr, 64)
				if err != nil {
					log.Println("Error in geting latStr", err.Error())
					err = nil
				}
				long, err := strconv.ParseFloat(longStr, 64)
				if err != nil {
					log.Println("Error in geting latStr", err.Error())
					err = nil
				}
				dwd.Location = models.Location{
					Type:        "point",
					Coordinates: []float64{long, lat},
				}
			}
			alto1 := lines[altoNo1:altoNo2]
			alto := strings.TrimSpace(alto1)
			if alto != "" {
				dwd.Alto, err = strconv.ParseFloat(alto, 64)
				if err != nil {
					log.Println("Error in geting alto", err.Error())
					err = nil
				}
			}
			Pcod := lines[pcodNo1:pcodNo2]
			dwd.Pcod = strings.TrimSpace(Pcod)

			years := lines[yearNo1:yearNo2]
			yearStr := strings.TrimSpace(years)

			months := lines[monthNo1:monthNo2]
			monthStr := strings.TrimSpace(months)

			days := lines[dayNo1:dayNo2]
			dayStr := strings.TrimSpace(days)
			if yearStr != "" {
				year, err := strconv.ParseInt(yearStr[0:4], 10, 0)
				if err != nil {
					log.Println("Error in geting year", err.Error())
					err = nil
				}
				month, err := strconv.ParseInt(monthStr, 10, 0)
				if err != nil {
					log.Println("Error in geting year", err.Error())
					err = nil
				}
				day, err := strconv.ParseInt(dayStr, 10, 0)
				if err != nil {
					log.Println("Error in geting year", err.Error())
					err = nil
				}
				//	loc, _ := time.LoadLocation("Asia/Kolkata")
				loc, _ := time.LoadLocation("Asia/Kolkata")

				t := time.Date(int(year), time.Month(int(month)), int(day), 0, 0, 0, 0, loc)
				fmt.Println("time=====>", t)
				dwd.Date = t.In(loc)

			}
			mspls := lines[msplNo1:msplNo2]
			mspl := strings.TrimSpace(mspls)

			if mspl != "" {
				dwd.Mspl, err = strconv.ParseFloat(mspl, 64)
				if err != nil {
					log.Println("Error in geting mspl", err.Error())
					err = nil
				}
			}
			icids := lines[icidNo1:icidNo2]
			icid := strings.TrimSpace(icids)

			if icid != "" {
				dwd.Icid, err = strconv.ParseFloat(icid, 64)
				if err != nil {
					log.Println("Error in geting icid", err.Error())
					err = nil
				}
			}
			rainfalls := lines[rainfallNo1:rainfallNo2]
			rainfall := strings.TrimSpace(rainfalls)
			if rainfall != "" {
				dwd.WeatherData.Rain, err = strconv.ParseFloat(rainfall, 64)
				if err != nil {
					log.Println("Error in geting rainfall", err.Error())
					err = nil
				}
			}
			maximumtemp := lines[maxtempNo1:maxtempNo2]
			maxtemp := strings.TrimSpace(maximumtemp)

			if maxtemp != "" {
				dwd.WeatherData.Temp.Max, err = strconv.ParseFloat(maxtemp, 64)
				if err != nil {
					log.Println("Error in geting maxtemp", err.Error())
					err = nil
				}
			}
			minimumtemp := lines[mintempNo1:mintempNo2]
			mintemp := strings.TrimSpace(minimumtemp)

			if mintemp != "" {
				dwd.WeatherData.Temp.Min, err = strconv.ParseFloat(mintemp, 64)
				if err != nil {
					log.Println("Error in geting mintemp", err.Error())
					err = nil
				}
			}
			maximumrelhum := lines[maxrelhumNo1:maxrelhumNo2]
			maxrelhum := strings.TrimSpace(maximumrelhum)

			if maxrelhum != "" {
				dwd.WeatherData.HumidityMax, err = strconv.ParseFloat(maxrelhum, 64)
				if err != nil {
					log.Println("Error in geting maxrelhum", err.Error())
					err = nil
				}
			}
			minimumremhum := lines[minremhumNo1:minremhumNo2]
			minremhum := strings.TrimSpace(minimumremhum)

			if minremhum != "" {
				dwd.WeatherData.HumidityMin, err = strconv.ParseFloat(minremhum, 64)
				if err != nil {
					log.Println("Error in geting minremhum", err.Error())
					err = nil
				}
			}
			windspeeds := lines[windspeedNo1:windspeedNo2]
			windspeed := strings.TrimSpace(windspeeds)

			if windspeed != "" {
				dwd.WeatherData.Windspeed, err = strconv.ParseFloat(windspeed, 64)
				if err != nil {
					log.Println("Error in geting windspeed", err.Error())
					err = nil
				}
			}
			wind := lines[winddirectionNo1:winddirectionNo2]
			winddirection := strings.TrimSpace(wind)

			if winddirection != "" {
				dwd.WeatherData.Winddeg, err = strconv.ParseFloat(winddirection, 64)
				if err != nil {
					log.Println("Error in geting winddirection", err.Error())
					err = nil
				}
			}
			dwd.UniqueID = fmt.Sprintf("%v%v%v", dwd.Date.Day(), dwd.Date.Month().String(), dwd.Date.Year())
			dberr := s.Daos.SaveDistrictWeatherDataNameWithUpsert(ctx, &dwd)
			if dberr != nil {

				return errors.New("Db Error" + dberr.Error())

			}

		} else {
			for k, v := range lines {
				fmt.Println(k+1, v)
			}
			log.Println("words not in range", lines)
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
func (s *Service) GetBlockWeatherDataWithImdWithStates(ctx *models.Context, ImdName string, ImdFile string) ([]string, error) {
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
	IMDWeatherDataUrl := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.IMD_BLOCKWEATHERDATA_URL)
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

	IMDURL := fmt.Sprintf("%v%v/%v%v", IMDWeatherDataUrl, ImdName, ImdFile, date)

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

	//	fmt.Println("lines===>", len(lines))
	return lines, nil
}

func (s *Service) GetBlockWeatherDataWithImdWithState(ctx *models.Context, ImdName string, ImdFile string) ([]string, error) {
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
	//date = fmt.Sprintf("%v%v0%v", year, month, 2)
	imdurl := "DIST_BLOCK_FT1534/2020/bfcst/"
	param1 := "madhya-pradesh"
	file := "/bfcst2000z"
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

func (s *Service) GetDistrictWeatherDataByDistrictId(ctx *models.Context, UniqueID string) ([]models.RefDistrictWeatherData, error) {
	districtweatherdata, err := s.Daos.GetDistrictWeatherDataByDistrictId(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return districtweatherdata, nil
}
