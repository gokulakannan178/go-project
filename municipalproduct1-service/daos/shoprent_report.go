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

func (d *Daos) WardWiseShoprentReport(ctx *models.Context, filter *models.WardWiseShoprentReportFilter) ([]models.WardWiseShoprentReport, error) {

	mainPipeline := []bson.M{}
	var sd, ed *time.Time
	query := []bson.M{}
	if filter != nil {
		if len(filter.ZoneCode) > 0 {
			query = append(query, bson.M{"zoneCode": bson.M{"$in": filter.ZoneCode}})
		}
		query = append(query, bson.M{"status": bson.M{"$in": []string{constants.WARDSTATUSACTIVE}}})
		if filter.StartDate != nil {
			sdt := time.Date(filter.StartDate.Year(), filter.StartDate.Month(), filter.StartDate.Day(), 0, 0, 0, 0, filter.StartDate.Location())
			edt := time.Date(filter.StartDate.Year(), filter.StartDate.Month(), filter.StartDate.Day(), 23, 59, 59, 0, filter.StartDate.Location())
			sd = &sdt
			ed = &edt
			if filter.EndDate != nil {
				edt := time.Date(filter.EndDate.Year(), filter.EndDate.Month(), filter.EndDate.Day(), 23, 59, 59, 0, filter.EndDate.Location())
				ed = &edt
			}
		}

	}
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if filter.SortBy != "" {
		mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{filter.SortBy: filter.SortOrder}})
	} else {
		mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"name": 1}})
	}
	var ShoprentreportPipelineAnd = []bson.M{
		bson.M{"$eq": []string{"$address.wardCode", "$$wardId"}},
		bson.M{"$eq": []string{"$status", "Completed"}},
	}

	if sd != nil {
		ShoprentreportPipelineAnd = append(ShoprentreportPipelineAnd,
			bson.M{"$gte": []string{"$completionDate", "$$sd"}},
			bson.M{"$lte": []string{"$completionDate", "$$ed"}})
	}
	mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{
		"from": "shoprent",
		"as":   "shoprentpayments",
		"let":  bson.M{"wardId": "$code", "sd": sd, "ed": ed},
		"pipeline": []bson.M{
			bson.M{"$match": bson.M{"$expr": bson.M{"$and": ShoprentreportPipelineAnd}}},
			bson.M{"$group": bson.M{"_id": "$shopRentId", "amount": bson.M{"$sum": "$details.amount"}}},
			bson.M{"$group": bson.M{"_id": nil, "propertCount": bson.M{"$sum": 1}, "amount": bson.M{"$sum": "$amount"}}},
		},
	}})

	mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{
		"from": "shoprent",
		"as":   "allPropertyCount",
		"let":  bson.M{"wardId": "$code"},
		"pipeline": []bson.M{
			bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
				bson.M{"$eq": []string{"$address.wardCode", "$$wardId"}},
				bson.M{"$in": []interface{}{"$status", "Active"}},
			}}}},
			bson.M{"$group": bson.M{"_id": nil, "totalCount": bson.M{"$sum": 1}}},
		},
	}})

	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{
		"shoprentpayments": bson.M{"$arrayElemAt": []interface{}{"$shoprentpayments", 0}},
		"allPropertyCount": bson.M{"$arrayElemAt": []interface{}{"$allPropertyCount", 0}},
	}})

	//Aggregation
	d.Shared.BsonToJSONPrint(mainPipeline)
	var data []models.WardWiseShoprentReport

	cursor, err := ctx.DB.Collection(constants.COLLECTIONWARD).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}

	if err = cursor.All(context.TODO(), &data); err != nil {
		return nil, err
	}

	return data, nil
}

func (d *Daos) DashboardDayWiseShoprentCollectionChart(ctx *models.Context, filter *models.DashboardDayWiseShoprentCollectionChartFilter) (models.DashboardDayWiseShoprentCollectionChart, error) {

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
		"shops": "$shopRentId",
		"day":   bson.M{"$dayOfMonth": "$completionDate"},
	}, "amount": bson.M{"$sum": "$details.amount"}}})
	mainPipeline = append(mainPipeline, bson.M{"$group": bson.M{"_id": "$_id.day", "shopRentCount": bson.M{"$sum": 1}, "amount": bson.M{"$sum": "$amount"}}})
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
									"shopRentCount": 0,
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
	var emptyData models.DashboardDayWiseShoprentCollectionChart

	cursor, err := ctx.DB.Collection(constants.COLLECTIONSHOPRENTPAYMENTS).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return emptyData, err
	}
	var data []models.DashboardDayWiseShoprentCollectionChart

	if err = cursor.All(context.TODO(), &data); err != nil {
		return emptyData, err
	}
	if len(data) > 0 {
		return data[0], nil
	}

	return emptyData, nil

}

// DayWiseShoprentDemandChart : ""
func (d *Daos) DayWiseShoprentDemandChart(ctx *models.Context, filter *models.DayWiseShoprentDemandChartFilter) (*models.DayWiseShoprentDemandChart, error) {

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
		"shopRentCount": bson.M{"$sum": 1}, "amount": bson.M{"$sum": "$demand.total.total"}}})
	// mainPipeline = append(mainPipeline, bson.M{"$group": bson.M{"_id": "$_id.day", "shopRentCount": bson.M{"$sum": 1}, "amount": bson.M{"$sum": "$amount"}}})
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
									"shopRentCount": 0,
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
	var emptyData *models.DayWiseShoprentDemandChart

	cursor, err := ctx.DB.Collection(constants.COLLECTIONSHOPRENT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return emptyData, err
	}
	var data []models.DayWiseShoprentDemandChart

	if err = cursor.All(context.TODO(), &data); err != nil {
		return emptyData, err
	}
	if len(data) > 0 {
		return &data[0], nil
	}

	return emptyData, nil

}

// FilterWardDayWiseShopRentCollectionReport : ""
func (d *Daos) FilterWardDayWiseShopRentCollectionReport(ctx *models.Context, filter *models.WardDayWiseShopRentCollectionReportFilter, pagination *models.Pagination) ([]models.WardDayWiseShopRentCollectionReport, error) {
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
		"from": constants.COLLECTIONSHOPRENTPAYMENTS,
		"as":   "report",
		"let":  bson.M{"status": "$status", "code": "$code"},
		"pipeline": []bson.M{{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
			{"$eq": []interface{}{"$status", constants.SHOPRENTPAYMENTSTATUSCOMPLETED}},
			{"$eq": []interface{}{"$$code", "$address.wardCode"}},
			{"$gte": []interface{}{"$completionDate", sd}},
			{"$lte": []interface{}{"$completionDate", ed}},
		}}}},
			{"$group": bson.M{"_id": "$shopRentId", "totalNoPayments": bson.M{"$sum": 1}, "totalCollection": bson.M{"$sum": "$details.amount"}}},
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
	var res []models.WardDayWiseShopRentCollectionReport
	if err = cursor.All(context.TODO(), &res); err != nil {
		return nil, err
	}
	return res, nil
}

// FilterWardMonthWiseShopRentCollectionReport : ""
func (d *Daos) FilterWardMonthWiseShopRentCollectionReport(ctx *models.Context, filter *models.WardMonthWiseShopRentCollectionReportFilter, pagination *models.Pagination) ([]models.WardMonthWiseShopRentCollectionReport, error) {
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
		"from": constants.COLLECTIONSHOPRENTPAYMENTS,
		"as":   "report",
		"let":  bson.M{"status": "$status", "code": "$code"},
		"pipeline": []bson.M{{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
			{"$eq": []interface{}{"$status", constants.SHOPRENTPAYMENTSTATUSCOMPLETED}},
			{"$eq": []interface{}{"$$code", "$address.wardCode"}},
			{"$gte": []interface{}{"$completionDate", sd}},
			{"$lte": []interface{}{"$completionDate", ed}},
		}}}},
			{"$group": bson.M{"_id": "$shopRentId", "totalNoPayments": bson.M{"$sum": 1}, "totalCollection": bson.M{"$sum": "$details.amount"}}},
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
	var res []models.WardMonthWiseShopRentCollectionReport
	if err = cursor.All(context.TODO(), &res); err != nil {
		return nil, err
	}
	return res, nil
}

// FilterWardDayWiseShopRentDemandReport : ""
func (d *Daos) FilterWardDayWiseShopRentDemandReport(ctx *models.Context, filter *models.WardDayWiseShopRentDemandReportFilter, pagination *models.Pagination) ([]models.WardDayWiseShopRentDemandReport, error) {
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
		"from": constants.COLLECTIONSHOPRENT,
		"as":   "report",
		"let":  bson.M{"status": "$status", "code": "$code"},
		"pipeline": []bson.M{bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
			bson.M{"$eq": []interface{}{"$status", constants.SHOPRENTSTATUSACTIVE}},
			{"$eq": []interface{}{"$$code", "$address.wardCode"}},
			bson.M{"$gte": []interface{}{"$created.on", sd}},
			bson.M{"$lte": []interface{}{"$created.on", ed}},
		}}}},
			bson.M{"$group": bson.M{"_id": nil, "shoprents": bson.M{"$sum": 1}, "totalDemand": bson.M{"$sum": "$demand.total.total"}}},
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
	var res []models.WardDayWiseShopRentDemandReport
	if err = cursor.All(context.TODO(), &res); err != nil {
		return nil, err
	}
	return res, nil
}

// FilterWardMonthWiseShopRentDemandReport : ""
func (d *Daos) FilterWardMonthWiseShopRentDemandReport(ctx *models.Context, filter *models.WardMonthWiseShopRentDemandReportFilter, pagination *models.Pagination) ([]models.WardMonthWiseShopRentDemandReport, error) {
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
		"from": constants.COLLECTIONSHOPRENT,
		"as":   "report",
		"let":  bson.M{"status": "$status", "code": "$code"},
		"pipeline": []bson.M{bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
			bson.M{"$eq": []interface{}{"$status", constants.SHOPRENTSTATUSACTIVE}},
			{"$eq": []interface{}{"$$code", "$address.wardCode"}},
			bson.M{"$gte": []interface{}{"$created.on", sd}},
			bson.M{"$lte": []interface{}{"$created.on", ed}},
		}}}},
			bson.M{"$group": bson.M{"_id": nil, "shoprents": bson.M{"$sum": 1}, "totalDemand": bson.M{"$sum": "$demand.total.total"}}},
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
	var res []models.WardMonthWiseShopRentDemandReport
	if err = cursor.All(context.TODO(), &res); err != nil {
		return nil, err
	}
	return res, nil
}

// FilterTeamDayWiseShopRentCollectionReport : ""
func (d *Daos) FilterTeamDayWiseShopRentCollectionReport(ctx *models.Context, filter *models.TeamDayWiseShopRentCollectionReportFilter, pagination *models.Pagination) ([]models.TeamDayWiseShopRentCollectionReport, error) {
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
		"from": constants.COLLECTIONSHOPRENTPAYMENTS,
		"as":   "report",
		"let":  bson.M{"userName": "$userName"},
		"pipeline": []bson.M{bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
			bson.M{"$eq": []interface{}{"$status", constants.SHOPRENTPAYMENTSTATUSCOMPLETED}},
			bson.M{"$eq": []interface{}{"$$userName", "$details.collector.by"}},

			bson.M{"$gte": []interface{}{"$completionDate", sd}},
			bson.M{"$lte": []interface{}{"$completionDate", ed}},
		}}}},
			bson.M{"$group": bson.M{"_id": "$shopRentId", "totalNoPayments": bson.M{"$sum": 1},
				"totalCollection": bson.M{"$sum": "$details.amount"}}},
			bson.M{"$group": bson.M{"_id": nil, "totalNoShopRents": bson.M{"$sum": 1}, "totalNoPayments": bson.M{"$sum": "$totalNoPayments"},
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
	var res []models.TeamDayWiseShopRentCollectionReport
	if err = cursor.All(context.TODO(), &res); err != nil {
		return nil, err
	}
	return res, nil
}

// FilterTeamMonthWiseShopRentCollectionReport : ""
func (d *Daos) FilterTeamMonthWiseShopRentCollectionReport(ctx *models.Context, filter *models.TeamMonthWiseShopRentCollectionReportFilter, pagination *models.Pagination) ([]models.TeamMonthWiseShopRentCollectionReport, error) {
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
		"from": constants.COLLECTIONSHOPRENTPAYMENTS,
		"as":   "report",
		"let":  bson.M{"userName": "$userName"},
		"pipeline": []bson.M{bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
			bson.M{"$eq": []interface{}{"$status", constants.SHOPRENTPAYMENTSTATUSCOMPLETED}},
			bson.M{"$eq": []interface{}{"$$userName", "$details.collector.by"}},

			bson.M{"$gte": []interface{}{"$completionDate", sd}},
			bson.M{"$lte": []interface{}{"$completionDate", ed}},
		}}}},
			bson.M{"$group": bson.M{"_id": "$shopRentId", "totalNoPayments": bson.M{"$sum": 1},
				"totalCollection": bson.M{"$sum": "$details.amount"}}},
			bson.M{"$group": bson.M{"_id": nil, "totalNoShopRents": bson.M{"$sum": 1}, "totalNoPayments": bson.M{"$sum": "$totalNoPayments"},
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
	var res []models.TeamMonthWiseShopRentCollectionReport
	if err = cursor.All(context.TODO(), &res); err != nil {
		return nil, err
	}
	return res, nil
}

// FilterWardYearWiseShopRentCollectionReport : ""
func (d *Daos) FilterWardYearWiseShopRentCollectionReport(ctx *models.Context, filter *models.WardYearWiseShopRentCollectionReportFilter, pagination *models.Pagination) ([]models.WardYearWiseShopRentCollectionReport, error) {
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
		"from": constants.COLLECTIONSHOPRENTPAYMENTS,
		"as":   "report",
		"let":  bson.M{"status": "$status", "code": "$code"},
		"pipeline": []bson.M{{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
			{"$eq": []interface{}{"$status", constants.SHOPRENTPAYMENTSTATUSCOMPLETED}},
			{"$eq": []interface{}{"$$code", "$address.wardCode"}},
			{"$gte": []interface{}{"$completionDate", sd}},
			{"$lte": []interface{}{"$completionDate", ed}},
		}}}},
			{"$group": bson.M{"_id": "$shopRentId", "totalNoPayments": bson.M{"$sum": 1}, "totalCollection": bson.M{"$sum": "$details.amount"}}},
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
	var res []models.WardYearWiseShopRentCollectionReport
	if err = cursor.All(context.TODO(), &res); err != nil {
		return nil, err
	}
	return res, nil
}

// FilterWardYearWiseShopRentDemandReport : ""
func (d *Daos) FilterWardYearWiseShopRentDemandReport(ctx *models.Context, filter *models.WardYearWiseShopRentDemandReportFilter, pagination *models.Pagination) ([]models.WardYearWiseShopRentDemandReport, error) {
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
		"from": constants.COLLECTIONSHOPRENT,
		"as":   "report",
		"let":  bson.M{"status": "$status", "code": "$code"},
		"pipeline": []bson.M{bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
			bson.M{"$eq": []interface{}{"$status", constants.SHOPRENTSTATUSACTIVE}},
			{"$eq": []interface{}{"$$code", "$address.wardCode"}},
			bson.M{"$gte": []interface{}{"$created.on", sd}},
			bson.M{"$lte": []interface{}{"$created.on", ed}},
		}}}},
			bson.M{"$group": bson.M{"_id": nil, "shopRents": bson.M{"$sum": 1}, "totalDemand": bson.M{"$sum": "$demand.total.total"}}},
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
	var res []models.WardYearWiseShopRentDemandReport
	if err = cursor.All(context.TODO(), &res); err != nil {
		return nil, err
	}
	return res, nil
}

// FilterTeamYearWiseShopRentCollectionReport : ""
func (d *Daos) FilterTeamYearWiseShopRentCollectionReport(ctx *models.Context, filter *models.TeamYearWiseShopRentCollectionReportFilter, pagination *models.Pagination) ([]models.TeamYearWiseShopRentCollectionReport, error) {
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
		"from": constants.COLLECTIONSHOPRENTPAYMENTS,
		"as":   "report",
		"let":  bson.M{"userName": "$userName"},
		"pipeline": []bson.M{bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
			bson.M{"$eq": []interface{}{"$status", constants.SHOPRENTPAYMENTSTATUSCOMPLETED}},
			bson.M{"$eq": []interface{}{"$$userName", "$details.collector.by"}},

			bson.M{"$gte": []interface{}{"$completionDate", sd}},
			bson.M{"$lte": []interface{}{"$completionDate", ed}},
		}}}},
			bson.M{"$group": bson.M{"_id": "$shopRentId", "totalNoPayments": bson.M{"$sum": 1},
				"totalCollection": bson.M{"$sum": "$details.amount"}}},
			bson.M{"$group": bson.M{"_id": nil, "totalNoShopRents": bson.M{"$sum": 1}, "totalNoPayments": bson.M{"$sum": "$totalNoPayments"},
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
	var res []models.TeamYearWiseShopRentCollectionReport
	if err = cursor.All(context.TODO(), &res); err != nil {
		return nil, err
	}
	return res, nil
}
