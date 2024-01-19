package daos

import (
	"context"
	"errors"
	"fmt"
	"log"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// SaveCronLog : ""
func (d *Daos) SaveCronLog(ctx *models.Context, cronLog *models.CronLog) error {
	d.Shared.BsonToJSONPrint(cronLog)
	_, err := ctx.DB.Collection(constants.COLLECTIONCRONLOG).InsertOne(ctx.CTX, cronLog)
	return err
}

// GetSingleCronLog : ""
func (d *Daos) GetSingleCronLog(ctx *models.Context, UniqueID string) (*models.RefCronLog, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONCRONLOG).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var towers []models.RefCronLog
	var tower *models.RefCronLog
	if err = cursor.All(ctx.CTX, &towers); err != nil {
		return nil, err
	}
	if len(towers) > 0 {
		tower = &towers[0]
	}
	return tower, nil
}

// UpdateCronLog : ""
func (d *Daos) UpdateCronLog(ctx *models.Context, cronLog *models.CronLog) error {
	selector := bson.M{"uniqueId": cronLog.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM

	data := bson.M{"$set": cronLog}
	_, err := ctx.DB.Collection(constants.COLLECTIONCRONLOG).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// EnableCronLog : ""
func (d *Daos) EnableCronLog(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.CRONLOGSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONCRONLOG).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DisableCronLog : ""
func (d *Daos) DisableCronLog(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.CRONLOGSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONCRONLOG).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DeleteCronLog : ""
func (d *Daos) DeleteCronLog(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.CRONLOGSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONCRONLOG).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterCronLog : ""
func (d *Daos) FilterCronLog(ctx *models.Context, filter *models.CronLogFilter, pagination *models.Pagination) ([]models.RefCronLog, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
		if len(filter.UniqueID) > 0 {
			query = append(query, bson.M{"uniqueId": bson.M{"$in": filter.UniqueID}})
		}
		if len(filter.Name) > 0 {
			query = append(query, bson.M{"name": bson.M{"$in": filter.Name}})
		}
		if filter.SearchBox.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: filter.SearchBox.Name, Options: "xi"}})
		}
		if filter.SearchBox.DateStr != "" {
			query = append(query, bson.M{"dateStr": primitive.Regex{Pattern: filter.SearchBox.DateStr, Options: "xi"}})
		}
		if filter.SearchBox.UniqueID != "" {
			query = append(query, bson.M{"uniqueId": primitive.Regex{Pattern: filter.SearchBox.UniqueID, Options: "xi"}})
		}
		if filter.StartTime != nil {
			//var sd,ed time.Time
			if filter.StartTime.From != nil {
				sd := time.Date(filter.StartTime.From.Year(), filter.StartTime.From.Month(), filter.StartTime.From.Day(), 0, 0, 0, 0, filter.StartTime.From.Location())
				ed := time.Date(filter.StartTime.From.Year(), filter.StartTime.From.Month(), filter.StartTime.From.Day(), 23, 59, 59, 0, filter.StartTime.From.Location())
				if filter.StartTime.To != nil {
					ed = time.Date(filter.StartTime.To.Year(), filter.StartTime.To.Month(), filter.StartTime.To.Day(), 23, 59, 59, 0, filter.StartTime.To.Location())
				}
				query = append(query, bson.M{"startTime": bson.M{"$gte": sd, "$lte": ed}})

			}
		}
		if filter.EndTime != nil {
			//var sd,ed time.Time
			if filter.EndTime.From != nil {
				sd := time.Date(filter.EndTime.From.Year(), filter.EndTime.From.Month(), filter.EndTime.From.Day(), 0, 0, 0, 0, filter.EndTime.From.Location())
				ed := time.Date(filter.EndTime.From.Year(), filter.EndTime.From.Month(), filter.EndTime.From.Day(), 23, 59, 59, 0, filter.EndTime.From.Location())
				if filter.EndTime.To != nil {
					ed = time.Date(filter.EndTime.To.Year(), filter.EndTime.To.Month(), filter.EndTime.To.Day(), 23, 59, 59, 0, filter.EndTime.To.Location())
				}
				query = append(query, bson.M{"endTime": bson.M{"$gte": sd, "$lte": ed}})

			}
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONCRONLOG).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONCRONLOG).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var cronLog []models.RefCronLog
	if err = cursor.All(context.TODO(), &cronLog); err != nil {
		return nil, err
	}
	return cronLog, nil
}

// InitCronLog : ""
func (d *Daos) InitCronLog(ctx *models.Context, cronLog *models.CronLog) error {
	d.Shared.BsonToJSONPrint(cronLog)
	_, err := ctx.DB.Collection(constants.COLLECTIONCRONLOG).InsertOne(ctx.CTX, cronLog)
	return err
}

// EndCronLog : ""
func (d *Daos) EndCronLog(ctx *models.Context, cronLog *models.CronLog) error {
	selector := bson.M{"name": cronLog.Name, "isCurrentScript": cronLog.IsCurrentScript}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	data := bson.M{"$set": bson.M{"endTime": cronLog.EndTime,
		"status":       cronLog.Status,
		"errorMessage": cronLog.ErrorMessage,
	}}
	_, err := ctx.DB.Collection(constants.COLLECTIONCRONLOG).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// RemoveCronLog : ""
func (d *Daos) CloseOldCron(ctx *models.Context, name string) error {
	selector := bson.M{"name": name}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	data := bson.M{"$set": bson.M{"isCurrentScript": false}}
	_, err := ctx.DB.Collection(constants.COLLECTIONCRONLOG).UpdateMany(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}
