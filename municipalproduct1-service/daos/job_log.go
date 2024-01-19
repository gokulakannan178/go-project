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
)

// SaveJobLog : ""
func (d *Daos) SaveJobLog(ctx *models.Context, jobLog *models.JobLog) error {
	d.Shared.BsonToJSONPrint(jobLog)
	_, err := ctx.DB.Collection(constants.COLLECTIONJOBLOG).InsertOne(ctx.CTX, jobLog)
	return err
}

// GetSingleJobLog : ""
func (d *Daos) GetSingleJobLog(ctx *models.Context, UniqueID string) (*models.RefJobLog, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})

	d.Shared.BsonToJSONPrintTag("getsinglejobLog query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONJOBLOG).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var towers []models.RefJobLog
	var tower *models.RefJobLog
	if err = cursor.All(ctx.CTX, &towers); err != nil {
		return nil, err
	}
	if len(towers) > 0 {
		tower = &towers[0]
	}
	return tower, nil
}

// UpdateJobLog : ""
func (d *Daos) UpdateJobLog(ctx *models.Context, joblog *models.JobLog) error {
	selector := bson.M{"uniqueId": joblog.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM

	data := bson.M{"$set": joblog}
	_, err := ctx.DB.Collection(constants.COLLECTIONJOBLOG).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// EnableJobLog : ""
func (d *Daos) EnableJobLog(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.JOBLOGSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONJOBLOG).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DisableJobLog : ""
func (d *Daos) DisableJobLog(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.JOBLOGSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONJOBLOG).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DeleteJobLog : ""
func (d *Daos) DeleteJobLog(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.JOBLOGSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONJOBLOG).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterJobLog : ""
func (d *Daos) FilterJobLog(ctx *models.Context, filter *models.JobLogFilter, pagination *models.Pagination) ([]models.RefJobLog, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
		if len(filter.UniqueID) > 0 {
			query = append(query, bson.M{"uniqueId": bson.M{"$in": filter.UniqueID}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONJOBLOG).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("filterjobLog query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONJOBLOG).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var jobLog []models.RefJobLog
	if err = cursor.All(context.TODO(), &jobLog); err != nil {
		return nil, err
	}
	return jobLog, nil
}
