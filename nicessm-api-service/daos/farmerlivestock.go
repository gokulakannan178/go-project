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
)

//SaveFarmerLiveStock :""
func (d *Daos) SaveFarmerLiveStock(ctx *models.Context, farmerLiveStock *models.FarmerLiveStock) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONFARMERLIVESTOCK).InsertOne(ctx.CTX, farmerLiveStock)
	if err != nil {
		return err
	}
	farmerLiveStock.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//UpdateFarmerLiveStock : ""
func (d *Daos) UpdateFarmerLiveStock(ctx *models.Context, farmerLiveStock *models.FarmerLiveStock) error {

	selector := bson.M{"_id": farmerLiveStock.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": farmerLiveStock}
	_, err := ctx.DB.Collection(constants.COLLECTIONFARMERLIVESTOCK).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//EnableFarmerLiveStock :""
func (d *Daos) EnableFarmerLiveStock(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.FARMERLIVESTOCKSTATUSACTIVE}}
	_, err = ctx.DB.Collection(constants.COLLECTIONFARMERLIVESTOCK).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableFarmerLiveStock :""
func (d *Daos) DisableFarmerLiveStock(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.FARMERLIVESTOCKSTATUSDISABLED}}
	_, err = ctx.DB.Collection(constants.COLLECTIONFARMERLIVESTOCK).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteFarmerLiveStock :""
func (d *Daos) DeleteFarmerLiveStock(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.FARMERLIVESTOCKSTATUSDELETED}}
	_, err = ctx.DB.Collection(constants.COLLECTIONFARMERLIVESTOCK).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//GetSingleFarmerLiveStock : ""
func (d *Daos) GetSingleFarmerLiveStock(ctx *models.Context, UniqueID string) (*models.RefFarmerLiveStock, error) {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})

	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONCOMMODITY, "liveStock", "_id", "ref.liveStock", "ref.liveStock")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONFARMER, "farmer", "_id", "ref.farmer", "ref.farmer")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONCOMMODITYSTAGE, "stage", "_id", "ref.stage", "ref.stage")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONCOMMODITYVARIETY, "veriety", "_id", "ref.veriety", "ref.veriety")...)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONFARMERLIVESTOCK).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var farmerLiveStocks []models.RefFarmerLiveStock
	var farmerLiveStock *models.RefFarmerLiveStock
	if err = cursor.All(ctx.CTX, &farmerLiveStocks); err != nil {
		return nil, err
	}
	if len(farmerLiveStocks) > 0 {
		farmerLiveStock = &farmerLiveStocks[0]
	}
	return farmerLiveStock, nil
}

//FilterFarmerLiveStock : ""
func (d *Daos) FilterFarmerLiveStock(ctx *models.Context, farmerLiveStockfilter *models.FarmerLiveStockFilter, pagination *models.Pagination) ([]models.RefFarmerLiveStock, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONCOMMODITY, "liveStock", "_id", "ref.liveStock", "ref.liveStock")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONCOMMODITYCATEGORY, "ref.liveStock.category", "_id", "ref.category", "ref.category")...)
	query := []bson.M{}
	if farmerLiveStockfilter != nil {
		if len(farmerLiveStockfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": farmerLiveStockfilter.Status}})
		}
		if len(farmerLiveStockfilter.Farmer) > 0 {
			query = append(query, bson.M{"farmer": bson.M{"$in": farmerLiveStockfilter.Farmer}})
		}
		if len(farmerLiveStockfilter.LiveStock) > 0 {
			query = append(query, bson.M{"liveStock": bson.M{"$in": farmerLiveStockfilter.LiveStock}})
		}
		if len(farmerLiveStockfilter.Veriety) > 0 {
			query = append(query, bson.M{"veriety": bson.M{"$in": farmerLiveStockfilter.Veriety}})
		}
		if len(farmerLiveStockfilter.Category) > 0 {
			query = append(query, bson.M{"ref.liveStock.category": bson.M{"$in": farmerLiveStockfilter.Category}})
		}
		if farmerLiveStockfilter.SearchBox.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: farmerLiveStockfilter.SearchBox.Name, Options: "xi"}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONFARMER, "farmer", "_id", "ref.farmer", "ref.farmer")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONCOMMODITYSTAGE, "stage", "_id", "ref.stage", "ref.stage")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONCOMMODITYVARIETY, "veriety", "_id", "ref.veriety", "ref.veriety")...)

	//Adding pagination if necessary
	if pagination != nil {
		paginationPipeline := []bson.M{}
		paginationPipeline = append(paginationPipeline, mainPipeline...)
		paginationPipeline = append(paginationPipeline, bson.M{"$group": bson.M{"_id": nil, "count": bson.M{"$sum": 1}}})
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		d.Shared.BsonToJSONPrintTag("farmercrop Pagination quary =>", paginationPipeline)

		countcursor, err := ctx.DB.Collection(constants.COLLECTIONFARMERLIVESTOCK).Aggregate(ctx.CTX, paginationPipeline, nil)
		if err != nil {
			log.Println("Error in geting pagination - " + err.Error())
		}
		countstruct := []models.Countstruct{}
		err = countcursor.All(ctx.CTX, &countstruct)
		if err != nil {
			return nil, err
		}
		var totalCount int64
		if len(countstruct) > 0 {
			totalCount = countstruct[0].Count
		}
		fmt.Println("count", totalCount)
		pagination.Count = int(totalCount)
		d.Shared.PaginationData(pagination)
	}
	//Aggregation
	d.Shared.BsonToJSONPrintTag("FarmerLiveStock query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONFARMERLIVESTOCK).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var FarmerLiveStocks []models.RefFarmerLiveStock
	if err = cursor.All(context.TODO(), &FarmerLiveStocks); err != nil {
		return nil, err
	}
	return FarmerLiveStocks, nil
}
func (d *Daos) GetFarmerLiveStockCount(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := []bson.M{}
	query = append(query, bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{

		{"$eq": []string{"$status", constants.FARMERLIVESTOCKSTATUSACTIVE}},
		{"$eq": []interface{}{"$farmer", id}},
	}}}})
	query = append(query, bson.M{"$group": bson.M{
		"_id":   nil,
		"count": bson.M{"$sum": 1},
	}})
	d.Shared.BsonToJSONPrintTag("GetFarmerLiveStockCount query =>", query)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONFARMERLIVESTOCK).Aggregate(ctx.CTX, query, nil)
	if err != nil {
		return err
	}
	var farmers []models.FarmerLiveStockCount
	var farmer *models.FarmerLiveStockCount
	if err = cursor.All(context.TODO(), &farmers); err != nil {
		return err
	}
	if len(farmers) > 0 {
		farmer = &farmers[0]
	} else {
		farmer = new(models.FarmerLiveStockCount)
	}
	query2 := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"liveStockCount": farmer.Count}}
	_, err = ctx.DB.Collection(constants.COLLECTIONFARMER).UpdateOne(ctx.CTX, query2, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return nil
}
