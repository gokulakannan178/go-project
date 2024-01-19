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

// SaveUserLocationLog : ""
func (d *Daos) SaveUserLocationLog(ctx *models.Context, locationLog *models.UserLocationLog) error {
	d.Shared.BsonToJSONPrint(locationLog)
	_, err := ctx.DB.Collection(constants.COLLECTIONUSERLOCATIONLOG).InsertOne(ctx.CTX, locationLog)
	return err
}

// GetSingleUserLocationLog : ""
func (d *Daos) GetSingleUserLocationLog(ctx *models.Context, UniqueID string) (*models.RefUserLocationLog, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONUSERLOCATIONLOG).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var towers []models.RefUserLocationLog
	var tower *models.RefUserLocationLog
	if err = cursor.All(ctx.CTX, &towers); err != nil {
		return nil, err
	}
	if len(towers) > 0 {
		tower = &towers[0]
	}
	return tower, nil
}

// UpdateUserLocationLog : ""
func (d *Daos) UpdateUserLocationLog(ctx *models.Context, locationLog *models.UserLocationLog) error {
	selector := bson.M{"uniqueId": locationLog.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM

	data := bson.M{"$set": locationLog}
	_, err := ctx.DB.Collection(constants.COLLECTIONUSERLOCATIONLOG).UpdateOne(ctx.CTX, selector, data)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// EnableUserLocationLog : ""
func (d *Daos) EnableUserLocationLog(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.USERLOCATIONLOGSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONUSERLOCATIONLOG).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DisableUserLocationLog : ""
func (d *Daos) DisableUserLocationLog(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.USERLOCATIONLOGSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONUSERLOCATIONLOG).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// DeleteUserLocationLog : ""
func (d *Daos) DeleteUserLocationLog(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.USERLOCATIONLOGSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONUSERLOCATIONLOG).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterUserLocationLog : ""
func (d *Daos) FilterUserLocationLog(ctx *models.Context, filter *models.UserLocationLogFilter, pagination *models.Pagination) ([]models.RefUserLocationLog, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
		if len(filter.UniqueID) > 0 {
			query = append(query, bson.M{"uniqueId": bson.M{"$in": filter.UniqueID}})
		}
		if len(filter.UserName) > 0 {
			query = append(query, bson.M{"userName": bson.M{"$in": filter.UserName}})
		}
		if len(filter.UserType) > 0 {
			query = append(query, bson.M{"userType": bson.M{"$in": filter.UserType}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONUSERLOCATIONLOG).CountDocuments(ctx.CTX, func() bson.M {
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
	cursor, err := ctx.DB.Collection(constants.COLLECTIONUSERLOCATIONLOG).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var locationLog []models.RefUserLocationLog
	if err = cursor.All(context.TODO(), &locationLog); err != nil {
		return nil, err
	}
	return locationLog, nil
}
