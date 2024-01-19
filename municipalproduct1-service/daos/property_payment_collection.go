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

func (d *Daos) CalcualteTotalCollectionForAllPayments(ctx *models.Context) error {
	var data []models.PropertyPaymentCalculate
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"status": bson.M{"$in": []string{"Completed"}}}})
	mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{

		"from":         "propertypaymentfys",
		"as":           "fys",
		"localField":   "tnxId",
		"foreignField": "tnxId",
	}})

	//Aggregation
	d.Shared.BsonToJSONPrintTag("collection dashboard chart query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYPAYMENT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return err
	}
	if err = cursor.All(context.TODO(), &data); err != nil {
		return err
	}
	if data != nil {
		for _, v := range data {
			fmt.Print("prop - ", v.PropertyID, " status - ", v.Status, " tranx - ", v.TnxID, " len fy", len(v.FYs))
			var arrear, current, total float64
			isCurrentAvailable := false
			for _, v2 := range v.FYs {
				fmt.Print(" ", v2.FY.IsCurrent)
				if v2.FY.IsCurrent == true {
					fmt.Print(" v2.IsCurrent == true")
					current = current + v2.FY.TotalTax
					isCurrentAvailable = true
				} else {
					arrear = arrear + v2.FY.TotalTax
				}

			}
			total = total + v.Demand.TotalTax
			if isCurrentAvailable {
				current = current + v.Demand.BoreCharge + v.Demand.FormFee
			} else {
				arrear = arrear + v.Demand.BoreCharge + v.Demand.FormFee
			}
			fmt.Println(" ==> A-", arrear, "C-", current, "T-", total)
			_, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYPAYMENT).UpdateOne(ctx.CTX,
				bson.M{"tnxId": v.TnxID},
				bson.M{"$set": bson.M{"demand.arrear": arrear, "demand.current": current}},
			)
			if err != nil {
				return err
			}
			// fmt.Println(res)
		}
	}
	return nil
}

func (d *Daos) PropertyUpdateCollection(ctx *models.Context, propertyID string) error {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"propertyId": propertyID, "status": "Completed"}})
	mainPipeline = append(mainPipeline, bson.M{"$group": bson.M{"_id": nil,
		"arrear":   bson.M{"$sum": "$demand.arrear"},
		"current":  bson.M{"$sum": "$demand.current"},
		"totalTax": bson.M{"$sum": "$demand.totalTax"},
		"other":    bson.M{"$sum": bson.M{"$add": []string{"$demand.formFee", "$boreCharge"}}},
		"penalty":  bson.M{"$sum": "$demand.penalCharge"},
	}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{
		"data.arrear":   "$arrear",
		"data.current":  "$current",
		"data.totalTax": "$totalTax",
		"data.other":    "$other",
		"data.penalty":  "$penalty",
	}})

	//Aggregation
	d.Shared.BsonToJSONPrintTag("property query =>", mainPipeline)

	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYPAYMENT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return err
	}
	var data []models.PropertyUpdateCollection
	if err = cursor.All(ctx.CTX, &data); err != nil {
		return err
	}
	if len(data) > 0 {
		datum := data[0]
		selector := bson.M{"uniqueId": propertyID}
		update := bson.M{"$set": bson.M{"collection": datum.Data}}
		fmt.Println("updating property - " + propertyID)
		res, err := ctx.DB.Collection(constants.COLLECTIONPROPERTY).UpdateOne(ctx.CTX, selector, update)
		if err != nil {
			return errors.New("Error in updating collection in property - " + err.Error())
		}
		fmt.Println(res)
	} else {
		updateData := models.UpdateCollection{}
		fmt.Println("updating 0")
		selector := bson.M{"uniqueId": propertyID}
		update := bson.M{"$set": bson.M{"collection": updateData}}
		fmt.Println("updating property - " + propertyID)
		res, err := ctx.DB.Collection(constants.COLLECTIONPROPERTY).UpdateOne(ctx.CTX, selector, update)
		if err != nil {
			return errors.New("Error in updating collection in property - " + err.Error())
		}
		fmt.Println(res)

	}
	return nil
}

//DashboardTotalCollectionChart : ""
func (d *Daos) DashboardTotalCollectionChart(ctx *models.Context, filter *models.DashboardTotalCollectionChartFilter) ([]models.DashboardTotalCollectionChart, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {

	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"fyOrder": 1}})
	mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{
		"from": "financialyears",
		"as":   "fy",
		"let":  bson.M{"fyId": filter.Fy},
		"pipeline": []bson.M{
			bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
				bson.M{"$eq": []string{"$uniqueId", "$$fyId"}},
			}}}},
		},
	}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"fy": bson.M{"$arrayElemAt": []interface{}{"$fy", 0}}}})

	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"sampleDate": bson.M{"$dateFromParts": bson.M{"day": 30, "month": "$month", "year": bson.M{
		"$cond": bson.M{"if": bson.M{"$lte": []interface{}{"$month", 3}}, "then": bson.M{"$year": "$fy.to"}, "else": bson.M{"$year": "$fy.from"}},
	}}}}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{
		"sd": bson.M{"$dateFromParts": bson.M{"year": bson.M{"$year": "$sampleDate"}, "month": "$month", "day": 1, "hour": 0, "minute": 0, "second": 0, "millisecond": 0}},
		"ed": bson.M{"$dateFromParts": bson.M{"year": bson.M{"$year": "$sampleDate"}, "month": bson.M{"$sum": []interface{}{"$month", 1}}, "day": 0, "hour": 23, "minute": 59, "second": 59, "millisecond": 0}},
	}})

	mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{
		"from": "propertypayments",
		"as":   "payments",
		"let":  bson.M{"sd": "$sd", "ed": "$ed"},
		"pipeline": []bson.M{
			bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
				bson.M{"$eq": []string{"$status", "Completed"}},
				bson.M{"$gte": []string{"$completionDate", "$$sd"}},
				bson.M{"$lte": []string{"$completionDate", "$$ed"}},
			}}}},
			bson.M{"$group": bson.M{"_id": nil, "collection": bson.M{"$sum": "$details.amount"}}},
		},
	}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"payments": bson.M{"$arrayElemAt": []interface{}{"$payments", 0}}}})
	//Aggregation
	d.Shared.BsonToJSONPrintTag("collection dashboard chart query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONMONTH).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var data []models.DashboardTotalCollectionChart
	if err = cursor.All(context.TODO(), &data); err != nil {
		return nil, err
	}
	if ctx.ProductConfig.LocationID == "Munger" {
		for k, v := range data {
			if filter.Fy == "GEN2023_2024" {
				if v.Name == "June" {
					data[k].Payments.Collection = data[k].Payments.Collection - 3817620

				}
			}

		}
	}

	return data, nil
}

func (d *Daos) DashboardTotalCollectionOverview(ctx *models.Context, filter *models.DashboardTotalCollectionOverviewFilter) (models.DashboardTotalCollectionOverview, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
	}

	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	mainPipeline = append(mainPipeline,
		bson.M{"$lookup": bson.M{
			"from": "propertypaymentfys",
			"as":   "propertypaymentfys",
			"let":  bson.M{"tnxId": "$tnxId"},
			"pipeline": []bson.M{
				bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
					bson.M{"$eq": []string{"$$tnxId", "$tnxId"}},
				}}}},

				bson.M{"$group": bson.M{
					"_id":             nil,
					"arrearTotalTax":  bson.M{"$sum": bson.M{"$cond": bson.M{"if": bson.M{"$eq": []interface{}{"$fy.isCurrent", false}}, "then": "$fy.totalTax", "else": 0}}},
					"currentTotalTax": bson.M{"$sum": bson.M{"$cond": bson.M{"if": bson.M{"$eq": []interface{}{"$fy.isCurrent", true}}, "then": "$fy.totalTax", "else": 0}}},
				}},
			},
		}})
	mainPipeline = append(mainPipeline,
		bson.M{"$addFields": bson.M{"propertypaymentfys": bson.M{"$arrayElemAt": []interface{}{"$propertypaymentfys", 0}}}})

	mainPipeline = append(mainPipeline,
		bson.M{"$addFields": bson.M{"genOtherCharges": bson.M{"$sum": []string{"$demand.formFee", "$demand.boreCharge"}},
			"checkingCurrent": bson.M{"$cond": bson.M{"if": bson.M{"$eq": []interface{}{"$propertypaymentfys.currentTotalTax", 0}}, "then": false, "else": true}},
		}})

	mainPipeline = append(mainPipeline,
		bson.M{"$addFields": bson.M{
			"propertypaymentfys.arrearTotalTax": bson.M{"$sum": []interface{}{"$propertypaymentfys.arrearTotalTax",
				bson.M{"$cond": bson.M{"if": bson.M{"$eq": []interface{}{"$checkingCurrent", false}}, "then": "$genOtherCharges", "else": 0}},
			}},
			"propertypaymentfys.currentTotalTax": bson.M{"$sum": []interface{}{"$propertypaymentfys.currentTotalTax",
				bson.M{"$cond": bson.M{"if": bson.M{"$eq": []interface{}{"$checkingCurrent", true}}, "then": "$genOtherCharges", "else": 0}},
			}},
		}})

	mainPipeline = append(mainPipeline,
		bson.M{"$group": bson.M{
			"_id": nil, "arrearTotalTax": bson.M{"$sum": "$propertypaymentfys.arrearTotalTax"},
			"currentTotalTax": bson.M{"$sum": "$propertypaymentfys.currentTotalTax"},
		}})

	//Aggregation
	d.Shared.BsonToJSONPrint(mainPipeline)
	var emptyData models.DashboardTotalCollectionOverview

	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYPAYMENT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return emptyData, err
	}
	var data []models.DashboardTotalCollectionOverview

	if err = cursor.All(context.TODO(), &data); err != nil {
		return emptyData, err
	}
	if len(data) > 0 {
		return data[0], nil
	}

	return emptyData, nil
}

func (d *Daos) DashboardDayWiseCollectionChart(ctx *models.Context, filter *models.DashboardDayWiseCollectionChartFilter) (models.DashboardDayWiseCollectionChart, error) {

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
	mainPipeline = append(mainPipeline,
		bson.M{"$lookup": bson.M{
			"from": "propertypaymentfys",
			"as":   "propertypaymentfys",
			"let":  bson.M{"tnxId": "$tnxId"},
			"pipeline": []bson.M{
				bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
					bson.M{"$eq": []string{"$$tnxId", "$tnxId"}},
				}}}},

				bson.M{"$group": bson.M{
					"_id":            nil,
					"currentPenalty": bson.M{"$sum": bson.M{"$cond": bson.M{"if": bson.M{"$eq": []interface{}{"$fy.isCurrent", true}}, "then": "$fy.penanty", "else": 0}}},
					"arrearPenalty":  bson.M{"$sum": bson.M{"$cond": bson.M{"if": bson.M{"$eq": []interface{}{"$fy.isCurrent", false}}, "then": "$fy.penanty", "else": 0}}},
					"rebateAmount":   bson.M{"$sum": bson.M{"$cond": bson.M{"if": bson.M{"$eq": []interface{}{"$fy.isCurrent", true}}, "then": "$fy.rebate", "else": 0}}},
				}},
			},
		}})
	mainPipeline = append(mainPipeline,
		bson.M{"$addFields": bson.M{"propertypaymentfys": bson.M{"$arrayElemAt": []interface{}{"$propertypaymentfys", 0}}}})

	mainPipeline = append(mainPipeline, bson.M{"$group": bson.M{
		"_id":               bson.M{"day": bson.M{"$dayOfMonth": "$completionDate"}},
		"totalTax":          bson.M{"$sum": "$demand.totalTax"},
		"propertyCount":     bson.M{"$sum": 1},
		"arrearCollection":  bson.M{"$sum": "$demand.arrear"},
		"arrearPenalty":     bson.M{"$sum": "$propertypaymentfys.arrearPenalty"},
		"currentCollection": bson.M{"$sum": "$demand.current"},
		"totalCollection":   bson.M{"$sum": "$details.amount"},
		"currentPenalty":    bson.M{"$sum": "$propertypaymentfys.currentPenalty"},
		"rebateAmount":      bson.M{"$sum": "$propertypaymentfys.rebateAmount"},
	}})

	mainPipeline = append(mainPipeline, bson.M{"$group": bson.M{"_id": nil, "days": bson.M{"$push": "$$ROOT"}}})

	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{
		"records": bson.M{
			"$map": bson.M{
				"input": bson.M{"$range": []interface{}{sd.Day(), ed.Day() + 1, 1}},
				"as":    "rangeDay",
				"in": bson.M{
					"$let": bson.M{
						"vars": bson.M{"index": bson.M{"$indexOfArray": []string{"$days._id.day", "$$rangeDay"}}},
						"in": bson.M{
							"$cond": bson.M{
								"if":   bson.M{"$eq": []interface{}{"$$index", -1}},
								"then": bson.M{"_id": bson.M{"day": "$$rangeDay"}},
								"else": bson.M{"$arrayElemAt": []string{"$days", "$$index"}}},
						},
					},
				},
			},
		},
	}})

	//Aggregation
	d.Shared.BsonToJSONPrint(mainPipeline)
	var emptyData models.DashboardDayWiseCollectionChart

	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYPAYMENT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return emptyData, err
	}
	var data []models.DashboardDayWiseCollectionChart

	if err = cursor.All(context.TODO(), &data); err != nil {
		return emptyData, err
	}
	if len(data) > 0 {
		return data[0], nil
	}

	return emptyData, nil

}

func (d *Daos) WardWiseCollectionReport(ctx *models.Context, filter *models.WardWiseCollectionReportFilter, pagination *models.Pagination) ([]models.WardWiseCollectionReport, error) {

	mainPipeline := []bson.M{}
	var sd, ed *time.Time
	query := []bson.M{}
	if filter != nil {
		if len(filter.Zone) > 0 {
			query = append(query, bson.M{"zoneCode": bson.M{"$in": filter.Zone}})
		}
		if len(filter.Ward) > 0 {
			query = append(query, bson.M{"code": bson.M{"$in": filter.Ward}})
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

	var propertypaymentsPipelineAnd = []bson.M{
		bson.M{"$eq": []string{"$$wardNo", "$address.wardCode"}},
		bson.M{"$eq": []string{"$status", "Completed"}},
	}

	if sd != nil {
		propertypaymentsPipelineAnd = append(propertypaymentsPipelineAnd,
			bson.M{"$gte": []string{"$completionDate", "$$sd"}},
			bson.M{"$lte": []string{"$completionDate", "$$ed"}})
	}

	mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{
		"from": "propertypayments",
		"as":   "propertypayments",
		"let":  bson.M{"wardNo": "$code", "sd": sd, "ed": ed},
		"pipeline": []bson.M{
			bson.M{"$match": bson.M{"$expr": bson.M{"$and": propertypaymentsPipelineAnd}}},
			bson.M{"$group": bson.M{"_id": "$propertyId", "payments": bson.M{"$sum": "$demand.totalTax"}}},
			bson.M{"$group": bson.M{"_id": nil, "properties": bson.M{"$sum": 1}, "payments": bson.M{"$sum": "$payments"}}},
		},
	}})

	mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{
		"from": "properties",
		"as":   "properties",
		"let":  bson.M{"wardNo": "$code"},
		"pipeline": []bson.M{
			bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
				bson.M{"$eq": []string{"$$wardNo", "$address.wardCode"}},
				bson.M{"$in": []interface{}{"$status", []string{"Active", "Init"}}},
			}}}},
			bson.M{"$group": bson.M{"_id": nil, "count": bson.M{"$sum": 1}}},
		},
	}})

	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{
		"propertypayments": bson.M{"$arrayElemAt": []interface{}{"$propertypayments", 0}},
		"properties":       bson.M{"$arrayElemAt": []interface{}{"$properties", 0}},
	}})

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

	//Aggregation
	d.Shared.BsonToJSONPrint(mainPipeline)
	var data []models.WardWiseCollectionReport

	cursor, err := ctx.DB.Collection(constants.COLLECTIONWARD).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}

	if err = cursor.All(context.TODO(), &data); err != nil {
		return nil, err
	}

	return data, nil
}

func (d *Daos) TCCollectionSummaryReport(ctx *models.Context, filter *models.TCCollectionSummaryFilter) ([]models.TCCollectionSummaryReport, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.UserType) > 0 {
			query = append(query, bson.M{"type": bson.M{"$in": filter.UserType}})
		}
		if len(filter.UserStatus) > 0 {
			//query = append(query, bson.M{"status": bson.M{"$in": filter.UserStatus}})
		}
		if len(filter.UserName) > 0 {
			query = append(query, bson.M{"userName": bson.M{"$in": filter.UserName}})
		}
	}
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{"from": "users", "as": "manager", "localField": "managerId", "foreignField": "userName"}})
	propertyPaymentsPipelineQuery := []bson.M{}

	propertyPaymentsPipelineQuery = append(propertyPaymentsPipelineQuery, bson.M{"$eq": []string{"$$userId", "$details.collector.id"}})
	propertyPaymentsPipelineQuery = append(propertyPaymentsPipelineQuery, bson.M{"$eq": []string{"$status", "Completed"}})
	if filter != nil {
		if filter.DateRange != nil {
			//var sd,ed time.Time
			if filter.DateRange.From != nil {
				sd := time.Date(filter.DateRange.From.Year(), filter.DateRange.From.Month(), filter.DateRange.From.Day(), 0, 0, 0, 0, filter.DateRange.From.Location())
				ed := time.Date(filter.DateRange.From.Year(), filter.DateRange.To.Month(), filter.DateRange.To.Day(), 23, 59, 59, 0, filter.DateRange.To.Location())
				if filter.DateRange.To != nil {
					ed = time.Date(filter.DateRange.To.Year(), filter.DateRange.To.Month(), filter.DateRange.To.Day(), 23, 59, 59, 0, filter.DateRange.To.Location())
				}
				propertyPaymentsPipelineQuery = append(propertyPaymentsPipelineQuery, bson.M{"$gte": []interface{}{"$completionDate", sd}})
				propertyPaymentsPipelineQuery = append(propertyPaymentsPipelineQuery, bson.M{"$lte": []interface{}{"$completionDate", ed}})

			}
		}
	}

	mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{
		"from": "propertypayments",
		"as":   "propertypayments",
		"let":  bson.M{"userId": "$userName"},
		"pipeline": []bson.M{
			func() bson.M {
				if len(propertyPaymentsPipelineQuery) > 0 {
					return bson.M{"$match": bson.M{
						"$expr": bson.M{"$and": propertyPaymentsPipelineQuery},
					}}
				}
				return bson.M{"$match": bson.M{}}
			}(),
			// bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
			// 	bson.M{"$eq": []string{"$$userId", "$details.collector.id"}},
			// 	bson.M{"$eq": []string{"$status", "Completed"}},

			// }}}},
			bson.M{"$sort": bson.M{"completionDate": -1}},
			bson.M{"$group": bson.M{"_id": "$propertyId", "recentTransaction": bson.M{"$first": "$completionDate"}, "payments": bson.M{"$sum": "$details.amount"}}},
			bson.M{"$sort": bson.M{"recentTransaction": -1}},
			bson.M{"$group": bson.M{"_id": nil, "propertyCount": bson.M{"$sum": 1}, "recentTransaction": bson.M{"$first": "$recentTransaction"}, "payments": bson.M{"$sum": "$payments"}}},
		},
	}})

	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{
		"manager":          bson.M{"$arrayElemAt": []interface{}{"$manager", 0}},
		"propertypayments": bson.M{"$arrayElemAt": []interface{}{"$propertypayments", 0}},
	}})
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDUSERTYPE, "type", "uniqueId", "userType", "userType")...)
	//Aggregation
	d.Shared.BsonToJSONPrint(mainPipeline)
	var data []models.TCCollectionSummaryReport

	cursor, err := ctx.DB.Collection(constants.COLLECTIONUSER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}

	if err = cursor.All(context.TODO(), &data); err != nil {
		return nil, err
	}

	return data, nil
}
func (d *Daos) TCCollectionSummaryReportV2(ctx *models.Context, filter *models.TCCollectionSummaryFilter) ([]models.TCCollectionSummaryReportV2, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.UserType) > 0 {
			query = append(query, bson.M{"type": bson.M{"$in": filter.UserType}})
		}
		if len(filter.UserStatus) > 0 {
			//query = append(query, bson.M{"status": bson.M{"$in": filter.UserStatus}})
		}
		if len(filter.UserName) > 0 {
			query = append(query, bson.M{"userName": bson.M{"$in": filter.UserName}})
		}
	}
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{"from": "users", "as": "manager", "localField": "managerId", "foreignField": "userName"}})
	propertyPaymentsPipelineQuery := []bson.M{}

	propertyPaymentsPipelineQuery = append(propertyPaymentsPipelineQuery, bson.M{"$eq": []string{"$$userId", "$details.collector.id"}})
	propertyPaymentsPipelineQuery = append(propertyPaymentsPipelineQuery, bson.M{"$eq": []string{"$status", "Completed"}})
	if filter != nil {
		if filter.DateRange != nil {
			//var sd,ed time.Time
			if filter.DateRange.From != nil {
				sd := time.Date(filter.DateRange.From.Year(), filter.DateRange.From.Month(), filter.DateRange.From.Day(), 0, 0, 0, 0, filter.DateRange.From.Location())
				ed := time.Date(filter.DateRange.From.Year(), filter.DateRange.To.Month(), filter.DateRange.To.Day(), 23, 59, 59, 0, filter.DateRange.To.Location())
				if filter.DateRange.To != nil {
					ed = time.Date(filter.DateRange.To.Year(), filter.DateRange.To.Month(), filter.DateRange.To.Day(), 23, 59, 59, 0, filter.DateRange.To.Location())
				}
				propertyPaymentsPipelineQuery = append(propertyPaymentsPipelineQuery, bson.M{"$gte": []interface{}{"$completionDate", sd}})
				propertyPaymentsPipelineQuery = append(propertyPaymentsPipelineQuery, bson.M{"$lte": []interface{}{"$completionDate", ed}})

			}
		}
	}

	mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{
		"from": "propertypayments",
		"as":   "propertypayments",
		"let":  bson.M{"userId": "$userName"},
		"pipeline": []bson.M{
			func() bson.M {
				if len(propertyPaymentsPipelineQuery) > 0 {
					return bson.M{"$match": bson.M{
						"$expr": bson.M{"$and": propertyPaymentsPipelineQuery},
					}}
				}
				return bson.M{"$match": bson.M{}}
			}(),
			// bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
			// 	bson.M{"$eq": []string{"$$userId", "$details.collector.id"}},
			// 	bson.M{"$eq": []string{"$status", "Completed"}},

			// }}}},
			bson.M{"$sort": bson.M{"completionDate": -1}},
			bson.M{"$group": bson.M{"_id": "$propertyId", "recentTransaction": bson.M{"$first": "$completionDate"}, "payments": bson.M{"$sum": "$demand.totalTax"}}},
			bson.M{"$sort": bson.M{"recentTransaction": -1}},
			bson.M{"$group": bson.M{"_id": nil, "propertyCount": bson.M{"$sum": 1}, "recentTransaction": bson.M{"$first": "$recentTransaction"}, "payments": bson.M{"$sum": "$payments"}}},
		},
	}})

	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{
		"manager":          bson.M{"$arrayElemAt": []interface{}{"$manager", 0}},
		"propertypayments": bson.M{"$arrayElemAt": []interface{}{"$propertypayments", 0}},
	}})
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDUSERTYPE, "type", "uniqueId", "userType", "userType")...)
	//Aggregation
	d.Shared.BsonToJSONPrint(mainPipeline)
	var data []models.TCCollectionSummaryReportV2

	cursor, err := ctx.DB.Collection(constants.COLLECTIONUSER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}

	if err = cursor.All(context.TODO(), &data); err != nil {
		return nil, err
	}

	return data, nil
}

// PropertyMonthWiseCollectionReportFilter : "this api is used for getting the records of totalarrear and totalcollection for particular year"
func (d *Daos) FilterPropertyMonthWiseCollectionReport(ctx *models.Context, filter *models.PropertyMonthWiseCollectionReportFilter) ([]models.PropertyMonthWiseCollectionReport, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	resFYs, err := d.GetSingleFinancialYear(ctx, filter.FYID)
	if err != nil {
		return nil, err
	}
	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
	}
	var sd, ed *time.Time
	fmt.Println("resFYs from ===>", resFYs.From)
	fmt.Println("resFYs to ===>", resFYs.To)
	sd = resFYs.From
	ed = resFYs.To
	fmt.Println("sd ===>", sd)
	fmt.Println("ed ===>", ed)
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	// lookup
	mainPipeline = append(mainPipeline, bson.M{"$group": bson.M{"_id": bson.M{
		"month": bson.M{"$month": "$completionDate"}},
		"propertyCount":         bson.M{"$sum": 1},
		"currentAmount":         bson.M{"$sum": "$demand.current"},
		"arrearAmount":          bson.M{"$sum": "$demand.arrear"},
		"totalDetailsAmount":    bson.M{"$sum": "$details.amount"},
		"totalCollectionAmount": bson.M{"$sum": "$demand.totalTax"}}})

	mainPipeline = append(mainPipeline, bson.M{"$group": bson.M{"_id": nil, "months": bson.M{"$push": "$$ROOT"}}})

	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{
		"records": bson.M{
			"$map": bson.M{
				"input": bson.M{"$range": []interface{}{1, 13, 1}},
				// "input": bson.M{"$range": []interface{}{sd.Month(), ed.Month() + 1, 1}},
				"as": "rangeDay",
				"in": bson.M{
					"$let": bson.M{
						"vars": bson.M{"index": bson.M{"$indexOfArray": []string{"$months._id.month", "$$rangeDay"}}},
						"in": bson.M{
							"$cond": bson.M{
								"if":   bson.M{"$eq": []interface{}{"$$index", -1}},
								"then": bson.M{"_id": bson.M{"month": "$$rangeDay"}},
								"else": bson.M{"$arrayElemAt": []string{"$months", "$$index"}}},
						},
					},
				},
			},
		},
	}})

	//Aggregation
	d.Shared.BsonToJSONPrint(mainPipeline)

	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYPAYMENT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var res []models.PropertyMonthWiseCollectionReport
	if err = cursor.All(context.TODO(), &res); err != nil {
		return nil, err
	}
	return res, nil

}

// FilterPropertyWiseCollectionReport : ""
func (d *Daos) FilterPropertyWiseCollectionReport(ctx *models.Context, filter *models.PropertyWiseCollectionReportFilter) ([]models.PropertyWiseCollectionReport, error) {

	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		// if len(filter.Status) > 0 {
		// 	query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		// }
	}
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"_id": -1}})

	mainPipeline = append(mainPipeline, d.PropertyOwnersLookup(constants.COLLECTIONPROPERTYOWNER, "uniqueId", "propertyId", "ref.propertyOwner", "ref.propertyOwner")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONWARD, "address.wardCode", "code", "ref.address.ward", "ref.address.ward")...)

	mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{
		"from": constants.COLLECTIONPROPERTYPAYMENTFY,
		"as":   "ref.payments",
		"let":  bson.M{"propertyId": "$uniqueId"},
		"pipeline": []bson.M{
			bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
				bson.M{"$eq": []string{"$status", "Completed"}},
				bson.M{"$eq": []string{"$$propertyId", "$propertyId"}},
			}}}},
			bson.M{"$sort": bson.M{"fy.from": 1}},
			bson.M{"$group": bson.M{"_id": nil,
				"fyFrom":          bson.M{"$first": bson.M{"$dateToString": bson.M{"format": "%Y-%m-%d", "date": "$fy.from"}}},
				"fyTo":            bson.M{"$last": bson.M{"$dateToString": bson.M{"format": "%Y-%m-%d", "date": "$fy.to"}}},
				"totalTax":        bson.M{"$sum": "$fy.tax"},
				"totalCollection": bson.M{"$sum": "$fy.totalTax"},
				"totalPenalty":    bson.M{"$sum": "$fy.penanty"},
			},
			},
		},
	}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"ref.payments": bson.M{"$arrayElemAt": []interface{}{"$ref.payments", 0}}}})

	//Aggregation
	d.Shared.BsonToJSONPrint(mainPipeline)

	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var res []models.PropertyWiseCollectionReport
	if err = cursor.All(context.TODO(), &res); err != nil {
		return nil, err
	}
	return res, nil

}
