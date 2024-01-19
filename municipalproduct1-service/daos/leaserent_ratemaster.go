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

//SaveLeaseRentRateMaster :""
func (d *Daos) SaveLeaseRentRateMaster(ctx *models.Context, ratemaster *models.LeaseRentRateMaster) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONLEASERENTRATEMASTER).InsertOne(ctx.CTX, ratemaster)
	return err
}

//GetSingleLeaseRentRateMaster : ""
func (d *Daos) GetSingleLeaseRentRateMaster(ctx *models.Context, UniqueID string) (*models.RefLeaseRentRateMaster, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})
	//Lookups
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONLEASERENTSHOPCATEGORY, "shopCategoryId", "code", "ref.leaseRentShopCategory", "ref.leaseRentShopCategory")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONLEASERENTSHOPSUBCATEGORY, "shopSubCategoryId", "code", "ref.leaseRentShopSubCategory", "ref.leaseRentShopSubCategory")...)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONLEASERENTRATEMASTER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var ratemasters []models.RefLeaseRentRateMaster
	var ratemaster *models.RefLeaseRentRateMaster
	if err = cursor.All(ctx.CTX, &ratemasters); err != nil {
		return nil, err
	}
	if len(ratemasters) > 0 {
		ratemaster = &ratemasters[0]
	}
	return ratemaster, nil
}

//UpdateLeaseRentRateMaster : ""
func (d *Daos) UpdateLeaseRentRateMaster(ctx *models.Context, ratemaster *models.LeaseRentRateMaster) error {
	selector := bson.M{"uniqueId": ratemaster.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": ratemaster, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONLEASERENTRATEMASTER).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterLeaseRentRateMaster : ""
func (d *Daos) FilterLeaseRentRateMaster(ctx *models.Context, ratemasterfilter *models.LeaseRentRateMasterFilter, pagination *models.Pagination) ([]models.RefLeaseRentRateMaster, error) {
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
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONLEASERENTRATEMASTER).CountDocuments(ctx.CTX, func() bson.M {
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
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONLEASERENTSHOPCATEGORY, "shopCategoryId", "code", "ref.leaseRentShopCategory", "ref.leaseRentShopCategory")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONLEASERENTSHOPSUBCATEGORY, "shopSubCategoryId", "code", "ref.leaseRentShopSubCategory", "ref.leaseRentShopSubCategory")...)
	//Aggregation
	d.Shared.BsonToJSONPrintTag("ratemaster query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONLEASERENTRATEMASTER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var ratemasters []models.RefLeaseRentRateMaster
	if err = cursor.All(context.TODO(), &ratemasters); err != nil {
		return nil, err
	}
	return ratemasters, nil
}

//EnableLeaseRentRateMaster :""
func (d *Daos) EnableLeaseRentRateMaster(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.LEASERENTRATEMASTERSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONLEASERENTRATEMASTER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableLeaseRentRateMaster :""
func (d *Daos) DisableLeaseRentRateMaster(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.LEASERENTRATEMASTERSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONLEASERENTRATEMASTER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteLeaseRentRateMaster :""
func (d *Daos) DeleteLeaseRentRateMaster(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.LEASERENTRATEMASTERSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONLEASERENTRATEMASTER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
