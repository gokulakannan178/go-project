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

//SaveEmployeeLeaveLog : ""
func (d *Daos) SaveEmployeeLeaveLog(ctx *models.Context, employeeLeaveLog *models.EmployeeLeaveLog) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEELEAVELOG).InsertOne(ctx.CTX, employeeLeaveLog)
	if err != nil {
		return err
	}
	employeeLeaveLog.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

// GetSingleEmployeeLeaveLog : ""
func (d *Daos) GetSingleEmployeeLeaveLog(ctx *models.Context, uniqueID string) (*models.RefEmployeeLeaveLog, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	//LookUp
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEELEAVELOG).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var EmployeeLeaveLogs []models.RefEmployeeLeaveLog
	var EmployeeLeaveLog *models.RefEmployeeLeaveLog
	if err = cursor.All(ctx.CTX, &EmployeeLeaveLogs); err != nil {
		return nil, err
	}
	if len(EmployeeLeaveLogs) > 0 {
		EmployeeLeaveLog = &EmployeeLeaveLogs[0]
	}
	return EmployeeLeaveLog, err
}

//EmployeeLeaveLogCount : ""
func (d *Daos) EmployeeLeaveLogCount(ctx *models.Context, employeeId string, OrganisationId string, LeaveType string) (*models.RefEmployeeLeaveLogCount, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"employeeId": employeeId, "organisationId": OrganisationId, "leaveType": LeaveType}})
	group := []bson.M{{"$group": bson.M{"_id": nil, "totalLeave": bson.M{"$sum": "$value"}}}}
	mainPipeline = append(mainPipeline, group...)

	d.Shared.BsonToJSONPrintTag("EmployeeLeaveLogCount =>", mainPipeline)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEELEAVELOG).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var EmployeeLeaveLogs []models.RefEmployeeLeaveLogCount
	var EmployeeLeaveLog *models.RefEmployeeLeaveLogCount
	if err = cursor.All(ctx.CTX, &EmployeeLeaveLogs); err != nil {
		return nil, err
	}

	if len(EmployeeLeaveLogs) > 0 {
		EmployeeLeaveLog = &EmployeeLeaveLogs[0]
	}
	return EmployeeLeaveLog, nil
}

//UpdateEmployeeLeaveLog : ""
func (d *Daos) UpdateEmployeeLeaveLog(ctx *models.Context, employeeLeaveLog *models.EmployeeLeaveLog) error {
	selector := bson.M{"uniqueId": employeeLeaveLog.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": employeeLeaveLog}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEELEAVELOG).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// EnableEmployeeLeaveLog : ""
func (d *Daos) EnableEmployeeLeaveLog(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.EMPLOYEELEAVELOGSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEELEAVELOG).UpdateOne(ctx.CTX, selector, data)
	return err
}

// DisableEmployeeLeaveLog : ""
func (d *Daos) DisableEmployeeLeaveLog(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.EMPLOYEELEAVELOGSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEELEAVELOG).UpdateOne(ctx.CTX, selector, data)
	return err
}

//DeleteEmployeeLeaveLog :""
func (d *Daos) DeleteEmployeeLeaveLog(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.EMPLOYEELEAVELOGSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEELEAVELOG).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterEmployeeLeaveLog : ""
func (d *Daos) FilterEmployeeLeaveLog(ctx *models.Context, employeeLeaveLog *models.FilterEmployeeLeaveLog, pagination *models.Pagination) ([]models.RefEmployeeLeaveLog, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if employeeLeaveLog != nil {
		if len(employeeLeaveLog.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": employeeLeaveLog.Status}})
		}
		if len(employeeLeaveLog.OrganisationId) > 0 {
			query = append(query, bson.M{"organisationId": bson.M{"$in": employeeLeaveLog.OrganisationId}})
		}
		if len(employeeLeaveLog.EmployeeId) > 0 {
			query = append(query, bson.M{"employeeId": bson.M{"$in": employeeLeaveLog.EmployeeId}})
		}
		//Regex
		if employeeLeaveLog.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: employeeLeaveLog.Regex.Name, Options: "xi"}})
		}
		if employeeLeaveLog.Regex.LeaveType != "" {
			query = append(query, bson.M{"leaveType": primitive.Regex{Pattern: employeeLeaveLog.Regex.LeaveType, Options: "xi"}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEELEAVELOG).CountDocuments(ctx.CTX, func() bson.M {
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
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("Feature query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEELEAVELOG).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var EmployeeLeaveLogFilter []models.RefEmployeeLeaveLog
	if err = cursor.All(context.TODO(), &EmployeeLeaveLogFilter); err != nil {
		return nil, err
	}
	return EmployeeLeaveLogFilter, nil
}
