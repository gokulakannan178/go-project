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

// SaveShopRentShopSubCategory : ""
func (d *Daos) SaveShopRentShopSubCategory(ctx *models.Context, shopsubcategory *models.ShopRentShopSubCategory) error {
	d.Shared.BsonToJSONPrint(shopsubcategory)
	_, err := ctx.DB.Collection(constants.COLLECTIONSHOPRENTSHOPSUBCATEGORY).InsertOne(ctx.CTX, shopsubcategory)
	return err
}

// GetSingleShopRentShopSubCategory : ""
func (d *Daos) GetSingleShopRentShopSubCategory(ctx *models.Context, UniqueID string) (*models.RefShopRentShopSubCategory, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSHOPRENTSHOPCATEGORY, "shopCategoryId", "uniqueId", "ref.shopRentShopCategory", "ref.shopRentShopCategory")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSHOPRENTSHOPSUBCATEGORY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var shopsubcategorys []models.RefShopRentShopSubCategory
	var shopsubcategory *models.RefShopRentShopSubCategory
	if err = cursor.All(ctx.CTX, &shopsubcategorys); err != nil {
		return nil, err
	}
	if len(shopsubcategorys) > 0 {
		shopsubcategory = &shopsubcategorys[0]
	}
	return shopsubcategory, nil
}

// UpdateShopRentShopSubCategory : ""
func (d *Daos) UpdateShopRentShopSubCategory(ctx *models.Context, shopsubcategory *models.ShopRentShopSubCategory) error {
	selector := bson.M{"uniqueId": shopsubcategory.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	data := bson.M{"$set": shopsubcategory}
	_, err := ctx.DB.Collection(constants.COLLECTIONSHOPRENTSHOPSUBCATEGORY).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// EnableShopRentShopSubCategory : ""
func (d *Daos) EnableShopRentShopSubCategory(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.SHOPRENTSHOPSUBCATEGORYSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONSHOPRENTSHOPSUBCATEGORY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DisableShopRentShopSubCategory : ""
func (d *Daos) DisableShopRentShopSubCategory(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.SHOPRENTSHOPSUBCATEGORYSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONSHOPRENTSHOPSUBCATEGORY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DeleteShopRentShopSubCategory : ""
func (d *Daos) DeleteShopRentShopSubCategory(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.SHOPRENTSHOPSUBCATEGORYSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONSHOPRENTSHOPSUBCATEGORY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterShopRentShopSubCategory : ""
func (d *Daos) FilterShopRentShopSubCategory(ctx *models.Context, shopsubcategoryfilter *models.ShopRentShopSubCategoryFilter, pagination *models.Pagination) ([]models.RefShopRentShopSubCategory, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if shopsubcategoryfilter != nil {
		if len(shopsubcategoryfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": shopsubcategoryfilter.Status}})
		}
		if len(shopsubcategoryfilter.UniqueID) > 0 {
			query = append(query, bson.M{"uniqueId": bson.M{"$in": shopsubcategoryfilter.UniqueID}})
		}
		if len(shopsubcategoryfilter.Name) > 0 {
			query = append(query, bson.M{"name": bson.M{"$in": shopsubcategoryfilter.Name}})
		}
		if len(shopsubcategoryfilter.ShopCategoryID) > 0 {
			query = append(query, bson.M{"shopCategoryId": bson.M{"$in": shopsubcategoryfilter.ShopCategoryID}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONSHOPRENTSHOPSUBCATEGORY).CountDocuments(ctx.CTX, func() bson.M {
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
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSHOPRENTSHOPCATEGORY, "shopCategoryId", "uniqueId", "ref.shopRentShopCategory", "ref.shopRentShopCategory")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSHOPRENTSHOPSUBCATEGORY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var shopsubcategorys []models.RefShopRentShopSubCategory
	if err = cursor.All(context.TODO(), &shopsubcategorys); err != nil {
		return nil, err
	}
	return shopsubcategorys, nil
}
