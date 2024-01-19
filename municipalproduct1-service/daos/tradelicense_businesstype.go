package daos

import (
	"context"
	"errors"
	"fmt"
	"log"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"

	"go.mongodb.org/mongo-driver/bson"
)

// SaveTradeLicenseBusinessType : ""
func (d *Daos) SaveTradeLicenseBusinessType(ctx *models.Context, businessType *models.TradeLicenseBusinessType) error {
	d.Shared.BsonToJSONPrint(businessType)
	_, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSEBUSINESSTYPE).InsertOne(ctx.CTX, businessType)
	return err
}

// GetSingleTradeLicenseBusinessType : ""
func (d *Daos) GetSingleTradeLicenseBusinessType(ctx *models.Context, UniqueID string) (*models.RefTradeLicenseBusinessType, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSEBUSINESSTYPE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var towers []models.RefTradeLicenseBusinessType
	var tower *models.RefTradeLicenseBusinessType
	if err = cursor.All(ctx.CTX, &towers); err != nil {
		return nil, err
	}
	if len(towers) > 0 {
		tower = &towers[0]
	}
	return tower, nil
}

// UpdateTradeLicenseBusinessType : ""
func (d *Daos) UpdateTradeLicenseBusinessType(ctx *models.Context, business *models.TradeLicenseBusinessType) error {
	selector := bson.M{"uniqueId": business.UniqueID}
	// t := time.Now()
	// update := models.Updated{}
	// update.On = &t
	// update.By = constants.SYSTEM

	data := bson.M{"$set": business}
	_, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSEBUSINESSTYPE).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// EnableTradeLicenseBusinessType : ""
func (d *Daos) EnableTradeLicenseBusinessType(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.TRADELICENSEBUSINESSTYPESTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSEBUSINESSTYPE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DisableTradeLicenseBusinessType : ""
func (d *Daos) DisableTradeLicenseBusinessType(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.TRADELICENSEBUSINESSTYPESTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSEBUSINESSTYPE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DeleteTradeLicenseBusinessType : ""
func (d *Daos) DeleteTradeLicenseBusinessType(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.TRADELICENSEBUSINESSTYPESTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSEBUSINESSTYPE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterTradeLicenseBusinessType : ""
func (d *Daos) FilterTradeLicenseBusinessType(ctx *models.Context, filter *models.TradeLicenseBusinessTypeFilter, pagination *models.Pagination) ([]models.RefTradeLicenseBusinessType, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSEBUSINESSTYPE).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSEBUSINESSTYPE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var businessType []models.RefTradeLicenseBusinessType
	if err = cursor.All(context.TODO(), &businessType); err != nil {
		return nil, err
	}
	return businessType, nil
}
