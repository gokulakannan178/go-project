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

//SavePayrollLog :""
func (d *Daos) SavePayrollLog(ctx *models.Context, payrollLog *models.PayrollLog) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONPAYROLLLOG).InsertOne(ctx.CTX, payrollLog)
	if err != nil {
		return err
	}
	payrollLog.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//GetSinglePayrollLog : ""
func (d *Daos) GetSinglePayrollLog(ctx *models.Context, uniqueID string) (*models.RefPayrollLog, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	// lookup
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPayrollLogCATEGORY, "categoryId", "uniqueId", "ref.category", "ref.category")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPAYROLLLOG).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var payrollLogs []models.RefPayrollLog
	var payrollLog *models.RefPayrollLog
	if err = cursor.All(ctx.CTX, &payrollLogs); err != nil {
		return nil, err
	}
	if len(payrollLogs) > 0 {
		payrollLog = &payrollLogs[0]
	}
	return payrollLog, nil
}

//GetSinglePayrollLog : ""
func (d *Daos) GetSinglePayrollLogWithEmployeeId(ctx *models.Context, uniqueID string) (*models.RefPayrollLog, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"employeeId": uniqueID, "status": constants.PAYROLLLOGSTATUSACTIVE}})
	// lookup
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPayrollLogCATEGORY, "categoryId", "uniqueId", "ref.category", "ref.category")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPAYROLLLOG).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var payrollLogs []models.RefPayrollLog
	var payrollLog *models.RefPayrollLog
	if err = cursor.All(ctx.CTX, &payrollLogs); err != nil {
		return nil, err
	}
	if len(payrollLogs) > 0 {
		payrollLog = &payrollLogs[0]
	}
	return payrollLog, nil
}

//UpdatePayrollLog : ""
func (d *Daos) UpdatePayrollLog(ctx *models.Context, payrollLog *models.PayrollLog) error {
	selector := bson.M{"uniqueId": payrollLog.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": payrollLog}
	_, err := ctx.DB.Collection(constants.COLLECTIONPAYROLLLOG).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//EnablePayrollLog :""
func (d *Daos) EnablePayrollLog(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PAYROLLLOGSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPAYROLLLOG).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisablePayrollLog :""
func (d *Daos) DisablePayrollLog(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PAYROLLLOGSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPAYROLLLOG).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeletePayrollLog :""
func (d *Daos) DeletePayrollLog(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PAYROLLLOGSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPAYROLLLOG).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//FilterPayrollLog : ""
func (d *Daos) FilterPayrollLog(ctx *models.Context, filter *models.FilterPayrollLog, pagination *models.Pagination) ([]models.RefPayrollLog, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONEMPLOYEE, "employeeId", "uniqueId", "ref.employeeId", "ref.employeeId")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDESIGNATION, "ref.employeeId.designationId", "uniqueId", "ref.designationId", "ref.designationId")...)

	query := []bson.M{}
	if filter != nil {

		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
		if len(filter.EmployeeId) > 0 {
			query = append(query, bson.M{"employeeId": bson.M{"$in": filter.EmployeeId}})
		}
		if len(filter.OrganisationID) > 0 {
			query = append(query, bson.M{"organisationId": bson.M{"$in": filter.OrganisationID}})
		}
		if len(filter.DesignationID) > 0 {
			query = append(query, bson.M{"ref.designationId.uniqueId": bson.M{"$in": filter.DesignationID}})
		}
		//Regex
		if filter.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: filter.Regex.Name, Options: "xi"}})
		}
		if filter.Regex.EmployeeName != "" {
			query = append(query, bson.M{"ref.employeeId.name": primitive.Regex{Pattern: filter.Regex.EmployeeName, Options: "xi"}})
		}
	}
	if filter.StartDate != nil {
		sd := time.Date(filter.StartDate.Year(), filter.StartDate.Month(), 1, 0, 0, 0, 0, filter.StartDate.Location())
		query = append(query, bson.M{"startDate": bson.M{"$lte": sd}})
		query = append(query, bson.M{"endDate": bson.M{"$gte": sd}})
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	if filter != nil {
		if filter.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{filter.SortBy: filter.SortOrder}})
		}
	}

	//Adding pagination if necessary
	if pagination != nil {
		paginationPipeline := []bson.M{}
		paginationPipeline = append(paginationPipeline, mainPipeline...)
		paginationPipeline = append(paginationPipeline, bson.M{"$group": bson.M{"_id": nil, "count": bson.M{"$sum": 1}}})
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		d.Shared.BsonToJSONPrintTag("employee timeOff Pagination quary =>", paginationPipeline)

		countcursor, err := ctx.DB.Collection(constants.COLLECTIONPAYROLLLOG).Aggregate(ctx.CTX, paginationPipeline, nil)
		if err != nil {
			log.Println("Error in geting pagination - " + err.Error())
		}
		countstruct := []models.Countstruct{}
		err = countcursor.All(ctx.CTX, &countstruct)
		if err != nil {
			return nil, err
		}
		var totalCount int64
		if len(countstruct) > 0 {
			totalCount = countstruct[0].Count
		}
		fmt.Println("count", totalCount)
		pagination.Count = int(totalCount)
		d.Shared.PaginationData(pagination)
	}
	// lookup
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPayrollLogCATEGORY, "categoryId", "uniqueId", "ref.category", "ref.category")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONBANKINFORMATION, "employeeId", "employeeId", "ref.bank", "ref.bank")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("PayrollLog query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPAYROLLLOG).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var payrollLogs []models.RefPayrollLog
	if err = cursor.All(context.TODO(), &payrollLogs); err != nil {
		return nil, err
	}
	return payrollLogs, nil
}
func (d *Daos) GetSinglePayrollLogWithDays(ctx *models.Context, uniqueID int64) (*models.RefPayrollLog, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"payrollLog": uniqueID}})
	// lookup
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPayrollLogCATEGORY, "categoryId", "uniqueId", "ref.category", "ref.category")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPAYROLLLOG).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var payrollLogs []models.RefPayrollLog
	var payrollLog *models.RefPayrollLog
	if err = cursor.All(ctx.CTX, &payrollLogs); err != nil {
		return nil, err
	}
	if len(payrollLogs) > 0 {
		payrollLog = &payrollLogs[0]
	}
	return payrollLog, nil
}
func (d *Daos) ArchivedPayrollLog(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	t := time.Now()
	t = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())

	update2 := models.Updated{}
	update2.On = &t
	update2.By = constants.SYSTEM
	update := bson.M{"$set": bson.M{"status": constants.EMPLOYEEPAYROLLSTATUSARCHIVED, "updated": update2, "endDate": t}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPAYROLLLOG).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
func (d *Daos) GetEmployeeSalarySlip(ctx *models.Context, filter *models.EmployeeSalarySlip) (*models.RefPayrollLog, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {

		if filter.EmployeeId != "" {
			query = append(query, bson.M{"employeeId": filter.EmployeeId})
		}
		// if filter.StartDate != nil {
		// 	sd := time.Date(filter.StartDate.Year(), filter.StartDate.Month(), filter.StartDate.Day(), 0, 0, 0, 0, filter.StartDate.Location())
		// 	query = append(query, bson.M{"startDate": bson.M{"$lte": sd}})
		// 	query = append(query, bson.M{"endDate": bson.M{"$gte": sd}})
		// }

	}

	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"startYear": bson.M{"$year": "$startDate"}}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"startMonth": bson.M{"$month": "$startDate"}}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"endYear": bson.M{"$year": "$endDate"}}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"endMonth": bson.M{"$month": "$endDate"}}})
	if filter.StartDate != nil {
		sd := time.Date(filter.StartDate.Year(), filter.StartDate.Month(), 1, 0, 0, 0, 0, filter.StartDate.Location())
		mainPipeline = append(mainPipeline, bson.M{
			"$match": bson.M{"$expr": bson.M{
				"$cond": bson.M{"if": bson.M{"$ne": []interface{}{"$endYear", nil}},
					"then": bson.M{"$and": []bson.M{
						{"$lte": []interface{}{"$startYear", sd.Year()}},
						{"$gte": []interface{}{"$endYear", sd.Year()}},
						{"$lte": []interface{}{"$startMonth", sd.Month()}},
						{"$gte": []interface{}{"$endMonth", sd.Month()}},
					}},

					"else": bson.M{"$and": []bson.M{
						{"$lte": []interface{}{"$startYear", sd.Year()}},
						{"$lte": []interface{}{"$startMonth", sd.Month()}},
					},
					}}}},
		})
	}
	// lookup
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPayrollLogCATEGORY, "categoryId", "uniqueId", "ref.category", "ref.category")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONEMPLOYEE, "employeeId", "uniqueId", "ref.employeeId", "ref.employeeId")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDESIGNATION, "ref.employeeId.designationId", "uniqueId", "ref.designationId", "ref.designationId")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONBANKINFORMATION, "employeeId", "employeeId", "ref.bank", "ref.bank")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPAYROLLLOG).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var payrollLogs []models.RefPayrollLog
	var payrollLog *models.RefPayrollLog
	if err = cursor.All(ctx.CTX, &payrollLogs); err != nil {
		return nil, err
	}
	if len(payrollLogs) > 0 {
		payrollLog = &payrollLogs[0]
	}
	return payrollLog, nil
}
func (d *Daos) PayrollLogList(ctx *models.Context, filter *models.FilterPayrollLog, pagination *models.Pagination) ([]models.RefPayrollLog, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONEMPLOYEE, "employeeId", "uniqueId", "ref.employeeId", "ref.employeeId")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDESIGNATION, "ref.employeeId.designationId", "uniqueId", "ref.designationId", "ref.designationId")...)

	query := []bson.M{}
	if filter != nil {

		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
		if len(filter.EmployeeId) > 0 {
			query = append(query, bson.M{"employeeId": bson.M{"$in": filter.EmployeeId}})
		}
		if len(filter.OrganisationID) > 0 {
			query = append(query, bson.M{"organisationId": bson.M{"$in": filter.OrganisationID}})
		}
		if len(filter.DesignationID) > 0 {
			query = append(query, bson.M{"ref.designationId.uniqueId": bson.M{"$in": filter.DesignationID}})
		}
		//Regex
		if filter.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: filter.Regex.Name, Options: "xi"}})
		}
		if filter.Regex.EmployeeName != "" {
			query = append(query, bson.M{"ref.employeeId.name": primitive.Regex{Pattern: filter.Regex.EmployeeName, Options: "xi"}})
		}
	}

	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	if filter != nil {
		if filter.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{filter.SortBy: filter.SortOrder}})
		}
	}
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"startYear": bson.M{"$year": "$startDate"}}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"startMonth": bson.M{"$month": "$startDate"}}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"endYear": bson.M{"$year": "$endDate"}}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"endMonth": bson.M{"$month": "$endDate"}}})
	if filter.StartDate != nil {
		sd := time.Date(filter.StartDate.Year(), filter.StartDate.Month(), 1, 0, 0, 0, 0, filter.StartDate.Location())
		mainPipeline = append(mainPipeline, bson.M{
			"$match": bson.M{"$expr": bson.M{
				"$cond": bson.M{"if": bson.M{"$ne": []interface{}{"$endYear", nil}},
					"then": bson.M{"$and": []bson.M{
						{"$lte": []interface{}{"$startYear", sd.Year()}},
						{"$gte": []interface{}{"$endYear", sd.Year()}},
						{"$lte": []interface{}{"$startMonth", sd.Month()}},
						{"$gte": []interface{}{"$endMonth", sd.Month()}},
					}},

					"else": bson.M{"$and": []bson.M{
						{"$lte": []interface{}{"$startYear", sd.Year()}},
						{"$lte": []interface{}{"$startMonth", sd.Month()}},
					},
					}}}},
		})
	}
	//	d.Shared.BsonToJSONPrintTag("PayrollLog query1111 =>", mainPipeline)
	//Adding pagination if necessary
	if pagination != nil {
		paginationPipeline := []bson.M{}
		paginationPipeline = append(paginationPipeline, mainPipeline...)
		paginationPipeline = append(paginationPipeline, bson.M{"$group": bson.M{"_id": nil, "count": bson.M{"$sum": 1}}})
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		d.Shared.BsonToJSONPrintTag("employee payroll Pagination quary =>", paginationPipeline)

		countcursor, err := ctx.DB.Collection(constants.COLLECTIONPAYROLLLOG).Aggregate(ctx.CTX, paginationPipeline, nil)
		if err != nil {
			log.Println("Error in geting pagination - " + err.Error())
		}
		countstruct := []models.Countstruct{}
		err = countcursor.All(ctx.CTX, &countstruct)
		if err != nil {
			return nil, err
		}
		var totalCount int64
		if len(countstruct) > 0 {
			totalCount = countstruct[0].Count
		}
		fmt.Println("count", totalCount)
		pagination.Count = int(totalCount)
		d.Shared.PaginationData(pagination)
	}
	// lookup
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPayrollLogCATEGORY, "categoryId", "uniqueId", "ref.category", "ref.category")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONBANKINFORMATION, "employeeId", "employeeId", "ref.bank", "ref.bank")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("PayrollLog query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPAYROLLLOG).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var payrollLogs []models.RefPayrollLog
	if err = cursor.All(context.TODO(), &payrollLogs); err != nil {
		return nil, err
	}
	return payrollLogs, nil
}
