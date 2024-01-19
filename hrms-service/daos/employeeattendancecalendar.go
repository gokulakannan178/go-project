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

//SaveEmployeeAttendanceCalendar : ""
func (d *Daos) SaveEmployeeAttendanceCalendar(ctx *models.Context, employeeAttendanceCalendar *models.EmployeeAttendanceCalendar) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEATTENDANCECALENDAR).InsertOne(ctx.CTX, employeeAttendanceCalendar)
	if err != nil {
		return err
	}
	employeeAttendanceCalendar.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//UpdateEmployeeAttendanceCalendar : ""
func (d *Daos) UpdateEmployeeAttendanceCalendar(ctx *models.Context, employeeAttendanceCalendar *models.EmployeeAttendanceCalendar) error {
	selector := bson.M{"uniqueId": employeeAttendanceCalendar.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": employeeAttendanceCalendar}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEATTENDANCECALENDAR).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// GetSingleEmployeeAttendanceCalendar : ""
func (d *Daos) GetSingleEmployeeAttendanceCalendar(ctx *models.Context, uniqueID string) (*models.RefEmployeeAttendanceCalendar, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	//LookUp
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEATTENDANCECALENDAR).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var EmployeeAttendanceCalendars []models.RefEmployeeAttendanceCalendar
	var EmployeeAttendanceCalendar *models.RefEmployeeAttendanceCalendar
	if err = cursor.All(ctx.CTX, &EmployeeAttendanceCalendars); err != nil {
		return nil, err
	}
	if len(EmployeeAttendanceCalendars) > 0 {
		EmployeeAttendanceCalendar = &EmployeeAttendanceCalendars[0]
	}
	return EmployeeAttendanceCalendar, err
}

// EnableEmployeeAttendanceCalendar : ""
func (d *Daos) EnableEmployeeAttendanceCalendar(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.EMPLOYEEATTENDANCECALENDARSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEATTENDANCECALENDAR).UpdateOne(ctx.CTX, selector, data)
	return err
}

// DisableEmployeeAttendanceCalendar : ""
func (d *Daos) DisableEmployeeAttendanceCalendar(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.EMPLOYEEATTENDANCECALENDARSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEATTENDANCECALENDAR).UpdateOne(ctx.CTX, selector, data)
	return err
}

//DeleteEmployeeAttendanceCalendar :""
func (d *Daos) DeleteEmployeeAttendanceCalendar(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.EMPLOYEEATTENDANCECALENDARSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEATTENDANCECALENDAR).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterEmployeeAttendanceCalendar : ""
func (d *Daos) FilterEmployeeAttendanceCalendar(ctx *models.Context, EmployeeAttendanceCalendar *models.FilterEmployeeAttendanceCalendar, pagination *models.Pagination) ([]models.RefEmployeeAttendanceCalendar, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if EmployeeAttendanceCalendar != nil {
		if len(EmployeeAttendanceCalendar.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": EmployeeAttendanceCalendar.Status}})
		}
		if len(EmployeeAttendanceCalendar.OrganisationId) > 0 {
			query = append(query, bson.M{"organisationId": bson.M{"$in": EmployeeAttendanceCalendar.OrganisationId}})
		}
		//Regex
		if EmployeeAttendanceCalendar.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: EmployeeAttendanceCalendar.Regex.Name, Options: "xi"}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if EmployeeAttendanceCalendar != nil {
		if EmployeeAttendanceCalendar.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{EmployeeAttendanceCalendar.SortBy: EmployeeAttendanceCalendar.SortOrder}})
		}
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEATTENDANCECALENDAR).CountDocuments(ctx.CTX, func() bson.M {
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
	cursor, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEATTENDANCECALENDAR).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var EmployeeAttendanceCalendarFilter []models.RefEmployeeAttendanceCalendar
	if err = cursor.All(context.TODO(), &EmployeeAttendanceCalendarFilter); err != nil {
		return nil, err
	}
	return EmployeeAttendanceCalendarFilter, nil
}
func (d *Daos) GetSingleEmployeeAttendanceCalendarWithCurrentMonth(ctx *models.Context) (*models.RefEmployeeAttendanceCalendar, error) {
	t := time.Now()
	sd := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
	ed := time.Date(t.Year(), t.Month()+1, 0, 23, 59, 59, 0, t.Location())
	query := []bson.M{}
	query = append(query, bson.M{})
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"startDate": bson.M{"$gte": sd, "$lte": ed}}})
	//LookUp
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)

	d.Shared.BsonToJSONPrintTag("EmployeeAttendanceCalendar query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEEATTENDANCECALENDAR).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var EmployeeAttendanceCalendars []models.RefEmployeeAttendanceCalendar
	var EmployeeAttendanceCalendar *models.RefEmployeeAttendanceCalendar
	if err = cursor.All(ctx.CTX, &EmployeeAttendanceCalendars); err != nil {
		return nil, err
	}
	if len(EmployeeAttendanceCalendars) > 0 {
		EmployeeAttendanceCalendar = &EmployeeAttendanceCalendars[0]
	}
	return EmployeeAttendanceCalendar, err
}
