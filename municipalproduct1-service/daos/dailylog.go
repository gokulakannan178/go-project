package daos

import (
	"errors"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// SaveLease : ""
func (d *Daos) InitiateDailylog(ctx *models.Context, dailyLog *models.DailyLog) error {
	opts := options.Update().SetUpsert(true)
	updateQuery := bson.M{"dateStr": dailyLog.Datestr}
	updateData := bson.M{"$set": dailyLog}
	if _, err := ctx.DB.Collection(constants.COLLECTIONDAILYLOG).UpdateOne(ctx.CTX, updateQuery, updateData, opts); err != nil {
		return errors.New("Error in upserting - " + err.Error())
	}
	return nil
}

func (d *Daos) DailylogGetTodaysCompletedPayments(ctx *models.Context, dailyLog *models.DailyLog) error {
	sd := time.Date(dailyLog.Date.Year(), dailyLog.Date.Month(), dailyLog.Date.Day(), 0, 0, 0, 0, dailyLog.Date.Location())
	ed := time.Date(dailyLog.Date.Year(), dailyLog.Date.Month(), dailyLog.Date.Day(), 23, 59, 59, 0, dailyLog.Date.Location())
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{
		"status":         constants.PROPERTYPAYMENTCOMPLETED,
		"completionDate": bson.M{"$gte": sd, "$lte": ed},
	}})

	mainPipeline = append(mainPipeline, bson.M{"$group": bson.M{
		"_id":       nil,
		"ptcurrent": bson.M{"$sum": "$demand.current"},
		"ptarrear":  bson.M{"$sum": "$demand.arrear"},
		"pttotal":   bson.M{"$sum": "$demand.totalTax"},
	}})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{
		"propertyTax.collections.current.total": "$ptcurrent",
		"propertyTax.collections.arrear.total":  "$ptarrear",
		"propertyTax.collections.total.total":   "$pttotal",
	}})

	//Aggregation
	d.Shared.BsonToJSONPrintTag("payment query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROPERTYPAYMENT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return err
	}

	var dls []models.DailyLog
	var dl *models.DailyLog
	if err = cursor.All(ctx.CTX, &dls); err != nil {
		return err
	}
	if len(dls) > 0 {
		dl = &dls[0]
	}
	if dl == nil {
		dl = new(models.DailyLog)
	}
	dailyLog.PropertyTax.Collections = dl.PropertyTax.Collections
	opts := options.Update().SetUpsert(true)
	updateQuery := bson.M{"dateStr": dailyLog.Datestr}
	updateData := bson.M{"$set": bson.M{"propertyTax.collections.todaysCompleted": dl.PropertyTax.Collections.TodaysCompleted}}
	if _, err := ctx.DB.Collection(constants.COLLECTIONDAILYLOG).UpdateOne(ctx.CTX, updateQuery, updateData, opts); err != nil {
		return errors.New("Error in upserting - " + err.Error())
	}
	return nil

}
