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

//SaveUsageType :""
func (d *Daos) SaveUsageType(ctx *models.Context, usageType *models.UsageType) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONUSAGETYPE).InsertOne(ctx.CTX, usageType)
	return err
}

//GetSingleUsageType : ""
func (d *Daos) GetSingleUsageType(ctx *models.Context, UniqueID string) (*models.RefUsageType, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONUSAGETYPE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var usageTypes []models.RefUsageType
	var usageType *models.RefUsageType
	if err = cursor.All(ctx.CTX, &usageTypes); err != nil {
		return nil, err
	}
	if len(usageTypes) > 0 {
		usageType = &usageTypes[0]
	}
	return usageType, nil
}

//UpdateUsageType : ""
func (d *Daos) UpdateUsageType(ctx *models.Context, usageType *models.UsageType) error {
	selector := bson.M{"uniqueId": usageType.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": usageType, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONUSAGETYPE).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterUsageType : ""
func (d *Daos) FilterUsageType(ctx *models.Context, usageTypefilter *models.UsageTypeFilter, pagination *models.Pagination) ([]models.RefUsageType, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if usageTypefilter != nil {

		if len(usageTypefilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": usageTypefilter.Status}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONUSAGETYPE).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("usageType query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONUSAGETYPE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var usageTypes []models.RefUsageType
	if err = cursor.All(context.TODO(), &usageTypes); err != nil {
		return nil, err
	}
	return usageTypes, nil
}

//EnableUsageType :""
func (d *Daos) EnableUsageType(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.USAGETYPESTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONUSAGETYPE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableUsageType :""
func (d *Daos) DisableUsageType(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.USAGETYPESTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONUSAGETYPE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteUsageType :""
func (d *Daos) DeleteUsageType(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.USAGETYPESTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONUSAGETYPE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
