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

//SaveTodayTips :""
func (d *Daos) SaveTodayTips(ctx *models.Context, todayTips *models.TodayTips) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONTODAYTIPS).InsertOne(ctx.CTX, todayTips)
	return err
}

//GetSingleTodayTips : ""
func (d *Daos) GetSingleTodayTips(ctx *models.Context, uniqueID string) (*models.RefTodayTips, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONTODAYTIPS).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var todayTipss []models.RefTodayTips
	var todayTips *models.RefTodayTips
	if err = cursor.All(ctx.CTX, &todayTipss); err != nil {
		return nil, err
	}
	if len(todayTipss) > 0 {
		todayTips = &todayTipss[0]
	}
	return todayTips, nil
}

//UpdateTodayTips : ""
func (d *Daos) UpdateTodayTips(ctx *models.Context, todayTips *models.TodayTips) error {
	selector := bson.M{"uniqueId": todayTips.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": todayTips, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONTODAYTIPS).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//EnableTodayTips :""
func (d *Daos) EnableTodayTips(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.TODAYTIPSSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONTODAYTIPS).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableTodayTips :""
func (d *Daos) DisableTodayTips(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.TODAYTIPSSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONTODAYTIPS).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteTodayTips :""
func (d *Daos) DeleteTodayTips(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.TODAYTIPSSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONTODAYTIPS).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//FilterTodayTips : ""
func (d *Daos) FilterTodayTips(ctx *models.Context, todayTipsfilter *models.TodayTipsFilter, pagination *models.Pagination) ([]models.RefTodayTips, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if todayTipsfilter != nil {
		if len(todayTipsfilter.UniqueID) > 0 {
			query = append(query, bson.M{"uniqueId": bson.M{"$in": todayTipsfilter.UniqueID}})
		}
		if len(todayTipsfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": todayTipsfilter.Status}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONTODAYTIPS).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("todayTips query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONTODAYTIPS).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var todayTipss []models.RefTodayTips
	if err = cursor.All(context.TODO(), &todayTipss); err != nil {
		return nil, err
	}
	return todayTipss, nil
}

// // GetTodayTips : ""
func (d *Daos) GetTodayTips(ctx *models.Context) (*models.RefTodayTips, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"status": constants.TODAYTIPSSTATUSACTIVE}})
	mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"_id": -1}})

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONTODAYTIPS).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var tips []models.RefTodayTips
	var tip *models.RefTodayTips
	if err = cursor.All(ctx.CTX, &tips); err != nil {
		return nil, err
	}
	if len(tips) > 0 {
		tip = &tips[0]
	}
	return tip, nil
}
