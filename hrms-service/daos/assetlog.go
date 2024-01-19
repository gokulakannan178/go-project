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

// SaveAssetLog : ""
func (d *Daos) SaveAssetLog(ctx *models.Context, assetlog *models.AssetLog) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONASSETLOG).InsertOne(ctx.CTX, assetlog)
	if err != nil {
		return err
	}
	assetlog.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//UpdateAssetLog : ""
func (d *Daos) UpdateAssetLog(ctx *models.Context, assetlog *models.AssetLog) error {
	fmt.Println("assetlog.UniqueId===>", assetlog.UniqueID)
	selector := bson.M{"uniqueId": assetlog.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": assetlog}
	_, err := ctx.DB.Collection(constants.COLLECTIONASSETLOG).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// func (d *Daos) UpdateAssetLogUsingUniqueId(ctx *models.Context, uniqueID string) error {
// 	selector := bson.M{"uniqueId": uniqueID}
// 	t := time.Now()
// 	update := models.Updated{}
// 	update.On = &t
// 	update.By = constants.SYSTEM
// 	updateInterface := bson.M{"$set": bson.M{"endDate":}}
// 	_, err := ctx.DB.Collection(constants.COLLECTIONASSETLOG).UpdateOne(ctx.CTX, selector, updateInterface)
// 	if err != nil {
// 		fmt.Println("Not changed", err.Error())
// 		return err
// 	}
// 	return err
// }

// GetSingleAssetLog : ""
func (d *Daos) GetSingleAssetLog(ctx *models.Context, uniqueID string) (*models.RefAssetLog, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	//LookUp
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONASSETLOG).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var assetlogs []models.RefAssetLog
	var assetlog *models.RefAssetLog
	if err = cursor.All(ctx.CTX, &assetlogs); err != nil {
		return nil, err
	}
	if len(assetlogs) > 0 {
		assetlog = &assetlogs[0]
	}
	return assetlog, err
}

// EnableAssetLog : ""
func (d *Daos) EnableAssetLog(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.ASSETLOGSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONASSETLOG).UpdateOne(ctx.CTX, selector, data)
	return err
}

// DisableAssetLog : ""
func (d *Daos) DisableAssetLog(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.ASSETLOGSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONASSETLOG).UpdateOne(ctx.CTX, selector, data)
	return err
}

//DeleteAssetLog :""
func (d *Daos) DeleteAssetLog(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.ASSETLOGSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONASSETLOG).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterAssetLog : ""
func (d *Daos) FilterAssetLog(ctx *models.Context, assetlog *models.FilterAssetLog, pagination *models.Pagination) ([]models.RefAssetLog, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if assetlog != nil {
		if len(assetlog.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": assetlog.Status}})
		}
		//Regex
		if assetlog.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: assetlog.Regex.Name, Options: "xi"}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONASSETLOG).CountDocuments(ctx.CTX, func() bson.M {
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
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("Feature query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONASSETLOG).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var rf []models.RefAssetLog
	if err = cursor.All(context.TODO(), &rf); err != nil {
		return nil, err
	}
	return rf, nil
}
func (d *Daos) SaveAssetAssignUpdert(ctx *models.Context, assetLog *models.AssetLog) error {
	// t := time.Now()
	opts := options.Update().SetUpsert(true)
	updateQuery := bson.M{"uniqueId": assetLog.UniqueID}
	fmt.Println("updateQuery===>", updateQuery)

	//fmt.Println("present added =======>", AssetPolicyAssets.UniqueID)
	updateData := bson.M{"$set": assetLog}
	if _, err := ctx.DB.Collection(constants.COLLECTIONASSETLOG).UpdateOne(ctx.CTX, updateQuery, updateData, opts); err != nil {
		return errors.New("Error in updating log - " + err.Error())
	}

	return nil
}

// GetSingleAssetLog : ""
func (d *Daos) GetSingleAssetLogUsingEmpID(ctx *models.Context, AssetID string) (*models.RefAssetLog, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"assetId": AssetID, "status": constants.ASSETASSIGNSTATUS}})
	//LookUp
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONASSETLOG).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var assetlogs []models.RefAssetLog
	var assetlog *models.RefAssetLog
	if err = cursor.All(ctx.CTX, &assetlogs); err != nil {
		return nil, err
	}
	if len(assetlogs) > 0 {
		assetlog = &assetlogs[0]
	}
	return assetlog, err
}
