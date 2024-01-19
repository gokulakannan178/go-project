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

//SaveAssetPropertys :""
func (d *Daos) SaveAssetPropertys(ctx *models.Context, assetPropertys *models.AssetPropertys) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONASSETPROPERTYS).InsertOne(ctx.CTX, assetPropertys)
	if err != nil {
		return err
	}
	assetPropertys.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//GetSingleAssetPropertys : ""
func (d *Daos) GetSingleAssetPropertys(ctx *models.Context, uniqueID string) (*models.RefAssetPropertys, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	// lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONASSETTYPE, "assetTypeId", "uniqueId", "ref.assetTypeId", "ref.assetTypeId")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONASSETPROPERTYS).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var AssetPropertys []models.RefAssetPropertys
	var AssetPolicyAsset *models.RefAssetPropertys
	if err = cursor.All(ctx.CTX, &AssetPropertys); err != nil {
		return nil, err
	}
	if len(AssetPropertys) > 0 {
		AssetPolicyAsset = &AssetPropertys[0]
	}
	return AssetPolicyAsset, nil
}

//UpdateAssetPropertys : ""
func (d *Daos) UpdateAssetPropertys(ctx *models.Context, assetPropertys *models.AssetPropertys) error {
	selector := bson.M{"uniqueId": assetPropertys.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": assetPropertys}
	_, err := ctx.DB.Collection(constants.COLLECTIONASSETPROPERTYS).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//EnableAssetPropertys :""
func (d *Daos) EnableAssetPropertys(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.ASSETPROPERTYSSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONASSETPROPERTYS).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableAssetPropertys :""
func (d *Daos) DisableAssetPropertys(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.ASSETPROPERTYSSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONASSETPROPERTYS).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteAssetPropertys :""
func (d *Daos) DeleteAssetPropertys(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.ASSETPROPERTYSSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONASSETPROPERTYS).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//FilterAssetPropertys : ""
func (d *Daos) FilterAssetPropertys(ctx *models.Context, assetPropertysFilter *models.FilterAssetPropertys, pagination *models.Pagination) ([]models.RefAssetPropertys, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if assetPropertysFilter != nil {

		if len(assetPropertysFilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": assetPropertysFilter.Status}})
		}
		if len(assetPropertysFilter.OrganisationID) > 0 {
			query = append(query, bson.M{"organisationId": bson.M{"$in": assetPropertysFilter.OrganisationID}})
		}
		//Regex
		if assetPropertysFilter.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: assetPropertysFilter.Regex.Name, Options: "xi"}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	if assetPropertysFilter != nil {
		if assetPropertysFilter.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{assetPropertysFilter.SortBy: assetPropertysFilter.SortOrder}})
		}
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONASSETPROPERTYS).CountDocuments(ctx.CTX, func() bson.M {
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

	//Aggregation
	d.Shared.BsonToJSONPrintTag("AssetPropertys query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONASSETPROPERTYS).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var assetPropertys []models.RefAssetPropertys
	if err = cursor.All(context.TODO(), &assetPropertys); err != nil {
		return nil, err
	}
	return assetPropertys, nil
}
func (d *Daos) AssetPropertysRemoveNotPresentValue(ctx *models.Context, assetID string, AssetPropertyId []string) error {
	selector := bson.M{"assetId": assetID, "uniqueId": bson.M{"$nin": AssetPropertyId}}
	d.Shared.BsonToJSONPrintTag("selector query in onboarding checklist =>", selector)
	data := bson.M{"$set": bson.M{"status": constants.ASSETPROPERTYSSTATUSDELETED}}
	d.Shared.BsonToJSONPrintTag("data query in onboarding checklist =>", data)
	_, err := ctx.DB.Collection(constants.COLLECTIONASSETPROPERTYS).UpdateMany(ctx.CTX, selector, data)
	return err
}

func (d *Daos) AssetPropertysUpsert(ctx *models.Context, asset *models.Asset) error {
	for _, v := range asset.AssetPropertysId {
		opts := options.Update().SetUpsert(true)
		updateQuery := bson.M{"assetId": asset.UniqueID, "uniqueId": v.UniqueID, "organisationId": asset.OrganisationID, "assetTypeId": asset.AssetTypeId}
		fmt.Println("updateQuery===>", updateQuery)
		assetPropertys := new(models.AssetPropertys)
		assetPropertys.Status = constants.ASSETPROPERTYSSTATUSACTIVE
		assetPropertys.Name = asset.Name
		assetPropertys.Description = asset.Description
		assetPropertys.OrganisationID = asset.OrganisationID
		assetPropertys.AssetTypeID = asset.AssetTypeId
		assetPropertys.Value = v.Value
		//	assetPropertys.UniqueID = d.GetUniqueID(ctx, constants.COLLECTIONASSETPROPERTYS)
		fmt.Println("present added =======>", assetPropertys.UniqueID)
		updateData := bson.M{"$set": assetPropertys}
		if _, err := ctx.DB.Collection(constants.COLLECTIONASSETPROPERTYS).UpdateMany(ctx.CTX, updateQuery, updateData, opts); err != nil {
			return errors.New("Error in updating log - " + err.Error())
		}
	}
	return nil
}
