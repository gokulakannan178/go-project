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
	"go.mongodb.org/mongo-driver/mongo/options"
)

// SaveWardWiseHouseVisited : ""
func (d *Daos) SaveWardWiseHouseVisited(ctx *models.Context, WardWiseHouseVisited *models.WardWiseHouseVisited) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONWARDWISEHOUSEVISITED).InsertOne(ctx.CTX, WardWiseHouseVisited)
	if err != nil {
		return err
	}
	WardWiseHouseVisited.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}
func (d *Daos) SaveWardWiseHouseVisitedWithUpsert(ctx *models.Context, WardWiseHouseVisited *models.WardWiseHouseVisited) error {
	opts := options.Update().SetUpsert(true)
	updateQuery := bson.M{"dateStr": WardWiseHouseVisited.Datestr, "wardCode": WardWiseHouseVisited.WardCode}
	updateData := bson.M{"$set": WardWiseHouseVisited, "quantity": bson.M{"$inc": 1}}
	if _, err := ctx.DB.Collection(constants.COLLECTIONWARDWISEHOUSEVISITED).UpdateOne(ctx.CTX, updateQuery, updateData, opts); err != nil {
		return errors.New("Error in updating log - " + err.Error())
	}

	return nil
}

//UpdateState : ""
func (d *Daos) UpdateWardWiseHouseVisited(ctx *models.Context, dept *models.WardWiseHouseVisited) error {
	selector := bson.M{"uniqueId": dept.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": dept, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONWARDWISEHOUSEVISITED).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// GetSingleWardWiseHouseVisited : ""
func (d *Daos) GetSingleWardWiseHouseVisited(ctx *models.Context, uniqueID string) (*models.RefWardWiseHouseVisited, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	//LookUp
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONWARDWISEHOUSEVISITED).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var WardWiseHouseVisiteds []models.RefWardWiseHouseVisited
	var WardWiseHouseVisited *models.RefWardWiseHouseVisited
	if err = cursor.All(ctx.CTX, &WardWiseHouseVisiteds); err != nil {
		return nil, err
	}
	if len(WardWiseHouseVisiteds) > 0 {
		WardWiseHouseVisited = &WardWiseHouseVisiteds[0]
	}
	return WardWiseHouseVisited, err
}

// EnableWardWiseHouseVisited : ""
func (d *Daos) EnableWardWiseHouseVisited(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.WARDWISEHOUSEVISITEDSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONWARDWISEHOUSEVISITED).UpdateOne(ctx.CTX, selector, data)
	return err
}

// DisableWardWiseHouseVisited : ""
func (d *Daos) DisableWardWiseHouseVisited(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.WARDWISEHOUSEVISITEDSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONWARDWISEHOUSEVISITED).UpdateOne(ctx.CTX, selector, data)
	return err
}

//DeleteWardWiseHouseVisited :""
func (d *Daos) DeleteWardWiseHouseVisited(ctx *models.Context, uniqueID string) error {
	query := bson.M{"uniqueId": uniqueID}
	update := bson.M{"$set": bson.M{"status": constants.WARDWISEHOUSEVISITEDSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONWARDWISEHOUSEVISITED).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterWardWiseHouseVisited : ""
func (d *Daos) FilterWardWiseHouseVisited(ctx *models.Context, filter *models.FilterWardWiseHouseVisited, pagination *models.Pagination) ([]models.RefWardWiseHouseVisited, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
		if len(filter.CircleCode) > 0 {
			query = append(query, bson.M{"circleCode": bson.M{"$in": filter.CircleCode}})
		}
		if len(filter.WardCode) > 0 {
			query = append(query, bson.M{"wardCode": bson.M{"$in": filter.WardCode}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONWARDWISEHOUSEVISITED).CountDocuments(ctx.CTX, func() bson.M {
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
	cursor, err := ctx.DB.Collection(constants.COLLECTIONWARDWISEHOUSEVISITED).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var WardWiseHouseVisitedFilter []models.RefWardWiseHouseVisited
	if err = cursor.All(context.TODO(), &WardWiseHouseVisitedFilter); err != nil {
		return nil, err
	}
	return WardWiseHouseVisitedFilter, nil
}

func (d *Daos) GetSingleWardWiseHouseVisitedWithDate(ctx *models.Context, date string, wardCode string) (*models.RefWardWiseHouseVisited, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"dateStr": date, "wardCode": wardCode}})

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONWARDWISEHOUSEVISITED).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var WardWiseHouseVisiteds []models.RefWardWiseHouseVisited
	var WardWiseHouseVisited *models.RefWardWiseHouseVisited
	if err = cursor.All(ctx.CTX, &WardWiseHouseVisiteds); err != nil {
		return nil, err
	}
	if len(WardWiseHouseVisiteds) > 0 {
		WardWiseHouseVisited = &WardWiseHouseVisiteds[0]
	}
	return WardWiseHouseVisited, err
}

func (d *Daos) IncreaseWardWiseHouseVisitedCount(ctx *models.Context, date string, wardCode string) error {
	selector := bson.M{"dateStr": date, "wardCode": wardCode}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$inc": bson.M{"todayCollection": 1}}
	_, err := ctx.DB.Collection(constants.COLLECTIONWARDWISEHOUSEVISITED).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}
