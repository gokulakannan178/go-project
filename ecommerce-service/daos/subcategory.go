package daos

import (
	"context"
	"ecommerce-service/constants"
	"ecommerce-service/models"
	"errors"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// SaveSubCategory : ""
func (d *Daos) SaveSubCategory(ctx *models.Context, subCategory *models.SubCategory) error {
	d.Shared.BsonToJSONPrint(subCategory)
	_, err := ctx.DB.Collection(constants.COLLECTIONSUBCATEGORY).InsertOne(ctx.CTX, subCategory)
	return err
}

// GetSingleSubCategory : ""
func (d *Daos) GetSingleSubCategory(ctx *models.Context, UniqueID string) (*models.RefSubCategory, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	//Lookups
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONCATEGORY, "categoryId", "uniqueId", "ref.category", "ref.category")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSUBCATEGORY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var subCategorys []models.RefSubCategory
	var subCategory *models.RefSubCategory
	if err = cursor.All(ctx.CTX, &subCategorys); err != nil {
		return nil, err
	}
	if len(subCategorys) > 0 {
		subCategory = &subCategorys[0]
	}
	return subCategory, nil
}

// UpdateSubCategory : ""
func (d *Daos) UpdateSubCategory(ctx *models.Context, subCategory *models.SubCategory) error {
	selector := bson.M{"uniqueId": subCategory.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	data := bson.M{"$set": subCategory}
	_, err := ctx.DB.Collection(constants.COLLECTIONSUBCATEGORY).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// EnableSubCategory : ""
func (d *Daos) EnableSubCategory(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.SUBCATEGORYSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONSUBCATEGORY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DisableSubCategory : ""
func (d *Daos) DisableSubCategory(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.SUBCATEGORYSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONSUBCATEGORY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DeleteSubCategory : ""
func (d *Daos) DeleteSubCategory(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.SUBCATEGORYSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONSUBCATEGORY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterSubCategory : ""
func (d *Daos) FilterSubCategory(ctx *models.Context, filter *models.SubCategoryFilter, pagination *models.Pagination) ([]models.RefSubCategory, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
		if len(filter.UniqueID) > 0 {
			query = append(query, bson.M{"uniqueId": bson.M{"$in": filter.UniqueID}})
		}
		if len(filter.CategoryID) > 0 {
			query = append(query, bson.M{"categoryId": bson.M{"$in": filter.CategoryID}})
		}

		//Regex Using searchBox Struct
		if filter.SearchText.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: filter.SearchText.Name, Options: "xi"}})
		}
		if filter.SearchText.UniqueID != "" {
			query = append(query, bson.M{"uniqueId": primitive.Regex{Pattern: filter.SearchText.UniqueID, Options: "xi"}})
		}
		if filter.DateRange != nil {
			//var sd,ed time.Time
			if filter.DateRange.From != nil {
				sd := time.Date(filter.DateRange.From.Year(), filter.DateRange.From.Month(), filter.DateRange.From.Day(), 0, 0, 0, 0, filter.DateRange.From.Location())
				ed := time.Date(filter.DateRange.From.Year(), filter.DateRange.From.Month(), filter.DateRange.From.Day(), 23, 59, 59, 0, filter.DateRange.From.Location())
				if filter.DateRange.To != nil {
					ed = time.Date(filter.DateRange.To.Year(), filter.DateRange.To.Month(), filter.DateRange.To.Day(), 23, 59, 59, 0, filter.DateRange.To.Location())
				}
				query = append(query, bson.M{"created.on": bson.M{"$gte": sd, "$lte": ed}})

			}
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONSUBCATEGORY).CountDocuments(ctx.CTX, func() bson.M {
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
	//Lookups
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONCATEGORY, "categoryId", "uniqueId", "ref.category", "ref.category")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSUBCATEGORY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var subCategory []models.RefSubCategory
	if err = cursor.All(context.TODO(), &subCategory); err != nil {
		return nil, err
	}
	return subCategory, nil
}
