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

//SaveAssetMaster :""
func (d *Daos) SaveAssetMaster(ctx *models.Context, assetMaster *models.AssetMaster) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONASSETMASTER).InsertOne(ctx.CTX, assetMaster)
	if err != nil {
		return err
	}
	assetMaster.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//GetSingleAssetMaster : ""
func (d *Daos) GetSingleAssetMaster(ctx *models.Context, uniqueID string) (*models.RefAssetMaster, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	// lookup
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisation", "ref.organisation")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONASSETMASTER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var AssetMasters []models.RefAssetMaster
	var AssetMaster *models.RefAssetMaster
	if err = cursor.All(ctx.CTX, &AssetMasters); err != nil {
		return nil, err
	}
	if len(AssetMasters) > 0 {
		AssetMaster = &AssetMasters[0]
	}
	return AssetMaster, nil
}

//GetSingleAssetMasterWithActive : ""
func (d *Daos) GetSingleAssetMasterWithActive(ctx *models.Context, uniqueID string, Status string) (*models.RefAssetMaster, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID, "status": Status}})
	// lookup
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisation", "ref.organisation")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONASSETMASTER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var AssetMasters []models.RefAssetMaster
	var AssetMaster *models.RefAssetMaster
	if err = cursor.All(ctx.CTX, &AssetMasters); err != nil {
		return nil, err
	}
	if len(AssetMasters) > 0 {
		AssetMaster = &AssetMasters[0]
	}
	return AssetMaster, nil
}

//UpdateAssetMaster : ""
func (d *Daos) UpdateAssetMaster(ctx *models.Context, assetMaster *models.AssetMaster) error {
	selector := bson.M{"uniqueId": assetMaster.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": assetMaster}
	_, err := ctx.DB.Collection(constants.COLLECTIONASSETMASTER).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//EnableAssetMaster :""
func (d *Daos) EnableAssetMaster(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.ASSETMASTERSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONASSETMASTER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableAssetMaster :""
func (d *Daos) DisableAssetMaster(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.ASSETMASTERSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONASSETMASTER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteAssetMaster :""
func (d *Daos) DeleteAssetMaster(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.ASSETMASTERSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONASSETMASTER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//FilterAssetMaster : ""
func (d *Daos) FilterAssetMaster(ctx *models.Context, assetMasterfilter *models.FilterAssetMaster, pagination *models.Pagination) ([]models.RefAssetMaster, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if assetMasterfilter != nil {

		if len(assetMasterfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": assetMasterfilter.Status}})
		}
		if len(assetMasterfilter.Organisation) > 0 {
			query = append(query, bson.M{"organisationId": bson.M{"$in": assetMasterfilter.Organisation}})
		}
		//Regex
		if assetMasterfilter.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: assetMasterfilter.Regex.Name, Options: "xi"}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	if assetMasterfilter != nil {
		if assetMasterfilter.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{assetMasterfilter.SortBy: assetMasterfilter.SortOrder}})
		}
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONASSETMASTER).CountDocuments(ctx.CTX, func() bson.M {
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
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisation", "ref.organisation")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("Asset Master query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONASSETMASTER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var assetMastersFilter []models.RefAssetMaster
	if err = cursor.All(context.TODO(), &assetMastersFilter); err != nil {
		return nil, err
	}
	return assetMastersFilter, nil
}
