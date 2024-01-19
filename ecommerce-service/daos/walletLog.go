package daos

import (
	"context"
	"ecommerce-service/constants"
	"ecommerce-service/models"
	"errors"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//SaveWalletLog :""
func (d *Daos) SaveWalletLog(ctx *models.Context, WalletLog *models.WalletLog) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONWALLETLOG).InsertOne(ctx.CTX, WalletLog)
	if err != nil {
		return err
	}
	WalletLog.ID = res.InsertedID.(primitive.ObjectID)
	return nil

}

// func (d *Daos) UpsertWalletLog(ctx *models.Context, WalletLog *models.WalletLog) error {

// 	d.Shared.BsonToJSONPrint(WalletLog)

// 	opts := options.Update().SetUpsert(true)

// 	query := bson.M{"mobile": WalletLog.Mobile}
// 	update := bson.M{"$set": WalletLog}
// 	_, err := ctx.DB.Collection(constants.COLLECTIONINVENTORY).UpdateOne(ctx.CTX, query, update, opts)
// 	if err != nil {
// 		return err
// 	}

// 	return nil

// }

//GetSingleWalletLog : ""
func (d *Daos) GetSingleWalletLog(ctx *models.Context, UniqueID string) (*models.RefWalletLog, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONWALLETLOG).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var WalletLogs []models.RefWalletLog
	var WalletLog *models.RefWalletLog
	if err = cursor.All(ctx.CTX, &WalletLogs); err != nil {
		return nil, err
	}
	if len(WalletLogs) > 0 {
		WalletLog = &WalletLogs[0]
	}
	return WalletLog, nil
}

// //GetSingleGetUsingMobileNumber : ""
// func (d *Daos) GetSingleGetUsingMobileNumber(ctx *models.Context, Mobile string) (*models.RefWalletLog, error) {
// 	mainPipeline := []bson.M{}
// 	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"mobile": Mobile}})
// 	//Aggregation
// 	cursor, err := ctx.DB.Collection(constants.COLLECTIONWalletLog).Aggregate(ctx.CTX, mainPipeline, nil)
// 	if err != nil {
// 		return nil, err
// 	}
// 	var WalletLogs []models.RefWalletLog
// 	var WalletLog *models.RefWalletLog
// 	if err = cursor.All(ctx.CTX, &WalletLogs); err != nil {
// 		return nil, err
// 	}
// 	if len(WalletLogs) > 0 {
// 		WalletLog = &WalletLogs[0]
// 	}
// 	return WalletLog, nil
// }

//UpdateWalletLog : ""
func (d *Daos) UpdateWalletLog(ctx *models.Context, WalletLog *models.WalletLog) error {
	selector := bson.M{"uniqueId": WalletLog.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": WalletLog, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONWALLETLOG).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterWalletLog : ""
func (d *Daos) FilterWalletLog(ctx *models.Context, WalletLogfilter *models.WalletLogFilter, pagination *models.Pagination) ([]models.RefWalletLog, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if WalletLogfilter != nil {
		if len(WalletLogfilter.Regex.UniqueID) > 0 {
			query = append(query, bson.M{"uniqueId": bson.M{"$in": WalletLogfilter.Regex.UniqueID}})
		}
		if len(WalletLogfilter.Regex.UserType) > 0 {
			query = append(query, bson.M{"userType": bson.M{"$in": WalletLogfilter.Regex.UniqueID}})
		}
		if len(WalletLogfilter.Regex.UserID) > 0 {
			query = append(query, bson.M{"userId": bson.M{"$in": WalletLogfilter.Regex.UserID}})
		}
		if len(WalletLogfilter.Regex.WalletID) > 0 {
			query = append(query, bson.M{"walletId": bson.M{"$in": WalletLogfilter.Regex.WalletID}})
		}
		if len(WalletLogfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": WalletLogfilter.Status}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONWALLETLOG).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("WalletLog query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONWALLETLOG).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var WalletLogs []models.RefWalletLog
	if err = cursor.All(context.TODO(), &WalletLogs); err != nil {
		return nil, err
	}
	return WalletLogs, nil
}

//EnableWalletLog :""
func (d *Daos) EnableWalletLog(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.WALLETLOGSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONWALLETLOG).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableWalletLog :""
func (d *Daos) DisableWalletLog(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.WALLETLOGSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONWALLETLOG).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteWalletLog :""
func (d *Daos) DeleteWalletLog(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.WALLETLOGSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONWALLETLOG).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
