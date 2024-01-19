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

//SavePropertyOtherTax :""
func (d *Daos) SavePropertyOtherTax(ctx *models.Context, propertyOtherTax *models.PropertyOtherTax) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYOTHERTAX).InsertOne(ctx.CTX, propertyOtherTax)
	return err
}

//GetSinglePropertyOtherTax : ""
func (d *Daos) GetSinglePropertyOtherTax(ctx *models.Context, UniqueID string) (*models.RefPropertyOtherTax, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYOTHERTAX).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var propertyOtherTaxs []models.RefPropertyOtherTax
	var propertyOtherTax *models.RefPropertyOtherTax
	if err = cursor.All(ctx.CTX, &propertyOtherTaxs); err != nil {
		return nil, err
	}
	if len(propertyOtherTaxs) > 0 {
		propertyOtherTax = &propertyOtherTaxs[0]
	}
	return propertyOtherTax, nil
}

//UpdatePropertyOtherTax : ""
func (d *Daos) UpdatePropertyOtherTax(ctx *models.Context, propertyOtherTax *models.PropertyOtherTax) error {
	selector := bson.M{"uniqueId": propertyOtherTax.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": propertyOtherTax, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYOTHERTAX).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterPropertyOtherTax : ""
func (d *Daos) FilterPropertyOtherTax(ctx *models.Context, propertyOtherTaxfilter *models.PropertyOtherTaxFilter, pagination *models.Pagination) ([]models.RefPropertyOtherTax, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if propertyOtherTaxfilter != nil {

		if len(propertyOtherTaxfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": propertyOtherTaxfilter.Status}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYOTHERTAX).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("propertyOtherTax query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYOTHERTAX).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var propertyOtherTaxs []models.RefPropertyOtherTax
	if err = cursor.All(context.TODO(), &propertyOtherTaxs); err != nil {
		return nil, err
	}
	return propertyOtherTaxs, nil
}

//EnablePropertyOtherTax :""
func (d *Daos) EnablePropertyOtherTax(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PROPERTYOTHERSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYOTHERTAX).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisablePropertyOtherTax :""
func (d *Daos) DisablePropertyOtherTax(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PROPERTYOTHERSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYOTHERTAX).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeletePropertyOtherTax :""
func (d *Daos) DeletePropertyOtherTax(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PROPERTYOTHERSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYOTHERTAX).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
