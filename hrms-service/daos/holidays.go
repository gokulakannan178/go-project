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

//SaveHolidays :""
func (d *Daos) SaveHolidays(ctx *models.Context, holidays *models.Holidays) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONHOLIDAYS).InsertOne(ctx.CTX, holidays)
	if err != nil {
		return err
	}
	holidays.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//GetSingleDocumentScenario : ""
func (d *Daos) GetSingleHolidays(ctx *models.Context, uniqueID string) (*models.RefHolidays, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID, "status": constants.HOLIDAYSSTATUSACTIVE}})
	// lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONHOLIDAYS).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var Holidayss []models.RefHolidays
	var Holidays *models.RefHolidays
	if err = cursor.All(ctx.CTX, &Holidayss); err != nil {
		return nil, err
	}
	if len(Holidayss) > 0 {
		Holidays = &Holidayss[0]
	}
	return Holidays, nil
}
func (d *Daos) GetSingleHolidaysWithOutStatus(ctx *models.Context, uniqueID string) (*models.RefHolidays, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	// lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONHOLIDAYS).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var Holidayss []models.RefHolidays
	var Holidays *models.RefHolidays
	if err = cursor.All(ctx.CTX, &Holidayss); err != nil {
		return nil, err
	}
	if len(Holidayss) > 0 {
		Holidays = &Holidayss[0]
	}
	return Holidays, nil
}

//UpdateHolidays : ""
func (d *Daos) UpdateHolidays(ctx *models.Context, holidays *models.Holidays) error {
	selector := bson.M{"uniqueId": holidays.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": holidays}
	_, err := ctx.DB.Collection(constants.COLLECTIONHOLIDAYS).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//EnableHolidays :""
func (d *Daos) EnableHolidays(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.HOLIDAYSSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONHOLIDAYS).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableHolidays :""
func (d *Daos) DisableHolidays(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.HOLIDAYSSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONHOLIDAYS).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteHolidays :""
func (d *Daos) DeleteHolidays(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.HOLIDAYSSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONHOLIDAYS).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//FilterHolidays : ""
func (d *Daos) FilterHolidays(ctx *models.Context, holidaysfilter *models.FilterHolidays, pagination *models.Pagination) ([]models.RefHolidays, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if holidaysfilter != nil {

		if len(holidaysfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": holidaysfilter.Status}})
		}
		if len(holidaysfilter.OrganisationID) > 0 {
			query = append(query, bson.M{"organisationId": bson.M{"$in": holidaysfilter.OrganisationID}})
		}
		//Regex
		if holidaysfilter.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: holidaysfilter.Regex.Name, Options: "xi"}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	if holidaysfilter != nil {
		if holidaysfilter.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{holidaysfilter.SortBy: holidaysfilter.SortOrder}})
		}
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONHOLIDAYS).CountDocuments(ctx.CTX, func() bson.M {
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
	// lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("DocumentScenario query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONHOLIDAYS).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var holidayssFilter []models.RefHolidays
	if err = cursor.All(context.TODO(), &holidayssFilter); err != nil {
		return nil, err
	}
	return holidayssFilter, nil
}

//GetHoildays With Dates
func (d *Daos) GetHolidaysWithDays(ctx *models.Context, holidaysfilter *models.FilterHolidays) ([]models.HolidaysList, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	//loc, _ := time.LoadLocation("Asia/Kolkata")
	//t := time.Now().In(loc)
	var sd, ed time.Time
	if holidaysfilter != nil {
		query = append(query, bson.M{"status": constants.HOLIDAYSSTATUSACTIVE})
		if len(holidaysfilter.OrganisationID) > 0 {
			query = append(query, bson.M{"organisationId": bson.M{"$in": holidaysfilter.OrganisationID}})
		}
		//Regex
		if holidaysfilter.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: holidaysfilter.Regex.Name, Options: "xi"}})
		}
		if holidaysfilter.Date.StartDate != nil {
			sd = time.Date(holidaysfilter.Date.StartDate.Year(), holidaysfilter.Date.StartDate.Month(), holidaysfilter.Date.StartDate.Day(), 0, 0, 0, 0, holidaysfilter.Date.StartDate.Location())
			if holidaysfilter.Date.EndDate == nil {
				ed = time.Date(holidaysfilter.Date.StartDate.Year(), holidaysfilter.Date.StartDate.Month(), holidaysfilter.Date.StartDate.Day(), 23, 59, 59, 999999999, holidaysfilter.Date.StartDate.Location())
			} else {
				ed = time.Date(holidaysfilter.Date.EndDate.Year(), holidaysfilter.Date.EndDate.Month(), holidaysfilter.Date.EndDate.Day(), 23, 59, 59, 999999999, holidaysfilter.Date.StartDate.Location())
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
	mainPipeline = append(mainPipeline, bson.M{"$group": bson.M{"_id": "$dom",
		"date":     bson.M{"$push": "$$ROOT.date"},
		"uniqueId": bson.M{"$push": "$$ROOT.uniqueId"},
		"hoildays": bson.M{"$push": "$$ROOT"}}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"date": bson.M{"$arrayElemAt": []interface{}{"$date", 0}}}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"uniqueId": bson.M{"$arrayElemAt": []interface{}{"$uniqueId", 0}}}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"dayOfWeek": bson.M{"$dayOfWeek": "$date"}}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"month1": bson.M{"$arrayElemAt": []interface{}{"$hoildays.month", 0}}}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"year1": bson.M{"$arrayElemAt": []interface{}{"$hoildays.year", 0}}}})
	mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"date": 1}})

	mainPipeline = append(mainPipeline, bson.M{"$group": bson.M{"_id": nil, "days": bson.M{"$push": "$$ROOT"}}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"month1": bson.M{"$arrayElemAt": []interface{}{"$days.month1", 0}}}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"year1": bson.M{"$arrayElemAt": []interface{}{"$days.year1", 0}}}})
	//	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"dom": bson.M{"$arrayElemAt": []interface{}{"$days.dom", 0}}}})

	//mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"dayOfWeek1": bson.M{"$arrayElemAt": []interface{}{"$days.dayOfWeek", 0}}}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{
		"days": bson.M{
			"$map": bson.M{
				"input": bson.M{"$range": []interface{}{sd.Day(), ed.Day() + 1, 1}},
				"as":    "rangeDay",
				"in": bson.M{
					"$let": bson.M{
						"vars": bson.M{"index": bson.M{"$indexOfArray": []string{"$days._id", "$$rangeDay"}}},
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
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{
		"date":     "$days.date",
		"name":     "$days.name",
		"uniqueId": "$days.uniqueId",
		"holidays": "$days.hoildays",
	}})

	// lookup
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("DocumentScenario query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONHOLIDAYS).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var holidayssFilter []models.HolidaysList
	if err = cursor.All(context.TODO(), &holidayssFilter); err != nil {
		return nil, err
	}
	return holidayssFilter, nil
}
func (d *Daos) GetSingleArrayHolidays(ctx *models.Context, uniqueID string) ([]models.Holidays, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID, "status": constants.HOLIDAYSSTATUSACTIVE}})
	// lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)
	d.Shared.BsonToJSONPrintTag("GetSingleArrayHolidays query =>", mainPipeline)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONHOLIDAYS).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var Holidayss []models.Holidays
	//var Holidays *models.RefHolidays
	if err = cursor.All(ctx.CTX, &Holidayss); err != nil {
		return nil, err
	}

	return Holidayss, nil
}
