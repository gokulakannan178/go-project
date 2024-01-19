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
func (d *Daos) SaveStateWeatherAlert(ctx *models.Context, StateWeatherAlert *models.StateWeatherAlert) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONSTATEWEATHERALERT).InsertOne(ctx.CTX, StateWeatherAlert)
	if err != nil {
		return err
	}
	StateWeatherAlert.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}
func (d *Daos) SaveStateWeatherAlertWithUpsert(ctx *models.Context, StateWeatherAlert *models.StateWeatherAlert) error {

	opts := options.Update().SetUpsert(true)
	updateQuery := bson.M{"uniqueId": StateWeatherAlert.UniqueID, "state._id": StateWeatherAlert.State.ID, "parameter._id": StateWeatherAlert.ParameterId.ID, "month._id": StateWeatherAlert.Month.ID, "severityType._id": StateWeatherAlert.SeverityType.ID}
	updateData := bson.M{"$set": StateWeatherAlert}
	_, err := ctx.DB.Collection(constants.COLLECTIONSTATEWEATHERALERT).UpdateOne(ctx.CTX, updateQuery, updateData, opts)
	return err
}
func (d *Daos) SaveStateWeatherAlertTempWithUpsert(ctx *models.Context, StateWeatherAlert *models.StateWeatherAlert) error {

	opts := options.Update().SetUpsert(true)
	updateQuery := bson.M{"uniqueId": StateWeatherAlert.UniqueID, "state._id": StateWeatherAlert.State.ID, "parameter._id": StateWeatherAlert.ParameterId.ID, "month._id": StateWeatherAlert.Month.ID, "severityType._id": StateWeatherAlert.SeverityType.ID, "tittle": StateWeatherAlert.Tittle}
	updateData := bson.M{"$set": StateWeatherAlert}
	_, err := ctx.DB.Collection(constants.COLLECTIONSTATEWEATHERALERT).UpdateOne(ctx.CTX, updateQuery, updateData, opts)
	return err
}

//GetSingleDisease : ""
func (d *Daos) GetSingleStateWeatherAlert(ctx *models.Context, UniqueID string) (*models.RefStateWeatherAlert, error) {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSTATEWEATHERALERT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var StateWeatherAlerts []models.RefStateWeatherAlert
	var StateWeatherAlert *models.RefStateWeatherAlert
	if err = cursor.All(ctx.CTX, &StateWeatherAlerts); err != nil {
		return nil, err
	}
	if len(StateWeatherAlerts) > 0 {
		StateWeatherAlert = &StateWeatherAlerts[0]
	}
	return StateWeatherAlert, nil
}

//UpdateStateWeatherAlert : ""
func (d *Daos) UpdateStateWeatherAlert(ctx *models.Context, StateWeatherAlert *models.StateWeatherAlert) error {

	selector := bson.M{"_id": StateWeatherAlert.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": StateWeatherAlert}
	_, err := ctx.DB.Collection(constants.COLLECTIONSTATEWEATHERALERT).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}
func (d *Daos) UpdateWeatherAlertMaster(ctx *models.Context, StateWeatherAlert *models.UpdateStateWeatherAlert) error {
	selector := bson.M{"weatherDataAlert._id": StateWeatherAlert.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": bson.M{"weatherDataAlert.isSms": StateWeatherAlert.IsSms, "weatherDataAlert.isNicessm": StateWeatherAlert.IsNicessm, "weatherDataAlert.isWhatsApp": StateWeatherAlert.IsWhatsApp, "weatherDataAlert.isTelegram": StateWeatherAlert.IsTelegram, "weatherDataAlert.WeatherAlert": StateWeatherAlert.WeatherAlert}}
	fmt.Println("selector==>", selector)
	fmt.Println("updateInterface==>", updateInterface)
	_, err := ctx.DB.Collection(constants.COLLECTIONSTATEWEATHERALERT).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}
func (d *Daos) IsUpdateStateWeatherAlertMaster(ctx *models.Context, StateWeatherAlert *models.UpdateStateWeatherAlert) error {
	selector := bson.M{"_id": StateWeatherAlert.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": bson.M{"isSms": StateWeatherAlert.IsSms, "isNicessm": StateWeatherAlert.IsNicessm, "isWhatsApp": StateWeatherAlert.IsWhatsApp, "isTelegram": StateWeatherAlert.IsTelegram, "WeatherAlert": StateWeatherAlert.WeatherAlert}}

	_, err := ctx.DB.Collection(constants.COLLECTIONSTATEWEATHERALERTMASTER).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterStateWeatherAlert : ""
func (d *Daos) FilterStateWeatherAlert(ctx *models.Context, StateWeatherAlertfilter *models.StateWeatherAlertFilter, pagination *models.Pagination) ([]models.RefStateWeatherAlert, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if StateWeatherAlertfilter != nil {

		if len(StateWeatherAlertfilter.ActiveStatus) > 0 {
			query = append(query, bson.M{"activeStatus": bson.M{"$in": StateWeatherAlertfilter.ActiveStatus}})
		}
		if len(StateWeatherAlertfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": StateWeatherAlertfilter.Status}})
		}
		if len(StateWeatherAlertfilter.State) > 0 {
			query = append(query, bson.M{"state._id": bson.M{"$in": StateWeatherAlertfilter.State}})
		}
		if len(StateWeatherAlertfilter.ParameterId) > 0 {
			query = append(query, bson.M{"parameter._id": bson.M{"$in": StateWeatherAlertfilter.ParameterId}})
		}
		if len(StateWeatherAlertfilter.SeverityType) > 0 {
			query = append(query, bson.M{"severityType._id": bson.M{"$in": StateWeatherAlertfilter.SeverityType}})
		}
		if len(StateWeatherAlertfilter.WeatherData) > 0 {
			query = append(query, bson.M{"weatherData._id": bson.M{"$in": StateWeatherAlertfilter.WeatherData}})
		}
		if len(StateWeatherAlertfilter.WeatherDataAlert) > 0 {
			query = append(query, bson.M{"weatherDataAlert._id": bson.M{"$in": StateWeatherAlertfilter.WeatherDataAlert}})
		}
		if len(StateWeatherAlertfilter.Month) > 0 {
			query = append(query, bson.M{"month._id": bson.M{"$in": StateWeatherAlertfilter.Month}})
		}
		//Regex
		if StateWeatherAlertfilter.SearchBox.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: StateWeatherAlertfilter.SearchBox.Name, Options: "xi"}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if StateWeatherAlertfilter != nil {
		if StateWeatherAlertfilter.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{StateWeatherAlertfilter.SortBy: StateWeatherAlertfilter.SortOrder}})

		}

	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONSTATEWEATHERALERT).CountDocuments(ctx.CTX, func() bson.M {
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

	//Aggregation
	d.Shared.BsonToJSONPrintTag("StateWeatherAlert query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSTATEWEATHERALERT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var StateWeatherAlerts []models.RefStateWeatherAlert
	if err = cursor.All(context.TODO(), &StateWeatherAlerts); err != nil {
		return nil, err
	}
	return StateWeatherAlerts, nil
}

//EnableStateWeatherAlert :""
func (d *Daos) EnableStateWeatherAlert(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.STATEWEATHERALERTSTATUSACTIVE, "activeStatus": true}}
	_, err = ctx.DB.Collection(constants.COLLECTIONSTATEWEATHERALERT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableDisease :""
func (d *Daos) DisableStateWeatherAlert(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.STATEWEATHERALERTSTATUSDISABLE, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONSTATEWEATHERALERT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteStateWeatherAlert :""
func (d *Daos) DeleteStateWeatherAlert(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.STATEWEATHERALERTSTATUSDELETED, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONSTATEWEATHERALERT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
func (d *Daos) StateWeatherAlertMaster(ctx *models.Context, Weatherdata *models.StateWeatherData) error {
	month, err := d.GetCurrentMonthSeason(ctx, Weatherdata.Date)
	if err != nil {
		return err
	}
	state, err := d.GetSingleState(ctx, Weatherdata.State.Hex())
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
	StateWeatherAlertWindSpeed, err := d.GetSingleStateWeatherAlertMasterWithSpecialIds(ctx, month.ID.Hex(), paramerterid.ID.Hex(), Weatherdata.State.Hex())
	if err != nil {
		return err
	}

	if len(StateWeatherAlertWindSpeed) > 0 {
		var ServentType primitive.ObjectID
		var StateStateWeatherAlertMaster models.StateWeatherAlertMaster
		for k, v := range StateWeatherAlertWindSpeed {
			if Weatherdata.WeatherData.Windspeed >= v.Min && Weatherdata.WeatherData.Windspeed <= v.Max {
				ServentType = StateWeatherAlertWindSpeed[k].SeverityType
				StateStateWeatherAlertMaster = v
				fmt.Println("windspeed===>", v.ParameterId)

			}
		}
		if !ServentType.IsZero() {
			serventtype, err := d.GetSingleWeatherAlertType(ctx, ServentType.Hex())
			if err != nil {
				return err
			}
			StateWeatherAlert := new(models.StateWeatherAlert)
			StateWeatherAlert.State = state.State
			StateWeatherAlert.UniqueID = fmt.Sprintf("%v_%v_%v", Weatherdata.Date.Day(), Weatherdata.Date.Month().String(), Weatherdata.Date.Year())
			StateWeatherAlert.ParameterId = paramerterid.WeatherParameter
			StateWeatherAlert.SeverityType = serventtype.WeatherAlertType
			StateWeatherAlert.WeatherDataAlert = StateStateWeatherAlertMaster
			StateWeatherAlert.WeatherData = *Weatherdata
			StateWeatherAlert.Month = *month
			StateWeatherAlert.Date = &Weatherdata.Date
			StateWeatherAlert.Status = constants.STATEWEATHERALERTSTATUSACTIVE
			StateWeatherAlert.Value = Weatherdata.WeatherData.Windspeed
			StateWeatherAlert.Tittle = fmt.Sprintf("This %v values above %v", constants.WEATHERPARAMETERWINDSPEED, serventtype.Name)
			t := time.Now()
			created := models.Created{}
			created.On = &t
			created.By = constants.SYSTEM
			log.Println("b4 StateWeatherAlert.created ", constants.WEATHERPARAMETERWINDSPEED)
			StateWeatherAlert.Created = &created
			err = d.SaveStateWeatherAlertWithUpsert(ctx, StateWeatherAlert)
			if err != nil {
				return err
			}
			//return nil

		} else {
			log.Println("value is not in range")
			StateWeatherAlertNotInRange := new(models.WeatherAlertNotInRange)
			StateWeatherAlertNotInRange.UniqueID = fmt.Sprintf("%v_%v_%v", Weatherdata.Date.Day(), Weatherdata.Date.Month().String(), Weatherdata.Date.Year())
			StateWeatherAlertNotInRange.State = state.State
			StateWeatherAlertNotInRange.ParameterId = paramerterid.WeatherParameter
			StateWeatherAlertNotInRange.WeatherData = *Weatherdata
			StateWeatherAlertNotInRange.Month = *month
			StateWeatherAlertNotInRange.Date = &Weatherdata.Date
			StateWeatherAlertNotInRange.Status = constants.STATEWEATHERALERTSTATUSACTIVE
			StateWeatherAlertNotInRange.Value = Weatherdata.WeatherData.Windspeed
			StateWeatherAlertNotInRange.Tittle = fmt.Sprintf("This %v values not in range", constants.WEATHERPARAMETERWINDSPEED)
			t := time.Now()
			created := models.Created{}
			created.On = &t
			created.By = constants.SYSTEM
			log.Println("b4 StateWeatherAlertNotInRange ", constants.WEATHERPARAMETERWINDSPEED)
			StateWeatherAlertNotInRange.Created = &created
			err = d.SaveWeatherAlertNotInRangeWithUpsert(ctx, StateWeatherAlertNotInRange)
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

	StateWeatherAlertRainFall, err := d.GetSingleStateWeatherAlertMasterWithSpecialIds(ctx, month.ID.Hex(), paramerterid.ID.Hex(), Weatherdata.State.Hex())
	if err != nil {
		return err
	}
	if len(StateWeatherAlertRainFall) > 0 {
		var ServentType primitive.ObjectID
		var StateStateWeatherAlertMaster models.StateWeatherAlertMaster
		for k, v := range StateWeatherAlertRainFall {
			if Weatherdata.WeatherData.Rain >= v.Min && Weatherdata.WeatherData.Rain <= v.Max {
				ServentType = StateWeatherAlertRainFall[k].SeverityType
				StateStateWeatherAlertMaster = v
				fmt.Println("RAinFall ===>", v.ParameterId)

			}
		}
		if !ServentType.IsZero() {
			serventtype, err := d.GetSingleWeatherAlertType(ctx, ServentType.Hex())
			if err != nil {
				return err
			}
			StateWeatherAlert := new(models.StateWeatherAlert)
			StateWeatherAlert.State = state.State
			StateWeatherAlert.UniqueID = fmt.Sprintf("%v_%v_%v", Weatherdata.Date.Day(), Weatherdata.Date.Month().String(), Weatherdata.Date.Year())
			StateWeatherAlert.ParameterId = paramerterid.WeatherParameter
			StateWeatherAlert.SeverityType = serventtype.WeatherAlertType
			StateWeatherAlert.WeatherDataAlert = StateStateWeatherAlertMaster
			StateWeatherAlert.WeatherData = *Weatherdata
			StateWeatherAlert.Month = *month
			StateWeatherAlert.Date = &Weatherdata.Date
			StateWeatherAlert.Status = constants.STATEWEATHERALERTSTATUSACTIVE
			StateWeatherAlert.Value = Weatherdata.WeatherData.Rain
			StateWeatherAlert.Tittle = fmt.Sprintf("This %v values  above %v", constants.WEATHERPARAMETERRAINFALL, serventtype.Name)
			t := time.Now()
			created := models.Created{}
			created.On = &t
			created.By = constants.SYSTEM
			log.Println("b4 StateWeatherAlert.created ", constants.WEATHERPARAMETERRAINFALL)
			StateWeatherAlert.Created = &created
			err = d.SaveStateWeatherAlertWithUpsert(ctx, StateWeatherAlert)
			if err != nil {
				return err
			}
			//	return nil

		} else {
			log.Println("value is not in range")
			StateWeatherAlertNotInRange := new(models.WeatherAlertNotInRange)
			StateWeatherAlertNotInRange.UniqueID = fmt.Sprintf("%v_%v_%v", Weatherdata.Date.Day(), Weatherdata.Date.Month().String(), Weatherdata.Date.Year())
			StateWeatherAlertNotInRange.State = state.State
			StateWeatherAlertNotInRange.ParameterId = paramerterid.WeatherParameter
			StateWeatherAlertNotInRange.WeatherData = *Weatherdata
			StateWeatherAlertNotInRange.Month = *month
			StateWeatherAlertNotInRange.Date = &Weatherdata.Date
			StateWeatherAlertNotInRange.Status = constants.STATEWEATHERALERTSTATUSACTIVE
			StateWeatherAlertNotInRange.Value = Weatherdata.WeatherData.Rain
			StateWeatherAlertNotInRange.Tittle = fmt.Sprintf("This %v values not in range", constants.WEATHERPARAMETERRAINFALL)
			t := time.Now()
			created := models.Created{}
			created.On = &t
			created.By = constants.SYSTEM
			log.Println("b4 StateWeatherAlertNotInRange ", constants.WEATHERPARAMETERRAINFALL)
			StateWeatherAlertNotInRange.Created = &created
			err = d.SaveWeatherAlertNotInRangeWithUpsert(ctx, StateWeatherAlertNotInRange)
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

	StateWeatherAlertHumidity, err := d.GetSingleStateWeatherAlertMasterWithSpecialIds(ctx, month.ID.Hex(), paramerterid.ID.Hex(), Weatherdata.State.Hex())
	if err != nil {
		return err
	}
	//var ServentType primitive.ObjectID
	//var StateStateWeatherAlertMaster models.StateStateWeatherAlertMaster
	if len(StateWeatherAlertHumidity) > 0 {
		var ServentType primitive.ObjectID
		var StateStateWeatherAlertMaster models.StateWeatherAlertMaster
		for k, v := range StateWeatherAlertHumidity {
			if Weatherdata.WeatherData.Humidity >= v.Min && Weatherdata.WeatherData.Humidity <= v.Max {
				ServentType = StateWeatherAlertHumidity[k].SeverityType
				StateStateWeatherAlertMaster = v
				fmt.Println("humidity===>", v.ParameterId)

			}
		}
		if !ServentType.IsZero() {
			serventtype, err := d.GetSingleWeatherAlertType(ctx, ServentType.Hex())
			if err != nil {
				return err
			}
			StateWeatherAlert := new(models.StateWeatherAlert)
			StateWeatherAlert.State = state.State
			StateWeatherAlert.UniqueID = fmt.Sprintf("%v_%v_%v", Weatherdata.Date.Day(), Weatherdata.Date.Month().String(), Weatherdata.Date.Year())
			StateWeatherAlert.ParameterId = paramerterid.WeatherParameter
			StateWeatherAlert.SeverityType = serventtype.WeatherAlertType
			StateWeatherAlert.WeatherDataAlert = StateStateWeatherAlertMaster
			StateWeatherAlert.WeatherData = *Weatherdata
			StateWeatherAlert.Month = *month
			StateWeatherAlert.Date = &Weatherdata.Date
			StateWeatherAlert.Status = constants.STATEWEATHERALERTSTATUSACTIVE
			StateWeatherAlert.Value = Weatherdata.WeatherData.Humidity
			StateWeatherAlert.Tittle = fmt.Sprintf("This %v values  above %v", constants.WEATHERPARAMETERHUMIDITY, serventtype.Name)
			t := time.Now()
			created := models.Created{}
			created.On = &t
			created.By = constants.SYSTEM
			log.Println("b4 StateWeatherAlert.created ", constants.WEATHERPARAMETERHUMIDITY)
			StateWeatherAlert.Created = &created
			err = d.SaveStateWeatherAlertWithUpsert(ctx, StateWeatherAlert)
			if err != nil {
				return err
			}
			//	return nil

		} else {
			log.Println("value is not in range")
			StateWeatherAlertNotInRange := new(models.WeatherAlertNotInRange)
			StateWeatherAlertNotInRange.UniqueID = fmt.Sprintf("%v_%v_%v", Weatherdata.Date.Day(), Weatherdata.Date.Month().String(), Weatherdata.Date.Year())
			StateWeatherAlertNotInRange.State = state.State
			StateWeatherAlertNotInRange.ParameterId = paramerterid.WeatherParameter
			StateWeatherAlertNotInRange.WeatherData = *Weatherdata
			StateWeatherAlertNotInRange.Month = *month
			StateWeatherAlertNotInRange.Date = &Weatherdata.Date
			StateWeatherAlertNotInRange.Status = constants.STATEWEATHERALERTSTATUSACTIVE
			StateWeatherAlertNotInRange.Value = Weatherdata.WeatherData.Humidity
			StateWeatherAlertNotInRange.Tittle = fmt.Sprintf("This %v values not in range", constants.WEATHERPARAMETERHUMIDITY)
			t := time.Now()
			created := models.Created{}
			created.On = &t
			created.By = constants.SYSTEM
			log.Println("b4 StateWeatherAlertNotInRange ", constants.WEATHERPARAMETERHUMIDITY)
			StateWeatherAlertNotInRange.Created = &created
			err = d.SaveWeatherAlertNotInRangeWithUpsert(ctx, StateWeatherAlertNotInRange)
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

	StateWeatherAlertTemp, err := d.GetSingleStateWeatherAlertMasterWithSpecialIds(ctx, month.ID.Hex(), paramerterid.ID.Hex(), Weatherdata.State.Hex())
	if err != nil {
		return err
	}
	//	var ServentType primitive.ObjectID
	//	var StateStateWeatherAlertMaster models.StateStateWeatherAlertMaster
	if len(StateWeatherAlertTemp) > 0 {
		var ServentType primitive.ObjectID
		var StateStateWeatherAlertMaster models.StateWeatherAlertMaster
		for k, v := range StateWeatherAlertTemp {
			if Weatherdata.WeatherData.Temp.Min >= v.Min && Weatherdata.WeatherData.Temp.Min <= v.Max {
				ServentType = StateWeatherAlertTemp[k].SeverityType
				StateStateWeatherAlertMaster = v
				fmt.Println("temp===>", v.ParameterId)

			}
		}
		if !ServentType.IsZero() {
			serventtype, err := d.GetSingleWeatherAlertType(ctx, ServentType.Hex())
			if err != nil {
				return err
			}
			StateWeatherAlert := new(models.StateWeatherAlert)
			StateWeatherAlert.State = state.State
			StateWeatherAlert.UniqueID = fmt.Sprintf("%v_%v_%v", Weatherdata.Date.Day(), Weatherdata.Date.Month().String(), Weatherdata.Date.Year())
			StateWeatherAlert.ParameterId = paramerterid.WeatherParameter
			StateWeatherAlert.SeverityType = serventtype.WeatherAlertType
			StateWeatherAlert.WeatherDataAlert = StateStateWeatherAlertMaster
			StateWeatherAlert.WeatherData = *Weatherdata
			StateWeatherAlert.Month = *month
			StateWeatherAlert.Date = &Weatherdata.Date
			StateWeatherAlert.Status = constants.STATEWEATHERALERTSTATUSACTIVE
			StateWeatherAlert.Value = Weatherdata.WeatherData.Temp.Min
			//	StateWeatherAlert.ValueMax = Weatherdata.WeatherData.Temp.Max
			StateWeatherAlert.Tittle = fmt.Sprintf("This %v min values  above %v", constants.WEATHERPARAMETERTEMPERATURE, serventtype.Name)
			t := time.Now()
			created := models.Created{}
			created.On = &t
			created.By = constants.SYSTEM
			log.Println("b4 StateWeatherAlert.created ", constants.WEATHERPARAMETERTEMPERATURE)
			StateWeatherAlert.Created = &created
			err = d.SaveStateWeatherAlertTempWithUpsert(ctx, StateWeatherAlert)
			if err != nil {
				return err
			}
			//	return nil

		} else {
			log.Println("value is not in range")
			StateWeatherAlertNotInRange := new(models.WeatherAlertNotInRange)
			StateWeatherAlertNotInRange.UniqueID = fmt.Sprintf("%v_%v_%v", Weatherdata.Date.Day(), Weatherdata.Date.Month().String(), Weatherdata.Date.Year())
			StateWeatherAlertNotInRange.State = state.State
			StateWeatherAlertNotInRange.ParameterId = paramerterid.WeatherParameter
			StateWeatherAlertNotInRange.WeatherData = *Weatherdata
			StateWeatherAlertNotInRange.Month = *month
			StateWeatherAlertNotInRange.Date = &Weatherdata.Date
			StateWeatherAlertNotInRange.Status = constants.STATEWEATHERALERTSTATUSACTIVE
			StateWeatherAlertNotInRange.Value = Weatherdata.WeatherData.Temp.Min
			//	StateWeatherAlertNotInRange.ValueMax = Weatherdata.WeatherData.Temp.Max
			StateWeatherAlertNotInRange.Tittle = fmt.Sprintf("This %v min values not in range", constants.WEATHERPARAMETERTEMPERATURE)
			t := time.Now()
			created := models.Created{}
			created.On = &t
			created.By = constants.SYSTEM
			log.Println("b4 StateWeatherAlertNotInRange ", constants.WEATHERPARAMETERTEMPERATURE)
			StateWeatherAlertNotInRange.Created = &created
			err = d.SaveWeatherAlertNotInRangeTempWithUpsert(ctx, StateWeatherAlertNotInRange)
			if err != nil {
				return err
			}
		}
	}
	if len(StateWeatherAlertTemp) > 0 {
		var ServentType primitive.ObjectID
		var StateStateWeatherAlertMaster models.StateWeatherAlertMaster
		for k, v := range StateWeatherAlertTemp {
			if Weatherdata.WeatherData.Temp.Max >= v.Min && Weatherdata.WeatherData.Temp.Max <= v.Max {
				ServentType = StateWeatherAlertTemp[k].SeverityType
				StateStateWeatherAlertMaster = v
				fmt.Println("temp===>", v.ParameterId)

			}
		}
		if !ServentType.IsZero() {
			serventtype, err := d.GetSingleWeatherAlertType(ctx, ServentType.Hex())
			if err != nil {
				return err
			}
			StateWeatherAlert := new(models.StateWeatherAlert)
			StateWeatherAlert.State = state.State
			StateWeatherAlert.UniqueID = fmt.Sprintf("%v_%v_%v", Weatherdata.Date.Day(), Weatherdata.Date.Month().String(), Weatherdata.Date.Year())
			StateWeatherAlert.ParameterId = paramerterid.WeatherParameter
			StateWeatherAlert.SeverityType = serventtype.WeatherAlertType
			StateWeatherAlert.WeatherDataAlert = StateStateWeatherAlertMaster
			StateWeatherAlert.WeatherData = *Weatherdata
			StateWeatherAlert.Month = *month
			StateWeatherAlert.Date = &Weatherdata.Date
			StateWeatherAlert.Status = constants.STATEWEATHERALERTSTATUSACTIVE
			StateWeatherAlert.Value = Weatherdata.WeatherData.Temp.Max
			//	StateWeatherAlert.ValueMax = Weatherdata.WeatherData.Temp.Max
			StateWeatherAlert.Tittle = fmt.Sprintf("This %v max values  above %v", constants.WEATHERPARAMETERTEMPERATURE, serventtype.Name)
			t := time.Now()
			created := models.Created{}
			created.On = &t
			created.By = constants.SYSTEM
			log.Println("b4 StateWeatherAlert.created ", constants.WEATHERPARAMETERTEMPERATURE)
			StateWeatherAlert.Created = &created
			err = d.SaveStateWeatherAlertTempWithUpsert(ctx, StateWeatherAlert)
			if err != nil {
				return err
			}
			//	return nil

		} else {
			log.Println("value is not in range")
			StateWeatherAlertNotInRange := new(models.WeatherAlertNotInRange)
			StateWeatherAlertNotInRange.UniqueID = fmt.Sprintf("%v_%v_%v", Weatherdata.Date.Day(), Weatherdata.Date.Month().String(), Weatherdata.Date.Year())
			StateWeatherAlertNotInRange.State = state.State
			StateWeatherAlertNotInRange.ParameterId = paramerterid.WeatherParameter
			StateWeatherAlertNotInRange.WeatherData = *Weatherdata
			StateWeatherAlertNotInRange.Month = *month
			StateWeatherAlertNotInRange.Date = &Weatherdata.Date
			StateWeatherAlertNotInRange.Status = constants.STATEWEATHERALERTSTATUSACTIVE
			StateWeatherAlertNotInRange.Value = Weatherdata.WeatherData.Temp.Max
			//	StateWeatherAlertNotInRange.ValueMax = Weatherdata.WeatherData.Temp.Max
			StateWeatherAlertNotInRange.Tittle = fmt.Sprintf("This %v max values not in range", constants.WEATHERPARAMETERTEMPERATURE)
			t := time.Now()
			created := models.Created{}
			created.On = &t
			created.By = constants.SYSTEM
			log.Println("b4 StateWeatherAlertNotInRange ", constants.WEATHERPARAMETERTEMPERATURE)
			StateWeatherAlertNotInRange.Created = &created
			err = d.SaveWeatherAlertNotInRangeTempWithUpsert(ctx, StateWeatherAlertNotInRange)
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

	StateWeatherAlertWindDeg, err := d.GetSingleStateWeatherAlertMasterWithSpecialIds(ctx, month.ID.Hex(), paramerterid.ID.Hex(), Weatherdata.State.Hex())
	if err != nil {
		return err
	}
	//	var ServentType primitive.ObjectID
	//	var StateStateWeatherAlertMaster models.StateStateWeatherAlertMaster
	if len(StateWeatherAlertWindDeg) > 0 {
		var ServentType primitive.ObjectID
		var StateStateWeatherAlertMaster models.StateWeatherAlertMaster
		for k, v := range StateWeatherAlertWindDeg {
			if Weatherdata.WeatherData.Winddeg >= v.Min && Weatherdata.WeatherData.Winddeg <= v.Max {
				ServentType = StateWeatherAlertWindDeg[k].SeverityType
				StateStateWeatherAlertMaster = v

				fmt.Println("winddeg===>", v.ParameterId)

			}
		}
		if !ServentType.IsZero() {
			serventtype, err := d.GetSingleWeatherAlertType(ctx, ServentType.Hex())
			if err != nil {
				return err
			}
			StateWeatherAlert := new(models.StateWeatherAlert)
			StateWeatherAlert.State = state.State
			StateWeatherAlert.UniqueID = fmt.Sprintf("%v_%v_%v", Weatherdata.Date.Day(), Weatherdata.Date.Month().String(), Weatherdata.Date.Year())
			StateWeatherAlert.ParameterId = paramerterid.WeatherParameter
			StateWeatherAlert.SeverityType = serventtype.WeatherAlertType
			StateWeatherAlert.WeatherDataAlert = StateStateWeatherAlertMaster
			StateWeatherAlert.WeatherData = *Weatherdata
			StateWeatherAlert.Month = *month
			StateWeatherAlert.Date = &Weatherdata.Date
			StateWeatherAlert.Status = constants.STATEWEATHERALERTSTATUSACTIVE
			StateWeatherAlert.Value = Weatherdata.WeatherData.Winddeg
			StateWeatherAlert.Tittle = fmt.Sprintf("This %v values  above %v", constants.WEATHERPARAMETERWINDDIRECTION, serventtype.Name)
			t := time.Now()
			created := models.Created{}
			created.On = &t
			created.By = constants.SYSTEM
			log.Println("b4 StateWeatherAlert ", constants.WEATHERPARAMETERWINDDIRECTION)
			StateWeatherAlert.Created = &created
			err = d.SaveStateWeatherAlertWithUpsert(ctx, StateWeatherAlert)
			if err != nil {
				return err
			}
			//	return nil

		} else {
			log.Println("value is not in range")
			StateWeatherAlertNotInRange := new(models.WeatherAlertNotInRange)
			StateWeatherAlertNotInRange.State = state.State
			StateWeatherAlertNotInRange.UniqueID = fmt.Sprintf("%v_%v_%v", Weatherdata.Date.Day(), Weatherdata.Date.Month().String(), Weatherdata.Date.Year())
			StateWeatherAlertNotInRange.ParameterId = paramerterid.WeatherParameter
			StateWeatherAlertNotInRange.WeatherData = *Weatherdata
			StateWeatherAlertNotInRange.Month = *month
			StateWeatherAlertNotInRange.Date = &Weatherdata.Date
			StateWeatherAlertNotInRange.Status = constants.STATEWEATHERALERTSTATUSACTIVE
			StateWeatherAlertNotInRange.Value = Weatherdata.WeatherData.Winddeg
			StateWeatherAlertNotInRange.Tittle = fmt.Sprintf("This %v values not in range", constants.WEATHERPARAMETERWINDDIRECTION)
			t := time.Now()
			created := models.Created{}
			created.On = &t
			created.By = constants.SYSTEM
			log.Println("b4 StateWeatherAlertNotinRange.created ", constants.WEATHERPARAMETERWINDDIRECTION)
			StateWeatherAlertNotInRange.Created = &created
			err = d.SaveWeatherAlertNotInRangeWithUpsert(ctx, StateWeatherAlertNotInRange)
			if err != nil {
				return err
			}
		}
	}
	//}

	return nil
}
func (d *Daos) GetTodayActiveStateWeatherAlert(ctx *models.Context) ([]models.StateWeatherAlert, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	t := time.Now()
	UniqueID := fmt.Sprintf("%v_%v_%v", t.Day(), t.Month().String(), t.Year())

	query = append(query, bson.M{"status": constants.STATESTATUSACTIVE})
	query = append(query, bson.M{"uniqueId": UniqueID})
	query = append(query, bson.M{"weatherDataAlert.isAutomatic": "Yes"})

	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	//Aggregation
	d.Shared.BsonToJSONPrintTag("GetTodayActiveStateWeatherAlert query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSTATEWEATHERALERT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var StateWeatherAlert []models.StateWeatherAlert
	if err = cursor.All(context.TODO(), &StateWeatherAlert); err != nil {
		return nil, err
	}
	return StateWeatherAlert, nil
}
