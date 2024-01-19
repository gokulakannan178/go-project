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

// SaveWaterTaxConnectionType : ""
func (d *Daos) SaveWaterTaxConnectionType(ctx *models.Context, watertaxconnectiontype *models.WaterTaxConnectionType) error {
	d.Shared.BsonToJSONPrint(watertaxconnectiontype)
	_, err := ctx.DB.Collection(constants.COLLECTIONWATERTAXCONNECTIONTYPE).InsertOne(ctx.CTX, watertaxconnectiontype)
	return err
}

// GetSingleWaterTaxConnectionType : ""
func (d *Daos) GetSingleWaterTaxConnectionType(ctx *models.Context, UniqueID string) (*models.RefWaterTaxConnectionType, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	//lookup
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONWATERTAXCONNECTIONTYPECATEGORY, "categoryId", "uniqueId", "ref.category", "ref.category")...)
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONWATERTAXCONNECTIONTYPESUBCATEGORY, "subCategoryId", "uniqueId", "ref.subCategory", "ref.subCategory")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONWATERTAXCONNECTIONTYPE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var watertaxconnectiontypes []models.RefWaterTaxConnectionType
	var watertaxconnectiontype *models.RefWaterTaxConnectionType
	if err = cursor.All(ctx.CTX, &watertaxconnectiontypes); err != nil {
		return nil, err
	}
	if len(watertaxconnectiontypes) > 0 {
		watertaxconnectiontype = &watertaxconnectiontypes[0]
	}
	return watertaxconnectiontype, nil
}

// UpdateWaterTaxConnectionType : ""
func (d *Daos) UpdateWaterTaxConnectionType(ctx *models.Context, watertaxconnectiontype *models.WaterTaxConnectionType) error {
	selector := bson.M{"uniqueId": watertaxconnectiontype.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM

	data := bson.M{"$set": watertaxconnectiontype}
	_, err := ctx.DB.Collection(constants.COLLECTIONWATERTAXCONNECTIONTYPE).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// EnableWaterTaxConnectionType : ""
func (d *Daos) EnableWaterTaxConnectionType(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.WATERTAXCONNECTIONTYPESTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONWATERTAXCONNECTIONTYPE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DisableWaterTaxConnectionType : ""
func (d *Daos) DisableWaterTaxConnectionType(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.WATERTAXCONNECTIONTYPESTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONWATERTAXCONNECTIONTYPE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DeleteWaterTaxConnectionType : ""
func (d *Daos) DeleteWaterTaxConnectionType(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.WATERTAXCONNECTIONTYPESTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONWATERTAXCONNECTIONTYPE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterWaterTaxConnectionType : ""
func (d *Daos) FilterWaterTaxConnectionType(ctx *models.Context, filter *models.WaterTaxConnectionTypeFilter, pagination *models.Pagination) ([]models.RefWaterTaxConnectionType, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
		if len(filter.UniqueIDs) > 0 {
			query = append(query, bson.M{"uniqueId": bson.M{"$in": filter.UniqueIDs}})
		}
		//regex

		if filter.Regex.UniqueID != "" {
			query = append(query, bson.M{"uniqueId": primitive.Regex{Pattern: filter.Regex.UniqueID, Options: "xi"}})
		}
		if filter.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: filter.Regex.Name, Options: "xi"}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONWATERTAXCONNECTIONTYPE).CountDocuments(ctx.CTX, func() bson.M {
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
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONWATERTAXCONNECTIONTYPECATEGORY, "categoryId", "uniqueId", "ref.category", "ref.category")...)
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONWATERTAXCONNECTIONTYPESUBCATEGORY, "subCategoryId", "uniqueId", "ref.subCategory", "ref.subCategory")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONWATERTAXCONNECTIONTYPE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var watertaxconnectiontype []models.RefWaterTaxConnectionType
	if err = cursor.All(context.TODO(), &watertaxconnectiontype); err != nil {
		return nil, err
	}
	return watertaxconnectiontype, nil
}
