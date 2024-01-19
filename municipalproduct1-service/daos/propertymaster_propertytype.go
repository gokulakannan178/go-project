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

//SavePropertyType :""
func (d *Daos) SavePropertyType(ctx *models.Context, propertyType *models.PropertyType) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYTYPE).InsertOne(ctx.CTX, propertyType)
	return err
}

//GetSinglePropertyType : ""
func (d *Daos) GetSinglePropertyType(ctx *models.Context, UniqueID string) (*models.RefPropertyType, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYTYPE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var propertyTypes []models.RefPropertyType
	var propertyType *models.RefPropertyType
	if err = cursor.All(ctx.CTX, &propertyTypes); err != nil {
		return nil, err
	}
	if len(propertyTypes) > 0 {
		propertyType = &propertyTypes[0]
	}
	return propertyType, nil
}

//UpdatePropertyType : ""
func (d *Daos) UpdatePropertyType(ctx *models.Context, propertyType *models.PropertyType) error {
	selector := bson.M{"uniqueId": propertyType.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": propertyType, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYTYPE).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterPropertyType : ""
func (d *Daos) FilterPropertyType(ctx *models.Context, propertyTypefilter *models.PropertyTypeFilter, pagination *models.Pagination) ([]models.RefPropertyType, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if propertyTypefilter != nil {

		if len(propertyTypefilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": propertyTypefilter.Status}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if propertyTypefilter != nil {
		if propertyTypefilter.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{propertyTypefilter.SortBy: propertyTypefilter.SortOrder}})

		}
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYTYPE).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("propertyType query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYTYPE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var propertyTypes []models.RefPropertyType
	if err = cursor.All(context.TODO(), &propertyTypes); err != nil {
		return nil, err
	}
	return propertyTypes, nil
}

//EnablePropertyType :""
func (d *Daos) EnablePropertyType(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PROPERTYTYPESTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYTYPE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisablePropertyType :""
func (d *Daos) DisablePropertyType(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PROPERTYTYPESTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYTYPE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeletePropertyType :""
func (d *Daos) DeletePropertyType(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PROPERTYTYPESTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYTYPE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
