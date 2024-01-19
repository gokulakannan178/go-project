package daos

import (
	"logikoof-echalan-service/constants"
	"logikoof-echalan-service/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

//PaymentWidget : ""
func (d *Daos) PaymentWidget(ctx *models.Context, paymentWidgetFilter *models.PaymentWidgetFilter) (*models.PaymentWidget, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}

	if paymentWidgetFilter != nil {
		if len(paymentWidgetFilter.VehicleType) > 0 {
			query = append(query, bson.M{"vehicle.type": bson.M{"$in": paymentWidgetFilter.VehicleType}})
		}
		if len(paymentWidgetFilter.OffenceType) > 0 {
			query = append(query, bson.M{"offenceType": bson.M{"$in": paymentWidgetFilter.OffenceType}})
		}
		if paymentWidgetFilter.DateRange != nil {
			if paymentWidgetFilter.DateRange.From != nil {
				ft := *paymentWidgetFilter.DateRange.From
				sd := time.Date(ft.Year(), ft.Month(), ft.Day(), 0, 0, 0, 0, ft.Location())
				ed := time.Now()
				if paymentWidgetFilter.DateRange.To != nil {
					tt := *paymentWidgetFilter.DateRange.To
					ed = time.Date(tt.Year(), tt.Month(), tt.Day(), 11, 59, 59, 0, tt.Location())

				}
				query = append(query, bson.M{"offenceDate": bson.M{"$gte": sd, "$lte": ed}})
			}
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	//Calculating Recovered and Pending Payments
	mainPipeline = append(mainPipeline, bson.M{"$group": bson.M{"_id": "$payment.status", "total": bson.M{"$sum": "$pelalty"}}})
	mainPipeline = append(mainPipeline, bson.M{"$group": bson.M{"_id": nil, "arr": bson.M{"$push": bson.M{"k": "$_id", "v": "$total"}}}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"payments": bson.M{"$arrayToObject": "$arr"}}})
	//Aggregation
	d.Shared.BsonToJSONPrintTag("Payment widget dashboard query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLVEHICLECHALLAN).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var paymentWidgets []models.PaymentWidget
	var paymentWidget *models.PaymentWidget
	if err = cursor.All(ctx.CTX, &paymentWidgets); err != nil {
		return nil, err
	}
	if len(paymentWidgets) > 0 {
		paymentWidget = &paymentWidgets[0]
	}
	return paymentWidget, nil
}

//TodaysOffenceWidget : ""
func (d *Daos) TodaysOffenceWidget(ctx *models.Context, filter *models.TodaysOffenceWidgetFilter) (*models.TodaysOffenceWidget, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}

	if filter != nil {
		if filter.Date != nil {
			ft := *filter.Date
			sd := time.Date(ft.Year(), ft.Month(), ft.Day(), 0, 0, 0, 0, ft.Location())
			ed := time.Date(ft.Year(), ft.Month(), ft.Day(), 23, 59, 59, 0, ft.Location())
			query = append(query, bson.M{"offenceDate": bson.M{"$gte": sd, "$lte": ed}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	//Calculating Recovered and Pending Payments
	mainPipeline = append(mainPipeline, bson.M{"$group": bson.M{"_id": "$vehicle.type", "total": bson.M{"$sum": 1}}})
	mainPipeline = append(mainPipeline, bson.M{"$group": bson.M{"_id": nil, "arr": bson.M{"$push": bson.M{"k": "$_id", "v": "$total"}}}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"types": bson.M{"$arrayToObject": "$arr"}}})

	//Aggregation
	d.Shared.BsonToJSONPrintTag("Payment widget dashboard query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLVEHICLECHALLAN).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var todaysOffenceWidgets []models.TodaysOffenceWidget
	var todaysOffenceWidget *models.TodaysOffenceWidget
	if err = cursor.All(ctx.CTX, &todaysOffenceWidgets); err != nil {
		return nil, err
	}
	if len(todaysOffenceWidgets) > 0 {
		todaysOffenceWidget = &todaysOffenceWidgets[0]
	}
	return todaysOffenceWidget, nil
}

//TopOffencesWidget : ""
func (d *Daos) TopOffencesWidget(ctx *models.Context, filter *models.TopOffencesWidgetFilter) ([]models.TopOffencesWidget, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
		if len(filter.VehicleType) > 0 {
			query = append(query, bson.M{"vehicleType": bson.M{"$in": filter.VehicleType}})
		}
		if len(filter.OffenceID) > 0 {
			query = append(query, bson.M{"uniqueId": bson.M{"$in": filter.OffenceID}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	//Calculating Top Offences
	mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{
		"from": "vehiclechallans",
		"as":   "challans",
		"let":  bson.M{"id": "$uniqueId"},
		"pipeline": []bson.M{
			bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
				bson.M{"$eq": []string{"$offenceType", "$$id"}},
			}}}},
			bson.M{"$group": bson.M{"_id": nil, "total": bson.M{"$sum": 1}, "penalty": bson.M{"$sum": "$pelalty"}}},
		},
	}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"challans": bson.M{"$arrayElemAt": []interface{}{"$challans", 0}}}})
	mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{"challans.total": -1}})

	//Aggregation
	d.Shared.BsonToJSONPrintTag("Payment widget dashboard query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLOFFENCETYPE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var topOffencesWidgets []models.TopOffencesWidget
	if err = cursor.All(ctx.CTX, &topOffencesWidgets); err != nil {
		return nil, err
	}
	return topOffencesWidgets, nil
}
