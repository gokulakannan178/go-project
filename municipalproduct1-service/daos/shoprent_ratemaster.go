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

//SaveShopRentRateMaster :""
func (d *Daos) SaveShopRentRateMaster(ctx *models.Context, ratemaster *models.ShopRentRateMaster) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONSHOPRENTRATEMASTER).InsertOne(ctx.CTX, ratemaster)
	return err
}

//GetSingleShopRentRateMaster : ""
func (d *Daos) GetSingleShopRentRateMaster(ctx *models.Context, UniqueID string) (*models.RefShopRentRateMaster, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})
	//Lookups                       shopCategoryId
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSHOPRENTSHOPCATEGORY, "shopCategoryId", "uniqueId", "ref.shopRentShopCategory", "ref.shopRentShopCategory")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSHOPRENTSHOPSUBCATEGORY, "shopSubCategoryId", "uniqueId", "ref.shopRentShopSubCategory", "ref.shopRentShopSubCategory")...)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSHOPRENTRATEMASTER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var ratemasters []models.RefShopRentRateMaster
	var ratemaster *models.RefShopRentRateMaster
	if err = cursor.All(ctx.CTX, &ratemasters); err != nil {
		return nil, err
	}
	if len(ratemasters) > 0 {
		ratemaster = &ratemasters[0]
	}
	return ratemaster, nil
}

//UpdateShopRentRateMaster : ""
func (d *Daos) UpdateShopRentRateMaster(ctx *models.Context, ratemaster *models.ShopRentRateMaster) error {
	selector := bson.M{"uniqueId": ratemaster.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": ratemaster, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONSHOPRENTRATEMASTER).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterShopRentRateMaster : ""
func (d *Daos) FilterShopRentRateMaster(ctx *models.Context, ratemasterfilter *models.ShopRentRateMasterFilter, pagination *models.Pagination) ([]models.RefShopRentRateMaster, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if ratemasterfilter != nil {
		if len(ratemasterfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": ratemasterfilter.Status}})
		}
		if len(ratemasterfilter.UniqueID) > 0 {
			query = append(query, bson.M{"uniqueId": bson.M{"$in": ratemasterfilter.UniqueID}})
		}
		if len(ratemasterfilter.ShopCategoryID) > 0 {
			query = append(query, bson.M{"shopCategoryId": bson.M{"$in": ratemasterfilter.ShopCategoryID}})
		}
		if len(ratemasterfilter.ShopSubCategoryID) > 0 {
			query = append(query, bson.M{"shopSubCategoryId": bson.M{"$in": ratemasterfilter.ShopSubCategoryID}})
		}
		//
		if len(ratemasterfilter.ShopSubCategoryID) > 0 {
			query = append(query, bson.M{"shopRentId": bson.M{"$in": ratemasterfilter.ShopRentID}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONSHOPRENTRATEMASTER).CountDocuments(ctx.CTX, func() bson.M {
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
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSHOPRENTSHOPCATEGORY, "shopCategoryId", "code", "ref.shopRentShopCategory", "ref.shopRentShopCategory")...)
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSHOPRENTSHOPSUBCATEGORY, "shopSubCategoryId", "code", "ref.shopRentShopSubCategory", "ref.shopRentShopSubCategory")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSHOPRENTSHOPCATEGORY, "shopCategoryId", "uniqueId", "ref.shopRentShopCategory", "ref.shopRentShopCategory")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSHOPRENTSHOPSUBCATEGORY, "shopSubCategoryId", "uniqueId", "ref.shopRentShopSubCategory", "ref.shopRentShopSubCategory")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("ratemaster query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSHOPRENTRATEMASTER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var ratemasters []models.RefShopRentRateMaster
	if err = cursor.All(context.TODO(), &ratemasters); err != nil {
		return nil, err
	}
	return ratemasters, nil
}

//EnableShopRentRateMaster :""
func (d *Daos) EnableShopRentRateMaster(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.SHOPRENTRATEMASTERSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONSHOPRENTRATEMASTER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableShopRentRateMaster :""
func (d *Daos) DisableShopRentRateMaster(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.SHOPRENTRATEMASTERSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONSHOPRENTRATEMASTER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteShopRentRateMaster :""
func (d *Daos) DeleteShopRentRateMaster(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.SHOPRENTRATEMASTERSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONSHOPRENTRATEMASTER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
