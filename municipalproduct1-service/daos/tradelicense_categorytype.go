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

// SaveTradeLicenseCategoryType : ""
func (d *Daos) SaveTradeLicenseCategoryType(ctx *models.Context, categoryType *models.TradeLicenseCategoryType) error {
	d.Shared.BsonToJSONPrint(categoryType)
	_, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSECATEGORYTYPE).InsertOne(ctx.CTX, categoryType)
	return err
}

// GetSingleTradeLicenseCategoryType : ""
func (d *Daos) GetSingleTradeLicenseCategoryType(ctx *models.Context, UniqueID string) (*models.RefTradeLicenseCategoryType, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	// LookUp
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONTRADELICENSEBUSINESSTYPE, "tlbtId", "uniqueId", "ref.tradeLicenseBusinessType", "ref.tradeLicenseBusinessType")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSECATEGORYTYPE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var towers []models.RefTradeLicenseCategoryType
	var tower *models.RefTradeLicenseCategoryType
	if err = cursor.All(ctx.CTX, &towers); err != nil {
		return nil, err
	}
	if len(towers) > 0 {
		tower = &towers[0]
	}
	return tower, nil
}

// UpdateTradeLicenseCategoryType : ""
func (d *Daos) UpdateTradeLicenseCategoryType(ctx *models.Context, categoryType *models.TradeLicenseCategoryType) error {
	selector := bson.M{"uniqueId": categoryType.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	data := bson.M{"$set": categoryType}
	_, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSECATEGORYTYPE).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// EnableTradeLicenseCategoryType : ""
func (d *Daos) EnableTradeLicenseCategoryType(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.TRADELICENSECATEGORYTYPESTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSECATEGORYTYPE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DisableTradeLicenseCategoryType : ""
func (d *Daos) DisableTradeLicenseCategoryType(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.TRADELICENSECATEGORYTYPESTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSECATEGORYTYPE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DeleteTradeLicenseCategoryType : ""
func (d *Daos) DeleteTradeLicenseCategoryType(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.TRADELICENSECATEGORYTYPESTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSECATEGORYTYPE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterTradeLicenseCategoryType : ""
func (d *Daos) FilterTradeLicenseCategoryType(ctx *models.Context, filter *models.TradeLicenseCategoryTypeFilter, pagination *models.Pagination) ([]models.RefTradeLicenseCategoryType, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
		if len(filter.TLBTID) > 0 {
			query = append(query, bson.M{"tlbtId": bson.M{"$in": filter.TLBTID}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSECATEGORYTYPE).CountDocuments(ctx.CTX, func() bson.M {
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

	// LookUp
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONTRADELICENSEBUSINESSTYPE, "tlbtId", "uniqueId", "ref.tradeLicenseBusinessType", "ref.tradeLicenseBusinessType")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSECATEGORYTYPE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var categoryType []models.RefTradeLicenseCategoryType
	if err = cursor.All(context.TODO(), &categoryType); err != nil {
		return nil, err
	}
	return categoryType, nil
}
