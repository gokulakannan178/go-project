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

//SaveAsset :""
func (d *Daos) SaveAsset(ctx *models.Context, asset *models.Asset) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONASSET).InsertOne(ctx.CTX, asset)
	if err != nil {
		return err
	}
	asset.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//GetSingleAsset : ""
func (d *Daos) GetSingleAsset(ctx *models.Context, uniqueID string) (*models.RefAsset, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	// lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)
	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"from": constants.COLLECTIONASSETLOG,
			"as":   "ref.assetlog",
			"let":  bson.M{"assetId": "$uniqueId"},
			"pipeline": []bson.M{
				{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
					{"$eq": []string{"$status", constants.ASSETASSIGNSTATUS}},
					{"$eq": []string{"$assetId", "$$assetId"}},
				}}},
				},
			},
		},
	})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"ref.assetlog": bson.M{"$arrayElemAt": []interface{}{"$ref.assetlog", 0}}}})
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONEMPLOYEE, "employeeId", "uniqueId", "ref.employee", "ref.employee")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONASSETTYPE, "assetTypeId", "uniqueId", "ref.assetTypeId", "ref.assetTypeId")...)
	//mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONASSETTYPEPROPERTYS, "assetTypeId", "assetTypeId", "ref.assetTypeId.ref.assetTypePepropertys", "ref.assetTypeId.ref.assetTypePepropertys")...)
	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"from": constants.COLLECTIONASSETPROPERTYS,
			"as":   "ref.assetTypePepropertys",
			"let":  bson.M{"assetId": "$uniqueId", "assetTypeId": "$assetTypeId"},
			"pipeline": []bson.M{
				{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
					///	{"$eq": []string{"$status", constants.ASSETASSIGNSTATUS}},
					{"$eq": []string{"$assetId", "$$assetId"}},
					{"$eq": []string{"$assetTypeId", "$$assetTypeId"}},
				}}},
				},
			},
		},
	})
	d.Shared.BsonToJSONPrintTag("Asset query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONASSET).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var Assets []models.RefAsset
	var Asset *models.RefAsset
	if err = cursor.All(ctx.CTX, &Assets); err != nil {
		return nil, err
	}
	if len(Assets) > 0 {
		Asset = &Assets[0]
	}
	return Asset, nil
}

//UpdateAsset : ""
func (d *Daos) UpdateAsset(ctx *models.Context, asset *models.Asset) error {
	selector := bson.M{"uniqueId": asset.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": asset}
	_, err := ctx.DB.Collection(constants.COLLECTIONASSET).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}
func (d *Daos) UpdateAssetAssign(ctx *models.Context, asset *models.AssetAssign) error {
	selector := bson.M{"uniqueId": asset.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": asset}
	_, err := ctx.DB.Collection(constants.COLLECTIONASSET).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//EnableAsset :""
func (d *Daos) EnableAsset(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.ASSETSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONASSET).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableAsset :""
func (d *Daos) DisableAsset(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.ASSETSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONASSET).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteAsset :""
func (d *Daos) DeleteAsset(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.ASSETSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONASSET).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//FilterAsset : ""
func (d *Daos) FilterAsset(ctx *models.Context, assetFilter *models.FilterAsset, pagination *models.Pagination) ([]models.RefAsset, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if assetFilter != nil {

		if len(assetFilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": assetFilter.Status}})
		}
		if len(assetFilter.OrganisationID) > 0 {
			query = append(query, bson.M{"organisationId": bson.M{"$in": assetFilter.OrganisationID}})
		}
		if len(assetFilter.EmployeeId) > 0 {
			query = append(query, bson.M{"employeeId": bson.M{"$in": assetFilter.EmployeeId}})
		}
		//Regex
		if assetFilter.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: assetFilter.Regex.Name, Options: "xi"}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	if assetFilter != nil {
		if assetFilter.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{assetFilter.SortBy: assetFilter.SortOrder}})
		}
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONASSET).CountDocuments(ctx.CTX, func() bson.M {
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
	// lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONEMPLOYEE, "employeeId", "uniqueId", "ref.employee", "ref.employee")...)

	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)
	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"from": constants.COLLECTIONASSETLOG,
			"as":   "ref.assetlog",
			"let":  bson.M{"assetId": "$uniqueId"},
			"pipeline": []bson.M{
				{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
					{"$eq": []string{"$status", constants.ASSETASSIGNSTATUS}},
					{"$eq": []string{"$assetId", "$$assetId"}},
				}}},
				},
			},
		},
	})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"ref.assetlog": bson.M{"$arrayElemAt": []interface{}{"$ref.assetlog", 0}}}})
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONASSETTYPE, "assetTypeId", "uniqueId", "ref.assetTypeId", "ref.assetTypeId")...)
	//	mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONASSETTYPEPROPERTYS, "assetTypeId", "assetTypeId", "ref.assetTypeId.ref.assetTypePepropertys", "ref.assetTypeId.ref.assetTypePepropertys")...)
	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"from": constants.COLLECTIONASSETPROPERTYS,
			"as":   "ref.assetTypePepropertys",
			"let":  bson.M{"assetId": "$uniqueId", "assetTypeId": "$assetTypeId"},
			"pipeline": []bson.M{
				{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
					//	{"$eq": []string{"$status", constants.ASSETASSIGNSTATUS}},
					{"$eq": []string{"$assetId", "$$assetId"}},
					{"$eq": []string{"$assetTypeId", "$$assetTypeId"}},
				}}},
				},
			},
		},
	})
	//Aggregation
	d.Shared.BsonToJSONPrintTag("Asset query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONASSET).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var assetsFilter []models.RefAsset
	if err = cursor.All(context.TODO(), &assetsFilter); err != nil {
		return nil, err
	}
	return assetsFilter, nil
}
func (d *Daos) AssetAssign(ctx *models.Context, asset *models.AssetAssign) error {
	selector := bson.M{"uniqueId": asset.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": asset}
	_, err := ctx.DB.Collection(constants.COLLECTIONASSET).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}
func (d *Daos) RevokeAsset(ctx *models.Context, asset *models.Asset) error {
	selector := bson.M{"employeeId": asset.EmployeeId}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": asset, "status": constants.ASSETREVOKESTATUS}
	_, err := ctx.DB.Collection(constants.COLLECTIONASSET).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// GetSingleAssetUsingEmpID : ""
func (d *Daos) GetSingleAssetUsingUniqueId(ctx *models.Context, UniqueID string) (*models.RefAsset, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID, "status": bson.M{"$in": []string{constants.ASSETSTATUSACTIVE, constants.ASSETASSIGNSTATUS}}}})
	// lookup
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONASSET).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var Assets []models.RefAsset
	var Asset *models.RefAsset
	if err = cursor.All(ctx.CTX, &Assets); err != nil {
		return nil, err
	}
	if len(Assets) > 0 {
		Asset = &Assets[0]
	}
	return Asset, nil
}
