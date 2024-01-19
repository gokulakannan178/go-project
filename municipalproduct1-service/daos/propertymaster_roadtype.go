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

//SaveRoadType :""
func (d *Daos) SaveRoadType(ctx *models.Context, roadType *models.RoadType) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONROADTYPE).InsertOne(ctx.CTX, roadType)
	return err
}

//GetSingleRoadType : ""
func (d *Daos) GetSingleRoadType(ctx *models.Context, UniqueID string) (*models.RefRoadType, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONROADTYPE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var roadTypes []models.RefRoadType
	var roadType *models.RefRoadType
	if err = cursor.All(ctx.CTX, &roadTypes); err != nil {
		return nil, err
	}
	if len(roadTypes) > 0 {
		roadType = &roadTypes[0]
	}
	return roadType, nil
}

//UpdateRoadType : ""
func (d *Daos) UpdateRoadType(ctx *models.Context, roadType *models.RoadType) error {
	selector := bson.M{"uniqueId": roadType.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": roadType, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONROADTYPE).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterRoadType : ""
func (d *Daos) FilterRoadType(ctx *models.Context, roadTypefilter *models.RoadTypeFilter, pagination *models.Pagination) ([]models.RefRoadType, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if roadTypefilter != nil {

		if len(roadTypefilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": roadTypefilter.Status}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONROADTYPE).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("roadType query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONROADTYPE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var roadTypes []models.RefRoadType
	if err = cursor.All(context.TODO(), &roadTypes); err != nil {
		return nil, err
	}
	return roadTypes, nil
}

//EnableRoadType :""
func (d *Daos) EnableRoadType(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.ROADTYPESTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONROADTYPE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableRoadType :""
func (d *Daos) DisableRoadType(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.ROADTYPESTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONROADTYPE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteRoadType :""
func (d *Daos) DeleteRoadType(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.ROADTYPESTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONROADTYPE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
