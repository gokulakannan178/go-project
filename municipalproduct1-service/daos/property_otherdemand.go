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
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// SavePropertyOtherDemand : ""
func (d *Daos) SavePropertyOtherDemand(ctx *models.Context, propertyotherdemand *models.PropertyOtherDemand) error {
	d.Shared.BsonToJSONPrint(propertyotherdemand)
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYOTHERDEMAND).InsertOne(ctx.CTX, propertyotherdemand)
	return err
}

// GetSinglePropertyOtherDemand : ""
func (d *Daos) GetSinglePropertyOtherDemand(ctx *models.Context, UniqueID string) (*models.RefPropertyOtherDemand, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	//lookup
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPROPERTYOTHERDEMANDCATEGORY, "categoryId", "uniqueId", "ref.category", "ref.category")...)
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPROPERTYOTHERDEMANDSUBCATEGORY, "subCategoryId", "uniqueId", "ref.subCategory", "ref.subCategory")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYOTHERDEMAND).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var propertyotherdemands []models.RefPropertyOtherDemand
	var propertyotherdemand *models.RefPropertyOtherDemand
	if err = cursor.All(ctx.CTX, &propertyotherdemands); err != nil {
		return nil, err
	}
	if len(propertyotherdemands) > 0 {
		propertyotherdemand = &propertyotherdemands[0]
	}
	return propertyotherdemand, nil
}

// UpdatePropertyOtherDemand : ""
func (d *Daos) UpdatePropertyOtherDemand(ctx *models.Context, propertyotherdemand *models.PropertyOtherDemand) error {
	selector := bson.M{"uniqueId": propertyotherdemand.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM

	data := bson.M{"$set": propertyotherdemand}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYOTHERDEMAND).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// EnablePropertyOtherDemand : ""
func (d *Daos) EnablePropertyOtherDemand(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PROPERTYOTHERDEMANDSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYOTHERDEMAND).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DisablePropertyOtherDemand : ""
func (d *Daos) DisablePropertyOtherDemand(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PROPERTYOTHERDEMANDSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYOTHERDEMAND).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DeletePropertyOtherDemand : ""
func (d *Daos) DeletePropertyOtherDemand(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PROPERTYOTHERDEMANDSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYOTHERDEMAND).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// UpdatePropertyOtherDemandStatus : ""
func (d *Daos) UpdatePropertyOtherDemandStatus(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"paymentStatus": constants.PROPERTYOTHERDEMANDPAYMENTSTATUSPAID}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYOTHERDEMAND).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterPropertyOtherDemand : ""
func (d *Daos) FilterPropertyOtherDemand(ctx *models.Context, filter *models.PropertyOtherDemandFilter, pagination *models.Pagination) ([]models.RefPropertyOtherDemand, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
		if len(filter.UniqueIDs) > 0 {
			query = append(query, bson.M{"uniqueId": bson.M{"$in": filter.UniqueIDs}})
		}
		if len(filter.FyID) > 0 {
			query = append(query, bson.M{"fyId": bson.M{"$in": filter.FyID}})
		}
		if len(filter.PropertyID) > 0 {
			query = append(query, bson.M{"propertyId": bson.M{"$in": filter.PropertyID}})
		}
		if len(filter.PaymentStatus) > 0 {
			query = append(query, bson.M{"paymentStatus": bson.M{"$in": filter.PaymentStatus}})
		}
		//regex

		if filter.Regex.UniqueID != "" {
			query = append(query, bson.M{"uniqueId": primitive.Regex{Pattern: filter.Regex.UniqueID, Options: "xi"}})
		}

	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	if filter.SortBy != "" {
		mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{filter.SortBy: filter.SortOrder}})
	} else {
		mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"_id": -1}})
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYOTHERDEMAND).CountDocuments(ctx.CTX, func() bson.M {
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
	// //Lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONFINANCIALYEAR, "fyId", "uniqueId", "ref.fy", "ref.fy")...)
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPROPERTYOTHERDEMANDSUBCATEGORY, "subCategoryId", "uniqueId", "ref.subCategory", "ref.subCategory")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYOTHERDEMAND).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var propertyotherdemand []models.RefPropertyOtherDemand
	if err = cursor.All(context.TODO(), &propertyotherdemand); err != nil {
		return nil, err
	}
	return propertyotherdemand, nil
}
