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
func (d *Daos) SaveWeatherParameter(ctx *models.Context, WeatherParameter *models.WeatherParameter) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONWEATHERPARAMETER).InsertOne(ctx.CTX, WeatherParameter)
	if err != nil {
		return err
	}
	WeatherParameter.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//GetSingleDisease : ""
func (d *Daos) GetSingleWeatherParameter(ctx *models.Context, UniqueID string) (*models.RefWeatherParameter, error) {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONWEATHERPARAMETER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var WeatherParameters []models.RefWeatherParameter
	var WeatherParameter *models.RefWeatherParameter
	if err = cursor.All(ctx.CTX, &WeatherParameters); err != nil {
		return nil, err
	}
	if len(WeatherParameters) > 0 {
		WeatherParameter = &WeatherParameters[0]
	}
	return WeatherParameter, nil
}

//UpdateWeatherParameter : ""
func (d *Daos) UpdateWeatherParameter(ctx *models.Context, WeatherParameter *models.WeatherParameter) error {

	selector := bson.M{"_id": WeatherParameter.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": WeatherParameter}
	_, err := ctx.DB.Collection(constants.COLLECTIONWEATHERPARAMETER).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterWeatherParameter : ""
func (d *Daos) FilterWeatherParameter(ctx *models.Context, WeatherParameterfilter *models.WeatherParameterFilter, pagination *models.Pagination) ([]models.RefWeatherParameter, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if WeatherParameterfilter != nil {

		if len(WeatherParameterfilter.ActiveStatus) > 0 {
			query = append(query, bson.M{"activeStatus": bson.M{"$in": WeatherParameterfilter.ActiveStatus}})
		}
		if len(WeatherParameterfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": WeatherParameterfilter.Status}})
		}
		//Regex
		if WeatherParameterfilter.SearchBox.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: WeatherParameterfilter.SearchBox.Name, Options: "xi"}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if WeatherParameterfilter != nil {
		if WeatherParameterfilter.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{WeatherParameterfilter.SortBy: WeatherParameterfilter.SortOrder}})

		}

	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONWEATHERPARAMETER).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("WeatherParameter query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONWEATHERPARAMETER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var WeatherParameters []models.RefWeatherParameter
	if err = cursor.All(context.TODO(), &WeatherParameters); err != nil {
		return nil, err
	}
	return WeatherParameters, nil
}

//EnableWeatherParameter :""
func (d *Daos) EnableWeatherParameter(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.WEATHERPARAMETERSTATUSACTIVE, "activeStatus": true}}
	_, err = ctx.DB.Collection(constants.COLLECTIONWEATHERPARAMETER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableDisease :""
func (d *Daos) DisableWeatherParameter(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.WEATHERPARAMETERSTATUSDISABLE, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONWEATHERPARAMETER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteWeatherParameter :""
func (d *Daos) DeleteWeatherParameter(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.WEATHERPARAMETERSTATUSDELETED, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONWEATHERPARAMETER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
func (d *Daos) GetSingleWeatherParameterWithName(ctx *models.Context, name string) (*models.RefWeatherParameter, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"name": name}})

	d.Shared.BsonToJSONPrintTag("GetSingleWeatherParameterWithName query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONWEATHERPARAMETER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var WeatherParameters []models.RefWeatherParameter
	var WeatherParameter *models.RefWeatherParameter
	if err = cursor.All(ctx.CTX, &WeatherParameters); err != nil {
		return nil, err
	}
	if len(WeatherParameters) > 0 {
		WeatherParameter = &WeatherParameters[0]
	}
	return WeatherParameter, nil
}
