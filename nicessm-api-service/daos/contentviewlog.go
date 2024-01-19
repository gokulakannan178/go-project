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
	"go.mongodb.org/mongo-driver/mongo/options"
)

//SaveDisease :""
func (d *Daos) SaveContentViewLog(ctx *models.Context, ContentViewLog *models.ContentViewLog) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONCONTENTVIEWLOG).InsertOne(ctx.CTX, ContentViewLog)
	if err != nil {
		return err
	}
	ContentViewLog.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}
func (d *Daos) SaveContentViewLogUpdertInc(ctx *models.Context, content *models.ContentViewLog) error {
	// t := time.Now()
	opts := options.Update().SetUpsert(true)
	updateQuery := bson.M{"uniqueId": content.UniqueId, "contentId": content.ContentId}
	fmt.Println("updateQuery===>", updateQuery)
	updateData := bson.M{"$inc": bson.M{"count": 1}}
	if _, err := ctx.DB.Collection(constants.COLLECTIONCONTENTVIEWLOG).UpdateOne(ctx.CTX, updateQuery, updateData, opts); err != nil {
		return errors.New("Error in updating log - " + err.Error())
	}

	return nil
}

//GetSingleDisease : ""
func (d *Daos) GetSingleContentViewLog(ctx *models.Context, UniqueID string) (*models.RefContentViewLog, error) {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONCONTENTVIEWLOG).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var ContentViewLogs []models.RefContentViewLog
	var ContentViewLog *models.RefContentViewLog
	if err = cursor.All(ctx.CTX, &ContentViewLogs); err != nil {
		return nil, err
	}
	if len(ContentViewLogs) > 0 {
		ContentViewLog = &ContentViewLogs[0]
	}
	return ContentViewLog, nil
}

//UpdateContentViewLog : ""
func (d *Daos) UpdateContentViewLog(ctx *models.Context, ContentViewLog *models.ContentViewLog) error {

	selector := bson.M{"_id": ContentViewLog.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": ContentViewLog}
	_, err := ctx.DB.Collection(constants.COLLECTIONCONTENTVIEWLOG).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterContentViewLog : ""
func (d *Daos) FilterContentViewLog(ctx *models.Context, ContentViewLogfilter *models.ContentViewLogFilter, pagination *models.Pagination) ([]models.RefContentViewLog, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if ContentViewLogfilter != nil {

		if len(ContentViewLogfilter.ActiveStatus) > 0 {
			query = append(query, bson.M{"activeStatus": bson.M{"$in": ContentViewLogfilter.ActiveStatus}})
		}
		if len(ContentViewLogfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": ContentViewLogfilter.Status}})
		}
		//Regex
		if ContentViewLogfilter.SearchBox.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: ContentViewLogfilter.SearchBox.Name, Options: "xi"}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if ContentViewLogfilter != nil {
		if ContentViewLogfilter.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{ContentViewLogfilter.SortBy: ContentViewLogfilter.SortOrder}})

		}

	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONCONTENTVIEWLOG).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("ContentViewLog query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONCONTENTVIEWLOG).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var ContentViewLogs []models.RefContentViewLog
	if err = cursor.All(context.TODO(), &ContentViewLogs); err != nil {
		return nil, err
	}
	return ContentViewLogs, nil
}

//EnableContentViewLog :""
func (d *Daos) EnableContentViewLog(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.CONTENTVIEWLOGSTATUSACTIVE, "activeStatus": true}}
	_, err = ctx.DB.Collection(constants.COLLECTIONCONTENTVIEWLOG).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableDisease :""
func (d *Daos) DisableContentViewLog(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.CONTENTVIEWLOGSTATUSDISABLE, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONCONTENTVIEWLOG).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteContentViewLog :""
func (d *Daos) DeleteContentViewLog(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.CONTENTVIEWLOGSTATUSDELETED, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONCONTENTVIEWLOG).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
func (d *Daos) DayWiseContentViewChart(ctx *models.Context, contentfilter *models.FilterDaywiseViewChart) (*models.DayWiseContentViewChartReport, error) {
	var sd, ed time.Time
	mainPipeline := []bson.M{}
	query := []bson.M{}
	query = append(query, bson.M{"contentId": contentfilter.ContentId})

	if contentfilter.CreatedFrom.StartDate != nil {
		var sdcondition, edcondition string = "gte", "lte"
		sd = time.Date(contentfilter.CreatedFrom.StartDate.Year(), contentfilter.CreatedFrom.StartDate.Month(), contentfilter.CreatedFrom.StartDate.Day(), 0, 0, 0, 0, contentfilter.CreatedFrom.StartDate.Location())
		ed = time.Date(contentfilter.CreatedFrom.EndDate.Year(), contentfilter.CreatedFrom.EndDate.Month(), contentfilter.CreatedFrom.EndDate.Day(), 23, 59, 59, 0, contentfilter.CreatedFrom.EndDate.Location())

		if contentfilter.CreatedFrom.EndDate != nil {
			ed = time.Date(contentfilter.CreatedFrom.EndDate.Year(), contentfilter.CreatedFrom.EndDate.Month(), contentfilter.CreatedFrom.EndDate.Day(), 23, 59, 59, 0, contentfilter.CreatedFrom.EndDate.Location())
			//edcondition = contentfilter.CreatedTo.Condition
		}
		fmt.Println("sd==>", sd)
		fmt.Println("ed==>", ed)
		query = append(query, bson.M{"date": bson.M{"$" + sdcondition: sd, "$" + edcondition: ed}})
	}
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"dom": bson.M{"$dayOfMonth": "$date"}}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"month": bson.M{"$month": "$date"}}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"year": bson.M{"$year": "$date"}}})
	//	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"dayOfWeek": bson.M{"$dayOfWeek": "$date"}}})
	//fmt.Println("dayofweek==>", attendance)
	mainPipeline = append(mainPipeline, bson.M{"$group": bson.M{"_id": nil, "days": bson.M{"$push": "$$ROOT"}}})
	mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"date": 1}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"month1": bson.M{"$arrayElemAt": []interface{}{"$days.month", 0}}}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"year1": bson.M{"$arrayElemAt": []interface{}{"$days.year", 0}}}})

	//mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"dayOfWeek1": bson.M{"$arrayElemAt": []interface{}{"$days.dayOfWeek", 0}}}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{
		"days": bson.M{
			"$map": bson.M{
				"input": bson.M{"$range": []interface{}{sd.Day(), ed.Day() + 1, 1}},
				"as":    "rangeDay",
				"in": bson.M{
					"$let": bson.M{
						"vars": bson.M{"index": bson.M{"$indexOfArray": []string{"$days.dom", "$$rangeDay"}}},
						"in": bson.M{
							"$cond": bson.M{
								"if": bson.M{"$eq": []interface{}{"$$index", -1}},
								"then": bson.M{
									"date":  bson.M{"$dateFromParts": bson.M{"year": "$year1", "month": "$month1", "day": "$$rangeDay", "hour": 9}},
									"count": 0,
								},
								"else": bson.M{"$arrayElemAt": []string{"$days", "$$index"}}},
						},
					},
				},
			},
		},
	}})
	//Aggregation
	d.Shared.BsonToJSONPrint(mainPipeline)
	var emptyData *models.DayWiseContentViewChartReport

	cursor, err := ctx.DB.Collection(constants.COLLECTIONCONTENTVIEWLOG).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return emptyData, err
	}
	var data []models.DayWiseContentViewChartReport

	if err = cursor.All(context.TODO(), &data); err != nil {
		return emptyData, err
	}
	if len(data) > 0 {
		return &data[0], nil
	}

	return emptyData, nil

}
