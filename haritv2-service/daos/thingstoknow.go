package daos

import (
	"context"
	"errors"
	"fmt"
	"haritv2-service/constants"
	"haritv2-service/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

//SaveThingsToKnow :""
func (d *Daos) SaveThingsToKnow(ctx *models.Context, thingskw *models.ThingsToKnow) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONTHINGSTOKNOW).InsertOne(ctx.CTX, thingskw)
	return err
}

//GetSingleThingsToKnow : ""
func (d *Daos) GetSingleThingsToKnow(ctx *models.Context, uniqueID string) (*models.ThingsToKnow, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONTHINGSTOKNOW).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var thingskws []models.ThingsToKnow
	var thingskw *models.ThingsToKnow
	if err = cursor.All(ctx.CTX, &thingskws); err != nil {
		return nil, err
	}
	if len(thingskws) > 0 {
		thingskw = &thingskws[0]
	}
	return thingskw, nil
}

//UpdateThingsToKnow : ""
func (d *Daos) UpdateThingsToKnow(ctx *models.Context, thingskw *models.ThingsToKnow) error {
	selector := bson.M{"uniqueId": thingskw.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": thingskw, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONTHINGSTOKNOW).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterThingsToKnow : ""
func (d *Daos) FilterThingsToKnow(ctx *models.Context, thingskwfilter *models.FilterThingsToKnow, pagination *models.Pagination) ([]models.ThingsToKnow, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if thingskwfilter != nil {

		if len(thingskwfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": thingskwfilter.Status}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONTHINGSTOKNOW).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("thingskw query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONTHINGSTOKNOW).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var thingskws []models.ThingsToKnow
	if err = cursor.All(context.TODO(), &thingskws); err != nil {
		return nil, err
	}
	return thingskws, nil
}

//EnableThingsToKnow :""
func (d *Daos) EnableThingsToKnow(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.THINGSTOKNOWSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONTHINGSTOKNOW).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableThingsToKnow :""
func (d *Daos) DisableThingsToKnow(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.THINGSTOKNOWSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONTHINGSTOKNOW).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteThingsToKnow :""
func (d *Daos) DeleteThingsToKnow(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.THINGSTOKNOWSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONTHINGSTOKNOW).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
