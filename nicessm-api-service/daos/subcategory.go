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

//SaveSubCategory :""
func (d *Daos) SaveSubCategory(ctx *models.Context, subcategory *models.SubCategory) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONSUBCATEGORY).InsertOne(ctx.CTX, subcategory)
	if err != nil {
		return err
	}
	subcategory.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//GetSingleSubCategory : ""
func (d *Daos) GetSingleSubCategory(ctx *models.Context, UniqueID string) (*models.RefSubCategory, error) {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSUBCATEGORY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var subcategorys []models.RefSubCategory
	var subcategory *models.RefSubCategory
	if err = cursor.All(ctx.CTX, &subcategorys); err != nil {
		return nil, err
	}
	if len(subcategorys) > 0 {
		subcategory = &subcategorys[0]
	}
	return subcategory, nil
}

//UpdateSubCategory : ""
func (d *Daos) UpdateSubCategory(ctx *models.Context, subcategory *models.SubCategory) error {

	selector := bson.M{"_id": subcategory.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": subcategory}
	_, err := ctx.DB.Collection(constants.COLLECTIONSUBCATEGORY).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterSubCategory : ""
func (d *Daos) FilterSubCategory(ctx *models.Context, filter *models.SubCategoryFilter, pagination *models.Pagination) ([]models.RefSubCategory, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {

		if len(filter.ActiveStatus) > 0 {
			query = append(query, bson.M{"activeStatus": bson.M{"$in": filter.ActiveStatus}})
		}
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}

		//Regex
		if filter.SearchBox.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: filter.SearchBox.Name, Options: "xi"}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if filter != nil {
		if filter.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{filter.SortBy: filter.SortOrder}})

		}

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

	//Aggregation
	d.Shared.BsonToJSONPrintTag("SubCategory query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSUBCATEGORY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var subcategorys []models.RefSubCategory
	if err = cursor.All(context.TODO(), &subcategorys); err != nil {
		return nil, err
	}
	return subcategorys, nil
}

//EnableSubCategory :""
func (d *Daos) EnableSubCategory(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.SUBCATEGORYSTATUSACTIVE, "activeStatus": true}}
	_, err = ctx.DB.Collection(constants.COLLECTIONSUBCATEGORY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableSubCategory :""
func (d *Daos) DisableSubCategory(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.SUBCATEGORYSTATUSDISABLED, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONSUBCATEGORY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteSubCategory :""
func (d *Daos) DeleteSubCategory(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.SUBCATEGORYSTATUSDELETED, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONSUBCATEGORY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
