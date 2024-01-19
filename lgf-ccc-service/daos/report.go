package daos

import (
	"fmt"
	"lgf-ccc-service/constants"
	"lgf-ccc-service/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func (d *Daos) DayWiseDumphistoryCount(ctx *models.Context, filter *models.FilterDumpHistory) ([]models.MonthWiseDumphistoryCount, error) {
	mainPipeline := []bson.M{}
	//var sd, ed *time.Time

	startmonth := d.Shared.BeginningOfMonth(*filter.Date)
	sm := &startmonth
	endmonth := d.Shared.EndOfMonth(*filter.Date)
	em := &endmonth
	fmt.Println("sm===========>", sm)
	fmt.Println("em===========>", em)
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": []bson.M{
		// bson.M{"$gte": []interface{}{"$date", sm}},
		// bson.M{"$lte": []interface{}{"$date", em}},{}
		bson.M{"date": bson.M{"$gte": sm,
			"$lte": em}},
	}}},
		bson.M{"$group": bson.M{"_id": bson.M{"$dateToString": bson.M{"format": "%Y-%m-%d", "date": "$date"}},
			"quantity": bson.M{"$sum": "$quantity"}}},
	)

	d.Shared.BsonToJSONPrintTag("query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDUMPHISTORY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var Reports []models.MonthWiseDumphistoryCount
	if err = cursor.All(ctx.CTX, &Reports); err != nil {
		return nil, err
	}

	return Reports, err
}

func (d *Daos) MonthWiseDumphistoryCount(ctx *models.Context, filter *models.FilterDumpHistory) ([]models.MonthWiseDumphistoryCount, error) {
	mainPipeline := []bson.M{}
	var sy, ey time.Time
	query := []bson.M{}
	if filter != nil {

		if filter.Date != nil {

			if filter.Date != nil {
				sy = time.Date(filter.Date.Year(), time.January, 1, 0, 0, 0, 0, filter.Date.Location())
				ey = time.Date(filter.Date.Year(), time.December, 31, 23, 59, 59, 0, filter.Date.Location())
				if filter.Date != nil {
					ey = time.Date(filter.Date.Year(), time.December, 31, 23, 59, 59, 0, filter.Date.Location())
				}
				query = append(query, bson.M{"date": bson.M{"$gte": sy, "$lte": ey}})

			}
		}

		if len(query) > 0 {
			mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
		}
	}

	// startmonth := d.Shared.BeginningOfMonth(*filter.Date)
	// sm := &startmonth
	// endmonth := d.Shared.EndOfMonth(*filter.Date)
	// em := &endmonth
	// fmt.Println("sm===========>", sm)
	// fmt.Println("em===========>", em)
	mainPipeline = append(mainPipeline, bson.M{"$group": bson.M{"_id": bson.M{
		"month": bson.M{"$month": "$date"},
	}, "quantity": bson.M{"$sum": "$quantity"},
	}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"date": "$_id.month"}})
	mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"date": 1}})
	mainPipeline = append(mainPipeline, bson.M{"$project": bson.M{"_id": 0}})

	// mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": []bson.M{
	// 	// bson.M{"$gte": []interface{}{"$date", sm}},
	// 	// bson.M{"$lte": []interface{}{"$date", em}},{}
	// 	bson.M{"date": bson.M{"$gte": sy,
	// 		"$lte": ey}},
	// }}},
	// 	bson.M{"$group": bson.M{"_id": bson.M{"$month": bson.M{"format": "%Y-%m-%d", "date": "$date"}},
	// 		"quantity": bson.M{"$sum": "$quantity"}}},
	// )

	d.Shared.BsonToJSONPrintTag("query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDUMPHISTORY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var Reports []models.MonthWiseDumphistoryCount
	if err = cursor.All(ctx.CTX, &Reports); err != nil {
		return nil, err
	}

	return Reports, err
}

func (d *Daos) CircleWiseHouseVisitedCount(ctx *models.Context, filter *models.FilterHouseVisited) ([]models.CircleWiseHouseVisitedv2, error) {
	mainPipeline := []bson.M{}
	//var sd, ed time.Time
	query := []bson.M{}
	if filter != nil {
		if len(filter.CircleNo) > 0 {
			query = append(query, bson.M{"code": bson.M{"$in": filter.CircleNo}})

		}
		if len(filter.WardNo) > 0 {
			query = append(query, bson.M{"code": bson.M{"$in": filter.WardNo}})

		}

	}
	startmonth := d.Shared.BeginningOfMonth(*filter.DateRange.From)
	sm := &startmonth
	endmonth := d.Shared.EndOfMonth(*filter.DateRange.From)
	em := &endmonth
	fmt.Println("sm===========>", sm)
	fmt.Println("em===========>", em)

	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	mainPipeline = append(mainPipeline,
		bson.M{"$lookup": bson.M{
			"from": "circlewisehousevisited",
			"as":   "report",
			"let":  bson.M{"code": "$code"},
			"pipeline": []bson.M{
				bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
					bson.M{"$eq": []string{"$status", "Active"}},
					bson.M{"$eq": []string{"$circleCode", "$$code"}},
					bson.M{"$gte": []interface{}{"$date", sm}},
					bson.M{"$lte": []interface{}{"$date", em}},
				}}}},
				bson.M{"$group": bson.M{"_id": bson.M{"$dateToString": bson.M{"format": "%Y-%m-%d", "date": "$date"}},
					"field":    bson.M{"$push": "$$ROOT.date"},
					"quantity": bson.M{"$sum": "$count"}}},
				bson.M{"$addFields": bson.M{"date": bson.M{"$arrayElemAt": []interface{}{"$field", 0}}}},
				bson.M{"$addFields": bson.M{"day": bson.M{"$dayOfMonth": "$date"}}},
			}}})
	mainPipeline = append(mainPipeline, bson.M{"$unwind": "$report"})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{
		"quantity": "$report.quantity",
		"date":     "$report.date",
	}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"dom": bson.M{"$dayOfMonth": "$report.date"}}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"month": bson.M{"$month": "$report.date"}}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"year": bson.M{"$year": "$report.date"}}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"dayOfWeek": bson.M{"$dayOfWeek": "$report.date"}}})
	mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"report.date": 1}})
	mainPipeline = append(mainPipeline, bson.M{"$group": bson.M{"_id": nil, "days": bson.M{"$push": "$$ROOT"}}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"month1": bson.M{"$arrayElemAt": []interface{}{"$days.month", 0}}}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"year1": bson.M{"$arrayElemAt": []interface{}{"$days.year", 0}}}})
	//mainPipeline = append(mainPipeline, bson.M{"$group": bson.M{"_id": nil, "days": bson.M{"$push": "$$ROOT"}}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{
		"days": bson.M{
			"$map": bson.M{
				"as": "rangeDay",
				"in": bson.M{
					"$let": bson.M{
						"in": bson.M{
							"$cond": bson.M{"else": bson.M{"$arrayElemAt": []string{"$days", "$$index"}},
								"if": bson.M{"$eq": []interface{}{"$$index", -1}},
								"then": bson.M{
									"date": bson.M{"$dateFromParts": bson.M{
										"day":   "$$rangeDay",
										"hour":  9,
										"month": "$month1",
										"year":  "$year1"}},
								}}},
						"vars": bson.M{
							"index": bson.M{"$indexOfArray": []string{"$days.dom", "$$rangeDay"}}}}},
				"input": bson.M{"$range": []interface{}{1, 32, 1}}}}}})
	mainPipeline = append(mainPipeline, bson.M{"$unwind": "$days"})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{
		"date":     "$days.date",
		"quantity": "$days.quantity",
		"code":     "$days.code",
	}})

	//Aggregation
	d.Shared.BsonToJSONPrint(mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONCIRCLE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var Reports []models.CircleWiseHouseVisitedv2
	if err = cursor.All(ctx.CTX, &Reports); err != nil {
		return nil, err
	}

	return Reports, err
}

func (d *Daos) DayWiseWardHouseVisitedCount(ctx *models.Context, filter *models.FilterHouseVisited) ([]models.CircleWiseHouseVisitedv2, error) {
	mainPipeline := []bson.M{}
	//var sd, ed time.Time
	query := []bson.M{}
	if filter != nil {
		if len(filter.CircleNo) > 0 {
			query = append(query, bson.M{"code": bson.M{"$in": filter.CircleNo}})

		}
		if len(filter.WardNo) > 0 {
			query = append(query, bson.M{"code": bson.M{"$in": filter.WardNo}})

		}

	}
	startmonth := d.Shared.BeginningOfMonth(*filter.DateRange.From)
	sm := &startmonth
	endmonth := d.Shared.EndOfMonth(*filter.DateRange.From)
	em := &endmonth
	fmt.Println("sm===========>", sm)
	fmt.Println("em===========>", em)

	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	mainPipeline = append(mainPipeline,
		bson.M{"$lookup": bson.M{
			"from": "wardwisehousevisited",
			"as":   "report",
			"let":  bson.M{"code": "$code"},
			"pipeline": []bson.M{
				bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
					bson.M{"$eq": []string{"$status", "Active"}},
					bson.M{"$eq": []string{"$wardCode", "$$code"}},
					bson.M{"$gte": []interface{}{"$date", sm}},
					bson.M{"$lte": []interface{}{"$date", em}},
				}}}},
				bson.M{"$group": bson.M{"_id": bson.M{"$dateToString": bson.M{"format": "%Y-%m-%d", "date": "$date"}},
					"field":    bson.M{"$push": "$$ROOT.date"},
					"quantity": bson.M{"$sum": "$todayCollection"}}},
				bson.M{"$addFields": bson.M{"date": bson.M{"$arrayElemAt": []interface{}{"$field", 0}}}},
				bson.M{"$addFields": bson.M{"day": bson.M{"$dayOfMonth": "$date"}}},
			}}})
	mainPipeline = append(mainPipeline, bson.M{"$unwind": "$report"})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{
		"quantity": "$report.quantity",
		"date":     "$report.date",
	}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"dom": bson.M{"$dayOfMonth": "$report.date"}}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"month": bson.M{"$month": "$report.date"}}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"year": bson.M{"$year": "$report.date"}}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"dayOfWeek": bson.M{"$dayOfWeek": "$report.date"}}})
	mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"report.date": 1}})
	mainPipeline = append(mainPipeline, bson.M{"$group": bson.M{"_id": nil, "days": bson.M{"$push": "$$ROOT"}}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"month1": bson.M{"$arrayElemAt": []interface{}{"$days.month", 0}}}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"year1": bson.M{"$arrayElemAt": []interface{}{"$days.year", 0}}}})
	//mainPipeline = append(mainPipeline, bson.M{"$group": bson.M{"_id": nil, "days": bson.M{"$push": "$$ROOT"}}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{
		"days": bson.M{
			"$map": bson.M{
				"as": "rangeDay",
				"in": bson.M{
					"$let": bson.M{
						"in": bson.M{
							"$cond": bson.M{"else": bson.M{"$arrayElemAt": []string{"$days", "$$index"}},
								"if": bson.M{"$eq": []interface{}{"$$index", -1}},
								"then": bson.M{
									"date": bson.M{"$dateFromParts": bson.M{
										"day":   "$$rangeDay",
										"hour":  9,
										"month": "$month1",
										"year":  "$year1"}},
								}}},
						"vars": bson.M{
							"index": bson.M{"$indexOfArray": []string{"$days.dom", "$$rangeDay"}}}}},
				"input": bson.M{"$range": []interface{}{1, 32, 1}}}}}})
	mainPipeline = append(mainPipeline, bson.M{"$unwind": "$days"})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{
		"date":     "$days.date",
		"quantity": "$days.quantity",
		"code":     "$days.code",
	}})

	//Aggregation
	d.Shared.BsonToJSONPrint(mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONWARD).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var Reports []models.CircleWiseHouseVisitedv2
	if err = cursor.All(ctx.CTX, &Reports); err != nil {
		return nil, err
	}

	return Reports, err
}

func (d *Daos) WardWiseHouseVisitedPercentage(ctx *models.Context, filter *models.FilterHouseVisited) ([]models.CircleWiseHouseVisitedv2, error) {
	mainPipeline := []bson.M{}
	//var sd, ed time.Time
	query := []bson.M{}
	if filter != nil {
		if len(filter.CircleNo) > 0 {
			query = append(query, bson.M{"code": bson.M{"$in": filter.CircleNo}})

		}
		if len(filter.WardNo) > 0 {
			query = append(query, bson.M{"code": bson.M{"$in": filter.WardNo}})

		}

	}
	startmonth := d.Shared.BeginningOfMonth(*filter.DateRange.From)
	sm := &startmonth
	endmonth := d.Shared.EndOfMonth(*filter.DateRange.From)
	em := &endmonth
	fmt.Println("sm===========>", sm)
	fmt.Println("em===========>", em)

	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	mainPipeline = append(mainPipeline,
		bson.M{"$lookup": bson.M{
			"from": "wardwisehousevisited",
			"as":   "report",
			"let":  bson.M{"code": "$code"},
			"pipeline": []bson.M{
				bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
					bson.M{"$eq": []string{"$status", "Active"}},
					bson.M{"$eq": []string{"$wardCode", "$$code"}},
					bson.M{"$gte": []interface{}{"$date", sm}},
					bson.M{"$lte": []interface{}{"$date", em}},
				}}}},
				bson.M{"$group": bson.M{"_id": bson.M{"$dateToString": bson.M{"format": "%Y-%m-%d", "date": "$date"}},
					"field":    bson.M{"$push": "$$ROOT.date"},
					"quantity": bson.M{"$sum": "$todayCollection"}}},
				bson.M{"$addFields": bson.M{"date": bson.M{"$arrayElemAt": []interface{}{"$field", 0}}}},
				bson.M{"$addFields": bson.M{"day": bson.M{"$dayOfMonth": "$date"}}},
			}}})
	mainPipeline = append(mainPipeline, bson.M{"$unwind": "$report"})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{
		"quantity": "$report.quantity",
		"date":     "$report.date",
	}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"dom": bson.M{"$dayOfMonth": "$report.date"}}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"month": bson.M{"$month": "$report.date"}}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"year": bson.M{"$year": "$report.date"}}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"dayOfWeek": bson.M{"$dayOfWeek": "$report.date"}}})
	mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"report.date": 1}})
	mainPipeline = append(mainPipeline, bson.M{"$group": bson.M{"_id": nil, "days": bson.M{"$push": "$$ROOT"}}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"month1": bson.M{"$arrayElemAt": []interface{}{"$days.month", 0}}}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"year1": bson.M{"$arrayElemAt": []interface{}{"$days.year", 0}}}})
	//mainPipeline = append(mainPipeline, bson.M{"$group": bson.M{"_id": nil, "days": bson.M{"$push": "$$ROOT"}}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{
		"days": bson.M{
			"$map": bson.M{
				"as": "rangeDay",
				"in": bson.M{
					"$let": bson.M{
						"in": bson.M{
							"$cond": bson.M{"else": bson.M{"$arrayElemAt": []string{"$days", "$$index"}},
								"if": bson.M{"$eq": []interface{}{"$$index", -1}},
								"then": bson.M{
									"date": bson.M{"$dateFromParts": bson.M{
										"day":   "$$rangeDay",
										"hour":  9,
										"month": "$month1",
										"year":  "$year1"}},
								}}},
						"vars": bson.M{
							"index": bson.M{"$indexOfArray": []string{"$days.dom", "$$rangeDay"}}}}},
				"input": bson.M{"$range": []interface{}{1, 32, 1}}}}}})
	mainPipeline = append(mainPipeline, bson.M{"$unwind": "$days"})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{
		"date":     "$days.date",
		"quantity": "$days.quantity",
		"code":     "$days.code",
	}})

	//Aggregation
	d.Shared.BsonToJSONPrint(mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONWARD).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var Reports []models.CircleWiseHouseVisitedv2
	if err = cursor.All(ctx.CTX, &Reports); err != nil {
		return nil, err
	}

	return Reports, err
}
