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

//SaveJobTimeline :""
func (d *Daos) SaveJobTimeline(ctx *models.Context, jobTimeline *models.JobTimeline) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONJOBTIMELINE).InsertOne(ctx.CTX, jobTimeline)
	if err != nil {
		return err
	}
	jobTimeline.ID = res.InsertedID.(primitive.ObjectID)
	return nil

}

//GetSingleJobTimeline : ""
func (d *Daos) GetSingleJobTimeline(ctx *models.Context, UniqueID string) (*models.RefJobTimeline, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONJOBTIMELINE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var jobTimelines []models.RefJobTimeline
	var jobTimeline *models.RefJobTimeline
	if err = cursor.All(ctx.CTX, &jobTimelines); err != nil {
		return nil, err
	}
	if len(jobTimelines) > 0 {
		jobTimeline = &jobTimelines[0]
	}
	return jobTimeline, nil
}

//GetSingleJobTimelineemployeeid : ""
func (d *Daos) GetSingleJobTimelineemployeeid(ctx *models.Context, employeeId string, status string) (*models.RefJobTimeline, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"employeeId": employeeId, "status": status}})
	//mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"status": status}})

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONJOBTIMELINE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var jobTimelines []models.RefJobTimeline
	var jobTimeline *models.RefJobTimeline
	if err = cursor.All(ctx.CTX, &jobTimelines); err != nil {
		return nil, err
	}
	if len(jobTimelines) > 0 {
		jobTimeline = &jobTimelines[0]
	}
	return jobTimeline, nil
}

//UpdateJobTimeline : ""
func (d *Daos) UpdateJobTimeline(ctx *models.Context, jobTimeline *models.JobTimeline) error {
	selector := bson.M{"uniqueId": jobTimeline.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": jobTimeline}
	_, err := ctx.DB.Collection(constants.COLLECTIONJOBTIMELINE).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterJobTimeline : ""
func (d *Daos) FilterJobTimeline(ctx *models.Context, filter *models.JobTimelineFilter, pagination *models.Pagination) ([]models.RefJobTimeline, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {

		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
		//Regex
		if filter.Regex.LineManager != "" {
			query = append(query, bson.M{"lineManager": primitive.Regex{Pattern: filter.Regex.LineManager, Options: "xi"}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONJOBTIMELINE).CountDocuments(ctx.CTX, func() bson.M {
			if query != nil {
				if len(query) > 0 {
					return bson.M{"$and": query}
				}
			}
			return bson.M{}
		}())
		if err != nil {
			log.Println("Error in getting pagination")
		}
		fmt.Println("count", totalCount)
		pagination.Count = int(totalCount)
		d.Shared.PaginationData(pagination)
	}

	//Aggregation
	d.Shared.BsonToJSONPrintTag("JobTimeline query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONJOBTIMELINE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var jobTimelines []models.RefJobTimeline
	if err = cursor.All(context.TODO(), &jobTimelines); err != nil {
		return nil, err
	}
	return jobTimelines, nil
}

//EnableJobTimeline :""
func (d *Daos) EnableJobTimeline(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.JOBTIMELINESTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONJOBTIMELINE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableJobTimeline :""
func (d *Daos) DisableJobTimeline(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.JOBTIMELINESTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONJOBTIMELINE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteJobTimeline :""
func (d *Daos) DeleteJobTimeline(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.JOBTIMELINESTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONJOBTIMELINE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
