package daos

import (
	"context"
	"errors"
	"fmt"
	"log"
	"nicessm-api-service/constants"
	"nicessm-api-service/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//SaveCategory :""
func (d *Daos) SaveCategory(ctx *models.Context, category *models.Category) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONCATEGORY).InsertOne(ctx.CTX, category)
	if err != nil {
		return err
	}
	category.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//GetSingleCategory : ""
func (d *Daos) GetSingleCategory(ctx *models.Context, UniqueID string) (*models.RefCategory, error) {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})
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

//UpdateCategory : ""
func (d *Daos) UpdateCategory(ctx *models.Context, category *models.Category) error {

	selector := bson.M{"_id": category.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": category}
	_, err := ctx.DB.Collection(constants.COLLECTIONCATEGORY).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterCategory : ""
func (d *Daos) FilterCategory(ctx *models.Context, Filter *models.CategoryFilter, pagination *models.Pagination) ([]models.RefCategory, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if Filter != nil {

		if len(Filter.ActiveStatus) > 0 {
			query = append(query, bson.M{"activeStatus": bson.M{"$in": Filter.ActiveStatus}})
		}
		if len(Filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": Filter.Status}})
		}

		//Regex
		if Filter.SearchBox.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: Filter.SearchBox.Name, Options: "xi"}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if Filter != nil {
		if Filter.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{Filter.SortBy: Filter.SortOrder}})

		}

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
	d.Shared.BsonToJSONPrintTag("Category query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONCATEGORY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var categorys []models.RefCategory
	if err = cursor.All(context.TODO(), &categorys); err != nil {
		return nil, err
	}
	return categorys, nil
}

//EnableCategory :""
func (d *Daos) EnableCategory(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.CATEGORYSTATUSACTIVE, "activeStatus": true}}
	_, err = ctx.DB.Collection(constants.COLLECTIONCATEGORY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableCategory :""
func (d *Daos) DisableCategory(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.CATEGORYSTATUSDISABLED, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONCATEGORY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteCategory :""
func (d *Daos) DeleteCategory(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.CATEGORYSTATUSDELETED, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONCATEGORY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
