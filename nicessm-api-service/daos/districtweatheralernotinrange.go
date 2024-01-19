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
func (d *Daos) SaveDistrictWeatherAlertNotInRange(ctx *models.Context, DistrictWeatherAlertNotInRange *models.DistrictWeatherAlertNotInRange) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONDISTRICTWEATHERALERTNOTINRANGE).InsertOne(ctx.CTX, DistrictWeatherAlertNotInRange)
	if err != nil {
		return err
	}
	DistrictWeatherAlertNotInRange.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

func (d *Daos) SaveDistrictWeatherAlertNotInRangeWithUpsert(ctx *models.Context, WeatherAlert *models.DistrictWeatherAlertNotInRange) error {

	opts := options.Update().SetUpsert(true)
	updateQuery := bson.M{"uniqueId": WeatherAlert.UniqueID, "district._id": WeatherAlert.District.ID, "parameter._id": WeatherAlert.ParameterId.ID, "month._id": WeatherAlert.Month.ID}
	updateData := bson.M{"$set": WeatherAlert}
	_, err := ctx.DB.Collection(constants.COLLECTIONDISTRICTWEATHERALERTNOTINRANGE).UpdateOne(ctx.CTX, updateQuery, updateData, opts)
	return err
}
func (d *Daos) SaveDistrictWeatherAlertNotInRangeTempWithUpsert(ctx *models.Context, WeatherAlert *models.DistrictWeatherAlertNotInRange) error {

	opts := options.Update().SetUpsert(true)
	updateQuery := bson.M{"uniqueId": WeatherAlert.UniqueID, "district._id": WeatherAlert.District.ID, "parameter._id": WeatherAlert.ParameterId.ID, "month._id": WeatherAlert.Month.ID, "tittle": WeatherAlert.Tittle}
	updateData := bson.M{"$set": WeatherAlert}
	_, err := ctx.DB.Collection(constants.COLLECTIONDISTRICTWEATHERALERTNOTINRANGE).UpdateOne(ctx.CTX, updateQuery, updateData, opts)
	return err
}

//GetSingleDisease : ""
func (d *Daos) GetSingleDistrictWeatherAlertNotInRange(ctx *models.Context, UniqueID string) (*models.RefDistrictWeatherAlertNotInRange, error) {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDISTRICTWEATHERALERTNOTINRANGE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var DistrictWeatherAlertNotInRanges []models.RefDistrictWeatherAlertNotInRange
	var DistrictWeatherAlertNotInRange *models.RefDistrictWeatherAlertNotInRange
	if err = cursor.All(ctx.CTX, &DistrictWeatherAlertNotInRanges); err != nil {
		return nil, err
	}
	if len(DistrictWeatherAlertNotInRanges) > 0 {
		DistrictWeatherAlertNotInRange = &DistrictWeatherAlertNotInRanges[0]
	}
	return DistrictWeatherAlertNotInRange, nil
}

//UpdateDistrictWeatherAlertNotInRange : ""
func (d *Daos) UpdateDistrictWeatherAlertNotInRange(ctx *models.Context, DistrictWeatherAlertNotInRange *models.DistrictWeatherAlertNotInRange) error {

	selector := bson.M{"_id": DistrictWeatherAlertNotInRange.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": DistrictWeatherAlertNotInRange}
	_, err := ctx.DB.Collection(constants.COLLECTIONDISTRICTWEATHERALERTNOTINRANGE).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterDistrictWeatherAlertNotInRange : ""
func (d *Daos) FilterDistrictWeatherAlertNotInRange(ctx *models.Context, DistrictWeatherAlertNotInRangefilter *models.DistrictWeatherAlertNotInRangeFilter, pagination *models.Pagination) ([]models.RefDistrictWeatherAlertNotInRange, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if DistrictWeatherAlertNotInRangefilter != nil {

		if len(DistrictWeatherAlertNotInRangefilter.ActiveStatus) > 0 {
			query = append(query, bson.M{"activeStatus": bson.M{"$in": DistrictWeatherAlertNotInRangefilter.ActiveStatus}})
		}
		if len(DistrictWeatherAlertNotInRangefilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": DistrictWeatherAlertNotInRangefilter.Status}})
		}
		//Regex
		if DistrictWeatherAlertNotInRangefilter.SearchBox.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: DistrictWeatherAlertNotInRangefilter.SearchBox.Name, Options: "xi"}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if DistrictWeatherAlertNotInRangefilter != nil {
		if DistrictWeatherAlertNotInRangefilter.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{DistrictWeatherAlertNotInRangefilter.SortBy: DistrictWeatherAlertNotInRangefilter.SortOrder}})

		}

	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONDISTRICTWEATHERALERTNOTINRANGE).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("DistrictWeatherAlertNotInRange query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDISTRICTWEATHERALERTNOTINRANGE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var DistrictWeatherAlertNotInRanges []models.RefDistrictWeatherAlertNotInRange
	if err = cursor.All(context.TODO(), &DistrictWeatherAlertNotInRanges); err != nil {
		return nil, err
	}
	return DistrictWeatherAlertNotInRanges, nil
}

//EnableDistrictWeatherAlertNotInRange :""
func (d *Daos) EnableDistrictWeatherAlertNotInRange(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.DISTRICTWEATHERALERTNOTINRANGESTATUSACTIVE, "activeStatus": true}}
	_, err = ctx.DB.Collection(constants.COLLECTIONDISTRICTWEATHERALERTNOTINRANGE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableDisease :""
func (d *Daos) DisableDistrictWeatherAlertNotInRange(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.DISTRICTWEATHERALERTNOTINRANGESTATUSDISABLE, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONDISTRICTWEATHERALERTNOTINRANGE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteDistrictWeatherAlertNotInRange :""
func (d *Daos) DeleteDistrictWeatherAlertNotInRange(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.DISTRICTWEATHERALERTNOTINRANGESTATUSDELETED, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONDISTRICTWEATHERALERTNOTINRANGE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
