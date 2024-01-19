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

// SaveLeaseRentShopSubCategory : ""
func (d *Daos) SaveLeaseRentShopSubCategory(ctx *models.Context, shopsubcategory *models.LeaseRentShopSubCategory) error {
	d.Shared.BsonToJSONPrint(shopsubcategory)
	_, err := ctx.DB.Collection(constants.COLLECTIONLEASERENTSHOPSUBCATEGORY).InsertOne(ctx.CTX, shopsubcategory)
	return err
}

// GetSingleLeaseRentShopSubCategory : ""
func (d *Daos) GetSingleLeaseRentShopSubCategory(ctx *models.Context, UniqueID string) (*models.LeaseRentShopSubCategory, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONLEASERENTSHOPSUBCATEGORY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var shopsubcategorys []models.LeaseRentShopSubCategory
	var shopsubcategory *models.LeaseRentShopSubCategory
	if err = cursor.All(ctx.CTX, &shopsubcategorys); err != nil {
		return nil, err
	}
	if len(shopsubcategorys) > 0 {
		shopsubcategory = &shopsubcategorys[0]
	}
	return shopsubcategory, nil
}

// UpdateLeaseRentShopSubCategory : ""
func (d *Daos) UpdateLeaseRentShopSubCategory(ctx *models.Context, shopsubcategory *models.LeaseRentShopSubCategory) error {
	selector := bson.M{"uniqueId": shopsubcategory.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	data := bson.M{"$set": shopsubcategory}
	_, err := ctx.DB.Collection(constants.COLLECTIONLEASERENTSHOPSUBCATEGORY).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// EnableLeaseRentShopSubCategory : ""
func (d *Daos) EnableLeaseRentShopSubCategory(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.LEASERENTSHOPSUBCATEGORYSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONLEASERENTSHOPSUBCATEGORY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DisableLeaseRentShopSubCategory : ""
func (d *Daos) DisableLeaseRentShopSubCategory(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.LEASERENTSHOPSUBCATEGORYSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONLEASERENTSHOPSUBCATEGORY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DeleteLeaseRentShopSubCategory : ""
func (d *Daos) DeleteLeaseRentShopSubCategory(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.LEASERENTSHOPSUBCATEGORYSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONLEASERENTSHOPSUBCATEGORY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterLeaseRentShopSubCategory : ""
func (d *Daos) FilterLeaseRentShopSubCategory(ctx *models.Context, shopsubcategoryfilter *models.LeaseRentShopSubCategoryFilter, pagination *models.Pagination) ([]models.LeaseRentShopSubCategory, error) {
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
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONLEASERENTSHOPSUBCATEGORY).CountDocuments(ctx.CTX, func() bson.M {
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
	cursor, err := ctx.DB.Collection(constants.COLLECTIONLEASERENTSHOPSUBCATEGORY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var shopsubcategorys []models.LeaseRentShopSubCategory
	if err = cursor.All(context.TODO(), &shopsubcategorys); err != nil {
		return nil, err
	}
	return shopsubcategorys, nil
}
