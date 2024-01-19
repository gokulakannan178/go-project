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
	"go.mongodb.org/mongo-driver/mongo/options"
)

//SaveAssetTypePropertys :""
func (d *Daos) SaveAssetTypePropertys(ctx *models.Context, assetTypePropertys *models.AssetTypePropertys) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONASSETTYPEPROPERTYS).InsertOne(ctx.CTX, assetTypePropertys)
	if err != nil {
		return err
	}
	assetTypePropertys.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}
func (d *Daos) SaveAssetTypePropertysWithUpsert(ctx *models.Context, assetTypePropertys *models.AssetTypePropertys) error {
	opts := options.Update().SetUpsert(true)
	updateQuery := bson.M{"assetTypeId": assetTypePropertys.AssetTypeID, "organisationId": assetTypePropertys.OrganisationID, "uniqueId": assetTypePropertys.UniqueID}
	updateData := bson.M{"$set": assetTypePropertys}
	if _, err := ctx.DB.Collection(constants.COLLECTIONASSETTYPEPROPERTYS).UpdateOne(ctx.CTX, updateQuery, updateData, opts); err != nil {
		return errors.New("Error in updating log - " + err.Error())
	}

	return nil
}

//GetSingleAssetTypePropertys : ""
func (d *Daos) GetSingleAssetTypePropertys(ctx *models.Context, uniqueID string) (*models.RefAssetTypePropertys, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	// lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONASSETTYPE, "assetTypeId", "uniqueId", "ref.assetTypeId", "ref.assetTypeId")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisation", "ref.organisation")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONASSETTYPEPROPERTYS).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var AssetTypePropertys []models.RefAssetTypePropertys
	var AssetPolicyAsset *models.RefAssetTypePropertys
	if err = cursor.All(ctx.CTX, &AssetTypePropertys); err != nil {
		return nil, err
	}
	if len(AssetTypePropertys) > 0 {
		AssetPolicyAsset = &AssetTypePropertys[0]
	}
	return AssetPolicyAsset, nil
}
func (d *Daos) GetSingleAssetTypePropertysWithAssetId(ctx *models.Context, uniqueID string, AssetTypeID string) (*models.RefAssetTypePropertys, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID, "assetTypeId": AssetTypeID}})
	// lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONASSETTYPE, "assetTypeId", "uniqueId", "ref.assetTypeId", "ref.assetTypeId")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisation", "ref.organisation")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONASSETTYPEPROPERTYS).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var AssetTypePropertys []models.RefAssetTypePropertys
	var AssetPolicyAsset *models.RefAssetTypePropertys
	if err = cursor.All(ctx.CTX, &AssetTypePropertys); err != nil {
		return nil, err
	}
	if len(AssetTypePropertys) > 0 {
		AssetPolicyAsset = &AssetTypePropertys[0]
	}
	return AssetPolicyAsset, nil
}

//UpdateAssetTypePropertys : ""
func (d *Daos) UpdateAssetTypePropertys(ctx *models.Context, assetTypePropertys *models.AssetTypePropertys) error {
	selector := bson.M{"uniqueId": assetTypePropertys.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": assetTypePropertys}
	_, err := ctx.DB.Collection(constants.COLLECTIONASSETTYPEPROPERTYS).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//EnableAssetTypePropertys :""
func (d *Daos) EnableAssetTypePropertys(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.ASSETTYPEPROPERTYSSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONASSETTYPEPROPERTYS).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableAssetTypePropertys :""
func (d *Daos) DisableAssetTypePropertys(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.ASSETTYPEPROPERTYSSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONASSETTYPEPROPERTYS).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteAssetTypePropertys :""
func (d *Daos) DeleteAssetTypePropertys(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.ASSETTYPEPROPERTYSSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONASSETTYPEPROPERTYS).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//FilterAssetTypePropertys : ""
func (d *Daos) FilterAssetTypePropertys(ctx *models.Context, assetTypePropertysFilter *models.FilterAssetTypePropertys, pagination *models.Pagination) ([]models.RefAssetTypePropertys, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if assetTypePropertysFilter != nil {

		if len(assetTypePropertysFilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": assetTypePropertysFilter.Status}})
		}
		if len(assetTypePropertysFilter.OrganisationID) > 0 {
			query = append(query, bson.M{"organisationId": bson.M{"$in": assetTypePropertysFilter.OrganisationID}})
		}
		//assetTypeId
		if len(assetTypePropertysFilter.AssetTypeIds) > 0 {
			query = append(query, bson.M{"assetTypeId": bson.M{"$in": assetTypePropertysFilter.AssetTypeIds}})
		}
		//Regex
		if assetTypePropertysFilter.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: assetTypePropertysFilter.Regex.Name, Options: "xi"}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	if assetTypePropertysFilter != nil {
		if assetTypePropertysFilter.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{assetTypePropertysFilter.SortBy: assetTypePropertysFilter.SortOrder}})
		}
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONASSETTYPEPROPERTYS).CountDocuments(ctx.CTX, func() bson.M {
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
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONASSETTYPE, "assetTypeId", "uniqueId", "ref.assetTypeId", "ref.assetTypeId")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisation", "ref.organisation")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("AssetTypePropertys query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONASSETTYPEPROPERTYS).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var assetTypePropertys []models.RefAssetTypePropertys
	if err = cursor.All(context.TODO(), &assetTypePropertys); err != nil {
		return nil, err
	}
	return assetTypePropertys, nil
}
func (d *Daos) GetSingleAssetTypePropertysWithActive(ctx *models.Context, uniqueID string, Status string) (*models.RefAssetTypePropertys, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID, "status": Status}})
	//LookUp
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONASSETTYPEPROPERTYS).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var assetTypePropertys []models.RefAssetTypePropertys
	var assetTypeProperty *models.RefAssetTypePropertys
	if err = cursor.All(ctx.CTX, &assetTypePropertys); err != nil {
		return nil, err
	}
	if len(assetTypePropertys) > 0 {
		assetTypeProperty = &assetTypePropertys[0]
	}
	return assetTypeProperty, err
}
