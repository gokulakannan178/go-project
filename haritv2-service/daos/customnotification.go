package daos

import (
	"context"
	"errors"
	"fmt"
	"haritv2-service/constants"
	"haritv2-service/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

//SaveCustomNotification :""
func (d *Daos) SaveCustomNotification(ctx *models.Context, notification *models.CustomNotification) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONCUSTOMNOTIFICATION).InsertOne(ctx.CTX, notification)
	return err
}

//GetSingleCustomNotification : ""
func (d *Daos) GetSingleCustomNotification(ctx *models.Context, uniqueID string) (*models.RefCustomNotification, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONCUSTOMNOTIFICATION).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var notifications []models.RefCustomNotification
	var notification *models.RefCustomNotification
	if err = cursor.All(ctx.CTX, &notifications); err != nil {
		return nil, err
	}
	if len(notifications) > 0 {
		notification = &notifications[0]
	}
	return notification, nil
}

//CustomNotification : ""
func (d *Daos) UpdateCustomNotification(ctx *models.Context, notification *models.CustomNotification) error {
	selector := bson.M{"uniqueId": notification.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": notification}
	_, err := ctx.DB.Collection(constants.COLLECTIONCUSTOMNOTIFICATION).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//EnableCustomNotification :""
func (d *Daos) EnableCustomNotification(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.CUSTOMNOTIFICATIONSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONCUSTOMNOTIFICATION).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableCustomNotification :""
func (d *Daos) DisableCustomNotification(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.CUSTOMNOTIFICATIONSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONCUSTOMNOTIFICATION).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteCustomNotification :""
func (d *Daos) DeleteCustomNotification(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.CUSTOMNOTIFICATIONSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONCUSTOMNOTIFICATION).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//FilterCustomNotification : ""
func (d *Daos) FilterCustomNotification(ctx *models.Context, filter *models.CustomNotificationFilter, pagination *models.Pagination) ([]models.RefCustomNotification, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.UniqueID) > 0 {
			query = append(query, bson.M{"uniqueId": bson.M{"$in": filter.UniqueID}})
		}
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONCUSTOMNOTIFICATION).CountDocuments(ctx.CTX, func() bson.M {
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
	cursor, err := ctx.DB.Collection(constants.COLLECTIONCUSTOMNOTIFICATION).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var notifications []models.RefCustomNotification
	if err = cursor.All(context.TODO(), &notifications); err != nil {
		return nil, err
	}
	return notifications, nil
}
