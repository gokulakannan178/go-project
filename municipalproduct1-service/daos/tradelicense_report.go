package daos

import (
	"context"
	"errors"
	"fmt"
	"log"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// DashboardDayWiseTradelicenseCollectionChart : ""
func (d *Daos) DashboardDayWiseTradelicenseCollectionChart(ctx *models.Context, filter *models.DashboardDayWiseTradeLicenseCollectionChartFilter) (models.DashboardDayWiseTradeLicenseCollectionChart, error) {

	mainPipeline := []bson.M{}
	var sd, ed *time.Time
	query := []bson.M{}
	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})

		}
		if filter.StartDate == nil {
			t := time.Now()
			filter.StartDate = &t
		}
		sdt := d.Shared.BeginningOfMonth(*filter.StartDate)
		sd = &sdt
		edt := d.Shared.EndOfMonth(*filter.StartDate)
		ed = &edt
		query = append(query, bson.M{"completionDate": bson.M{"$gte": sd, "$lte": ed}})
	}

	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	mainPipeline = append(mainPipeline, bson.M{"$group": bson.M{"_id": bson.M{
		"shops": "$tradeLicenseId",
		"day":   bson.M{"$dayOfMonth": "$completionDate"},
	}, "amount": bson.M{"$sum": "$details.amount"}}})
	mainPipeline = append(mainPipeline, bson.M{"$group": bson.M{"_id": "$_id.day", "tradeLicenseCount": bson.M{"$sum": 1}, "amount": bson.M{"$sum": "$amount"}}})
	mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"_id": 1}})
	mainPipeline = append(mainPipeline, bson.M{"$group": bson.M{"_id": nil, "days": bson.M{"$push": "$$ROOT"}}})

	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{
		"records": bson.M{
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
									"_id":               "$$rangeDay",
									"tradeLicenseCount": 0,
									"amount":            0.0,
								},
								"else": bson.M{"$arrayElemAt": []string{"$days", "$$index"}}},
						},
					},
				},
			},
		},
	}})

	//Aggregation
	d.Shared.BsonToJSONPrint(mainPipeline)
	var emptyData models.DashboardDayWiseTradeLicenseCollectionChart

	cursor, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSEPAYMENTS).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return emptyData, err
	}
	var data []models.DashboardDayWiseTradeLicenseCollectionChart

	if err = cursor.All(context.TODO(), &data); err != nil {
		return emptyData, err
	}
	if len(data) > 0 {
		return data[0], nil
	}

	return emptyData, nil

}

// DayWiseTradeLicenseDemandChart : ""
func (d *Daos) DayWiseTradeLicenseDemandChart(ctx *models.Context, filter *models.DayWiseTradeLicenseDemandChartFilter) (*models.DayWiseTradeLicenseDemandChart, error) {

	mainPipeline := []bson.M{}
	var sd, ed *time.Time
	query := []bson.M{}
	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})

		}
		if filter.StartDate == nil {
			t := time.Now()
			filter.StartDate = &t
		}
		sdt := d.Shared.BeginningOfMonth(*filter.StartDate)
		sd = &sdt
		edt := d.Shared.EndOfMonth(*filter.StartDate)
		ed = &edt
		query = append(query, bson.M{"created.on": bson.M{"$gte": sd, "$lte": ed}})
	}

	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	mainPipeline = append(mainPipeline, bson.M{"$group": bson.M{"_id": bson.M{"$dayOfMonth": "$created.on"},
		"tradeLicenseCount": bson.M{"$sum": 1}, "amount": bson.M{"$sum": "$demand.total.total"}}})
	// mainPipeline = append(mainPipeline, bson.M{"$group": bson.M{"_id": "$_id.day", "tradeLicenseCount": bson.M{"$sum": 1}, "amount": bson.M{"$sum": "$amount"}}})
	mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"_id": 1}})
	mainPipeline = append(mainPipeline, bson.M{"$group": bson.M{"_id": nil, "days": bson.M{"$push": "$$ROOT"}}})

	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{
		"records": bson.M{
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
									"_id":               "$$rangeDay",
									"tradeLicenseCount": 0,
									"amount":            0.0,
								},
								"else": bson.M{"$arrayElemAt": []string{"$days", "$$index"}}},
						},
					},
				},
			},
		},
	}})

	//Aggregation
	d.Shared.BsonToJSONPrint(mainPipeline)
	var emptyData *models.DayWiseTradeLicenseDemandChart

	cursor, err := ctx.DB.Collection(constants.COLLECTIONTRADELICENSE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return emptyData, err
	}
	var data []models.DayWiseTradeLicenseDemandChart

	if err = cursor.All(context.TODO(), &data); err != nil {
		return emptyData, err
	}
	if len(data) > 0 {
		return &data[0], nil
	}

	return emptyData, nil

}

// FilterWardDayWiseTradeLicenseCollectionReport : ""
func (d *Daos) FilterWardDayWiseTradeLicenseCollectionReport(ctx *models.Context, filter *models.WardDayWiseTradeLicenseCollectionReportFilter, pagination *models.Pagination) ([]models.WardDayWiseTradeLicenseCollectionReport, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.Ward) > 0 {
			query = append(query, bson.M{"code": bson.M{"$in": filter.Ward}})
		}
		if len(filter.Zone) > 0 {
			query = append(query, bson.M{"zoneCode": bson.M{"$in": filter.Zone}})
		}
	}
	var sd, ed time.Time
	if filter != nil {
		if filter.Date == nil {
			return nil, errors.New("please select a date")
		}
		sd = time.Date(filter.Date.Year(), filter.Date.Month(), filter.Date.Day(), 0, 0, 0, 0, sd.Location())
		ed = time.Date(filter.Date.Year(), filter.Date.Month(), filter.Date.Day(), 23, 59, 59, 0, ed.Location())
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if filter.SortBy != "" {
		mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{filter.SortBy: filter.SortOrder}})
	} else {
		mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"sortOrder": 1}})
	}
	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONWARD).CountDocuments(ctx.CTX, func() bson.M {
			if query != nil {
				if len(query) > 0 {
					return bson.M{"$and": query}
				}
			}
			return bson.M{}
		}())
		if err != nil {
			log.Println("Error in geting pagination")
		}
		fmt.Println("count", totalCount)
		pagination.Count = int(totalCount)
		d.Shared.PaginationData(pagination)
	}

	// LookUp
	mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"sortOrder": 1}})
	mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{
		"from": constants.COLLECTIONTRADELICENSEPAYMENTS,
		"as":   "report",
		"let":  bson.M{"status": "$status", "code": "$code"},
		"pipeline": []bson.M{{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
			{"$eq": []interface{}{"$status", constants.TRADELICENSEPAYMENRSTATUSCOMPLETED}},
			{"$eq": []interface{}{"$$code", "$address.wardCode"}},
			{"$gte": []interface{}{"$completionDate", sd}},
			{"$lte": []interface{}{"$completionDate", ed}},
		}}}},
			{"$group": bson.M{"_id": "$tradeLicenseId", "totalNoPayments": bson.M{"$sum": 1}, "totalCollection": bson.M{"$sum": "$details.amount"}}},
			{"$group": bson.M{"_id": nil, "totalNoProperties": bson.M{"$sum": 1}, "totalNoPayments": bson.M{"$sum": "$totalNoPayments"}, "totalCollections": bson.M{"$sum": "$totalCollection"}}},
		},
	},
	},
		bson.M{"$addFields": bson.M{"report": bson.M{"$arrayElemAt": []interface{}{"$report", 0}}}})

	//Aggregation
	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONWARD).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var res []models.WardDayWiseTradeLicenseCollectionReport
	if err = cursor.All(context.TODO(), &res); err != nil {
		return nil, err
	}
	return res, nil
}

// FilterWardMonthWiseTradeLicenseCollectionReport : ""
func (d *Daos) FilterWardMonthWiseTradeLicenseCollectionReport(ctx *models.Context, filter *models.WardMonthWiseTradeLicenseCollectionReportFilter, pagination *models.Pagination) ([]models.WardMonthWiseTradeLicenseCollectionReport, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.Ward) > 0 {
			query = append(query, bson.M{"code": bson.M{"$in": filter.Ward}})
		}
		if len(filter.Zone) > 0 {
			query = append(query, bson.M{"zoneCode": bson.M{"$in": filter.Zone}})
		}
	}
	var sd, ed *time.Time
	if filter.Date == nil {
		t := time.Now()
		filter.Date = &t
	}
	sdt := d.Shared.BeginningOfMonth(*filter.Date)
	sd = &sdt
	edt := d.Shared.EndOfMonth(*filter.Date)
	ed = &edt
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if filter.SortBy != "" {
		mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{filter.SortBy: filter.SortOrder}})
	} else {
		mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"sortOrder": 1}})
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONWARD).CountDocuments(ctx.CTX, func() bson.M {
			if query != nil {
				if len(query) > 0 {
					return bson.M{"$and": query}
				}
			}
			return bson.M{}
		}())
		if err != nil {
			log.Println("Error in geting pagination")
		}
		fmt.Println("count", totalCount)
		pagination.Count = int(totalCount)
		d.Shared.PaginationData(pagination)
	}

	// LookUp
	mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"sortOrder": 1}})
	mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{
		"from": constants.COLLECTIONTRADELICENSEPAYMENTS,
		"as":   "report",
		"let":  bson.M{"status": "$status", "code": "$code"},
		"pipeline": []bson.M{{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
			{"$eq": []interface{}{"$status", constants.TRADELICENSEPAYMENRSTATUSCOMPLETED}},
			{"$eq": []interface{}{"$$code", "$address.wardCode"}},
			{"$gte": []interface{}{"$completionDate", sd}},
			{"$lte": []interface{}{"$completionDate", ed}},
		}}}},
			{"$group": bson.M{"_id": "$tradeLicenseId", "totalNoPayments": bson.M{"$sum": 1}, "totalCollection": bson.M{"$sum": "$details.amount"}}},
			{"$group": bson.M{"_id": nil, "totalNoProperties": bson.M{"$sum": 1}, "totalNoPayments": bson.M{"$sum": "$totalNoPayments"}, "totalCollections": bson.M{"$sum": "$totalCollection"}}},
		},
	},
	},
		bson.M{"$addFields": bson.M{"report": bson.M{"$arrayElemAt": []interface{}{"$report", 0}}}})

	//Aggregation
	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONWARD).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var res []models.WardMonthWiseTradeLicenseCollectionReport
	if err = cursor.All(context.TODO(), &res); err != nil {
		return nil, err
	}
	return res, nil
}

// FilterWardDayWiseTradeLicenseDemandReport : ""
func (d *Daos) FilterWardDayWiseTradeLicenseDemandReport(ctx *models.Context, filter *models.WardDayWiseTradeLicenseDemandReportFilter, pagination *models.Pagination) ([]models.WardDayWiseTradeLicenseDemandReport, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.Ward) > 0 {
			query = append(query, bson.M{"code": bson.M{"$in": filter.Ward}})
		}
		if len(filter.Zone) > 0 {
			query = append(query, bson.M{"zoneCode": bson.M{"$in": filter.Zone}})
		}
	}
	var sd, ed time.Time
	if filter != nil {
		if filter.Date == nil {
			return nil, errors.New("please select a date")
		}
		sd = time.Date(filter.Date.Year(), filter.Date.Month(), filter.Date.Day(), 0, 0, 0, 0, sd.Location())
		ed = time.Date(filter.Date.Year(), filter.Date.Month(), filter.Date.Day(), 23, 59, 59, 0, ed.Location())
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if filter.SortBy != "" {
		mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{filter.SortBy: filter.SortOrder}})
	} else {
		mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"sortOrder": 1}})
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONWARD).CountDocuments(ctx.CTX, func() bson.M {
			if query != nil {
				if len(query) > 0 {
					return bson.M{"$and": query}
				}
			}
			return bson.M{}
		}())
		if err != nil {
			log.Println("Error in geting pagination")
		}
		fmt.Println("count", totalCount)
		pagination.Count = int(totalCount)
		d.Shared.PaginationData(pagination)
	}

	// LookUp
	mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"sortOrder": 1}})
	mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{
		"from": constants.COLLECTIONTRADELICENSE,
		"as":   "report",
		"let":  bson.M{"status": "$status", "code": "$code"},
		"pipeline": []bson.M{bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
			bson.M{"$eq": []interface{}{"$status", constants.TRADELICENSESTATUSACTIVE}},
			{"$eq": []interface{}{"$$code", "$address.wardCode"}},
			bson.M{"$gte": []interface{}{"$created.on", sd}},
			bson.M{"$lte": []interface{}{"$created.on", ed}},
		}}}},
			bson.M{"$group": bson.M{"_id": nil, "tradeLicenses": bson.M{"$sum": 1}, "totalDemand": bson.M{"$sum": "$demand.total.total"}}},
		},
	},
	},
		bson.M{"$addFields": bson.M{"report": bson.M{"$arrayElemAt": []interface{}{"$report", 0}}}})

	//Aggregation
	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONWARD).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var res []models.WardDayWiseTradeLicenseDemandReport
	if err = cursor.All(context.TODO(), &res); err != nil {
		return nil, err
	}
	return res, nil
}

// FilterWardMonthWiseTradeLicenseDemandReport : ""
func (d *Daos) FilterWardMonthWiseTradeLicenseDemandReport(ctx *models.Context, filter *models.WardMonthWiseTradeLicenseDemandReportFilter, pagination *models.Pagination) ([]models.WardMonthWiseTradeLicenseDemandReport, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.Ward) > 0 {
			query = append(query, bson.M{"code": bson.M{"$in": filter.Ward}})
		}
		if len(filter.Zone) > 0 {
			query = append(query, bson.M{"zoneCode": bson.M{"$in": filter.Zone}})
		}
	}
	var sd, ed *time.Time
	if filter.Date == nil {
		t := time.Now()
		filter.Date = &t
	}
	sdt := d.Shared.BeginningOfMonth(*filter.Date)
	sd = &sdt
	edt := d.Shared.EndOfMonth(*filter.Date)
	ed = &edt
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if filter.SortBy != "" {
		mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{filter.SortBy: filter.SortOrder}})
	} else {
		mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"sortOrder": 1}})
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONWARD).CountDocuments(ctx.CTX, func() bson.M {
			if query != nil {
				if len(query) > 0 {
					return bson.M{"$and": query}
				}
			}
			return bson.M{}
		}())
		if err != nil {
			log.Println("Error in geting pagination")
		}
		fmt.Println("count", totalCount)
		pagination.Count = int(totalCount)
		d.Shared.PaginationData(pagination)
	}

	// LookUp
	mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"sortOrder": 1}})
	mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{
		"from": constants.COLLECTIONTRADELICENSE,
		"as":   "report",
		"let":  bson.M{"status": "$status", "code": "$code"},
		"pipeline": []bson.M{bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
			bson.M{"$eq": []interface{}{"$status", constants.TRADELICENSESTATUSACTIVE}},
			{"$eq": []interface{}{"$$code", "$address.wardCode"}},
			bson.M{"$gte": []interface{}{"$created.on", sd}},
			bson.M{"$lte": []interface{}{"$created.on", ed}},
		}}}},
			bson.M{"$group": bson.M{"_id": nil, "tradeLicenses": bson.M{"$sum": 1}, "totalDemand": bson.M{"$sum": "$demand.total.total"}}},
		},
	},
	},
		bson.M{"$addFields": bson.M{"report": bson.M{"$arrayElemAt": []interface{}{"$report", 0}}}})

	//Aggregation
	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONWARD).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var res []models.WardMonthWiseTradeLicenseDemandReport
	if err = cursor.All(context.TODO(), &res); err != nil {
		return nil, err
	}
	return res, nil
}

// FilterTeamDayWiseTradeLicenseCollectionReport : ""
func (d *Daos) FilterTeamDayWiseTradeLicenseCollectionReport(ctx *models.Context, filter *models.TeamDayWiseTradeLicenseCollectionReportFilter, pagination *models.Pagination) ([]models.TeamDayWiseTradeLicenseCollectionReport, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.UserName) > 0 {
			query = append(query, bson.M{"userName": bson.M{"$in": filter.UserName}})
		}
		if len(filter.Type) > 0 {
			query = append(query, bson.M{"type": bson.M{"$in": filter.Type}})
		}
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
	}
	var sd, ed time.Time
	if filter != nil {
		if filter.Date == nil {
			return nil, errors.New("please select a date")
		}
		sd = time.Date(filter.Date.Year(), filter.Date.Month(), filter.Date.Day(), 0, 0, 0, 0, sd.Location())
		ed = time.Date(filter.Date.Year(), filter.Date.Month(), filter.Date.Day(), 23, 59, 59, 0, ed.Location())
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONUSER).CountDocuments(ctx.CTX, func() bson.M {
			if query != nil {
				if len(query) > 0 {
					return bson.M{"$and": query}
				}
			}
			return bson.M{}
		}())
		if err != nil {
			log.Println("Error in geting pagination")
		}
		fmt.Println("count", totalCount)
		pagination.Count = int(totalCount)
		d.Shared.PaginationData(pagination)
	}

	// LookUp
	mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"name": 1}})
	mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{
		"from": constants.COLLECTIONTRADELICENSEPAYMENTS,
		"as":   "report",
		"let":  bson.M{"userName": "$userName"},
		"pipeline": []bson.M{bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
			bson.M{"$eq": []interface{}{"$status", constants.TRADELICENSEPAYMENRSTATUSCOMPLETED}},
			bson.M{"$eq": []interface{}{"$$userName", "$details.collector.by"}},

			bson.M{"$gte": []interface{}{"$completionDate", sd}},
			bson.M{"$lte": []interface{}{"$completionDate", ed}},
		}}}},
			bson.M{"$group": bson.M{"_id": "$tradeLicenseId", "totalNoPayments": bson.M{"$sum": 1},
				"totalCollection": bson.M{"$sum": "$details.amount"}}},
			bson.M{"$group": bson.M{"_id": nil, "totalNoTradeLicenses": bson.M{"$sum": 1}, "totalNoPayments": bson.M{"$sum": "$totalNoPayments"},
				"totalCollections": bson.M{"$sum": "$totalCollection"}}},
		},
	},
	},
		bson.M{"$addFields": bson.M{"report": bson.M{"$arrayElemAt": []interface{}{"$report", 0}}}})

	//Aggregation
	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONUSER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var res []models.TeamDayWiseTradeLicenseCollectionReport
	if err = cursor.All(context.TODO(), &res); err != nil {
		return nil, err
	}
	return res, nil
}

// FilterTeamMonthWiseTradeLicenseCollectionReport : ""
func (d *Daos) FilterTeamMonthWiseTradeLicenseCollectionReport(ctx *models.Context, filter *models.TeamMonthWiseTradeLicenseCollectionReportFilter, pagination *models.Pagination) ([]models.TeamMonthWiseTradeLicenseCollectionReport, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.UserName) > 0 {
			query = append(query, bson.M{"userName": bson.M{"$in": filter.UserName}})
		}
		if len(filter.Type) > 0 {
			query = append(query, bson.M{"type": bson.M{"$in": filter.Type}})
		}
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
	}
	var sd, ed *time.Time
	if filter.Date == nil {
		t := time.Now()
		filter.Date = &t
	}
	sdt := d.Shared.BeginningOfMonth(*filter.Date)
	sd = &sdt
	edt := d.Shared.EndOfMonth(*filter.Date)
	ed = &edt
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONUSER).CountDocuments(ctx.CTX, func() bson.M {
			if query != nil {
				if len(query) > 0 {
					return bson.M{"$and": query}
				}
			}
			return bson.M{}
		}())
		if err != nil {
			log.Println("Error in geting pagination")
		}
		fmt.Println("count", totalCount)
		pagination.Count = int(totalCount)
		d.Shared.PaginationData(pagination)
	}

	// LookUp
	mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"name": 1}})
	mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{
		"from": constants.COLLECTIONTRADELICENSEPAYMENTS,
		"as":   "report",
		"let":  bson.M{"userName": "$userName"},
		"pipeline": []bson.M{bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
			bson.M{"$eq": []interface{}{"$status", constants.TRADELICENSEPAYMENRSTATUSCOMPLETED}},
			bson.M{"$eq": []interface{}{"$$userName", "$details.collector.by"}},

			bson.M{"$gte": []interface{}{"$completionDate", sd}},
			bson.M{"$lte": []interface{}{"$completionDate", ed}},
		}}}},
			bson.M{"$group": bson.M{"_id": "$tradeLicenseId", "totalNoPayments": bson.M{"$sum": 1},
				"totalCollection": bson.M{"$sum": "$details.amount"}}},
			bson.M{"$group": bson.M{"_id": nil, "totalNoTradeLicenses": bson.M{"$sum": 1}, "totalNoPayments": bson.M{"$sum": "$totalNoPayments"},
				"totalCollections": bson.M{"$sum": "$totalCollection"}}},
		},
	},
	},
		bson.M{"$addFields": bson.M{"report": bson.M{"$arrayElemAt": []interface{}{"$report", 0}}}})

	//Aggregation
	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONUSER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var res []models.TeamMonthWiseTradeLicenseCollectionReport
	if err = cursor.All(context.TODO(), &res); err != nil {
		return nil, err
	}
	return res, nil
}

// FilterWardYearWiseTradeLicenseCollectionReport : ""
func (d *Daos) FilterWardYearWiseTradeLicenseCollectionReport(ctx *models.Context, filter *models.WardYearWiseTradeLicenseCollectionReportFilter, pagination *models.Pagination) ([]models.WardYearWiseTradeLicenseCollectionReport, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	resFYs, err := d.GetSingleFinancialYear(ctx, filter.UniqueID)
	if err != nil {
		return nil, err
	}
	if filter != nil {
		if len(filter.Ward) > 0 {
			query = append(query, bson.M{"code": bson.M{"$in": filter.Ward}})
		}
		if len(filter.Zone) > 0 {
			query = append(query, bson.M{"zoneCode": bson.M{"$in": filter.Zone}})
		}
	}
	var sd, ed *time.Time
	// if filter.Date == nil {
	// 	t := time.Now()
	// 	filter.Date = &t
	// }
	// sdt := d.Shared.BeginningOfYear(*filter.Date)
	// sd = &sdt
	// edt := d.Shared.EndOfYear(*filter.Date)
	// ed = &edt
	sd = resFYs.From
	ed = resFYs.To
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if filter.SortBy != "" {
		mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{filter.SortBy: filter.SortOrder}})
	} else {
		mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"sortOrder": 1}})
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONWARD).CountDocuments(ctx.CTX, func() bson.M {
			if query != nil {
				if len(query) > 0 {
					return bson.M{"$and": query}
				}
			}
			return bson.M{}
		}())
		if err != nil {
			log.Println("Error in geting pagination")
		}
		fmt.Println("count", totalCount)
		pagination.Count = int(totalCount)
		d.Shared.PaginationData(pagination)
	}

	// LookUp
	mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"sortOrder": 1}})
	mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{
		"from": constants.COLLECTIONTRADELICENSEPAYMENTS,
		"as":   "report",
		"let":  bson.M{"status": "$status", "code": "$code"},
		"pipeline": []bson.M{{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
			{"$eq": []interface{}{"$status", constants.TRADELICENSEPAYMENRSTATUSCOMPLETED}},
			{"$eq": []interface{}{"$$code", "$address.wardCode"}},
			{"$gte": []interface{}{"$completionDate", sd}},
			{"$lte": []interface{}{"$completionDate", ed}},
		}}}},
			{"$group": bson.M{"_id": "$tradeLicenseId", "totalNoPayments": bson.M{"$sum": 1}, "totalCollection": bson.M{"$sum": "$details.amount"}}},
			{"$group": bson.M{"_id": nil, "totalNoProperties": bson.M{"$sum": 1}, "totalNoPayments": bson.M{"$sum": "$totalNoPayments"}, "totalCollections": bson.M{"$sum": "$totalCollection"}}},
		},
	},
	},
		bson.M{"$addFields": bson.M{"report": bson.M{"$arrayElemAt": []interface{}{"$report", 0}}}})

	//Aggregation
	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONWARD).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var res []models.WardYearWiseTradeLicenseCollectionReport
	if err = cursor.All(context.TODO(), &res); err != nil {
		return nil, err
	}
	return res, nil
}

// FilterWardYearWiseTradeLicenseDemandReport : ""
func (d *Daos) FilterWardYearWiseTradeLicenseDemandReport(ctx *models.Context, filter *models.WardYearWiseTradeLicenseDemandReportFilter, pagination *models.Pagination) ([]models.WardYearWiseTradeLicenseDemandReport, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	resFYs, err := d.GetSingleFinancialYear(ctx, filter.UniqueID)
	if err != nil {
		return nil, err
	}
	if filter != nil {
		if len(filter.Ward) > 0 {
			query = append(query, bson.M{"code": bson.M{"$in": filter.Ward}})
		}
		if len(filter.Zone) > 0 {
			query = append(query, bson.M{"zoneCode": bson.M{"$in": filter.Zone}})
		}
	}
	var sd, ed *time.Time
	// if filter.Date == nil {
	// 	t := time.Now()
	// 	filter.Date = &t
	// }
	// sdt := d.Shared.BeginningOfYear(*filter.Date)
	// sd = &sdt
	// edt := d.Shared.EndOfYear(*filter.Date)
	// ed = &edt
	sd = resFYs.From
	ed = resFYs.To
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if filter.SortBy != "" {
		mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{filter.SortBy: filter.SortOrder}})
	} else {
		mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"sortOrder": 1}})
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONWARD).CountDocuments(ctx.CTX, func() bson.M {
			if query != nil {
				if len(query) > 0 {
					return bson.M{"$and": query}
				}
			}
			return bson.M{}
		}())
		if err != nil {
			log.Println("Error in geting pagination")
		}
		fmt.Println("count", totalCount)
		pagination.Count = int(totalCount)
		d.Shared.PaginationData(pagination)
	}

	// LookUp
	mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"sortOrder": 1}})
	mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{
		"from": constants.COLLECTIONTRADELICENSE,
		"as":   "report",
		"let":  bson.M{"status": "$status", "code": "$code"},
		"pipeline": []bson.M{bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
			bson.M{"$eq": []interface{}{"$status", constants.TRADELICENSESTATUSACTIVE}},
			{"$eq": []interface{}{"$$code", "$address.wardCode"}},
			bson.M{"$gte": []interface{}{"$created.on", sd}},
			bson.M{"$lte": []interface{}{"$created.on", ed}},
		}}}},
			bson.M{"$group": bson.M{"_id": nil, "tradeLicenses": bson.M{"$sum": 1}, "totalDemand": bson.M{"$sum": "$demand.total.total"}}},
		},
	},
	},
		bson.M{"$addFields": bson.M{"report": bson.M{"$arrayElemAt": []interface{}{"$report", 0}}}})

	//Aggregation
	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONWARD).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var res []models.WardYearWiseTradeLicenseDemandReport
	if err = cursor.All(context.TODO(), &res); err != nil {
		return nil, err
	}
	return res, nil
}

// FilterTeamYearWiseTradeLicenseCollectionReport : ""
func (d *Daos) FilterTeamYearWiseTradeLicenseCollectionReport(ctx *models.Context, filter *models.TeamYearWiseTradeLicenseCollectionReportFilter, pagination *models.Pagination) ([]models.TeamYearWiseTradeLicenseCollectionReport, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	resFYs, err := d.GetSingleFinancialYear(ctx, filter.UniqueID)
	if err != nil {
		return nil, err
	}
	if filter != nil {
		if len(filter.UserName) > 0 {
			query = append(query, bson.M{"userName": bson.M{"$in": filter.UserName}})
		}
		if len(filter.Type) > 0 {
			query = append(query, bson.M{"type": bson.M{"$in": filter.Type}})
		}
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
	}
	var sd, ed *time.Time
	// if filter.Date == nil {
	// 	t := time.Now()
	// 	filter.Date = &t
	// }
	// sdt := d.Shared.BeginningOfYear(*filter.Date)
	// sd = &sdt
	// edt := d.Shared.EndOfYear(*filter.Date)
	// ed = &edt
	sd = resFYs.From
	ed = resFYs.To
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONUSER).CountDocuments(ctx.CTX, func() bson.M {
			if query != nil {
				if len(query) > 0 {
					return bson.M{"$and": query}
				}
			}
			return bson.M{}
		}())
		if err != nil {
			log.Println("Error in geting pagination")
		}
		fmt.Println("count", totalCount)
		pagination.Count = int(totalCount)
		d.Shared.PaginationData(pagination)
	}

	// LookUp
	mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"name": 1}})
	mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{
		"from": constants.COLLECTIONTRADELICENSEPAYMENTS,
		"as":   "report",
		"let":  bson.M{"userName": "$userName"},
		"pipeline": []bson.M{bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
			bson.M{"$eq": []interface{}{"$status", constants.TRADELICENSEPAYMENRSTATUSCOMPLETED}},
			bson.M{"$eq": []interface{}{"$$userName", "$details.collector.by"}},

			bson.M{"$gte": []interface{}{"$completionDate", sd}},
			bson.M{"$lte": []interface{}{"$completionDate", ed}},
		}}}},
			bson.M{"$group": bson.M{"_id": "$tradeLicenseId", "totalNoPayments": bson.M{"$sum": 1},
				"totalCollection": bson.M{"$sum": "$details.amount"}}},
			bson.M{"$group": bson.M{"_id": nil, "totalNoTradeLicenses": bson.M{"$sum": 1}, "totalNoPayments": bson.M{"$sum": "$totalNoPayments"},
				"totalCollections": bson.M{"$sum": "$totalCollection"}}},
		},
	},
	},
		bson.M{"$addFields": bson.M{"report": bson.M{"$arrayElemAt": []interface{}{"$report", 0}}}})

	//Aggregation
	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONUSER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var res []models.TeamYearWiseTradeLicenseCollectionReport
	if err = cursor.All(context.TODO(), &res); err != nil {
		return nil, err
	}
	return res, nil
}
