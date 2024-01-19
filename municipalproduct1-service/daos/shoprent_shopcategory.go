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

// SaveShopRentShopCategory : ""
func (d *Daos) SaveShopRentShopCategory(ctx *models.Context, shopcategory *models.ShopRentShopCategory) error {
	d.Shared.BsonToJSONPrint(shopcategory)
	_, err := ctx.DB.Collection(constants.COLLECTIONSHOPRENTSHOPCATEGORY).InsertOne(ctx.CTX, shopcategory)
	return err
}

// GetSingleShopRentShopCategory : ""
func (d *Daos) GetSingleShopRentShopCategory(ctx *models.Context, UniqueID string) (*models.RefShopRentShopCategory, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSHOPRENTSHOPCATEGORY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var shopcategorys []models.RefShopRentShopCategory
	var shopcategory *models.RefShopRentShopCategory
	if err = cursor.All(ctx.CTX, &shopcategorys); err != nil {
		return nil, err
	}
	if len(shopcategorys) > 0 {
		shopcategory = &shopcategorys[0]
	}
	return shopcategory, nil
}

// UpdateShopRentShopCategory : ""
func (d *Daos) UpdateShopRentShopCategory(ctx *models.Context, shopcategory *models.ShopRentShopCategory) error {
	selector := bson.M{"uniqueId": shopcategory.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	data := bson.M{"$set": shopcategory}
	_, err := ctx.DB.Collection(constants.COLLECTIONSHOPRENTSHOPCATEGORY).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// EnableShopRentShopCategory : ""
func (d *Daos) EnableShopRentShopCategory(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.SHOPRENTSHOPCATEGORYSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONSHOPRENTSHOPCATEGORY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DisableShopRentShopCategory : ""
func (d *Daos) DisableShopRentShopCategory(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.SHOPRENTSHOPCATEGORYSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONSHOPRENTSHOPCATEGORY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DeleteShopRentShopCategory : ""
func (d *Daos) DeleteShopRentShopCategory(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.SHOPRENTSHOPCATEGORYSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONSHOPRENTSHOPCATEGORY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterShopRentShopCategory : ""
func (d *Daos) FilterShopRentShopCategory(ctx *models.Context, shopcategoryfilter *models.ShopRentShopCategoryFilter, pagination *models.Pagination) ([]models.RefShopRentShopCategory, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if shopcategoryfilter != nil {
		if len(shopcategoryfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": shopcategoryfilter.Status}})
		}
		if len(shopcategoryfilter.UniqueID) > 0 {
			query = append(query, bson.M{"uniqueId": bson.M{"$in": shopcategoryfilter.UniqueID}})
		}
		if len(shopcategoryfilter.Name) > 0 {
			query = append(query, bson.M{"name": bson.M{"$in": shopcategoryfilter.Name}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONSHOPRENTSHOPCATEGORY).CountDocuments(ctx.CTX, func() bson.M {
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
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSHOPRENTSHOPCATEGORY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var shopcategorys []models.RefShopRentShopCategory
	if err = cursor.All(context.TODO(), &shopcategorys); err != nil {
		return nil, err
	}
	return shopcategorys, nil
}
