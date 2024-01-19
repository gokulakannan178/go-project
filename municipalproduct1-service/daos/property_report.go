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

//PropertyDayWiseDemandReport : ""
func (d *Daos) DayWisePropertyDemandReport(ctx *models.Context, filter *models.DayWisePropertyDemandChartFilter) (*models.DayWisePropertyDemandChart, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	query = d.FilterPropertyQuery(ctx, &filter.PropertyFilter)
	var sd, ed *time.Time
	if filter != nil {
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
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if filter.SortBy != "" {
		mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{filter.SortBy: filter.SortOrder}})
	} else {
		mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"_id": -1}})
	}

	//Lookups
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONOVERALLPROPERTYDEMAND, "uniqueId", "propertyId", "ref.demand", "ref.demand")...)

	mainPipeline = append(mainPipeline, bson.M{"$group": bson.M{"_id": bson.M{"$dayOfMonth": "$created.on"},
		"propertyCount": bson.M{"$sum": 1}, "amount": bson.M{"$sum": "$ref.demand.total.totalTax"}}})
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
									"_id":           "$$rangeDay",
									"propertyCount": 0,
									"amount":        0.0,
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
	var emptyData *models.DayWisePropertyDemandChart

	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return emptyData, err
	}
	var data []models.DayWisePropertyDemandChart

	if err = cursor.All(context.TODO(), &data); err != nil {
		return emptyData, err
	}
	if len(data) > 0 {
		return &data[0], nil
	}

	return emptyData, nil

}

// FilterWardDayWisePropertyCollectionReport : ""
func (d *Daos) FilterWardDayWisePropertyCollectionReport(ctx *models.Context, filter *models.WardDayWisePropertyCollectionReportFilter, pagination *models.Pagination) ([]models.WardDayWisePropertyCollectionReport, error) {
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
		"from": constants.COLLECTIONPROPERTYPAYMENT,
		"as":   "report",
		"let":  bson.M{"status": "$status", "code": "$code"},
		"pipeline": []bson.M{{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
			{"$eq": []interface{}{"$status", constants.PROPERTYPAYMENTCOMPLETED}},
			{"$eq": []interface{}{"$$code", "$address.wardCode"}},
			{"$gte": []interface{}{"$completionDate", sd}},
			{"$lte": []interface{}{"$completionDate", ed}},
		}}}},
			{"$group": bson.M{"_id": "$propertyId", "totalNoPayments": bson.M{"$sum": 1}, "totalCollection": bson.M{"$sum": "$details.amount"}}},
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
	var res []models.WardDayWisePropertyCollectionReport
	if err = cursor.All(context.TODO(), &res); err != nil {
		return nil, err
	}
	return res, nil
}

// FilterWardMonthWisePropertyCollectionReport : ""
func (d *Daos) FilterWardMonthWisePropertyCollectionReport(ctx *models.Context, filter *models.WardMonthWisePropertyCollectionReportFilter, pagination *models.Pagination) ([]models.WardMonthWisePropertyCollectionReport, error) {
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
		"from": constants.COLLECTIONPROPERTYPAYMENT,
		"as":   "report",
		"let":  bson.M{"status": "$status", "code": "$code"},
		"pipeline": []bson.M{{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
			{"$eq": []interface{}{"$status", constants.PROPERTYPAYMENTCOMPLETED}},
			{"$eq": []interface{}{"$$code", "$address.wardCode"}},
			{"$gte": []interface{}{"$completionDate", sd}},
			{"$lte": []interface{}{"$completionDate", ed}},
		}}}},
			{"$group": bson.M{"_id": "$propertyId", "totalNoPayments": bson.M{"$sum": 1}, "totalCollection": bson.M{"$sum": "$details.amount"}}},
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
	var res []models.WardMonthWisePropertyCollectionReport
	if err = cursor.All(context.TODO(), &res); err != nil {
		return nil, err
	}
	return res, nil
}

// FilterWardDayWisePropertyDemandReport : ""
func (d *Daos) FilterWardDayWisePropertyDemandReport(ctx *models.Context, filter *models.WardDayWisePropertyDemandReportFilter, pagination *models.Pagination) ([]models.WardDayWisePropertyDemandReport, error) {
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
		"from": constants.COLLECTIONPROPERTY,
		"as":   "report",
		"let":  bson.M{"status": "$status", "code": "$code"},
		"pipeline": []bson.M{{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
			{"$eq": []interface{}{"$status", constants.PROPERTYSTATUSACTIVE}},
			{"$eq": []interface{}{"$$code", "$address.wardCode"}},
			{"$gte": []interface{}{"$created.on", sd}},
			{"$lte": []interface{}{"$created.on", ed}},
		}}}},
			{"$lookup": bson.M{
				"from": constants.COLLECTIONOVERALLPROPERTYDEMAND,
				"as":   "demandReport",
				"let":  bson.M{"uniqueId": "$uniqueId"},
				"pipeline": []bson.M{{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
					{"$eq": []interface{}{"$$uniqueId", "$propertyId"}},
				}}}}},
			},
			},

			{"$addFields": bson.M{"demandReport": bson.M{"$arrayElemAt": []interface{}{"$demandReport", 0}}}},

			{"$group": bson.M{"_id": nil,
				"totalNoProperties": bson.M{"$sum": 1},
				"totalDemand":       bson.M{"$sum": "$demandReport.total.totalTax"}}},
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
	var res []models.WardDayWisePropertyDemandReport
	if err = cursor.All(context.TODO(), &res); err != nil {
		return nil, err
	}
	return res, nil
}

// FilterWardMonthWisePropertyDemandReport : ""
func (d *Daos) FilterWardMonthWisePropertyDemandReport(ctx *models.Context, filter *models.WardMonthWisePropertyDemandReportFilter, pagination *models.Pagination) ([]models.WardMonthWisePropertyDemandReport, error) {
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
		"from": constants.COLLECTIONPROPERTY,
		"as":   "report",
		"let":  bson.M{"status": "$status", "code": "$code"},
		"pipeline": []bson.M{{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
			{"$eq": []interface{}{"$status", constants.PROPERTYSTATUSACTIVE}},
			{"$eq": []interface{}{"$$code", "$address.wardCode"}},
			{"$gte": []interface{}{"$created.on", sd}},
			{"$lte": []interface{}{"$created.on", ed}},
		}}}},
			{"$lookup": bson.M{
				"from": constants.COLLECTIONOVERALLPROPERTYDEMAND,
				"as":   "demandReport",
				"let":  bson.M{"uniqueId": "$uniqueId"},
				"pipeline": []bson.M{{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
					{"$eq": []interface{}{"$$uniqueId", "$propertyId"}},
				}}}}},
			},
			},

			{"$addFields": bson.M{"demandReport": bson.M{"$arrayElemAt": []interface{}{"$demandReport", 0}}}},

			{"$group": bson.M{"_id": nil,
				"totalNoProperties": bson.M{"$sum": 1},
				"totalDemand":       bson.M{"$sum": "$demandReport.total.totalTax"}}},
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
	var res []models.WardMonthWisePropertyDemandReport
	if err = cursor.All(context.TODO(), &res); err != nil {
		return nil, err
	}
	return res, nil
}

// FilterTeamDayWisePropertyCollectionReport : ""
func (d *Daos) FilterTeamDayWisePropertyCollectionReport(ctx *models.Context, filter *models.TeamDayWisePropertyCollectionReportFilter, pagination *models.Pagination) ([]models.TeamDayWisePropertyCollectionReport, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.UserName) > 0 {
			query = append(query, bson.M{"userName": bson.M{"$in": filter.UserName}})
		}
		if len(filter.Type) > 0 {
			query = append(query, bson.M{"type": bson.M{"$in": filter.Type}})
		}
		if len(filter.ManagerID) > 0 {
			query = append(query, bson.M{"managerId": bson.M{"$in": filter.ManagerID}})
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
		"from": constants.COLLECTIONPROPERTYPAYMENT,
		"as":   "report",
		"let":  bson.M{"userName": "$userName"},
		"pipeline": []bson.M{bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
			bson.M{"$eq": []interface{}{"$status", constants.PROPERTYPAYMENTCOMPLETED}},
			bson.M{"$eq": []interface{}{"$$userName", "$details.collector.id"}},

			bson.M{"$gte": []interface{}{"$completionDate", sd}},
			bson.M{"$lte": []interface{}{"$completionDate", ed}},
		}}}},
			bson.M{"$group": bson.M{"_id": "$propertyId",
				"totalNoPayments": bson.M{"$sum": 1},
				"totalCollection": bson.M{"$sum": "$details.amount"}}},
			bson.M{"$group": bson.M{"_id": nil,
				"totalNoProperties": bson.M{"$sum": 1}, "totalNoPayments": bson.M{"$sum": "$totalNoPayments"},
				"totalCollections": bson.M{"$sum": "$totalCollection"}}},
		},
	},
	},
		bson.M{"$addFields": bson.M{"report": bson.M{"$arrayElemAt": []interface{}{"$report", 0}}}},
	)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONUSER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var res []models.TeamDayWisePropertyCollectionReport
	if err = cursor.All(context.TODO(), &res); err != nil {
		return nil, err
	}
	return res, nil
}

// FilterTeamMonthWisePropertyCollectionReport : ""
func (d *Daos) FilterTeamMonthWisePropertyCollectionReport(ctx *models.Context, filter *models.TeamMonthWisePropertyCollectionReportFilter, pagination *models.Pagination) ([]models.TeamMonthWisePropertyCollectionReport, error) {
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
		if len(filter.ManagerID) > 0 {
			query = append(query, bson.M{"managerId": bson.M{"$in": filter.ManagerID}})
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
		"from": constants.COLLECTIONPROPERTYPAYMENT,
		"as":   "report",
		"let":  bson.M{"userName": "$userName"},
		"pipeline": []bson.M{bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
			bson.M{"$eq": []interface{}{"$status", constants.PROPERTYPAYMENTCOMPLETED}},
			bson.M{"$eq": []interface{}{"$$userName", "$details.collector.id"}},

			bson.M{"$gte": []interface{}{"$completionDate", sd}},
			bson.M{"$lte": []interface{}{"$completionDate", ed}},
		}}}},
			bson.M{"$group": bson.M{"_id": "$propertyId",
				"totalNoPayments": bson.M{"$sum": 1},
				"totalCollection": bson.M{"$sum": "$details.amount"}}},
			bson.M{"$group": bson.M{"_id": nil,
				"totalNoProperties": bson.M{"$sum": 1}, "totalNoPayments": bson.M{"$sum": "$totalNoPayments"},
				"totalCollections": bson.M{"$sum": "$totalCollection"}}},
		},
	},
	},
		bson.M{"$addFields": bson.M{"report": bson.M{"$arrayElemAt": []interface{}{"$report", 0}}}},
	)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONUSER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var res []models.TeamMonthWisePropertyCollectionReport
	if err = cursor.All(context.TODO(), &res); err != nil {
		return nil, err
	}
	return res, nil
}

// FilterWardYearWisePropertyCollectionReport : ""
func (d *Daos) FilterWardYearWisePropertyCollectionReport(ctx *models.Context, filter *models.WardYearWisePropertyCollectionReportFilter, pagination *models.Pagination) ([]models.WardYearWisePropertyCollectionReport, error) {
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
		"from": constants.COLLECTIONPROPERTYPAYMENT,
		"as":   "report",
		"let":  bson.M{"status": "$status", "code": "$code"},
		"pipeline": []bson.M{{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
			{"$eq": []interface{}{"$status", constants.PROPERTYPAYMENTCOMPLETED}},
			{"$eq": []interface{}{"$$code", "$address.wardCode"}},
			{"$gte": []interface{}{"$completionDate", sd}},
			{"$lte": []interface{}{"$completionDate", ed}},
		}}}},
			{"$group": bson.M{"_id": "$propertyId", "totalNoPayments": bson.M{"$sum": 1}, "totalCollection": bson.M{"$sum": "$details.amount"}}},
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
	var res []models.WardYearWisePropertyCollectionReport
	if err = cursor.All(context.TODO(), &res); err != nil {
		return nil, err
	}
	return res, nil
}

// FilterWardYearWisePropertyDemandReport : ""
func (d *Daos) FilterWardYearWisePropertyDemandReport(ctx *models.Context, filter *models.WardYearWisePropertyDemandReportFilter, pagination *models.Pagination) ([]models.WardYearWisePropertyDemandReport, error) {
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
		"from": constants.COLLECTIONPROPERTY,
		"as":   "report",
		"let":  bson.M{"status": "$status", "code": "$code"},
		"pipeline": []bson.M{{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
			{"$eq": []interface{}{"$status", constants.PROPERTYSTATUSACTIVE}},
			{"$eq": []interface{}{"$$code", "$address.wardCode"}},
			{"$gte": []interface{}{"$created.on", sd}},
			{"$lte": []interface{}{"$created.on", ed}},
		}}}},
			{"$lookup": bson.M{
				"from": constants.COLLECTIONOVERALLPROPERTYDEMAND,
				"as":   "demandReport",
				"let":  bson.M{"uniqueId": "$uniqueId"},
				"pipeline": []bson.M{{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
					{"$eq": []interface{}{"$$uniqueId", "$propertyId"}},
				}}}}},
			},
			},

			{"$addFields": bson.M{"demandReport": bson.M{"$arrayElemAt": []interface{}{"$demandReport", 0}}}},

			{"$group": bson.M{"_id": nil,
				"totalNoProperties": bson.M{"$sum": 1},
				"totalDemand":       bson.M{"$sum": "$demandReport.total.totalTax"}}},
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
	var res []models.WardYearWisePropertyDemandReport
	if err = cursor.All(context.TODO(), &res); err != nil {
		return nil, err
	}
	return res, nil
}

// FilterTeamYearWisePropertyCollectionReport : ""
func (d *Daos) FilterTeamYearWisePropertyCollectionReport(ctx *models.Context, filter *models.TeamYearWisePropertyCollectionReportFilter, pagination *models.Pagination) ([]models.TeamYearWisePropertyCollectionReport, error) {
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
		if len(filter.ManagerID) > 0 {
			query = append(query, bson.M{"managerId": bson.M{"$in": filter.ManagerID}})
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
		"from": constants.COLLECTIONPROPERTYPAYMENT,
		"as":   "report",
		"let":  bson.M{"userName": "$userName"},
		"pipeline": []bson.M{bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
			bson.M{"$eq": []interface{}{"$status", constants.PROPERTYPAYMENTCOMPLETED}},
			bson.M{"$eq": []interface{}{"$$userName", "$details.collector.id"}},

			bson.M{"$gte": []interface{}{"$completionDate", sd}},
			bson.M{"$lte": []interface{}{"$completionDate", ed}},
		}}}},
			bson.M{"$group": bson.M{"_id": "$propertyId",
				"totalNoPayments": bson.M{"$sum": 1},
				"totalCollection": bson.M{"$sum": "$details.amount"}}},
			bson.M{"$group": bson.M{"_id": nil,
				"totalNoProperties": bson.M{"$sum": 1}, "totalNoPayments": bson.M{"$sum": "$totalNoPayments"},
				"totalCollections": bson.M{"$sum": "$totalCollection"}}},
		},
	},
	},
		bson.M{"$addFields": bson.M{"report": bson.M{"$arrayElemAt": []interface{}{"$report", 0}}}},
	)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONUSER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var res []models.TeamYearWisePropertyCollectionReport
	if err = cursor.All(context.TODO(), &res); err != nil {
		return nil, err
	}
	return res, nil
}

// FilterWardYearWisePropertyCollectionReport : ""
func (d *Daos) FilterYearWisePropertyDemandReport(ctx *models.Context, filter *models.YearWisePropertyDemandReportFilter, pagination *models.Pagination) ([]models.YearWisePropertyDemandReport, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	resFYs, err := d.GetSingleFinancialYear(ctx, filter.FYID)
	if err != nil {
		return nil, err
	}
	if filter != nil {
		// if len(filter.Status) > 0 {
		// 	query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		// }
	}
	var sd, ed *time.Time
	fmt.Println("resFYs from ===>", resFYs.From)
	fmt.Println("resFYs to ===>", resFYs.To)
	sd = resFYs.From
	ed = resFYs.To
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	// LookUp
	mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"sortOrder": 1}})
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"status": bson.M{"$in": []string{constants.PROPERTYSTATUSACTIVE, constants.PROPERTYSTATUSPENDING}},
		"created.on": bson.M{"$gte": sd, "$lte": ed}}},
		bson.M{"$lookup": bson.M{
			"from":         constants.COLLECTIONOVERALLPROPERTYDEMAND,
			"as":           "ref.demand",
			"localField":   "uniqueId",
			"foreignField": "propertyId"}},
		bson.M{"$addFields": bson.M{"ref.demand": bson.M{"$arrayElemAt": []interface{}{"$ref.demand", 0}}}},

		bson.M{"$project": bson.M{"uniqueId": 1, "created": 1, "ref": 1}},
		bson.M{"$group": bson.M{"_id": bson.M{"$month": "$created.on"}, "noOfProperties": bson.M{"$sum": 1}, "demand": bson.M{"$sum": "$ref.demand.actual.total.totalTax"}}},
		bson.M{"$sort": bson.M{"_id": 1}},
		bson.M{"$addFields": bson.M{"month": "$_id"}},
		bson.M{"$project": bson.M{"_id": 0, "month": 1, "noOfProperties": 1, "demand": 1}})

	//Aggregation
	d.Shared.BsonToJSONPrintTag("property yearwise report query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var res []models.YearWisePropertyDemandReport
	if err = cursor.All(context.TODO(), &res); err != nil {
		return nil, err
	}
	fmt.Println("res===>", res)
	return res, nil

}

// DashboardYearWiseCollectionChart : "This api is used to get the records for property yearwise collection"
func (d *Daos) FilterYearWisePropertyCollectionReport(ctx *models.Context, filter *models.YearWisePropertyCollectionReportFilter) ([]models.YearWisePropertyCollectionReport, error) {

	mainPipeline := []bson.M{}
	query := []bson.M{}
	resFYs, err := d.GetSingleFinancialYear(ctx, filter.FYID)
	if err != nil {
		return nil, err
	}
	if filter != nil {
		// if len(filter.Status) > 0 {
		// 	query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		// }
	}
	var sd, ed *time.Time
	fmt.Println("resFYs from ===>", resFYs.From)
	fmt.Println("resFYs to ===>", resFYs.To)
	sd = resFYs.From
	ed = resFYs.To
	query = append(query, bson.M{"completionDate": bson.M{"$gte": sd, "$lte": ed}})

	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	// Lookup
	mainPipeline = append(mainPipeline, bson.M{"$group": bson.M{"_id": bson.M{
		"month": bson.M{"$month": "$completionDate"},
	}, "totalTax": bson.M{"$sum": "$demand.totalTax"},
		"propertyCount": bson.M{"$sum": 1},
	}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"month": "$_id.month"}})
	mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"month": 1}})
	mainPipeline = append(mainPipeline, bson.M{"$project": bson.M{"_id": 0}})

	// mainPipeline = append(mainPipeline, bson.M{"$group": bson.M{"_id": nil, "months": bson.M{"$push": "$$ROOT"}}})

	// mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{
	// 	"records": bson.M{
	// 		"$map": bson.M{
	// 			"input": bson.M{"$range": []interface{}{sd.Day(), ed.Day() + 1, 1}},
	// 			"as":    "rangeDay",
	// 			"in": bson.M{
	// 				"$let": bson.M{
	// 					"vars": bson.M{"index": bson.M{"$indexOfArray": []string{"$months._id.month", "$$rangeDay"}}},
	// 					"in": bson.M{
	// 						"$cond": bson.M{
	// 							"if":   bson.M{"$eq": []interface{}{"$$index", -1}},
	// 							"then": bson.M{"_id": bson.M{"month": "$$rangeDay"}},
	// 							"else": bson.M{"$arrayElemAt": []string{"$months", "$$index"}}},
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }})

	//Aggregation
	d.Shared.BsonToJSONPrint(mainPipeline)

	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYPAYMENT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var res []models.YearWisePropertyCollectionReport
	if err = cursor.All(context.TODO(), &res); err != nil {
		return nil, err
	}
	fmt.Println("res===>", res)
	return res, nil

}

// DashboardYearWiseCollectionChart : "This api is used to get the records for property yearwise collection"
func (d *Daos) FilterWardWisePropertyDemandAndCollectionReport(ctx *models.Context, filter *models.YearWisePropertyCollectionReportFilter) (models.WardWisePropertyDemandAndCollectionReport, error) {

	mainPipeline := []bson.M{}
	//var sd, ed *time.Time
	query := []bson.M{}
	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})

		}
		// if filter.StartDate == nil {
		// 	t := time.Now()
		// 	filter.StartDate = &t
		// }
		// sdt := d.Shared.BeginningOfMonth(*filter.StartDate)
		// sd = &sdt
		// edt := d.Shared.EndOfMonth(*filter.StartDate)
		// ed = &edt
		// query = append(query, bson.M{"completionDate": bson.M{"$gte": sd, "$lte": ed}})
	}

	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	mainPipeline = append(mainPipeline,
		bson.M{"$lookup": bson.M{
			"from": "propertypayments",
			"as":   "report",
			"let":  bson.M{"code": "$code"},
			"pipeline": []bson.M{
				bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
					bson.M{"$eq": []string{"$status", "Completed"}},
					bson.M{"$eq": []string{"$$code", "$address.wardCode"}},
				}}}},
			},
		}})
	mainPipeline = append(mainPipeline,
		bson.M{"$addFields": bson.M{"report": bson.M{"$arrayElemAt": []interface{}{"$report", 0}}}})

	mainPipeline = append(mainPipeline, bson.M{"$group": bson.M{
		"_id":                    "$code",
		"outstandingDemand":      bson.M{"$sum": "$report.demand.totalTax"},
		"currentDemand":          bson.M{"$sum": "$report.demand.current"},
		"totalAmount":            bson.M{"$sum": "$report.details.amount"},
		"paidAmount":             bson.M{"$sum": "$report.details.amountPaid"},
		"totalDemand":            bson.M{"$sum": "$report.demand.totalTax"},
		"arrearCollection":       bson.M{"$sum": "$report.details.amountPaid"},
		"currentCollection":      bson.M{"$sum": "$report.demand.totalTax"},
		"penalty":                bson.M{"$sum": "$report.demand.current"},
		"rebate":                 bson.M{"$sum": "$report.details.amount"},
		"advance":                bson.M{"$sum": "$report.details.amountPaid"},
		"totalCollection":        bson.M{"$sum": "$report.demand.totalTax"},
		"totalOutstandingDemand": bson.M{"$sum": "$report.details.amountPaid"},
	}})

	mainPipeline = append(mainPipeline, bson.M{"$group": bson.M{"_id": nil, "report": bson.M{"$push": "$$ROOT"}}})

	//Aggregation
	d.Shared.BsonToJSONPrint(mainPipeline)
	var emptyData models.WardWisePropertyDemandAndCollectionReport

	cursor, err := ctx.DB.Collection(constants.COLLECTIONWARD).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return emptyData, err
	}
	var data []models.WardWisePropertyDemandAndCollectionReport

	if err = cursor.All(context.TODO(), &data); err != nil {
		return emptyData, err
	}
	if len(data) > 0 {
		return data[0], nil
	}
	return emptyData, nil

}

// UserWisePropertyCollectionReport : ""
func (d *Daos) UserWisePropertyCollectionReport(ctx *models.Context, filter *models.UserWisePropertyCollectionReportFilter) ([]models.UserWisePropertyCollectionReport, error) {

	mainPipeline := []bson.M{}
	var sd, ed *time.Time
	query := []bson.M{}
	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
		if len(filter.UserType) > 0 {
			query = append(query, bson.M{"type": bson.M{"$in": filter.UserType}})

		}
		if filter.StartDate == nil {
			t := time.Now()
			filter.StartDate = &t

		}
		sdt := d.Shared.BeginningOfMonth(*filter.StartDate)
		sd = &sdt
		edt := d.Shared.EndOfMonth(*filter.StartDate)
		ed = &edt
		fmt.Println("sd ====>", sd)
		fmt.Println("ed ====>", ed)
		// query = append(query, bson.M{"completionDate": bson.M{"$gte": sd, "$lte": ed}})
	}

	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{
		"from": constants.COLLECTIONPROPERTYPAYMENT,
		"as":   "payments",
		"let":  bson.M{"userType": "$type", "name": "$userName"},
		"pipeline": []bson.M{
			bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
				bson.M{"$eq": []interface{}{"$status", constants.PROPERTYPAYMENTCOMPLETED}},
				bson.M{"$eq": []interface{}{"$details.collector.type", "$$userType"}},
				bson.M{"$eq": []interface{}{"$details.collector.id", "$$name"}},
				bson.M{"$gte": []interface{}{"$completionDate", sd}},
				bson.M{"$lte": []interface{}{"$completionDate", ed}},
			},
			},
			},
			},

			bson.M{"$group": bson.M{"_id": bson.M{"$dayOfMonth": "$completionDate"},
				"totalCollection": bson.M{"$sum": "$details.amount"},
			}},
			bson.M{"$sort": bson.M{"_id": 1}},

			bson.M{"$group": bson.M{"_id": nil,
				"days": bson.M{"$push": "$$ROOT"},
			}},
			bson.M{
				"$addFields": bson.M{
					"records": bson.M{
						"$map": bson.M{
							"as": "rangeDay",
							"in": bson.M{
								"$let": bson.M{
									"in": bson.M{
										"$cond": bson.M{
											"else": bson.M{
												"$arrayElemAt": []interface{}{"$days", "$$index"},
											},
											"if": bson.M{
												"$eq": []interface{}{"$$index", -1},
											},
											"then": bson.M{
												"_id":             "$$rangeDay",
												"totalCollection": 0,
											},
										},
									},
									"vars": bson.M{
										"index": bson.M{
											"$indexOfArray": []interface{}{"$days._id", "$$rangeDay"},
										},
									},
								},
							},
							"input": bson.M{
								"$range": []interface{}{1, 31, 1},
							},
						},
					},
				},
			},

			bson.M{"$project": bson.M{"days": 0}},
		},
	}})

	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"payments": bson.M{"$arrayElemAt": []interface{}{"$payments.records", 0}}}})

	//Aggregation
	d.Shared.BsonToJSONPrintTag("userwise report query", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONUSER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var res []models.UserWisePropertyCollectionReport
	if err = cursor.All(context.TODO(), &res); err != nil {
		return nil, err
	}

	return res, nil

}

// FilterPropertyArrearAndCurrentCollectionReportJSON : ""
func (d *Daos) FilterPropertyArrearAndCurrentCollectionReportJSON(ctx *models.Context, filter *models.PropertyArrearAndCurrentCollectionFilter) ([]models.PropertyArrearAndCurrentCollectionReport, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	var sd, ed time.Time

	if filter != nil {

		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
		if len(filter.Ward) > 0 {
			query = append(query, bson.M{"code": bson.M{"$in": filter.Ward}})
		}
	}
	if filter.DateRange.From != nil {
		sd = time.Date(filter.DateRange.From.Year(), filter.DateRange.From.Month(), filter.DateRange.From.Day(), 0, 0, 0, 0, filter.DateRange.From.Location())
		ed = time.Date(filter.DateRange.From.Year(), filter.DateRange.From.Month(), filter.DateRange.From.Day(), 23, 59, 59, 0, filter.DateRange.From.Location())
		if filter.DateRange.To != nil {
			ed = time.Date(filter.DateRange.To.Year(), filter.DateRange.To.Month(), filter.DateRange.To.Day(), 23, 59, 59, 0, filter.DateRange.To.Location())
		}

	}
	//Adding $match from filter

	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	fmt.Println("fromDate =====> ", sd)
	fmt.Println("toDate =====> ", ed)
	// LookUp
	mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"sortOrder": 1}})
	mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{
		"from": "propertypayments",
		"as":   "payments",
		"let":  bson.M{"varCode": "$code"},
		"pipeline": []bson.M{
			bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
				bson.M{"$eq": []string{"$status", constants.PROPERTYPAYMENTCOMPLETED}},
				bson.M{"$eq": []string{"$$varCode", "$address.wardCode"}},
				bson.M{"$gte": []interface{}{"$completionDate", sd}},
				bson.M{"$lte": []interface{}{"$completionDate", ed}},
			}}}},

			bson.M{"$lookup": bson.M{
				"from": "propertypaymentfys",
				"as":   "fys",
				"let":  bson.M{"varTnxId": "$tnxId"},
				"pipeline": []bson.M{
					bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
						bson.M{"$eq": []string{"$$varTnxId", "$tnxId"}},
					}}}},

					bson.M{"$group": bson.M{"_id": "$tnxId",
						"arrearTax": bson.M{"$sum": bson.M{"$cond": bson.M{"else": 0,
							"if":   bson.M{"$eq": []interface{}{"$fy.isCurrent", false}},
							"then": bson.M{"$sum": bson.M{"$add": []interface{}{"$fy.vacantLandTax", "$fy.tax"}}}}}},
						"currentTax": bson.M{"$sum": bson.M{"$cond": bson.M{"else": 0,
							"if":   bson.M{"$eq": []interface{}{"$fy.isCurrent", true}},
							"then": bson.M{"$sum": bson.M{"$add": []interface{}{"$fy.vacantLandTax", "$fy.tax"}}}}}},
						"arrearPenalty": bson.M{"$sum": bson.M{"$cond": bson.M{"else": 0,
							"if":   bson.M{"$eq": []interface{}{"$fy.isCurrent", false}},
							"then": bson.M{"$sum": bson.M{"$add": []interface{}{"$fy.penanty"}}}}}},
						"currentPenalty": bson.M{"$sum": bson.M{"$cond": bson.M{"else": 0,
							"if":   bson.M{"$eq": []interface{}{"$fy.isCurrent", true}},
							"then": bson.M{"$sum": bson.M{"$add": []interface{}{"$fy.penanty"}}}}}},
						"arrearRebate": bson.M{"$sum": bson.M{"$cond": bson.M{"else": 0,
							"if":   bson.M{"$eq": []interface{}{"$fy.isCurrent", false}},
							"then": bson.M{"$sum": bson.M{"$add": []interface{}{"$fy.rebate"}}}}}},
						"currentRebate": bson.M{"$sum": bson.M{"$cond": bson.M{"else": 0,
							"if":   bson.M{"$eq": []interface{}{"$fy.isCurrent", true}},
							"then": bson.M{"$sum": bson.M{"$add": []interface{}{"$fy.rebate"}}}}}},
						"arrearAlreadyPaid": bson.M{"$sum": bson.M{"$cond": bson.M{"else": 0,
							"if":   bson.M{"$eq": []interface{}{"$fy.isCurrent", false}},
							"then": bson.M{"$sum": bson.M{"$add": []interface{}{"$fy.alreadyPayed.fyTax", "$fy.alreadyPayed.vlTax"}}}}}},
						"currentAlreadyPaid": bson.M{"$sum": bson.M{"$cond": bson.M{"else": 0,
							"if":   bson.M{"$eq": []interface{}{"$fy.isCurrent", true}},
							"then": bson.M{"$sum": bson.M{"$add": []interface{}{"$fy.alreadyPayed.fyTax", "$fy.alreadyPayed.vlTax"}}}}}},
						"totalTax":    bson.M{"$sum": "$fy.totalTax"},
						"otherDemand": bson.M{"$sum": "$fy.otherDemand"},
					}},
				},
			}},
			bson.M{"$addFields": bson.M{"fys": bson.M{"$arrayElemAt": []interface{}{"$fys", 0}}}},
			bson.M{"$group": bson.M{"_id": "$propertyId",
				"totalProperties":    bson.M{"$sum": 1},
				"formFee":            bson.M{"$sum": "$demand.formFee"},
				"arrearTax":          bson.M{"$sum": "$fys.arrearTax"},
				"currentTax":         bson.M{"$sum": "$fys.currentTax"},
				"arrearPenalty":      bson.M{"$sum": "$fys.arrearPenalty"},
				"currentPenalty":     bson.M{"$sum": "$fys.currentPenalty"},
				"currentRebate":      bson.M{"$sum": "$fys.currentRebate"},
				"arrearRebate":       bson.M{"$sum": "$fys.arrearRebate"},
				"currentAlreadyPaid": bson.M{"$sum": "$fys.currentAlreadyPaid"},
				"arrearAlreadyPaid":  bson.M{"$sum": "$fys.arrearAlreadyPaid"},
				"totalTax":           bson.M{"$sum": "$fys.totalTax"},
				"otherDemand":        bson.M{"$sum": "$fys.otherDemand"},
			}},
			bson.M{"$group": bson.M{"_id": nil,
				"totalProperties":    bson.M{"$sum": "$totalProperties"},
				"formFee":            bson.M{"$sum": "$formFee"},
				"arrearTax":          bson.M{"$sum": "$arrearTax"},
				"currentTax":         bson.M{"$sum": "$currentTax"},
				"arrearPenalty":      bson.M{"$sum": "$arrearPenalty"},
				"currentPenalty":     bson.M{"$sum": "$currentPenalty"},
				"currentRebate":      bson.M{"$sum": "$currentRebate"},
				"arrearRebate":       bson.M{"$sum": "$arrearRebate"},
				"currentAlreadyPaid": bson.M{"$sum": "$currentAlreadyPaid"},
				"arrearAlreadyPaid":  bson.M{"$sum": "$arrearAlreadyPaid"},
				"totalTax":           bson.M{"$sum": "$totalTax"},
				"otherDemand":        bson.M{"$sum": "$otherDemand"},
			}},
		},
	}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"payments": bson.M{"$arrayElemAt": []interface{}{"$payments", 0}}}})
	//Aggregation
	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONWARD).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var res []models.PropertyArrearAndCurrentCollectionReport
	if err = cursor.All(context.TODO(), &res); err != nil {
		return nil, err
	}
	return res, nil
}

// FilterCounterReportV2JSON : ""
func (d *Daos) FilterCounterReportV2JSON(ctx *models.Context, filter *models.PropertyPaymentFilter, pagination *models.Pagination) ([]models.RefCounterReportV2, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	// mainPipeline = append(mainPipeline, d.FilterPropertyPaymentQuery(ctx, filter)...)

	query = d.FilterPropertyPaymentQuery(ctx, filter)
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if filter.SortBy != "" {
		mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{filter.SortBy: filter.SortOrder}})
	} else {
		mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"_id": -1}})
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYPAYMENT).CountDocuments(ctx.CTX, func() bson.M {
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

	mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{
		"from": "propertypaymentfys",
		"as":   "paymentFys",
		"let":  bson.M{"varTnxId": "$tnxId"},
		"pipeline": []bson.M{
			bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
				bson.M{"$eq": []string{"$$varTnxId", "$tnxId"}},
			}}}},
			bson.M{"$group": bson.M{"_id": "$tnxId",
				"arrearPenalty":      bson.M{"$sum": bson.M{"$cond": bson.M{"else": 0, "if": bson.M{"$eq": []interface{}{"$fy.isCurrent", false}}, "then": bson.M{"$sum": bson.M{"$add": []interface{}{"$fy.penanty"}}}}}},
				"arrearRebate":       bson.M{"$sum": bson.M{"$cond": bson.M{"else": 0, "if": bson.M{"$eq": []interface{}{"$fy.isCurrent", false}}, "then": bson.M{"$sum": bson.M{"$add": []interface{}{"$fy.rebate"}}}}}},
				"arrearTax":          bson.M{"$sum": bson.M{"$cond": bson.M{"else": 0, "if": bson.M{"$eq": []interface{}{"$fy.isCurrent", false}}, "then": bson.M{"$sum": bson.M{"$add": []interface{}{"$fy.vacantLandTax", "$fy.tax"}}}}}},
				"currentPenalty":     bson.M{"$sum": bson.M{"$cond": bson.M{"else": 0, "if": bson.M{"$eq": []interface{}{"$fy.isCurrent", true}}, "then": bson.M{"$sum": bson.M{"$add": []interface{}{"$fy.penanty"}}}}}},
				"currentRebate":      bson.M{"$sum": bson.M{"$cond": bson.M{"else": 0, "if": bson.M{"$eq": []interface{}{"$fy.isCurrent", true}}, "then": bson.M{"$sum": bson.M{"$add": []interface{}{"$fy.rebate"}}}}}},
				"currentTax":         bson.M{"$sum": bson.M{"$cond": bson.M{"else": 0, "if": bson.M{"$eq": []interface{}{"$fy.isCurrent", true}}, "then": bson.M{"$sum": bson.M{"$add": []interface{}{"$fy.vacantLandTax", "$fy.tax"}}}}}},
				"arrearAlreadyPaid":  bson.M{"$sum": bson.M{"$cond": bson.M{"else": 0, "if": bson.M{"$eq": []interface{}{"$fy.isCurrent", false}}, "then": bson.M{"$sum": bson.M{"$add": []interface{}{"$fy.alreadyPayed.fyTax", "$fy.alreadyPayed.vlTax"}}}}}},
				"currentAlreadyPaid": bson.M{"$sum": bson.M{"$cond": bson.M{"else": 0, "if": bson.M{"$eq": []interface{}{"$fy.isCurrent", true}}, "then": bson.M{"$sum": bson.M{"$add": []interface{}{"$fy.alreadyPayed.fyTax", "$fy.alreadyPayed.vlTax"}}}}}},
				"totalTax":           bson.M{"$sum": "$fy.totalTax"},
				"otherDemand":        bson.M{"$sum": "$fy.otherDemand"},
				"fys":                bson.M{"$push": "$$ROOT"},
			}},
		},
	}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"paymentFys": bson.M{"$arrayElemAt": []interface{}{"$paymentFys", 0}}}})
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPROPERTYPAYMENTBASIC, "tnxId", "tnxId", "basic", "basic")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONWARD, "basic.property.address.wardCode", "code", "ref.ward", "ref.ward")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPROPERTYOWNER, "basic.property.uniqueId", "propertyId", "ref.owner", "ref.owner")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "basic.property.log.by.id", "userName", "ref.activator", "ref.activator")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "details.collector.id", "userName", "ref.collector", "ref.collector")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPROPERTY, "propertyId", "uniqueId", "ref.propertyDetails", "ref.propertyDetails")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "ref.propertyDetails.created.by", "userName", "ref.creator", "ref.creator")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("payment query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYPAYMENT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var data []models.RefCounterReportV2
	if err = cursor.All(context.TODO(), &data); err != nil {
		return nil, err
	}
	return data, nil
}

// FilterUserWisePropertyCollectionReport : ""
func (d *Daos) FilterUserWisePropertyCollectionReport(ctx *models.Context, filter *models.UserWisePropertyCollectionFilter, pagination *models.Pagination) ([]models.RefUserWisePropertyCollection, error) {
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
		if filter.DateFrom == nil {
			return nil, errors.New("please select a date")
		}
		if filter.DateFrom != nil {
			sd = time.Date(filter.DateFrom.Year(), filter.DateFrom.Month(), filter.DateFrom.Day(), 0, 0, 0, 0, filter.DateFrom.Location())
			// var ed time.Time
			if filter.DateTo != nil {
				ed = time.Date(filter.DateTo.Year(), filter.DateTo.Month(), filter.DateTo.Day(), 23, 59, 59, 0, filter.DateTo.Location())
			} else {
				ed = time.Date(filter.DateFrom.Year(), filter.DateFrom.Month(), filter.DateFrom.Day(), 23, 59, 59, 0, filter.DateFrom.Location())
			}
		}
	}
	fmt.Println("sd ===>", sd)
	fmt.Println("ed ===>", ed)
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
	mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{
		"from": constants.COLLECTIONPROPERTYPAYMENT,
		"as":   "payments",
		"let":  bson.M{"varUser": "$userName"},
		"pipeline": []bson.M{
			bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
				bson.M{"$eq": []string{"$status", constants.PROPERTYPAYMENTCOMPLETED}},
				bson.M{"$ne": []string{"$collectionReceived.status", "Collected"}},
				bson.M{"$ne": []string{"$collectionReceived.status", "Rejected"}},
				bson.M{"$eq": []string{"$$varUser", "$details.collector.id"}},
				// bson.M{"$eq": []string{"$$varUser", "$details.collector.id"}},
				bson.M{"$gte": []interface{}{"$completionDate", sd}},
				bson.M{"$lte": []interface{}{"$completionDate", ed}},
			}}}},
			bson.M{"$group": bson.M{"_id": bson.M{"mop": "$details.mop.mode"},
				"noOfPayments":     bson.M{"$sum": 1},
				"totalAmount":      bson.M{"$sum": "$details.amount"},
				"propertypayments": bson.M{"$push": "$$ROOT"},
			}},
			bson.M{"$addFields": bson.M{"mode": "$_id.mop"}},
			bson.M{"$project": bson.M{"_id": 0, "k": "$mode", "v": "$$ROOT"}},
		},
	}})

	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"payments": bson.M{"$arrayToObject": "$payments"}}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"propertyCash": "$payments.Cash", "propertyCheque": "$payments.Cheque", "propertynb": "$payments.NB", "propertydd": "$payments.DD"}})

	mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{
		"from": constants.COLLECTIONTRADELICENSEPAYMENTS,
		"as":   "tlPayments",
		"let":  bson.M{"varUser": "$userName"},
		"pipeline": []bson.M{
			bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
				bson.M{"$eq": []string{"$status", constants.TRADELICENSEPAYMENRSTATUSCOMPLETED}},
				bson.M{"$ne": []string{"$collectionReceived.status", "Collected"}},
				bson.M{"$ne": []string{"$collectionReceived.status", "Rejected"}},
				bson.M{"$eq": []string{"$$varUser", "$details.collector.by"}},
				bson.M{"$gte": []interface{}{"$completionDate", sd}},
				bson.M{"$lte": []interface{}{"$completionDate", ed}},
			}}}},
			bson.M{"$group": bson.M{"_id": bson.M{"mop": "$details.mop.mode"},
				"noOfPayments": bson.M{"$sum": 1},
				"totalAmount":  bson.M{"$sum": "$details.amount"},
				"tlPayments":   bson.M{"$push": "$$ROOT"},
			}},
			bson.M{"$addFields": bson.M{"mode": "$_id.mop"}},
			bson.M{"$project": bson.M{"_id": 0, "k": "$mode", "v": "$$ROOT"}},
		},
	}})

	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"tlPayments": bson.M{"$arrayToObject": "$tlPayments"}}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"tlCash": "$tlPayments.Cash", "tlCheque": "$tlPayments.Cheque", "tlnb": "$tlPayments.NB", "tldd": "$tlPayments.DD"}})

	mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{
		"from": constants.COLLECTIONSHOPRENTPAYMENTS,
		"as":   "srPayments",
		"let":  bson.M{"varUser": "$userName"},
		"pipeline": []bson.M{
			bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
				bson.M{"$eq": []string{"$status", constants.SHOPRENTPAYMENTSTATUSCOMPLETED}},
				bson.M{"$ne": []string{"$collectionReceived.status", "Collected"}},
				bson.M{"$ne": []string{"$collectionReceived.status", "Rejected"}},
				bson.M{"$eq": []string{"$$varUser", "$details.collector.by"}},
				bson.M{"$gte": []interface{}{"$completionDate", sd}},
				bson.M{"$lte": []interface{}{"$completionDate", ed}},
			}}}},
			bson.M{"$group": bson.M{"_id": bson.M{"mop": "$details.mop.mode"},
				"noOfPayments": bson.M{"$sum": 1},
				"totalAmount":  bson.M{"$sum": "$details.amount"},
				"srPayments":   bson.M{"$push": "$$ROOT"},
			}},
			bson.M{"$addFields": bson.M{"mode": "$_id.mop"}},
			bson.M{"$project": bson.M{"_id": 0, "k": "$mode", "v": "$$ROOT"}},
		},
	}})

	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"srPayments": bson.M{"$arrayToObject": "$srPayments"}}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"srCash": "$srPayments.Cash", "srCheque": "$srPayments.Cheque", "srnb": "$srPayments.NB", "srdd": "$srPayments.DD"}})

	mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{
		"from": constants.COLLECTIONMOBILETOWERPAYMENTS,
		"as":   "mtPayments",
		"let":  bson.M{"varUser": "$userName"},
		"pipeline": []bson.M{
			bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
				bson.M{"$eq": []string{"$status", constants.MOBILETOWERPAYMENRSTATUSCOMPLETED}},
				bson.M{"$ne": []string{"$collectionReceived.status", "Collected"}},
				bson.M{"$ne": []string{"$collectionReceived.status", "Rejected"}},
				bson.M{"$eq": []string{"$$varUser", "$details.collector.by"}},
				bson.M{"$gte": []interface{}{"$completionDate", sd}},
				bson.M{"$lte": []interface{}{"$completionDate", ed}},
			}}}},
			bson.M{"$group": bson.M{"_id": bson.M{"mop": "$details.mop.mode"},
				"noOfPayments": bson.M{"$sum": 1},
				"totalAmount":  bson.M{"$sum": "$details.amount"},
				"mtPayments":   bson.M{"$push": "$$ROOT"},
			}},
			bson.M{"$addFields": bson.M{"mode": "$_id.mop"}},
			bson.M{"$project": bson.M{"_id": 0, "k": "$mode", "v": "$$ROOT"}},
		},
	}})

	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"mtPayments": bson.M{"$arrayToObject": "$mtPayments"}}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"mtCash": "$mtPayments.Cash", "mtCheque": "$mtPayments.Cheque", "mtnb": "$mtPayments.NB", "mtdd": "$mtPayments.DD"}})

	mainPipeline = append(mainPipeline, bson.M{"$project": bson.M{"payments": 0, "tlPayments": 0, "srPayments": 0, "mtPayments": 0}})
	//Aggregation
	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONUSER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var res []models.RefUserWisePropertyCollection
	if err = cursor.All(context.TODO(), &res); err != nil {
		return nil, err
	}
	return res, nil
}

// FilterZoneDayWisePropertyCollectionReport : ""
func (d *Daos) FilterZoneDayWisePropertyCollectionReport(ctx *models.Context, filter *models.TeamDayWisePropertyCollectionReportFilter, pagination *models.Pagination) ([]models.TeamDayWisePropertyCollectionReport, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.UserName) > 0 {
			query = append(query, bson.M{"managerId": bson.M{"$in": filter.UserName}})
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
		"from": constants.COLLECTIONPROPERTYPAYMENT,
		"as":   "report",
		"let":  bson.M{"userName": "$userName"},
		"pipeline": []bson.M{bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
			bson.M{"$eq": []interface{}{"$status", constants.PROPERTYPAYMENTCOMPLETED}},
			bson.M{"$eq": []interface{}{"$$userName", "$details.collector.id"}},

			bson.M{"$gte": []interface{}{"$completionDate", sd}},
			bson.M{"$lte": []interface{}{"$completionDate", ed}},
		}}}},
			bson.M{"$group": bson.M{"_id": "$propertyId",
				"totalNoPayments": bson.M{"$sum": 1},
				"totalCollection": bson.M{"$sum": "$details.amount"}}},
			bson.M{"$group": bson.M{"_id": nil,
				"totalNoProperties": bson.M{"$sum": 1}, "totalNoPayments": bson.M{"$sum": "$totalNoPayments"},
				"totalCollections": bson.M{"$sum": "$totalCollection"}}},
		},
	},
	},
		bson.M{"$addFields": bson.M{"report": bson.M{"$arrayElemAt": []interface{}{"$report", 0}}}},
	)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONUSER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var res []models.TeamDayWisePropertyCollectionReport
	if err = cursor.All(context.TODO(), &res); err != nil {
		return nil, err
	}
	return res, nil
}

// FilterHoldingWiseCollectionReportJSON : ""
func (d *Daos) FilterHoldingWiseCollectionReportJSON(ctx *models.Context, filter *models.PropertyPaymentFilter, pagination *models.Pagination) ([]models.RefHoldingWiseCollectionReport, error) {
	timedaosStart := time.Now()
	mainPipeline := []bson.M{}
	query := []bson.M{}
	// mainPipeline = append(mainPipeline, d.FilterPropertyPaymentQuery(ctx, filter)...)

	query = d.FilterPropertyPaymentQuery(ctx, filter)
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if filter.SortBy != "" {
		mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{filter.SortBy: filter.SortOrder}})
	} else {
		mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"_id": -1}})
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYPAYMENT).CountDocuments(ctx.CTX, func() bson.M {
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

	mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{
		"from": "propertypaymentfys",
		"as":   "paymentFys",
		"let":  bson.M{"varTnxId": "$tnxId"},
		"pipeline": []bson.M{
			bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
				bson.M{"$eq": []string{"$$varTnxId", "$tnxId"}},
			}}}},
			bson.M{"$group": bson.M{"_id": "$tnxId",
				"arrearPenalty":      bson.M{"$sum": bson.M{"$cond": bson.M{"else": 0, "if": bson.M{"$eq": []interface{}{"$fy.isCurrent", false}}, "then": bson.M{"$sum": bson.M{"$add": []interface{}{"$fy.penanty"}}}}}},
				"arrearRebate":       bson.M{"$sum": bson.M{"$cond": bson.M{"else": 0, "if": bson.M{"$eq": []interface{}{"$fy.isCurrent", false}}, "then": bson.M{"$sum": bson.M{"$add": []interface{}{"$fy.rebate"}}}}}},
				"arrearTax":          bson.M{"$sum": bson.M{"$cond": bson.M{"else": 0, "if": bson.M{"$eq": []interface{}{"$fy.isCurrent", false}}, "then": bson.M{"$sum": bson.M{"$add": []interface{}{"$fy.vacantLandTax", "$fy.tax"}}}}}},
				"currentPenalty":     bson.M{"$sum": bson.M{"$cond": bson.M{"else": 0, "if": bson.M{"$eq": []interface{}{"$fy.isCurrent", true}}, "then": bson.M{"$sum": bson.M{"$add": []interface{}{"$fy.penanty"}}}}}},
				"currentRebate":      bson.M{"$sum": bson.M{"$cond": bson.M{"else": 0, "if": bson.M{"$eq": []interface{}{"$fy.isCurrent", true}}, "then": bson.M{"$sum": bson.M{"$add": []interface{}{"$fy.rebate"}}}}}},
				"currentTax":         bson.M{"$sum": bson.M{"$cond": bson.M{"else": 0, "if": bson.M{"$eq": []interface{}{"$fy.isCurrent", true}}, "then": bson.M{"$sum": bson.M{"$add": []interface{}{"$fy.vacantLandTax", "$fy.tax"}}}}}},
				"arrearAlreadyPaid":  bson.M{"$sum": bson.M{"$cond": bson.M{"else": 0, "if": bson.M{"$eq": []interface{}{"$fy.isCurrent", false}}, "then": bson.M{"$sum": bson.M{"$add": []interface{}{"$fy.alreadyPayed.fyTax", "$fy.alreadyPayed.vlTax"}}}}}},
				"currentAlreadyPaid": bson.M{"$sum": bson.M{"$cond": bson.M{"else": 0, "if": bson.M{"$eq": []interface{}{"$fy.isCurrent", true}}, "then": bson.M{"$sum": bson.M{"$add": []interface{}{"$fy.alreadyPayed.fyTax", "$fy.alreadyPayed.vlTax"}}}}}},
				"totalTax":           bson.M{"$sum": "$fy.totalTax"},
				"otherDemand":        bson.M{"$sum": "$fy.otherDemand"},
				"fys":                bson.M{"$push": "$$ROOT"},
			}},
		},
	}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"paymentFys": bson.M{"$arrayElemAt": []interface{}{"$paymentFys", 0}}}})

	mainPipeline = append(mainPipeline, bson.M{
		"$group": bson.M{
			"_id":                bson.M{"propertyId": "$propertyId"},
			"arrearAlreadyPaid":  bson.M{"$sum": "$paymentFys.arrearAlreadyPaid"},
			"arrearPenalty":      bson.M{"$sum": "$paymentFys.arrearPenalty"},
			"arrearRebate":       bson.M{"$sum": "$paymentFys.arrearRebate"},
			"arrearTax":          bson.M{"$sum": "$paymentFys.arrearTax"},
			"currentAlreadyPaid": bson.M{"$sum": "$paymentFys.currentAlreadyPaid"},
			"currentPenalty":     bson.M{"$sum": "$paymentFys.currentPenalty"},
			"currentRebate":      bson.M{"$sum": "$paymentFys.currentRebate"},
			"currentTax":         bson.M{"$sum": "$paymentFys.currentTax"},
			"otherDemand":        bson.M{"$sum": "$paymentFys.otherDemand"},
			"totalTax":           bson.M{"$sum": "$paymentFys.totalTax"},
			"totalAmount":        bson.M{"$sum": "$details.amount"},
		},
	})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"propertyId": "$_id.propertyId"}})

	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPROPERTYPAYMENTBASIC, "propertyId", "property.uniqueId", "basic", "basic")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONWARD, "basic.property.address.wardCode", "code", "ref.ward", "ref.ward")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPROPERTYOWNER, "basic.property.uniqueId", "propertyId", "ref.owner", "ref.owner")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONROADTYPE, "basic.property.roadTypeId", "uniqueId", "ref.roadType", "ref.roadType")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPROPERTYTYPE, "basic.property.propertyTypeId", "uniqueId", "ref.propertyType", "ref.propertyType")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "basic.property.log.by.id", "userName", "ref.activator", "ref.activator")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPROPERTY, "propertyId", "uniqueId", "ref.propertyDetails", "ref.propertyDetails")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "ref.propertyDetails.created.by", "userName", "ref.creator", "ref.creator")...)
	// mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{
	// 	"payment.arrearAlreadyPaid":  "$arrearAlreadyPaid",
	// 	"payment.arrearPenalty":      "$arrearPenalty",
	// 	"payment.arrearRebate":       "$arrearRebate",
	// 	"payment.arrearTax":          "$arrearTax",
	// 	"payment.currentAlreadyPaid": "$currentAlreadyPaid",
	// 	"payment.currentPenalty":     "$currentPenalty",
	// 	"payment.currentRebate":      "$currentRebate",
	// 	"payment.currentTax":         "$currentTax",
	// 	"payment.otherDemand":        "$otherDemand",
	// 	"payment.totalTax":           "$totalTax",
	// 	"payment.totalAmount":        "$totalAmount",
	// }})
	//Aggregation
	d.Shared.BsonToJSONPrintTag("payment query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYPAYMENT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var data []models.RefHoldingWiseCollectionReport
	if err = cursor.All(context.TODO(), &data); err != nil {
		return nil, err
	}
	timeDaosEnd := time.Now()
	fmt.Printf("The call took %v to run.\n", timeDaosEnd.Sub(timedaosStart))
	return data, nil
}

// PropertyWiseDemandandCollectionV2JSON : ""
func (d *Daos) PropertyWiseDemandandCollectionV2JSON(ctx *models.Context, propertyfilter *models.PropertyFilter, pagination *models.Pagination) ([]models.ResPropertyWiseDemandandCollectionV2Report, error) {
	resFYs, err := d.GetCurrentFinancialYear(ctx)
	if err != nil {
		return nil, errors.New("error in getting current financial year" + err.Error())
	}
	mainPipeline := []bson.M{}
	query := []bson.M{}

	query = d.FilterPropertyQuery(ctx, propertyfilter)
	var sd, ed time.Time

	if propertyfilter.AppliedRange != nil {
		if propertyfilter.AppliedRange.From != nil {
			sd = time.Date(propertyfilter.AppliedRange.From.Year(), propertyfilter.AppliedRange.From.Month(), propertyfilter.AppliedRange.From.Day(), 0, 0, 0, 0, propertyfilter.AppliedRange.From.Location())
			if propertyfilter.AppliedRange.To != nil {
				ed = time.Date(propertyfilter.AppliedRange.To.Year(), propertyfilter.AppliedRange.To.Month(), propertyfilter.AppliedRange.To.Day(), 23, 59, 59, 0, propertyfilter.AppliedRange.To.Location())
			} else {
				ed = time.Date(propertyfilter.AppliedRange.From.Year(), propertyfilter.AppliedRange.From.Month(), propertyfilter.AppliedRange.From.Day(), 23, 59, 59, 0, propertyfilter.AppliedRange.From.Location())
			}
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if propertyfilter.SortBy != "" {
		mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{propertyfilter.SortBy: propertyfilter.SortOrder}})
	} else {
		mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"_id": -1}})
	}
	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONPROPERTY).CountDocuments(ctx.CTX, func() bson.M {
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

	//Lookups
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONWARD, "address.wardCode", "code", "ref.address.ward", "ref.address.ward")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPROPERTYOWNER, "uniqueId", "propertyId", "ref.propertyOwner", "ref.propertyOwner")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPROPERTYTYPE, "propertyTypeId", "uniqueId", "ref.propertyType", "ref.propertyType")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONROADTYPE, "roadTypeId", "uniqueId", "ref.roadType", "ref.roadType")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONOVERALLPROPERTYDEMAND, "uniqueId", "propertyId", "ref.demand", "ref.demand")...)

	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"from": "propertypayments",
			"as":   "ref.payments",
			"let":  bson.M{"varUniqueId": "$uniqueId"},
			"pipeline": []bson.M{bson.M{
				"$match": bson.M{
					"$expr": bson.M{
						"$and": []bson.M{
							bson.M{"$eq": []string{"$propertyId", "$$varUniqueId"}},
							bson.M{"$gte": []interface{}{"$completionDate", sd}},
							bson.M{"$lte": []interface{}{"$completionDate", ed}},
							// query1,
							// query2,
						}}},
			},

				bson.M{
					"$group": bson.M{
						"_id":     "$tnxId",
						"rebate":  bson.M{"$sum": "$demand.rebate"},
						"formFee": bson.M{"$sum": "$demand.formFee"},
					},
				},
				bson.M{
					"$lookup": bson.M{
						"from": "propertypaymentfys",
						"as":   "paymentFys",
						"let":  bson.M{"varTnxId": "$_id"},
						"pipeline": []bson.M{bson.M{
							"$match": bson.M{
								"$expr": bson.M{
									"$and": []bson.M{
										bson.M{"$eq": []string{"$tnxId", "$$varTnxId"}},
										bson.M{"$eq": []string{"$status", "Completed"}},
									}}},
						},

							bson.M{"$group": bson.M{"_id": "$propertyId",
								"arrearTax": bson.M{"$sum": bson.M{"$cond": bson.M{"else": 0,
									"if":   bson.M{"$ne": []interface{}{"$fy.uniqueId", resFYs.UniqueID}},
									"then": bson.M{"$sum": bson.M{"$add": []interface{}{"$fy.vacantLandTax", "$fy.tax"}}}}}},
								"currentTax": bson.M{"$sum": bson.M{"$cond": bson.M{"else": 0,
									"if":   bson.M{"$eq": []interface{}{"$fy.uniqueId", resFYs.UniqueID}},
									"then": bson.M{"$sum": bson.M{"$add": []interface{}{"$fy.vacantLandTax", "$fy.tax"}}}}}},
								"arrearPenalty": bson.M{"$sum": bson.M{"$cond": bson.M{"else": 0,
									"if":   bson.M{"$ne": []interface{}{"$fy.uniqueId", resFYs.UniqueID}},
									"then": bson.M{"$sum": bson.M{"$add": []interface{}{"$fy.penanty"}}}}}},
								"currentPenalty": bson.M{"$sum": bson.M{"$cond": bson.M{"else": 0,
									"if":   bson.M{"$eq": []interface{}{"$fy.uniqueId", resFYs.UniqueID}},
									"then": bson.M{"$sum": bson.M{"$add": []interface{}{"$fy.penanty"}}}}}},
								"arrearRebate": bson.M{"$sum": bson.M{"$cond": bson.M{"else": 0,
									"if":   bson.M{"$ne": []interface{}{"$fy.uniqueId", resFYs.UniqueID}},
									"then": bson.M{"$sum": bson.M{"$add": []interface{}{"$fy.rebate"}}}}}},
								"currentRebate": bson.M{"$sum": bson.M{"$cond": bson.M{"else": 0,
									"if":   bson.M{"$eq": []interface{}{"$fy.uniqueId", resFYs.UniqueID}},
									"then": bson.M{"$sum": bson.M{"$add": []interface{}{"$fy.rebate"}}}}}},
								"arrearAlreadyPaid": bson.M{"$sum": bson.M{"$cond": bson.M{"else": 0,
									"if":   bson.M{"$ne": []interface{}{"$fy.uniqueId", resFYs.UniqueID}},
									"then": bson.M{"$sum": bson.M{"$add": []interface{}{"$fy.alreadyPayed.fyTax", "$fy.alreadyPayed.vlTax"}}}}}},
								"currentAlreadyPaid": bson.M{"$sum": bson.M{"$cond": bson.M{"else": 0,
									"if":   bson.M{"$eq": []interface{}{"$fy.uniqueId", resFYs.UniqueID}},
									"then": bson.M{"$sum": bson.M{"$add": []interface{}{"$fy.alreadyPayed.fyTax", "$fy.alreadyPayed.vlTax"}}}}}},
								"totalTax":    bson.M{"$sum": "$fy.totalTax"},
								"otherDemand": bson.M{"$sum": "$fy.otherDemand"},
							}},
						},
					},
				},
				bson.M{"$addFields": bson.M{"paymentFys": bson.M{"$arrayElemAt": []interface{}{"$paymentFys", 0}}}}}},
	},
		bson.M{"$addFields": bson.M{"ref.payments": bson.M{"$arrayElemAt": []interface{}{"$ref.payments", 0}}}},
	)
	// mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"payments": bson.M{"$arrayElemAt": []interface{}{"$payments", 0}}}})
	//Aggregation
	d.Shared.BsonToJSONPrintTag("property query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var propertys []models.ResPropertyWiseDemandandCollectionV2Report
	if err = cursor.All(context.TODO(), &propertys); err != nil {
		return nil, err
	}
	return propertys, nil
}

// PropertyDemandAndCollectionReportJSON : ""
func (d *Daos) PropertyDemandAndCollectionReportJSON(ctx *models.Context, filter *models.PropertyFilter, pagination *models.Pagination) ([]models.RefPropertyDemandAndCollectionReport, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}

	resFYs, err := d.GetSingleFinancialYearUsingDateV2(ctx, filter.Date)
	if err != nil {
		return nil, errors.New("error in getting current financial year" + err.Error())
	}

	query = d.FilterPropertyQuery(ctx, filter)
	// var sd, ed time.Time

	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if filter.SortBy != "" {
		mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{filter.SortBy: filter.SortOrder}})
	} else {
		mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"_id": -1}})
	}
	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONPROPERTY).CountDocuments(ctx.CTX, func() bson.M {
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

	//Lookups
	mainPipeline = append(mainPipeline, bson.M{"$project": bson.M{"uniqueId": 1, "propertyTypeId": 1}})
	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"from": "propertypayments",
			"as":   "collections",
			"let": bson.M{
				"propertyId": "$uniqueId",
			},
			"pipeline": []bson.M{
				bson.M{
					"$match": bson.M{
						"$expr": bson.M{
							"$and": []bson.M{
								bson.M{
									"$eq": []string{"$status", "Completed"},
								},
								bson.M{
									"$eq": []string{
										"$propertyId",
										"$$propertyId",
									},
								},
								bson.M{"$gte": []interface{}{"$completionDate", resFYs.From}},
								bson.M{"$lte": []interface{}{"$completionDate", resFYs.To}},
							},
						},
					},
				},
				bson.M{
					"$lookup": bson.M{
						"from": "propertypaymentfys",
						"as":   "fys",
						"let": bson.M{
							"tnxId": "$tnxId",
						},
						"pipeline": []bson.M{
							bson.M{
								"$match": bson.M{
									"$expr": bson.M{
										"$and": []bson.M{
											bson.M{
												"$eq": []string{
													"$tnxId",
													"$$tnxId",
												},
											},
										},
									},
								},
							},
							bson.M{
								"$group": bson.M{
									"_id": "$fy.isCurrent",
									"tax": bson.M{
										"$sum": bson.M{
											"$add": []interface{}{
												"$fy.vacantLandTax",
												"$fy.tax",
											},
										},
									},
									"rebate": bson.M{
										"$sum": "$fy.rebate",
									},
									"penalty": bson.M{
										"$sum": "$fy.penanty",
									},
									"otherDemand": bson.M{
										"$sum": "$fy.otherDemand",
									},
								},
							},
							bson.M{
								"$group": bson.M{
									"_id": nil,
									"arrearTax": bson.M{"$sum": bson.M{
										"$cond": bson.M{
											"if": bson.M{
												"$eq": []interface{}{
													"$_id",
													false,
												},
											},
											"then": "$tax",
											"else": 0,
										},
									}},
									"arrearRebate": bson.M{"$sum": bson.M{
										"$cond": bson.M{
											"if": bson.M{
												"$eq": []interface{}{
													"$_id",
													false,
												},
											},
											"then": "$rebate",
											"else": 0,
										},
									}},
									"arrearPenalty": bson.M{"$sum": bson.M{
										"$cond": bson.M{
											"if": bson.M{
												"$eq": []interface{}{
													"$_id",
													false,
												},
											},
											"then": "$penalty",
											"else": 0,
										},
									}},
									"currentTax": bson.M{"$sum": bson.M{
										"$cond": bson.M{
											"if": bson.M{
												"$eq": []interface{}{
													"$_id",
													true,
												},
											},
											"then": "$tax",
											"else": 0,
										},
									}},
									"currentRebate": bson.M{"$sum": bson.M{
										"$cond": bson.M{
											"if": bson.M{
												"$eq": []interface{}{
													"$_id",
													true,
												},
											},
											"then": "$rebate",
											"else": 0,
										},
									}},
									"currentPenalty": bson.M{"$sum": bson.M{
										"$cond": bson.M{
											"if": bson.M{
												"$eq": []interface{}{
													"$_id",
													true,
												},
											},
											"then": "$penalty",
											"else": 0,
										},
									}},
									"currentOtherDemand": bson.M{"$sum": bson.M{
										"$cond": bson.M{
											"if": bson.M{
												"$eq": []interface{}{
													"$_id",
													true,
												},
											},
											"then": "$otherDemand",
											"else": 0,
										},
									}},
									"arrearOtherDemand": bson.M{"$sum": bson.M{
										"$cond": bson.M{
											"if": bson.M{
												"$eq": []interface{}{
													"$_id",
													false,
												},
											},
											"then": "$otherDemand",
											"else": 0,
										},
									}},
								},
							},
						},
					},
				},
				bson.M{
					"$addFields": bson.M{
						"fys": bson.M{"$arrayElemAt": []interface{}{"$fys", 0}},
					},
				},
				bson.M{"$group": bson.M{
					"_id":                nil,
					"arrearTax":          bson.M{"$sum": "$fys.arrearTax"},
					"arrearRebate":       bson.M{"$sum": "$fys.arrearRebate"},
					"arrearPenalty":      bson.M{"$sum": "$fys.arrearPenalty"},
					"currentTax":         bson.M{"$sum": "$fys.currentTax"},
					"currentRebate":      bson.M{"$sum": "$fys.currentRebate"},
					"currentPenalty":     bson.M{"$sum": "$fys.currentPenalty"},
					"arrearOtherDemand":  bson.M{"$sum": "$fys.arrearOtherDemand"},
					"currentOtherDemand": bson.M{"$sum": "$fys.currentOtherDemand"},
					"boreCharge":         bson.M{"$sum": "$demand.boreCharge"},
					"formFee":            bson.M{"$sum": "$demand.formFee"},
					// "formFee":        bson.M{"$sum": "$demand.formFee"},
				}},
			},
		},
	},
		bson.M{"$addFields": bson.M{
			"collections": bson.M{"$arrayElemAt": []interface{}{"$collections", 0}},
		}})
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONWARD, "address.wardCode", "code", "ref.address.ward", "ref.address.ward")...)
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPROPERTYOWNER, "uniqueId", "propertyId", "ref.propertyOwner", "ref.propertyOwner")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPROPERTYTYPE, "propertyTypeId", "uniqueId", "propertyType", "propertyType")...)
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONROADTYPE, "roadTypeId", "uniqueId", "ref.roadType", "ref.roadType")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONOVERALLPROPERTYDEMAND, "uniqueId", "propertyId", "demand", "demand")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("property query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var propertys []models.RefPropertyDemandAndCollectionReport
	if err = cursor.All(context.TODO(), &propertys); err != nil {
		return nil, err
	}
	return propertys, nil
}

func (d *Daos) FilterPaymentCOllection(ctx *models.Context, filter *models.PropertyPaymentFilter, pagination *models.Pagination) ([]models.RefPropertyCollectionReport, error) {
	timedaosStart := time.Now()
	mainPipeline := []bson.M{}
	query := []bson.M{}
	// mainPipeline = append(mainPipeline, d.FilterPropertyPaymentQuery(ctx, filter)...)

	query = d.FilterPropertyPaymentQuery(ctx, filter)
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if filter.SortBy != "" {
		mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{filter.SortBy: filter.SortOrder}})
	} else {
		mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"completionDate": 1}})
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYPAYMENT).CountDocuments(ctx.CTX, func() bson.M {
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

	mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{
		"from": "propertypaymentfys",
		"as":   "paymentFys",
		"let":  bson.M{"varTnxId": "$tnxId"},
		"pipeline": []bson.M{
			bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
				bson.M{"$eq": []string{"$$varTnxId", "$tnxId"}},
				//bson.M{"$eq": []interface{}{"$fy.isCurrent", false}},
			}}}},
			bson.M{"$group": bson.M{"_id": "$tnxId",
				"arrearPenalty":      bson.M{"$sum": bson.M{"$cond": bson.M{"else": 0, "if": bson.M{"$eq": []interface{}{"$fy.isCurrent", false}}, "then": bson.M{"$sum": bson.M{"$add": []interface{}{"$fy.penanty"}}}}}},
				"arrearRebate":       bson.M{"$sum": bson.M{"$cond": bson.M{"else": 0, "if": bson.M{"$eq": []interface{}{"$fy.isCurrent", false}}, "then": bson.M{"$sum": bson.M{"$add": []interface{}{"$fy.rebate"}}}}}},
				"arrearTax":          bson.M{"$sum": bson.M{"$cond": bson.M{"else": 0, "if": bson.M{"$eq": []interface{}{"$fy.isCurrent", false}}, "then": bson.M{"$sum": bson.M{"$add": []interface{}{"$fy.vacantLandTax", "$fy.tax"}}}}}},
				"currentPenalty":     bson.M{"$sum": bson.M{"$cond": bson.M{"else": 0, "if": bson.M{"$eq": []interface{}{"$fy.isCurrent", true}}, "then": bson.M{"$sum": bson.M{"$add": []interface{}{"$fy.penanty"}}}}}},
				"currentRebate":      bson.M{"$sum": bson.M{"$cond": bson.M{"else": 0, "if": bson.M{"$eq": []interface{}{"$fy.isCurrent", true}}, "then": bson.M{"$sum": bson.M{"$add": []interface{}{"$fy.rebate"}}}}}},
				"currentTax":         bson.M{"$sum": bson.M{"$cond": bson.M{"else": 0, "if": bson.M{"$eq": []interface{}{"$fy.isCurrent", true}}, "then": bson.M{"$sum": bson.M{"$add": []interface{}{"$fy.vacantLandTax", "$fy.tax"}}}}}},
				"arrearAlreadyPaid":  bson.M{"$sum": bson.M{"$cond": bson.M{"else": 0, "if": bson.M{"$eq": []interface{}{"$fy.isCurrent", false}}, "then": bson.M{"$sum": bson.M{"$add": []interface{}{"$fy.alreadyPayed.fyTax", "$fy.alreadyPayed.vlTax"}}}}}},
				"currentAlreadyPaid": bson.M{"$sum": bson.M{"$cond": bson.M{"else": 0, "if": bson.M{"$eq": []interface{}{"$fy.isCurrent", true}}, "then": bson.M{"$sum": bson.M{"$add": []interface{}{"$fy.alreadyPayed.fyTax", "$fy.alreadyPayed.vlTax"}}}}}},
				"totalTax":           bson.M{"$sum": "$fy.totalTax"},
				"otherDemand":        bson.M{"$sum": "$fy.otherDemand"},
				"fys":                bson.M{"$push": "$$ROOT"},
			}},
		},
	}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"paymentFys": bson.M{"$arrayElemAt": []interface{}{"$paymentFys", 0}}}})
	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"as":           "property",
			"foreignField": "uniqueId",
			"from":         "properties",
			"localField":   "propertyId",
		}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"property": bson.M{"$arrayElemAt": []interface{}{"$property", 0}}}})
	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"as":           "owners",
			"foreignField": "propertyId",
			"from":         "owners",
			"localField":   "propertyId",
		}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"owners": bson.M{"$arrayElemAt": []interface{}{"$owners", 0}}}})
	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"as":           "activator",
			"foreignField": "userName",
			"from":         "users",
			"localField":   "property.log.by.id",
		}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"activator": bson.M{"$arrayElemAt": []interface{}{"$activator", 0}}}})
	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"as":           "creator",
			"foreignField": "userName",
			"from":         "users",
			"localField":   "property.created.by",
		}})

	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"creator": bson.M{"$arrayElemAt": []interface{}{"$creator", 0}}}})
	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"as":           "collector",
			"foreignField": "userName",
			"from":         "users",
			"localField":   "details.collector.id",
		}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"collector": bson.M{"$arrayElemAt": []interface{}{"$collector", 0}}}})
	mainPipeline = append(mainPipeline, bson.M{"$project": bson.M{
		"_id":                           false,
		"paymentFys.arrearAlreadyPaid":  true,
		"paymentFys.arrearTax":          true,
		"paymentFys.arrearPenalty":      true,
		"paymentFys.arrearRebate":       true,
		"paymentFys.currentAlreadyPaid": true,
		"paymentFys.currentPenalty":     true,
		"paymentFys.currentRebate":      true,
		"paymentFys.currentTax":         true,
		"paymentFys.otherDemand":        true,
		"paymentFys.totalTax":           true,
		"paymentFys.fys.fy.name":        true,
		"address.al1":                   true,
		"address.al2":                   true,
		"demand.formFee":                true,
		"demand.previousCollection":     true,
		"demand.boreCharge":             true,
		"demand.otherDemand":            true,
		"details.madeAt.at":             true,
		"details.amount":                true,
		"details.mop":                   true,
		"property.applicationNo":        true,
		"property.oldHoldingNumber":     true,
		"property.address.wardCode":     true,
		"completionDate":                true,
		"reciptNo":                      true,
		"propertyId":                    true,
		"tnxId":                         true,
		"status":                        true,
		"collector.name":                true,
		"owners.name":                   true,
		"owners.mobile":                 true,
		"activator.name":                true,
		"creator.name":                  true,
	}})

	//	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPROPERTYOWNER, "propertyId", "propertyId", "ref.owner", "ref.owner")...)
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "property.log.by.id", "userName", "ref.activator", "ref.activator")...)
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "details.collector.id", "userName", "ref.collector", "ref.collector")...)
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPROPERTY, "propertyId", "uniqueId", "ref.propertyDetails", "ref.propertyDetails")...)
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "property.created.by", "userName", "ref.creator", "ref.creator")...)
	//Aggregation
	d.Shared.BsonToJSONPrintTag("payment query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYPAYMENT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var data []models.RefPropertyCollectionReport
	if err = cursor.All(context.TODO(), &data); err != nil {
		return nil, err
	}
	timeDaosEnd := time.Now()
	fmt.Printf("The call took %v to run.\n", timeDaosEnd.Sub(timedaosStart))
	return data, nil
}

func (d *Daos) FilterPaymentSummary(ctx *models.Context) ([]models.Summary, error) {
	timedaosStart := time.Now()
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{
		"from": "propertypaymentfys",
		"as":   "paymentFys",
		"let":  bson.M{"varTnxId": "$tnxId"},
		"pipeline": []bson.M{
			bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
				bson.M{"$eq": []string{"$$varTnxId", "$tnxId"}},
			}}}},
			bson.M{"$sort": bson.M{"fy.order": 1}},
			bson.M{"$group": bson.M{"_id": "$tnxId",
				"arrearPenalty":      bson.M{"$sum": bson.M{"$cond": bson.M{"else": 0, "if": bson.M{"$eq": []interface{}{"$fy.isCurrent", false}}, "then": bson.M{"$sum": bson.M{"$add": []interface{}{"$fy.penanty"}}}}}},
				"arrearRebate":       bson.M{"$sum": bson.M{"$cond": bson.M{"else": 0, "if": bson.M{"$eq": []interface{}{"$fy.isCurrent", false}}, "then": bson.M{"$sum": bson.M{"$add": []interface{}{"$fy.rebate"}}}}}},
				"arrearTax":          bson.M{"$sum": bson.M{"$cond": bson.M{"else": 0, "if": bson.M{"$eq": []interface{}{"$fy.isCurrent", false}}, "then": bson.M{"$sum": bson.M{"$add": []interface{}{"$fy.vacantLandTax", "$fy.tax"}}}}}},
				"currentPenalty":     bson.M{"$sum": bson.M{"$cond": bson.M{"else": 0, "if": bson.M{"$eq": []interface{}{"$fy.isCurrent", true}}, "then": bson.M{"$sum": bson.M{"$add": []interface{}{"$fy.penanty"}}}}}},
				"currentRebate":      bson.M{"$sum": bson.M{"$cond": bson.M{"else": 0, "if": bson.M{"$eq": []interface{}{"$fy.isCurrent", true}}, "then": bson.M{"$sum": bson.M{"$add": []interface{}{"$fy.rebate"}}}}}},
				"currentTax":         bson.M{"$sum": bson.M{"$cond": bson.M{"else": 0, "if": bson.M{"$eq": []interface{}{"$fy.isCurrent", true}}, "then": bson.M{"$sum": bson.M{"$add": []interface{}{"$fy.vacantLandTax", "$fy.tax"}}}}}},
				"arrearOtherDemand":  bson.M{"$sum": bson.M{"$cond": bson.M{"else": 0, "if": bson.M{"$eq": []interface{}{"$fy.isCurrent", false}}, "then": bson.M{"$sum": "$fy.otherDemand"}}}},
				"currentOtherDemand": bson.M{"$sum": bson.M{"$cond": bson.M{"else": 0, "if": bson.M{"$eq": []interface{}{"$fy.isCurrent", true}}, "then": bson.M{"$sum": "$fy.otherDemand"}}}},
				"fromFy":             bson.M{"$first": "$fy.name"},
				"toFy":               bson.M{"$last": "$fy.name"},
			}},
		},
	}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"paymentFys": bson.M{"$arrayElemAt": []interface{}{"$paymentFys", 0}}}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{
		"paymentFys.boreCharge": "$demand.boreCharge",
		"paymentFys.formFee":    "$demand.formFee",
	}})
	mainPipeline = append(mainPipeline, bson.M{"$project": bson.M{"paymentFys": 1}})
	mainPipeline = append(mainPipeline, bson.M{"$replaceRoot": bson.M{"newRoot": "$paymentFys"}})

	//Aggregation
	d.Shared.BsonToJSONPrintTag("payment query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYPAYMENT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var data []models.Summary
	if err = cursor.All(context.TODO(), &data); err != nil {
		return nil, err
	}
	timeDaosEnd := time.Now()
	fmt.Printf("The call took %v to run.\n", timeDaosEnd.Sub(timedaosStart))
	return data, nil
}

func (d *Daos) UpdatePaymentSummary(ctx *models.Context, data models.Summary) error {
	query := bson.M{"tnxId": data.TnxID}
	update := bson.M{"$set": bson.M{"summary": data}}
	rs, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYPAYMENT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return err
	}
	fmt.Println(rs)
	return nil
}

func (d *Daos) FilterHoldingWiseCollectionReportJSONV2(ctx *models.Context, filter *models.PropertyPaymentFilter, pagination *models.Pagination) ([]models.RefHoldingWiseCollectionReportV2, error) {
	timedaosStart := time.Now()
	mainPipeline := []bson.M{}
	query := []bson.M{}
	// mainPipeline = append(mainPipeline, d.FilterPropertyPaymentQuery(ctx, filter)...)

	query = d.FilterPropertyPaymentQuery(ctx, filter)
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if filter.SortBy != "" {
		mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{filter.SortBy: filter.SortOrder}})
	} else {
		mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"_id": -1}})
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYPAYMENT).CountDocuments(ctx.CTX, func() bson.M {
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

	mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{
		"from": "propertypaymentfys",
		"as":   "paymentFys",
		"let":  bson.M{"varTnxId": "$tnxId"},
		"pipeline": []bson.M{
			bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
				bson.M{"$eq": []string{"$$varTnxId", "$tnxId"}},
			}}}},
			bson.M{"$group": bson.M{"_id": "$tnxId",
				"arrearPenalty":      bson.M{"$sum": bson.M{"$cond": bson.M{"else": 0, "if": bson.M{"$eq": []interface{}{"$fy.isCurrent", false}}, "then": bson.M{"$sum": bson.M{"$add": []interface{}{"$fy.penanty"}}}}}},
				"arrearRebate":       bson.M{"$sum": bson.M{"$cond": bson.M{"else": 0, "if": bson.M{"$eq": []interface{}{"$fy.isCurrent", false}}, "then": bson.M{"$sum": bson.M{"$add": []interface{}{"$fy.rebate"}}}}}},
				"arrearTax":          bson.M{"$sum": bson.M{"$cond": bson.M{"else": 0, "if": bson.M{"$eq": []interface{}{"$fy.isCurrent", false}}, "then": bson.M{"$sum": bson.M{"$add": []interface{}{"$fy.vacantLandTax", "$fy.tax"}}}}}},
				"currentPenalty":     bson.M{"$sum": bson.M{"$cond": bson.M{"else": 0, "if": bson.M{"$eq": []interface{}{"$fy.isCurrent", true}}, "then": bson.M{"$sum": bson.M{"$add": []interface{}{"$fy.penanty"}}}}}},
				"currentRebate":      bson.M{"$sum": bson.M{"$cond": bson.M{"else": 0, "if": bson.M{"$eq": []interface{}{"$fy.isCurrent", true}}, "then": bson.M{"$sum": bson.M{"$add": []interface{}{"$fy.rebate"}}}}}},
				"currentTax":         bson.M{"$sum": bson.M{"$cond": bson.M{"else": 0, "if": bson.M{"$eq": []interface{}{"$fy.isCurrent", true}}, "then": bson.M{"$sum": bson.M{"$add": []interface{}{"$fy.vacantLandTax", "$fy.tax"}}}}}},
				"arrearAlreadyPaid":  bson.M{"$sum": bson.M{"$cond": bson.M{"else": 0, "if": bson.M{"$eq": []interface{}{"$fy.isCurrent", false}}, "then": bson.M{"$sum": bson.M{"$add": []interface{}{"$fy.alreadyPayed.fyTax", "$fy.alreadyPayed.vlTax"}}}}}},
				"currentAlreadyPaid": bson.M{"$sum": bson.M{"$cond": bson.M{"else": 0, "if": bson.M{"$eq": []interface{}{"$fy.isCurrent", true}}, "then": bson.M{"$sum": bson.M{"$add": []interface{}{"$fy.alreadyPayed.fyTax", "$fy.alreadyPayed.vlTax"}}}}}},
				"totalTax":           bson.M{"$sum": "$fy.totalTax"},
				"otherDemand":        bson.M{"$sum": "$fy.otherDemand"},
				"fys":                bson.M{"$push": "$$ROOT"},
			}},
		},
	}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"paymentFys": bson.M{"$arrayElemAt": []interface{}{"$paymentFys", 0}}}})

	mainPipeline = append(mainPipeline, bson.M{
		"$group": bson.M{
			"_id":                bson.M{"propertyId": "$propertyId"},
			"arrearAlreadyPaid":  bson.M{"$sum": "$paymentFys.arrearAlreadyPaid"},
			"arrearPenalty":      bson.M{"$sum": "$paymentFys.arrearPenalty"},
			"arrearRebate":       bson.M{"$sum": "$paymentFys.arrearRebate"},
			"arrearTax":          bson.M{"$sum": "$paymentFys.arrearTax"},
			"currentAlreadyPaid": bson.M{"$sum": "$paymentFys.currentAlreadyPaid"},
			"currentPenalty":     bson.M{"$sum": "$paymentFys.currentPenalty"},
			"currentRebate":      bson.M{"$sum": "$paymentFys.currentRebate"},
			"currentTax":         bson.M{"$sum": "$paymentFys.currentTax"},
			"otherDemand":        bson.M{"$sum": "$paymentFys.otherDemand"},
			"totalTax":           bson.M{"$sum": "$paymentFys.totalTax"},
			"totalAmount":        bson.M{"$sum": "$details.amount"},
		},
	})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"propertyId": "$_id.propertyId"}})

	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPROPERTYPAYMENTBASIC, "propertyId", "property.uniqueId", "basic", "basic")...)
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONWARD, "basic.property.address.wardCode", "code", "ref.ward", "ref.ward")...)
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPROPERTYOWNER, "basic.property.uniqueId", "propertyId", "ref.owner", "ref.owner")...)
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONROADTYPE, "basic.property.roadTypeId", "uniqueId", "ref.roadType", "ref.roadType")...)
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPROPERTYTYPE, "basic.property.propertyTypeId", "uniqueId", "ref.propertyType", "ref.propertyType")...)
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "basic.property.log.by.id", "userName", "ref.activator", "ref.activator")...)
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPROPERTY, "propertyId", "uniqueId", "ref.propertyDetails", "ref.propertyDetails")...)
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "ref.propertyDetails.created.by", "userName", "ref.creator", "ref.creator")...)
	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"as":           "basic",
			"foreignField": "uniqueId",
			"from":         "properties",
			"localField":   "propertyId",
		}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"basic": bson.M{"$arrayElemAt": []interface{}{"$basic", 0}}}})
	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"as":           "owners",
			"foreignField": "propertyId",
			"from":         "owners",
			"localField":   "propertyId",
		}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"owners": bson.M{"$arrayElemAt": []interface{}{"$owners", 0}}}})
	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"as":           "ward",
			"foreignField": "code",
			"from":         "wards",
			"localField":   "basic.address.wardCode",
		}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"ward": bson.M{"$arrayElemAt": []interface{}{"$ward", 0}}}})

	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"as":           "roadType",
			"foreignField": "uniqueId",
			"from":         "roadtypes",
			"localField":   "basic.roadTypeId",
		}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"roadType": bson.M{"$arrayElemAt": []interface{}{"$roadType", 0}}}})

	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"as":           "propertyType",
			"foreignField": "uniqueId",
			"from":         "propertytypes",
			"localField":   "basic.propertyTypeId",
		}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"propertyType": bson.M{"$arrayElemAt": []interface{}{"$propertyType", 0}}}})

	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"as":           "activator",
			"foreignField": "userName",
			"from":         "users",
			"localField":   "basic.log.by.id",
		}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"activator": bson.M{"$arrayElemAt": []interface{}{"$activator", 0}}}})
	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"as":           "property",
			"foreignField": "uniqueId",
			"from":         "properties",
			"localField":   "propertyId",
		}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"property": bson.M{"$arrayElemAt": []interface{}{"$property", 0}}}})
	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"as":           "creator",
			"foreignField": "userName",
			"from":         "users",
			"localField":   "property.created.by",
		}})

	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"creator": bson.M{"$arrayElemAt": []interface{}{"$creator", 0}}}})
	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"as":           "collector",
			"foreignField": "userName",
			"from":         "users",
			"localField":   "details.collector.id",
		}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"collector": bson.M{"$arrayElemAt": []interface{}{"$collector", 0}}}})
	mainPipeline = append(mainPipeline, bson.M{
		"$addFields": bson.M{
			"payment.arrearAlreadyPaid":  "$arrearAlreadyPaid",
			"payment.arrearPenalty":      "$arrearPenalty",
			"payment.arrearRebate":       "$arrearRebate",
			"payment.arrearTax":          "$arrearTax",
			"payment.currentAlreadyPaid": "$currentAlreadyPaid",
			"payment.currentPenalty":     "$currentPenalty",
			"payment.currentRebate":      "$currentRebate",
			"payment.currentTax":         "$currentTax",
			"payment.otherDemand":        "$otherDemand",
			"payment.totalAmount":        "$totalAmount",
			"payment.totalTax":           "$totalTax",
		},
	})
	mainPipeline = append(mainPipeline, bson.M{"$project": bson.M{
		"_id":                       false,
		"propertyId":                true,
		"payment":                   true,
		"basic.applicationNo":       true,
		"basic.oldHoldingNumber":    true,
		"basic.address.al1":         true,
		"basic.address.al2":         true,
		"ward.code":                 true,
		"owners.mobile":             true,
		"owners.name":               true,
		"owners.fatherRpanRhusband": true,
		"roadType.name":             true,
		"propertyType.name":         true,
		"activator.name":            true,
		"creator.name":              true,
	}})
	//Aggregation
	d.Shared.BsonToJSONPrintTag("payment query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYPAYMENT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var data []models.RefHoldingWiseCollectionReportV2
	if err = cursor.All(context.TODO(), &data); err != nil {
		return nil, err
	}
	timeDaosEnd := time.Now()
	fmt.Printf("The call took %v to run.\n", timeDaosEnd.Sub(timedaosStart))
	return data, nil
}

func (d *Daos) PropertyWiseDemandandCollectionExcelV2(ctx *models.Context, propertyfilter *models.PropertyFilter, pagination *models.Pagination) ([]models.RefPropertyV2, error) {
	resFYs, err := d.GetCurrentFinancialYear(ctx)
	if err != nil {
		return nil, errors.New("error in getting current financial year" + err.Error())
	}
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if propertyfilter != nil {
		if propertyfilter.IsLocation {
			mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"address.location.coordinates": bson.M{"$exists": true}}})
			mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"locSize": bson.M{"$size": "$address.location.coordinates"}}})
		}
	}

	query = d.FilterPropertyQuery(ctx, propertyfilter)
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if propertyfilter.SortBy != "" {
		mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{propertyfilter.SortBy: propertyfilter.SortOrder}})
	} else {
		mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"_id": -1}})
	}
	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONPROPERTY).CountDocuments(ctx.CTX, func() bson.M {
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

	//Lookups
	if propertyfilter != nil {
		if !propertyfilter.RemoveLookup.PropertyOwner {
			mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPROPERTYOWNER, "uniqueId", "propertyId", "ref.propertyOwner", "ref.propertyOwner")...)
		}
		if !propertyfilter.RemoveLookup.State {
			mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "address.stateCode", "code", "ref.address.state", "ref.address.state")...)
		}
		if !propertyfilter.RemoveLookup.District {
			mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "address.districtCode", "code", "ref.address.district", "ref.address.district")...)
		}
		if !propertyfilter.RemoveLookup.Village {
			mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVILLAGE, "address.villageCode", "code", "ref.address.village", "ref.address.village")...)
		}
		if !propertyfilter.RemoveLookup.Zone {
			mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONZONE, "address.zoneCode", "code", "ref.address.zone", "ref.address.zone")...)
		}
		if !propertyfilter.RemoveLookup.Ward {
			mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONWARD, "address.wardCode", "code", "ref.address.ward", "ref.address.ward")...)
		}
		if !propertyfilter.RemoveLookup.PropertyFloor {
			mainPipeline = append(mainPipeline, d.PropertyFloorsLookup(constants.COLLECTIONPROPERTYFLOOR, "uniqueId", "propertyId", "ref.floors", "ref.floors")...)
		}
		if !propertyfilter.RemoveLookup.PropertyType {
			mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPROPERTYTYPE, "propertyTypeId", "uniqueId", "ref.propertyType", "ref.propertyType")...)
		}
		if !propertyfilter.RemoveLookup.RoadType {
			mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONROADTYPE, "roadTypeId", "uniqueId", "ref.roadType", "ref.roadType")...)
		}
		if !propertyfilter.RemoveLookup.MunicipalType {
			mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONMUNICIPALTYPES, "municipalityId", "uniqueId", "ref.municipalType", "ref.municipalType")...)
		}
		if !propertyfilter.RemoveLookup.Wallet {
			mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPROPERTYWALLET, "uniqueId", "propertyId", "ref.wallet", "ref.wallet")...)
		}
		if !propertyfilter.RemoveLookup.User {
			mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "created.by", "userName", "ref.user", "ref.user")...)
		}
		if !propertyfilter.RemoveLookup.Activator {
			mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "log.by.id", "userName", "ref.activator", "ref.activator")...)
		}
	}

	mainPipeline = append(mainPipeline, d.PropertyOwnersLookup(constants.COLLECTIONPROPERTYOWNER, "uniqueId", "propertyId", "ref.propertyOwner", "ref.propertyOwner")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONOVERALLPROPERTYDEMAND, "uniqueId", "propertyId", "ref.demand", "ref.demand")...)
	if propertyfilter.OmitZeroDemand {
		query = append(query, bson.M{"ref.demand.total.totalTax": bson.M{"$gte": 0}})
	}
	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"as":           "ref.propertyOwner",
			"foreignField": "propertyId",
			"from":         "owners",
			"localField":   "uniqueId",
		},
	})
	mainPipeline = append(mainPipeline, bson.M{
		"$addFields": bson.M{
			"ref.propertyOwner": bson.M{
				"$arrayElemAt": []interface{}{"$ref.propertyOwner", 0},
			},
		},
	})
	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"as":           "ref.address.state",
			"foreignField": "code",
			"from":         "states",
			"localField":   "address.stateCode",
		},
	})
	mainPipeline = append(mainPipeline, bson.M{
		"$addFields": bson.M{
			"ref.address.state": bson.M{
				"$arrayElemAt": []interface{}{"$ref.address.state", 0},
			},
		},
	})
	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"as":           "ref.address.district",
			"foreignField": "code",
			"from":         "districts",
			"localField":   "address.districtCode",
		},
	})
	mainPipeline = append(mainPipeline, bson.M{
		"$addFields": bson.M{
			"ref.address.district": bson.M{
				"$arrayElemAt": []interface{}{
					"$ref.address.district",
					0,
				},
			},
		},
	})
	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"as":           "ref.address.village",
			"foreignField": "code",
			"from":         "villages",
			"localField":   "address.villageCode",
		},
	})
	mainPipeline = append(mainPipeline, bson.M{
		"$addFields": bson.M{
			"ref.address.village": bson.M{
				"$arrayElemAt": []interface{}{"$ref.address.village", 0},
			},
		},
	})
	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"as":           "ref.address.zone",
			"foreignField": "code",
			"from":         "zones",
			"localField":   "address.zoneCode",
		},
	})
	mainPipeline = append(mainPipeline, bson.M{
		"$addFields": bson.M{
			"ref.address.zone": bson.M{
				"$arrayElemAt": []interface{}{"$ref.address.zone", 0},
			},
		},
	})
	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"as":           "ref.address.ward",
			"foreignField": "code",
			"from":         "wards",
			"localField":   "address.wardCode",
		},
	})
	mainPipeline = append(mainPipeline, bson.M{
		"$addFields": bson.M{
			"ref.address.ward": bson.M{
				"$arrayElemAt": []interface{}{"$ref.address.ward", 0},
			},
		},
	})
	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"as":   "ref.floors",
			"from": "floors",
			"let": bson.M{
				"propertyId": "$uniqueId",
			},
			"pipeline": []bson.M{
				bson.M{
					"$match": bson.M{
						"$expr": bson.M{
							"$and": []bson.M{
								bson.M{
									"$eq": []string{
										"$propertyId",
										"$$propertyId",
									},
								},
								bson.M{
									"$eq": []string{"$status", "Active"},
								},
							},
						},
					},
				},
				{
					"$lookup": bson.M{
						"as":           "ref.usageType",
						"foreignField": "uniqueId",
						"from":         "usagetypes",
						"localField":   "usageType",
					},
				},
				{
					"$addFields": bson.M{
						"ref.usageType": bson.M{
							"$arrayElemAt": []interface{}{"$ref.usageType", 0},
						},
					},
				},
				{
					"$lookup": bson.M{
						"as":           "ref.constructionType",
						"foreignField": "uniqueId",
						"from":         "constructiontypes",
						"localField":   "constructionType",
					},
				},
				{
					"$addFields": bson.M{
						"ref.constructionType": bson.M{
							"$arrayElemAt": []interface{}{
								"$ref.constructionType",
								0,
							},
						},
					},
				},
				{
					"$lookup": bson.M{
						"as":           "ref.occupancyType",
						"foreignField": "uniqueId",
						"from":         "occupancytypes",
						"localField":   "occupancyType",
					},
				},
				{
					"$addFields": bson.M{
						"ref.occupancyType": bson.M{
							"$arrayElemAt": []interface{}{
								"$ref.occupancyType",
								0,
							},
						},
					},
				},
				{
					"$lookup": bson.M{
						"as":           "ref.nonResUsageType",
						"foreignField": "uniqueId",
						"from":         "nonresidentialusagefactors",
						"localField":   "nonResUsageType",
					},
				},
				{
					"$addFields": bson.M{
						"ref.nonResUsageType": bson.M{
							"$arrayElemAt": []interface{}{
								"$ref.nonResUsageType",
								0,
							},
						},
					},
				},
				{
					"$lookup": bson.M{
						"as":           "ref.floorRatableArea",
						"foreignField": "uniqueId",
						"from":         "floorratableareas",
						"localField":   "ratableAreaType",
					},
				},
				{
					"$addFields": bson.M{
						"ref.floorRatableArea": bson.M{
							"$arrayElemAt": []interface{}{
								"$ref.floorRatableArea",
								0,
							},
						},
					},
				},
				{
					"$lookup": bson.M{
						"as":           "ref.floorNo",
						"foreignField": "uniqueId",
						"from":         "floortypes",
						"localField":   "no",
					},
				},
				{
					"$addFields": bson.M{
						"ref.floorNo": bson.M{
							"$arrayElemAt": []interface{}{"$ref.floorNo", 0},
						},
					},
				},
			},
		},
	})
	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"as":           "ref.propertyType",
			"foreignField": "uniqueId",
			"from":         "propertytypes",
			"localField":   "propertyTypeId",
		},
	})
	mainPipeline = append(mainPipeline, bson.M{
		"$addFields": bson.M{
			"ref.propertyType": bson.M{
				"$arrayElemAt": []interface{}{"$ref.propertyType", 0},
			},
		},
	})
	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"as":           "ref.roadType",
			"foreignField": "uniqueId",
			"from":         "roadtypes",
			"localField":   "roadTypeId",
		},
	})
	mainPipeline = append(mainPipeline, bson.M{
		"$addFields": bson.M{
			"ref.roadType": bson.M{
				"$arrayElemAt": []interface{}{"$ref.roadType", 0},
			},
		},
	})
	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"as":           "ref.municipalType",
			"foreignField": "uniqueId",
			"from":         "municipaltypes",
			"localField":   "municipalityId",
		},
	})
	mainPipeline = append(mainPipeline, bson.M{
		"$addFields": bson.M{
			"ref.municipalType": bson.M{
				"$arrayElemAt": []interface{}{"$ref.municipalType", 0},
			},
		},
	})
	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"as":           "ref.wallet",
			"foreignField": "propertyId",
			"from":         "propertywallet",
			"localField":   "uniqueId",
		},
	})
	mainPipeline = append(mainPipeline, bson.M{
		"$addFields": bson.M{
			"ref.wallet": bson.M{
				"$arrayElemAt": []interface{}{"$ref.wallet", 0},
			},
		},
	})
	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"as":           "ref.user",
			"foreignField": "userName",
			"from":         "users",
			"localField":   "created.by",
		},
	})
	mainPipeline = append(mainPipeline, bson.M{
		"$addFields": bson.M{
			"ref.user": bson.M{
				"$arrayElemAt": []interface{}{"$ref.user", 0},
			},
		},
	})
	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"as":           "ref.activator",
			"foreignField": "userName",
			"from":         "users",
			"localField":   "log.by.id",
		},
	})
	mainPipeline = append(mainPipeline, bson.M{
		"$addFields": bson.M{
			"ref.activator": bson.M{
				"$arrayElemAt": []interface{}{"$ref.activator", 0},
			},
		},
	})
	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"as":   "ref.propertyOwner",
			"from": "owners",
			"let": bson.M{
				"propertyId": "$uniqueId",
			},
			"pipeline": []bson.M{
				bson.M{
					"$match": bson.M{
						"$expr": bson.M{
							"$and": []bson.M{
								bson.M{
									"$eq": []string{"$propertyId",
										"$$propertyId"},
								},
								bson.M{
									"$eq": []string{"$status", "Active"},
								},
							},
						},
					},
				},
			},
		},
	})
	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"as":           "ref.demand",
			"foreignField": "propertyId",
			"from":         "overallpropertydemand",
			"localField":   "uniqueId",
		},
	})
	mainPipeline = append(mainPipeline, bson.M{
		"$addFields": bson.M{
			"ref.demand": bson.M{
				"$arrayElemAt": []interface{}{"$ref.demand", 0},
			},
		},
	})
	mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{
		"from": constants.COLLECTIONPROPERTYPAYMENTFY,
		"as":   "ref.collections",
		"let":  bson.M{"varUniqueId": "$uniqueId"},
		"pipeline": []bson.M{bson.M{
			"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
				bson.M{"$eq": []string{"$status", "Completed"}},
				bson.M{"$eq": []string{"$propertyId", "$$varUniqueId"}},
			}}},
		},
			bson.M{"$group": bson.M{"_id": "$propertyId",
				"arrearTax": bson.M{"$sum": bson.M{"$cond": bson.M{"else": 0,
					"if":   bson.M{"$ne": []interface{}{"$fy.uniqueId", resFYs.UniqueID}},
					"then": bson.M{"$sum": bson.M{"$add": []interface{}{"$fy.vacantLandTax", "$fy.tax"}}}}}},
				"currentTax": bson.M{"$sum": bson.M{"$cond": bson.M{"else": 0,
					"if":   bson.M{"$eq": []interface{}{"$fy.uniqueId", resFYs.UniqueID}},
					"then": bson.M{"$sum": bson.M{"$add": []interface{}{"$fy.vacantLandTax", "$fy.tax"}}}}}},
				"arrearPenalty": bson.M{"$sum": bson.M{"$cond": bson.M{"else": 0,
					"if":   bson.M{"$ne": []interface{}{"$fy.uniqueId", resFYs.UniqueID}},
					"then": bson.M{"$sum": bson.M{"$add": []interface{}{"$fy.penanty"}}}}}},
				"currentPenalty": bson.M{"$sum": bson.M{"$cond": bson.M{"else": 0,
					"if":   bson.M{"$eq": []interface{}{"$fy.uniqueId", resFYs.UniqueID}},
					"then": bson.M{"$sum": bson.M{"$add": []interface{}{"$fy.penanty"}}}}}},
				"arrearRebate": bson.M{"$sum": bson.M{"$cond": bson.M{"else": 0,
					"if":   bson.M{"$ne": []interface{}{"$fy.uniqueId", resFYs.UniqueID}},
					"then": bson.M{"$sum": bson.M{"$add": []interface{}{"$fy.rebate"}}}}}},
				"currentRebate": bson.M{"$sum": bson.M{"$cond": bson.M{"else": 0,
					"if":   bson.M{"$eq": []interface{}{"$fy.uniqueId", resFYs.UniqueID}},
					"then": bson.M{"$sum": bson.M{"$add": []interface{}{"$fy.rebate"}}}}}},
				"arrearAlreadyPaid": bson.M{"$sum": bson.M{"$cond": bson.M{"else": 0,
					"if":   bson.M{"$ne": []interface{}{"$fy.uniqueId", resFYs.UniqueID}},
					"then": bson.M{"$sum": bson.M{"$add": []interface{}{"$fy.alreadyPayed.fyTax", "$fy.alreadyPayed.vlTax"}}}}}},
				"currentAlreadyPaid": bson.M{"$sum": bson.M{"$cond": bson.M{"else": 0,
					"if":   bson.M{"$eq": []interface{}{"$fy.uniqueId", resFYs.UniqueID}},
					"then": bson.M{"$sum": bson.M{"$add": []interface{}{"$fy.alreadyPayed.fyTax", "$fy.alreadyPayed.vlTax"}}}}}},
				"totalTax":    bson.M{"$sum": "$fy.totalTax"},
				"otherDemand": bson.M{"$sum": "$fy.otherDemand"},
			}},
		},
	},
	})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"ref.collections": bson.M{"$arrayElemAt": []interface{}{"$ref.collections", 0}}}})
	mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{
		"from": constants.COLLECTIONPROPERTYPAYMENT,
		"as":   "ref.propertyPayments",
		"let":  bson.M{"varUniqueId": "$uniqueId"},
		"pipeline": []bson.M{
			bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
				bson.M{"$eq": []string{"$status", "Completed"}},
				bson.M{"$eq": []string{"$$varUniqueId", "$propertyId"}},
			}}}},
			bson.M{"$group": bson.M{"_id": "$propertyId",
				"rebate":  bson.M{"$sum": "$demand.rebate"},
				"formFee": bson.M{"$sum": "$demand.formFee"},
			}},
		},
	}})
	mainPipeline = append(mainPipeline, bson.M{
		"$addFields": bson.M{
			"ref.propertyPayments": bson.M{
				"$arrayElemAt": []interface{}{"$ref.propertyPayments", 0},
			},
		},
	})
	mainPipeline = append(mainPipeline, bson.M{
		"$project": bson.M{
			"_id":                                  false,
			"uniqueId":                             true,
			"applicationNo":                        true,
			"advance":                              true,
			"address.al1":                          true,
			"address.al2":                          true,
			"ref.address.ward.name":                true,
			"ref.propertyOwner.mobile":             true,
			"ref.propertyOwner.name":               true,
			"ref.propertyOwner.fatherRpanRhusband": true,
			"areaOfPlot":                           true,
			"ref.demand.arrear.totalTax":           true,
			"ref.demand.current.totalTax":          true,
			"ref.demand.total.totalTax":            true,
			"ref.demand.total.ecess":               true,
			"ref.propertyPayments.formFee":         true,
			"ref.propertyPayments.rebate":          true,
			"ref.roadType.name":                    true,
			"ref.collections":                      true,
			"ref.propertyType.name":                true,
		},
	})
	//Aggregation
	d.Shared.BsonToJSONPrintTag("property query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var propertys []models.RefPropertyV2
	if err = cursor.All(context.TODO(), &propertys); err != nil {
		return nil, err
	}
	return propertys, nil
}

func (d *Daos) PropertyWiseDemandCollectionandBalanceReport(ctx *models.Context, propertyfilter *models.PropertyFilter, pagination *models.Pagination) ([]models.RefProperty, error) {
	// resFYs, err := d.GetCurrentFinancialYear(ctx)
	// if err != nil {
	// 	return nil, errors.New("error in getting current financial year" + err.Error())
	// }
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if propertyfilter != nil {
		if propertyfilter.IsLocation {
			mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"address.location.coordinates": bson.M{"$exists": true}}})
			mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"locSize": bson.M{"$size": "$address.location.coordinates"}}})
		}
	}

	query = d.FilterPropertyQuery(ctx, propertyfilter)
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if propertyfilter.SortBy != "" {
		mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{propertyfilter.SortBy: propertyfilter.SortOrder}})
	} else {
		mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"_id": -1}})
	}
	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONPROPERTY).CountDocuments(ctx.CTX, func() bson.M {
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

	//Lookups
	/*
		if propertyfilter != nil {
			if !propertyfilter.RemoveLookup.PropertyOwner {
				mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPROPERTYOWNER, "uniqueId", "propertyId", "ref.propertyOwner", "ref.propertyOwner")...)
			}
			if !propertyfilter.RemoveLookup.State {
				mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "address.stateCode", "code", "ref.address.state", "ref.address.state")...)
			}
			if !propertyfilter.RemoveLookup.District {
				mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "address.districtCode", "code", "ref.address.district", "ref.address.district")...)
			}
			if !propertyfilter.RemoveLookup.Village {
				mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVILLAGE, "address.villageCode", "code", "ref.address.village", "ref.address.village")...)
			}
			if !propertyfilter.RemoveLookup.Zone {
				mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONZONE, "address.zoneCode", "code", "ref.address.zone", "ref.address.zone")...)
			}
			if !propertyfilter.RemoveLookup.Ward {
				mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONWARD, "address.wardCode", "code", "ref.address.ward", "ref.address.ward")...)
			}
			if !propertyfilter.RemoveLookup.PropertyFloor {
				mainPipeline = append(mainPipeline, d.PropertyFloorsLookup(constants.COLLECTIONPROPERTYFLOOR, "uniqueId", "propertyId", "ref.floors", "ref.floors")...)
			}
			if !propertyfilter.RemoveLookup.PropertyType {
				mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPROPERTYTYPE, "propertyTypeId", "uniqueId", "ref.propertyType", "ref.propertyType")...)
			}
			if !propertyfilter.RemoveLookup.RoadType {
				mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONROADTYPE, "roadTypeId", "uniqueId", "ref.roadType", "ref.roadType")...)
			}
			if !propertyfilter.RemoveLookup.MunicipalType {
				mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONMUNICIPALTYPES, "municipalityId", "uniqueId", "ref.municipalType", "ref.municipalType")...)
			}
			if !propertyfilter.RemoveLookup.Wallet {
				mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPROPERTYWALLET, "uniqueId", "propertyId", "ref.wallet", "ref.wallet")...)
			}
			if !propertyfilter.RemoveLookup.User {
				mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "created.by", "userName", "ref.user", "ref.user")...)
			}
			if !propertyfilter.RemoveLookup.Activator {
				mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "log.by.id", "userName", "ref.activator", "ref.activator")...)
			}
		}
	*/

	mainPipeline = append(mainPipeline, d.PropertyOwnersLookup(constants.COLLECTIONPROPERTYOWNER, "uniqueId", "propertyId", "ref.propertyOwner", "ref.propertyOwner")...)
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPROPERTYPAYMENTBASIC, "uniqueId", "property.uniqueId", "basic", "basic")...)
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONOVERALLPROPERTYDEMAND, "uniqueId", "propertyId", "ref.demand", "ref.demand")...)
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "address.stateCode", "code", "ref.address.state", "ref.address.state")...)
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "address.districtCode", "code", "ref.address.district", "ref.address.district")...)
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVILLAGE, "address.villageCode", "code", "ref.address.village", "ref.address.village")...)
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONZONE, "address.zoneCode", "code", "ref.address.zone", "ref.address.zone")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONWARD, "address.wardCode", "code", "ref.address.ward", "ref.address.ward")...)

	//mainPipeline = append(mainPipeline, d.CommonLookup(constants., "uniqueId", "propertyId", "ref.basic", "ref.basic")...)
	if propertyfilter.OmitZeroDemand {
		query = append(query, bson.M{"ref.demand.total.totalTax": bson.M{"$gte": 0}})
	}

	//Aggregation
	d.Shared.BsonToJSONPrintTag("property query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var propertys []models.RefProperty
	if err = cursor.All(context.TODO(), &propertys); err != nil {
		return nil, err
	}
	return propertys, nil
}
