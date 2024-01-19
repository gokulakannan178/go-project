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

// SaveEmployeeLog : ""
func (d *Daos) SaveEmployeeLog(ctx *models.Context, employeelog *models.EmployeeLog) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEELOG).InsertOne(ctx.CTX, employeelog)
	if err != nil {
		return err
	}
	employeelog.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//UpdateEmployeeLog : ""
func (d *Daos) UpdateEmployeeLog(ctx *models.Context, employeelog *models.EmployeeLog) error {
	selector := bson.M{"uniqueId": employeelog.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": employeelog}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEELOG).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// GetSingleEmployeeLog : ""
func (d *Daos) GetSingleEmployeeLog(ctx *models.Context, uniqueID string) (*models.RefEmployeeLog, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	//LookUp
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEELOG).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var employeelogs []models.RefEmployeeLog
	var employeelog *models.RefEmployeeLog
	if err = cursor.All(ctx.CTX, &employeelogs); err != nil {
		return nil, err
	}
	if len(employeelogs) > 0 {
		employeelog = &employeelogs[0]
	}
	return employeelog, err
}

// EnableEmployeeLog : ""
func (d *Daos) EnableEmployeeLog(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.EMPLOYEELOGSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEELOG).UpdateOne(ctx.CTX, selector, data)
	return err
}

// DisableEmployeeLog : ""
func (d *Daos) DisableEmployeeLog(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.EMPLOYEELOGSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEELOG).UpdateOne(ctx.CTX, selector, data)
	return err
}

//DeleteEmployeeLog :""
func (d *Daos) DeleteEmployeeLog(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.EMPLOYEELOGSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEELOG).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterEmployeeLog : ""
func (d *Daos) FilterEmployeeLog(ctx *models.Context, employeelog *models.FilterEmployeeLog, pagination *models.Pagination) ([]models.RefEmployeeLog, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if employeelog != nil {
		if len(employeelog.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": employeelog.Status}})
		}
		if len(employeelog.OrganisationId) > 0 {
			query = append(query, bson.M{"organisationId": bson.M{"$in": employeelog.OrganisationId}})
		}
		//Regex
		if employeelog.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: employeelog.Regex.Name, Options: "xi"}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEELOG).CountDocuments(ctx.CTX, func() bson.M {
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
	//LookUp
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("Feature query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEELOG).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var rf []models.RefEmployeeLog
	if err = cursor.All(context.TODO(), &rf); err != nil {
		return nil, err
	}
	return rf, nil
}
