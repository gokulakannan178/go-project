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
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//SavePropertyOwner :""
func (d *Daos) SavePropertyOwner(ctx *models.Context, propertyOwners []models.PropertyOwner) error {
	isertData := []interface{}{}
	for _, v := range propertyOwners {
		isertData = append(isertData, v)
	}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYOWNER).InsertMany(ctx.SC, isertData)
	return err
}

//SavePropertyOwner :""
func (d *Daos) SavePropertyOwnerV2(ctx *models.Context, db *mongo.Database, sc *mongo.SessionContext, propertyOwners []models.PropertyOwner) error {
	insertData := []interface{}{}
	for _, v := range propertyOwners {
		insertData = append(insertData, v)
	}
	_, err := db.Collection(constants.COLLECTIONPROPERTYOWNER).InsertMany(ctx.CTX, insertData)
	return err
}

//SavePropertyOwner :""
func (d *Daos) SavePropertyOwnerV3(db *mongo.Database, sc mongo.SessionContext, propertyOwner *models.PropertyOwner) error {
	_, err := db.Collection(constants.COLLECTIONPROPERTYOWNER).InsertOne(sc, propertyOwner)
	return err
}

//GetSinglePropertyOwner : ""
func (d *Daos) GetSinglePropertyOwner(ctx *models.Context, UniqueID string) (*models.RefPropertyOwner, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYOWNER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var propertyOwners []models.RefPropertyOwner
	var propertyOwner *models.RefPropertyOwner
	if err = cursor.All(ctx.CTX, &propertyOwners); err != nil {
		return nil, err
	}
	if len(propertyOwners) > 0 {
		propertyOwner = &propertyOwners[0]
	}
	return propertyOwner, nil
}

//UpdatePropertyOwner : ""
func (d *Daos) UpdatePropertyOwner(ctx *models.Context, propertyOwner *models.PropertyOwner) error {
	selector := bson.M{"uniqueId": propertyOwner.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": propertyOwner, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYOWNER).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//UpdatePropertyOwnerV2 : ""
func (d *Daos) UpdatePropertyOwnerV2(ctx *models.Context, propertyOwners []models.PropertyOwner) error {
	ownerIDs := []string{}
	if len(propertyOwners) > 0 {
		for _, v := range propertyOwners {
			ownerIDs = append(ownerIDs, v.UniqueID)
			selector := bson.M{"uniqueId": v.UniqueID}
			opts := options.Update().SetUpsert(true)
			data := bson.M{"$set": v}
			if _, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYOWNER).UpdateOne(ctx.CTX, selector, data, opts); err != nil {
				return errors.New("Error in updating owner : " + "unique ID => " + v.UniqueID + " floor no => " + v.Name + " - " + err.Error())
			}
		}
		deleteOpts := options.Update().SetUpsert(false)
		deleteSelector := bson.M{"uniqueId": bson.M{"$nin": ownerIDs}}
		deleteData := bson.M{"$set": bson.M{"status": constants.PROPERTYOWNERSTATUSDELETED}}
		if _, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYOWNER).UpdateOne(ctx.CTX, deleteSelector, deleteData, deleteOpts); err != nil {
			return errors.New("Error in updating deleted owner - " + err.Error())
		}
	}
	return nil
}

//FilterPropertyOwner : ""
func (d *Daos) FilterPropertyOwner(ctx *models.Context, propertyOwnerfilter *models.PropertyOwnerFilter, pagination *models.Pagination) ([]models.RefPropertyOwner, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if propertyOwnerfilter != nil {

		if len(propertyOwnerfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": propertyOwnerfilter.Status}})
		}
		if len(propertyOwnerfilter.Mobile) > 0 {
			query = append(query, bson.M{"mobile": bson.M{"$in": propertyOwnerfilter.Mobile}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYOWNER).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("propertyOwner query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYOWNER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var propertyOwners []models.RefPropertyOwner
	if err = cursor.All(context.TODO(), &propertyOwners); err != nil {
		return nil, err
	}
	return propertyOwners, nil
}

//EnablePropertyOwner :""
func (d *Daos) EnablePropertyOwner(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PROPERTYOWNERSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYOWNER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisablePropertyOwner :""
func (d *Daos) DisablePropertyOwner(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PROPERTYOWNERSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYOWNER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeletePropertyOwner :""
func (d *Daos) DeletePropertyOwner(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PROPERTYOWNERSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYOWNER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

func (d *Daos) GetOwnersOfProperty(ctx *models.Context, propertyID string) ([]models.RefPropertyOwner, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	query = append(query, bson.M{"status": bson.M{"$in": []string{constants.PROPERTYOWNERSTATUSACTIVE}}})
	query = append(query, bson.M{"propertyId": bson.M{"$in": []string{propertyID}}})
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	//Aggregation
	d.Shared.BsonToJSONPrintTag("propertyOwner query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYOWNER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var propertyOwners []models.RefPropertyOwner
	if err = cursor.All(context.TODO(), &propertyOwners); err != nil {
		return nil, err
	}
	return propertyOwners, nil
}

// GetPropertyIDsWithOwnerNames : ""
func (d *Daos) GetPropertyIDsWithOwnerNames(ctx *models.Context, propertyOwnerName string) ([]string, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"status": constants.PROPERTYOWNERSTATUSACTIVE,
		"name": primitive.Regex{Pattern: propertyOwnerName, Options: "sxi"},
	}})
	mainPipeline = append(mainPipeline, bson.M{
		"$group": bson.M{"_id": "$propertyId"},
	})
	mainPipeline = append(mainPipeline, bson.M{
		"$group": bson.M{"_id": nil, "propertyIds": bson.M{"$push": "$_id"}},
	})
	var propertyIDsWithOwnerNames []models.PropertyIDsWithOwnerNames
	// var data models.PropertyIDsWithOwnerNames
	d.Shared.BsonToJSONPrintTag("propertyOwner query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYOWNER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	if err = cursor.All(context.TODO(), &propertyIDsWithOwnerNames); err != nil {
		return nil, err
	}
	if len(propertyIDsWithOwnerNames) > 0 {
		return propertyIDsWithOwnerNames[0].PropertyIDs, nil
	}
	return []string{}, nil
}

// GetPropertyIDsWithRegex : ""
func (d *Daos) GetPropertyIDsWithRegex(ctx *models.Context, key, value string) ([]string, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"status": constants.PROPERTYOWNERSTATUSACTIVE,
		key: primitive.Regex{Pattern: value, Options: "sxi"},
	}})
	mainPipeline = append(mainPipeline, bson.M{
		"$group": bson.M{"_id": "$propertyId"},
	})
	mainPipeline = append(mainPipeline, bson.M{
		"$group": bson.M{"_id": nil, "propertyIds": bson.M{"$push": "$_id"}},
	})
	var propertyIDsWithOwnerNames []models.PropertyIDsWithOwnerNames
	// var data models.PropertyIDsWithOwnerNames
	d.Shared.BsonToJSONPrintTag("propertyOwner query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYOWNER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	if err = cursor.All(context.TODO(), &propertyIDsWithOwnerNames); err != nil {
		return nil, err
	}
	if len(propertyIDsWithOwnerNames) > 0 {
		return propertyIDsWithOwnerNames[0].PropertyIDs, nil
	}
	return []string{}, nil
}

// GetPropertyIDsWithMobileNos : ""
func (d *Daos) GetPropertyIDsWithMobileNos(ctx *models.Context, propertyMobileNo string) ([]string, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"status": constants.PROPERTYOWNERSTATUSACTIVE,
		"mobile": primitive.Regex{Pattern: propertyMobileNo, Options: "sxi"},
	}})
	mainPipeline = append(mainPipeline, bson.M{
		"$group": bson.M{"_id": "$propertyId"},
	})
	mainPipeline = append(mainPipeline, bson.M{
		"$group": bson.M{"_id": nil, "propertyIds": bson.M{"$push": "$_id"}},
	})
	var propertyIDsWithMobileNos []models.PropertyIDsWithMobileNos
	// var data models.PropertyIDsWithOwnerNames
	d.Shared.BsonToJSONPrintTag("propertyOwner query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYOWNER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	if err = cursor.All(context.TODO(), &propertyIDsWithMobileNos); err != nil {
		return nil, err
	}
	if len(propertyIDsWithMobileNos) > 0 {
		return propertyIDsWithMobileNos[0].PropertyIDs, nil
	}
	return []string{}, nil
}

//GetSinglePropertyOwnerWithMobileNo : ""
func (d *Daos) GetSinglePropertyOwnerWithMobileNo(ctx *models.Context, MobileNo string) (*models.RefPropertyOwner, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"mobile": MobileNo}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYOWNER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var propertyOwners []models.RefPropertyOwner
	var propertyOwner *models.RefPropertyOwner
	if err = cursor.All(ctx.CTX, &propertyOwners); err != nil {
		return nil, err
	}
	if len(propertyOwners) > 0 {
		propertyOwner = &propertyOwners[0]
	}
	return propertyOwner, nil
}

// UpdateOwnerPropertyID :""
func (d *Daos) UpdateOwnerPropertyID(ctx *models.Context, uniqueIds *models.UpdatePropertyUniqueID) error {
	query := bson.M{"propertyId": uniqueIds.UniqueID}
	update := bson.M{"$set": bson.M{"oldPropertyId": uniqueIds.OldUniqueID, "newPropertyId": uniqueIds.NewUniqueID}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYOWNER).UpdateMany(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
