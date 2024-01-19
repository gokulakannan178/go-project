package daos

import (
	"context"
	"errors"
	"fmt"
	"lgfweather-service/constants"
	"lgfweather-service/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//SaveNotification :""
func (d *Daos) SaveNotification(ctx *models.Context, notification *models.Notification) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONNOTIFICATION).InsertOne(ctx.CTX, notification)
	return err
}

//GetSingleNotification : ""
func (d *Daos) GetSingleNotification(ctx *models.Context, uniqueID string) (*models.RefNotification, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONNOTIFICATION).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var notifications []models.RefNotification
	var notification *models.RefNotification
	if err = cursor.All(ctx.CTX, &notifications); err != nil {
		return nil, err
	}
	if len(notifications) > 0 {
		notification = &notifications[0]
	}
	return notification, nil
}

//UpdateNotification : ""
func (d *Daos) UpdateNotification(ctx *models.Context, notification *models.Notification) error {
	selector := bson.M{"uniqueId": notification.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": notification}
	_, err := ctx.DB.Collection(constants.COLLECTIONNOTIFICATION).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//EnableNotification :""
func (d *Daos) EnableNotification(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.NOTIFICATIONSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONNOTIFICATION).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableNotification :""
func (d *Daos) DisableNotification(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.NOTIFICATIONSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONNOTIFICATION).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteNotification :""
func (d *Daos) DeleteNotification(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.NOTIFICATIONSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONNOTIFICATION).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//FilterNotification : ""
func (d *Daos) FilterNotification(ctx *models.Context, filter *models.NotificationFilter, pagination *models.Pagination) ([]models.RefNotification, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.UniqueID) > 0 {
			query = append(query, bson.M{"uniqueId": bson.M{"$in": filter.UniqueID}})
		}
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
		if filter.SearchText.Content != "" {
			query = append(query, bson.M{"content": primitive.Regex{Pattern: filter.SearchText.Content, Options: "xi"}})
		}
		if filter.SearchText.UniqueID != "" {
			query = append(query, bson.M{"uniqueId": primitive.Regex{Pattern: filter.SearchText.UniqueID, Options: "xi"}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONNOTIFICATION).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("notification query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONNOTIFICATION).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var notifications []models.RefNotification
	if err = cursor.All(context.TODO(), &notifications); err != nil {
		return nil, err
	}
	return notifications, nil
}
