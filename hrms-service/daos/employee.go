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

//SaveEmployee :""
func (d *Daos) SaveEmployee(ctx *models.Context, employee *models.Employee) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEE).InsertOne(ctx.CTX, employee)
	if err != nil {
		return err
	}
	employee.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//GetSingleEmployee : ""
func (d *Daos) GetSingleEmployee(ctx *models.Context, uniqueID string) (*models.RefEmployee, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	//LookUp
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONGRADE, "grade", "uniqueId", "ref.grade", "ref.grade")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDEPARTMENT, "departmentId", "uniqueId", "ref.departmentId", "ref.departmentId")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONBRANCH, "branchId", "uniqueId", "ref.branchId", "ref.branchId")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONBANKINFORMATION, "uniqueId", "employeeId", "ref.bank", "ref.bank")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDESIGNATION, "designationId", "uniqueId", "ref.designationId", "ref.designationId")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "uniqueId", "employeeId", "ref.userId", "ref.userId")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONWORKSCHEDULE, "workscheduleId", "uniqueId", "ref.workschedule", "ref.workschedule")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "lineManager", "userName", "ref.lineManager", "ref.lineManager")...)
	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"from": constants.COLLECTIONEMPLOYEEFAMILYMEMBERS,
			"as":   "ref.familyMembers",
			"let":  bson.M{"employeeID": "$uniqueId"},
			"pipeline": []bson.M{
				{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
					bson.M{"$eq": []string{"$status", "Active"}},
					bson.M{"$eq": []string{"$employeeID", "$$employeeID"}},
				}}}}},
		},
	})
	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"from": constants.COLLECTIONEMPLOYEEEDUCATION,
			"as":   "ref.education",
			"let":  bson.M{"employeeID": "$uniqueId"},
			"pipeline": []bson.M{
				{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
					bson.M{"$eq": []string{"$status", "Active"}},
					bson.M{"$eq": []string{"$employeeID", "$$employeeID"}},
				}}}}},
		},
	})
	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"from": constants.COLLECTIONEMPLOYEEEXPERIENCE,
			"as":   "ref.experience",
			"let":  bson.M{"employeeID": "$uniqueId"},
			"pipeline": []bson.M{
				{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
					bson.M{"$eq": []string{"$status", "Active"}},
					bson.M{"$eq": []string{"$employeeID", "$$employeeID"}},
				}}}}},
		},
	})
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONNOTICEPOLICY, "noticeId", "uniqueId", "ref.notice", "ref.notice")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDOCUMENTPOLICY, "documentPolicyID", "uniqueId", "ref.documentPolicy", "ref.documentPolicy")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONLEAVEPOLICY, "leavePolicyID", "uniqueId", "ref.leavePolicy", "ref.leavePolicy")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPROBATIONARY, "probationaryId", "uniqueId", "ref.probationary", "ref.probationary")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPAYROLLPOLICY, "payrollPolicyId", "uniqueId", "ref.payrollPolicyId", "ref.payrollPolicyId")...)

	d.Shared.BsonToJSONPrintTag("Employee query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var Employees []models.RefEmployee
	var Employee *models.RefEmployee
	if err = cursor.All(ctx.CTX, &Employees); err != nil {
		return nil, err
	}
	if len(Employees) > 0 {
		Employee = &Employees[0]
	}
	return Employee, nil
}

//UpdateEmployee : ""
func (d *Daos) UpdateEmployee(ctx *models.Context, Employee *models.Employee) error {
	selector := bson.M{"uniqueId": Employee.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": Employee}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEE).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//EnableEmployee :""
func (d *Daos) EnableEmployee(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.EMPLOYEESTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableEmployee :""
func (d *Daos) DisableEmployee(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.EMPLOYEESTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteEmployee :""
func (d *Daos) DeleteEmployee(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.EMPLOYEESTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//FilterEmployee : ""
func (d *Daos) FilterEmployee(ctx *models.Context, employeefilter *models.FilterEmployee, pagination *models.Pagination) ([]models.RefEmployee, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if employeefilter != nil {

		if len(employeefilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": employeefilter.Status}})
		}
		if len(employeefilter.OrganisationID) > 0 {
			query = append(query, bson.M{"organisationId": bson.M{"$in": employeefilter.OrganisationID}})
		}
		if len(employeefilter.DepartmentID) > 0 {
			query = append(query, bson.M{"departmentId": bson.M{"$in": employeefilter.DepartmentID}})
		}
		if len(employeefilter.BranchID) > 0 {
			query = append(query, bson.M{"branchId": bson.M{"$in": employeefilter.BranchID}})
		}
		if len(employeefilter.DesignationID) > 0 {
			query = append(query, bson.M{"designationId": bson.M{"$in": employeefilter.DesignationID}})
		}
		if len(employeefilter.UniqueID) > 0 {
			query = append(query, bson.M{"uniqueId": bson.M{"$in": employeefilter.UniqueID}})
		}
		if len(employeefilter.Grade) > 0 {
			query = append(query, bson.M{"grade": bson.M{"$in": employeefilter.Grade}})
		}

		//Regex
		if employeefilter.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: employeefilter.Regex.Name, Options: "xi"}})
		}

	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if employeefilter != nil {
		if employeefilter.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{employeefilter.SortBy: employeefilter.SortOrder}})
		}
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEE).CountDocuments(ctx.CTX, func() bson.M {
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
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONGRADE, "grade", "uniqueId", "ref.grade", "ref.grade")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDEPARTMENT, "departmentId", "uniqueId", "ref.departmentId", "ref.departmentId")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONBRANCH, "branchId", "uniqueId", "ref.branchId", "ref.branchId")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDESIGNATION, "designationId", "uniqueId", "ref.designationId", "ref.designationId")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "uniqueId", "employeeId", "ref.userId", "ref.userId")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONWORKSCHEDULE, "workscheduleId", "uniqueId", "ref.workschedule", "ref.workschedule")...)
	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"from": constants.COLLECTIONEMPLOYEEFAMILYMEMBERS,
			"as":   "ref.familyMembers",
			"let":  bson.M{"employeeID": "$uniqueId"},
			"pipeline": []bson.M{
				{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
					bson.M{"$eq": []string{"$status", "Active"}},
					bson.M{"$eq": []string{"$employeeID", "$$employeeID"}},
				}}}}},
		},
	})
	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"from": constants.COLLECTIONEMPLOYEEEDUCATION,
			"as":   "ref.education",
			"let":  bson.M{"employeeID": "$uniqueId"},
			"pipeline": []bson.M{
				{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
					bson.M{"$eq": []string{"$status", "Active"}},
					bson.M{"$eq": []string{"$employeeID", "$$employeeID"}},
				}}}}},
		},
	})
	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"from": constants.COLLECTIONEMPLOYEEEXPERIENCE,
			"as":   "ref.experience",
			"let":  bson.M{"employeeID": "$uniqueId"},
			"pipeline": []bson.M{
				{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
					bson.M{"$eq": []string{"$status", "Active"}},
					bson.M{"$eq": []string{"$employeeID", "$$employeeID"}},
				}}}}},
		},
	})
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONNOTICEPOLICY, "noticeId", "uniqueId", "ref.notice", "ref.notice")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDOCUMENTPOLICY, "documentPolicyID", "uniqueId", "ref.documentPolicy", "ref.documentPolicy")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONLEAVEPOLICY, "leavePolicyID", "uniqueId", "ref.leavePolicy", "ref.leavePolicy")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPROBATIONARY, "probationaryId", "uniqueId", "ref.probationary", "ref.probationary")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPAYROLLPOLICY, "payrollPolicyId", "uniqueId", "ref.payrollPolicyId", "ref.payrollPolicyId")...)
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONEMPLOYEE, "lineManager", "uniqueId", "ref.lineManager", "ref.lineManager")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "lineManager", "userName", "ref.lineManager", "ref.lineManager")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("Employee query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var Employees []models.RefEmployee
	if err = cursor.All(context.TODO(), &Employees); err != nil {
		return nil, err
	}
	return Employees, nil
}

//EmployeeReject : ""
func (d *Daos) EmployeeReject(ctx *models.Context, Employee *models.Employee) error {
	selector := bson.M{"uniqueId": Employee.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": Employee}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEE).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//EmployeeOnboarding : ""
func (d *Daos) EmployeeOnboarding(ctx *models.Context, Employee *models.Employee) error {
	selector := bson.M{"uniqueId": Employee.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": Employee}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEE).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//EmployeeProbationary : ""
func (d *Daos) EmployeeProbationary(ctx *models.Context, Employee *models.Employee) error {
	selector := bson.M{"uniqueId": Employee.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": Employee}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEE).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//EmployeeActive : ""
func (d *Daos) EmployeeActive(ctx *models.Context, Employee *models.Employee) error {
	selector := bson.M{"uniqueId": Employee.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": Employee}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEE).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//EmployeeBench : ""
func (d *Daos) EmployeeBench(ctx *models.Context, Employee *models.Employee) error {
	selector := bson.M{"uniqueId": Employee.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": Employee}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEE).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//EmployeeNotice : ""
func (d *Daos) EmployeeNotice(ctx *models.Context, Employee *models.Employee) error {
	selector := bson.M{"uniqueId": Employee.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": Employee}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEE).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//EmployeeOffboard : ""
func (d *Daos) EmployeeOffboard(ctx *models.Context, Employee *models.Employee) error {
	selector := bson.M{"uniqueId": Employee.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": Employee}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEE).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//EmployeeRelieve : ""
func (d *Daos) EmployeeRelieve(ctx *models.Context, Employee *models.Employee) error {
	selector := bson.M{"uniqueId": Employee.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": Employee}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEE).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}
func (d *Daos) UpdateEmployeeBioData(ctx *models.Context, Employee *models.UpdateBioData, UniqueID string) error {
	selector := bson.M{"uniqueId": UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": Employee}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEE).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}
func (d *Daos) UpdateEmployeeEmergencyContact(ctx *models.Context, Employee *models.UpdateEmergencyContact, UniqueID string) error {
	selector := bson.M{"uniqueId": UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": bson.M{"emergencyContact": Employee}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEE).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}
func (d *Daos) UpdateEmployeePersonalInformation(ctx *models.Context, Employee *models.PersonalInformation, UniqueID string) error {
	selector := bson.M{"uniqueId": UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": bson.M{"personalInformation": Employee}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEE).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}
func (d *Daos) EmployeeDayWiseAttendanceReportV2(ctx *models.Context, filter *models.DayWiseAttendanceReportFilter) ([]models.EmployeeDayWiseAttendanceReport, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	var sd, ed time.Time
	//	var attendances *models.DayWiseAttendanceReport

	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONATTENDANCE, "employeeId", "uniqueId", "attendances", "attendances")...)

	if filter != nil {
		query = append(query, bson.M{"status": "Active"})
		if filter.StartDate != nil {
			sd = time.Date(filter.StartDate.Year(), filter.StartDate.Month(), 1, 0, 0, 0, 0, filter.StartDate.Location())
			ed = time.Date(filter.StartDate.Year(), filter.StartDate.Month()+1, 0, 23, 59, 59, 999999999, filter.StartDate.Location())

			query = append(query, bson.M{"attendances.date": bson.M{"$gte": sd, "$lte": ed}})
			fmt.Println("sd===>", sd)
			fmt.Println("ed===>", ed)

		}
	}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})

	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"dom": bson.M{"$dayOfMonth": "$date"}}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"month": bson.M{"$month": "$date"}}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"year": bson.M{"$year": "$date"}}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"dayOfWeek": bson.M{"$dayOfWeek": "$date"}}})
	//fmt.Println("dayofweek==>", attendance)
	mainPipeline = append(mainPipeline, bson.M{"$group": bson.M{"_id": bson.M{"employeeId": "$employeeId"}, "days": bson.M{"$push": "$$ROOT"}}})
	mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"date": 1}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"employeeId": "$_id.employeeId"}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"month1": bson.M{"$arrayElemAt": []interface{}{"$days.month", 0}}}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"year1": bson.M{"$arrayElemAt": []interface{}{"$days.year", 0}}}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"workscheduleId": bson.M{"$arrayElemAt": []interface{}{"$days.workscheduleId", 0}}}})
	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"from":         constants.COLLECTIONWORKSCHEDULE,
			"as":           "workinghourse",
			"localField":   "workscheduleId",
			"foreignField": "uniqueId",
		},
	})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"days.workinghourse": bson.M{"$arrayElemAt": []interface{}{"$workinghourse.workingHoursinDay", 0}}}})

	//mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"dayOfWeek1": bson.M{"$arrayElemAt": []interface{}{"$days.dayOfWeek", 0}}}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{
		"days": bson.M{
			"$map": bson.M{
				"input": bson.M{"$range": []interface{}{sd.Day(), ed.Day() + 1, 1}},
				"as":    "rangeDay",
				"in": bson.M{
					"$let": bson.M{
						"vars": bson.M{"index": bson.M{"$indexOfArray": []string{"$days.dom", "$$rangeDay"}}},
						"in": bson.M{
							"$cond": bson.M{
								"if": bson.M{"$eq": []interface{}{"$$index", -1}},
								"then": bson.M{
									"date":          bson.M{"$dateFromParts": bson.M{"year": "$year1", "month": "$month1", "day": "$$rangeDay", "hour": 9}},
									"workinghourse": bson.M{"$arrayElemAt": []interface{}{"$days.workinghourse", 0}},
									//"attendance":    false,
									//"dayOfWeek": bson.M{"$dayOfWeek": bson.M{"$dateFromParts": bson.M{"year": "$year1", "month": "$month1", "day": "$$rangeDay", "hour": 9}}},
								},
								"else": bson.M{
									"date":             bson.M{"$dateFromParts": bson.M{"year": "$year1", "month": "$month1", "day": "$$rangeDay", "hour": 9}},
									"workinghourse":    bson.M{"$arrayElemAt": []interface{}{"$days.workinghourse", 0}},
									"totalWorkingMins": bson.M{"$arrayElemAt": []interface{}{"$days.totalWorkingMins", 0}},
									//	"attendance": true,
									//"dayOfWeek": bson.M{"$dayOfWeek": bson.M{"$dateFromParts": bson.M{"year": "$year1", "month": "$month1", "day": "$$rangeDay", "hour": 9}}},

								}},
						},
					},
				},
			},
		},
	}})

	//Aggregation
	d.Shared.BsonToJSONPrint(mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var attendance []models.EmployeeDayWiseAttendanceReport

	if err = cursor.All(context.TODO(), &attendance); err != nil {
		return nil, err
	}
	return attendance, nil
}
func (d *Daos) EmployeeDayWiseAttendanceReport(ctx *models.Context, filter *models.DayWiseAttendanceReportFilter, pagination *models.Pagination) ([]models.EmployeeDayWiseAttendanceReport, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	var sd, ed time.Time
	//	var attendances *models.DayWiseAttendanceReport
	var attendance []models.EmployeeDayWiseAttendanceReport
	t := time.Now()
	if filter != nil {
		if len(filter.OrganisationId) > 0 {
			query = append(query, bson.M{"organisationId": bson.M{"$in": filter.OrganisationId}})
		}
		if len(filter.DepartmentId) > 0 {
			query = append(query, bson.M{"departmentId": bson.M{"$in": filter.DepartmentId}})
		}
		if filter.EmployeeId != "" {
			query = append(query, bson.M{"uniqueId": filter.EmployeeId})
		}
		if filter.Manager != "" {
			query = append(query, bson.M{"lineManager": filter.Manager})
		}
		query = append(query, bson.M{"status": bson.M{"$nin": []string{constants.EMPLOYEESTATUSDELETED, constants.EMPLOYEESTATUSREJECT, constants.EMPLOYEESTATUSRELIEVE, constants.EMPLOYEESTATUSOFFBOARD, constants.EMPLOYEESTATUSONBORADING}}})

		if filter.StartDate != nil {
			sd = time.Date(filter.StartDate.Year(), filter.StartDate.Month(), 1, 0, 0, 0, 0, filter.StartDate.Location())
			ed = time.Date(filter.StartDate.Year(), filter.StartDate.Month()+1, 0, 23, 59, 59, 999999999, filter.StartDate.Location())
			if t.Month() == filter.StartDate.Month() && t.Year() == filter.StartDate.Year() {
				ed = time.Date(filter.StartDate.Year(), filter.StartDate.Month(), t.Day(), 23, 59, 59, 999999999, filter.StartDate.Location())

			}
			//	query = append(query, bson.M{"date": bson.M{"$gte": sd, "$lte": ed}})
			fmt.Println("sd===>", sd)
			fmt.Println("ed===>", ed)

		}
		if filter.SearchBox.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: filter.SearchBox.Name, Options: "xi"}})
		}
	}
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if filter != nil {
		if filter.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{filter.SortBy: filter.SortOrder}})
		}
	}
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"employeeId": "$uniqueId"}})
	mainPipeline = append(mainPipeline, bson.M{

		"$lookup": bson.M{
			"as":   "days",
			"from": constants.COLLECTIONATTENDANCE,
			"let":  bson.M{"employeeId": "$uniqueId"},
			"pipeline": []bson.M{
				{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
					//	bson.M{"$eq": []string{"$status", "Active"}},
					{"$eq": []string{"$employeeId", "$$employeeId"}},
					{"$gte": []interface{}{"$date", sd}},
					{"$lte": []interface{}{"$date", ed}},
				}}}},
				{"$addFields": bson.M{"dom": bson.M{"$dayOfMonth": "$date"}}},
				{"$addFields": bson.M{"month": bson.M{"$month": "$date"}}},
				{"$addFields": bson.M{"year": bson.M{"$year": "$date"}}},
				{"$lookup": bson.M{
					"as":   "attandanceLog",
					"from": constants.COLLECTIONATTENDANCELOG,
					"let":  bson.M{"attedanceId": "$uniqueId", "employeeId": "$employeeId"},
					"pipeline": []bson.M{
						{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
							//	bson.M{"$eq": []string{"$status", "Active"}},
							{"$eq": []string{"$employeeId", "$$employeeId"}},
							{"$eq": []string{"$attendanceId", "$$attedanceId"}},
						}}}},
					},
				}},
			}},
	})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"month1": bson.M{"$arrayElemAt": []interface{}{"$days.month", 0}}}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"year1": bson.M{"$arrayElemAt": []interface{}{"$days.year", 0}}}})

	//mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"dayOfWeek1": bson.M{"$arrayElemAt": []interface{}{"$days.dayOfWeek", 0}}}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{
		"days": bson.M{
			"$map": bson.M{
				"input": bson.M{"$range": []interface{}{sd.Day(), ed.Day() + 1, 1}},
				"as":    "rangeDay",
				"in": bson.M{
					"$let": bson.M{
						"vars": bson.M{
							"index": bson.M{"$indexOfArray": []string{"$days.dom", "$$rangeDay"}},
							//	"workingHours":""

						},
						"in": bson.M{
							"$cond": bson.M{
								"if": bson.M{"$eq": []interface{}{"$$index", -1}},
								"then": bson.M{
									"date":       bson.M{"$dateFromParts": bson.M{"year": "$year1", "month": "$month1", "day": "$$rangeDay", "hour": 9}},
									"attendance": bson.M{"payRoll": constants.ATTENDANCESTATUSLOP},
								},
								"else": bson.M{
									"date": bson.M{"$dateFromParts": bson.M{"year": "$year1", "month": "$month1", "day": "$$rangeDay", "hour": 9}},

									"attendance": bson.M{"$arrayElemAt": []string{"$days", "$$index"}}}},
						},
					},
				},
			},
		},
	}})
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDEPARTMENT, "departmentId", "uniqueId", "ref.departmentId", "ref.departmentId")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONBRANCH, "branchId", "uniqueId", "ref.branchId", "ref.branchId")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDESIGNATION, "designationId", "uniqueId", "ref.designationId", "ref.designationId")...)

	//Adding pagination if necessary

	if pagination != nil {
		// mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		// //Getting Total count
		// totalCount, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEE).CountDocuments(ctx.CTX, func() bson.M {
		// 	if query != nil {
		// 		if len(query) > 0 {
		// 			return bson.M{"$and": query}
		// 		}
		// 	}
		// 	return bson.M{}
		// }())
		// if err != nil {
		// 	log.Println("Error in getting pagination")
		// }
		// fmt.Println("count", totalCount)
		// pagination.Count = int(totalCount)
		// d.Shared.PaginationData(pagination)
		paginationPipeline := []bson.M{}
		paginationPipeline = append(paginationPipeline, mainPipeline...)
		paginationPipeline = append(paginationPipeline, bson.M{"$group": bson.M{"_id": nil, "count": bson.M{"$sum": 1}}})
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		d.Shared.BsonToJSONPrintTag("Commodity Pagination quary =>", paginationPipeline)

		countcursor, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEE).Aggregate(ctx.CTX, paginationPipeline, nil)
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

	//Aggregation
	d.Shared.BsonToJSONPrint(mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}

	if err = cursor.All(context.TODO(), &attendance); err != nil {
		return nil, err
	}
	fmt.Println("len attendance", len(attendance))
	return attendance, nil
}
func (d *Daos) GetActiveEmployee(ctx *models.Context) ([]models.Employee, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	query = append(query, bson.M{"status": bson.M{"$nin": []string{constants.EMPLOYEESTATUSDELETED, constants.EMPLOYEESTATUSREJECT, constants.EMPLOYEESTATUSRELIEVE, constants.EMPLOYEESTATUSOFFBOARD, constants.EMPLOYEESTATUSONBORADING}}})

	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	//mainPipeline = append(mainPipeline, bson.M{"$project":bson.M{"uniqueId":}})
	//Aggregation
	d.Shared.BsonToJSONPrintTag("activestate query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var Employees []models.Employee
	if err = cursor.All(context.TODO(), &Employees); err != nil {
		return nil, err
	}
	return Employees, nil
}
func (d *Daos) EmployeePayrollWithEarningDeduction(ctx *models.Context, filter *models.FilterEmployeeSalary) (*models.EmployeePayrollWithEmployee, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter.Employee != "" {
		query = append(query, bson.M{"uniqueId": filter.Employee})
	}
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"from": constants.COLLECTIONEMPLOYEEPAYROLL,
			"as":   "payRoll",
			"let":  bson.M{"employeeId": "$uniqueId"},
			"pipeline": []bson.M{
				{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
					{"$eq": []string{"$status", constants.EMPLOYEEPAYROLLSTATUSACTIVE}},
					{"$eq": []string{"$employeeId", "$$employeeId"}},
				}}}},
			},
		},
	})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"payRoll": bson.M{"$arrayElemAt": []interface{}{"$payRoll", 0}}}})
	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"from": constants.COLLECTIONEMPLOYEEEARNING,
			"as":   "earning",
			"let":  bson.M{"employeeId": "$uniqueId"},
			"pipeline": []bson.M{
				{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
					{"$eq": []string{"$status", constants.EMPLOYEEPAYROLLSTATUSACTIVE}},
					{"$eq": []string{"$employeeId", "$$employeeId"}},
				}}}},
				{"$lookup": bson.M{
					"from": constants.COLLECTIONEMPLOYEEEARNINGMASTER,
					"as":   "name",
					"let":  bson.M{"masterId": "$earningId"},
					"pipeline": []bson.M{
						{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
							{"$eq": []string{"$status", constants.EMPLOYEEPAYROLLSTATUSACTIVE}},
							{"$eq": []string{"$uniqueId", "$$masterId"}},
						}}}},
					},
				}},
				{"$addFields": bson.M{"name": bson.M{"$arrayElemAt": []interface{}{"$name.title", 0}}}},
			},
		},
	})
	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"from": constants.COLLECTIONEMPLOYEEDEDUCTION,
			"as":   "deduction",
			"let":  bson.M{"employeeId": "$uniqueId"},
			"pipeline": []bson.M{
				{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
					{"$eq": []string{"$status", constants.EMPLOYEEPAYROLLSTATUSACTIVE}},
					{"$eq": []string{"$employeeId", "$$employeeId"}},
				}}}},
				{"$lookup": bson.M{
					"from": constants.COLLECTIONEMPLOYEEDEDUCTIONMASTER,
					"as":   "name",
					"let":  bson.M{"masterId": "$deductionId"},
					"pipeline": []bson.M{
						{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
							{"$eq": []string{"$status", constants.EMPLOYEEPAYROLLSTATUSACTIVE}},
							{"$eq": []string{"$uniqueId", "$$masterId"}},
						}}}},
					},
				}},
				{"$addFields": bson.M{"name": bson.M{"$arrayElemAt": []interface{}{"$name.title", 0}}}},
			},
		},
	})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"employeeId": "$uniqueId"}})
	d.Shared.BsonToJSONPrintTag("EmployeePayrollWithEarningDeduction", mainPipeline)
	var Employees []models.EmployeePayrollWithEmployee
	var Employee *models.EmployeePayrollWithEmployee
	cursor, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}

	if err = cursor.All(context.TODO(), &Employees); err != nil {
		return nil, err
	}
	if len(Employees) > 0 {
		Employee = &Employees[0]
	}
	return Employee, nil
}
func (d *Daos) UpdateEmployeeProfileImage(ctx *models.Context, Employee *models.UpdateBioData, UniqueID string) error {
	selector := bson.M{"uniqueId": UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": bson.M{"profileImg": Employee.ProfileImg}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEE).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}
func (d *Daos) GetsingleEmployeeWithMobileNumber(ctx *models.Context, UniqueId string) (*models.RefEmployee, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"mobile": UniqueId}})
	//LookUp
	d.Shared.BsonToJSONPrintTag("Employee query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var Employees []models.RefEmployee
	var Employee *models.RefEmployee
	if err = cursor.All(ctx.CTX, &Employees); err != nil {
		return nil, err
	}
	if len(Employees) > 0 {
		Employee = &Employees[0]
	}
	return Employee, nil
}
func (d *Daos) DashboardEmployeeCount(ctx *models.Context, employeefilter *models.FilterEmployee) ([]models.DashboardEmployeeCount, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if employeefilter != nil {

		if len(employeefilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": employeefilter.Status}})
		}
		if len(employeefilter.OrganisationID) > 0 {
			query = append(query, bson.M{"organisationId": bson.M{"$in": employeefilter.OrganisationID}})
		}
		if len(employeefilter.DepartmentID) > 0 {
			query = append(query, bson.M{"departmentId": bson.M{"$in": employeefilter.DepartmentID}})
		}
		if len(employeefilter.BranchID) > 0 {
			query = append(query, bson.M{"branchId": bson.M{"$in": employeefilter.BranchID}})
		}
		if len(employeefilter.DesignationID) > 0 {
			query = append(query, bson.M{"designationId": bson.M{"$in": employeefilter.DesignationID}})
		}
		if len(employeefilter.OmitStatus) > 0 {
			query = append(query, bson.M{"status": bson.M{"$nin": employeefilter.OmitStatus}})
		}
		// if employeefilter.Manager != "" {
		// 	query = append(query, bson.M{"lineManager": employeefilter.Manager})
		// }
		//Regex
		if employeefilter.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: employeefilter.Regex.Name, Options: "xi"}})
		}

	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if employeefilter != nil {
		if employeefilter.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{employeefilter.SortBy: employeefilter.SortOrder}})
		}
	}

	mainPipeline = append(mainPipeline, bson.M{
		"$group": bson.M{
			"_id":   nil,
			"count": bson.M{"$sum": 1},
		},
	})

	//LookUp

	//Aggregation
	d.Shared.BsonToJSONPrintTag("Employee query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var Employees []models.DashboardEmployeeCount
	if err = cursor.All(context.TODO(), &Employees); err != nil {
		return nil, err
	}
	return Employees, nil
}
func (d *Daos) GetChildEmployee(ctx *models.Context, uniqueID string) ([]models.LineManagerEmployee, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"lineManager": uniqueID}})
	//LookUp
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)
	//Aggregation
	//mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"employee": "$$ROOT"}})
	mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{
		"from": "employee",
		"as":   "child",
		"let":  bson.M{"employeeId": "$uniqueId"},
		"pipeline": []bson.M{
			{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
				//{"$eq":["$status","Active"]},
				{"$eq": []string{"$lineManager", "$$employeeId"}},
			}}}},
			{"$group": bson.M{
				"_id":   nil,
				"count": bson.M{"$sum": 1},
			}},
		}}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"child": bson.M{"$arrayElemAt": []interface{}{"$child.count", 0}}}})

	d.Shared.BsonToJSONPrintTag("Employee query =>", mainPipeline)

	cursor, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var employee []models.LineManagerEmployee
	//var emaillog *models.RefEmailLog
	if err = cursor.All(ctx.CTX, &employee); err != nil {
		return nil, err
	}

	return employee, nil
}
func (d *Daos) GetParetentEmployee(ctx *models.Context) (*models.EmployeeTree, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"lineManager": ""}})
	//LookUp
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)
	mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{
		"from": "employee",
		"as":   "child",
		"let":  bson.M{"employeeId": "$uniqueId"},
		"pipeline": []bson.M{
			{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
				//{"$eq":["$status","Active"]},
				{"$eq": []string{"$lineManager", "$$employeeId"}},
			}}}},
			{"$group": bson.M{
				"_id":   nil,
				"count": bson.M{"$sum": 1},
			}},
		}}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"child": bson.M{"$arrayElemAt": []interface{}{"$child.count", 0}}}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"employee": "$$ROOT"}})

	//Aggregation
	d.Shared.BsonToJSONPrintTag("Employee query =>", mainPipeline)

	cursor, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var employees []models.EmployeeTree
	fmt.Println("employee=====>,", employees)
	var employee *models.EmployeeTree
	if err = cursor.All(ctx.CTX, &employees); err != nil {
		return nil, err
	}
	if len(employees) > 0 {
		employee = &employees[0]
	}
	return employee, nil
}
func (d *Daos) GetLineManagerEmployee(ctx *models.Context, uniqueID string) ([]models.LineManagerEmployee, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"lineManager": uniqueID}})

	mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{
		"from": "employee",
		"as":   "child",
		"let":  bson.M{"employeeId": "$uniqueId"},
		"pipeline": []bson.M{
			{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
				//{"$eq":["$status","Active"]},
				{"$eq": []string{"$lineManager", "$$employeeId"}},
			}}}},
			{"$group": bson.M{
				"_id":   nil,
				"count": bson.M{"$sum": 1},
			}},
		}}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"child": bson.M{"$arrayElemAt": []interface{}{"$child.count", 0}}}})

	d.Shared.BsonToJSONPrintTag("Employee query =>", mainPipeline)

	cursor, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var employee []models.LineManagerEmployee
	//var emaillog *models.RefEmailLog
	if err = cursor.All(ctx.CTX, &employee); err != nil {
		return nil, err
	}
	return employee, nil
}
func (d *Daos) GetChildEmployeeWithtree(ctx *models.Context, uniqueID string) ([]models.EmployeeTreev2, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"lineManager": uniqueID}})
	//LookUp
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)
	//Aggregation
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"employee": "$$ROOT"}})
	mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{
		"from": "employee",
		"as":   "child",
		"let":  bson.M{"employeeId": "$uniqueId"},
		"pipeline": []bson.M{
			{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
				//{"$eq":["$status","Active"]},
				{"$eq": []string{"$lineManager", "$$employeeId"}},
			}}}},
			{"$group": bson.M{
				"_id":   nil,
				"count": bson.M{"$sum": 1},
			}},
		}}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"child": bson.M{"$arrayElemAt": []interface{}{"$child.count", 0}}}})

	d.Shared.BsonToJSONPrintTag("Employee query =>", mainPipeline)

	cursor, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var employee []models.EmployeeTreev2
	//var emaillog *models.RefEmailLog
	if err = cursor.All(ctx.CTX, &employee); err != nil {
		return nil, err
	}

	return employee, nil
}
func (d *Daos) GetallEmployee(ctx *models.Context) ([]models.AllEmployees, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	query = append(query, bson.M{"status": bson.M{"$nin": []string{constants.EMPLOYEESTATUSDELETED, constants.EMPLOYEESTATUSREJECT, constants.EMPLOYEESTATUSRELIEVE, constants.EMPLOYEESTATUSOFFBOARD, constants.EMPLOYEESTATUSONBORADING}}})
	//query = append(query, bson.M{"$sort": bson.M{"name": 1}})

	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{
		"from": "employee",
		"as":   "employee",
		"let":  bson.M{"employeeId": "$uniqueId"},
		"pipeline": []bson.M{
			{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
				//{"$eq":["$status","Active"]},
				{"$eq": []string{"$uniqueId", "$$employeeId"}},
			}}}},
		}}})
	mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"name": 1}})

	//LookUp
	//Aggregation
	d.Shared.BsonToJSONPrintTag("Employee query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var Employees []models.AllEmployees
	if err = cursor.All(context.TODO(), &Employees); err != nil {
		return nil, err
	}
	return Employees, nil
}
func (d *Daos) GetEmployeeLinemanagerCheck(ctx *models.Context, uniqueID string) (*models.EmployeeLinemanagerCheck, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"lineManager": uniqueID}})
	//LookUp
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONEMPLOYEE, uniqueID, "uniqueId", "ref.employee", "ref.employee")...)
	//Aggregation

	d.Shared.BsonToJSONPrintTag("Employee query =>", mainPipeline)

	cursor, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	linemanager := new(models.EmployeeLinemanagerCheck)
	var employee []models.Employee
	//var emaillog *models.RefEmailLog
	if err = cursor.All(ctx.CTX, &employee); err != nil {
		return nil, err
	}
	fmt.Println("Length of employee=====>", len(employee))
	if len(employee) > 0 {
		linemanager.LineManager = true
	} else {
		linemanager.LineManager = false
	}
	return linemanager, nil
}

//GetSingleEmployeeActiveWithName : ""
func (d *Daos) GetSingleEmployeeActiveWithName(ctx *models.Context, uniqueID string) (*models.RefEmployee, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"name": uniqueID, "status": constants.EMPLOYEESTATUSACTIVEEMPLOYEE}})

	d.Shared.BsonToJSONPrintTag("Employee query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var Employees []models.RefEmployee
	var Employee *models.RefEmployee
	if err = cursor.All(ctx.CTX, &Employees); err != nil {
		return nil, err
	}
	if len(Employees) > 0 {
		Employee = &Employees[0]
	}
	return Employee, nil
}

//GetBrithdayEmployee
func (d *Daos) GetBrithdayEmployees(ctx *models.Context, employeefilter *models.FilterEmployee, date *time.Time) ([]models.WeekCalEmployee, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{
		"$addFields": bson.M{
			"dobMonth": bson.M{"$month": "$dob"},
			"dobDay":   bson.M{"$dayOfMonth": "$dob"},
		},
	})
	query := []bson.M{}
	query = append(query, bson.M{"status": bson.M{"$nin": []string{constants.EMPLOYEESTATUSDELETED, constants.EMPLOYEESTATUSREJECT, constants.EMPLOYEESTATUSRELIEVE, constants.EMPLOYEESTATUSOFFBOARD, constants.EMPLOYEESTATUSONBORADING}}})

	if employeefilter != nil {

		if len(employeefilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": employeefilter.Status}})
		}
		if len(employeefilter.OrganisationID) > 0 {
			query = append(query, bson.M{"organisationId": bson.M{"$in": employeefilter.OrganisationID}})
		}
		if len(employeefilter.DepartmentID) > 0 {
			query = append(query, bson.M{"departmentId": bson.M{"$in": employeefilter.DepartmentID}})
		}
		if len(employeefilter.BranchID) > 0 {
			query = append(query, bson.M{"branchId": bson.M{"$in": employeefilter.BranchID}})
		}
		if len(employeefilter.DesignationID) > 0 {
			query = append(query, bson.M{"designationId": bson.M{"$in": employeefilter.DesignationID}})
		}
		if len(employeefilter.OmitStatus) > 0 {
			query = append(query, bson.M{"status": bson.M{"$nin": employeefilter.OmitStatus}})
		}
		// if employeefilter.Manager != "" {
		// 	query = append(query, bson.M{"lineManager": employeefilter.Manager})
		// }
		//Regex
		if employeefilter.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: employeefilter.Regex.Name, Options: "xi"}})
		}

	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	mainPipeline = append(mainPipeline, bson.M{
		"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
			{"$eq": []interface{}{"$dobMonth", int(date.Month())}},
			{"$eq": []interface{}{"$dobDay", date.Day()}},
		}}},
	})
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDESIGNATION, "designationId", "uniqueId", "ref.designationId", "ref.designationId")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("Employee query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var Employees []models.WeekCalEmployee
	if err = cursor.All(context.TODO(), &Employees); err != nil {
		return nil, err
	}
	return Employees, nil
}

//GetAnnivesaryEmployee
func (d *Daos) GetAnnivesaryEmployee(ctx *models.Context, employeefilter *models.FilterEmployee, date *time.Time) ([]models.WeekCalEmployee, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{
		"$addFields": bson.M{
			"joinMonth": bson.M{"$month": "$joiningDate"},
			"joinDay":   bson.M{"$dayOfMonth": "$joiningDate"},
			"joinyear":  bson.M{"$year": "$joiningDate"},
			//"year":bson.M{"$add":}
		},
	})
	query := []bson.M{}
	query = append(query, bson.M{"status": bson.M{"$nin": []string{constants.EMPLOYEESTATUSDELETED, constants.EMPLOYEESTATUSREJECT, constants.EMPLOYEESTATUSRELIEVE, constants.EMPLOYEESTATUSOFFBOARD, constants.EMPLOYEESTATUSONBORADING}}})

	if employeefilter != nil {

		if len(employeefilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": employeefilter.Status}})
		}
		if len(employeefilter.OrganisationID) > 0 {
			query = append(query, bson.M{"organisationId": bson.M{"$in": employeefilter.OrganisationID}})
		}
		if len(employeefilter.DepartmentID) > 0 {
			query = append(query, bson.M{"departmentId": bson.M{"$in": employeefilter.DepartmentID}})
		}
		if len(employeefilter.BranchID) > 0 {
			query = append(query, bson.M{"branchId": bson.M{"$in": employeefilter.BranchID}})
		}
		if len(employeefilter.DesignationID) > 0 {
			query = append(query, bson.M{"designationId": bson.M{"$in": employeefilter.DesignationID}})
		}
		if len(employeefilter.OmitStatus) > 0 {
			query = append(query, bson.M{"status": bson.M{"$nin": employeefilter.OmitStatus}})
		}
		// if employeefilter.Manager != "" {
		// 	query = append(query, bson.M{"lineManager": employeefilter.Manager})
		// }
		//Regex
		if employeefilter.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: employeefilter.Regex.Name, Options: "xi"}})
		}

	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	mainPipeline = append(mainPipeline, bson.M{
		"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
			{"$eq": []interface{}{"$joinMonth", int(date.Month())}},
			{"$eq": []interface{}{"$joinDay", date.Day()}},
		}}},
	})
	t := time.Now()
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"yearOfJoining": bson.M{"$subtract": []interface{}{t.Year(), "$joinyear"}}}})
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDESIGNATION, "designationId", "uniqueId", "ref.designationId", "ref.designationId")...)
	//Aggregation
	d.Shared.BsonToJSONPrintTag("Employee query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var Employees []models.WeekCalEmployee
	if err = cursor.All(context.TODO(), &Employees); err != nil {
		return nil, err
	}
	return Employees, nil
}
func (d *Daos) GetSingleEmployeeWithUserName(ctx *models.Context, uniqueID string) (*models.RefEmployee, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"userName": uniqueID}})
	//LookUp

	d.Shared.BsonToJSONPrintTag("Employee query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var Employees []models.RefEmployee
	var Employee *models.RefEmployee
	if err = cursor.All(ctx.CTX, &Employees); err != nil {
		return nil, err
	}
	if len(Employees) > 0 {
		Employee = &Employees[0]
	}
	return Employee, nil
}
func (d *Daos) UpdateEmployeeLoginId(ctx *models.Context, username string, loginid string) error {
	selector := bson.M{"userName": username}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": bson.M{"loginId": loginid}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEE).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}
func (d *Daos) GetSingleEmployeeWithLoginId(ctx *models.Context, uniqueID string) (*models.RefEmployee, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"loginId": uniqueID}})
	//LookUp

	d.Shared.BsonToJSONPrintTag("Employee query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var Employees []models.RefEmployee
	var Employee *models.RefEmployee
	if err = cursor.All(ctx.CTX, &Employees); err != nil {
		return nil, err
	}
	if len(Employees) > 0 {
		Employee = &Employees[0]
	}
	return Employee, nil
}
