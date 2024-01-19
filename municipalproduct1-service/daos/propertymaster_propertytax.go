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

//SavePropertyTax :""
func (d *Daos) SavePropertyTax(ctx *models.Context, propertyTax *models.PropertyTax) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYTAX).InsertOne(ctx.CTX, propertyTax)
	return err
}

//GetSinglePropertyTax : ""
func (d *Daos) GetSinglePropertyTax(ctx *models.Context, UniqueID string) (*models.RefPropertyTax, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYTAX).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var propertyTaxs []models.RefPropertyTax
	var propertyTax *models.RefPropertyTax
	if err = cursor.All(ctx.CTX, &propertyTaxs); err != nil {
		return nil, err
	}
	if len(propertyTaxs) > 0 {
		propertyTax = &propertyTaxs[0]
	}
	return propertyTax, nil
}

//UpdatePropertyTax : ""
func (d *Daos) UpdatePropertyTax(ctx *models.Context, propertyTax *models.PropertyTax) error {
	selector := bson.M{"uniqueId": propertyTax.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": propertyTax, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYTAX).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterPropertyTax : ""
func (d *Daos) FilterPropertyTax(ctx *models.Context, propertyTaxfilter *models.PropertyTaxFilter, pagination *models.Pagination) ([]models.RefPropertyTax, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if propertyTaxfilter != nil {

		if len(propertyTaxfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": propertyTaxfilter.Status}})
		}
		if len(propertyTaxfilter.UniqueID) > 0 {
			query = append(query, bson.M{"uniqueId": bson.M{"$in": propertyTaxfilter.UniqueID}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if propertyTaxfilter != nil {
		if propertyTaxfilter.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{propertyTaxfilter.SortBy: propertyTaxfilter.SortOrder}})

		}
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYTAX).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("propertyTax query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYTAX).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var propertyTaxs []models.RefPropertyTax
	if err = cursor.All(context.TODO(), &propertyTaxs); err != nil {
		return nil, err
	}
	return propertyTaxs, nil
}

//EnablePropertyTax :""
func (d *Daos) EnablePropertyTax(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PROPERTYTAXSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYTAX).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisablePropertyTax :""
func (d *Daos) DisablePropertyTax(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PROPERTYTAXSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYTAX).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeletePropertyTax :""
func (d *Daos) DeletePropertyTax(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PROPERTYTAXSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYTAX).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
