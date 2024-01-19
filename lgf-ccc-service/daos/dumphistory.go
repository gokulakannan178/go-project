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

//SaveDumpHistory :""
func (d *Daos) SaveDumpHistory(ctx *models.Context, dumpHistory *models.DumpHistory) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONDUMPHISTORY).InsertOne(ctx.CTX, dumpHistory)

	//DumpHistory.ID = res.InsertedID.(primitive.ObjectID)
	return err
}

//GetSingleDumpHistory : ""
func (d *Daos) GetSingleDumpHistory(ctx *models.Context, uniqueID string) (*models.RefDumpHistory, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	// lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDUMPSITE, "name", "uniqueId", "ref.dumbsite", "ref.dumbsite")...)
	// //Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDUMPHISTORY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var DumpHistorys []models.RefDumpHistory
	var DumpHistory *models.RefDumpHistory
	if err = cursor.All(ctx.CTX, &DumpHistorys); err != nil {
		return nil, err
	}
	if len(DumpHistorys) > 0 {
		DumpHistory = &DumpHistorys[0]
	}
	return DumpHistory, nil
}

//UpdateDumpHistory : ""
func (d *Daos) UpdateDumpHistory(ctx *models.Context, DumpHistory *models.DumpHistory) error {
	selector := bson.M{"uniqueId": DumpHistory.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": DumpHistory}
	_, err := ctx.DB.Collection(constants.COLLECTIONDUMPHISTORY).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// func (d *Daos) UpdateDumpHistoryAssign(ctx *models.Context, DumpHistory *models.DumpHistoryAssign) error {
// 	selector := bson.M{"uniqueId": DumpHistory.UniqueID}
// 	t := time.Now()
// 	update := models.Updated{}
// 	update.On = &t
// 	update.By = constants.SYSTEM
// 	updateInterface := bson.M{"$set": DumpHistory}
// 	_, err := ctx.DB.Collection(constants.COLLECTIONDumpHistory).UpdateOne(ctx.CTX, selector, updateInterface)
// 	if err != nil {
// 		fmt.Println("Not changed", err.Error())
// 		return err
// 	}
// 	return err
// }

//EnableDumpHistory :""
func (d *Daos) EnableDumpHistory(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.DUMPHISTORYSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONDUMPHISTORY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableDumpHistory :""
func (d *Daos) DisableDumpHistory(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.DUMPHISTORYSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONDUMPHISTORY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteDumpHistory :""
func (d *Daos) DeleteDumpHistory(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.DUMPHISTORYSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONDUMPHISTORY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//FilterDumpHistory : ""
func (d *Daos) FilterDumpHistory(ctx *models.Context, DumpHistoryFilter *models.FilterDumpHistory, pagination *models.Pagination) ([]models.RefDumpHistory, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if DumpHistoryFilter != nil {

		if len(DumpHistoryFilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": DumpHistoryFilter.Status}})
		}
		if len(DumpHistoryFilter.ManagerId) > 0 {
			query = append(query, bson.M{"minUser.id": bson.M{"$in": DumpHistoryFilter.ManagerId}})
		}
		if len(DumpHistoryFilter.DumbUser) > 0 {
			query = append(query, bson.M{"dumbUser.id": bson.M{"$in": DumpHistoryFilter.DumbUser}})
		}
		if len(DumpHistoryFilter.Name) > 0 {
			query = append(query, bson.M{"name": bson.M{"$in": DumpHistoryFilter.Name}})
		}
		//Regex
		if DumpHistoryFilter.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: DumpHistoryFilter.Regex.Name, Options: "xi"}})
		}
		if DumpHistoryFilter.Regex.DriverName != "" {
			query = append(query, bson.M{"driver.name": primitive.Regex{Pattern: DumpHistoryFilter.Regex.DriverName, Options: "xi"}})
		}
		if DumpHistoryFilter.Regex.VehicleName != "" {
			query = append(query, bson.M{"vehicleId.name": primitive.Regex{Pattern: DumpHistoryFilter.Regex.VehicleName, Options: "xi"}})
		}

	}
	if DumpHistoryFilter.DateRange.From != nil {
		t := *DumpHistoryFilter.DateRange.From
		FromDate := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
		ToDate := time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, t.Location())
		if DumpHistoryFilter.DateRange.To != nil {
			t2 := *DumpHistoryFilter.DateRange.To
			ToDate = time.Date(t2.Year(), t2.Month(), t2.Day(), 23, 59, 59, 0, t2.Location())
		}
		query = append(query, bson.M{"date": bson.M{"$gte": FromDate, "$lte": ToDate}})

	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	if DumpHistoryFilter != nil {
		if DumpHistoryFilter.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{DumpHistoryFilter.SortBy: DumpHistoryFilter.SortOrder}})
		}
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONDUMPHISTORY).CountDocuments(ctx.CTX, func() bson.M {
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
	// lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDUMPSITE, "name", "uniqueId", "ref.dumbsite", "ref.dumbsite")...)
	//Aggregation
	d.Shared.BsonToJSONPrintTag("DumpHistory query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDUMPHISTORY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var DumpHistorysFilter []models.RefDumpHistory
	if err = cursor.All(context.TODO(), &DumpHistorysFilter); err != nil {
		return nil, err
	}
	return DumpHistorysFilter, nil
}
func (d *Daos) GetQuantityByManagerId(ctx *models.Context, DumpHistoryFilter *models.FilterDumpHistory) ([]models.GetQuantity, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if DumpHistoryFilter != nil {

		if len(DumpHistoryFilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": DumpHistoryFilter.Status}})
		}
		if len(DumpHistoryFilter.ManagerId) > 0 {
			query = append(query, bson.M{"minUser.id": bson.M{"$in": DumpHistoryFilter.ManagerId}})
		}
		if len(DumpHistoryFilter.DumbUser) > 0 {
			query = append(query, bson.M{"dumbUser.id": bson.M{"$in": DumpHistoryFilter.DumbUser}})
		}
		//Regex
		if DumpHistoryFilter.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: DumpHistoryFilter.Regex.Name, Options: "xi"}})
		}
		if DumpHistoryFilter.Regex.DriverName != "" {
			query = append(query, bson.M{"driver.name": primitive.Regex{Pattern: DumpHistoryFilter.Regex.DriverName, Options: "xi"}})
		}
		if DumpHistoryFilter.Regex.VehicleName != "" {
			query = append(query, bson.M{"vehicle.name": primitive.Regex{Pattern: DumpHistoryFilter.Regex.VehicleName, Options: "xi"}})
		}

	}
	if DumpHistoryFilter.DateRange.From != nil {
		t := *DumpHistoryFilter.DateRange.From
		FromDate := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
		ToDate := time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, t.Location())
		if DumpHistoryFilter.DateRange.To != nil {
			t2 := *DumpHistoryFilter.DateRange.To
			ToDate = time.Date(t2.Year(), t2.Month(), t2.Day(), 23, 59, 59, 0, t2.Location())
		}
		query = append(query, bson.M{"date": bson.M{"$gte": FromDate, "$lte": ToDate}})

	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if DumpHistoryFilter.Status != nil {
		mainPipeline = append(mainPipeline, bson.M{"$group": bson.M{
			"_id":      nil,
			"quantity": bson.M{"$sum": "$quantity"},
		}})
	}
	if DumpHistoryFilter.ManagerId != nil {
		mainPipeline = append(mainPipeline, bson.M{"$group": bson.M{
			"_id":      "$minUser.id",
			"quantity": bson.M{"$sum": "$quantity"},
		}})
	}
	if DumpHistoryFilter.DumbUser != nil {
		mainPipeline = append(mainPipeline, bson.M{"$group": bson.M{
			"_id":      "$dumbUser.id",
			"quantity": bson.M{"$sum": "$quantity"},
		}})
	}
	if DumpHistoryFilter.DateRange.From != nil {
		mainPipeline = append(mainPipeline, bson.M{"$group": bson.M{
			"_id":      nil,
			"quantity": bson.M{"$sum": "$quantity"},
		}})
	}

	//Aggregation
	d.Shared.BsonToJSONPrintTag("DumpHistory query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDUMPHISTORY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var GetQuantity []models.GetQuantity
	if err = cursor.All(context.TODO(), &GetQuantity); err != nil {
		return nil, err
	}
	return GetQuantity, nil
}

// func (d *Daos) DumpHistoryAssign(ctx *models.Context, DumpHistory *models.DumpHistoryAssign) error {
// 	selector := bson.M{"uniqueId": DumpHistory.UniqueID}
// 	t := time.Now()
// 	update := models.Updated{}
// 	update.On = &t
// 	update.By = constants.SYSTEM
// 	updateInterface := bson.M{"$set": DumpHistory}
// 	_, err := ctx.DB.Collection(constants.COLLECTIONDumpHistory).UpdateOne(ctx.CTX, selector, updateInterface)
// 	if err != nil {
// 		fmt.Println("Not changed", err.Error())
// 		return err
// 	}
// 	return err
// }

// func (d *Daos) RevokeDumpHistory(ctx *models.Context, DumpHistory *models.DumpHistory) error {
// 	selector := bson.M{"employeeId": DumpHistory.EmployeeId}
// 	t := time.Now()
// 	update := models.Updated{}
// 	update.On = &t
// 	update.By = constants.SYSTEM
// 	updateInterface := bson.M{"$set": DumpHistory, "status": constants.DumpHistoryREVOKESTATUS}
// 	_, err := ctx.DB.Collection(constants.COLLECTIONDumpHistory).UpdateOne(ctx.CTX, selector, updateInterface)
// 	if err != nil {
// 		fmt.Println("Not changed", err.Error())
// 		return err
// 	}
// 	return err
// }

// GetSingleDumpHistoryUsingEmpID : ""
func (d *Daos) GetSingleDumpHistoryUsingUniqueId(ctx *models.Context, UniqueID string) (*models.RefDumpHistory, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID, "status": bson.M{"$in": []string{constants.DUMPHISTORYSTATUSACTIVE}}}})
	// lookup
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDUMPHISTORY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var DumpHistorys []models.RefDumpHistory
	var DumpHistory *models.RefDumpHistory
	if err = cursor.All(ctx.CTX, &DumpHistorys); err != nil {
		return nil, err
	}
	if len(DumpHistorys) > 0 {
		DumpHistory = &DumpHistorys[0]
	}
	return DumpHistory, nil
}

func (d *Daos) DateWiseDumpHistory(ctx *models.Context, filter *models.DayWiseDumpHistory) ([]models.GetQuantity, error) {
	mainPipeline := []bson.M{}
	// query := []bson.M{}

	var daysd, dayed time.Time
	if filter != nil {
		daysd = time.Date(filter.Date.Year(), filter.Date.Month(), filter.Date.Day(), 0, 0, 0, 0, daysd.Location())
		dayed = time.Date(filter.Date.Year(), filter.Date.Month(), filter.Date.Day(), 23, 59, 59, 0, dayed.Location())
	}
	//Adding $match from filter
	// if len(query) > 0 {
	// 	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	// }
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": []bson.M{
		bson.M{"status": bson.M{"$in": []string{constants.DUMPHISTORYSTATUSACTIVE}}},
		bson.M{"date": bson.M{"$gte": daysd,
			"$lte": dayed}},
	}}},
		bson.M{"$group": bson.M{"_id": nil,
			"quantity": bson.M{"$sum": 1},
		},
		},
		//bson.M{"$addFields": bson.M{"date": "$_id"}},
		bson.M{"$project": bson.M{"_id": 0}},
	)
	// ============================================>

	//Aggregation
	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDUMPHISTORY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var report []models.GetQuantity
	//var data []models.DateWisePropertyPaymentReport
	if err = cursor.All(context.TODO(), &report); err != nil {
		return nil, err
	}

	return report, nil
}
