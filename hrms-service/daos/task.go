package daos

import (
	"context"
	"errors"
	"fmt"
	"hrms-services/constants"
	"hrms-services/models"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

// SaveTask : ""
func (d *Daos) SaveTask(ctx *models.Context, tk *models.Task) error {
	d.Shared.BsonToJSONPrint(tk)
	_, err := ctx.DB.Collection(constants.COLLECTIONTASK).InsertOne(ctx.CTX, tk)
	return err
}

//SaveTaskMembers :""
func (d *Daos) SaveTaskMembers(ctx *models.Context, tm []models.TaskMember) error {
	var data []interface{}

	for _, v := range tm {
		data = append(data, v)
	}
	_, err := ctx.DB.Collection(constants.COLLECTIONTASKMEMBER).InsertMany(ctx.CTX, data)
	return err
}

// SaveTaskTeamMember : ""
func (d *Daos) SaveTaskTeamMember(ctx *models.Context, tm *models.TaskMember) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONTASKMEMBER).InsertOne(ctx.CTX, tm)
	return err
}

// DisableTaskTeamMember : ""
func (d *Daos) DisableTaskTeamMember(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.TASKTEAMMEMBERSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONTASKMEMBER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// GetSingleTask : ""
func (d *Daos) GetSingleTask(ctx *models.Context, uniqueID string) (*models.RefTask, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONTASK).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var tks []models.RefTask
	var tk *models.RefTask
	if err = cursor.All(ctx.CTX, &tks); err != nil {
		return nil, err
	}
	if len(tks) > 0 {
		tk = &tks[0]
	}
	return tk, err
}

// EnableTask : ""
func (d *Daos) EnableTask(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.TASKSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONTASK).UpdateOne(ctx.CTX, selector, data)
	return err
}

// DisableTask : ""
func (d *Daos) DisableTask(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.TASKSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONTASK).UpdateOne(ctx.CTX, selector, data)
	return err
}

// DeleteTask : ""
func (d *Daos) DeleteTask(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.TASKSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONTASK).UpdateOne(ctx.CTX, selector, data)
	return err
}

// UpdateTask : ""
func (d *Daos) UpdateTask(ctx *models.Context, p *models.Task) error {
	selector := bson.M{"uniqueId": p.UniqueID}
	data := bson.M{"$set": p}
	_, err := ctx.DB.Collection(constants.COLLECTIONTASK).UpdateOne(ctx.CTX, selector, data)
	return err
}

// FilterTask : ""
func (d *Daos) FilterTask(ctx *models.Context, ft *models.FilterTask, pagination *models.Pagination) ([]models.RefTask, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if ft != nil {
		if len(ft.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": ft.Status}})
		}
		if len(ft.ProjectId) > 0 {
			query = append(query, bson.M{"projectId": bson.M{"$in": ft.ProjectId}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	//{"$sort":{"on":-1}}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONTASK).CountDocuments(ctx.CTX, func() bson.M {
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
	cursor, err := ctx.DB.Collection(constants.COLLECTIONTASK).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var taskFilter []models.RefTask
	if err = cursor.All(context.TODO(), &taskFilter); err != nil {
		return nil, err
	}
	return taskFilter, nil
}
