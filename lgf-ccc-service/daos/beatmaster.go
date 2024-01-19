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

// SaveBeatMaster : ""
func (d *Daos) SaveBeatMaster(ctx *models.Context, BeatMaster *models.BeatMaster) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONBEATMASTER).InsertOne(ctx.CTX, BeatMaster)
	if err != nil {
		return err
	}
	BeatMaster.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

// UpdateBeatMaster : ""
func (d *Daos) UpdateBeatMaster(ctx *models.Context, BeatMaster *models.BeatMaster) error {
	selector := bson.M{"uniqueId": BeatMaster.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": BeatMaster}
	_, err := ctx.DB.Collection(constants.COLLECTIONBEATMASTER).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// GetSingleBeatMaster : ""
func (d *Daos) GetSingleBeatMaster(ctx *models.Context, uniqueID string) (*models.RefBeatMaster, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	//LookUp
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONBEATMASTER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var BeatMasters []models.RefBeatMaster
	var BeatMaster *models.RefBeatMaster
	if err = cursor.All(ctx.CTX, &BeatMasters); err != nil {
		return nil, err
	}
	if len(BeatMasters) > 0 {
		BeatMaster = &BeatMasters[0]
	}
	return BeatMaster, err
}

// EnableBeatMaster : ""
func (d *Daos) EnableBeatMaster(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.BEATMASTERSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONBEATMASTER).UpdateOne(ctx.CTX, selector, data)
	return err
}

// DisableBeatMaster : ""
func (d *Daos) DisableBeatMaster(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.BEATMASTERSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONBEATMASTER).UpdateOne(ctx.CTX, selector, data)
	return err
}

// DeleteState :""
func (d *Daos) DeleteBeatMaster(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.BEATMASTERSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONBEATMASTER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterBeatMaster : ""
func (d *Daos) FilterBeatMaster(ctx *models.Context, filter *models.FilterBeatMaster, pagination *models.Pagination) ([]models.RefBeatMaster, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONROUTEMASTER, "routeId", "uniqueId", "ref.route", "ref.route")...)

	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
		if len(filter.EmployeeId) > 0 {
			query = append(query, bson.M{"organisationId": bson.M{"$in": filter.EmployeeId}})
		}
		if len(filter.RouteID) > 0 {
			query = append(query, bson.M{"routeId": bson.M{"$in": filter.RouteID}})
		}
		if len(filter.WardCode) > 0 {
			query = append(query, bson.M{"ref.route.area.wardCode": bson.M{"$in": filter.WardCode}})
		}
		if len(filter.ZoneCode) > 0 {
			query = append(query, bson.M{"ref.route.area.zoneCone": bson.M{"$in": filter.ZoneCode}})
		}
		if len(filter.VehicleId) > 0 {
			query = append(query, bson.M{"vehicle.id": bson.M{"$in": filter.VehicleId}})
		}
		if len(filter.DriverId) > 0 {
			query = append(query, bson.M{"driver.id": bson.M{"$in": filter.DriverId}})
		}
		//Regex
		if filter.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: filter.Regex.Name, Options: "xi"}})
		}
		if filter.Regex.ContactNo != "" {
			query = append(query, bson.M{"mobile": primitive.Regex{Pattern: filter.Regex.ContactNo, Options: "xi"}})
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
	// if pagination != nil {
	// 	mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
	// 	//Getting Total count
	// 	totalCount, err := ctx.DB.Collection(constants.COLLECTIONBEATMASTER).CountDocuments(ctx.CTX, func() bson.M {
	// 		if query != nil {
	// 			if len(query) > 0 {
	// 				return bson.M{"$and": query}
	// 			}
	// 		}
	// 		return bson.M{}
	// 	}())
	// 	if err != nil {
	// 		log.Println("Error in getting pagination")
	// 	}
	// 	fmt.Println("count", totalCount)
	// 	pagination.Count = int(totalCount)
	// 	d.Shared.PaginationData(pagination)
	// }
	if pagination != nil {
		paginationPipeline := []bson.M{}
		paginationPipeline = append(paginationPipeline, mainPipeline...)
		paginationPipeline = append(paginationPipeline, bson.M{"$group": bson.M{"_id": nil, "count": bson.M{"$sum": 1}}})
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		d.Shared.BsonToJSONPrintTag("Commodity Pagination quary =>", paginationPipeline)

		countcursor, err := ctx.DB.Collection(constants.COLLECTIONBEATMASTER).Aggregate(ctx.CTX, paginationPipeline, nil)
		if err != nil {
			log.Println("Error in geting pagination - " + err.Error())
		}
		countstruct := []models.Countstruct{}
		err = countcursor.All(ctx.CTX, &countstruct)
		if err != nil {
			return nil, err
		}
		var totalCount int64
		if len(countstruct) > 0 {
			totalCount = countstruct[0].Count
		}
		fmt.Println("count", totalCount)
		pagination.Count = int(totalCount)
		d.Shared.PaginationData(pagination)
	}
	//LookUp
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONROUTEMASTER, "routeId", "uniqueId", "ref.route", "ref.route")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("Feature query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONBEATMASTER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var BeatMasterFilter []models.RefBeatMaster
	if err = cursor.All(context.TODO(), &BeatMasterFilter); err != nil {
		return nil, err
	}
	return BeatMasterFilter, nil
}

func (d *Daos) VehicleAssignForBeatMaster(ctx *models.Context, BeatMaster *models.BeatMaster) error {
	selector := bson.M{"uniqueId": BeatMaster.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": bson.M{"driver": BeatMaster.Driver, "vehicle": BeatMaster.Vehicle}}
	_, err := ctx.DB.Collection(constants.COLLECTIONBEATMASTER).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}
