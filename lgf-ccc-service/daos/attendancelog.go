package daos

import (
	"context"
	"errors"
	"fmt"
	"lgf-ccc-service/constants"
	"lgf-ccc-service/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// SaveAttendanceLog : ""
func (d *Daos) SaveAttendanceLog(ctx *models.Context, attendanceLog *models.AttendanceLog) error {
	d.Shared.BsonToJSONPrint(attendanceLog)
	_, err := ctx.DB.Collection(constants.COLLECTIONATTENDANCELOG).InsertOne(ctx.CTX, attendanceLog)
	return err
}
func (d *Daos) SaveAttendanceLogWithUpsert(ctx *models.Context, attendance *models.AttendanceLog) error {
	opts := options.Update().SetUpsert(true)
	updateQuery := bson.M{"uniqueId": attendance.UniqueID, "employeeId": attendance.EmployeeId}
	updateData := bson.M{"$set": attendance}
	if _, err := ctx.DB.Collection(constants.COLLECTIONATTENDANCELOG).UpdateOne(ctx.CTX, updateQuery, updateData, opts); err != nil {
		return errors.New("Error in updating log - " + err.Error())
	}

	return nil
}

// GetSingleAttendanceLog : ""
func (d *Daos) GetSingleAttendanceLog(ctx *models.Context, UniqueID string) (*models.RefAttendanceLog, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONATTENDANCELOG).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var attendanceLogs []models.RefAttendanceLog
	var attendanceLog *models.RefAttendanceLog
	if err = cursor.All(ctx.CTX, &attendanceLogs); err != nil {
		return nil, err
	}
	if len(attendanceLogs) > 0 {
		attendanceLog = &attendanceLogs[0]
	}
	return attendanceLog, err
}

//GetSingleAttendanceLogEmployeeId : ""
func (d *Daos) GetSingleAttendanceLogEmployeeId(ctx *models.Context, employeeId string, uniqueID string) (*models.RefAttendanceLog, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"employeeId": employeeId, "uniqueId": uniqueID}})
	//mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"status": status}})

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONATTENDANCELOG).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var AttendanceLogs []models.RefAttendanceLog
	var AttendanceLog *models.RefAttendanceLog
	if err = cursor.All(ctx.CTX, &AttendanceLogs); err != nil {
		return nil, err
	}
	if len(AttendanceLogs) > 0 {
		AttendanceLog = &AttendanceLogs[0]
	}
	return AttendanceLog, nil
}

//GetSingleAttendanceLoglast : ""
func (d *Daos) GetSingleAttendanceLoglast(ctx *models.Context, employeeId string, uniqueID string) (*models.RefAttendanceLog, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"employeeId": employeeId, "attendanceId": uniqueID}})
	mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"_id": -1}})
	//mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"status": status}})
	//mainPipeline = append(mainPipeline, []bson.M{bson.M{"$limit": 1}."$sort": bson.M{"_id": -1}}...)

	d.Shared.BsonToJSONPrintTag("GetSingleAttendanceLoglast =>", mainPipeline)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONATTENDANCELOG).Aggregate(ctx.CTX, mainPipeline, nil)
	fmt.Println("cursor value", cursor)
	if err != nil {
		return nil, err
	}
	var AttendanceLogs []models.RefAttendanceLog
	var AttendanceLog *models.RefAttendanceLog
	if err = cursor.All(ctx.CTX, &AttendanceLogs); err != nil {
		return nil, err
	}
	if len(AttendanceLogs) > 0 {
		AttendanceLog = &AttendanceLogs[0]
	}
	fmt.Println("AttendanceLog", AttendanceLog)
	return AttendanceLog, nil
}

//AttendanceEmployeeTodayStatus : ""
func (d *Daos) AttendanceEmployeeTodayStatus(ctx *models.Context, employeeId string, uniqueID string) (*models.AttendanceEmployeeTodayStatus, error) {
	//var attlog models.AttendanceEmployeeTodayStatus
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"employeeId": employeeId, "attendanceId": uniqueID}})
	mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"date": 1}})
	group := []bson.M{{"$group": bson.M{"_id": nil, "firstpunchinTime": bson.M{"$first": "$punchinTime"}, "lastpunchoutTime": bson.M{"$last": "$punchoutTime"}, "totalworkingHours": bson.M{"$sum": "$workingHours"}, "totalbreakHours": bson.M{"$sum": "$breakHours"}}}}
	mainPipeline = append(mainPipeline, group...)

	d.Shared.BsonToJSONPrintTag("GetSingleAttendanceLoglast =>", mainPipeline)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONATTENDANCELOG).Aggregate(ctx.CTX, mainPipeline, nil)
	//fmt.Println("cursor value", cursor)
	if err != nil {
		return nil, err
	}
	var AttendanceLogs []models.AttendanceEmployeeTodayStatus
	var AttendanceLog *models.AttendanceEmployeeTodayStatus
	if err = cursor.All(ctx.CTX, &AttendanceLogs); err != nil {
		return nil, err
	}

	if len(AttendanceLogs) > 0 {
		AttendanceLog = &AttendanceLogs[0]
	}
	return AttendanceLog, nil
}
func (d *Daos) AttendanceEmployeeTodayLogs(ctx *models.Context, employeeId string, uniqueID string) ([]models.AttendanceLog, error) {
	//var attlog models.AttendanceEmployeeTodayStatus
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"employeeId": employeeId, "attendanceId": uniqueID}})
	mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"date": 1}})
	// group := []bson.M{{"$group": bson.M{"_id": nil, "firstpunchinTime": bson.M{"$first": "$punchinTime"}, "lastpunchoutTime": bson.M{"$last": "$punchoutTime"}, "totalworkingHours": bson.M{"$sum": "$workingHours"}, "totalbreakHours": bson.M{"$sum": "$breakHours"}}}}
	// mainPipeline = append(mainPipeline, group...)

	d.Shared.BsonToJSONPrintTag("GetSingleAttendanceLogs =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONATTENDANCELOG).Aggregate(ctx.CTX, mainPipeline, nil)
	//fmt.Println("cursor value", cursor)
	if err != nil {
		return nil, err
	}
	var AttendanceLogs []models.AttendanceLog
	if err = cursor.All(ctx.CTX, &AttendanceLogs); err != nil {
		return nil, err
	}
	return AttendanceLogs, nil
}

//AttendanceEmployeeTodaystatuswithouttotaltime : ""
func (d *Daos) AttendanceEmployeeTodaystatuswithouttotaltime(ctx *models.Context, employeeId string, uniqueID string) (*models.AttendanceEmployeeTodayStatus, error) {
	//var attlog models.AttendanceEmployeeTodayStatus
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"employeeId": employeeId, "uniqueId": uniqueID}})
	mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"date": 1}})
	group := []bson.M{{"$group": bson.M{"_id": nil, "firstpunchinTime": bson.M{"$first": "$punchinTime"}, "lastpunchoutTime": bson.M{"$last": "$punchoutTime"}}}}
	mainPipeline = append(mainPipeline, group...)

	d.Shared.BsonToJSONPrintTag("GetSingleAttendanceLoglast =>", mainPipeline)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONATTENDANCELOG).Aggregate(ctx.CTX, mainPipeline, nil)
	//fmt.Println("cursor value", cursor)
	if err != nil {
		return nil, err
	}
	var AttendanceLogs []models.AttendanceEmployeeTodayStatus
	var AttendanceLog *models.AttendanceEmployeeTodayStatus
	if err = cursor.All(ctx.CTX, &AttendanceLogs); err != nil {
		return nil, err
	}

	if len(AttendanceLogs) > 0 {
		AttendanceLog = &AttendanceLogs[0]
	}
	return AttendanceLog, nil
}

// GetSingleAttendanceLogByEmployeeIdAndState : ""
func (d *Daos) GetSingleLastAttendanceLogByEmployeeIdAndState(ctx *models.Context, uniqueID string, employeeId string, stateId string) (*models.RefAttendanceLog, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"attendanceId": uniqueID, "employeeId": employeeId, "loginMode": stateId}})
	mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"_id": -1}})
	//Aggregation
	d.Shared.BsonToJSONPrintTag("GetSingleLastAttendanceLogByEmployeeIdAndState =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONATTENDANCELOG).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}

	var attendanceLogs []models.RefAttendanceLog
	var attendanceLog *models.RefAttendanceLog
	if err = cursor.All(ctx.CTX, &attendanceLogs); err != nil {
		return nil, err
	}
	if len(attendanceLogs) > 0 {
		attendanceLog = &attendanceLogs[0]
	}
	return attendanceLog, err
}

//UpdateAttendanceLog : ""
func (d *Daos) UpdateAttendanceLog(ctx *models.Context, AttendanceLog *models.AttendanceLog) error {
	selector := bson.M{"uniqueId": AttendanceLog.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": AttendanceLog}
	_, err := ctx.DB.Collection(constants.COLLECTIONATTENDANCELOG).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// EnableAttendanceLog : ""
func (d *Daos) EnableAttendanceLog(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.ATTENDANCELOGSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONATTENDANCELOG).UpdateOne(ctx.CTX, selector, data)
	return err
}

// DisableAttendanceLog : ""
func (d *Daos) DisableAttendanceLog(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.ATTENDANCELOGSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONATTENDANCELOG).UpdateOne(ctx.CTX, selector, data)
	return err
}

// DeleteAttendanceLog : ""
func (d *Daos) DeleteAttendanceLog(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.ATTENDANCELOGSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONATTENDANCELOG).UpdateOne(ctx.CTX, selector, data)
	return err
}

// FilterAttendanceLog : ""
func (d *Daos) FilterAttendanceLog(ctx *models.Context, filter *models.FilterAttendanceLog, pagination *models.Pagination) ([]models.RefAttendanceLog, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
		if len(filter.State) > 0 {
			query = append(query, bson.M{"state": bson.M{"$in": filter.State}})
		}
		if len(filter.EmployeeId) > 0 {
			query = append(query, bson.M{"employeeId": bson.M{"$in": filter.EmployeeId}})
		}
		if filter.UniqueID != nil {
			query = append(query, bson.M{"uniqueId": bson.M{"$in": filter.UniqueID}})
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

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONATTENDANCELOG).CountDocuments(ctx.CTX, func() bson.M {
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
	cursor, err := ctx.DB.Collection(constants.COLLECTIONATTENDANCELOG).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var attendanceLogFilter []models.RefAttendanceLog
	if err = cursor.All(context.TODO(), &attendanceLogFilter); err != nil {
		return nil, err
	}
	return attendanceLogFilter, nil
}
func (d *Daos) RecentLoginCheck(ctx *models.Context, employeeId string, uniqueID string) (*models.AttendanceLog, error) {

	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"employeeId": employeeId, "attendanceId": uniqueID, "loginMode": constants.ATTENDANCESTATUSLOGIN}})
	mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"_id": -1}})
	d.Shared.BsonToJSONPrintTag("RecentLoginCheck =>", mainPipeline)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONATTENDANCELOG).Aggregate(ctx.CTX, mainPipeline, nil)
	//fmt.Println("cursor value", cursor)
	if err != nil {
		return nil, err
	}
	var AttendanceLogs []models.AttendanceLog
	var AttendanceLog *models.AttendanceLog
	if err = cursor.All(ctx.CTX, &AttendanceLogs); err != nil {
		return nil, err
	}
	if len(AttendanceLogs) > 0 {
		AttendanceLog = &AttendanceLogs[0]
	}
	return AttendanceLog, nil
}
