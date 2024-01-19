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

//SaveDisease :""
func (d *Daos) SaveContentCountLog(ctx *models.Context, ContentCountLog *models.ContentCountLog) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONCONTENTCOUNTLOG).InsertOne(ctx.CTX, ContentCountLog)
	if err != nil {
		return err
	}
	ContentCountLog.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//GetSingleDisease : ""
func (d *Daos) GetSingleContentCountLog(ctx *models.Context, UniqueID string) (*models.RefContentCountLog, error) {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONCONTENTCOUNTLOG).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var ContentCountLogs []models.RefContentCountLog
	var ContentCountLog *models.RefContentCountLog
	if err = cursor.All(ctx.CTX, &ContentCountLogs); err != nil {
		return nil, err
	}
	if len(ContentCountLogs) > 0 {
		ContentCountLog = &ContentCountLogs[0]
	}
	return ContentCountLog, nil
}

//UpdateContentCountLog : ""
func (d *Daos) UpdateContentCountLog(ctx *models.Context, ContentCountLog *models.ContentCountLog) error {

	selector := bson.M{"_id": ContentCountLog.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": ContentCountLog}
	_, err := ctx.DB.Collection(constants.COLLECTIONCONTENTCOUNTLOG).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterContentCountLog : ""
func (d *Daos) FilterContentCountLog(ctx *models.Context, ContentCountLogfilter *models.ContentCountLogFilter, pagination *models.Pagination) ([]models.RefContentCountLog, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if ContentCountLogfilter != nil {

		if len(ContentCountLogfilter.ActiveStatus) > 0 {
			query = append(query, bson.M{"activeStatus": bson.M{"$in": ContentCountLogfilter.ActiveStatus}})
		}
		if len(ContentCountLogfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": ContentCountLogfilter.Status}})
		}
		//Regex
		if ContentCountLogfilter.SearchBox.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: ContentCountLogfilter.SearchBox.Name, Options: "xi"}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if ContentCountLogfilter != nil {
		if ContentCountLogfilter.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{ContentCountLogfilter.SortBy: ContentCountLogfilter.SortOrder}})

		}

	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONCONTENTCOUNTLOG).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("ContentCountLog query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONCONTENTCOUNTLOG).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var ContentCountLogs []models.RefContentCountLog
	if err = cursor.All(context.TODO(), &ContentCountLogs); err != nil {
		return nil, err
	}
	return ContentCountLogs, nil
}

//EnableContentCountLog :""
func (d *Daos) EnableContentCountLog(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.CONTENTCOUNTLOGSTATUSACTIVE, "activeStatus": true}}
	_, err = ctx.DB.Collection(constants.COLLECTIONCONTENTCOUNTLOG).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableDisease :""
func (d *Daos) DisableContentCountLog(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.CONTENTCOUNTLOGSTATUSDISABLE, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONCONTENTCOUNTLOG).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteContentCountLog :""
func (d *Daos) DeleteContentCountLog(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.CONTENTCOUNTLOGSTATUSDELETED, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONCONTENTCOUNTLOG).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
