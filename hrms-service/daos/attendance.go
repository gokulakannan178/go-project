package daos

import (
	"context"
	"errors"
	"fmt"
	"hrms-services/constants"
	"hrms-services/models"
	"log"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// SaveAttendance : ""
func (d *Daos) SaveAttendance(ctx *models.Context, attendance *models.Attendance) error {
	d.Shared.BsonToJSONPrint(attendance)
	_, err := ctx.DB.Collection(constants.COLLECTIONATTENDANCE).InsertOne(ctx.CTX, attendance)
	return err
}
func (d *Daos) SaveAttendanceWithUpsert(ctx *models.Context, attendance *models.Attendance) error {
	opts := options.Update().SetUpsert(true)
	updateQuery := bson.M{"uniqueId": attendance.UniqueID, "employeeId": attendance.EmployeeId}

	updateData := bson.M{"$set": attendance}
	if _, err := ctx.DB.Collection(constants.COLLECTIONATTENDANCE).UpdateOne(ctx.CTX, updateQuery, updateData, opts); err != nil {
		return errors.New("Error in updating log - " + err.Error())
	}

	return nil
}
func (d *Daos) SaveAttendanceWithUpsertUnPalanned(ctx *models.Context, attendance *models.Attendance) error {
	opts := options.Update().SetUpsert(true)
	updateQuery := bson.M{"uniqueId": attendance.UniqueID, "employeeId": attendance.EmployeeId, "loginMode": constants.ATTENDANCESTATUSAUTO}
	updateData := bson.M{"$set": attendance}
	if _, err := ctx.DB.Collection(constants.COLLECTIONATTENDANCE).UpdateOne(ctx.CTX, updateQuery, updateData, opts); err != nil {
		return errors.New("Error in updating log - " + err.Error())
	}

	return nil
}

// GetSingleAttendance : ""
func (d *Daos) GetSingleAttendance(ctx *models.Context, uniqueID string) (*models.RefAttendance, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONATTENDANCE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var attendances []models.RefAttendance
	var attendance *models.RefAttendance
	if err = cursor.All(ctx.CTX, &attendances); err != nil {
		return nil, err
	}
	if len(attendances) > 0 {
		attendance = &attendances[0]
	}
	return attendance, err
}

// GetSingleAttendanceByEmployeeId : ""
func (d *Daos) GetSingleAttendanceByEmployeeId(ctx *models.Context, uniqueID string, employeeId string) (*models.RefAttendance, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID, "employeeId": employeeId}})
	//Aggregation

	//d.Shared.BsonToJSONPrintTag("Attendance query =>", mainPipeline)

	cursor, err := ctx.DB.Collection(constants.COLLECTIONATTENDANCE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var attendances []models.RefAttendance
	var attendance *models.RefAttendance
	if err = cursor.All(ctx.CTX, &attendances); err != nil {
		return nil, err
	}
	if len(attendances) > 0 {
		attendance = &attendances[0]
	}
	return attendance, err
}
func (d *Daos) GetSingleAttendanceByEmployeeIdWithLoginMode(ctx *models.Context, uniqueID string, employeeId string, loginmode string) (*models.RefAttendance, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID, "employeeId": employeeId, "loginMode": loginmode}})
	//Aggregation

	d.Shared.BsonToJSONPrintTag("Attendance query =>", mainPipeline)

	cursor, err := ctx.DB.Collection(constants.COLLECTIONATTENDANCE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var attendances []models.RefAttendance
	var attendance *models.RefAttendance
	if err = cursor.All(ctx.CTX, &attendances); err != nil {
		return nil, err
	}
	if len(attendances) > 0 {
		attendance = &attendances[0]
	}
	return attendance, err
}

// GetSingleEmployeeAttendanceTodayStatus : ""
func (d *Daos) GetSingleEmployeeAttendanceTodayStatus(ctx *models.Context, uniqueID string, employeeId string) (*models.EmployeeAttendanceTodayStatus, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID, "employeeId": employeeId}})
	//Aggregation

	//d.Shared.BsonToJSONPrintTag("Attendance query =>", mainPipeline)

	cursor, err := ctx.DB.Collection(constants.COLLECTIONATTENDANCE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var attendances []models.EmployeeAttendanceTodayStatus
	var attendance *models.EmployeeAttendanceTodayStatus
	if err = cursor.All(ctx.CTX, &attendances); err != nil {
		return nil, err
	}
	if len(attendances) > 0 {
		attendance = &attendances[0]
	}
	return attendance, err
}

// GetSingleAttendanceByEmployeeIdAndStateValue : ""
func (d *Daos) GetSingleAttendanceByEmployeeIdAndStateValue(ctx *models.Context, uniqueID string, employeeId string, loginMode string) (*models.RefAttendance, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID, "employeeId": employeeId, "loginMode": loginMode}})
	//Aggregation

	d.Shared.BsonToJSONPrintTag("Attendance query =>", mainPipeline)

	cursor, err := ctx.DB.Collection(constants.COLLECTIONATTENDANCE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var attendances []models.RefAttendance
	var attendance *models.RefAttendance
	if err = cursor.All(ctx.CTX, &attendances); err != nil {
		return nil, err
	}
	if len(attendances) > 0 {
		attendance = &attendances[0]
	}
	return attendance, err
}

//UpdateAttendance : ""
func (d *Daos) UpdateAttendance(ctx *models.Context, attendance *models.Attendance) error {
	selector := bson.M{"uniqueId": attendance.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": attendance}
	_, err := ctx.DB.Collection(constants.COLLECTIONATTENDANCE).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// EnableAttendance : ""
func (d *Daos) EnableAttendance(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.ATTENDANCESTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONATTENDANCE).UpdateOne(ctx.CTX, selector, data)
	return err
}

// DisableAttendance : ""
func (d *Daos) DisableAttendance(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.ATTENDANCESTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONATTENDANCE).UpdateOne(ctx.CTX, selector, data)
	return err
}

// FilterAttendance : ""
func (d *Daos) FilterAttendance(ctx *models.Context, attendance *models.FilterAttendance, pagination *models.Pagination) ([]models.RefAttendance, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if attendance != nil {
		if len(attendance.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": attendance.Status}})
		}
		if len(attendance.EmployeeId) > 0 {
			query = append(query, bson.M{"employeeId": bson.M{"$in": attendance.EmployeeId}})

		}
		// if ft.DateRange.From != nil {
		// 	t := *ft.DateRange.From
		// 	FromDate := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
		// 	ToDate := time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, t.Location())
		// 	if ft.DateRange.To != nil {
		// 		t2 := *ft.DateRange.To
		// 		ToDate = time.Date(t2.Year(), t2.Month(), t2.Day(), 23, 59, 59, 0, t2.Location())
		// 	}
		// 	query = append(query, bson.M{"on": bson.M{"$gte": FromDate, "$lte": ToDate}})

		// }
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if attendance != nil {
		if attendance.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{attendance.SortBy: attendance.SortOrder}})
		}
	}
	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONATTENDANCE).CountDocuments(ctx.CTX, func() bson.M {
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
	cursor, err := ctx.DB.Collection(constants.COLLECTIONATTENDANCE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var attendanceFilter []models.RefAttendance
	if err = cursor.All(context.TODO(), &attendanceFilter); err != nil {
		return nil, err
	}
	return attendanceFilter, nil
}

// ClockinAttendance : ""
func (d *Daos) ClockinAttendance(ctx *models.Context, attendance *models.Attendance) error {
	opts := options.Update().SetUpsert(true)
	currentTime := time.Now()
	Day := currentTime.Day()
	Month := currentTime.Month()
	Year := currentTime.Year()
	strDay := strconv.Itoa(Day)
	strMonth := Month.String()
	strYear := strconv.Itoa(Year)

	uniqueIdcurrentformat := strDay + strMonth + strYear
	attendance.UniqueID = uniqueIdcurrentformat
	fmt.Println("attendance unique id ==>", attendance.UniqueID)
	updateQuery := bson.M{"uniqueId": uniqueIdcurrentformat, "employeeId": attendance.EmployeeId}
	fmt.Println("update Query ==>", updateQuery)
	updateData := bson.M{"$set": attendance}
	if _, err := ctx.DB.Collection(constants.COLLECTIONATTENDANCE).UpdateOne(ctx.CTX, updateQuery, updateData, opts); err != nil {
		return errors.New("Error in updating log - " + err.Error())
	}
	return nil

}
func (d *Daos) ClockinAttendancev2(ctx *models.Context, attendance *models.Attendance) error {
	opts := options.Update().SetUpsert(true)
	currentTime := time.Now()
	Day := currentTime.Day()
	Month := currentTime.Month()
	Year := currentTime.Year()
	strDay := strconv.Itoa(Day)
	strMonth := Month.String()
	strYear := strconv.Itoa(Year)

	uniqueIdcurrentformat := strDay + strMonth + strYear
	updateQuery := bson.M{"uniqueId": uniqueIdcurrentformat, "employeeId": attendance.EmployeeId}
	d.Shared.BsonToJSONPrintTag("clockin query =>", updateQuery)
	updateData := bson.M{"$set": attendance}
	_, err := ctx.DB.Collection(constants.COLLECTIONATTENDANCE).UpdateOne(ctx.CTX, updateQuery, updateData, opts)
	return err
}

// ClockoutAttendance : ""
func (d *Daos) ClockoutAttendance(ctx *models.Context, attendance *models.Attendance) error {
	selector := bson.M{"uniqueId": attendance.UniqueID, "employeeId": attendance.EmployeeId}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": attendance}
	_, err := ctx.DB.Collection(constants.COLLECTIONATTENDANCE).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}
func (d *Daos) DayWiseAttendanceReport(ctx *models.Context, filter *models.DayWiseAttendanceReportFilter) (*models.DayWiseAttendanceReport, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	loc, _ := time.LoadLocation("Asia/Kolkata")
	t := time.Now().In(loc)
	var sd, ed time.Time
	//	var attendances *models.DayWiseAttendanceReport
	attendance := new(models.DayWiseAttendanceReport)
	if filter != nil {
		if filter.EmployeeId != "" {
			query = append(query, bson.M{"employeeId": filter.EmployeeId})
		}
		if filter.Status != "" {
			query = append(query, bson.M{"status": filter.Status})
		}
		if len(filter.Manager) > 0 {
			query = append(query, bson.M{"lineManager": bson.M{"$in": filter.Manager}})
		}
		if filter.StartDate != nil {
			sd = time.Date(filter.StartDate.Year(), filter.StartDate.Month(), filter.StartDate.Day(), 0, 0, 0, 0, filter.StartDate.Location())
			if t.Month() == filter.StartDate.Month() && t.Year() == filter.StartDate.Year() {
				ed = time.Date(filter.StartDate.Year(), filter.StartDate.Month(), t.Day(), 23, 59, 59, 999999999, filter.StartDate.Location())
			} else {
				ed = time.Date(filter.StartDate.Year(), filter.StartDate.Month()+1, 0, 23, 59, 59, 999999999, filter.StartDate.Location())
			}
			query = append(query, bson.M{"date": bson.M{"$gte": sd, "$lte": ed}})
			fmt.Println("sd===>", sd)
			fmt.Println("ed===>", ed)

		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"dom": bson.M{"$dayOfMonth": "$date"}}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"month": bson.M{"$month": "$date"}}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"year": bson.M{"$year": "$date"}}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"dayOfWeek": bson.M{"$dayOfWeek": "$date"}}})
	//fmt.Println("dayofweek==>", attendance)
	mainPipeline = append(mainPipeline, bson.M{"$group": bson.M{"_id": nil, "days": bson.M{"$push": "$$ROOT"}}})
	mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"date": 1}})
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
						"vars": bson.M{"index": bson.M{"$indexOfArray": []string{"$days.dom", "$$rangeDay"}}},
						"in": bson.M{
							"$cond": bson.M{
								"if": bson.M{"$eq": []interface{}{"$$index", -1}},
								"then": bson.M{
									"date":      bson.M{"$dateFromParts": bson.M{"year": "$year1", "month": "$month1", "day": "$$rangeDay", "hour": 9}},
									"dayOfWeek": bson.M{"$dayOfWeek": bson.M{"$dateFromParts": bson.M{"year": "$year1", "month": "$month1", "day": "$$rangeDay", "hour": 9}}},
								},
								"else": bson.M{"$arrayElemAt": []string{"$days", "$$index"}}},
						},
					},
				},
			},
		},
	}})
	mainPipeline = append(mainPipeline, bson.M{"$unwind": "$days"})
	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"from":         constants.COLLECTIONDAYOFWEEK,
			"as":           "days.dayOfWeek",
			"localField":   "days.dayOfWeek",
			"foreignField": "dayOfWeek",
		},
	})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"days.dayOfWeek": bson.M{"$arrayElemAt": []interface{}{"$days.dayOfWeek.name", 0}}}})

	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONATTENDANCELOG, "attendanceId", "days.uniqueId", "days.attendanceLog", "days.attendanceLog")...)
	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"from": constants.COLLECTIONATTENDANCELOG,
			"as":   "days.attendanceLog",
			"let":  bson.M{"uniqueId": "$days.uniqueId", "employeeId": "$days.employeeId"},
			"pipeline": []bson.M{
				{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
					{"$eq": []string{"$attendanceId", "$$uniqueId"}},
					{"$eq": []string{"$employeeId", "$$employeeId"}},
				}}}},
			},
		},
	})
	mainPipeline = append(mainPipeline, bson.M{
		"$group": bson.M{
			"_id":              nil,
			"days":             bson.M{"$push": "$$ROOT.days"},
			"totalWorkingMins": bson.M{"$sum": "$days.totalWorkingMins"},
			"totalOvertime":    bson.M{"$sum": "$days.overtime"},
		},
	})
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDAYOFWEEK, "dayOfWeek", "dayOfWeek", "dayOfWeek", "dayOfWeek.name")...)

	//Aggregation
	d.Shared.BsonToJSONPrint(mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONATTENDANCE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var attendances []models.DayWiseAttendanceReport
	if err = cursor.All(ctx.CTX, &attendances); err != nil {
		return nil, err
	}
	if len(attendances) > 0 {
		attendance = &attendances[0]
	}
	return attendance, nil
}
func (d *Daos) WeaklyWiseAttendanceReport(ctx *models.Context, filter *models.DayWiseAttendanceReportFilter) (*models.DayWiseAttendanceReport, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	var sd, ed time.Time
	//	var attendances *models.DayWiseAttendanceReport
	attendance := new(models.DayWiseAttendanceReport)
	if filter != nil {
		query = append(query, bson.M{"status": constants.ATTENDANCESTATUSPENDING})

		if filter.EmployeeId != "" {
			query = append(query, bson.M{"employeeId": filter.EmployeeId})
		}
		if filter.StartDate != nil {
			days := filter.StartDate.Day()
			switch {
			case days <= 7:
				sd = time.Date(filter.StartDate.Year(), filter.StartDate.Month(), 1, 0, 0, 0, 0, filter.StartDate.Location())
				ed = time.Date(filter.StartDate.Year(), filter.StartDate.Month(), 7, 23, 59, 59, 999999999, filter.StartDate.Location())
			case days <= 14:
				sd = time.Date(filter.StartDate.Year(), filter.StartDate.Month(), 8, 0, 0, 0, 0, filter.StartDate.Location())
				ed = time.Date(filter.StartDate.Year(), filter.StartDate.Month(), 14, 23, 59, 59, 999999999, filter.StartDate.Location())
			case days <= 21:
				sd = time.Date(filter.StartDate.Year(), filter.StartDate.Month(), 15, 0, 0, 0, 0, filter.StartDate.Location())
				ed = time.Date(filter.StartDate.Year(), filter.StartDate.Month(), 21, 23, 59, 59, 999999999, filter.StartDate.Location())
			case days <= 28:
				sd = time.Date(filter.StartDate.Year(), filter.StartDate.Month(), 22, 0, 0, 0, 0, filter.StartDate.Location())
				ed = time.Date(filter.StartDate.Year(), filter.StartDate.Month(), 28, 23, 59, 59, 999999999, filter.StartDate.Location())
			default:
				sd = time.Date(filter.StartDate.Year(), filter.StartDate.Month(), 29, 0, 0, 0, 0, filter.StartDate.Location())
				ed = time.Date(filter.StartDate.Year(), filter.StartDate.Month()+1, 0, 23, 59, 59, 999999999, filter.StartDate.Location())

			}
			query = append(query, bson.M{"date": bson.M{"$gte": sd, "$lte": ed}})
			fmt.Println("sd===>", sd)
			fmt.Println("ed===>", ed)

		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"dom": bson.M{"$dayOfMonth": "$date"}}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"month": bson.M{"$month": "$date"}}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"year": bson.M{"$year": "$date"}}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"dayOfWeek": bson.M{"$dayOfWeek": "$date"}}})
	//fmt.Println("dayofweek==>", attendance)
	mainPipeline = append(mainPipeline, bson.M{"$group": bson.M{"_id": nil, "days": bson.M{"$push": "$$ROOT"}}})
	mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"date": 1}})
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
						"vars": bson.M{"index": bson.M{"$indexOfArray": []string{"$days.dom", "$$rangeDay"}}},
						"in": bson.M{
							"$cond": bson.M{
								"if": bson.M{"$eq": []interface{}{"$$index", -1}},
								"then": bson.M{
									"date":      bson.M{"$dateFromParts": bson.M{"year": "$year1", "month": "$month1", "day": "$$rangeDay", "hour": 9}},
									"dayOfWeek": bson.M{"$dayOfWeek": bson.M{"$dateFromParts": bson.M{"year": "$year1", "month": "$month1", "day": "$$rangeDay", "hour": 9}}},
								},
								"else": bson.M{"$arrayElemAt": []string{"$days", "$$index"}}},
						},
					},
				},
			},
		},
	}})
	mainPipeline = append(mainPipeline, bson.M{"$unwind": "$days"})
	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"from":         constants.COLLECTIONDAYOFWEEK,
			"as":           "days.dayOfWeek",
			"localField":   "days.dayOfWeek",
			"foreignField": "dayOfWeek",
		},
	})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"days.dayOfWeek": bson.M{"$arrayElemAt": []interface{}{"$days.dayOfWeek.name", 0}}}})

	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONATTENDANCELOG, "attendanceId", "days.uniqueId", "days.attendanceLog", "days.attendanceLog")...)
	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"from":         constants.COLLECTIONATTENDANCELOG,
			"as":           "days.attendanceLog",
			"localField":   "days.uniqueId",
			"foreignField": "attendanceId",
		},
	})
	mainPipeline = append(mainPipeline, bson.M{
		"$group": bson.M{
			"_id":              nil,
			"days":             bson.M{"$push": "$$ROOT.days"},
			"totalWorkingMins": bson.M{"$sum": "$days.totalWorkingMins"},
			"totalOvertime":    bson.M{"$sum": "$days.overtime"},
		},
	})
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDAYOFWEEK, "dayOfWeek", "dayOfWeek", "dayOfWeek", "dayOfWeek.name")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("WeakEmployessworkinghours =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONATTENDANCE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var attendances []models.DayWiseAttendanceReport
	if err = cursor.All(ctx.CTX, &attendances); err != nil {
		return nil, err
	}
	if len(attendances) > 0 {
		attendance = &attendances[0]
	}
	return attendance, nil

}
func (d *Daos) AttendanceEmployeeStatistics(ctx *models.Context, employeeId string, uniqueID string) (*models.AttendanceEmployeeStatistics, error) {

	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"employeeId": employeeId, "uniqueId": uniqueID}})

	d.Shared.BsonToJSONPrintTag("GetSingleAttendanceLoglast =>", mainPipeline)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONATTENDANCE).Aggregate(ctx.CTX, mainPipeline, nil)
	//fmt.Println("cursor value", cursor)
	if err != nil {
		return nil, err
	}
	var AttendanceLogs []models.AttendanceEmployeeStatistics
	var AttendanceLog *models.AttendanceEmployeeStatistics
	if err = cursor.All(ctx.CTX, &AttendanceLogs); err != nil {
		return nil, err
	}
	if len(AttendanceLogs) > 0 {
		AttendanceLog = &AttendanceLogs[0]
	}
	return AttendanceLog, nil
}
func (d *Daos) TodayEmployessLeave(ctx *models.Context, uniqueID string) (*models.TodayEmployessLeave, error) {
	currentTime := time.Now()
	Day := currentTime.Day()
	Month := currentTime.Month()
	Year := currentTime.Year()
	strDay := strconv.Itoa(Day)
	strMonth := Month.String()
	strYear := strconv.Itoa(Year)
	uniqueID = strDay + strMonth + strYear
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	mainPipeline = append(mainPipeline, bson.M{
		"$group": bson.M{"_id": bson.M{"caseLOP": "$caseLOP"},
			"count": bson.M{"$sum": 1},
		}})
	mainPipeline = append(mainPipeline, bson.M{
		"$project": bson.M{
			"k":   "$_id.caseLOP",
			"v":   "$count",
			"_id": 0,
		}})
	mainPipeline = append(mainPipeline, bson.M{"$group": bson.M{
		"_id":  nil,
		"data": bson.M{"$push": "$$ROOT"},
	}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"attendance": bson.M{"$arrayToObject": "$data"}}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"planned": "$attendance.Planned"}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"unPlanned": "$attendance.UnPlanned"}})

	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"from": constants.COLLECTIONEMPLOYEETIMEOFF,
			"as":   "pendingTimeOff",
			"let":  bson.M{"uniqueId": "$uniqueId"},
			"pipeline": []bson.M{
				{"$match": bson.M{"status": constants.EMPLOYEETIMEOFFSTATUSREQUEST}},
				{"$group": bson.M{
					"_id":   nil,
					"count": bson.M{"$sum": 1}}},
			},
		}})

	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"pendingTimeOff": bson.M{"$arrayElemAt": []interface{}{"$pendingTimeOff.count", 0}}}})
	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"from": constants.COLLECTIONATTENDANCE,
			"as":   "todayPresent",
			"let":  bson.M{"uniqueId": "$uniqueId"},
			"pipeline": []bson.M{
				{"$match": bson.M{"loginMode": bson.M{"$in": []string{constants.ATTENDANCESTATUSLOGIN, constants.ATTENDANCESTATUSLOGOUT}}, "uniqueId": uniqueID}},
				{"$group": bson.M{
					"_id":   nil,
					"count": bson.M{"$sum": 1}}},
			},
		}})

	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"todayPresent": bson.M{"$arrayElemAt": []interface{}{"$todayPresent.count", 0}}}})
	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"from": constants.COLLECTIONEMPLOYEE,
			"as":   "totalEmployee",
			"let":  bson.M{"uniqueId": "$uniqueId"},
			"pipeline": []bson.M{
				{"$match": bson.M{"status": bson.M{"$nin": []string{constants.EMPLOYEESTATUSREJECT, constants.EMPLOYEESTATUSONBORADING, constants.EMPLOYEESTATUSRELIEVE}}}},
				{"$group": bson.M{
					"_id":   nil,
					"count": bson.M{"$sum": 1}}},
			},
		}})

	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"totalEmployee": bson.M{"$arrayElemAt": []interface{}{"$totalEmployee.count", 0}}}})

	d.Shared.BsonToJSONPrintTag("TodayEmployessLeave =>", mainPipeline)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONATTENDANCE).Aggregate(ctx.CTX, mainPipeline, nil)
	//fmt.Println("cursor value", cursor)
	if err != nil {
		return nil, err
	}
	var AttendanceLogs []models.TodayEmployessLeave
	var AttendanceLog *models.TodayEmployessLeave
	if err = cursor.All(ctx.CTX, &AttendanceLogs); err != nil {
		return nil, err
	}
	if len(AttendanceLogs) > 0 {
		AttendanceLog = &AttendanceLogs[0]
	}
	return AttendanceLog, nil
}
func (d *Daos) EmployeeAttendanceParRoll(ctx *models.Context, attendance *models.FilterAttendance) (*models.PayRollEmployessLeave, error) {

	mainPipeline := []bson.M{}
	query := []bson.M{}
	var sd, ed time.Time
	if attendance != nil {
		if attendance.Employee != "" {
			query = append(query, bson.M{"uniqueId": attendance.Employee})
		}
		if len(attendance.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": attendance.Status}})
		}
	}
	if attendance.StartDate != nil {
		t := *attendance.StartDate
		sd = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
		ed = time.Date(t.Year(), t.Month()+1, 0, 23, 59, 59, 0, t.Location())
		//query = append(query, bson.M{"date": bson.M{"$gte": FromDate, "$lte": ToDate}})
	}
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	//	mainPipeline = append(mainPipeline, bson.M)
	//mainPipeline = append(mainPipeline, bson.M)
	//mainPipeline = append(mainPipeline, bson.M)
	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"from": constants.COLLECTIONATTENDANCE,
			"as":   "attendance",
			"let":  bson.M{"employeeId": "$uniqueId"},
			"pipeline": []bson.M{
				{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
					{"$eq": []string{"$employeeId", "$$employeeId"}},
					{"$gte": []interface{}{"$date", sd}},
					{"$lte": []interface{}{"$date", ed}},
				},
				}}},
				{"$group": bson.M{"_id": bson.M{"payRoll": "$payRoll"},
					"count": bson.M{"$sum": 1},
				}},
				{"$project": bson.M{
					"k":   "$_id.payRoll",
					"v":   "$count",
					"_id": 0,
				}},
				{"$group": bson.M{
					"_id":  nil,
					"data": bson.M{"$push": "$$ROOT"},
				}},
				{"$addFields": bson.M{"data": bson.M{"$arrayToObject": "$data"}}},
			},
		},
	})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"attendance": bson.M{"$arrayElemAt": []interface{}{"$attendance.data", 0}}}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"lossOfPay": "$attendance.LOP"}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"paid": "$attendance.Paid"}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"partialPay": "$attendance.PartialPay"}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"employeeId": "$uniqueId"}})

	d.Shared.BsonToJSONPrintTag("EmployeeParRoll =>", mainPipeline)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEE).Aggregate(ctx.CTX, mainPipeline, nil)
	//fmt.Println("cursor value", cursor)
	if err != nil {
		return nil, err
	}
	var payRolls []models.PayRollEmployessLeave
	var payRoll *models.PayRollEmployessLeave
	//	var AttendanceLog *models.TodayEmployessLeave
	if err = cursor.All(ctx.CTX, &payRolls); err != nil {
		return nil, err
	}
	if len(payRolls) > 0 {
		payRoll = &payRolls[0]
	}
	return payRoll, nil
}
func (d *Daos) EmployeeAttendanceApprove(ctx *models.Context, employeeid string, Attendance string) error {
	selector := bson.M{"employeeId": employeeid, "uniqueId": Attendance}
	data := bson.M{"$set": bson.M{"status": constants.ATTENDANCESTATUSAPPROVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONATTENDANCE).UpdateOne(ctx.CTX, selector, data)
	return err
}
func (d *Daos) EmployeeAttendanceReject(ctx *models.Context, employeeid string, Attendance string) error {
	selector := bson.M{"employeeId": employeeid, "uniqueId": Attendance}
	data := bson.M{"$set": bson.M{"status": constants.ATTENDANCESTATUSREJECT}}
	_, err := ctx.DB.Collection(constants.COLLECTIONATTENDANCE).UpdateOne(ctx.CTX, selector, data)
	return err
}
func (d *Daos) GetTodayEmployeeTimeOff(ctx *models.Context) ([]models.TodayTimeoff, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	t := time.Now()
	sd := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	ed := time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, t.Location())
	query = append(query, bson.M{"date": bson.M{"$gte": sd, "$lte": ed}})

	query = append(query, bson.M{"loginMode": constants.ATTENDANCESTATUSTIMEOFF})
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	//lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONEMPLOYEE, "employeeId", "uniqueId", "employeeId", "employeeId")...)
	d.Shared.BsonToJSONPrintTag("Today time off =>", mainPipeline)
	// mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{
	// "as":   "PendingLeaveRequest",
	// "from": "employeetimeoff",
	// "let":  bson.M{"uniqueId": "$employeeId"},
	// "pipeline": []bson.M{
	// 	{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
	// 		{"$eq": []string{"$status", "Request"}},
	// 		{"$gte": []interface{}{"$startDate", sd}},
	// 		{"$lte": []interface{}{"$endDate", ed}}}}}},
	// 	bson.M{"$lookup": bson.M{
	// 		"as":           "ref.employee",
	// 		"foreignField": "uniqueId",
	// 		"from":         "employee",
	// 		"localField":   "employeeId"}},
	// 	bson.M{"$addFields": bson.M{"ref.employee": bson.M{"$arrayElemAt": []interface{}{"$ref.employee", 0}}}},
	// }}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONATTENDANCE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var attendances []models.TodayTimeoff
	if err = cursor.All(ctx.CTX, &attendances); err != nil {
		return nil, err
	}
	return attendances, err
}
func (d *Daos) EmployeeAttendanceRejected(ctx *models.Context, Attendance *models.EmployeeAttendanceApprove) error {
	selector := bson.M{"employeeId": Attendance.EmployeeId, "uniqueId": bson.M{"$in": Attendance.UniqueID}}
	data := bson.M{"$set": bson.M{"status": constants.ATTENDANCESTATUSREJECT}}
	_, err := ctx.DB.Collection(constants.COLLECTIONATTENDANCE).UpdateOne(ctx.CTX, selector, data)
	return err
}
func (d *Daos) GetTodayEmployeePunchin(ctx *models.Context) ([]models.TodayPunchin, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	t := time.Now()
	sd := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	ed := time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, t.Location())
	query = append(query, bson.M{"date": bson.M{"$gte": sd, "$lte": ed}})

	query = append(query, bson.M{"loginMode": bson.M{"$in": []string{constants.ATTENDANCESTATUSLOGIN, constants.ATTENDANCESTATUSLOGOUT}}})
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	mainPipeline = append(mainPipeline, bson.M{
		"$addFields": bson.M{
			"firstpunchinTime": "$punchIn",
			"lastpunchoutTime": "$punchOut",
		},
	})
	//lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONEMPLOYEE, "employeeId", "uniqueId", "employeeId", "employeeId")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONATTENDANCE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var attendances []models.TodayPunchin
	if err = cursor.All(ctx.CTX, &attendances); err != nil {
		return nil, err
	}
	return attendances, err
}
func (d *Daos) GetTodayEmployeeUplannedLeave(ctx *models.Context) ([]models.TodayTimeoff, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	t := time.Now()
	sd := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	ed := time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, t.Location())
	query = append(query, bson.M{"date": bson.M{"$gte": sd, "$lte": ed}})

	query = append(query, bson.M{"caseLOP": constants.EMPLOYEELOPCASEUNPLANNED})
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	//lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONEMPLOYEE, "employeeId", "uniqueId", "employeeId", "employeeId")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONATTENDANCE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var attendances []models.TodayTimeoff
	if err = cursor.All(ctx.CTX, &attendances); err != nil {
		return nil, err
	}
	return attendances, err
}
