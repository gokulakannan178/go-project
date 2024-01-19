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

// SaveCategory : ""
func (d *Daos) SaveCategory(ctx *models.Context, category *models.Category) error {
	d.Shared.BsonToJSONPrint(category)
	_, err := ctx.DB.Collection(constants.COLLECTIONCATEGORY).InsertOne(ctx.CTX, category)
	return err
}

// GetSingleCategory : ""
func (d *Daos) GetSingleCategory(ctx *models.Context, UniqueID string) (*models.RefCategory, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONCATEGORY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var categorys []models.RefCategory
	var category *models.RefCategory
	if err = cursor.All(ctx.CTX, &categorys); err != nil {
		return nil, err
	}
	if len(categorys) > 0 {
		category = &categorys[0]
	}
	return category, nil
}

// UpdateCategory : ""
func (d *Daos) UpdateCategory(ctx *models.Context, category *models.Category) error {
	selector := bson.M{"uniqueId": category.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	data := bson.M{"$set": category}
	_, err := ctx.DB.Collection(constants.COLLECTIONCATEGORY).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// EnableCategory : ""
func (d *Daos) EnableCategory(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.CATEGORYSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONCATEGORY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DisableCategory : ""
func (d *Daos) DisableCategory(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.CATEGORYSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONCATEGORY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DeleteCategory : ""
func (d *Daos) DeleteCategory(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.CATEGORYSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONCATEGORY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterCategory : ""
func (d *Daos) FilterCategory(ctx *models.Context, filter *models.CategoryFilter, pagination *models.Pagination) ([]models.RefCategory, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
		if len(filter.UniqueID) > 0 {
			query = append(query, bson.M{"uniqueId": bson.M{"$in": filter.UniqueID}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONCATEGORY).CountDocuments(ctx.CTX, func() bson.M {
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
	cursor, err := ctx.DB.Collection(constants.COLLECTIONCATEGORY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var category []models.RefCategory
	if err = cursor.All(context.TODO(), &category); err != nil {
		return nil, err
	}
	return category, nil
}
