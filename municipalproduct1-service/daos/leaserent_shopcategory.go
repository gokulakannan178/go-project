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

// SaveLeaseRentShopCategory : ""
func (d *Daos) SaveLeaseRentShopCategory(ctx *models.Context, shopcategory *models.LeaseRentShopCategory) error {
	d.Shared.BsonToJSONPrint(shopcategory)
	_, err := ctx.DB.Collection(constants.COLLECTIONLEASERENTSHOPCATEGORY).InsertOne(ctx.CTX, shopcategory)
	return err
}

// GetSingleLeaseRentShopCategory : ""
func (d *Daos) GetSingleLeaseRentShopCategory(ctx *models.Context, UniqueID string) (*models.LeaseRentShopCategory, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONLEASERENTSHOPCATEGORY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var shopcategorys []models.LeaseRentShopCategory
	var shopcategory *models.LeaseRentShopCategory
	if err = cursor.All(ctx.CTX, &shopcategorys); err != nil {
		return nil, err
	}
	if len(shopcategorys) > 0 {
		shopcategory = &shopcategorys[0]
	}
	return shopcategory, nil
}

// UpdateLeaseRentShopCategory : ""
func (d *Daos) UpdateLeaseRentShopCategory(ctx *models.Context, shopcategory *models.LeaseRentShopCategory) error {
	selector := bson.M{"uniqueId": shopcategory.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	data := bson.M{"$set": shopcategory}
	_, err := ctx.DB.Collection(constants.COLLECTIONLEASERENTSHOPCATEGORY).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// EnableLeaseRentShopCategory : ""
func (d *Daos) EnableLeaseRentShopCategory(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.LEASERENTSHOPCATEGORYSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONLEASERENTSHOPCATEGORY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DisableLeaseRentShopCategory : ""
func (d *Daos) DisableLeaseRentShopCategory(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.LEASERENTSHOPCATEGORYSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONLEASERENTSHOPCATEGORY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DeleteLeaseRentShopCategory : ""
func (d *Daos) DeleteLeaseRentShopCategory(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.LEASERENTSHOPCATEGORYSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONLEASERENTSHOPCATEGORY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterLeaseRentShopCategory : ""
func (d *Daos) FilterLeaseRentShopCategory(ctx *models.Context, shopcategoryfilter *models.LeaseRentShopCategoryFilter, pagination *models.Pagination) ([]models.LeaseRentShopCategory, error) {
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONLEASERENTSHOPCATEGORY).CountDocuments(ctx.CTX, func() bson.M {
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
	cursor, err := ctx.DB.Collection(constants.COLLECTIONLEASERENTSHOPCATEGORY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var shopcategorys []models.LeaseRentShopCategory
	if err = cursor.All(context.TODO(), &shopcategorys); err != nil {
		return nil, err
	}
	return shopcategorys, nil
}
