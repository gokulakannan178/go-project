package daos

import (
	"context"
	"errors"
	"fmt"
	"lgfweather-service/constants"
	"lgfweather-service/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//SaveDisease :""
func (d *Daos) SaveWeatherAlertNotInRange(ctx *models.Context, WeatherAlertNotInRange *models.WeatherAlertNotInRange) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONWEATHERALERTNOTINRANGE).InsertOne(ctx.CTX, WeatherAlertNotInRange)
	if err != nil {
		return err
	}
	WeatherAlertNotInRange.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

func (d *Daos) SaveWeatherAlertNotInRangeWithUpsert(ctx *models.Context, WeatherAlert *models.WeatherAlertNotInRange) error {

	opts := options.Update().SetUpsert(true)
	updateQuery := bson.M{"uniqueId": WeatherAlert.UniqueID, "state._id": WeatherAlert.State.ID, "parameter._id": WeatherAlert.ParameterId.ID, "month._id": WeatherAlert.Month.ID}
	updateData := bson.M{"$set": WeatherAlert}
	_, err := ctx.DB.Collection(constants.COLLECTIONWEATHERALERTNOTINRANGE).UpdateOne(ctx.CTX, updateQuery, updateData, opts)
	return err
}
func (d *Daos) SaveWeatherAlertNotInRangeTempWithUpsert(ctx *models.Context, WeatherAlert *models.WeatherAlertNotInRange) error {

	opts := options.Update().SetUpsert(true)
	updateQuery := bson.M{"uniqueId": WeatherAlert.UniqueID, "state._id": WeatherAlert.State.ID, "parameter._id": WeatherAlert.ParameterId.ID, "month._id": WeatherAlert.Month.ID, "tittle": WeatherAlert.Tittle}
	updateData := bson.M{"$set": WeatherAlert}
	_, err := ctx.DB.Collection(constants.COLLECTIONWEATHERALERTNOTINRANGE).UpdateOne(ctx.CTX, updateQuery, updateData, opts)
	return err
}

//GetSingleDisease : ""
func (d *Daos) GetSingleWeatherAlertNotInRange(ctx *models.Context, UniqueID string) (*models.RefWeatherAlertNotInRange, error) {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONWEATHERALERTNOTINRANGE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var WeatherAlertNotInRanges []models.RefWeatherAlertNotInRange
	var WeatherAlertNotInRange *models.RefWeatherAlertNotInRange
	if err = cursor.All(ctx.CTX, &WeatherAlertNotInRanges); err != nil {
		return nil, err
	}
	if len(WeatherAlertNotInRanges) > 0 {
		WeatherAlertNotInRange = &WeatherAlertNotInRanges[0]
	}
	return WeatherAlertNotInRange, nil
}

//UpdateWeatherAlertNotInRange : ""
func (d *Daos) UpdateWeatherAlertNotInRange(ctx *models.Context, WeatherAlertNotInRange *models.WeatherAlertNotInRange) error {

	selector := bson.M{"_id": WeatherAlertNotInRange.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": WeatherAlertNotInRange}
	_, err := ctx.DB.Collection(constants.COLLECTIONWEATHERALERTNOTINRANGE).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterWeatherAlertNotInRange : ""
func (d *Daos) FilterWeatherAlertNotInRange(ctx *models.Context, WeatherAlertNotInRangefilter *models.WeatherAlertNotInRangeFilter, pagination *models.Pagination) ([]models.RefWeatherAlertNotInRange, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if WeatherAlertNotInRangefilter != nil {

		if len(WeatherAlertNotInRangefilter.ActiveStatus) > 0 {
			query = append(query, bson.M{"activeStatus": bson.M{"$in": WeatherAlertNotInRangefilter.ActiveStatus}})
		}
		if len(WeatherAlertNotInRangefilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": WeatherAlertNotInRangefilter.Status}})
		}
		//Regex
		if WeatherAlertNotInRangefilter.SearchBox.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: WeatherAlertNotInRangefilter.SearchBox.Name, Options: "xi"}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if WeatherAlertNotInRangefilter != nil {
		if WeatherAlertNotInRangefilter.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{WeatherAlertNotInRangefilter.SortBy: WeatherAlertNotInRangefilter.SortOrder}})

		}

	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONWEATHERALERTNOTINRANGE).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("WeatherAlertNotInRange query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONWEATHERALERTNOTINRANGE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var WeatherAlertNotInRanges []models.RefWeatherAlertNotInRange
	if err = cursor.All(context.TODO(), &WeatherAlertNotInRanges); err != nil {
		return nil, err
	}
	return WeatherAlertNotInRanges, nil
}

//EnableWeatherAlertNotInRange :""
func (d *Daos) EnableWeatherAlertNotInRange(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.WEATHERALERTNOTINRANGESTATUSACTIVE, "activeStatus": true}}
	_, err = ctx.DB.Collection(constants.COLLECTIONWEATHERALERTNOTINRANGE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableDisease :""
func (d *Daos) DisableWeatherAlertNotInRange(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.WEATHERALERTNOTINRANGESTATUSDISABLE, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONWEATHERALERTNOTINRANGE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteWeatherAlertNotInRange :""
func (d *Daos) DeleteWeatherAlertNotInRange(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.WEATHERALERTNOTINRANGESTATUSDELETED, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONWEATHERALERTNOTINRANGE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
