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
func (d *Daos) SaveBlockWeatherData(ctx *models.Context, blockweatherdata *models.BlockWeatherData) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONBLOCKWEATHERDATA).InsertOne(ctx.CTX, blockweatherdata)
	if err != nil {
		return err
	}
	blockweatherdata.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}
func (d *Daos) SaveBlockWeatherDataWithUpsert(ctx *models.Context, blockweatherdata *models.BlockWeatherData) error {

	opts := options.Update().SetUpsert(true)
	updateQuery := bson.M{"block": blockweatherdata.Block, "uniqueId": blockweatherdata.UniqueID}
	updateData := bson.M{"$set": blockweatherdata}
	_, err := ctx.DB.Collection(constants.COLLECTIONBLOCKWEATHERDATA).UpdateOne(ctx.CTX, updateQuery, updateData, opts)
	return err
}
func (d *Daos) SaveBlockWeatherDataNameWithUpsert(ctx *models.Context, blockweatherdata *models.BlockWeatherData) error {

	opts := options.Update().SetUpsert(true)
	updateQuery := bson.M{"name": blockweatherdata.Name, "uniqueId": blockweatherdata.UniqueID, "district": blockweatherdata.District, "state": blockweatherdata.State}
	updateData := bson.M{"$set": blockweatherdata}
	_, err := ctx.DB.Collection(constants.COLLECTIONBLOCKWEATHERDATA).UpdateOne(ctx.CTX, updateQuery, updateData, opts)
	return err
}

//GetSingleBlockWeatherData : ""
func (d *Daos) GetSingleBlockWeatherData(ctx *models.Context, UniqueID string) (*models.RefBlockWeatherData, error) {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONBLOCKWEATHERDATA).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var blockweatherdatas []models.RefBlockWeatherData
	var blockweatherdata *models.RefBlockWeatherData
	if err = cursor.All(ctx.CTX, &blockweatherdatas); err != nil {
		return nil, err
	}
	if len(blockweatherdatas) > 0 {
		blockweatherdata = &blockweatherdatas[0]
	}
	return blockweatherdata, nil
}

//UpdateBlockWeatherData : ""
func (d *Daos) UpdateBlockWeatherData(ctx *models.Context, blockweatherdata *models.BlockWeatherData) error {

	selector := bson.M{"_id": blockweatherdata.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": blockweatherdata}
	_, err := ctx.DB.Collection(constants.COLLECTIONBLOCKWEATHERDATA).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterBlockWeatherData : ""
func (d *Daos) FilterBlockWeatherData(ctx *models.Context, blockweatherdatafilter *models.BlockWeatherDataFilter, pagination *models.Pagination) ([]models.RefBlockWeatherData, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if blockweatherdatafilter != nil {

		if len(blockweatherdatafilter.ActiveStatus) > 0 {
			query = append(query, bson.M{"activeStatus": bson.M{"$in": blockweatherdatafilter.ActiveStatus}})
		}
		if len(blockweatherdatafilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": blockweatherdatafilter.Status}})
		}
		if len(blockweatherdatafilter.Block) > 0 {
			query = append(query, bson.M{"Block": bson.M{"$in": blockweatherdatafilter.Block}})
		}
		if len(blockweatherdatafilter.Block) > 0 {
			query = append(query, bson.M{"Block": bson.M{"$in": blockweatherdatafilter.Block}})
		}
		//Regex
		if blockweatherdatafilter.SearchBox.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: blockweatherdatafilter.SearchBox.Name, Options: "xi"}})
		}
	}
	if blockweatherdatafilter.DateRange != nil {
		//var sd,ed time.blockweatherdatafilter
		if blockweatherdatafilter.DateRange.From != nil {
			sd := time.Date(blockweatherdatafilter.DateRange.From.Year(), blockweatherdatafilter.DateRange.From.Month(), blockweatherdatafilter.DateRange.From.Day(), 0, 0, 0, 0, blockweatherdatafilter.DateRange.From.Location())
			ed := time.Date(blockweatherdatafilter.DateRange.From.Year(), blockweatherdatafilter.DateRange.From.Month(), blockweatherdatafilter.DateRange.From.Day(), 23, 59, 59, 0, blockweatherdatafilter.DateRange.From.Location())
			if blockweatherdatafilter.DateRange.To != nil {
				ed = time.Date(blockweatherdatafilter.DateRange.To.Year(), blockweatherdatafilter.DateRange.To.Month(), blockweatherdatafilter.DateRange.To.Day(), 23, 59, 59, 0, blockweatherdatafilter.DateRange.To.Location())
			}
			query = append(query, bson.M{"date": bson.M{"$gte": sd, "$lte": ed}})

		}
	}
	mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"date": 1}})

	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	// if blockweatherdatafilter != nil {
	// 	if blockweatherdatafilter.SortBy != "" {
	// 		mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{blockweatherdatafilter.SortBy: blockweatherdatafilter.SortOrder}})

	// 	}

	// }

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONBLOCKWEATHERDATA).CountDocuments(ctx.CTX, func() bson.M {
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
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONBLOCK, "Block", "_id", "ref.Block", "ref.Block")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("Blockweatherdata query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONBLOCKWEATHERDATA).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var blockweatherdatas []models.RefBlockWeatherData
	if err = cursor.All(context.TODO(), &blockweatherdatas); err != nil {
		return nil, err
	}
	return blockweatherdatas, nil
}

//EnableBlockWeatherData :""
func (d *Daos) EnableBlockWeatherData(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.BLOCKWEATHERDATASTATUSACTIVE, "activeStatus": true}}
	_, err = ctx.DB.Collection(constants.COLLECTIONBLOCKWEATHERDATA).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableBlockWeatherData :""
func (d *Daos) DisableBlockWeatherData(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.BLOCKWEATHERDATASTATUSDISABLE, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONBLOCKWEATHERDATA).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteBlockWeatherData :""
func (d *Daos) DeleteBlockWeatherData(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.BLOCKWEATHERDATASTATUSDELETED, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONBLOCKWEATHERDATA).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
func (d *Daos) GetSingleBlockWeatherDataWithCurrentDate(ctx *models.Context, UniqueID string) (*models.RefBlockWeatherData, error) {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return nil, err
	}
	t := time.Now()
	fmt.Println("Block time==>", t)
	query := []bson.M{}
	query = append(query, bson.M{"Block": id})
	query = append(query, bson.M{"uniqueId": fmt.Sprintf("%v_%v_%v", t.Day(), t.Month().String(), t.Year())})
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})

	d.Shared.BsonToJSONPrintTag("Blockweatherdata query =>", mainPipeline)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONBLOCKWEATHERDATA).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var blockweatherdatas []models.RefBlockWeatherData
	var blockweatherdata *models.RefBlockWeatherData
	if err = cursor.All(ctx.CTX, &blockweatherdatas); err != nil {
		return nil, err
	}
	if len(blockweatherdatas) > 0 {
		blockweatherdata = &blockweatherdatas[0]
	}
	return blockweatherdata, nil
}

func (d *Daos) GetBlockWeatherDataByBlockId(ctx *models.Context, UniqueID string) ([]models.RefBlockWeatherData, error) {
	// id, err := primitive.ObjectIDFromHex(UniqueID)
	// if err != nil {
	// 	return nil, err
	// }
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"name": UniqueID}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONBLOCKWEATHERDATA).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var blockweatherdatas []models.RefBlockWeatherData
	if err = cursor.All(context.TODO(), &blockweatherdatas); err != nil {
		return nil, err
	}
	return blockweatherdatas, nil
}
