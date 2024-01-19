package daos

import (
	"context"
	"fmt"
	"hrms-services/constants"
	"hrms-services/models"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

// SaveTaskMessage : ""
func (d *Daos) SaveTaskMessage(ctx *models.Context, taskMessage *models.TaskMessage) error {
	d.Shared.BsonToJSONPrint(taskMessage)
	_, err := ctx.DB.Collection(constants.COLLECTIONTASKMESSAGE).InsertOne(ctx.CTX, taskMessage)
	return err
}

// FilterTaskMessage : ""
func (d *Daos) FilterTaskMessage(ctx *models.Context, ftm *models.FilterTaskMessage, pagination *models.Pagination) ([]models.RefTaskMessage, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if ftm != nil {
		if len(ftm.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": ftm.Status}})
		}
		if len(ftm.ProjectID) > 0 {
			query = append(query, bson.M{"projectId": bson.M{"$in": ftm.ProjectID}})
		}
		if len(ftm.TaskID) > 0 {
			query = append(query, bson.M{"projectId": bson.M{"$in": ftm.TaskID}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"on": -1}})

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONTASKMESSAGE).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("Feature query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONTASKMESSAGE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var taskmessageFilter []models.RefTaskMessage
	if err = cursor.All(context.TODO(), &taskmessageFilter); err != nil {
		return nil, err
	}
	return taskmessageFilter, nil
}
