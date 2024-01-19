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

//SaveFloorType :""
func (d *Daos) SaveFloorType(ctx *models.Context, floorType *models.FloorType) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONFLOORTYPE).InsertOne(ctx.CTX, floorType)
	return err
}

//GetSingleFloorType : ""
func (d *Daos) GetSingleFloorType(ctx *models.Context, UniqueID string) (*models.RefFloorType, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONFLOORTYPE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var floorTypes []models.RefFloorType
	var floorType *models.RefFloorType
	if err = cursor.All(ctx.CTX, &floorTypes); err != nil {
		return nil, err
	}
	if len(floorTypes) > 0 {
		floorType = &floorTypes[0]
	}
	return floorType, nil
}

//UpdateFloorType : ""
func (d *Daos) UpdateFloorType(ctx *models.Context, floorType *models.FloorType) error {
	selector := bson.M{"uniqueId": floorType.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": floorType, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONFLOORTYPE).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterFloorType : ""
func (d *Daos) FilterFloorType(ctx *models.Context, floorTypefilter *models.FloorTypeFilter, pagination *models.Pagination) ([]models.RefFloorType, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if floorTypefilter != nil {

		if len(floorTypefilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": floorTypefilter.Status}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if floorTypefilter.SortBy != "" {
		mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{floorTypefilter.SortBy: floorTypefilter.SortOrder}})
	} else {
		mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"sortOrder": 1}})
	}
	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONFLOORTYPE).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("floorType query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONFLOORTYPE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var floorTypes []models.RefFloorType
	if err = cursor.All(context.TODO(), &floorTypes); err != nil {
		return nil, err
	}
	return floorTypes, nil
}

//EnableFloorType :""
func (d *Daos) EnableFloorType(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.FLOORTYPESTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONFLOORTYPE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableFloorType :""
func (d *Daos) DisableFloorType(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.FLOORTYPESTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONFLOORTYPE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteFloorType :""
func (d *Daos) DeleteFloorType(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.FLOORTYPESTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONFLOORTYPE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
