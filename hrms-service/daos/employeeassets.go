package daos

import (
	"context"
	"errors"
	"fmt"
	"hrms-services/constants"
	"hrms-services/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//SaveEmployeeAssets : ""
func (d *Daos) SaveEmployeeAssets(ctx *models.Context, employeeAssets *models.EmployeeAssets) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEASSETS).InsertOne(ctx.CTX, employeeAssets)
	if err != nil {
		return err
	}
	employeeAssets.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//UpdateEmployeeAssets : ""
func (d *Daos) UpdateEmployeeAssets(ctx *models.Context, employeeAssets *models.EmployeeAssets) error {
	selector := bson.M{"uniqueId": employeeAssets.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": employeeAssets}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEASSETS).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// GetSingleEmployeeAssets : ""
func (d *Daos) GetSingleEmployeeAssets(ctx *models.Context, uniqueID string) (*models.RefEmployeeAssets, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	//LookUp
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONASSETPROPERTYS, "assetPropertyId", "uniqueId", "ref.assetPropertyId", "ref.assetPropertyId")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEASSETS).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var EmployeeAssetss []models.RefEmployeeAssets
	var EmployeeAssets *models.RefEmployeeAssets
	if err = cursor.All(ctx.CTX, &EmployeeAssetss); err != nil {
		return nil, err
	}
	if len(EmployeeAssetss) > 0 {
		EmployeeAssets = &EmployeeAssetss[0]
	}
	return EmployeeAssets, err
}

// EnableEmployeeAssets : ""
func (d *Daos) EnableEmployeeAssets(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.EMPLOYEEASSETSSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEASSETS).UpdateOne(ctx.CTX, selector, data)
	return err
}

// DisableEmployeeAssets : ""
func (d *Daos) DisableEmployeeAssets(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.EMPLOYEEASSETSSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEASSETS).UpdateOne(ctx.CTX, selector, data)
	return err
}

//DeleteEmployeeAssets :""
func (d *Daos) DeleteEmployeeAssets(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.EMPLOYEEASSETSSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEASSETS).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterEmployeeAssets : ""
func (d *Daos) FilterEmployeeAssets(ctx *models.Context, employeeAssets *models.FilterEmployeeAssets, pagination *models.Pagination) ([]models.RefEmployeeAssets, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if employeeAssets != nil {
		if len(employeeAssets.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": employeeAssets.Status}})
		}
		if len(employeeAssets.OrganisationId) > 0 {
			query = append(query, bson.M{"organisationId": bson.M{"$in": employeeAssets.OrganisationId}})
		}
		if len(employeeAssets.EmployeeId) > 0 {
			query = append(query, bson.M{"employeeId": bson.M{"$in": employeeAssets.EmployeeId}})
		}
		//Regex
		if employeeAssets.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: employeeAssets.Regex.Name, Options: "xi"}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEASSETS).CountDocuments(ctx.CTX, func() bson.M {
			if query != nil {
				if len(query) > 0 {
					return bson.M{"$and": query}
				}
			}
			return bson.M{}
		}())
		if err != nil {
			log.Println("Error in getting pagination")
		}
		fmt.Println("count", totalCount)
		pagination.Count = int(totalCount)
		d.Shared.PaginationData(pagination)
	}
	//LookUp
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONASSETPROPERTYS, "assetPropertyId", "uniqueId", "ref.assetPropertyId", "ref.assetPropertyId")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("Feature query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEASSETS).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var EmployeeAssetsFilter []models.RefEmployeeAssets
	if err = cursor.All(context.TODO(), &EmployeeAssetsFilter); err != nil {
		return nil, err
	}
	return EmployeeAssetsFilter, nil
}
