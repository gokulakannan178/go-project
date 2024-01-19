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

//SaveAssetPolicyAssets :""
func (d *Daos) SaveAssetPolicyAssets(ctx *models.Context, assetPolicyAssets *models.AssetPolicyAssets) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONASSETPOLICYASSETS).InsertOne(ctx.CTX, assetPolicyAssets)
	if err != nil {
		return err
	}
	assetPolicyAssets.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//GetSingleAssetPolicyAssets : ""
func (d *Daos) GetSingleAssetPolicyAssets(ctx *models.Context, uniqueID string) (*models.RefAssetPolicyAssets, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	// lookup
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDocumentMuxMasterCATEGORY, "categoryId", "uniqueId", "ref.category", "ref.category")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONASSETPOLICYASSETS).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var AssetPolicyAssets []models.RefAssetPolicyAssets
	var AssetPolicyAsset *models.RefAssetPolicyAssets
	if err = cursor.All(ctx.CTX, &AssetPolicyAssets); err != nil {
		return nil, err
	}
	if len(AssetPolicyAssets) > 0 {
		AssetPolicyAsset = &AssetPolicyAssets[0]
	}
	return AssetPolicyAsset, nil
}

// AssetPolicyAssetsRemoveNotPresentValue : ""
func (d *Daos) AssetPolicyAssetsRemoveNotPresentValue(ctx *models.Context, assetpolicyId string, arrayValue []string) error {
	selector := bson.M{"assetPolicyId": assetpolicyId, "assetMasterId": bson.M{"$nin": arrayValue}}
	//d.Shared.BsonToJSONPrintTag("selector query in AssetPolicyAssets =>", selector)
	data := bson.M{"$set": bson.M{"status": constants.ASSETPOLICYASSETSSTATUSDELETED}}
	//d.Shared.BsonToJSONPrintTag("data query in AssetPolicyAssets  =>", data)
	_, err := ctx.DB.Collection(constants.COLLECTIONASSETPOLICYASSETS).UpdateMany(ctx.CTX, selector, data)
	return err
}

// AssetPolicyAssetsUpsert : ""
func (d *Daos) AssetPolicyAssetsUpsert(ctx *models.Context, assetpolicyId string, arrayValue []string, name string) error {
	//fmt.Println("arrayValue", arrayValue)
	for _, v := range arrayValue {
		opts := options.Update().SetUpsert(true)
		updateQuery := bson.M{"assetPolicyId": assetpolicyId, "assetMasterId": v}
		fmt.Println("updateQuery===>", updateQuery)
		AssetPolicyAssets := new(models.AssetPolicyAssets)
		AssetPolicyAssets.Status = constants.ASSETPOLICYASSETSSTATUSACTIVE
		AssetPolicyAssets.Name = name
		AssetPolicyAssets.UniqueID = d.GetUniqueID(ctx, constants.COLLECTIONASSETPOLICYASSETS)
		//fmt.Println("present added =======>", AssetPolicyAssets.UniqueID)
		updateData := bson.M{"$set": AssetPolicyAssets}
		if _, err := ctx.DB.Collection(constants.COLLECTIONASSETPOLICYASSETS).UpdateMany(ctx.CTX, updateQuery, updateData, opts); err != nil {
			return errors.New("Error in updating log - " + err.Error())
		}
	}
	return nil
}

//UpdateAssetPolicyAssets : ""
func (d *Daos) UpdateAssetPolicyAssets(ctx *models.Context, assetPolicyAssets *models.AssetPolicyAssets) error {
	selector := bson.M{"uniqueId": assetPolicyAssets.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": assetPolicyAssets}
	_, err := ctx.DB.Collection(constants.COLLECTIONASSETPOLICYASSETS).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//EnableAssetPolicyAssets :""
func (d *Daos) EnableAssetPolicyAssets(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.ASSETPOLICYASSETSSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONASSETPOLICYASSETS).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableAssetPolicyAssets :""
func (d *Daos) DisableAssetPolicyAssets(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.ASSETPOLICYASSETSSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONASSETPOLICYASSETS).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteAssetPolicyAssets :""
func (d *Daos) DeleteAssetPolicyAssets(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.ASSETPOLICYASSETSSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONASSETPOLICYASSETS).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//FilterAssetPolicyAssets : ""
func (d *Daos) FilterAssetPolicyAssets(ctx *models.Context, assetPolicyAssetsFilter *models.FilterAssetPolicyAssets, pagination *models.Pagination) ([]models.RefAssetPolicyAssets, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if assetPolicyAssetsFilter != nil {

		if len(assetPolicyAssetsFilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": assetPolicyAssetsFilter.Status}})
		}
		//Regex
		if assetPolicyAssetsFilter.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: assetPolicyAssetsFilter.Regex.Name, Options: "xi"}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	if assetPolicyAssetsFilter != nil {
		if assetPolicyAssetsFilter.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{assetPolicyAssetsFilter.SortBy: assetPolicyAssetsFilter.SortOrder}})
		}
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONASSETPOLICYASSETS).CountDocuments(ctx.CTX, func() bson.M {
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
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDocumentMuxMasterCATEGORY, "categoryId", "uniqueId", "ref.category", "ref.category")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("AssetPolicyAssets query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONASSETPOLICYASSETS).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var AssetPolicyAssets []models.RefAssetPolicyAssets
	if err = cursor.All(context.TODO(), &AssetPolicyAssets); err != nil {
		return nil, err
	}
	return AssetPolicyAssets, nil
}
