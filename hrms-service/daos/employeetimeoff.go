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

// SaveEmployeeTimeOff : ""
func (d *Daos) SaveEmployeeTimeOff(ctx *models.Context, employeeTimeOff *models.EmployeeTimeOff) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEETIMEOFF).InsertOne(ctx.CTX, employeeTimeOff)
	if err != nil {
		return err
	}
	employeeTimeOff.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

// GetSingleEmployeeTimeOff : ""
func (d *Daos) GetSingleEmployeeTimeOff(ctx *models.Context, uniqueID string) (*models.RefEmployeeTimeOff, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	//LookUp
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "approvedBy", "uniqueId", "ref.approvedUser", "ref.approvedUser")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "rejectedBy", "uniqueId", "ref.rejectedUser", "ref.rejectedUser")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONLEAVEMASTER, "leaveType", "uniqueId", "ref.leaveType", "ref.leaveType")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONEMPLOYEE, "employeeId", "uniqueId", "ref.employee", "ref.employee")...)
	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"from": constants.COLLECTIONEMPLOYEELEAVE,
			"as":   "ref.employeeLeave",
			"let":  bson.M{"employeeId": "$employeeId", "leaveType": "$leaveType"},
			"pipeline": []bson.M{
				{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
					{"$eq": []string{"$employeeId", "$$employeeId"}},
					{"$eq": []string{"$leaveType", "$$leaveType"}},
				}}}},
			},
		},
	})
	//
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"ref.employeeLeave": bson.M{"$arrayElemAt": []interface{}{"$ref.employeeLeave", 0}}}})
	d.Shared.BsonToJSONPrintTag("EmployeeTimeOff =>", mainPipeline)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEETIMEOFF).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var EmployeeTimeOffs []models.RefEmployeeTimeOff
	var EmployeeTimeOff *models.RefEmployeeTimeOff
	if err = cursor.All(ctx.CTX, &EmployeeTimeOffs); err != nil {
		return nil, err
	}
	if len(EmployeeTimeOffs) > 0 {
		EmployeeTimeOff = &EmployeeTimeOffs[0]
	}
	return EmployeeTimeOff, err
}

// EmployeeTimeOffCount : ""
func (d *Daos) EmployeeTimeOffCount(ctx *models.Context, employeeId string, OrganisationId string, TimeOffType string) (*models.RefEmployeeTimeOffCount, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"employeeId": employeeId, "organisationId": OrganisationId, "timeoffType": TimeOffType}})
	group := []bson.M{{"$group": bson.M{"_id": nil, "totalTimeOff": bson.M{"$sum": "$numberOfDays"}}}}
	mainPipeline = append(mainPipeline, group...)

	d.Shared.BsonToJSONPrintTag("EmployeeTimeOffCount =>", mainPipeline)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEETIMEOFF).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var EmployeeTimeOffs []models.RefEmployeeTimeOffCount
	var EmployeeTimeOff *models.RefEmployeeTimeOffCount
	if err = cursor.All(ctx.CTX, &EmployeeTimeOffs); err != nil {
		return nil, err
	}

	if len(EmployeeTimeOffs) > 0 {
		EmployeeTimeOff = &EmployeeTimeOffs[0]
	}
	return EmployeeTimeOff, nil
}

// UpdateEmployeeTimeOff : ""
func (d *Daos) UpdateEmployeeTimeOff(ctx *models.Context, employeeTimeOff *models.EmployeeTimeOff) error {
	selector := bson.M{"uniqueId": employeeTimeOff.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": employeeTimeOff}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEETIMEOFF).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// EnableEmployeeTimeOff : ""
func (d *Daos) EnableEmployeeTimeOff(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.EMPLOYEETIMEOFFSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEETIMEOFF).UpdateOne(ctx.CTX, selector, data)
	return err
}

// DisableEmployeeTimeOff : ""
func (d *Daos) DisableEmployeeTimeOff(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.EMPLOYEETIMEOFFSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEETIMEOFF).UpdateOne(ctx.CTX, selector, data)
	return err
}

// DeleteEmployeeTimeOff :""
func (d *Daos) DeleteEmployeeTimeOff(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.EMPLOYEETIMEOFFSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEETIMEOFF).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterEmployeeTimeOff : ""
func (d *Daos) FilterEmployeeTimeOff(ctx *models.Context, employeeTimeOff *models.FilterEmployeeTimeOff, pagination *models.Pagination) ([]models.RefEmployeeTimeOff, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONEMPLOYEE, "employeeId", "uniqueId", "ref.employee", "ref.employee")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONLEAVEMASTER, "leaveType", "uniqueId", "ref.leaveType", "ref.leaveType")...)
	if employeeTimeOff != nil {
		if len(employeeTimeOff.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": employeeTimeOff.Status}})
		}
		if len(employeeTimeOff.OmitStatus) > 0 {
			query = append(query, bson.M{"status": bson.M{"$nin": employeeTimeOff.OmitStatus}})
		}
		if len(employeeTimeOff.OrganisationId) > 0 {
			query = append(query, bson.M{"organisationId": bson.M{"$in": employeeTimeOff.OrganisationId}})
		}
		if len(employeeTimeOff.EmployeeId) > 0 {
			query = append(query, bson.M{"employeeId": bson.M{"$in": employeeTimeOff.EmployeeId}})
		}
		if len(employeeTimeOff.RevokeRequest) > 0 {
			query = append(query, bson.M{"revoke": bson.M{"$in": employeeTimeOff.RevokeRequest}})
		}
		if len(employeeTimeOff.OmitRevokeRequest) > 0 {
			query = append(query, bson.M{"revoke": bson.M{"$nin": employeeTimeOff.OmitRevokeRequest}})
		}
		if len(employeeTimeOff.LeaveType) > 0 {
			query = append(query, bson.M{"ref.leaveType.uniqueId": bson.M{"$in": employeeTimeOff.LeaveType}})
		}
		if employeeTimeOff.Manager != "" {
			query = append(query, bson.M{"ref.employee.lineManager": employeeTimeOff.Manager})
		}
		//Regex
		if employeeTimeOff.Regex.EmployeeName != "" {
			query = append(query, bson.M{"ref.employee.name": primitive.Regex{Pattern: employeeTimeOff.Regex.EmployeeName, Options: "xi"}})
		}
	}
	if employeeTimeOff.DateRange.From != nil {
		t := *employeeTimeOff.DateRange.From
		FromDate := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
		ToDate := time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, t.Location())
		if employeeTimeOff.DateRange.To != nil {
			t2 := *employeeTimeOff.DateRange.To
			ToDate = time.Date(t2.Year(), t2.Month(), t2.Day(), 23, 59, 59, 0, t2.Location())
		}
		query = append(query, bson.M{"startDate": bson.M{"$gte": FromDate, "$lte": ToDate}})
		query = append(query, bson.M{"endDate": bson.M{"$gte": FromDate, "$lte": ToDate}})
	}

	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if employeeTimeOff != nil {
		if employeeTimeOff.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{employeeTimeOff.SortBy: employeeTimeOff.SortOrder}})
		}
	}
	mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"_id": -1}})
	//Adding pagination if necessary
	if pagination != nil {
		paginationPipeline := []bson.M{}
		paginationPipeline = append(paginationPipeline, mainPipeline...)
		paginationPipeline = append(paginationPipeline, bson.M{"$group": bson.M{"_id": nil, "count": bson.M{"$sum": 1}}})
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		d.Shared.BsonToJSONPrintTag("employee timeOff Pagination quary =>", paginationPipeline)

		countcursor, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEETIMEOFF).Aggregate(ctx.CTX, paginationPipeline, nil)
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
	//LookUp
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "approvedBy", "employeeId", "ref.approvedUser", "ref.approvedUser")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "rejectedBy", "employeeId", "ref.rejectedUser", "ref.rejectedUser")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONBRANCH, "ref.employee.branchId", "uniqueId", "ref.branch", "ref.branch")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDESIGNATION, "ref.employee.designationId", "uniqueId", "ref.designation", "ref.designation")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("Employeetimeoff query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEETIMEOFF).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var EmployeeTimeOffFilter []models.RefEmployeeTimeOff
	if err = cursor.All(context.TODO(), &EmployeeTimeOffFilter); err != nil {
		return nil, err
	}
	return EmployeeTimeOffFilter, nil
}

// EmployeeTimeOffRequest : ""
func (d *Daos) EmployeeTimeOffRequest(ctx *models.Context, employeeTimeOff *models.EmployeeTimeOff) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEETIMEOFF).InsertOne(ctx.CTX, employeeTimeOff)
	if err != nil {
		return err
	}
	employeeTimeOff.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

// EmployeeTimeOffApprove : ""
func (d *Daos) EmployeeTimeOffApprove(ctx *models.Context, employeeTimeOff *models.ReviewedEmployeeTimeOff) error {
	selector := bson.M{"uniqueId": employeeTimeOff.EmployeeTimeOff}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": bson.M{"approvedBy": employeeTimeOff.ReviewedBy, "approvedDate": t, "status": constants.EMPLOYEETIMEOFFSTATUSAPPROVE, "revoke": constants.EMPLOYEETIMEOFFSTATUSAPPROVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEETIMEOFF).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// EmployeeTimeOffReject : ""
func (d *Daos) EmployeeTimeOffReject(ctx *models.Context, employeeTimeOff *models.ReviewedEmployeeTimeOff) error {
	selector := bson.M{"uniqueId": employeeTimeOff.EmployeeTimeOff}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	//updateInterface := bson.M{"$set": employeeTimeOff}
	updateInterface := bson.M{"$set": bson.M{"status": constants.EMPLOYEETIMEOFFSTATUSREJECT, "rejectedBy": employeeTimeOff.ReviewedBy, "rejectedDate": &t, "remarks": employeeTimeOff.Remarks, "revoke": constants.EMPLOYEETIMEOFFSTATUSREJECT}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEETIMEOFF).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// EmployeeTimeOffRevoke : ""
func (d *Daos) EmployeeTimeOffRevoke(ctx *models.Context, employeeTimeOff *models.ReviewedEmployeeTimeOff) error {
	selector := bson.M{"uniqueId": employeeTimeOff.EmployeeTimeOff}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	//updateInterface := bson.M{"$set": employeeTimeOff}
	updateInterface := bson.M{"$set": bson.M{"status": constants.EMPLOYEETIMEOFFSTATUSREVOKE, "rejectedBy": employeeTimeOff.ReviewedBy, "rejectedDate": &t, "remarks": employeeTimeOff.Remarks, "revoke": constants.EMPLOYEETIMEOFFSTATUSREVOKE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEETIMEOFF).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}
func (d *Daos) GetTodayApprovedEmployeeTimeOff(ctx *models.Context) ([]models.EmployeeTimeOff, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	query = append(query, bson.M{"status": constants.EMPLOYEETIMEOFFSTATUSAPPROVE})
	//query = append(query, bson.M{"status": constants.EMPLOYEETIMEOFFSTATUSAPPROVE})
	t := time.Now()
	FromDate := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	ToDate := time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, t.Location())
	query = append(query, bson.M{"startDate": bson.M{"$gte": FromDate, "$lte": ToDate}})
	query = append(query, bson.M{"endDate": bson.M{"$gte": FromDate, "$lte": ToDate}})

	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	//Aggregation
	d.Shared.BsonToJSONPrintTag("approved time off query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEETIMEOFF).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var Employees []models.EmployeeTimeOff
	if err = cursor.All(context.TODO(), &Employees); err != nil {
		return nil, err
	}
	return Employees, nil
}

// CancelEmployeeTimeOff
func (d *Daos) CancelEmployeeTimeOff(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.EMPLOYEETIMEOFFSTATUSCANCEL, "revoke": constants.EMPLOYEETIMEOFFSTATUSCANCEL}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEETIMEOFF).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// RevokeRequestEmployeeTimeOff
func (d *Daos) RevokeRequestEmployeeTimeOff(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"revoke": constants.EMPLOYEETIMEOFFSTATUSREVOKEREQUEST}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEETIMEOFF).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
func (d *Daos) GetTodayEmployeeApprovedTimeOff(ctx *models.Context, filter *models.DayWiseAttendanceReportFilter) ([]models.RefEmployeeTimeOff, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	t := time.Now()
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONEMPLOYEE, "employeeId", "uniqueId", "ref.employee", "ref.employee")...)

	sd := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)
	// ed := time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, t.Location())
	query = append(query, bson.M{"status": constants.EMPLOYEETIMEOFFSTATUSAPPROVE})
	query = append(query, bson.M{"startDate": bson.M{"$lte": sd}})
	query = append(query, bson.M{"endDate": bson.M{"$gte": sd}})
	if filter != nil {
		if len(filter.OrganisationId) > 0 {
			query = append(query, bson.M{"organisationId": bson.M{"$in": filter.OrganisationId}})
		}
		if len(filter.DepartmentId) > 0 {
			query = append(query, bson.M{"departmentId": bson.M{"$in": filter.DepartmentId}})
		}
		if filter.Manager != "" {
			query = append(query, bson.M{"lineManager": filter.Manager})
		}
	}
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	//lookup

	d.Shared.BsonToJSONPrintTag("GetTodayEmployeeApprovedTimeOff =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEETIMEOFF).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var attendances []models.RefEmployeeTimeOff
	if err = cursor.All(ctx.CTX, &attendances); err != nil {
		return nil, err
	}
	return attendances, err
}
func (d *Daos) GetTodayEmployeePendingTimeOff(ctx *models.Context, filter *models.DayWiseAttendanceReportFilter) ([]models.RefEmployeeTimeOff, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	t := time.Now()
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONEMPLOYEE, "employeeId", "uniqueId", "ref.employee", "ref.employee")...)

	sd := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)
	// ed := time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, t.Location())
	query = append(query, bson.M{"status": constants.EMPLOYEETIMEOFFSTATUSREQUEST})
	query = append(query, bson.M{"startDate": bson.M{"$lte": sd}})
	query = append(query, bson.M{"endDate": bson.M{"$gte": sd}})
	if filter != nil {
		if len(filter.OrganisationId) > 0 {
			query = append(query, bson.M{"organisationId": bson.M{"$in": filter.OrganisationId}})
		}
		if len(filter.DepartmentId) > 0 {
			query = append(query, bson.M{"departmentId": bson.M{"$in": filter.DepartmentId}})
		}
		if filter.Manager != "" {
			query = append(query, bson.M{"lineManager": filter.Manager})
		}
	}
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	//lookup

	d.Shared.BsonToJSONPrintTag("GetTodayEmployeePendingTimeOff =>", mainPipeline)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEETIMEOFF).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var attendances []models.RefEmployeeTimeOff
	if err = cursor.All(ctx.CTX, &attendances); err != nil {
		return nil, err
	}
	return attendances, err
}
func (d *Daos) EmployeeTimeoffDateCheck(ctx *models.Context, filter *models.DayWiseAttendanceReportFilter) (*models.RefEmployeeTimeOff, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}

	if filter != nil {
		if filter.EmployeeId != "" {
			query = append(query, bson.M{"employeeId": filter.EmployeeId})
		}
		if filter.StartDate != nil {
			sd := time.Date(filter.StartDate.Year(), filter.StartDate.Month(), filter.StartDate.Day(), 0, 0, 0, 0, time.UTC)
			query = append(query, bson.M{"startDate": bson.M{"$lte": sd}})
			query = append(query, bson.M{"endDate": bson.M{"$gte": sd}})
		}
		query = append(query, bson.M{"status": bson.M{"$in": []string{constants.EMPLOYEETIMEOFFSTATUSREQUEST, constants.EMPLOYEETIMEOFFSTATUSAPPROVE}}})
	}
	//lookup
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	d.Shared.BsonToJSONPrintTag("GetTodayEmployeePendingTimeOff =>", mainPipeline)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEETIMEOFF).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var attendances []models.RefEmployeeTimeOff
	var attendance *models.RefEmployeeTimeOff
	if err = cursor.All(ctx.CTX, &attendances); err != nil {
		return nil, err
	}
	if len(attendances) > 0 {
		attendance = &attendances[0]
	}
	return attendance, err
}
