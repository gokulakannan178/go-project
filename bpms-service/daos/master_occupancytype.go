package daos

import (
	"bpms-service/constants"
	"bpms-service/models"
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

//SaveOccupancyType :""
func (d *Daos) SaveOccupancyType(ctx *models.Context, OccupancyType *models.OccupancyType) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONOCCUPANCYTYPE).InsertOne(ctx.CTX, OccupancyType)
	return err
}

//GetSingleOccupancyType : ""
func (d *Daos) GetSingleOccupancyType(ctx *models.Context, UniqueID string) (*models.RefOccupancyType, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONOCCUPANCYTYPE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var occupancyTypes []models.RefOccupancyType
	var OccupancyType *models.RefOccupancyType
	if err = cursor.All(ctx.CTX, &occupancyTypes); err != nil {
		return nil, err
	}
	if len(occupancyTypes) > 0 {
		OccupancyType = &occupancyTypes[0]
	}
	return OccupancyType, nil
}

//UpdateOccupancyType : ""
func (d *Daos) UpdateOccupancyType(ctx *models.Context, OccupancyType *models.OccupancyType) error {
	selector := bson.M{"uniqueId": OccupancyType.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": OccupancyType, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONOCCUPANCYTYPE).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterOccupancyType : ""
func (d *Daos) FilterOccupancyType(ctx *models.Context, occupancyTypefilter *models.OccupancyTypeFilter, pagination *models.Pagination) ([]models.RefOccupancyType, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if occupancyTypefilter != nil {

		if len(occupancyTypefilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": occupancyTypefilter.Status}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONOCCUPANCYTYPE).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("OccupancyType query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONOCCUPANCYTYPE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var occupancyTypes []models.RefOccupancyType
	if err = cursor.All(context.TODO(), &occupancyTypes); err != nil {
		return nil, err
	}
	return occupancyTypes, nil
}

//EnableOccupancyType :""
func (d *Daos) EnableOccupancyType(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.OCCUPANCYTYPESTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONOCCUPANCYTYPE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableOccupancyType :""
func (d *Daos) DisableOccupancyType(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.OCCUPANCYTYPESTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONOCCUPANCYTYPE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteOccupancyType :""
func (d *Daos) DeleteOccupancyType(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.OCCUPANCYTYPESTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONOCCUPANCYTYPE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
