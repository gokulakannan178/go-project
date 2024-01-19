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

//SaveDayOfWeek :""
func (d *Daos) SaveDayOfWeek(ctx *models.Context, dayOfWeek *models.DayOfWeek) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONDAYOFWEEK).InsertOne(ctx.CTX, dayOfWeek)
	if err != nil {
		return err
	}
	dayOfWeek.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//GetSingleDayOfWeek : ""
func (d *Daos) GetSingleDayOfWeek(ctx *models.Context, uniqueID string) (*models.RefDayOfWeek, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	// lookup
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDayOfWeekCATEGORY, "categoryId", "uniqueId", "ref.category", "ref.category")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDAYOFWEEK).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var dayOfWeeks []models.RefDayOfWeek
	var dayOfWeek *models.RefDayOfWeek
	if err = cursor.All(ctx.CTX, &dayOfWeeks); err != nil {
		return nil, err
	}
	if len(dayOfWeeks) > 0 {
		dayOfWeek = &dayOfWeeks[0]
	}
	return dayOfWeek, nil
}

//UpdateDayOfWeek : ""
func (d *Daos) UpdateDayOfWeek(ctx *models.Context, dayOfWeek *models.DayOfWeek) error {
	selector := bson.M{"uniqueId": dayOfWeek.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": dayOfWeek}
	_, err := ctx.DB.Collection(constants.COLLECTIONDAYOFWEEK).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//EnableDayOfWeek :""
func (d *Daos) EnableDayOfWeek(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.DAYOFWEEKSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONDAYOFWEEK).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableDayOfWeek :""
func (d *Daos) DisableDayOfWeek(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.DAYOFWEEKSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONDAYOFWEEK).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteDayOfWeek :""
func (d *Daos) DeleteDayOfWeek(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.DAYOFWEEKSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONDAYOFWEEK).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//FilterDayOfWeek : ""
func (d *Daos) FilterDayOfWeek(ctx *models.Context, filter *models.FilterDayOfWeek, pagination *models.Pagination) ([]models.RefDayOfWeek, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {

		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
		//Regex
		if filter.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: filter.Regex.Name, Options: "xi"}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	if filter != nil {
		if filter.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{filter.SortBy: filter.SortOrder}})
		}
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONDAYOFWEEK).CountDocuments(ctx.CTX, func() bson.M {
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
	// lookup
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDayOfWeekCATEGORY, "categoryId", "uniqueId", "ref.category", "ref.category")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("DayOfWeek query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDAYOFWEEK).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var dayOfWeeks []models.RefDayOfWeek
	if err = cursor.All(context.TODO(), &dayOfWeeks); err != nil {
		return nil, err
	}
	return dayOfWeeks, nil
}
func (d *Daos) GetSingleDayOfWeekWithDays(ctx *models.Context, uniqueID int64) (*models.RefDayOfWeek, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"dayOfWeek": uniqueID}})
	// lookup
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDayOfWeekCATEGORY, "categoryId", "uniqueId", "ref.category", "ref.category")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDAYOFWEEK).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var dayOfWeeks []models.RefDayOfWeek
	var dayOfWeek *models.RefDayOfWeek
	if err = cursor.All(ctx.CTX, &dayOfWeeks); err != nil {
		return nil, err
	}
	if len(dayOfWeeks) > 0 {
		dayOfWeek = &dayOfWeeks[0]
	}
	return dayOfWeek, nil
}
