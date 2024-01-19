package daos

import (
	"context"
	"fmt"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func (d *Daos) GetUserChargeSAFDashboard(ctx *models.Context, filter *models.GetUserChargeSAFDashboardFilter) (*models.UserChargeSAFDashboard, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$facet": bson.M{

		"pending": []bson.M{
			bson.M{"$match": bson.M{"userCharge.status": "Init", "userCharge.isUserCharge": "Yes"}},
			bson.M{"$group": bson.M{"_id": nil, "count": bson.M{"$sum": 1}}},
		},
		"active": []bson.M{
			bson.M{"$match": bson.M{"userCharge.status": "Active", "userCharge.isUserCharge": "Yes"}},
			bson.M{"$group": bson.M{"_id": nil, "count": bson.M{"$sum": 1}}},
		},
		"rejected": []bson.M{
			bson.M{"$match": bson.M{"userCharge.status": "NotApproved", "userCharge.isUserCharge": "Yes"}},
			bson.M{"$group": bson.M{"_id": nil, "count": bson.M{"$sum": 1}}},
		},
		"deleted": []bson.M{
			bson.M{"$match": bson.M{
				"status": "Deleted", "userCharge.userCharge.isUserCharge": "Yes",
			}},
			bson.M{"$group": bson.M{"_id": nil, "count": bson.M{"$sum": 1}}},
		},
	}},
		bson.M{
			"$addFields": bson.M{

				"pending":  bson.M{"$arrayElemAt": []interface{}{"$pending", 0}},
				"active":   bson.M{"$arrayElemAt": []interface{}{"$active", 0}},
				"rejected": bson.M{"$arrayElemAt": []interface{}{"$rejected", 0}},
				"deleted":  bson.M{"$arrayElemAt": []interface{}{"$deleted", 0}},
			},
		},
	)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("GetUserChargeSAFDashboard query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var results []models.UserChargeSAFDashboard
	var result *models.UserChargeSAFDashboard
	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}
	if len(results) > 0 {
		result = &results[0]
	}
	return result, nil
}

// DashBoardStatusWiseShopRentCollectionAndChart : ""
func (d *Daos) UserwiseUserChargeReport(ctx *models.Context, filter *models.UserFilter) ([]models.UserwiseUsercharge, error) {

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
	t := time.Now()
	if filter != nil {
		if filter.DateRange == nil {
			sd = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
			ed = time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, t.Location())
		}
		if filter.DateRange.From != nil {
			sd = time.Date(filter.DateRange.From.Year(), filter.DateRange.From.Month(), filter.DateRange.From.Day(), 0, 0, 0, 0, filter.DateRange.From.Location())
			// var ed time.Time
			if filter.DateRange.To != nil {
				ed = time.Date(filter.DateRange.To.Year(), filter.DateRange.To.Month(), filter.DateRange.To.Day(), 23, 59, 59, 0, filter.DateRange.To.Location())
			} else {
				ed = time.Date(filter.DateRange.From.Year(), filter.DateRange.From.Month(), filter.DateRange.From.Day(), 23, 59, 59, 0, filter.DateRange.From.Location())
			}
		}
	}
	fmt.Println("sd ===>", sd)
	fmt.Println("ed ===>", ed)
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{
		"as":   "userchargepayments",
		"from": "userchargepayments",
		"let":  bson.M{"userId": "$userName"},
		"pipeline": []bson.M{
			bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
				bson.M{"$eq": []string{"$status", "Completed"}},
				bson.M{"$eq": []string{"$details.collector.by", "$$userId"}},
				bson.M{"$gte": []interface{}{"$completionDate", sd}},
				bson.M{"$lte": []interface{}{"$completionDate", ed}},
			}}}},
			bson.M{"$facet": bson.M{
				"cash": []bson.M{
					bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
						bson.M{"$eq": []string{"$details.mop.mode", "Cash"}},
					}}}},
					bson.M{"$group": bson.M{"_id": nil,
						"count":       bson.M{"$sum": 1},
						"totalAmount": bson.M{"$sum": "$details.amount"},
					}}},
				"cheque": []bson.M{
					bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
						bson.M{"$eq": []string{"$details.mop.mode", "Cheque"}},
					}}}},
					bson.M{"$group": bson.M{"_id": nil,
						"count":       bson.M{"$sum": 1},
						"totalAmount": bson.M{"$sum": "$details.amount"},
					}}},
				"netbanking": []bson.M{
					bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
						bson.M{"$eq": []string{"$details.mop.mode", "NB"}},
					}}}},
					bson.M{"$group": bson.M{"_id": nil,
						"count":       bson.M{"$sum": 1},
						"totalAmount": bson.M{"$sum": "$details.amount"},
					}}},
				"dd": []bson.M{
					bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
						bson.M{"$eq": []string{"$details.mop.mode", "DD"}},
					}}}},
					bson.M{"$group": bson.M{"_id": nil,
						"count":       bson.M{"$sum": 1},
						"totalAmount": bson.M{"$sum": "$details.amount"},
					}}},
			}},
			bson.M{"$addFields": bson.M{"cash": bson.M{"$arrayElemAt": []interface{}{"$cash", 0}}}},
			bson.M{"$addFields": bson.M{"cheque": bson.M{"$arrayElemAt": []interface{}{"$cheque", 0}}}},
			bson.M{"$addFields": bson.M{"netbanking": bson.M{"$arrayElemAt": []interface{}{"$netbanking", 0}}}},
			bson.M{"$addFields": bson.M{"dd": bson.M{"$arrayElemAt": []interface{}{"$dd", 0}}}},
		}}},
	)
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"userchargepayments": bson.M{"$arrayElemAt": []interface{}{"$userchargepayments", 0}}}})

	//Aggregation
	d.Shared.BsonToJSONPrintTag("shoprent Query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONUSER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var ddac []models.UserwiseUsercharge
	if err := cursor.All(ctx.CTX, &ddac); err != nil {
		return nil, err
	}
	return ddac, nil

}
