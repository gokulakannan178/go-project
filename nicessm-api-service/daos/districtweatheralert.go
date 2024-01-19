package daos

import (
	"context"
	"errors"
	"fmt"
	"log"
	"nicessm-api-service/constants"
	"nicessm-api-service/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//SaveDisease :""
func (d *Daos) SaveDistrictWeatherAlert(ctx *models.Context, DistrictWeatherAlert *models.DistrictWeatherAlert) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONDISTRICTWEATHERALERT).InsertOne(ctx.CTX, DistrictWeatherAlert)
	if err != nil {
		return err
	}
	DistrictWeatherAlert.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}
func (d *Daos) SaveDistrictWeatherAlertWithUpsert(ctx *models.Context, DistrictWeatherAlert *models.DistrictWeatherAlert) error {

	opts := options.Update().SetUpsert(true)
	updateQuery := bson.M{"uniqueId": DistrictWeatherAlert.UniqueID, "district._id": DistrictWeatherAlert.District.ID, "parameter._id": DistrictWeatherAlert.ParameterId.ID, "month._id": DistrictWeatherAlert.Month.ID, "severityType._id": DistrictWeatherAlert.SeverityType.ID}
	updateData := bson.M{"$set": DistrictWeatherAlert}
	_, err := ctx.DB.Collection(constants.COLLECTIONDISTRICTWEATHERALERT).UpdateOne(ctx.CTX, updateQuery, updateData, opts)
	return err
}
func (d *Daos) SaveDistrictWeatherAlertTempWithUpsert(ctx *models.Context, DistrictWeatherAlert *models.DistrictWeatherAlert) error {

	opts := options.Update().SetUpsert(true)
	updateQuery := bson.M{"uniqueId": DistrictWeatherAlert.UniqueID, "district._id": DistrictWeatherAlert.District.ID, "parameter._id": DistrictWeatherAlert.ParameterId.ID, "month._id": DistrictWeatherAlert.Month.ID, "severityType._id": DistrictWeatherAlert.SeverityType.ID, "tittle": DistrictWeatherAlert.Tittle}
	updateData := bson.M{"$set": DistrictWeatherAlert}
	_, err := ctx.DB.Collection(constants.COLLECTIONDISTRICTWEATHERALERT).UpdateOne(ctx.CTX, updateQuery, updateData, opts)
	return err
}

//GetSingleDisease : ""
func (d *Daos) GetSingleDistrictWeatherAlert(ctx *models.Context, UniqueID string) (*models.RefDistrictWeatherAlert, error) {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})
	//lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "district.state", "_id", "ref.state", "ref.state")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDISTRICTWEATHERALERT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var DistrictWeatherAlerts []models.RefDistrictWeatherAlert
	var DistrictWeatherAlert *models.RefDistrictWeatherAlert
	if err = cursor.All(ctx.CTX, &DistrictWeatherAlerts); err != nil {
		return nil, err
	}
	if len(DistrictWeatherAlerts) > 0 {
		DistrictWeatherAlert = &DistrictWeatherAlerts[0]
	}
	return DistrictWeatherAlert, nil
}

//UpdateDistrictWeatherAlert : ""
func (d *Daos) UpdateDistrictWeatherAlert(ctx *models.Context, DistrictWeatherAlert *models.DistrictWeatherAlert) error {

	selector := bson.M{"_id": DistrictWeatherAlert.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": DistrictWeatherAlert}
	_, err := ctx.DB.Collection(constants.COLLECTIONDISTRICTWEATHERALERT).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}
func (d *Daos) UpdateDistrictWeatherAlertMasterwithWeatheralert(ctx *models.Context, DistrictWeatherAlert *models.UpdateDistrictWeatherAlert) error {
	selector := bson.M{"weatherDataAlert._id": DistrictWeatherAlert.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": bson.M{"weatherDataAlert.isSms": DistrictWeatherAlert.IsSms, "weatherDataAlert.isNicessm": DistrictWeatherAlert.IsNicessm, "weatherDataAlert.isWhatsApp": DistrictWeatherAlert.IsWhatsApp, "weatherDataAlert.isTelegram": DistrictWeatherAlert.IsTelegram, "weatherDataAlert.WeatherAlert": DistrictWeatherAlert.WeatherAlert}}
	fmt.Println("selector==>", selector)
	fmt.Println("updateInterface==>", updateInterface)
	_, err := ctx.DB.Collection(constants.COLLECTIONDISTRICTWEATHERALERT).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}
func (d *Daos) IsUpdateDistrictWeatherAlertMaster(ctx *models.Context, DistrictWeatherAlert *models.UpdateDistrictWeatherAlert) error {
	selector := bson.M{"_id": DistrictWeatherAlert.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": bson.M{"isSms": DistrictWeatherAlert.IsSms, "isNicessm": DistrictWeatherAlert.IsNicessm, "isWhatsApp": DistrictWeatherAlert.IsWhatsApp, "isTelegram": DistrictWeatherAlert.IsTelegram, "WeatherAlert": DistrictWeatherAlert.WeatherAlert}}

	_, err := ctx.DB.Collection(constants.COLLECTIONDISTRICTWEATHERALERTMASTER).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterDistrictWeatherAlert : ""
func (d *Daos) FilterDistrictWeatherAlert(ctx *models.Context, DistrictWeatherAlertfilter *models.DistrictWeatherAlertFilter, pagination *models.Pagination) ([]models.RefDistrictWeatherAlert, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if DistrictWeatherAlertfilter != nil {

		if len(DistrictWeatherAlertfilter.ActiveStatus) > 0 {
			query = append(query, bson.M{"activeStatus": bson.M{"$in": DistrictWeatherAlertfilter.ActiveStatus}})
		}
		if len(DistrictWeatherAlertfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": DistrictWeatherAlertfilter.Status}})
		}
		if len(DistrictWeatherAlertfilter.District) > 0 {
			query = append(query, bson.M{"district._id": bson.M{"$in": DistrictWeatherAlertfilter.District}})
		}
		if len(DistrictWeatherAlertfilter.ParameterId) > 0 {
			query = append(query, bson.M{"parameter._id": bson.M{"$in": DistrictWeatherAlertfilter.ParameterId}})
		}
		if len(DistrictWeatherAlertfilter.SeverityType) > 0 {
			query = append(query, bson.M{"severityType._id": bson.M{"$in": DistrictWeatherAlertfilter.SeverityType}})
		}
		if len(DistrictWeatherAlertfilter.WeatherData) > 0 {
			query = append(query, bson.M{"weatherData._id": bson.M{"$in": DistrictWeatherAlertfilter.WeatherData}})
		}
		if len(DistrictWeatherAlertfilter.WeatherDataAlert) > 0 {
			query = append(query, bson.M{"weatherDataAlert._id": bson.M{"$in": DistrictWeatherAlertfilter.WeatherDataAlert}})
		}
		if len(DistrictWeatherAlertfilter.Month) > 0 {
			query = append(query, bson.M{"month._id": bson.M{"$in": DistrictWeatherAlertfilter.Month}})
		}
		//Regex
		if DistrictWeatherAlertfilter.SearchBox.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: DistrictWeatherAlertfilter.SearchBox.Name, Options: "xi"}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if DistrictWeatherAlertfilter != nil {
		if DistrictWeatherAlertfilter.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{DistrictWeatherAlertfilter.SortBy: DistrictWeatherAlertfilter.SortOrder}})

		}

	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONDISTRICTWEATHERALERT).CountDocuments(ctx.CTX, func() bson.M {
			if query != nil {
				if len(query) > 0 {
					return bson.M{"$and": query}
				}
			}
			return bson.M{}
		}())
		if err != nil {
			log.Println("Error in geting pagination")
		}
		fmt.Println("count", totalCount)
		pagination.Count = int(totalCount)
		d.Shared.PaginationData(pagination)
	}
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "district.state", "_id", "ref.state", "ref.state")...)
	//Aggregation
	d.Shared.BsonToJSONPrintTag("DistrictWeatherAlert query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDISTRICTWEATHERALERT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var DistrictWeatherAlerts []models.RefDistrictWeatherAlert
	if err = cursor.All(context.TODO(), &DistrictWeatherAlerts); err != nil {
		return nil, err
	}
	return DistrictWeatherAlerts, nil
}

//EnableDistrictWeatherAlert :""
func (d *Daos) EnableDistrictWeatherAlert(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.DISTRICTWEATHERALERTSTATUSACTIVE, "activeStatus": true}}
	_, err = ctx.DB.Collection(constants.COLLECTIONDISTRICTWEATHERALERT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableDisease :""
func (d *Daos) DisableDistrictWeatherAlert(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.DISTRICTWEATHERALERTSTATUSDISABLE, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONDISTRICTWEATHERALERT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteDistrictWeatherAlert :""
func (d *Daos) DeleteDistrictWeatherAlert(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.DISTRICTWEATHERALERTSTATUSDELETED, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONDISTRICTWEATHERALERT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
func (d *Daos) DistrictWeatherAlertMaster(ctx *models.Context, Weatherdata *models.DistrictWeatherData) error {
	month, err := d.GetCurrentMonthSeason(ctx, Weatherdata.Date)
	if err != nil {
		return err
	}
	district, err := d.GetSingleDistrict(ctx, Weatherdata.District.Hex())
	if err != nil {
		return err
	}
	fmt.Println("windSpeed==>", Weatherdata.WeatherData.Windspeed)
	fmt.Println("Humidity==>", Weatherdata.WeatherData.Humidity)
	fmt.Println("Winddeg==>", Weatherdata.WeatherData.Winddeg)
	fmt.Println("Rain==>", Weatherdata.WeatherData.Rain)
	fmt.Println("TempMax==>", Weatherdata.WeatherData.Temp.Max)
	fmt.Println("TempMin==>", Weatherdata.WeatherData.Temp.Min)
	//	if Weatherdata.WeatherData.Windspeed > 0 {
	//value := Weatherdata.WeatherData.Windspeed
	fmt.Println("Windspeed===>", Weatherdata.WeatherData.Windspeed)
	//find a Windspeed Weather alert
	paramerterid, err := d.GetSingleWeatherParameterWithName(ctx, constants.WEATHERPARAMETERWINDSPEED)
	if err != nil {
		return err
	}
	fmt.Println("windspeed===>", paramerterid.Name, paramerterid.ID)
	DistrictWeatherAlertWindSpeed, err := d.GetSingleDistrictWeatherAlertMasterWithSpecialIds(ctx, month.ID.Hex(), paramerterid.ID.Hex(), Weatherdata.District.Hex())
	if err != nil {
		return err
	}

	if len(DistrictWeatherAlertWindSpeed) > 0 {
		var ServentType primitive.ObjectID
		var DistrictDistrictWeatherAlertMaster models.DistrictWeatherAlertMaster
		for k, v := range DistrictWeatherAlertWindSpeed {
			if Weatherdata.WeatherData.Windspeed >= v.Min && Weatherdata.WeatherData.Windspeed <= v.Max {
				ServentType = DistrictWeatherAlertWindSpeed[k].SeverityType
				DistrictDistrictWeatherAlertMaster = v
				fmt.Println("windspeed===>", v.ParameterId)

			}
		}
		if !ServentType.IsZero() {
			serventtype, err := d.GetSingleWeatherAlertType(ctx, ServentType.Hex())
			if err != nil {
				return err
			}
			DistrictWeatherAlert := new(models.DistrictWeatherAlert)
			DistrictWeatherAlert.District = district.District
			DistrictWeatherAlert.UniqueID = fmt.Sprintf("%v_%v_%v", Weatherdata.Date.Day(), Weatherdata.Date.Month().String(), Weatherdata.Date.Year())
			DistrictWeatherAlert.ParameterId = paramerterid.WeatherParameter
			DistrictWeatherAlert.SeverityType = serventtype.WeatherAlertType
			DistrictWeatherAlert.WeatherDataAlert = DistrictDistrictWeatherAlertMaster
			DistrictWeatherAlert.WeatherData = *Weatherdata
			DistrictWeatherAlert.Month = *month
			DistrictWeatherAlert.Date = &Weatherdata.Date
			DistrictWeatherAlert.Status = constants.DISTRICTWEATHERALERTSTATUSACTIVE
			DistrictWeatherAlert.Value = Weatherdata.WeatherData.Windspeed
			DistrictWeatherAlert.Tittle = fmt.Sprintf("This %v values above %v", constants.WEATHERPARAMETERWINDSPEED, serventtype.Name)
			t := time.Now()
			created := models.Created{}
			created.On = &t
			created.By = constants.SYSTEM
			log.Println("b4 DistrictWeatherAlert.created ", constants.WEATHERPARAMETERWINDSPEED)
			DistrictWeatherAlert.Created = &created
			err = d.SaveDistrictWeatherAlertWithUpsert(ctx, DistrictWeatherAlert)
			if err != nil {
				return err
			}
			//return nil

		} else {
			log.Println("value is not in range")
			DistrictWeatherAlertNotInRange := new(models.DistrictWeatherAlertNotInRange)
			DistrictWeatherAlertNotInRange.UniqueID = fmt.Sprintf("%v_%v_%v", Weatherdata.Date.Day(), Weatherdata.Date.Month().String(), Weatherdata.Date.Year())
			DistrictWeatherAlertNotInRange.District = district.District
			DistrictWeatherAlertNotInRange.ParameterId = paramerterid.WeatherParameter
			DistrictWeatherAlertNotInRange.WeatherData = *Weatherdata
			DistrictWeatherAlertNotInRange.Month = *month
			DistrictWeatherAlertNotInRange.Date = &Weatherdata.Date
			DistrictWeatherAlertNotInRange.Status = constants.DISTRICTWEATHERALERTSTATUSACTIVE
			DistrictWeatherAlertNotInRange.Value = Weatherdata.WeatherData.Windspeed
			DistrictWeatherAlertNotInRange.Tittle = fmt.Sprintf("This %v values not in range", constants.WEATHERPARAMETERWINDSPEED)
			t := time.Now()
			created := models.Created{}
			created.On = &t
			created.By = constants.SYSTEM
			log.Println("b4 DistrictWeatherAlertNotInRange ", constants.WEATHERPARAMETERWINDSPEED)
			DistrictWeatherAlertNotInRange.Created = &created
			err = d.SaveDistrictWeatherAlertNotInRangeWithUpsert(ctx, DistrictWeatherAlertNotInRange)
			if err != nil {
				return err
			}
		}

	}
	//}
	//	if Weatherdata.WeatherData.Rain > 0 {
	//value := Weatherdata.WeatherData.Rain
	fmt.Println("rain===>", Weatherdata.WeatherData.Rain)
	//find a RainFall Weather alert
	paramerterid, err = d.GetSingleWeatherParameterWithName(ctx, constants.WEATHERPARAMETERRAINFALL)
	if err != nil {
		return err
	}
	fmt.Println("RainFAll===>", paramerterid.Name, paramerterid.ID)

	DistrictWeatherAlertRainFall, err := d.GetSingleDistrictWeatherAlertMasterWithSpecialIds(ctx, month.ID.Hex(), paramerterid.ID.Hex(), Weatherdata.District.Hex())
	if err != nil {
		return err
	}
	if len(DistrictWeatherAlertRainFall) > 0 {
		var ServentType primitive.ObjectID
		var DistrictDistrictWeatherAlertMaster models.DistrictWeatherAlertMaster
		for k, v := range DistrictWeatherAlertRainFall {
			if Weatherdata.WeatherData.Rain >= v.Min && Weatherdata.WeatherData.Rain <= v.Max {
				ServentType = DistrictWeatherAlertRainFall[k].SeverityType
				DistrictDistrictWeatherAlertMaster = v
				fmt.Println("RAinFall ===>", v.ParameterId)

			}
		}
		if !ServentType.IsZero() {
			serventtype, err := d.GetSingleWeatherAlertType(ctx, ServentType.Hex())
			if err != nil {
				return err
			}
			DistrictWeatherAlert := new(models.DistrictWeatherAlert)
			DistrictWeatherAlert.District = district.District
			DistrictWeatherAlert.UniqueID = fmt.Sprintf("%v_%v_%v", Weatherdata.Date.Day(), Weatherdata.Date.Month().String(), Weatherdata.Date.Year())
			DistrictWeatherAlert.ParameterId = paramerterid.WeatherParameter
			DistrictWeatherAlert.SeverityType = serventtype.WeatherAlertType
			DistrictWeatherAlert.WeatherDataAlert = DistrictDistrictWeatherAlertMaster
			DistrictWeatherAlert.WeatherData = *Weatherdata
			DistrictWeatherAlert.Month = *month
			DistrictWeatherAlert.Date = &Weatherdata.Date
			DistrictWeatherAlert.Status = constants.DISTRICTWEATHERALERTSTATUSACTIVE
			DistrictWeatherAlert.Value = Weatherdata.WeatherData.Rain
			DistrictWeatherAlert.Tittle = fmt.Sprintf("This %v values  above %v", constants.WEATHERPARAMETERRAINFALL, serventtype.Name)
			t := time.Now()
			created := models.Created{}
			created.On = &t
			created.By = constants.SYSTEM
			log.Println("b4 DistrictWeatherAlert.created ", constants.WEATHERPARAMETERRAINFALL)
			DistrictWeatherAlert.Created = &created
			err = d.SaveDistrictWeatherAlertWithUpsert(ctx, DistrictWeatherAlert)
			if err != nil {
				return err
			}
			//	return nil

		} else {
			log.Println("value is not in range")
			DistrictWeatherAlertNotInRange := new(models.DistrictWeatherAlertNotInRange)
			DistrictWeatherAlertNotInRange.UniqueID = fmt.Sprintf("%v_%v_%v", Weatherdata.Date.Day(), Weatherdata.Date.Month().String(), Weatherdata.Date.Year())
			DistrictWeatherAlertNotInRange.District = district.District
			DistrictWeatherAlertNotInRange.ParameterId = paramerterid.WeatherParameter
			DistrictWeatherAlertNotInRange.WeatherData = *Weatherdata
			DistrictWeatherAlertNotInRange.Month = *month
			DistrictWeatherAlertNotInRange.Date = &Weatherdata.Date
			DistrictWeatherAlertNotInRange.Status = constants.DISTRICTWEATHERALERTSTATUSACTIVE
			DistrictWeatherAlertNotInRange.Value = Weatherdata.WeatherData.Rain
			DistrictWeatherAlertNotInRange.Tittle = fmt.Sprintf("This %v values not in range", constants.WEATHERPARAMETERRAINFALL)
			t := time.Now()
			created := models.Created{}
			created.On = &t
			created.By = constants.SYSTEM
			log.Println("b4 DistrictWeatherAlertNotInRange ", constants.WEATHERPARAMETERRAINFALL)
			DistrictWeatherAlertNotInRange.Created = &created
			err = d.SaveDistrictWeatherAlertNotInRangeWithUpsert(ctx, DistrictWeatherAlertNotInRange)
			if err != nil {
				return err
			}
		}
	}
	//	}
	//	if Weatherdata.WeatherData.Humidity > 0 {
	//value := Weatherdata.WeatherData.Humidity
	//fmt.Println("Humidity===>", Weatherdata.WeatherData.Humidity)
	//find a humidity Weather alert

	paramerterid, err = d.GetSingleWeatherParameterWithName(ctx, constants.WEATHERPARAMETERHUMIDITY)
	if err != nil {
		return err
	}
	fmt.Println("humidity===>", paramerterid.Name, paramerterid.ID)

	DistrictWeatherAlertHumidity, err := d.GetSingleDistrictWeatherAlertMasterWithSpecialIds(ctx, month.ID.Hex(), paramerterid.ID.Hex(), Weatherdata.District.Hex())
	if err != nil {
		return err
	}
	//var ServentType primitive.ObjectID
	//var DistrictDistrictWeatherAlertMaster models.DistrictDistrictWeatherAlertMaster
	if len(DistrictWeatherAlertHumidity) > 0 {
		var ServentType primitive.ObjectID
		var DistrictDistrictWeatherAlertMaster models.DistrictWeatherAlertMaster
		for k, v := range DistrictWeatherAlertHumidity {
			if Weatherdata.WeatherData.Humidity >= v.Min && Weatherdata.WeatherData.Humidity <= v.Max {
				ServentType = DistrictWeatherAlertHumidity[k].SeverityType
				DistrictDistrictWeatherAlertMaster = v
				fmt.Println("humidity===>", v.ParameterId)

			}
		}
		if !ServentType.IsZero() {
			serventtype, err := d.GetSingleWeatherAlertType(ctx, ServentType.Hex())
			if err != nil {
				return err
			}
			DistrictWeatherAlert := new(models.DistrictWeatherAlert)
			DistrictWeatherAlert.District = district.District
			DistrictWeatherAlert.UniqueID = fmt.Sprintf("%v_%v_%v", Weatherdata.Date.Day(), Weatherdata.Date.Month().String(), Weatherdata.Date.Year())
			DistrictWeatherAlert.ParameterId = paramerterid.WeatherParameter
			DistrictWeatherAlert.SeverityType = serventtype.WeatherAlertType
			DistrictWeatherAlert.WeatherDataAlert = DistrictDistrictWeatherAlertMaster
			DistrictWeatherAlert.WeatherData = *Weatherdata
			DistrictWeatherAlert.Month = *month
			DistrictWeatherAlert.Date = &Weatherdata.Date
			DistrictWeatherAlert.Status = constants.DISTRICTWEATHERALERTSTATUSACTIVE
			DistrictWeatherAlert.Value = Weatherdata.WeatherData.Humidity
			DistrictWeatherAlert.Tittle = fmt.Sprintf("This %v values  above %v", constants.WEATHERPARAMETERHUMIDITY, serventtype.Name)
			t := time.Now()
			created := models.Created{}
			created.On = &t
			created.By = constants.SYSTEM
			log.Println("b4 DistrictWeatherAlert.created ", constants.WEATHERPARAMETERHUMIDITY)
			DistrictWeatherAlert.Created = &created
			err = d.SaveDistrictWeatherAlertWithUpsert(ctx, DistrictWeatherAlert)
			if err != nil {
				return err
			}
			//	return nil

		} else {
			log.Println("value is not in range")
			DistrictWeatherAlertNotInRange := new(models.DistrictWeatherAlertNotInRange)
			DistrictWeatherAlertNotInRange.UniqueID = fmt.Sprintf("%v_%v_%v", Weatherdata.Date.Day(), Weatherdata.Date.Month().String(), Weatherdata.Date.Year())
			DistrictWeatherAlertNotInRange.District = district.District
			DistrictWeatherAlertNotInRange.ParameterId = paramerterid.WeatherParameter
			DistrictWeatherAlertNotInRange.WeatherData = *Weatherdata
			DistrictWeatherAlertNotInRange.Month = *month
			DistrictWeatherAlertNotInRange.Date = &Weatherdata.Date
			DistrictWeatherAlertNotInRange.Status = constants.DISTRICTWEATHERALERTSTATUSACTIVE
			DistrictWeatherAlertNotInRange.Value = Weatherdata.WeatherData.Humidity
			DistrictWeatherAlertNotInRange.Tittle = fmt.Sprintf("This %v values not in range", constants.WEATHERPARAMETERHUMIDITY)
			t := time.Now()
			created := models.Created{}
			created.On = &t
			created.By = constants.SYSTEM
			log.Println("b4 DistrictWeatherAlertNotInRange ", constants.WEATHERPARAMETERHUMIDITY)
			DistrictWeatherAlertNotInRange.Created = &created
			err = d.SaveDistrictWeatherAlertNotInRangeWithUpsert(ctx, DistrictWeatherAlertNotInRange)
			if err != nil {
				return err
			}
		}
	}
	//	}
	//if Weatherdata.WeatherData.Temp.Day > 0 {
	//	valuemax := Weatherdata.WeatherData.Temp.Max
	//valuemin := Weatherdata.WeatherData.Temp.Min
	fmt.Println("Tempmin===>", Weatherdata.WeatherData.Temp.Min)
	fmt.Println("Tempmax===>", Weatherdata.WeatherData.Temp.Max)
	//find a temp Weather alert
	paramerterid, err = d.GetSingleWeatherParameterWithName(ctx, constants.WEATHERPARAMETERTEMPERATURE)
	if err != nil {
		return err
	}
	fmt.Println("temp===>", paramerterid.Name, paramerterid.ID)

	DistrictWeatherAlertTemp, err := d.GetSingleDistrictWeatherAlertMasterWithSpecialIds(ctx, month.ID.Hex(), paramerterid.ID.Hex(), Weatherdata.District.Hex())
	if err != nil {
		return err
	}
	//	var ServentType primitive.ObjectID
	//	var DistrictDistrictWeatherAlertMaster models.DistrictDistrictWeatherAlertMaster
	if len(DistrictWeatherAlertTemp) > 0 {
		var ServentType primitive.ObjectID
		var DistrictDistrictWeatherAlertMaster models.DistrictWeatherAlertMaster
		for k, v := range DistrictWeatherAlertTemp {
			if Weatherdata.WeatherData.Temp.Min >= v.Min && Weatherdata.WeatherData.Temp.Min <= v.Max {
				ServentType = DistrictWeatherAlertTemp[k].SeverityType
				DistrictDistrictWeatherAlertMaster = v
				fmt.Println("temp===>", v.ParameterId)

			}
		}
		if !ServentType.IsZero() {
			serventtype, err := d.GetSingleWeatherAlertType(ctx, ServentType.Hex())
			if err != nil {
				return err
			}
			DistrictWeatherAlert := new(models.DistrictWeatherAlert)
			DistrictWeatherAlert.District = district.District
			DistrictWeatherAlert.UniqueID = fmt.Sprintf("%v_%v_%v", Weatherdata.Date.Day(), Weatherdata.Date.Month().String(), Weatherdata.Date.Year())
			DistrictWeatherAlert.ParameterId = paramerterid.WeatherParameter
			DistrictWeatherAlert.SeverityType = serventtype.WeatherAlertType
			DistrictWeatherAlert.WeatherDataAlert = DistrictDistrictWeatherAlertMaster
			DistrictWeatherAlert.WeatherData = *Weatherdata
			DistrictWeatherAlert.Month = *month
			DistrictWeatherAlert.Date = &Weatherdata.Date
			DistrictWeatherAlert.Status = constants.DISTRICTWEATHERALERTSTATUSACTIVE
			DistrictWeatherAlert.Value = Weatherdata.WeatherData.Temp.Min
			//	DistrictWeatherAlert.ValueMax = Weatherdata.WeatherData.Temp.Max
			DistrictWeatherAlert.Tittle = fmt.Sprintf("This %v min values  above %v", constants.WEATHERPARAMETERTEMPERATURE, serventtype.Name)
			t := time.Now()
			created := models.Created{}
			created.On = &t
			created.By = constants.SYSTEM
			log.Println("b4 DistrictWeatherAlert.created ", constants.WEATHERPARAMETERTEMPERATURE)
			DistrictWeatherAlert.Created = &created
			err = d.SaveDistrictWeatherAlertTempWithUpsert(ctx, DistrictWeatherAlert)
			if err != nil {
				return err
			}
			//	return nil

		} else {
			log.Println("value is not in range")
			DistrictWeatherAlertNotInRange := new(models.DistrictWeatherAlertNotInRange)
			DistrictWeatherAlertNotInRange.UniqueID = fmt.Sprintf("%v_%v_%v", Weatherdata.Date.Day(), Weatherdata.Date.Month().String(), Weatherdata.Date.Year())
			DistrictWeatherAlertNotInRange.District = district.District
			DistrictWeatherAlertNotInRange.ParameterId = paramerterid.WeatherParameter
			DistrictWeatherAlertNotInRange.WeatherData = *Weatherdata
			DistrictWeatherAlertNotInRange.Month = *month
			DistrictWeatherAlertNotInRange.Date = &Weatherdata.Date
			DistrictWeatherAlertNotInRange.Status = constants.DISTRICTWEATHERALERTSTATUSACTIVE
			DistrictWeatherAlertNotInRange.Value = Weatherdata.WeatherData.Temp.Min
			//	DistrictWeatherAlertNotInRange.ValueMax = Weatherdata.WeatherData.Temp.Max
			DistrictWeatherAlertNotInRange.Tittle = fmt.Sprintf("This %v min values not in range", constants.WEATHERPARAMETERTEMPERATURE)
			t := time.Now()
			created := models.Created{}
			created.On = &t
			created.By = constants.SYSTEM
			log.Println("b4 DistrictWeatherAlertNotInRange ", constants.WEATHERPARAMETERTEMPERATURE)
			DistrictWeatherAlertNotInRange.Created = &created
			err = d.SaveDistrictWeatherAlertNotInRangeWithUpsert(ctx, DistrictWeatherAlertNotInRange)
			if err != nil {
				return err
			}
		}
	}
	if len(DistrictWeatherAlertTemp) > 0 {
		var ServentType primitive.ObjectID
		var DistrictDistrictWeatherAlertMaster models.DistrictWeatherAlertMaster
		for k, v := range DistrictWeatherAlertTemp {
			if Weatherdata.WeatherData.Temp.Max >= v.Min && Weatherdata.WeatherData.Temp.Max <= v.Max {
				ServentType = DistrictWeatherAlertTemp[k].SeverityType
				DistrictDistrictWeatherAlertMaster = v
				fmt.Println("temp===>", v.ParameterId)

			}
		}
		if !ServentType.IsZero() {
			serventtype, err := d.GetSingleWeatherAlertType(ctx, ServentType.Hex())
			if err != nil {
				return err
			}
			DistrictWeatherAlert := new(models.DistrictWeatherAlert)
			DistrictWeatherAlert.District = district.District
			DistrictWeatherAlert.UniqueID = fmt.Sprintf("%v_%v_%v", Weatherdata.Date.Day(), Weatherdata.Date.Month().String(), Weatherdata.Date.Year())
			DistrictWeatherAlert.ParameterId = paramerterid.WeatherParameter
			DistrictWeatherAlert.SeverityType = serventtype.WeatherAlertType
			DistrictWeatherAlert.WeatherDataAlert = DistrictDistrictWeatherAlertMaster
			DistrictWeatherAlert.WeatherData = *Weatherdata
			DistrictWeatherAlert.Month = *month
			DistrictWeatherAlert.Date = &Weatherdata.Date
			DistrictWeatherAlert.Status = constants.DISTRICTWEATHERALERTSTATUSACTIVE
			DistrictWeatherAlert.Value = Weatherdata.WeatherData.Temp.Max
			//	DistrictWeatherAlert.ValueMax = Weatherdata.WeatherData.Temp.Max
			DistrictWeatherAlert.Tittle = fmt.Sprintf("This %v max values  above %v", constants.WEATHERPARAMETERTEMPERATURE, serventtype.Name)
			t := time.Now()
			created := models.Created{}
			created.On = &t
			created.By = constants.SYSTEM
			log.Println("b4 DistrictWeatherAlert.created ", constants.WEATHERPARAMETERTEMPERATURE)
			DistrictWeatherAlert.Created = &created
			err = d.SaveDistrictWeatherAlertTempWithUpsert(ctx, DistrictWeatherAlert)
			if err != nil {
				return err
			}
			//	return nil

		} else {
			log.Println("value is not in range")
			DistrictWeatherAlertNotInRange := new(models.DistrictWeatherAlertNotInRange)
			DistrictWeatherAlertNotInRange.UniqueID = fmt.Sprintf("%v_%v_%v", Weatherdata.Date.Day(), Weatherdata.Date.Month().String(), Weatherdata.Date.Year())
			DistrictWeatherAlertNotInRange.District = district.District
			DistrictWeatherAlertNotInRange.ParameterId = paramerterid.WeatherParameter
			DistrictWeatherAlertNotInRange.WeatherData = *Weatherdata
			DistrictWeatherAlertNotInRange.Month = *month
			DistrictWeatherAlertNotInRange.Date = &Weatherdata.Date
			DistrictWeatherAlertNotInRange.Status = constants.DISTRICTWEATHERALERTSTATUSACTIVE
			DistrictWeatherAlertNotInRange.Value = Weatherdata.WeatherData.Temp.Max
			//	DistrictWeatherAlertNotInRange.ValueMax = Weatherdata.WeatherData.Temp.Max
			DistrictWeatherAlertNotInRange.Tittle = fmt.Sprintf("This %v max values not in range", constants.WEATHERPARAMETERTEMPERATURE)
			t := time.Now()
			created := models.Created{}
			created.On = &t
			created.By = constants.SYSTEM
			log.Println("b4 DistrictWeatherAlertNotInRange ", constants.WEATHERPARAMETERTEMPERATURE)
			DistrictWeatherAlertNotInRange.Created = &created
			err = d.SaveDistrictWeatherAlertNotInRangeWithUpsert(ctx, DistrictWeatherAlertNotInRange)
			if err != nil {
				return err
			}
		}
	}
	//	}
	//	if Weatherdata.WeatherData.Winddeg > 0 {
	//	value := Weatherdata.WeatherData.Winddeg
	fmt.Println("Winddeg===>", Weatherdata.WeatherData.Winddeg)
	//find a windDeg Weather alert

	paramerterid, err = d.GetSingleWeatherParameterWithName(ctx, constants.WEATHERPARAMETERWINDDIRECTION)
	if err != nil {
		return err
	}
	fmt.Println("windDeg===>", paramerterid.Name, paramerterid.ID)

	DistrictWeatherAlertWindDeg, err := d.GetSingleDistrictWeatherAlertMasterWithSpecialIds(ctx, month.ID.Hex(), paramerterid.ID.Hex(), Weatherdata.District.Hex())
	if err != nil {
		return err
	}
	//	var ServentType primitive.ObjectID
	//	var DistrictDistrictWeatherAlertMaster models.DistrictDistrictWeatherAlertMaster
	if len(DistrictWeatherAlertWindDeg) > 0 {
		var ServentType primitive.ObjectID
		var DistrictDistrictWeatherAlertMaster models.DistrictWeatherAlertMaster
		for k, v := range DistrictWeatherAlertWindDeg {
			if Weatherdata.WeatherData.Winddeg >= v.Min && Weatherdata.WeatherData.Winddeg <= v.Max {
				ServentType = DistrictWeatherAlertWindDeg[k].SeverityType
				DistrictDistrictWeatherAlertMaster = v

				fmt.Println("winddeg===>", v.ParameterId)

			}
		}
		if !ServentType.IsZero() {
			serventtype, err := d.GetSingleWeatherAlertType(ctx, ServentType.Hex())
			if err != nil {
				return err
			}
			DistrictWeatherAlert := new(models.DistrictWeatherAlert)
			DistrictWeatherAlert.District = district.District
			DistrictWeatherAlert.UniqueID = fmt.Sprintf("%v_%v_%v", Weatherdata.Date.Day(), Weatherdata.Date.Month().String(), Weatherdata.Date.Year())
			DistrictWeatherAlert.ParameterId = paramerterid.WeatherParameter
			DistrictWeatherAlert.SeverityType = serventtype.WeatherAlertType
			DistrictWeatherAlert.WeatherDataAlert = DistrictDistrictWeatherAlertMaster
			DistrictWeatherAlert.WeatherData = *Weatherdata
			DistrictWeatherAlert.Month = *month
			DistrictWeatherAlert.Date = &Weatherdata.Date
			DistrictWeatherAlert.Status = constants.DISTRICTWEATHERALERTSTATUSACTIVE
			DistrictWeatherAlert.Value = Weatherdata.WeatherData.Winddeg
			DistrictWeatherAlert.Tittle = fmt.Sprintf("This %v values  above %v", constants.WEATHERPARAMETERWINDDIRECTION, serventtype.Name)
			t := time.Now()
			created := models.Created{}
			created.On = &t
			created.By = constants.SYSTEM
			log.Println("b4 DistrictWeatherAlert ", constants.WEATHERPARAMETERWINDDIRECTION)
			DistrictWeatherAlert.Created = &created
			err = d.SaveDistrictWeatherAlertWithUpsert(ctx, DistrictWeatherAlert)
			if err != nil {
				return err
			}
			//	return nil

		} else {
			log.Println("value is not in range")
			DistrictWeatherAlertNotInRange := new(models.DistrictWeatherAlertNotInRange)
			DistrictWeatherAlertNotInRange.District = district.District
			DistrictWeatherAlertNotInRange.UniqueID = fmt.Sprintf("%v_%v_%v", Weatherdata.Date.Day(), Weatherdata.Date.Month().String(), Weatherdata.Date.Year())
			DistrictWeatherAlertNotInRange.ParameterId = paramerterid.WeatherParameter
			DistrictWeatherAlertNotInRange.WeatherData = *Weatherdata
			DistrictWeatherAlertNotInRange.Month = *month
			DistrictWeatherAlertNotInRange.Date = &Weatherdata.Date
			DistrictWeatherAlertNotInRange.Status = constants.DISTRICTWEATHERALERTSTATUSACTIVE
			DistrictWeatherAlertNotInRange.Value = Weatherdata.WeatherData.Winddeg
			DistrictWeatherAlertNotInRange.Tittle = fmt.Sprintf("This %v values not in range", constants.WEATHERPARAMETERWINDDIRECTION)
			t := time.Now()
			created := models.Created{}
			created.On = &t
			created.By = constants.SYSTEM
			log.Println("b4 DistrictWeatherAlertNotinRange.created ", constants.WEATHERPARAMETERWINDDIRECTION)
			DistrictWeatherAlertNotInRange.Created = &created
			err = d.SaveDistrictWeatherAlertNotInRangeWithUpsert(ctx, DistrictWeatherAlertNotInRange)
			if err != nil {
				return err
			}
		}
	}
	//}

	return nil
}
func (d *Daos) GetTodayActiveDistrictWeatherAlert(ctx *models.Context) ([]models.DistrictWeatherAlert, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	t := time.Now()
	UniqueID := fmt.Sprintf("%v_%v_%v", t.Day(), t.Month().String(), t.Year())

	query = append(query, bson.M{"status": constants.DISTRICTSTATUSACTIVE})
	query = append(query, bson.M{"uniqueId": UniqueID})
	query = append(query, bson.M{"weatherDataAlert.isAutomatic": "Yes"})

	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	//Aggregation
	d.Shared.BsonToJSONPrintTag("GetTodayActiveDistrictWeatherAlert query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDISTRICTWEATHERALERT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var DistrictWeatherAlert []models.DistrictWeatherAlert
	if err = cursor.All(context.TODO(), &DistrictWeatherAlert); err != nil {
		return nil, err
	}
	return DistrictWeatherAlert, nil
}
