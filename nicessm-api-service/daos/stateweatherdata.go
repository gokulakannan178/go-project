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
func (d *Daos) SaveStateWeatherData(ctx *models.Context, stateweatherdata *models.StateWeatherData) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONSTATEWEATHERDATA).InsertOne(ctx.CTX, stateweatherdata)
	if err != nil {
		return err
	}
	stateweatherdata.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}
func (d *Daos) SaveStateWeatherData2(ctx *models.Context, stateweatherdata *models.StateWeatherData) error {

	opts := options.Update().SetUpsert(true)
	updateQuery := bson.M{"state": stateweatherdata.State, "uniqueId": stateweatherdata.UniqueID}
	updateData := bson.M{"$set": stateweatherdata}
	_, err := ctx.DB.Collection(constants.COLLECTIONSTATEWEATHERDATA).UpdateOne(ctx.CTX, updateQuery, updateData, opts)
	return err
}

//GetSingleStateWeatherData : ""
func (d *Daos) GetSingleStateWeatherData(ctx *models.Context, UniqueID string) (*models.RefStateWeatherData, error) {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSTATEWEATHERDATA).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var stateweatherdatas []models.RefStateWeatherData
	var stateweatherdata *models.RefStateWeatherData
	if err = cursor.All(ctx.CTX, &stateweatherdatas); err != nil {
		return nil, err
	}
	if len(stateweatherdatas) > 0 {
		stateweatherdata = &stateweatherdatas[0]
	}
	return stateweatherdata, nil
}

//UpdateStateWeatherData : ""
func (d *Daos) UpdateStateWeatherData(ctx *models.Context, stateweatherdata *models.StateWeatherData) error {

	selector := bson.M{"_id": stateweatherdata.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": stateweatherdata}
	_, err := ctx.DB.Collection(constants.COLLECTIONSTATEWEATHERDATA).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterStateWeatherData : ""
func (d *Daos) FilterStateWeatherData(ctx *models.Context, stateweatherdatafilter *models.StateWeatherDataFilter, pagination *models.Pagination) ([]models.RefStateWeatherData, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if stateweatherdatafilter != nil {

		if len(stateweatherdatafilter.ActiveStatus) > 0 {
			query = append(query, bson.M{"activeStatus": bson.M{"$in": stateweatherdatafilter.ActiveStatus}})
		}
		if len(stateweatherdatafilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": stateweatherdatafilter.Status}})
		}
		if len(stateweatherdatafilter.State) > 0 {
			query = append(query, bson.M{"state": bson.M{"$in": stateweatherdatafilter.State}})
		}
		if len(stateweatherdatafilter.State) > 0 {
			query = append(query, bson.M{"state": bson.M{"$in": stateweatherdatafilter.State}})
		}
		//Regex
		if stateweatherdatafilter.SearchBox.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: stateweatherdatafilter.SearchBox.Name, Options: "xi"}})
		}
	}
	if stateweatherdatafilter.DateRange != nil {
		//var sd,ed time.stateweatherdatafilter
		if stateweatherdatafilter.DateRange.From != nil {
			sd := time.Date(stateweatherdatafilter.DateRange.From.Year(), stateweatherdatafilter.DateRange.From.Month(), stateweatherdatafilter.DateRange.From.Day(), 0, 0, 0, 0, stateweatherdatafilter.DateRange.From.Location())
			ed := time.Date(stateweatherdatafilter.DateRange.From.Year(), stateweatherdatafilter.DateRange.From.Month(), stateweatherdatafilter.DateRange.From.Day(), 23, 59, 59, 0, stateweatherdatafilter.DateRange.From.Location())
			if stateweatherdatafilter.DateRange.To != nil {
				ed = time.Date(stateweatherdatafilter.DateRange.To.Year(), stateweatherdatafilter.DateRange.To.Month(), stateweatherdatafilter.DateRange.To.Day(), 23, 59, 59, 0, stateweatherdatafilter.DateRange.To.Location())
			}
			query = append(query, bson.M{"date": bson.M{"$gte": sd, "$lte": ed}})

		}
	}
	mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"date": 1}})

	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	// if stateweatherdatafilter != nil {
	// 	if stateweatherdatafilter.SortBy != "" {
	// 		mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{stateweatherdatafilter.SortBy: stateweatherdatafilter.SortOrder}})

	// 	}

	// }

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONSTATEWEATHERDATA).CountDocuments(ctx.CTX, func() bson.M {
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
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "state", "_id", "ref.state", "ref.state")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("stateweatherdata query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSTATEWEATHERDATA).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var stateweatherdatas []models.RefStateWeatherData
	if err = cursor.All(context.TODO(), &stateweatherdatas); err != nil {
		return nil, err
	}
	return stateweatherdatas, nil
}

//EnableStateWeatherData :""
func (d *Daos) EnableStateWeatherData(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.STATEWEATHERDATASTATUSACTIVE, "activeStatus": true}}
	_, err = ctx.DB.Collection(constants.COLLECTIONSTATEWEATHERDATA).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableStateWeatherData :""
func (d *Daos) DisableStateWeatherData(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.STATEWEATHERDATASTATUSDISABLE, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONSTATEWEATHERDATA).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteStateWeatherData :""
func (d *Daos) DeleteStateWeatherData(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.STATEWEATHERDATASTATUSDELETED, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONSTATEWEATHERDATA).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
func (d *Daos) GetSingleStateWeatherDataWithCureentDate(ctx *models.Context, UniqueID string) (*models.RefStateWeatherData, error) {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return nil, err
	}
	t := time.Now()
	fmt.Println("state time==>", t)
	query := []bson.M{}
	query = append(query, bson.M{"state": id})
	query = append(query, bson.M{"uniqueId": fmt.Sprintf("%v_%v_%v", t.Day(), t.Month().String(), t.Year())})
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})

	d.Shared.BsonToJSONPrintTag("stateweatherdata query =>", mainPipeline)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSTATEWEATHERDATA).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var stateweatherdatas []models.RefStateWeatherData
	var stateweatherdata *models.RefStateWeatherData
	if err = cursor.All(ctx.CTX, &stateweatherdatas); err != nil {
		return nil, err
	}
	if len(stateweatherdatas) > 0 {
		stateweatherdata = &stateweatherdatas[0]
	}
	return stateweatherdata, nil
}
