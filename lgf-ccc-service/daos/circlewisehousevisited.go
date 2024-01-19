package daos

import (
	"context"
	"errors"
	"fmt"
	"lgf-ccc-service/constants"
	"lgf-ccc-service/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// SaveCircleWiseHouseVisited : ""
func (d *Daos) SaveCircleWiseHouseVisited(ctx *models.Context, circlewisehousevisited *models.CircleWiseHouseVisited) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONCIRCLEWISEHOUSEVISITED).InsertOne(ctx.CTX, circlewisehousevisited)
	if err != nil {
		return err
	}
	circlewisehousevisited.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//UpdateState : ""
func (d *Daos) UpdateCircleWiseHouseVisited(ctx *models.Context, dept *models.CircleWiseHouseVisited) error {
	selector := bson.M{"uniqueId": dept.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": dept, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONCIRCLEWISEHOUSEVISITED).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// GetSingleCircleWiseHouseVisited : ""
func (d *Daos) GetSingleCircleWiseHouseVisited(ctx *models.Context, uniqueID string) (*models.RefCircleWiseHouseVisited, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	//LookUp
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONCIRCLEWISEHOUSEVISITED).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var circlewisehousevisiteds []models.RefCircleWiseHouseVisited
	var circlewisehousevisited *models.RefCircleWiseHouseVisited
	if err = cursor.All(ctx.CTX, &circlewisehousevisiteds); err != nil {
		return nil, err
	}
	if len(circlewisehousevisiteds) > 0 {
		circlewisehousevisited = &circlewisehousevisiteds[0]
	}
	return circlewisehousevisited, err
}

// EnableCircleWiseHouseVisited : ""
func (d *Daos) EnableCircleWiseHouseVisited(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.CIRCLEWISEHOUSEVISITEDSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONCIRCLEWISEHOUSEVISITED).UpdateOne(ctx.CTX, selector, data)
	return err
}

// DisableCircleWiseHouseVisited : ""
func (d *Daos) DisableCircleWiseHouseVisited(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.CIRCLEWISEHOUSEVISITEDSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONCIRCLEWISEHOUSEVISITED).UpdateOne(ctx.CTX, selector, data)
	return err
}

//DeleteCircleWiseHouseVisited :""
func (d *Daos) DeleteCircleWiseHouseVisited(ctx *models.Context, uniqueID string) error {
	query := bson.M{"uniqueId": uniqueID}
	update := bson.M{"$set": bson.M{"status": constants.CIRCLEWISEHOUSEVISITEDSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONCIRCLEWISEHOUSEVISITED).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterCircleWiseHouseVisited : ""
func (d *Daos) FilterCircleWiseHouseVisited(ctx *models.Context, filter *models.FilterCircleWiseHouseVisited, pagination *models.Pagination) ([]models.RefCircleWiseHouseVisited, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
		if len(filter.CircleCode) > 0 {
			query = append(query, bson.M{"circleCode": bson.M{"$in": filter.CircleCode}})
		}
		if len(filter.Datestr) > 0 {
			query = append(query, bson.M{"datestr": bson.M{"$in": filter.Datestr}})
		}
		//Regex
		if filter.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: filter.Regex.Name, Options: "xi"}})
		}

		if filter.Regex.Type != "" {
			query = append(query, bson.M{"userName": primitive.Regex{Pattern: filter.Regex.Type, Options: "xi"}})
		}
	}
	if filter.DateRange.From != nil {
		t := *filter.DateRange.From
		FromDate := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
		ToDate := time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, t.Location())
		if filter.DateRange.To != nil {
			t2 := *filter.DateRange.To
			ToDate = time.Date(t2.Year(), t2.Month(), t2.Day(), 23, 59, 59, 0, t2.Location())
		}
		query = append(query, bson.M{"date": bson.M{"$gte": FromDate, "$lte": ToDate}})

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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONCIRCLEWISEHOUSEVISITED).CountDocuments(ctx.CTX, func() bson.M {
			if query != nil {
				if len(query) > 0 {
					return bson.M{"$and": query}
				}
			}
			return bson.M{}
		}())
		if err != nil {
			log.Println("Error in getting pagination")
		}
		fmt.Println("count", totalCount)
		pagination.Count = int(totalCount)
		d.Shared.PaginationData(pagination)
	}
	//LookUp
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("Feature query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONCIRCLEWISEHOUSEVISITED).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var circlewisehousevisited []models.RefCircleWiseHouseVisited
	if err = cursor.All(context.TODO(), &circlewisehousevisited); err != nil {
		return nil, err
	}
	return circlewisehousevisited, nil
}

func (d *Daos) GetSingleCircleWiseHouseVisitedWithDate(ctx *models.Context, date string, circleCode string) (*models.RefCircleWiseHouseVisited, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"dateStr": date, "circleCode": circleCode}})

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONCIRCLEWISEHOUSEVISITED).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var WardWiseHouseVisiteds []models.RefCircleWiseHouseVisited
	var WardWiseHouseVisited *models.RefCircleWiseHouseVisited
	if err = cursor.All(ctx.CTX, &WardWiseHouseVisiteds); err != nil {
		return nil, err
	}
	if len(WardWiseHouseVisiteds) > 0 {
		WardWiseHouseVisited = &WardWiseHouseVisiteds[0]
	}
	return WardWiseHouseVisited, err
}

func (d *Daos) IncreaseCircleWiseHouseVisitedCount(ctx *models.Context, date string, circleCode string) error {
	selector := bson.M{"dateStr": date, "circleCode": circleCode}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$inc": bson.M{"todayCollection": 1}}
	_, err := ctx.DB.Collection(constants.COLLECTIONCIRCLEWISEHOUSEVISITED).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}
