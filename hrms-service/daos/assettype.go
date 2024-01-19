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

//SaveAssetType :""
func (d *Daos) SaveAssetType(ctx *models.Context, assetType *models.AssetType) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONASSETTYPE).InsertOne(ctx.CTX, assetType)
	if err != nil {
		return err
	}
	assetType.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}
func (d *Daos) SaveAssetTypeWithPropertys(ctx *models.Context, assetType *models.AssetType) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONASSETTYPE).InsertOne(ctx.CTX, assetType)
	if err != nil {
		return err
	}
	assetType.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}
func (d *Daos) SaveAssetTypeProperty(ctx *models.Context, assetTypePropertys []models.AssetTypePropertys) error {
	var data []interface{}
	//array
	for _, v := range assetTypePropertys {
		data = append(data, v)
	}
	_, err := ctx.DB.Collection(constants.COLLECTIONASSETTYPEPROPERTYS).InsertMany(ctx.CTX, data)
	return err
}

//GetSingleAssetType : ""
func (d *Daos) GetSingleAssetType(ctx *models.Context, uniqueID string) (*models.RefAssetType, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	// lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisation", "ref.organisation")...)
	mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONASSETTYPEPROPERTYS, "uniqueId", "assetTypeId", "ref.assetTypePepropertys", "ref.assetTypePepropertys")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONASSETTYPE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var assetTypes []models.RefAssetType
	var assetType *models.RefAssetType
	if err = cursor.All(ctx.CTX, &assetTypes); err != nil {
		return nil, err
	}
	if len(assetTypes) > 0 {
		assetType = &assetTypes[0]
	}
	return assetType, nil
}

//GetSingleAssetTypeWithActive : ""
func (d *Daos) GetSingleAssetTypeWithActive(ctx *models.Context, uniqueID string, Status string) (*models.RefAssetType, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID, "status": Status}})
	// lookup
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONASSETTYPE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var assetTypes []models.RefAssetType
	var assetType *models.RefAssetType
	if err = cursor.All(ctx.CTX, &assetTypes); err != nil {
		return nil, err
	}
	if len(assetTypes) > 0 {
		assetType = &assetTypes[0]
	}
	return assetType, nil
}

//UpdateAssetType : ""
func (d *Daos) UpdateAssetType(ctx *models.Context, assetType *models.AssetType) error {
	selector := bson.M{"uniqueId": assetType.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": assetType}
	_, err := ctx.DB.Collection(constants.COLLECTIONASSETTYPE).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//EnableAssetType :""
func (d *Daos) EnableAssetType(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.ASSETTYPESTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONASSETTYPE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableAssetType :""
func (d *Daos) DisableAssetType(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.ASSETTYPESTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONASSETTYPE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteAssetType :""
func (d *Daos) DeleteAssetType(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.ASSETTYPESTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONASSETTYPE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//FilterAssetType : ""
func (d *Daos) FilterAssetType(ctx *models.Context, assetTypefilter *models.FilterAssetType, pagination *models.Pagination) ([]models.RefAssetType, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if assetTypefilter != nil {

		if len(assetTypefilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": assetTypefilter.Status}})
		}
		if len(assetTypefilter.OrganisationID) > 0 {
			query = append(query, bson.M{"organisationId": bson.M{"$in": assetTypefilter.OrganisationID}})
		}
		//Regex
		if assetTypefilter.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: assetTypefilter.Regex.Name, Options: "xi"}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	if assetTypefilter != nil {
		if assetTypefilter.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{assetTypefilter.SortBy: assetTypefilter.SortOrder}})
		}
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONASSETTYPE).CountDocuments(ctx.CTX, func() bson.M {
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
	mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONASSETTYPEPROPERTYS, "uniqueId", "assetTypeId", "ref.assetTypePepropertys", "ref.assetTypePepropertys")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("DocumentScenario query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONASSETTYPE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var assetTypesFilter []models.RefAssetType
	if err = cursor.All(context.TODO(), &assetTypesFilter); err != nil {
		return nil, err
	}
	return assetTypesFilter, nil
}

// func (d *Daos) GetAssetTypePropertys(ctx *models.Context, uniqueID string) (*models.RefAssetType, error) {
// 	mainPipeline := []bson.M{}
// 	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
// 	//lookup
// 	mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONASSETTYPEPROPERTYS, "assetTypePropertysId", "uniqueId", "ref.assetTypePropertysId", "ref.assetTypePropertysId")...)

// 	//Aggregation
// 	cursor, err := ctx.DB.Collection(constants.COLLECTIONASSETTYPE).Aggregate(ctx.CTX, mainPipeline, nil)
// 	if err != nil {
// 		return nil, err
// 	}
// 	var assetTypes []models.RefAssetType
// 	var assetType *models.RefAssetType
// 	if err = cursor.All(ctx.CTX, &assetTypes); err != nil {
// 		return nil, err
// 	}
// 	if len(assetTypes) > 0 {
// 		assetType = &assetTypes[0]
// 	}
// 	return assetType, nil
// }
