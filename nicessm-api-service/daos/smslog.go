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

//SaveSmsLog :""
func (d *Daos) SaveSmsLog(ctx *models.Context, smslog *models.SmsLog) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONSMSLOG).InsertOne(ctx.CTX, smslog)
	if err != nil {
		return err
	}
	smslog.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//GetSingleSmsLog : ""
func (d *Daos) GetSingleSmsLog(ctx *models.Context, UniqueID string) (*models.RefSmsLog, error) {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSMSLOG).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var smslogs []models.RefSmsLog
	var smslog *models.RefSmsLog
	if err = cursor.All(ctx.CTX, &smslogs); err != nil {
		return nil, err
	}
	if len(smslogs) > 0 {
		smslog = &smslogs[0]
	}
	return smslog, nil
}

//UpdateSmsLog : ""
func (d *Daos) UpdateSmsLog(ctx *models.Context, smslog *models.SmsLog) error {

	selector := bson.M{"_id": smslog.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": smslog}
	_, err := ctx.DB.Collection(constants.COLLECTIONSMSLOG).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterSmsLog : ""
func (d *Daos) FilterSmsLog(ctx *models.Context, smslogfilter *models.SmsLogFilter, pagination *models.Pagination) ([]models.RefSmsLog, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if smslogfilter != nil {

		if len(smslogfilter.IsJob) > 0 {
			query = append(query, bson.M{"isJob": bson.M{"$in": smslogfilter.IsJob}})
		}
		if len(smslogfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": smslogfilter.Status}})
		}
		if len(smslogfilter.No) > 0 {
			query = append(query, bson.M{"to.no": bson.M{"$in": smslogfilter.No}})
		}
		if len(smslogfilter.UserName) > 0 {
			query = append(query, bson.M{"to.userName": bson.M{"$in": smslogfilter.UserName}})
		}
		if len(smslogfilter.UserType) > 0 {
			query = append(query, bson.M{"to.userType": bson.M{"$in": smslogfilter.UserType}})
		}
		if len(smslogfilter.Name) > 0 {
			query = append(query, bson.M{"to.name": bson.M{"$in": smslogfilter.Name}})
		}
		//Regex
		if smslogfilter.Regex.SentFor != "" {
			query = append(query, bson.M{"sentFor": primitive.Regex{Pattern: smslogfilter.Regex.SentFor, Options: "xi"}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if smslogfilter != nil {
		if smslogfilter.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{smslogfilter.SortBy: smslogfilter.SortOrder}})

		}

	}
	mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"_id": -1}})

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONSMSLOG).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("SmsLog query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSMSLOG).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var smslogs []models.RefSmsLog
	if err = cursor.All(context.TODO(), &smslogs); err != nil {
		return nil, err
	}
	return smslogs, nil
}

//EnableSmsLog :""
func (d *Daos) EnableSmsLog(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.SMSLOGSTATUSACTIVE}}
	_, err = ctx.DB.Collection(constants.COLLECTIONSMSLOG).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableSmsLog :""
func (d *Daos) DisableSmsLog(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.SMSLOGSTATUSDISABLE}}
	_, err = ctx.DB.Collection(constants.COLLECTIONSMSLOG).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteSmsLog :""
func (d *Daos) DeleteSmsLog(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.SMSLOGSTATUSDELETED}}
	_, err = ctx.DB.Collection(constants.COLLECTIONSMSLOG).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
