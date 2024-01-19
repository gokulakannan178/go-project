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

// SaveWaterTaxArv : ""
func (d *Daos) SaveWaterTaxArv(ctx *models.Context, watertaxarv *models.WaterTaxArv) error {
	d.Shared.BsonToJSONPrint(watertaxarv)
	_, err := ctx.DB.Collection(constants.COLLECTIONWATERTAXARV).InsertOne(ctx.CTX, watertaxarv)
	return err
}

// GetSingleWaterTaxArv : ""
func (d *Daos) GetSingleWaterTaxArv(ctx *models.Context, UniqueID string) (*models.RefWaterTaxArv, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	//lookup
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONWATERTAXARVCATEGORY, "categoryId", "uniqueId", "ref.category", "ref.category")...)
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONWATERTAXARVSUBCATEGORY, "subCategoryId", "uniqueId", "ref.subCategory", "ref.subCategory")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONWATERTAXARV).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var watertaxarvs []models.RefWaterTaxArv
	var watertaxarv *models.RefWaterTaxArv
	if err = cursor.All(ctx.CTX, &watertaxarvs); err != nil {
		return nil, err
	}
	if len(watertaxarvs) > 0 {
		watertaxarv = &watertaxarvs[0]
	}
	return watertaxarv, nil
}

// UpdateWaterTaxArv : ""
func (d *Daos) UpdateWaterTaxArv(ctx *models.Context, watertaxarv *models.WaterTaxArv) error {
	selector := bson.M{"uniqueId": watertaxarv.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM

	data := bson.M{"$set": watertaxarv}
	_, err := ctx.DB.Collection(constants.COLLECTIONWATERTAXARV).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// EnableWaterTaxArv : ""
func (d *Daos) EnableWaterTaxArv(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.WATERTAXARVSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONWATERTAXARV).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DisableWaterTaxArv : ""
func (d *Daos) DisableWaterTaxArv(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.WATERTAXARVSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONWATERTAXARV).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DeleteWaterTaxArv : ""
func (d *Daos) DeleteWaterTaxArv(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.WATERTAXARVSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONWATERTAXARV).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterWaterTaxArv : ""
func (d *Daos) FilterWaterTaxArv(ctx *models.Context, filter *models.WaterTaxArvFilter, pagination *models.Pagination) ([]models.RefWaterTaxArv, error) {
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONWATERTAXARV).CountDocuments(ctx.CTX, func() bson.M {
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
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONWATERTAXARVCATEGORY, "categoryId", "uniqueId", "ref.category", "ref.category")...)
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONWATERTAXARVSUBCATEGORY, "subCategoryId", "uniqueId", "ref.subCategory", "ref.subCategory")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONWATERTAXARV).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var watertaxarv []models.RefWaterTaxArv
	if err = cursor.All(context.TODO(), &watertaxarv); err != nil {
		return nil, err
	}
	return watertaxarv, nil
}
