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

//SaveAssetPolicy :""
func (d *Daos) SaveAssetPolicy(ctx *models.Context, assetPolicy *models.AssetPolicy) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONASSETPOLICY).InsertOne(ctx.CTX, assetPolicy)
	if err != nil {
		return err
	}
	assetPolicy.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//GetSingleAssetPolicy : ""
func (d *Daos) GetSingleAssetPolicy(ctx *models.Context, uniqueID string) (*models.RefAssetPolicy, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	// lookup
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)
	mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONASSETPOLICYASSETS, "uniqueId", "AssetPolicyID", "ref.assetPolicyAssetsId", "ref.assetPolicyAssetsId")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONASSETPOLICY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var assetPolicys []models.RefAssetPolicy
	var assetPolicy *models.RefAssetPolicy
	if err = cursor.All(ctx.CTX, &assetPolicys); err != nil {
		return nil, err
	}
	if len(assetPolicys) > 0 {
		assetPolicy = &assetPolicys[0]
	}
	return assetPolicy, nil
}

//UpdateAssetPolicy : ""
func (d *Daos) UpdateAssetPolicy(ctx *models.Context, assetPolicy *models.AssetPolicy) error {
	selector := bson.M{"uniqueId": assetPolicy.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": assetPolicy}
	_, err := ctx.DB.Collection(constants.COLLECTIONASSETPOLICY).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//EnableAssetPolicy :""
func (d *Daos) EnableAssetPolicy(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.ASSETPOLICYSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONASSETPOLICY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableAssetPolicy :""
func (d *Daos) DisableAssetPolicy(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.ASSETPOLICYSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONASSETPOLICY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteAssetPolicy :""
func (d *Daos) DeleteAssetPolicy(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.ASSETPOLICYSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONASSETPOLICY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//FilterAssetPolicy : ""
func (d *Daos) FilterAssetPolicy(ctx *models.Context, assetPolicyFilter *models.FilterAssetPolicy, pagination *models.Pagination) ([]models.RefAssetPolicy, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if assetPolicyFilter != nil {

		if len(assetPolicyFilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": assetPolicyFilter.Status}})
		}
		if len(assetPolicyFilter.OrganisationID) > 0 {
			query = append(query, bson.M{"organisationId": bson.M{"$in": assetPolicyFilter.OrganisationID}})
		}
		//Regex
		if assetPolicyFilter.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: assetPolicyFilter.Regex.Name, Options: "xi"}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	if assetPolicyFilter != nil {
		if assetPolicyFilter.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{assetPolicyFilter.SortBy: assetPolicyFilter.SortOrder}})
		}
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONASSETPOLICY).CountDocuments(ctx.CTX, func() bson.M {
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
	mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONASSETPOLICYASSETS, "uniqueId", "AssetPolicyID", "ref.assetPolicyAssetsId", "ref.assetPolicyAssetsId")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("DocumentType query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONASSETPOLICY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var assetPolicysFilter []models.RefAssetPolicy
	if err = cursor.All(context.TODO(), &assetPolicyFilter); err != nil {
		return nil, err
	}
	return assetPolicysFilter, nil
}
