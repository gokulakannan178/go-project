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

// SaveVendorInfo : ""
func (d *Daos) SaveVendorInfo(ctx *models.Context, vendor *models.VendorInfo) error {
	d.Shared.BsonToJSONPrint(vendor)
	_, err := ctx.DB.Collection(constants.COLLECTIONVENDORINFO).InsertOne(ctx.CTX, vendor)
	return err
}

// GetSingleVendorInfo : ""
func (d *Daos) GetSingleVendorInfo(ctx *models.Context, UniqueID string) (*models.RefVendorInfo, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONVENDORINFO).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var towers []models.RefVendorInfo
	var tower *models.RefVendorInfo
	if err = cursor.All(ctx.CTX, &towers); err != nil {
		return nil, err
	}
	if len(towers) > 0 {
		tower = &towers[0]
	}
	return tower, nil
}

// UpdateVendorInfo : ""
func (d *Daos) UpdateVendorInfo(ctx *models.Context, crop *models.VendorInfo) error {
	selector := bson.M{"uniqueId": crop.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	data := bson.M{"$set": crop}
	_, err := ctx.DB.Collection(constants.COLLECTIONVENDORINFO).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// EnableVendorInfo : ""
func (d *Daos) EnableVendorInfo(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.VENDORINFOSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONVENDORINFO).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DisableVendorInfo : ""
func (d *Daos) DisableVendorInfo(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.VENDORINFOSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONVENDORINFO).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DeleteVendorInfo : ""
func (d *Daos) DeleteVendorInfo(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.VENDORINFOSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONVENDORINFO).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterVendorInfo : ""
func (d *Daos) FilterVendorInfo(ctx *models.Context, filter *models.VendorInfoFilter, pagination *models.Pagination) ([]models.RefVendorInfo, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
		if len(filter.UniqueID) > 0 {
			query = append(query, bson.M{"uniqueId": bson.M{"$in": filter.UniqueID}})
		}
		if len(filter.VendorID) > 0 {
			query = append(query, bson.M{"vendorId": bson.M{"$in": filter.VendorID}})
		}
		if filter.SearchText.UniqueID != "" {
			query = append(query, bson.M{"uniqueId": primitive.Regex{Pattern: filter.SearchText.UniqueID, Options: "xi"}})

		}
		if filter.SearchText.VendorID != "" {
			query = append(query, bson.M{"vendorId": primitive.Regex{Pattern: filter.SearchText.VendorID, Options: "xi"}})

		}
		if filter.SearchText.GSTNo != "" {
			query = append(query, bson.M{"gstNo": primitive.Regex{Pattern: filter.SearchText.GSTNo, Options: "xi"}})

		}
		if filter.SearchText.PanNo != "" {
			query = append(query, bson.M{"panNo": primitive.Regex{Pattern: filter.SearchText.PanNo, Options: "xi"}})

		}
		if filter.SearchText.TaxNo != "" {
			query = append(query, bson.M{"taxNo": primitive.Regex{Pattern: filter.SearchText.TaxNo, Options: "xi"}})

		}
		if filter.SearchText.AadhaarNo != "" {
			query = append(query, bson.M{"aadhaarNo": primitive.Regex{Pattern: filter.SearchText.AadhaarNo, Options: "xi"}})

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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONVENDORINFO).CountDocuments(ctx.CTX, func() bson.M {
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
	// Lookup
	if filter != nil {
		if filter.WantVendorInfo {
			// if filter.WantVendorInfo == true {
			mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVENDOR, "vendorId", "uniqueId", "ref.vendor", "ref.vendor")...)
		}
	}

	//Aggregation
	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONVENDORINFO).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var vendorCrop []models.RefVendorInfo
	if err = cursor.All(context.TODO(), &vendorCrop); err != nil {
		return nil, err
	}
	return vendorCrop, nil
}
