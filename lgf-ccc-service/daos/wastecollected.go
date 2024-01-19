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

//SaveWasteCollected :""
func (d *Daos) SaveWasteCollected(ctx *models.Context, WasteCollected *models.WasteCollected) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONWASTECOLLECTED).InsertOne(ctx.CTX, WasteCollected)

	//WasteCollected.ID = res.InsertedID.(primitive.ObjectID)
	return err
}

//GetSingleWasteCollected : ""
func (d *Daos) GetSingleWasteCollected(ctx *models.Context, uniqueID string) (*models.RefWasteCollected, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	// lookup
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)
	// mainPipeline = append(mainPipeline, bson.M{
	// 	"$lookup": bson.M{
	// 		"from": constants.COLLECTIONWasteCollectedLOG,
	// 		"as":   "ref.WasteCollectedlog",
	// 		"let":  bson.M{"WasteCollectedId": "$uniqueId"},
	// 		"pipeline": []bson.M{
	// 			{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
	// 				{"$eq": []string{"$status", constants.WasteCollectedASSIGNSTATUS}},
	// 				{"$eq": []string{"$WasteCollectedId", "$$WasteCollectedId"}},
	// 			}}},
	// 			},
	// 		},
	// 	},
	// })
	// mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"ref.WasteCollectedlog": bson.M{"$arrayElemAt": []interface{}{"$ref.WasteCollectedlog", 0}}}})
	// //Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONWASTECOLLECTED).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var WasteCollecteds []models.RefWasteCollected
	var WasteCollected *models.RefWasteCollected
	if err = cursor.All(ctx.CTX, &WasteCollecteds); err != nil {
		return nil, err
	}
	if len(WasteCollecteds) > 0 {
		WasteCollected = &WasteCollecteds[0]
	}
	return WasteCollected, nil
}

//UpdateWasteCollected : ""
func (d *Daos) UpdateWasteCollected(ctx *models.Context, WasteCollected *models.WasteCollected) error {
	selector := bson.M{"uniqueId": WasteCollected.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": WasteCollected}
	_, err := ctx.DB.Collection(constants.COLLECTIONWASTECOLLECTED).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// func (d *Daos) UpdateWasteCollectedAssign(ctx *models.Context, WasteCollected *models.WasteCollectedAssign) error {
// 	selector := bson.M{"uniqueId": WasteCollected.UniqueID}
// 	t := time.Now()
// 	update := models.Updated{}
// 	update.On = &t
// 	update.By = constants.SYSTEM
// 	updateInterface := bson.M{"$set": WasteCollected}
// 	_, err := ctx.DB.Collection(constants.COLLECTIONWASTECOLLECTED).UpdateOne(ctx.CTX, selector, updateInterface)
// 	if err != nil {
// 		fmt.Println("Not changed", err.Error())
// 		return err
// 	}
// 	return err
// }

//EnableWasteCollected :""
func (d *Daos) EnableWasteCollected(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.WASTECOLLECTEDSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONWASTECOLLECTED).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableWasteCollected :""
func (d *Daos) DisableWasteCollected(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.WASTECOLLECTEDSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONWASTECOLLECTED).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//EnableWasteCollected :""
func (d *Daos) WasteCollectedCompleted(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.WASTECOLLECTEDSTATUSCOMPLETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONWASTECOLLECTED).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableWasteCollected :""
func (d *Daos) WasteCollectedPending(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.WASTECOLLECTEDSTATUSPENDING}}
	_, err := ctx.DB.Collection(constants.COLLECTIONWASTECOLLECTED).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteWasteCollected :""
func (d *Daos) DeleteWasteCollected(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.WASTECOLLECTEDSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONWASTECOLLECTED).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//FilterWasteCollected : ""
func (d *Daos) FilterWasteCollected(ctx *models.Context, WasteCollectedFilter *models.FilterWasteCollected, pagination *models.Pagination) ([]models.RefWasteCollected, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if WasteCollectedFilter != nil {

		if len(WasteCollectedFilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": WasteCollectedFilter.Status}})
		}
		if len(WasteCollectedFilter.ManagerID) > 0 {
			query = append(query, bson.M{"minUser.id": bson.M{"$in": WasteCollectedFilter.ManagerID}})
		}
		if len(WasteCollectedFilter.DumbsiteID) > 0 {
			query = append(query, bson.M{"dumbSite.id": bson.M{"$in": WasteCollectedFilter.DumbsiteID}})
		}
		if len(WasteCollectedFilter.GCID) > 0 {
			query = append(query, bson.M{"gcUser.id": bson.M{"$in": WasteCollectedFilter.GCID}})
		}
		//Regex
		if WasteCollectedFilter.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: WasteCollectedFilter.Regex.Name, Options: "xi"}})
		}
		if WasteCollectedFilter.Regex.ManagerName.Name != "" {
			query = append(query, bson.M{"minUser.name": primitive.Regex{Pattern: WasteCollectedFilter.Regex.ManagerName.Name, Options: "xi"}})
		}
		if WasteCollectedFilter.Regex.GCName.Name != "" {
			query = append(query, bson.M{"gcUser.name": primitive.Regex{Pattern: WasteCollectedFilter.Regex.GCName.Name, Options: "xi"}})
		}
		if WasteCollectedFilter.Regex.CitizenName.Name != "" {
			query = append(query, bson.M{"citizen.name": primitive.Regex{Pattern: WasteCollectedFilter.Regex.CitizenName.Name, Options: "xi"}})
		}
		if WasteCollectedFilter.Regex.DumbSite.Name != "" {
			query = append(query, bson.M{"dumbSite.name": primitive.Regex{Pattern: WasteCollectedFilter.Regex.DumbSite.Name, Options: "xi"}})
		}

	}
	if WasteCollectedFilter.DateRange.From != nil {
		t := *WasteCollectedFilter.DateRange.From
		FromDate := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
		ToDate := time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, t.Location())
		if WasteCollectedFilter.DateRange.To != nil {
			t2 := *WasteCollectedFilter.DateRange.To
			ToDate = time.Date(t2.Year(), t2.Month(), t2.Day(), 23, 59, 59, 0, t2.Location())
		}
		query = append(query, bson.M{"date": bson.M{"$gte": FromDate, "$lte": ToDate}})

	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	if WasteCollectedFilter != nil {
		if WasteCollectedFilter.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{WasteCollectedFilter.SortBy: WasteCollectedFilter.SortOrder}})
		}
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONWASTECOLLECTED).CountDocuments(ctx.CTX, func() bson.M {
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
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)
	// mainPipeline = append(mainPipeline, bson.M{
	// 	"$lookup": bson.M{
	// 		"from": constants.COLLECTIONWasteCollectedLOG,
	// 		"as":   "ref.WasteCollectedlog",
	// 		"let":  bson.M{"WasteCollectedId": "$uniqueId"},
	// 		"pipeline": []bson.M{
	// 			{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
	// 				{"$eq": []string{"$status", constants.WasteCollectedASSIGNSTATUS}},
	// 				{"$eq": []string{"$WasteCollectedId", "$$WasteCollectedId"}},
	// 			}}},
	// 			},
	// 		},
	// 	},
	// })
	// mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"ref.WasteCollectedlog": bson.M{"$arrayElemAt": []interface{}{"$ref.WasteCollectedlog", 0}}}})

	//Aggregation
	d.Shared.BsonToJSONPrintTag("WasteCollected query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONWASTECOLLECTED).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var WasteCollectedsFilter []models.RefWasteCollected
	if err = cursor.All(context.TODO(), &WasteCollectedsFilter); err != nil {
		return nil, err
	}
	return WasteCollectedsFilter, nil
}

// func (d *Daos) WasteCollectedAssign(ctx *models.Context, WasteCollected *models.WasteCollectedAssign) error {
// 	selector := bson.M{"uniqueId": WasteCollected.UniqueID}
// 	t := time.Now()
// 	update := models.Updated{}
// 	update.On = &t
// 	update.By = constants.SYSTEM
// 	updateInterface := bson.M{"$set": WasteCollected}
// 	_, err := ctx.DB.Collection(constants.COLLECTIONWASTECOLLECTED).UpdateOne(ctx.CTX, selector, updateInterface)
// 	if err != nil {
// 		fmt.Println("Not changed", err.Error())
// 		return err
// 	}
// 	return err
// }

// func (d *Daos) RevokeWasteCollected(ctx *models.Context, WasteCollected *models.WasteCollected) error {
// 	selector := bson.M{"employeeId": WasteCollected.EmployeeId}
// 	t := time.Now()
// 	update := models.Updated{}
// 	update.On = &t
// 	update.By = constants.SYSTEM
// 	updateInterface := bson.M{"$set": WasteCollected, "status": constants.WASTECOLLECTEDREVOKESTATUS}
// 	_, err := ctx.DB.Collection(constants.COLLECTIONWASTECOLLECTED).UpdateOne(ctx.CTX, selector, updateInterface)
// 	if err != nil {
// 		fmt.Println("Not changed", err.Error())
// 		return err
// 	}
// 	return err
// }

// GetSingleWasteCollectedUsingEmpID : ""
func (d *Daos) GetSingleWasteCollectedUsingUniqueId(ctx *models.Context, UniqueID string) (*models.RefWasteCollected, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID, "status": bson.M{"$in": []string{constants.WASTECOLLECTEDSTATUSACTIVE}}}})
	// lookup
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONWASTECOLLECTED).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var WasteCollecteds []models.RefWasteCollected
	var WasteCollected *models.RefWasteCollected
	if err = cursor.All(ctx.CTX, &WasteCollecteds); err != nil {
		return nil, err
	}
	if len(WasteCollecteds) > 0 {
		WasteCollected = &WasteCollecteds[0]
	}
	return WasteCollected, nil
}
