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

//SaveNotificationLog :""
func (d *Daos) SaveNotificationLog(ctx *models.Context, notificationlog *models.NotificationLog) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONNOTIFICATIONLOG).InsertOne(ctx.CTX, notificationlog)
	if err != nil {
		return err
	}
	notificationlog.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//GetSingleNotificationLog : ""
func (d *Daos) GetSingleNotificationLog(ctx *models.Context, UniqueID string) (*models.RefNotificationLog, error) {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONNOTIFICATIONLOG).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var notificationlogs []models.RefNotificationLog
	var notificationlog *models.RefNotificationLog
	if err = cursor.All(ctx.CTX, &notificationlogs); err != nil {
		return nil, err
	}
	if len(notificationlogs) > 0 {
		notificationlog = &notificationlogs[0]
	}
	return notificationlog, nil
}

//UpdateNotificationLog : ""
func (d *Daos) UpdateNotificationLog(ctx *models.Context, notificationlog *models.NotificationLog) error {

	selector := bson.M{"_id": notificationlog.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": notificationlog}
	_, err := ctx.DB.Collection(constants.COLLECTIONNOTIFICATIONLOG).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//Isread
func (d *Daos) UpdateIsreadNotificationLog(ctx *models.Context, notificationlog *models.NotificationLog) error {

	selector := bson.M{"_id": notificationlog.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": bson.M{"isRead": true}}
	_, err := ctx.DB.Collection(constants.COLLECTIONNOTIFICATIONLOG).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterNotificationLog : ""
func (d *Daos) FilterNotificationLog(ctx *models.Context, notificationlogfilter *models.NotificationLogFilter, pagination *models.Pagination) ([]models.RefNotificationLog, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if notificationlogfilter != nil {

		if len(notificationlogfilter.IsJob) > 0 {
			query = append(query, bson.M{"isJob": bson.M{"$in": notificationlogfilter.IsJob}})
		}
		if len(notificationlogfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": notificationlogfilter.Status}})
		}
		if len(notificationlogfilter.No) > 0 {
			query = append(query, bson.M{"to.no": bson.M{"$in": notificationlogfilter.No}})
		}
		if len(notificationlogfilter.UserName) > 0 {
			query = append(query, bson.M{"to.userName": bson.M{"$in": notificationlogfilter.UserName}})
		}
		if len(notificationlogfilter.UserType) > 0 {
			query = append(query, bson.M{"to.userType": bson.M{"$in": notificationlogfilter.UserType}})
		}
		if len(notificationlogfilter.Name) > 0 {
			query = append(query, bson.M{"to.name": bson.M{"$in": notificationlogfilter.Name}})
		}
		if len(notificationlogfilter.AppRegistrationToken) > 0 {
			query = append(query, bson.M{"to.appRegistrationToken": bson.M{"$in": notificationlogfilter.AppRegistrationToken}})
		}
		//Regex
		if notificationlogfilter.Regex.SentFor != "" {
			query = append(query, bson.M{"sentFor": primitive.Regex{Pattern: notificationlogfilter.Regex.SentFor, Options: "xi"}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"_id": -1}})

	if notificationlogfilter != nil {
		if notificationlogfilter.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{notificationlogfilter.SortBy: notificationlogfilter.SortOrder}})

		}

	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONNOTIFICATIONLOG).CountDocuments(ctx.CTX, func() bson.M {
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
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONEMPLOYEE, "to.userName", "userName", "ref.userId", "ref.userId")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("NotificationLog query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONNOTIFICATIONLOG).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var notificationlogs []models.RefNotificationLog
	if err = cursor.All(context.TODO(), &notificationlogs); err != nil {
		return nil, err
	}
	return notificationlogs, nil
}

//EnableNotificationLog :""
func (d *Daos) EnableNotificationLog(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.NOTIFICATIONLOGSTATUSACTIVE}}
	_, err = ctx.DB.Collection(constants.COLLECTIONNOTIFICATIONLOG).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableNotificationLog :""
func (d *Daos) DisableNotificationLog(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.NOTIFICATIONLOGSTATUSDISABLE}}
	_, err = ctx.DB.Collection(constants.COLLECTIONNOTIFICATIONLOG).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteNotificationLog :""
func (d *Daos) DeleteNotificationLog(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.NOTIFICATIONLOGSTATUSDELETED}}
	_, err = ctx.DB.Collection(constants.COLLECTIONNOTIFICATIONLOG).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
func (d *Daos) GetSingleNotification(ctx *models.Context, uniqueID string) (*models.RefNotificationLog, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONNOTIFICATIONLOG).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var notifications []models.RefNotificationLog
	var notification *models.RefNotificationLog
	if err = cursor.All(ctx.CTX, &notifications); err != nil {
		return nil, err
	}
	if len(notifications) > 0 {
		notification = &notifications[0]
	}
	return notification, nil
}
