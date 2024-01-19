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

//SaveRoofType :""
func (d *Daos) SaveRoofType(ctx *models.Context, RoofType *models.RoofType) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONROOFTYPE).InsertOne(ctx.CTX, RoofType)
	return err
}

//GetSingleRoofType : ""
func (d *Daos) GetSingleRoofType(ctx *models.Context, UniqueID string) (*models.RefRoofType, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONROOFTYPE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var roofTypes []models.RefRoofType
	var RoofType *models.RefRoofType
	if err = cursor.All(ctx.CTX, &roofTypes); err != nil {
		return nil, err
	}
	if len(roofTypes) > 0 {
		RoofType = &roofTypes[0]
	}
	return RoofType, nil
}

//UpdateRoofType : ""
func (d *Daos) UpdateRoofType(ctx *models.Context, RoofType *models.RoofType) error {
	selector := bson.M{"uniqueId": RoofType.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": RoofType, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONROOFTYPE).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterRoofType : ""
func (d *Daos) FilterRoofType(ctx *models.Context, roofTypefilter *models.RoofTypeFilter, pagination *models.Pagination) ([]models.RefRoofType, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if roofTypefilter != nil {

		if len(roofTypefilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": roofTypefilter.Status}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONROOFTYPE).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("RoofType query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONROOFTYPE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var roofTypes []models.RefRoofType
	if err = cursor.All(context.TODO(), &roofTypes); err != nil {
		return nil, err
	}
	return roofTypes, nil
}

//EnableRoofType :""
func (d *Daos) EnableRoofType(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.ROOFTYPESTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONROOFTYPE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableRoofType :""
func (d *Daos) DisableRoofType(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.ROOFTYPESTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONROOFTYPE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteRoofType :""
func (d *Daos) DeleteRoofType(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.ROOFTYPESTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONROOFTYPE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
