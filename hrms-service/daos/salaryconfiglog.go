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

//SaveSalaryConfigLog : ""
func (d *Daos) SaveSalaryConfigLog(ctx *models.Context, salaryConfigLog *models.SalaryConfigLog) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONSALARYCONFIGLOG).InsertOne(ctx.CTX, salaryConfigLog)
	if err != nil {
		return err
	}
	salaryConfigLog.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//UpdateSalaryConfigLog : ""
func (d *Daos) UpdateSalaryConfigLog(ctx *models.Context, SalaryConfigLog *models.SalaryConfigLog) error {
	selector := bson.M{"uniqueId": SalaryConfigLog.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": SalaryConfigLog}
	_, err := ctx.DB.Collection(constants.COLLECTIONSALARYCONFIGLOG).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// GetSingleSalaryConfigLog : ""
func (d *Daos) GetSingleSalaryConfigLog(ctx *models.Context, uniqueID string) (*models.RefSalaryConfigLog, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	//LookUp
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSALARYCONFIGLOG).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var SalaryConfigLogs []models.RefSalaryConfigLog
	var SalaryConfigLog *models.RefSalaryConfigLog
	if err = cursor.All(ctx.CTX, &SalaryConfigLogs); err != nil {
		return nil, err
	}
	if len(SalaryConfigLogs) > 0 {
		SalaryConfigLog = &SalaryConfigLogs[0]
	}
	return SalaryConfigLog, err
}

// EnableSalaryConfigLog : ""
func (d *Daos) EnableSalaryConfigLog(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.SALARYCONFIGLOGSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONSALARYCONFIGLOG).UpdateOne(ctx.CTX, selector, data)
	return err
}

// DisableSalaryConfigLog : ""
func (d *Daos) DisableSalaryConfigLog(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.SALARYCONFIGLOGSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONSALARYCONFIGLOG).UpdateOne(ctx.CTX, selector, data)
	return err
}

//DeleteSalaryConfigLog :""
func (d *Daos) DeleteSalaryConfigLog(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.SALARYCONFIGLOGSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONSALARYCONFIGLOG).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterSalaryConfigLog : ""
func (d *Daos) FilterSalaryConfigLog(ctx *models.Context, salaryConfigLog *models.FilterSalaryConfigLog, pagination *models.Pagination) ([]models.RefSalaryConfigLog, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if salaryConfigLog != nil {
		if len(salaryConfigLog.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": salaryConfigLog.Status}})
		}
		if len(salaryConfigLog.EmployeeID) > 0 {
			query = append(query, bson.M{"employeeID": bson.M{"$in": salaryConfigLog.EmployeeID}})
		}
		if len(salaryConfigLog.OrganisationID) > 0 {
			query = append(query, bson.M{"organisationId": bson.M{"$in": salaryConfigLog.OrganisationID}})
		}
		//Regex

	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if salaryConfigLog != nil {
		if salaryConfigLog.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{salaryConfigLog.SortBy: salaryConfigLog.SortOrder}})
		}
	}
	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONSALARYCONFIGLOG).CountDocuments(ctx.CTX, func() bson.M {
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
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSALARYCONFIGLOG).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var salaryConfigLogFilter []models.RefSalaryConfigLog
	if err = cursor.All(context.TODO(), &salaryConfigLogFilter); err != nil {
		return nil, err
	}
	return salaryConfigLogFilter, nil
}
func (d *Daos) GetSingleSalaryConfigLogWithEmployeeType(ctx *models.Context, uniqueID string) (*models.RefSalaryConfigLog, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"employeeType": uniqueID, "status": constants.SALARYCONFIGLOGSTATUSACTIVE}})
	//LookUp
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSALARYCONFIGLOG).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var SalaryConfigLogs []models.RefSalaryConfigLog
	var SalaryConfigLog *models.RefSalaryConfigLog
	if err = cursor.All(ctx.CTX, &SalaryConfigLogs); err != nil {
		return nil, err
	}
	if len(SalaryConfigLogs) > 0 {
		SalaryConfigLog = &SalaryConfigLogs[0]
	}
	return SalaryConfigLog, err
}
