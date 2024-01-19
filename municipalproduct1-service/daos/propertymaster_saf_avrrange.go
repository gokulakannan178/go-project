package daos

import (
	"context"
	"errors"
	"fmt"
	"log"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

//SaveAVRRange :""
func (d *Daos) SaveAVRRange(ctx *models.Context, avrRange *models.AVRRange) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONAVRRANGE).InsertOne(ctx.CTX, avrRange)
	return err
}

//GetSingleAVRRange : ""
func (d *Daos) GetSingleAVRRange(ctx *models.Context, UniqueID string) (*models.RefAVRRange, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONAVRRANGE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var avrRanges []models.RefAVRRange
	var avrRange *models.RefAVRRange
	if err = cursor.All(ctx.CTX, &avrRanges); err != nil {
		return nil, err
	}
	if len(avrRanges) > 0 {
		avrRange = &avrRanges[0]
	}
	return avrRange, nil
}

//UpdateAVRRange : ""
func (d *Daos) UpdateAVRRange(ctx *models.Context, avrRange *models.AVRRange) error {
	selector := bson.M{"uniqueId": avrRange.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": avrRange, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONAVRRANGE).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterAVRRange : ""
func (d *Daos) FilterAVRRange(ctx *models.Context, avrRangefilter *models.AVRRangeFilter, pagination *models.Pagination) ([]models.RefAVRRange, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if avrRangefilter != nil {

		if len(avrRangefilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": avrRangefilter.Status}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONAVRRANGE).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("avrRange query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONAVRRANGE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var avrRanges []models.RefAVRRange
	if err = cursor.All(context.TODO(), &avrRanges); err != nil {
		return nil, err
	}
	return avrRanges, nil
}

//EnableAVRRange :""
func (d *Daos) EnableAVRRange(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.AVRRANGESTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONAVRRANGE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableAVRRange :""
func (d *Daos) DisableAVRRange(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.AVRRANGESTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONAVRRANGE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteAVRRange :""
func (d *Daos) DeleteAVRRange(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.AVRRANGESTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONAVRRANGE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

func (d *Daos) GetCGPropertyTaxRate(ctx *models.Context, arv float64, date *time.Duration) (*models.RefAVRRange, error) {
	query := []bson.M{
		bson.M{
			"from": bson.M{"$gte": arv}, "to": bson.M{"$lte": arv}, "doe": bson.M{"$lte": date},
		},
		bson.M{"$sort": bson.M{"doe": -1}},
	}
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONAVRRANGE).Aggregate(ctx.CTX, query, nil)
	if err != nil {
		return nil, err
	}
	var avrRanges []models.RefAVRRange
	var avrRange *models.RefAVRRange
	if err = cursor.All(ctx.CTX, &avrRanges); err != nil {
		return nil, err
	}
	if len(avrRanges) > 0 {
		avrRange = &avrRanges[0]
	} else {
		avrRange = new(models.RefAVRRange)
	}
	return avrRange, nil
}
