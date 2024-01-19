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

//SaveOtherCharges :""
func (d *Daos) SaveOtherCharges(ctx *models.Context, otherCharges *models.OtherCharges) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONOTHERCHARGES).InsertOne(ctx.CTX, otherCharges)
	return err
}

//GetSingleOtherCharges : ""
func (d *Daos) GetSingleOtherCharges(ctx *models.Context, UniqueID string) (*models.RefOtherCharges, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONOTHERCHARGES).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var otherChargess []models.RefOtherCharges
	var otherCharges *models.RefOtherCharges
	if err = cursor.All(ctx.CTX, &otherChargess); err != nil {
		return nil, err
	}
	if len(otherChargess) > 0 {
		otherCharges = &otherChargess[0]
	}
	return otherCharges, nil
}

//UpdateOtherCharges : ""
func (d *Daos) UpdateOtherCharges(ctx *models.Context, otherCharges *models.OtherCharges) error {
	selector := bson.M{"uniqueId": otherCharges.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": otherCharges, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONOTHERCHARGES).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterOtherCharges : ""
func (d *Daos) FilterOtherCharges(ctx *models.Context, otherChargesfilter *models.OtherChargesFilter, pagination *models.Pagination) ([]models.RefOtherCharges, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if otherChargesfilter != nil {

		if len(otherChargesfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": otherChargesfilter.Status}})
		}
		if len(otherChargesfilter.UniqueID) > 0 {
			query = append(query, bson.M{"uniqueId": bson.M{"$in": otherChargesfilter.UniqueID}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONOTHERCHARGES).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("otherCharges query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONOTHERCHARGES).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var otherChargess []models.RefOtherCharges
	if err = cursor.All(context.TODO(), &otherChargess); err != nil {
		return nil, err
	}
	return otherChargess, nil
}

//EnableOtherCharges :""
func (d *Daos) EnableOtherCharges(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.OTHERCHARGESSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONOTHERCHARGES).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableOtherCharges :""
func (d *Daos) DisableOtherCharges(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.OTHERCHARGESSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONOTHERCHARGES).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteOtherCharges :""
func (d *Daos) DeleteOtherCharges(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.OTHERCHARGESSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONOTHERCHARGES).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
