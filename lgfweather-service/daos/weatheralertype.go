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
)

//SaveDisease :""
func (d *Daos) SaveWeatherAlertType(ctx *models.Context, WeatherAlertType *models.WeatherAlertType) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONWEATHERALERTTYPE).InsertOne(ctx.CTX, WeatherAlertType)
	if err != nil {
		return err
	}
	WeatherAlertType.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//GetSingleDisease : ""
func (d *Daos) GetSingleWeatherAlertType(ctx *models.Context, UniqueID string) (*models.RefWeatherAlertType, error) {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})
	d.Shared.BsonToJSONPrintTag("WeatherAlertType query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONWEATHERALERTTYPE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var WeatherAlertTypes []models.RefWeatherAlertType
	var WeatherAlertType *models.RefWeatherAlertType
	if err = cursor.All(ctx.CTX, &WeatherAlertTypes); err != nil {
		return nil, err
	}
	if len(WeatherAlertTypes) > 0 {
		WeatherAlertType = &WeatherAlertTypes[0]
	}
	return WeatherAlertType, nil
}

//UpdateWeatherAlertType : ""
func (d *Daos) UpdateWeatherAlertType(ctx *models.Context, WeatherAlertType *models.WeatherAlertType) error {

	selector := bson.M{"_id": WeatherAlertType.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": WeatherAlertType}
	_, err := ctx.DB.Collection(constants.COLLECTIONWEATHERALERTTYPE).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterWeatherAlertType : ""
func (d *Daos) FilterWeatherAlertType(ctx *models.Context, WeatherAlertTypefilter *models.WeatherAlertTypeFilter, pagination *models.Pagination) ([]models.RefWeatherAlertType, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if WeatherAlertTypefilter != nil {

		if len(WeatherAlertTypefilter.ActiveStatus) > 0 {
			query = append(query, bson.M{"activeStatus": bson.M{"$in": WeatherAlertTypefilter.ActiveStatus}})
		}
		if len(WeatherAlertTypefilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": WeatherAlertTypefilter.Status}})
		}
		//Regex
		if WeatherAlertTypefilter.SearchBox.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: WeatherAlertTypefilter.SearchBox.Name, Options: "xi"}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if WeatherAlertTypefilter != nil {
		if WeatherAlertTypefilter.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{WeatherAlertTypefilter.SortBy: WeatherAlertTypefilter.SortOrder}})

		}

	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONWEATHERALERTTYPE).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("WeatherAlertType query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONWEATHERALERTTYPE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var WeatherAlertTypes []models.RefWeatherAlertType
	if err = cursor.All(context.TODO(), &WeatherAlertTypes); err != nil {
		return nil, err
	}
	return WeatherAlertTypes, nil
}

//EnableWeatherAlertType :""
func (d *Daos) EnableWeatherAlertType(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.WEATHERALERTTYPESTATUSACTIVE, "activeStatus": true}}
	_, err = ctx.DB.Collection(constants.COLLECTIONWEATHERALERTTYPE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableDisease :""
func (d *Daos) DisableWeatherAlertType(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.WEATHERALERTTYPESTATUSDISABLE, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONWEATHERALERTTYPE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteWeatherAlertType :""
func (d *Daos) DeleteWeatherAlertType(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.WEATHERALERTTYPESTATUSDELETED, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONWEATHERALERTTYPE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
