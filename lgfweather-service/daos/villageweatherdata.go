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
func (d *Daos) SaveVillageWeatherData(ctx *models.Context, villageweatherdata *models.VillageWeatherData) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONVILLAGEWEATHERDATA).InsertOne(ctx.CTX, villageweatherdata)
	if err != nil {
		return err
	}
	villageweatherdata.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}
func (d *Daos) SaveVillageWeatherDataWithUpsert(ctx *models.Context, villageweatherdata *models.VillageWeatherData) error {

	opts := options.Update().SetUpsert(true)
	updateQuery := bson.M{"Village": villageweatherdata.Village, "uniqueId": villageweatherdata.UniqueID}
	updateData := bson.M{"$set": villageweatherdata}
	_, err := ctx.DB.Collection(constants.COLLECTIONVILLAGEWEATHERDATA).UpdateOne(ctx.CTX, updateQuery, updateData, opts)
	return err
}

//GetSingleVillageWeatherData : ""
func (d *Daos) GetSingleVillageWeatherData(ctx *models.Context, UniqueID string) (*models.RefVillageWeatherData, error) {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONVILLAGEWEATHERDATA).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var villageweatherdatas []models.RefVillageWeatherData
	var villageweatherdata *models.RefVillageWeatherData
	if err = cursor.All(ctx.CTX, &villageweatherdatas); err != nil {
		return nil, err
	}
	if len(villageweatherdatas) > 0 {
		villageweatherdata = &villageweatherdatas[0]
	}
	return villageweatherdata, nil
}

//UpdateVillageWeatherData : ""
func (d *Daos) UpdateVillageWeatherData(ctx *models.Context, villageweatherdata *models.VillageWeatherData) error {

	selector := bson.M{"_id": villageweatherdata.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": villageweatherdata}
	_, err := ctx.DB.Collection(constants.COLLECTIONVILLAGEWEATHERDATA).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterVillageWeatherData : ""
func (d *Daos) FilterVillageWeatherData(ctx *models.Context, villageweatherdatafilter *models.VillageWeatherDataFilter, pagination *models.Pagination) ([]models.RefVillageWeatherData, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if villageweatherdatafilter != nil {

		if len(villageweatherdatafilter.ActiveStatus) > 0 {
			query = append(query, bson.M{"activeStatus": bson.M{"$in": villageweatherdatafilter.ActiveStatus}})
		}
		if len(villageweatherdatafilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": villageweatherdatafilter.Status}})
		}
		if len(villageweatherdatafilter.Village) > 0 {
			query = append(query, bson.M{"Village": bson.M{"$in": villageweatherdatafilter.Village}})
		}
		if len(villageweatherdatafilter.Village) > 0 {
			query = append(query, bson.M{"Village": bson.M{"$in": villageweatherdatafilter.Village}})
		}
		//Regex
		if villageweatherdatafilter.SearchBox.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: villageweatherdatafilter.SearchBox.Name, Options: "xi"}})
		}
	}
	if villageweatherdatafilter.DateRange != nil {
		//var sd,ed time.villageweatherdatafilter
		if villageweatherdatafilter.DateRange.From != nil {
			sd := time.Date(villageweatherdatafilter.DateRange.From.Year(), villageweatherdatafilter.DateRange.From.Month(), villageweatherdatafilter.DateRange.From.Day(), 0, 0, 0, 0, villageweatherdatafilter.DateRange.From.Location())
			ed := time.Date(villageweatherdatafilter.DateRange.From.Year(), villageweatherdatafilter.DateRange.From.Month(), villageweatherdatafilter.DateRange.From.Day(), 23, 59, 59, 0, villageweatherdatafilter.DateRange.From.Location())
			if villageweatherdatafilter.DateRange.To != nil {
				ed = time.Date(villageweatherdatafilter.DateRange.To.Year(), villageweatherdatafilter.DateRange.To.Month(), villageweatherdatafilter.DateRange.To.Day(), 23, 59, 59, 0, villageweatherdatafilter.DateRange.To.Location())
			}
			query = append(query, bson.M{"date": bson.M{"$gte": sd, "$lte": ed}})

		}
	}
	mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"date": 1}})

	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	// if villageweatherdatafilter != nil {
	// 	if villageweatherdatafilter.SortBy != "" {
	// 		mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{villageweatherdatafilter.SortBy: villageweatherdatafilter.SortOrder}})

	// 	}

	// }

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONVILLAGEWEATHERDATA).CountDocuments(ctx.CTX, func() bson.M {
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
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVILLAGE, "Village", "_id", "ref.Village", "ref.Village")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("Villageweatherdata query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONVILLAGEWEATHERDATA).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var villageweatherdatas []models.RefVillageWeatherData
	if err = cursor.All(context.TODO(), &villageweatherdatas); err != nil {
		return nil, err
	}
	return villageweatherdatas, nil
}

//EnableVillageWeatherData :""
func (d *Daos) EnableVillageWeatherData(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.VILLAGEWEATHERDATASTATUSACTIVE, "activeStatus": true}}
	_, err = ctx.DB.Collection(constants.COLLECTIONVILLAGEWEATHERDATA).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableVillageWeatherData :""
func (d *Daos) DisableVillageWeatherData(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.VILLAGEWEATHERDATASTATUSDISABLE, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONVILLAGEWEATHERDATA).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteVillageWeatherData :""
func (d *Daos) DeleteVillageWeatherData(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.VILLAGEWEATHERDATASTATUSDELETED, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONVILLAGEWEATHERDATA).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
func (d *Daos) GetSingleVillageWeatherDataWithCurrentDate(ctx *models.Context, UniqueID string) (*models.RefVillageWeatherData, error) {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return nil, err
	}
	t := time.Now()
	fmt.Println("Village time==>", t)
	query := []bson.M{}
	query = append(query, bson.M{"Village": id})
	query = append(query, bson.M{"uniqueId": fmt.Sprintf("%v_%v_%v", t.Day(), t.Month().String(), t.Year())})
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})

	d.Shared.BsonToJSONPrintTag("Villageweatherdata query =>", mainPipeline)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONVILLAGEWEATHERDATA).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var villageweatherdats []models.RefVillageWeatherData
	var villageweatherdata *models.RefVillageWeatherData
	if err = cursor.All(ctx.CTX, &villageweatherdats); err != nil {
		return nil, err
	}
	if len(villageweatherdats) > 0 {
		villageweatherdata = &villageweatherdats[0]
	}
	return villageweatherdata, nil
}
