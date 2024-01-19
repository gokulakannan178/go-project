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
func (d *Daos) SaveGramPanchayatWeatherData(ctx *models.Context, gramPanchayatweatherdata *models.GramPanchayatWeatherData) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONGRAMPANCHAYATWEATHERDATA).InsertOne(ctx.CTX, gramPanchayatweatherdata)
	if err != nil {
		return err
	}
	gramPanchayatweatherdata.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}
func (d *Daos) SaveGramPanchayatWeatherDataWithUpsert(ctx *models.Context, gramPanchayatweatherdata *models.GramPanchayatWeatherData) error {

	opts := options.Update().SetUpsert(true)
	updateQuery := bson.M{"GramPanchayat": gramPanchayatweatherdata.GramPanchayat, "uniqueId": gramPanchayatweatherdata.UniqueID}
	updateData := bson.M{"$set": gramPanchayatweatherdata}
	_, err := ctx.DB.Collection(constants.COLLECTIONGRAMPANCHAYATWEATHERDATA).UpdateOne(ctx.CTX, updateQuery, updateData, opts)
	return err
}

//GetSingleGramPanchayatWeatherData : ""
func (d *Daos) GetSingleGramPanchayatWeatherData(ctx *models.Context, UniqueID string) (*models.RefGramPanchayatWeatherData, error) {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONGRAMPANCHAYATWEATHERDATA).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var gramPanchayatweatherdatas []models.RefGramPanchayatWeatherData
	var gramPanchayatweatherdata *models.RefGramPanchayatWeatherData
	if err = cursor.All(ctx.CTX, &gramPanchayatweatherdatas); err != nil {
		return nil, err
	}
	if len(gramPanchayatweatherdatas) > 0 {
		gramPanchayatweatherdata = &gramPanchayatweatherdatas[0]
	}
	return gramPanchayatweatherdata, nil
}

//UpdateGramPanchayatWeatherData : ""
func (d *Daos) UpdateGramPanchayatWeatherData(ctx *models.Context, gramPanchayatweatherdata *models.GramPanchayatWeatherData) error {

	selector := bson.M{"_id": gramPanchayatweatherdata.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": gramPanchayatweatherdata}
	_, err := ctx.DB.Collection(constants.COLLECTIONGRAMPANCHAYATWEATHERDATA).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterGramPanchayatWeatherData : ""
func (d *Daos) FilterGramPanchayatWeatherData(ctx *models.Context, gramPanchayatweatherdatafilter *models.GramPanchayatWeatherDataFilter, pagination *models.Pagination) ([]models.RefGramPanchayatWeatherData, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if gramPanchayatweatherdatafilter != nil {

		if len(gramPanchayatweatherdatafilter.ActiveStatus) > 0 {
			query = append(query, bson.M{"activeStatus": bson.M{"$in": gramPanchayatweatherdatafilter.ActiveStatus}})
		}
		if len(gramPanchayatweatherdatafilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": gramPanchayatweatherdatafilter.Status}})
		}
		if len(gramPanchayatweatherdatafilter.GramPanchayat) > 0 {
			query = append(query, bson.M{"GramPanchayat": bson.M{"$in": gramPanchayatweatherdatafilter.GramPanchayat}})
		}
		if len(gramPanchayatweatherdatafilter.GramPanchayat) > 0 {
			query = append(query, bson.M{"GramPanchayat": bson.M{"$in": gramPanchayatweatherdatafilter.GramPanchayat}})
		}
		//Regex
		if gramPanchayatweatherdatafilter.SearchBox.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: gramPanchayatweatherdatafilter.SearchBox.Name, Options: "xi"}})
		}
	}
	if gramPanchayatweatherdatafilter.DateRange != nil {
		//var sd,ed time.GramPanchayatweatherdatafilter
		if gramPanchayatweatherdatafilter.DateRange.From != nil {
			sd := time.Date(gramPanchayatweatherdatafilter.DateRange.From.Year(), gramPanchayatweatherdatafilter.DateRange.From.Month(), gramPanchayatweatherdatafilter.DateRange.From.Day(), 0, 0, 0, 0, gramPanchayatweatherdatafilter.DateRange.From.Location())
			ed := time.Date(gramPanchayatweatherdatafilter.DateRange.From.Year(), gramPanchayatweatherdatafilter.DateRange.From.Month(), gramPanchayatweatherdatafilter.DateRange.From.Day(), 23, 59, 59, 0, gramPanchayatweatherdatafilter.DateRange.From.Location())
			if gramPanchayatweatherdatafilter.DateRange.To != nil {
				ed = time.Date(gramPanchayatweatherdatafilter.DateRange.To.Year(), gramPanchayatweatherdatafilter.DateRange.To.Month(), gramPanchayatweatherdatafilter.DateRange.To.Day(), 23, 59, 59, 0, gramPanchayatweatherdatafilter.DateRange.To.Location())
			}
			query = append(query, bson.M{"date": bson.M{"$gte": sd, "$lte": ed}})

		}
	}
	mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"date": 1}})

	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	// if GramPanchayatweatherdatafilter != nil {
	// 	if GramPanchayatweatherdatafilter.SortBy != "" {
	// 		mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{GramPanchayatweatherdatafilter.SortBy: GramPanchayatweatherdatafilter.SortOrder}})

	// 	}

	// }

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONGRAMPANCHAYATWEATHERDATA).CountDocuments(ctx.CTX, func() bson.M {
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
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONGRAMPANCHAYAT, "GramPanchayat", "_id", "ref.GramPanchayat", "ref.GramPanchayat")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("GramPanchayatweatherdata query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONGRAMPANCHAYATWEATHERDATA).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var gramPanchayatweatherdatas []models.RefGramPanchayatWeatherData
	if err = cursor.All(context.TODO(), &gramPanchayatweatherdatas); err != nil {
		return nil, err
	}
	return gramPanchayatweatherdatas, nil
}

//EnableGramPanchayatWeatherData :""
func (d *Daos) EnableGramPanchayatWeatherData(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.GRAMPANCHAYATWEATHERDATASTATUSACTIVE, "activeStatus": true}}
	_, err = ctx.DB.Collection(constants.COLLECTIONGRAMPANCHAYATWEATHERDATA).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableGramPanchayatWeatherData :""
func (d *Daos) DisableGramPanchayatWeatherData(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.GRAMPANCHAYATWEATHERDATASTATUSDISABLE, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONGRAMPANCHAYATWEATHERDATA).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteGramPanchayatWeatherData :""
func (d *Daos) DeleteGramPanchayatWeatherData(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.GRAMPANCHAYATWEATHERDATASTATUSDELETED, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONGRAMPANCHAYATWEATHERDATA).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
func (d *Daos) GetSingleGramPanchayatWeatherDataWithCurrentDate(ctx *models.Context, UniqueID string) (*models.RefGramPanchayatWeatherData, error) {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return nil, err
	}
	t := time.Now()
	fmt.Println("GramPanchayat time==>", t)
	query := []bson.M{}
	query = append(query, bson.M{"GramPanchayat": id})
	query = append(query, bson.M{"uniqueId": fmt.Sprintf("%v_%v_%v", t.Day(), t.Month().String(), t.Year())})
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})

	d.Shared.BsonToJSONPrintTag("GramPanchayatweatherdata query =>", mainPipeline)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONGRAMPANCHAYATWEATHERDATA).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var gramPanchayatweatherdatas []models.RefGramPanchayatWeatherData
	var gramPanchayatweatherdata *models.RefGramPanchayatWeatherData
	if err = cursor.All(ctx.CTX, &gramPanchayatweatherdatas); err != nil {
		return nil, err
	}
	if len(gramPanchayatweatherdatas) > 0 {
		gramPanchayatweatherdata = &gramPanchayatweatherdatas[0]
	}
	return gramPanchayatweatherdata, nil
}
