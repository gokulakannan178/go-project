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

//SaveBlockWeatherData :""
func (s *Service) SaveBlockWeatherData(ctx *models.Context, blockweatherdata *models.BlockWeatherData) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	//organisation.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONORGANISATION)

	blockweatherdata.Status = constants.BLOCKWEATHERDATASTATUSACTIVE
	blockweatherdata.ActiveStatus = true
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 Blockweatherdata.created")
	blockweatherdata.Created = &created
	log.Println("b4 Blockweatherdata.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveBlockWeatherData(ctx, blockweatherdata)
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

//UpdateBlockWeatherData : ""
func (s *Service) UpdateBlockWeatherData(ctx *models.Context, blockweatherdata *models.BlockWeatherData) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateBlockWeatherData(ctx, blockweatherdata)
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

//EnableBlockWeatherData : ""
func (s *Service) EnableBlockWeatherData(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableBlockWeatherData(ctx, UniqueID)
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

//DisableBlockWeatherData : ""
func (s *Service) DisableBlockWeatherData(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableBlockWeatherData(ctx, UniqueID)
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

//DeleteBlockWeatherData : ""
func (s *Service) DeleteBlockWeatherData(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteBlockWeatherData(ctx, UniqueID)
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
func (s *Service) GetSingleBlockWeatherData(ctx *models.Context, UniqueID string) (*models.RefBlockWeatherData, error) {
	Blockweatherdata, err := s.Daos.GetSingleBlockWeatherData(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return Blockweatherdata, nil
}

//FilterBlockWeatherData :""
func (s *Service) FilterBlockWeatherData(ctx *models.Context, blockweatherdatafilter *models.BlockWeatherDataFilter, pagination *models.Pagination) (Blockweatherdata []models.RefBlockWeatherData, err error) {
	return s.Daos.FilterBlockWeatherData(ctx, blockweatherdatafilter, pagination)
}
func (s *Service) SaveBlockWeatherDataWithOpenWebsite(ctx *models.Context, lat string, lon string) error {
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
			blockweatherdata := new(models.BlockWeatherData)
			blockweatherdata.Status = constants.DISTRICWEATHERDATASTATUSACTIVE
			blockweatherdata.ActiveStatus = true
			t := time.Now()
			created := models.Created{}
			created.On = &t
			created.By = constants.SYSTEM
			blockweatherdata.CreatedDate = &t
			blockweatherdata.Created = &created
			//blockweatherdata.WeatherData = v
			blockweatherdata.Date = time.Unix(int64(v.Dt), 0)

			dberr := s.Daos.SaveBlockWeatherData(ctx, blockweatherdata)
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
func (s *Service) SaveBlockWeatherDataWithBlock(ctx *models.Context, Block *models.Block) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	var lat string
	var long string
	//organisation.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONORGANISATION)
	if len(Block.Location.Coordinates) > 0 {
		lat = fmt.Sprintf("%v", Block.Location.Coordinates[1])
		long = fmt.Sprintf("%v", Block.Location.Coordinates[0])
	} else {
		log.Println("pls add a location latitude and longitude---" + Block.Name + "")
		log.Println("pls add a location latitude and longitude---" + Block.Name + "")

	}
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		if lat != "" && long != "" {
			weatherDataMaster, err := s.GetWeatherData(ctx, lat, long)
			if err != nil {
				return errors.New("weather data not found")
			}
			for _, v := range weatherDataMaster.Daily {
				blockweatherdata := new(models.BlockWeatherData)
				blockweatherdata.Status = constants.DISTRICWEATHERDATASTATUSACTIVE
				blockweatherdata.ActiveStatus = true
				//Blockweatherdata.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONBlockWEATHERDATA)
				t := time.Now()
				created := models.Created{}
				created.On = &t
				//	t.Month().String()
				created.By = constants.SYSTEM
				blockweatherdata.CreatedDate = &t
				blockweatherdata.Created = &created
				//blockweatherdata.WeatherData = weatherDataMaster.Daily[k]
				blockweatherdata.Date = time.Unix(int64(v.Dt), 0)
				blockweatherdata.Block = Block.ID
				blockweatherdata.Name = Block.Name
				blockweatherdata.UniqueID = fmt.Sprintf("%v_%v_%v", blockweatherdata.Date.Day(), blockweatherdata.Date.Month().String(), blockweatherdata.Date.Year())
				//	Blockweatherdata.WeatherData.Temp.Min
				dberr := s.Daos.SaveBlockWeatherDataWithUpsert(ctx, blockweatherdata)
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
func (s *Service) GetSingleBlockWeatherDataWithCurrentDate(ctx *models.Context, UniqueID string) (*models.RefBlockWeatherData, error) {
	Blockweatherdata, err := s.Daos.GetSingleBlockWeatherDataWithCurrentDate(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return Blockweatherdata, nil
}
func (s *Service) SaveBlockWeatherDataCron() {
	c := context.TODO()
	ctx := app.GetApp(c, s.Daos)
	defer ctx.Client.Disconnect(c)
	//Dissemination := new(models.Dissemination)
	Blocks, err := s.Daos.GetActiveBlock(ctx)
	if err != nil {
		log.Println("dissemination not found" + err.Error())
	}
	for _, v := range Blocks {
		err := s.SaveBlockWeatherDataWithBlock(ctx, &v)
		if err != nil {
			log.Println("not save Weather data this Block---" + v.Name + "" + err.Error())
			continue
		}
	}
}

func (s *Service) LoadBlockWeatherReport(ctx *models.Context) {
	client, err := ftp.Dial("103.215.208.49:21")
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := client.Login("anonymous", "anonymous"); err != nil {
		fmt.Println(err)
		return
	}
	r, err := client.Retr("/pub/DIST_BLOCK_FT1534/2020/dfcst/andhra-pradesh/dfcst200z200101")
	if err != nil {
		panic(err)
	}
	defer r.Close()

	buf, err := ioutil.ReadAll(r)
	// println(string(buf))
	//var lbwr []models.LoadBlockWeatherReport
	lines := strings.Split(string(buf), "\n")

	for _, v := range lines {
		words := strings.Fields(v)
		fmt.Println(words, len(words))
	}
	//	err = s.Daos.LoadBlockWeatherReport(ctx)
	if err != nil {
		return
	}
	//	fmt.Println("weather data===>", r)
	if err := client.Quit(); err != nil {
		log.Fatal(err)
	}
}

// func (s *Service) SaveBlockWeatherDataWithImd(ctx *models.Context, lines []string) error {
// 	log.Println("transaction start")
// 	//Start Transaction
// 	if err := ctx.Session.StartTransaction(); err != nil {
// 		return err
// 	}
// 	defer ctx.Session.EndSession(ctx.CTX)
// 	config := config.NewConfig("blockimd", "config")
//	maxLength := config.GetInt(constants.IMDDISTRICT + "." + constants.IMDmaxlength)
// stateNo := config.GetInt(constants.IMDDISTRICT + "." + constants.IMDstate)
// 	blockNo := config.GetInt(constants.IMDBLOCK + "." + constants.IMDBlock)
// 	latNo := config.GetInt(constants.IMDBLOCK + "." + constants.IMDlat)
// 	longNo := config.GetInt(constants.IMDBLOCK + "." + constants.IMDlong)
// 	altoNo := config.GetInt(constants.IMDBLOCK + "." + constants.IMDalto)
// 	pcodNo := config.GetInt(constants.IMDBLOCK + "." + constants.IMDpcod)
// 	dayNo := config.GetInt(constants.IMDBLOCK + "." + constants.IMDday)
// 	monthNo := config.GetInt(constants.IMDBLOCK + "." + constants.IMDmonth)
// 	yearNo := config.GetInt(constants.IMDBLOCK + "." + constants.IMDyear)
// 	msplNo := config.GetInt(constants.IMDBLOCK + "." + constants.IMDmspl)
// 	icidNo := config.GetInt(constants.IMDBLOCK + "." + constants.IMDicid)
// 	rainfallNo := config.GetInt(constants.IMDBLOCK + "." + constants.IMDrainfall)
// 	maxtempNo := config.GetInt(constants.IMDBLOCK + "." + constants.IMDmaxtemp)
// 	mintempNo := config.GetInt(constants.IMDBLOCK + "." + constants.IMDmintemp)
// 	maxrelhumNo := config.GetInt(constants.IMDBLOCK + "." + constants.IMDmaxrelhum)
// 	minremhumNo := config.GetInt(constants.IMDBLOCK + "." + constants.IMDminremhum)
// 	windspeedNo := config.GetInt(constants.IMDBLOCK + "." + constants.IMDwindspeed)
// 	winddirectionNo := config.GetInt(constants.IMDBLOCK + "." + constants.IMDwinddirection)
// 	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
// 		if len(lines) > 0 {
// 			for k, _ := range lines {
// 				words := strings.Fields(lines[k])
// 				dwd := models.BlockWeatherData{}
// 				// state := words[stateNo]
// 				// district := words[districtNo]
// 				// sfmt.Println("words[longNo]===", words[longNo])
// 				latStr := words[longNo]
// 				lat, err := strconv.ParseFloat(latStr, 64)
// 				if err != nil {
// 					log.Println("Error in geting latStr", err.Error())
// 					err = nil
// 				}
// 				longStr := words[latNo]
// 				long, err := strconv.ParseFloat(longStr, 64)
// 				if err != nil {
// 					log.Println("Error in geting longStr", err.Error())
// 					err = nil
// 				}
// 				dwd.Name = words[blockNo]
// 				dwd.Location = models.Location{
// 					Type:        "point",
// 					Coordinates: []float64{long, lat},
// 				}
// 				alto := words[altoNo]
// 				dwd.Alto, err = strconv.ParseFloat(alto, 64)
// 				if err != nil {
// 					log.Println("Error in geting alto", err.Error())
// 					err = nil
// 				}
// 				dwd.Pcod = words[pcodNo]
// 				yearStr := words[yearNo]
// 				year, err := strconv.ParseInt(yearStr[0:4], 10, 0)
// 				if err != nil {
// 					log.Println("Error in geting year", err.Error())
// 					err = nil
// 				}
// 				monthStr := words[monthNo]
// 				month, err := strconv.ParseInt(monthStr, 10, 0)
// 				if err != nil {
// 					log.Println("Error in geting year", err.Error())
// 					err = nil
// 				}
// 				dayStr := words[dayNo]
// 				day, err := strconv.ParseInt(dayStr, 10, 0)
// 				if err != nil {
// 					log.Println("Error in geting year", err.Error())
// 					err = nil
// 				}
// 				loc, _ := time.LoadLocation("Asia/Kolkata")
// 				t := time.Date(int(year), time.Month(int(month)), int(day), 0, 0, 0, 0, loc)
// 				dwd.Date = t
// 				mspl := words[msplNo]
// 				dwd.Mspl, err = strconv.ParseFloat(mspl, 64)
// 				if err != nil {
// 					log.Println("Error in geting mspl", err.Error())
// 					err = nil
// 				}
// 				icid := words[icidNo]
// 				dwd.Icid, err = strconv.ParseFloat(icid, 64)
// 				if err != nil {
// 					log.Println("Error in geting icid", err.Error())
// 					err = nil
// 				}
// 				rainfall := words[rainfallNo]
// 				dwd.WeatherData.Rain, err = strconv.ParseFloat(rainfall, 64)
// 				if err != nil {
// 					log.Println("Error in geting rainfall", err.Error())
// 					err = nil
// 				}
// 				maxtemp := words[maxtempNo]
// 				dwd.WeatherData.Temp.Max, err = strconv.ParseFloat(maxtemp, 64)
// 				if err != nil {
// 					log.Println("Error in geting maxtemp", err.Error())
// 					err = nil
// 				}
// 				mintemp := words[mintempNo]
// 				dwd.WeatherData.Temp.Min, err = strconv.ParseFloat(mintemp, 64)
// 				if err != nil {
// 					log.Println("Error in geting mintemp", err.Error())
// 					err = nil
// 				}
// 				maxrelhum := words[maxrelhumNo]
// 				dwd.WeatherData.HumidityMax, err = strconv.ParseFloat(maxrelhum, 64)
// 				if err != nil {
// 					log.Println("Error in geting maxrelhum", err.Error())
// 					err = nil
// 				}
// 				minremhum := words[minremhumNo]
// 				dwd.WeatherData.HumidityMin, err = strconv.ParseFloat(minremhum, 64)
// 				if err != nil {
// 					log.Println("Error in geting minremhum", err.Error())
// 					err = nil
// 				}
// 				windspeed := words[windspeedNo]
// 				dwd.WeatherData.Windspeed, err = strconv.ParseFloat(windspeed, 64)
// 				if err != nil {
// 					log.Println("Error in geting windspeed", err.Error())
// 					err = nil
// 				}
// 				winddirection := words[winddirectionNo]
// 				dwd.WeatherData.Winddeg, err = strconv.ParseFloat(winddirection, 64)
// 				if err != nil {
// 					log.Println("Error in geting winddirection", err.Error())
// 					err = nil
// 				}
// 				dberr := s.Daos.SaveBlockWeatherData(ctx, &dwd)
// 				if dberr != nil {
// 					return errors.New("Db Error" + dberr.Error())
// 				}
// 			}
// 		}
// 		if err := ctx.Session.CommitTransaction(sc); err != nil {
// 			return errors.New("Not able to commit - " + err.Error())
// 		}
// 		return nil
// 	}); err != nil {
// 		log.Println("Transaction start aborting")
// 		if abortError := ctx.Session.AbortTransaction(ctx.CTX); abortError != nil {
// 			return errors.New("Error while aborting transaction" + abortError.Error())
// 		}
// 		log.Println("Transaction aborting completed successfully")
// 		return err
// 	}
// 	return nil
// }

func (s *Service) SaveBlockWeatherDataWithImdWithState(ctx *models.Context, lines string, state string) error {
	//log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	//districtRefMap := make(map[string]primitive.ObjectID)
	config := config.NewConfig("blockimd", "config")
	//maxLength := config.GetInt(constants.IMDBLOCK + "." + constants.IMDmaxlength)

	// stateNo := config.GetInt(constants.IMDDISTRICT + "." + constants.IMDstate)
	districtNo2 := config.GetInt(constants.IMDBLOCK + "." + constants.IMDdistrict2)
	blockNo2 := config.GetInt(constants.IMDBLOCK + "." + constants.IMDBlock2)
	blockNo1 := config.GetInt(constants.IMDBLOCK + "." + constants.IMDBlock1)
	latNo2 := config.GetInt(constants.IMDBLOCK + "." + constants.IMDlat2)
	longNo2 := config.GetInt(constants.IMDBLOCK + "." + constants.IMDlong2)
	altoNo2 := config.GetInt(constants.IMDBLOCK + "." + constants.IMDalto2)
	pcodNo2 := config.GetInt(constants.IMDBLOCK + "." + constants.IMDpcod2)
	dayNo2 := config.GetInt(constants.IMDBLOCK + "." + constants.IMDday2)
	monthNo2 := config.GetInt(constants.IMDBLOCK + "." + constants.IMDmonth2)
	yearNo2 := config.GetInt(constants.IMDBLOCK + "." + constants.IMDyear2)
	msplNo2 := config.GetInt(constants.IMDBLOCK + "." + constants.IMDmspl2)
	icidNo2 := config.GetInt(constants.IMDBLOCK + "." + constants.IMDicid2)
	rainfallNo2 := config.GetInt(constants.IMDBLOCK + "." + constants.IMDrainfall2)
	maxtempNo2 := config.GetInt(constants.IMDBLOCK + "." + constants.IMDmaxtemp2)
	mintempNo2 := config.GetInt(constants.IMDBLOCK + "." + constants.IMDmintemp2)
	maxrelhumNo2 := config.GetInt(constants.IMDBLOCK + "." + constants.IMDmaxrelhum2)
	minremhumNo2 := config.GetInt(constants.IMDBLOCK + "." + constants.IMDminremhum2)
	windspeedNo2 := config.GetInt(constants.IMDBLOCK + "." + constants.IMDwindspeed2)
	winddirectionNo2 := config.GetInt(constants.IMDBLOCK + "." + constants.IMDwinddirection2)
	districtNo1 := config.GetInt(constants.IMDBLOCK + "." + constants.IMDdistrict1)
	latNo1 := config.GetInt(constants.IMDBLOCK + "." + constants.IMDlat1)
	longNo1 := config.GetInt(constants.IMDBLOCK + "." + constants.IMDlong1)
	altoNo1 := config.GetInt(constants.IMDBLOCK + "." + constants.IMDalto1)
	pcodNo1 := config.GetInt(constants.IMDBLOCK + "." + constants.IMDpcod1)
	dayNo1 := config.GetInt(constants.IMDBLOCK + "." + constants.IMDday1)
	monthNo1 := config.GetInt(constants.IMDBLOCK + "." + constants.IMDmonth1)
	yearNo1 := config.GetInt(constants.IMDBLOCK + "." + constants.IMDyear1)
	msplNo1 := config.GetInt(constants.IMDBLOCK + "." + constants.IMDmspl1)
	icidNo1 := config.GetInt(constants.IMDBLOCK + "." + constants.IMDicid1)
	rainfallNo1 := config.GetInt(constants.IMDBLOCK + "." + constants.IMDrainfall1)
	maxtempNo1 := config.GetInt(constants.IMDBLOCK + "." + constants.IMDmaxtemp1)
	mintempNo1 := config.GetInt(constants.IMDBLOCK + "." + constants.IMDmintemp1)
	maxrelhumNo1 := config.GetInt(constants.IMDBLOCK + "." + constants.IMDmaxrelhum1)
	minremhumNo1 := config.GetInt(constants.IMDBLOCK + "." + constants.IMDminremhum1)
	windspeedNo1 := config.GetInt(constants.IMDBLOCK + "." + constants.IMDwindspeed1)
	winddirectionNo1 := config.GetInt(constants.IMDBLOCK + "." + constants.IMDwinddirection1)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		if len(lines) > 0 {
			bwd := models.BlockWeatherData{}
			var err error

			name := lines[blockNo1:blockNo2]
			blockName := strings.TrimSpace(name)
			refstate, err := s.Daos.GetSingleState(ctx, state)
			if err != nil {
				return err
			}

			district := lines[districtNo1:districtNo2]
			districtName := strings.TrimSpace(district)
			refdistrct, err := s.Daos.GetSingleDistrict(ctx, districtName)
			if err != nil {
				return err
			}
			// districtid, err := primitive.ObjectIDFromHex(districtName)
			// if err != nil {
			// 	return err
			// }
			bwd.State = refstate.ID
			bwd.District = refdistrct.ID
			bwd.Name = blockName
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
				bwd.Location = models.Location{
					Type:        "point",
					Coordinates: []float64{long, lat},
				}
			}
			alto1 := lines[altoNo1:altoNo2]
			alto := strings.TrimSpace(alto1)
			if alto != "" {
				bwd.Alto, err = strconv.ParseFloat(alto, 64)
				if err != nil {
					log.Println("Error in geting alto", err.Error())
					err = nil
				}
			}
			Pcod := lines[pcodNo1:pcodNo2]
			bwd.Pcod = strings.TrimSpace(Pcod)

			year := lines[yearNo1:yearNo2]
			yearStr := strings.TrimSpace(year)

			month := lines[monthNo1:monthNo2]
			monthStr := strings.TrimSpace(month)

			day := lines[dayNo1:dayNo2]
			dayStr := strings.TrimSpace(day)
			if yearStr != "" {
				year, err := strconv.ParseInt(yearStr[0:4], 10, 0)
				if err != nil {
					log.Println("Error in geting year", err.Error())
					err = nil
				}
				month, err := strconv.ParseInt(monthStr, 10, 0)
				if err != nil {
					log.Println("Error in geting month", err.Error())
					err = nil
				}
				day, err := strconv.ParseInt(dayStr, 10, 0)
				if err != nil {
					log.Println("Error in geting day", err.Error())
					err = nil
				}
				//	loc, _ := time.LoadLocation("Asia/Kolkata")
				loc, _ := time.LoadLocation("Asia/Kolkata")

				t := time.Date(int(year), time.Month(int(month)), int(day), 0, 0, 0, 0, loc)
				fmt.Println("time=====>", t)
				bwd.Date = t.In(loc)

			}
			mspls := lines[msplNo1:msplNo2]
			mspl := strings.TrimSpace(mspls)

			if mspl != "" {
				bwd.Mspl, err = strconv.ParseFloat(mspl, 64)
				if err != nil {
					log.Println("Error in geting mspl", err.Error())
					err = nil
				}
			}
			icids := lines[icidNo1:icidNo2]
			icid := strings.TrimSpace(icids)

			if icid != "" {
				bwd.Icid, err = strconv.ParseFloat(icid, 64)
				if err != nil {
					log.Println("Error in geting icid", err.Error())
					err = nil
				}
			}
			rainfalls := lines[rainfallNo1:rainfallNo2]
			rainfall := strings.TrimSpace(rainfalls)
			if rainfall != "" {
				bwd.WeatherData.Rain, err = strconv.ParseFloat(rainfall, 64)
				if err != nil {
					log.Println("Error in geting rainfall", err.Error())
					err = nil
				}
			}
			maximumtemp := lines[maxtempNo1:maxtempNo2]
			maxtemp := strings.TrimSpace(maximumtemp)

			if maxtemp != "" {
				bwd.WeatherData.Temp.Max, err = strconv.ParseFloat(maxtemp, 64)
				if err != nil {
					log.Println("Error in geting maxtemp", err.Error())
					err = nil
				}
			}
			minimumtemp := lines[mintempNo1:mintempNo2]
			mintemp := strings.TrimSpace(minimumtemp)

			if mintemp != "" {
				bwd.WeatherData.Temp.Min, err = strconv.ParseFloat(mintemp, 64)
				if err != nil {
					log.Println("Error in geting mintemp", err.Error())
					err = nil
				}
			}
			maximumrelhum := lines[maxrelhumNo1:maxrelhumNo2]
			maxrelhum := strings.TrimSpace(maximumrelhum)

			if maxrelhum != "" {
				bwd.WeatherData.HumidityMax, err = strconv.ParseFloat(maxrelhum, 64)
				if err != nil {
					log.Println("Error in geting maxrelhum", err.Error())
					err = nil
				}
			}
			minimumremhum := lines[minremhumNo1:minremhumNo2]
			minremhum := strings.TrimSpace(minimumremhum)

			if minremhum != "" {
				bwd.WeatherData.HumidityMin, err = strconv.ParseFloat(minremhum, 64)
				if err != nil {
					log.Println("Error in geting minremhum", err.Error())
					err = nil
				}
			}
			windspeeds := lines[windspeedNo1:windspeedNo2]
			windspeed := strings.TrimSpace(windspeeds)

			if windspeed != "" {
				bwd.WeatherData.Windspeed, err = strconv.ParseFloat(windspeed, 64)
				if err != nil {
					log.Println("Error in geting windspeed", err.Error())
					err = nil
				}
			}
			wind := lines[winddirectionNo1:winddirectionNo2]
			winddirection := strings.TrimSpace(wind)

			if winddirection != "" {
				bwd.WeatherData.Winddeg, err = strconv.ParseFloat(winddirection, 64)
				if err != nil {
					log.Println("Error in geting winddirection", err.Error())
					err = nil
				}
			}
			bwd.UniqueID = fmt.Sprintf("%v%v%v", bwd.Date.Day(), bwd.Date.Month().String(), bwd.Date.Year())
			dberr := s.Daos.SaveBlockWeatherDataNameWithUpsert(ctx, &bwd)
			if dberr != nil {

				return errors.New("Db Error" + dberr.Error())

			}

		} else {
			log.Println("words not in range", len(lines))
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

func (s *Service) GetBlockWeatherDataByBlockId(ctx *models.Context, UniqueID string) ([]models.RefBlockWeatherData, error) {
	blockweatherdata, err := s.Daos.GetBlockWeatherDataByBlockId(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return blockweatherdata, nil
}
