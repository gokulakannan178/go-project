package daos

import (
	"context"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// FilterBulkPrint : ""
func (d *Daos) BulkPrintGetDetailsForProperty(ctx *models.Context, filter *models.BulkPrintFilter) ([]models.BulkPrintDetail, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
		if len(filter.UserType) > 0 {
			query = append(query, bson.M{"type": bson.M{"$in": filter.UserType}})
		}

	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	var sd, ed time.Time
	if filter.Date != nil {
		sd = time.Date(filter.Date.Year(), filter.Date.Month(), filter.Date.Day(), 0, 0, 0, 0, filter.Date.Location())

		ed = time.Date(filter.Date.Year(), filter.Date.Month(), filter.Date.Day(), 23, 59, 59, 999999999, filter.Date.Location())

	}
	if filter.SortBy != "" {
		mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{filter.SortBy: filter.SortOrder}})
	}
	//Adding pagination if necessary
	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"from": "propertypayments",
			"as":   "propertypayments",
			"let":  bson.M{"username": "$userName"},
			"pipeline": []bson.M{
				{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
					{"$eq": []string{"$details.collector.id", "$$username"}},
					{"$eq": []string{"$status", "Completed"}},
					{"$gte": []interface{}{"$completionDate", sd}},
					{"$lte": []interface{}{"$completionDate", ed}},
				}}}},
				{"$group": bson.M{
					// "_id":      "$propertyId",
					"_id":      nil,
					"amount":   bson.M{"$sum": "$details.amount"},
					"payments": bson.M{"$sum": 1},
					"tnxId":    bson.M{"$push": "$tnxId"},
				}},

				// {"$group": bson.M{
				// 	"_id":      nil,
				// 	"amount":   bson.M{"$sum": "$amount"},
				// 	"payments": bson.M{"$sum": "$payments"},
				// 	"tnxId":    bson.M{"$push": "$tnxId"},
				// }},
			},
		},
	})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"propertypayments": bson.M{"$arrayElemAt": []interface{}{"$propertypayments", 0}}}})

	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"amount": "$propertypayments.amount"}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"tnxId": "$propertypayments.tnxId"}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"payments": "$propertypayments.payments"}})
	mainPipeline = append(mainPipeline, bson.M{"$project": bson.M{
		"name":     1,
		"userName": 1,
		"amount":   1,
		"tnxId":    1,
		"payments": 1,
	}})
	mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"payments": -1}})

	//Aggregation
	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONUSER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var BulkPrint []models.BulkPrintDetail
	if err = cursor.All(context.TODO(), &BulkPrint); err != nil {
		return nil, err
	}
	return BulkPrint, nil
}

// FilterBulkPrint : ""
func (d *Daos) BulkPrintGetDetailsForTradelicense(ctx *models.Context, filter *models.BulkPrintFilter) ([]models.BulkPrintDetail, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
		if len(filter.UserType) > 0 {
			query = append(query, bson.M{"type": bson.M{"$in": filter.UserType}})
		}

	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	var sd, ed time.Time
	if filter.Date != nil {
		sd = time.Date(filter.Date.Year(), filter.Date.Month(), filter.Date.Day(), 0, 0, 0, 0, filter.Date.Location())

		ed = time.Date(filter.Date.Year(), filter.Date.Month(), filter.Date.Day(), 23, 59, 59, 999999999, filter.Date.Location())

	}
	if filter.SortBy != "" {
		mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{filter.SortBy: filter.SortOrder}})
	}
	//Adding pagination if necessary
	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"from": "tradelicensepayments",
			"as":   "tradelicensepayments",
			"let":  bson.M{"username": "$userName"},
			"pipeline": []bson.M{
				{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
					{"$eq": []string{"$details.collector.by", "$$username"}},
					{"$eq": []string{"$status", "Completed"}},
					{"$gte": []interface{}{"$completionDate", sd}},
					{"$lte": []interface{}{"$completionDate", ed}},
				}}}},
				{"$group": bson.M{
					// "_id":      "$propertyId",
					"_id":      nil,
					"amount":   bson.M{"$sum": "$details.amount"},
					"payments": bson.M{"$sum": 1},
					"tnxId":    bson.M{"$push": "$tnxId"},
				}},

				// {"$group": bson.M{
				// 	"_id":      nil,
				// 	"amount":   bson.M{"$sum": "$amount"},
				// 	"payments": bson.M{"$sum": "$payments"},
				// 	"tnxId":    bson.M{"$push": "$tnxId"},
				// }},
			},
		},
	})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"tradelicensepayments": bson.M{"$arrayElemAt": []interface{}{"$tradelicensepayments", 0}}}})

	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"amount": "$tradelicensepayments.amount"}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"tnxId": "$tradelicensepayments.tnxId"}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"payments": "$tradelicensepayments.payments"}})
	mainPipeline = append(mainPipeline, bson.M{"$project": bson.M{
		"name":     1,
		"userName": 1,
		"amount":   1,
		"tnxId":    1,
		"payments": 1,
	}})

	//Aggregation
	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONUSER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var BulkPrint []models.BulkPrintDetail
	if err = cursor.All(context.TODO(), &BulkPrint); err != nil {
		return nil, err
	}
	return BulkPrint, nil
}
