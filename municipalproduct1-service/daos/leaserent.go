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

//SaveLeaseRent :""
func (d *Daos) SaveLeaseRent(ctx *models.Context, leaserent *models.LeaseRent) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONLEASERENT).InsertOne(ctx.CTX, leaserent)
	return err
}

//GetSingleLeaseRent : ""
func (d *Daos) GetSingleLeaseRent(ctx *models.Context, UniqueID string) (*models.RefLeaseRent, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})
	//Lookups
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONLEASERENTSHOPCATEGORY, "shopCategoryId", "code", "ref.leaseRentShopCategory", "ref.leaseRentShopCategory")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONLEASERENTSHOPSUBCATEGORY, "shopSubCategoryId", "code", "ref.leaseRentShopSubCategory", "ref.leaseRentShopSubCategory")...)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONLEASERENT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var leaserents []models.RefLeaseRent
	var leaserent *models.RefLeaseRent
	if err = cursor.All(ctx.CTX, &leaserents); err != nil {
		return nil, err
	}
	if len(leaserents) > 0 {
		leaserent = &leaserents[0]
	}
	return leaserent, nil
}

//UpdateLeaseRent : ""
func (d *Daos) UpdateLeaseRent(ctx *models.Context, leaserent *models.LeaseRent) error {
	selector := bson.M{"uniqueId": leaserent.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": leaserent, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONLEASERENT).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterLeaseRent : ""
func (d *Daos) FilterLeaseRent(ctx *models.Context, leaserentfilter *models.LeaseRentFilter, pagination *models.Pagination) ([]models.RefLeaseRent, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if leaserentfilter != nil {
		if len(leaserentfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": leaserentfilter.Status}})
		}
		if len(leaserentfilter.UniqueID) > 0 {
			query = append(query, bson.M{"uniqueId": bson.M{"$in": leaserentfilter.UniqueID}})
		}
		if len(leaserentfilter.ShopCategoryID) > 0 {
			query = append(query, bson.M{"shopCategoryId": bson.M{"$in": leaserentfilter.ShopCategoryID}})
		}
		if len(leaserentfilter.ShopSubCategoryID) > 0 {
			query = append(query, bson.M{"shopSubCategoryId": bson.M{"$in": leaserentfilter.ShopSubCategoryID}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONLEASERENT).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("leaserent query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONLEASERENT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var leaserents []models.RefLeaseRent
	if err = cursor.All(context.TODO(), &leaserents); err != nil {
		return nil, err
	}
	return leaserents, nil
}

//EnableLeaseRent :""
func (d *Daos) EnableLeaseRent(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.LEASERENTSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONLEASERENT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableLeaseRent :""
func (d *Daos) DisableLeaseRent(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.LEASERENTSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONLEASERENT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteLeaseRent :""
func (d *Daos) DeleteLeaseRent(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.LEASERENTSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONLEASERENT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
