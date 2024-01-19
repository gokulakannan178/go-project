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

//SaveTradeLicenseRebate :""
func (d *Daos) SaveTradeLicenseRebate(ctx *models.Context, tlRebate *models.TradeLicenseRebate) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSEREBATE).InsertOne(ctx.CTX, tlRebate)
	return err
}

//GetSingleTradeLicenseRebate : ""
func (d *Daos) GetSingleTradeLicenseRebate(ctx *models.Context, UniqueID string) (*models.RefTradeLicenseRebate, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSEREBATE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var tlRebates []models.RefTradeLicenseRebate
	var tlRebate *models.RefTradeLicenseRebate
	if err = cursor.All(ctx.CTX, &tlRebates); err != nil {
		return nil, err
	}
	if len(tlRebates) > 0 {
		tlRebate = &tlRebates[0]
	}
	return tlRebate, nil
}

//UpdateTradeLicenseRebate : ""
func (d *Daos) UpdateTradeLicenseRebate(ctx *models.Context, tlRebate *models.TradeLicenseRebate) error {
	selector := bson.M{"uniqueId": tlRebate.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": tlRebate, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSEREBATE).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterTradeLicenseRebate : ""
func (d *Daos) FilterTradeLicenseRebate(ctx *models.Context, tlRebatefilter *models.TradeLicenseRebateFilter, pagination *models.Pagination) ([]models.RefTradeLicenseRebate, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if tlRebatefilter != nil {

		if len(tlRebatefilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": tlRebatefilter.Status}})
		}
		if len(tlRebatefilter.UniqueID) > 0 {
			query = append(query, bson.M{"uniqueId": bson.M{"$in": tlRebatefilter.UniqueID}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSEREBATE).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("tlRebate query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSEREBATE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var tlRebates []models.RefTradeLicenseRebate
	if err = cursor.All(context.TODO(), &tlRebates); err != nil {
		return nil, err
	}
	return tlRebates, nil
}

//EnableTradeLicenseRebate :""
func (d *Daos) EnableTradeLicenseRebate(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.TRADELICENSEREBATESTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSEREBATE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableTradeLicenseRebate :""
func (d *Daos) DisableTradeLicenseRebate(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.TRADELICENSEREBATESTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSEREBATE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteTradeLicenseRebate :""
func (d *Daos) DeleteTradeLicenseRebate(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.TRADELICENSEREBATESTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSEREBATE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
