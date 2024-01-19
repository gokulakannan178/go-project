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

//SaveRebate :""
func (d *Daos) SaveRebate(ctx *models.Context, rebate *models.Rebate) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONREBATE).InsertOne(ctx.CTX, rebate)
	return err
}

//GetSingleRebate : ""
func (d *Daos) GetSingleRebate(ctx *models.Context, UniqueID string) (*models.RefRebate, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONREBATE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var rebates []models.RefRebate
	var rebate *models.RefRebate
	if err = cursor.All(ctx.CTX, &rebates); err != nil {
		return nil, err
	}
	if len(rebates) > 0 {
		rebate = &rebates[0]
	}
	return rebate, nil
}

//UpdateRebate : ""
func (d *Daos) UpdateRebate(ctx *models.Context, rebate *models.Rebate) error {
	selector := bson.M{"uniqueId": rebate.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": rebate, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONREBATE).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterRebate : ""
func (d *Daos) FilterRebate(ctx *models.Context, rebatefilter *models.RebateFilter, pagination *models.Pagination) ([]models.RefRebate, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if rebatefilter != nil {

		if len(rebatefilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": rebatefilter.Status}})
		}
		if len(rebatefilter.UniqueID) > 0 {
			query = append(query, bson.M{"uniqueId": bson.M{"$in": rebatefilter.UniqueID}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONREBATE).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("rebate query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONREBATE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var rebates []models.RefRebate
	if err = cursor.All(context.TODO(), &rebates); err != nil {
		return nil, err
	}
	return rebates, nil
}

//EnableRebate :""
func (d *Daos) EnableRebate(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.REBATESTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONREBATE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableRebate :""
func (d *Daos) DisableRebate(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.REBATESTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONREBATE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteRebate :""
func (d *Daos) DeleteRebate(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.REBATESTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONREBATE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
