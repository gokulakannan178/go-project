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

//SaveTelegramLog :""
func (d *Daos) SaveTelegramLog(ctx *models.Context, telegramlog *models.TelegramLog) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONTELEGRAMLOG).InsertOne(ctx.CTX, telegramlog)
	if err != nil {
		return err
	}
	telegramlog.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//GetSingleTelegramLog : ""
func (d *Daos) GetSingleTelegramLog(ctx *models.Context, UniqueID string) (*models.RefTelegramLog, error) {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONTELEGRAMLOG).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var telegramlogs []models.RefTelegramLog
	var telegramlog *models.RefTelegramLog
	if err = cursor.All(ctx.CTX, &telegramlogs); err != nil {
		return nil, err
	}
	if len(telegramlogs) > 0 {
		telegramlog = &telegramlogs[0]
	}
	return telegramlog, nil
}

//UpdateTelegramLog : ""
func (d *Daos) UpdateTelegramLog(ctx *models.Context, telegramlog *models.TelegramLog) error {

	selector := bson.M{"_id": telegramlog.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": telegramlog}
	_, err := ctx.DB.Collection(constants.COLLECTIONTELEGRAMLOG).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterTelegramLog : ""
func (d *Daos) FilterTelegramLog(ctx *models.Context, telegramlogfilter *models.TelegramLogFilter, pagination *models.Pagination) ([]models.RefTelegramLog, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if telegramlogfilter != nil {

		if len(telegramlogfilter.IsJob) > 0 {
			query = append(query, bson.M{"isJob": bson.M{"$in": telegramlogfilter.IsJob}})
		}
		if len(telegramlogfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": telegramlogfilter.Status}})
		}
		if len(telegramlogfilter.No) > 0 {
			query = append(query, bson.M{"to.no": bson.M{"$in": telegramlogfilter.No}})
		}
		if len(telegramlogfilter.UserName) > 0 {
			query = append(query, bson.M{"to.userName": bson.M{"$in": telegramlogfilter.UserName}})
		}
		if len(telegramlogfilter.UserType) > 0 {
			query = append(query, bson.M{"to.userType": bson.M{"$in": telegramlogfilter.UserType}})
		}
		if len(telegramlogfilter.Name) > 0 {
			query = append(query, bson.M{"to.name": bson.M{"$in": telegramlogfilter.Name}})
		}
		//Regex
		if telegramlogfilter.Regex.SentFor != "" {
			query = append(query, bson.M{"sentFor": primitive.Regex{Pattern: telegramlogfilter.Regex.SentFor, Options: "xi"}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if telegramlogfilter != nil {
		if telegramlogfilter.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{telegramlogfilter.SortBy: telegramlogfilter.SortOrder}})

		}

	}
	mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"_id": -1}})

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONTELEGRAMLOG).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("TelegramLog query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONTELEGRAMLOG).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var telegramlogs []models.RefTelegramLog
	if err = cursor.All(context.TODO(), &telegramlogs); err != nil {
		return nil, err
	}
	return telegramlogs, nil
}

//EnableTelegramLog :""
func (d *Daos) EnableTelegramLog(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.TELEGRAMLOGSTATUSACTIVE}}
	_, err = ctx.DB.Collection(constants.COLLECTIONTELEGRAMLOG).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableTelegramLog :""
func (d *Daos) DisableTelegramLog(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.TELEGRAMLOGSTATUSDISABLE}}
	_, err = ctx.DB.Collection(constants.COLLECTIONTELEGRAMLOG).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteTelegramLog :""
func (d *Daos) DeleteTelegramLog(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.TELEGRAMLOGSTATUSDELETED}}
	_, err = ctx.DB.Collection(constants.COLLECTIONTELEGRAMLOG).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
