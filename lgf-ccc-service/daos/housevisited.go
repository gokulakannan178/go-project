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

// HouseVisited : ""
func (d *Daos) SaveHouseVisited(ctx *models.Context, housevisited *models.HouseVisited) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONHOUSEVISITED).InsertOne(ctx.CTX, housevisited)
	if err != nil {
		return err
	}
	//housevisited.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

// GetSingleHouseVisited : ""
func (d *Daos) GetSingleHouseVisited(ctx *models.Context, uniqueID string) (*models.RefHouseVisited, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	//LookUp
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)

	//query2 := []bson.M{}
	// query2 = append(query2, d.CommonLookup(constants.COLLECTIONLEAVEMASTER, "leavemasterId", "uniqueId", "leavemaster", "leavemaster")...)
	// mainPipeline = append(mainPipeline, bson.M{
	// 	"$lookup": bson.M{
	// 		"from": constants.COLLECTIONPOLICYRULE,
	// 		"as":   "ref.leavemaster",
	// 		"let":  bson.M{"HouseVisited": "$uniqueId"},
	// 		"pipeline": []bson.M{
	// 			{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
	// 				{"$eq": []string{"$$HouseVisited", "$HouseVisitedId"}},
	// 			}}}},
	// 			query2[0],
	// 			{"$addFields": bson.M{"leavemaster": bson.M{"$arrayElemAt": []interface{}{"$leavemaster", 0}}}},
	// 			{"$project": bson.M{"leavemaster": 1}},
	// 		},
	// 	},
	// })
	// mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"ref.leavemaster": "$ref.leavemaster.leavemaster"}})

	// mainPipeline = append(mainPipeline, []bson.M{
	// 	bson.M{"$lookup": bson.M{
	// 		"from": "policyrule", "as": "ref.policyrule", "let": bson.M{"policyId": "$uniqueId"},
	// 		"pipeline": []bson.M{
	// 			bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{

	// 				bson.M{"$eq": []string{"$HouseVisitedId", "$$policyId"}},
	// 			}}}},
	// 			bson.M{"$group": bson.M{"_id": nil, "masterid": bson.M{"$push": "$leavemasterId"}}},
	// 		},
	// 	}},
	// 	bson.M{"$addFields": bson.M{"ref.policyrule": bson.M{"$arrayElemAt": []interface{}{"$ref.policyrule", 0}}}},
	// 	bson.M{"$lookup": bson.M{
	// 		"from": "leavemaster", "as": "ref.policyrule", "let": bson.M{"leaveIDs": "$ref.policyrule.masterid"},
	// 		"pipeline": []bson.M{

	// 			bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
	// 				bson.M{"$in": []string{"$uniqueId", "$$leaveIDs"}},
	// 			}}}},
	// 		},
	// 	},
	// 	},
	// }...)
	d.Shared.BsonToJSONPrintTag("get single leave master", mainPipeline)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONHOUSEVISITED).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var HouseVisiteds []models.RefHouseVisited
	var HouseVisited *models.RefHouseVisited
	if err = cursor.All(ctx.CTX, &HouseVisiteds); err != nil {
		return nil, err
	}
	if len(HouseVisiteds) > 0 {
		HouseVisited = &HouseVisiteds[0]
	}
	return HouseVisited, err
}

//UpdateHouseVisited : ""
func (d *Daos) UpdateHouseVisited(ctx *models.Context, HouseVisited *models.HouseVisited) error {
	selector := bson.M{"uniqueId": HouseVisited.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": HouseVisited}
	_, err := ctx.DB.Collection(constants.COLLECTIONHOUSEVISITED).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// EnableHouseVisited : ""
func (d *Daos) EnableHouseVisited(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.HOUSEVISITEDSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONHOUSEVISITED).UpdateOne(ctx.CTX, selector, data)
	return err
}

// DisableHouseVisited : ""
func (d *Daos) DisableHouseVisited(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.HOUSEVISITEDSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONHOUSEVISITED).UpdateOne(ctx.CTX, selector, data)
	return err
}

//DeleteHouseVisited :""
func (d *Daos) DeleteHouseVisited(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.HOUSEVISITEDSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONHOUSEVISITED).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
func (d *Daos) CollectedHouseVisited(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.HOUSEVISITEDSTATUSCOLLECTED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONHOUSEVISITED).UpdateOne(ctx.CTX, selector, data)
	return err
}

// FilterHouseVisited : ""
func (d *Daos) FilterHouseVisited(ctx *models.Context, houseVisited *models.FilterHouseVisited, pagination *models.Pagination) ([]models.RefHouseVisited, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if houseVisited != nil {
		if len(houseVisited.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": houseVisited.Status}})
		}
		if len(houseVisited.IsStatus) > 0 {
			query = append(query, bson.M{"isStatus": bson.M{"$in": houseVisited.IsStatus}})
		}
		if len(houseVisited.ManagerID) > 0 {
			query = append(query, bson.M{"minUser.id": bson.M{"$in": houseVisited.ManagerID}})
		}
		if len(houseVisited.DumbsiteID) > 0 {
			query = append(query, bson.M{"dumbSite.id": bson.M{"$in": houseVisited.DumbsiteID}})
		}
		if len(houseVisited.GCID) > 0 {
			query = append(query, bson.M{"gcUser.id": bson.M{"$in": houseVisited.GCID}})
		}
		if len(houseVisited.CitizenID) > 0 {
			query = append(query, bson.M{"property.id": bson.M{"$in": houseVisited.CitizenID}})
		}
		//Regex

		if houseVisited.Regex.ManagerName != "" {
			query = append(query, bson.M{"minUser.name": primitive.Regex{Pattern: houseVisited.Regex.ManagerName, Options: "xi"}})
		}
		if houseVisited.Regex.GCName != "" {
			query = append(query, bson.M{"gcUser.name": primitive.Regex{Pattern: houseVisited.Regex.GCName, Options: "xi"}})
		}
		if houseVisited.Regex.CitizenName != "" {
			query = append(query, bson.M{"property.name": primitive.Regex{Pattern: houseVisited.Regex.CitizenName, Options: "xi"}})
		}
	}
	if houseVisited.DateRange.From != nil {
		t := *houseVisited.DateRange.From
		FromDate := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
		ToDate := time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, t.Location())
		if houseVisited.DateRange.To != nil {
			t2 := *houseVisited.DateRange.To
			ToDate = time.Date(t2.Year(), t2.Month(), t2.Day(), 23, 59, 59, 0, t2.Location())
		}
		query = append(query, bson.M{"date": bson.M{"$gte": FromDate, "$lte": ToDate}})

	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	if houseVisited != nil {
		if houseVisited.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{houseVisited.SortBy: houseVisited.SortOrder}})
		}
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONHOUSEVISITED).CountDocuments(ctx.CTX, func() bson.M {
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
	// query2 := []bson.M{}
	// query2 = append(query2, d.CommonLookup(constants.COLLECTIONLEAVEMASTER, "leavemasterId", "uniqueId", "leavemaster", "leavemaster")...)
	// mainPipeline = append(mainPipeline, bson.M{
	// 	"$lookup": bson.M{
	// 		"from": constants.COLLECTIONPOLICYRULE,
	// 		"as":   "ref.leavemaster",
	// 		"let":  bson.M{"HouseVisited": "$uniqueId"},
	// 		"pipeline": []bson.M{
	// 			{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
	// 				{"$eq": []string{"$$HouseVisited", "$HouseVisitedId"}},
	// 			}}}},
	// 			query2[0],
	// 			{"$addFields": bson.M{"leavemaster": bson.M{"$arrayElemAt": []interface{}{"$leavemaster", 0}}}},
	// 			{"$project": bson.M{"leavemaster": 1}},
	// 		},
	// 	},
	// })
	// mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"ref.leavemaster": "$ref.leavemaster.leavemaster"}})
	//Aggregation
	d.Shared.BsonToJSONPrintTag("Feature query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONHOUSEVISITED).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var HouseVisited []models.RefHouseVisited
	if err = cursor.All(context.TODO(), &HouseVisited); err != nil {
		return nil, err
	}
	return HouseVisited, nil
}
