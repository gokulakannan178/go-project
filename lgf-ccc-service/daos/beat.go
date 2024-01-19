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

// SaveBeat : ""
func (d *Daos) SaveBeat(ctx *models.Context, beat *models.Beat) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONBEAT).InsertOne(ctx.CTX, beat)
	if err != nil {
		return err
	}
	beat.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

// UpdateBeat : ""
func (d *Daos) UpdateBeat(ctx *models.Context, beat *models.Beat) error {
	selector := bson.M{"uniqueId": beat.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": beat, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONBEAT).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// GetSingleBeat : ""
func (d *Daos) GetSingleBeat(ctx *models.Context, uniqueID string) (*models.RefBeat, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	//LookUp
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONBEAT, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONBEATMASTER, "beatMasterId", "uniqueId", "ref.beatmaster", "ref.beatmaster")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONROUTEMASTER, "beatMasterId.routeId", "uniqueId", "ref.route", "ref.route")...)
	mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONVEHICLELOCATION, "uniqueId", "uniqueId", "ref.vehicleLocation", "ref.vehicleLocation")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONBEAT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var beats []models.RefBeat
	var beat *models.RefBeat
	if err = cursor.All(ctx.CTX, &beats); err != nil {
		return nil, err
	}
	if len(beats) > 0 {
		beat = &beats[0]
	}
	return beat, err
}

// EnableBeat : ""
func (d *Daos) EnableBeat(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.BEATSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONBEAT).UpdateOne(ctx.CTX, selector, data)
	return err
}

// DisableBeat : ""
func (d *Daos) DisableBeat(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.BEATSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONBEAT).UpdateOne(ctx.CTX, selector, data)
	return err
}

// DeleteBeat :""
func (d *Daos) DeleteBeat(ctx *models.Context, uniqueID string) error {
	query := bson.M{"uniqueId": uniqueID}
	update := bson.M{"$set": bson.M{"status": constants.BEATSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONBEAT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterBeat : ""
func (d *Daos) FilterBeat(ctx *models.Context, filter *models.FilterBeat, pagination *models.Pagination) ([]models.RefBeat, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
		if len(filter.VehicleId) > 0 {
			query = append(query, bson.M{"assignBeatdetails.vehicle.id": bson.M{"$in": filter.VehicleId}})
		}
		// if len(filter.BranchId) > 0 {
		// 	query = append(query, bson.M{"branch": bson.M{"$in": filter.BranchId}})
		// }
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONBEAT).CountDocuments(ctx.CTX, func() bson.M {
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
	// if pagination != nil {
	// 	paginationPipeline := []bson.M{}
	// 	paginationPipeline = append(paginationPipeline, mainPipeline...)
	// 	paginationPipeline = append(paginationPipeline, bson.M{"$group": bson.M{"_id": nil, "count": bson.M{"$sum": 1}}})
	// 	mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
	// 	//Getting Total count
	// 	d.Shared.BsonToJSONPrintTag("Commodity Pagination quary =>", paginationPipeline)

	// 	countcursor, err := ctx.DB.Collection(constants.COLLECTIONBEAT).Aggregate(ctx.CTX, paginationPipeline, nil)
	// 	if err != nil {
	// 		log.Println("Error in geting pagination - " + err.Error())
	// 	}
	// 	countstruct := []models.Countstruct{}
	// 	err = countcursor.All(ctx.CTX, &countstruct)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	var totalCount int64
	// 	if len(countstruct) > 0 {
	// 		totalCount = countstruct[0].Count
	// 	}
	// 	fmt.Println("count", totalCount)
	// 	pagination.Count = int(totalCount)
	// 	d.Shared.PaginationData(pagination)
	// }
	//LookUp
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONBEATMASTER, "beatMasterId", "uniqueId", "ref.beatmaster", "ref.beatmaster")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONROUTEMASTER, "beatMasterId.routeId", "uniqueId", "ref.route", "ref.route")...)
	mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONVEHICLELOCATION, "uniqueId", "uniqueId", "ref.vehicleLocation", "ref.vehicleLocation")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("Feature query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONBEAT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var beatFilter []models.RefBeat
	if err = cursor.All(context.TODO(), &beatFilter); err != nil {
		return nil, err
	}
	return beatFilter, nil
}

func (d *Daos) EndBeat(ctx *models.Context, uniqueID string) error {
	t := time.Now()
	query := bson.M{"uniqueId": uniqueID}
	update := bson.M{"$set": bson.M{"status": constants.BEATSTATUSCOMPLETED, "endTime": &t}}
	_, err := ctx.DB.Collection(constants.COLLECTIONBEAT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
