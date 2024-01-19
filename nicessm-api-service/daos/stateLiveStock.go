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

//SaveStateLiveStock :""
func (d *Daos) SaveStateLiveStock(ctx *models.Context, StateLiveStock *models.StateLiveStock) error {

	res, err := ctx.DB.Collection(constants.COLLECTIONSTATELIVESTOCK).InsertOne(ctx.CTX, StateLiveStock)
	if err != nil {
		return err
	}
	StateLiveStock.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//GetSingleStateLiveStock : ""
func (d *Daos) GetSingleStateLiveStock(ctx *models.Context, code string) (*models.RefStateLiveStock, error) {
	id, err := primitive.ObjectIDFromHex(code)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})
	//Lookups
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "state", "_id", "ref.state", "ref.state")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONCOMMODITY, "commodity", "_id", "ref.commodity", "ref.commodity")...)
	// //Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSTATELIVESTOCK).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var StateLiveStocks []models.RefStateLiveStock
	var StateLiveStock *models.RefStateLiveStock
	if err = cursor.All(ctx.CTX, &StateLiveStocks); err != nil {
		return nil, err
	}
	if len(StateLiveStocks) > 0 {
		StateLiveStock = &StateLiveStocks[0]
	}
	return StateLiveStock, nil
}

//UpdateStateLiveStock : ""
func (d *Daos) UpdateStateLiveStock(ctx *models.Context, StateLiveStock *models.StateLiveStock) error {

	selector := bson.M{"_id": StateLiveStock.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": StateLiveStock, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONSTATELIVESTOCK).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterStateLiveStock : ""
func (d *Daos) FilterStateLiveStock(ctx *models.Context, StateLiveStockfilter *models.StateLiveStockFilter, pagination *models.Pagination) ([]models.RefStateLiveStock, error) {
	mainPipeline := []bson.M{}
	paginationPipeline := []bson.M{}
	query := []bson.M{}
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONCOMMODITY, "commodity", "_id", "ref.commodity", "ref.commodity")...)
	if StateLiveStockfilter != nil {
		if len(StateLiveStockfilter.ActiveStatus) > 0 {
			query = append(query, bson.M{"activeStatus": bson.M{"$in": StateLiveStockfilter.ActiveStatus}})
		}
		if len(StateLiveStockfilter.State) > 0 {
			query = append(query, bson.M{"state": bson.M{"$in": StateLiveStockfilter.State}})
		}
		if len(StateLiveStockfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": StateLiveStockfilter.Status}})
		}
		// if StateLiveStockfilter.Regex.Name != "" {
		// 	query = append(query, bson.M{"name": primitive.Regex{Pattern: StateLiveStockfilter.Regex.Name, Options: "xi"}})
		// }
		if StateLiveStockfilter.Regex.NameInLocalLanguage != "" {
			query = append(query, bson.M{"nameInLocalLanguage": primitive.Regex{Pattern: StateLiveStockfilter.Regex.NameInLocalLanguage, Options: "xi"}})
		}
		if StateLiveStockfilter.Regex.Name != "" {
			mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"ref.commodity.commonName": primitive.Regex{Pattern: StateLiveStockfilter.Regex.Name, Options: "xi"}}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if StateLiveStockfilter != nil {
		if StateLiveStockfilter.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{StateLiveStockfilter.SortBy: StateLiveStockfilter.SortOrder}})
		}
	}

	//Adding pagination if necessary
	if pagination != nil {
		paginationPipeline = append(paginationPipeline, mainPipeline...)
		paginationPipeline = append(paginationPipeline, bson.M{"$group": bson.M{"_id": nil, "count": bson.M{"$sum": 1}}})
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		d.Shared.BsonToJSONPrintTag("statelivestock pagenation query =>", paginationPipeline)

		//Getting Total count
		paginationCursor, err := ctx.DB.Collection(constants.COLLECTIONSTATELIVESTOCK).Aggregate(ctx.CTX, paginationPipeline, nil)
		if err != nil {
			log.Println("Error in geting pagination - " + err.Error())
		}
		var totalCount int64
		cs := []models.Countstruct{}
		if err = paginationCursor.All(context.TODO(), &cs); err != nil {
			return nil, err
		}
		if len(cs) > 0 {
			totalCount = cs[0].Count
		}
		fmt.Println("count", totalCount)
		pagination.Count = int(totalCount)
		d.Shared.PaginationData(pagination)
	}

	//Lookups
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "state", "_id", "ref.state", "ref.state")...)

	//mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONCOMMODITY, "commodity", "_id", "ref.commodity", "ref.commodity")...)
	// mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONAGROECOLOGICALZONE, "agroEcologicalZones", "_id", "ref.agroEcologicalZones", "ref.agroEcologicalZones")...)
	// //Aggregation
	d.Shared.BsonToJSONPrintTag("StateLiveStock query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSTATELIVESTOCK).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var StateLiveStocks []models.RefStateLiveStock
	if err = cursor.All(context.TODO(), &StateLiveStocks); err != nil {
		return nil, err
	}
	return StateLiveStocks, nil
}

//EnableStateLiveStock :""
func (d *Daos) EnableStateLiveStock(ctx *models.Context, code string) error {
	id, err := primitive.ObjectIDFromHex(code)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.STATELIVESTOCKSTATUSACTIVE, "activeStatus": true}}
	_, err = ctx.DB.Collection(constants.COLLECTIONSTATELIVESTOCK).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableStateLiveStock :""
func (d *Daos) DisableStateLiveStock(ctx *models.Context, code string) error {
	id, err := primitive.ObjectIDFromHex(code)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.STATELIVESTOCKSTATUSDISABLED, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONSTATELIVESTOCK).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteStateLiveStock :""
func (d *Daos) DeleteStateLiveStock(ctx *models.Context, code string) error {
	id, err := primitive.ObjectIDFromHex(code)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.STATELIVESTOCKSTATUSDELETED, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONSTATELIVESTOCK).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
