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

//SaveFloorRatableArea :""
func (d *Daos) SaveFloorRatableArea(ctx *models.Context, floorRatableArea *models.FloorRatableArea) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONFLOORRATABLEAREA).InsertOne(ctx.CTX, floorRatableArea)
	return err
}

//GetSingleFloorRatableArea : ""
func (d *Daos) GetSingleFloorRatableArea(ctx *models.Context, UniqueID string) (*models.RefFloorRatableArea, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONFLOORRATABLEAREA).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var floorRatableAreas []models.RefFloorRatableArea
	var floorRatableArea *models.RefFloorRatableArea
	if err = cursor.All(ctx.CTX, &floorRatableAreas); err != nil {
		return nil, err
	}
	if len(floorRatableAreas) > 0 {
		floorRatableArea = &floorRatableAreas[0]
	}
	return floorRatableArea, nil
}

//UpdateFloorRatableArea : ""
func (d *Daos) UpdateFloorRatableArea(ctx *models.Context, floorRatableArea *models.FloorRatableArea) error {
	selector := bson.M{"uniqueId": floorRatableArea.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": floorRatableArea, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONFLOORRATABLEAREA).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterFloorRatableArea : ""
func (d *Daos) FilterFloorRatableArea(ctx *models.Context, floorRatableAreafilter *models.FloorRatableAreaFilter, pagination *models.Pagination) ([]models.RefFloorRatableArea, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if floorRatableAreafilter != nil {

		if len(floorRatableAreafilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": floorRatableAreafilter.Status}})
		}
		if len(floorRatableAreafilter.UniqueID) > 0 {
			query = append(query, bson.M{"uniqueId": bson.M{"$in": floorRatableAreafilter.UniqueID}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if floorRatableAreafilter != nil {
		if floorRatableAreafilter.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{floorRatableAreafilter.SortBy: floorRatableAreafilter.SortOrder}})

		}
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONFLOORRATABLEAREA).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("floorRatableArea query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONFLOORRATABLEAREA).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var floorRatableAreas []models.RefFloorRatableArea
	if err = cursor.All(context.TODO(), &floorRatableAreas); err != nil {
		return nil, err
	}
	return floorRatableAreas, nil
}

//EnableFloorRatableArea :""
func (d *Daos) EnableFloorRatableArea(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.FLOORRATABLEAREASTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONFLOORRATABLEAREA).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableFloorRatableArea :""
func (d *Daos) DisableFloorRatableArea(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.FLOORRATABLEAREASTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONFLOORRATABLEAREA).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteFloorRatableArea :""
func (d *Daos) DeleteFloorRatableArea(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.FLOORRATABLEAREASTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONFLOORRATABLEAREA).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
