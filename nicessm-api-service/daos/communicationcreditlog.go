package daos

import (
	"context"
	"errors"
	"fmt"
	"log"
	"nicessm-api-service/constants"
	"nicessm-api-service/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//SaveCommunicationCreditLog :""
func (d *Daos) SaveCommunicationCreditLog(ctx *models.Context, CommunicationCreditLog *models.CommunicationCreditLog) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONCOMMUNICATIONCREDITLOG).InsertOne(ctx.CTX, CommunicationCreditLog)
	if err != nil {
		return err
	}
	CommunicationCreditLog.ID = res.InsertedID.(primitive.ObjectID)
	return nil

}

//GetSingleCommunicationCreditLog : ""
func (d *Daos) GetSingleCommunicationCreditLog(ctx *models.Context, code string) (*models.RefCommunicationCreditLog, error) {
	id, err := primitive.ObjectIDFromHex(code)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONCOMMUNICATIONCREDITLOG).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var CommunicationCreditLogs []models.RefCommunicationCreditLog
	var CommunicationCreditLog *models.RefCommunicationCreditLog
	if err = cursor.All(ctx.CTX, &CommunicationCreditLogs); err != nil {
		return nil, err
	}
	if len(CommunicationCreditLogs) > 0 {
		CommunicationCreditLog = &CommunicationCreditLogs[0]
	}
	return CommunicationCreditLog, nil
}

//UpdateCommunicationCreditLog : ""
func (d *Daos) UpdateCommunicationCreditLog(ctx *models.Context, CommunicationCreditLog *models.CommunicationCreditLog) error {
	selector := bson.M{"_id": CommunicationCreditLog.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": CommunicationCreditLog, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONCOMMUNICATIONCREDITLOG).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterCommunicationCreditLog : ""
func (d *Daos) FilterCommunicationCreditLog(ctx *models.Context, CommunicationCreditLogfilter *models.CommunicationCreditLogFilter, pagination *models.Pagination) ([]models.RefCommunicationCreditLog, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if CommunicationCreditLogfilter != nil {

		if len(CommunicationCreditLogfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": CommunicationCreditLogfilter.Status}})
		}
		if len(CommunicationCreditLogfilter.CommunicationMode) > 0 {
			query = append(query, bson.M{"communicationMode": bson.M{"$in": CommunicationCreditLogfilter.CommunicationMode}})
		}

	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"_id": -1}})

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONCOMMUNICATIONCREDITLOG).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("CommunicationCreditLog query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONCOMMUNICATIONCREDITLOG).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var CommunicationCreditLogs []models.RefCommunicationCreditLog
	if err = cursor.All(context.TODO(), &CommunicationCreditLogs); err != nil {
		return nil, err
	}
	return CommunicationCreditLogs, nil
}

//EnableCommunicationCreditLog :""
func (d *Daos) EnableCommunicationCreditLog(ctx *models.Context, code string) error {
	id, err := primitive.ObjectIDFromHex(code)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.COMMUNICATIONCREDITLOGSTATUSACTIVE, "activeStatus": true}}
	_, err = ctx.DB.Collection(constants.COLLECTIONCOMMUNICATIONCREDITLOG).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableCommunicationCreditLog :""
func (d *Daos) DisableCommunicationCreditLog(ctx *models.Context, code string) error {
	id, err := primitive.ObjectIDFromHex(code)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.COMMUNICATIONCREDITLOGSTATUSDISABLED, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONCOMMUNICATIONCREDITLOG).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteCommunicationCreditLog :""
func (d *Daos) DeleteCommunicationCreditLog(ctx *models.Context, code string) error {
	id, err := primitive.ObjectIDFromHex(code)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.COMMUNICATIONCREDITLOGSTATUSDELETED, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONCOMMUNICATIONCREDITLOG).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
func (d *Daos) GetSingleCommunicationCreditLogWithMode(ctx *models.Context, code string) (*models.RefCommunicationCreditLog, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"communicationMode": code}})
	mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"_id": -1}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONCOMMUNICATIONCREDITLOG).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var CommunicationCreditLogs []models.RefCommunicationCreditLog
	var CommunicationCreditLog *models.RefCommunicationCreditLog
	if err = cursor.All(ctx.CTX, &CommunicationCreditLogs); err != nil {
		return nil, err
	}
	if len(CommunicationCreditLogs) > 0 {
		CommunicationCreditLog = &CommunicationCreditLogs[0]
	}
	return CommunicationCreditLog, nil
}
