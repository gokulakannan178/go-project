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

// SaveVendor : ""
func (d *Daos) SaveVendor(ctx *models.Context, vendor *models.Vendor) error {
	d.Shared.BsonToJSONPrint(vendor)
	_, err := ctx.DB.Collection(constants.COLLECTIONVENDOR).InsertOne(ctx.CTX, vendor)
	return err
}

// GetSingleVendor : ""
func (d *Daos) GetSingleVendor(ctx *models.Context, UniqueID string) (*models.RefVendor, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONVENDOR).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var towers []models.RefVendor
	var tower *models.RefVendor
	if err = cursor.All(ctx.CTX, &towers); err != nil {
		return nil, err
	}
	if len(towers) > 0 {
		tower = &towers[0]
	}
	return tower, nil
}

// UpdateVendor : ""
func (d *Daos) UpdateVendor(ctx *models.Context, crop *models.Vendor) error {
	selector := bson.M{"uniqueId": crop.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	data := bson.M{"$set": crop}
	_, err := ctx.DB.Collection(constants.COLLECTIONVENDOR).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// EnableVendor : ""
func (d *Daos) EnableVendor(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.VENDORSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONVENDOR).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DisableVendor : ""
func (d *Daos) DisableVendor(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.VENDORSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONVENDOR).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DeleteVendor : ""
func (d *Daos) DeleteVendor(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.VENDORSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONVENDOR).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterVendor : ""
func (d *Daos) FilterVendor(ctx *models.Context, filter *models.VendorFilter, pagination *models.Pagination) ([]models.RefVendor, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
		if len(filter.UniqueID) > 0 {
			query = append(query, bson.M{"uniqueId": bson.M{"$in": filter.UniqueID}})
		}
		if filter.SearchText.UniqueID != "" {
			query = append(query, bson.M{"uniqueId": primitive.Regex{Pattern: filter.SearchText.UniqueID, Options: "xi"}})

		}
		if filter.SearchText.MobileNo != "" {
			query = append(query, bson.M{"mobileNo": primitive.Regex{Pattern: filter.SearchText.MobileNo, Options: "xi"}})

		}
		if filter.SearchText.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: filter.SearchText.Name, Options: "xi"}})

		}
		if filter.SearchText.StoreName != "" {
			query = append(query, bson.M{"storeName": primitive.Regex{Pattern: filter.SearchText.StoreName, Options: "xi"}})

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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONVENDOR).CountDocuments(ctx.CTX, func() bson.M {
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
	cursor, err := ctx.DB.Collection(constants.COLLECTIONVENDOR).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var vendorCrop []models.RefVendor
	if err = cursor.All(context.TODO(), &vendorCrop); err != nil {
		return nil, err
	}
	return vendorCrop, nil
}

// // GetSingleVendor : ""
// func (d *Daos) GetSingleVendorUsingMoblienumber(ctx *models.Context, UniqueID string) (*models.RefVendor, error) {
// 	mainPipeline := []bson.M{}
// 	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"mobileNo": UniqueID}})
// GetSingleVendorwithMobileNo : ""
func (d *Daos) GetSingleVendorWithMobileNo(ctx *models.Context, MobileNo string) (*models.RefVendor, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"mobileNo": MobileNo}})

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONVENDOR).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var towers []models.RefVendor
	var tower *models.RefVendor
	if err = cursor.All(ctx.CTX, &towers); err != nil {
		return nil, err
	}
	if len(towers) > 0 {
		tower = &towers[0]
	}
	return tower, nil
}

// GetSingleVendorwithMobileNo : ""
func (d *Daos) GetSingleVendorWithMobileNoV2(ctx *models.Context, MobileNo string) (*models.RefVendor, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"mobileNo": MobileNo}})

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONVENDOR).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var towers []models.RefVendor
	var tower *models.RefVendor
	if err = cursor.All(ctx.CTX, &towers); err != nil {
		return nil, err
	}
	if len(towers) > 0 {
		tower = &towers[0]
	}
	return tower, nil
}
