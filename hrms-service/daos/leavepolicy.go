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

// SaveLeavePolicy : ""
func (d *Daos) SaveLeavePolicy(ctx *models.Context, leavepolicy *models.LeavePolicy) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONLEAVEPOLICY).InsertOne(ctx.CTX, leavepolicy)
	if err != nil {
		return err
	}
	leavepolicy.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

// GetSingleLeavePolicy : ""
func (d *Daos) GetSingleLeavePolicy(ctx *models.Context, uniqueID string) (*models.RefLeavePolicy, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	//LookUp
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)

	query2 := []bson.M{}
	query2 = append(query2, d.CommonLookup(constants.COLLECTIONLEAVEMASTER, "leavemasterId", "uniqueId", "leavemaster", "leavemaster")...)
	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"from": constants.COLLECTIONPOLICYRULE,
			"as":   "ref.leavemasters",
			"let":  bson.M{"leavePolicy": "$uniqueId"},
			"pipeline": []bson.M{
				{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
					{"$eq": []string{"$$leavePolicy", "$leavepolicyId"}},
					{"$eq": []string{"Active", "$status"}},
				}}}},
				query2[0],
				{"$addFields": bson.M{"leavemaster": bson.M{"$arrayElemAt": []interface{}{"$leavemaster", 0}}}},
				{"$addFields": bson.M{"uniqueId": "$leavemaster.uniqueId"}},
				{"$addFields": bson.M{"name": "$leavemaster.name"}},
				{"$addFields": bson.M{"leaveType": "$leavemaster.leaveType"}},
				{"$project": bson.M{"uniqueId": 1, "name": 1, "value": 1, "leaveType": 1}},
				//  {"$project": bson.M{"leavemaster": 1}},
			},
		},
	})
	//mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"ref.leavemasters": "$ref.leavemaster.leavemaster"}})

	// mainPipeline = append(mainPipeline, []bson.M{
	// 	bson.M{"$lookup": bson.M{
	// 		"from": "policyrule", "as": "ref.policyrule", "let": bson.M{"policyId": "$uniqueId"},
	// 		"pipeline": []bson.M{
	// 			bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{

	// 				bson.M{"$eq": []string{"$leavepolicyId", "$$policyId"}},
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
	cursor, err := ctx.DB.Collection(constants.COLLECTIONLEAVEPOLICY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var leavepolicys []models.RefLeavePolicy
	var leavepolicy *models.RefLeavePolicy
	if err = cursor.All(ctx.CTX, &leavepolicys); err != nil {
		return nil, err
	}
	if len(leavepolicys) > 0 {
		leavepolicy = &leavepolicys[0]
	}
	return leavepolicy, err
}

//UpdateLeavePolicy : ""
func (d *Daos) UpdateLeavePolicy(ctx *models.Context, leavepolicy *models.LeavePolicy) error {
	selector := bson.M{"uniqueId": leavepolicy.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": leavepolicy}
	_, err := ctx.DB.Collection(constants.COLLECTIONLEAVEPOLICY).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// EnableLeavePolicy : ""
func (d *Daos) EnableLeavePolicy(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.LEAVEPOLICYSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONLEAVEPOLICY).UpdateOne(ctx.CTX, selector, data)
	return err
}

// DisableLeavePolicy : ""
func (d *Daos) DisableLeavePolicy(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.LEAVEPOLICYSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONLEAVEPOLICY).UpdateOne(ctx.CTX, selector, data)
	return err
}

//DeleteLeavePolicy :""
func (d *Daos) DeleteLeavePolicy(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.LEAVEPOLICYSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONLEAVEPOLICY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterLeavePolicy : ""
func (d *Daos) FilterLeavePolicy(ctx *models.Context, leavepolicy *models.FilterLeavePolicy, pagination *models.Pagination) ([]models.RefLeavePolicy, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if leavepolicy != nil {
		if len(leavepolicy.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": leavepolicy.Status}})
		}
		if len(leavepolicy.OrganisationID) > 0 {
			query = append(query, bson.M{"organisationId": bson.M{"$in": leavepolicy.OrganisationID}})
		}
		//Regex
		if leavepolicy.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: leavepolicy.Regex.Name, Options: "xi"}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	if leavepolicy != nil {
		if leavepolicy.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{leavepolicy.SortBy: leavepolicy.SortOrder}})
		}
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONLEAVEPOLICY).CountDocuments(ctx.CTX, func() bson.M {
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
	query2 := []bson.M{}
	query2 = append(query2, d.CommonLookup(constants.COLLECTIONLEAVEMASTER, "leavemasterId", "uniqueId", "leavemaster", "leavemaster")...)
	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"from": constants.COLLECTIONPOLICYRULE,
			"as":   "ref.leavemasters",
			"let":  bson.M{"leavePolicy": "$uniqueId"},
			"pipeline": []bson.M{
				{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
					{"$eq": []string{"$$leavePolicy", "$leavepolicyId"}},
					{"$eq": []string{"Active", "$status"}},
				}}}},
				query2[0],
				{"$addFields": bson.M{"leavemaster": bson.M{"$arrayElemAt": []interface{}{"$leavemaster", 0}}}},
				{"$addFields": bson.M{"uniqueId": "$leavemaster.uniqueId"}},
				{"$addFields": bson.M{"name": "$leavemaster.name"}},
				{"$addFields": bson.M{"leaveType": "$leavemaster.leaveType"}},
				{"$project": bson.M{"name": 1, "value": 1, "leaveType": 1, "uniqueId": 1}},
				//  {"$project": bson.M{"leavemaster": 1}},
			},
		},
	})
	//mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"ref.leavemaster": "$ref.leavemaster.leavemaster"}})
	//Aggregation
	d.Shared.BsonToJSONPrintTag("Feature query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONLEAVEPOLICY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var leavePolicy []models.RefLeavePolicy
	if err = cursor.All(context.TODO(), &leavePolicy); err != nil {
		return nil, err
	}
	return leavePolicy, nil
}

// GetSingleLeavePolicyWithActiveName : ""
func (d *Daos) GetSingleLeavePolicyWithActiveName(ctx *models.Context, uniqueID string) (*models.RefLeavePolicy, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"name": uniqueID, "status": constants.ONBOARDINGPOLICYSTATUSACTIVE}})
	//LookUp
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)

	query2 := []bson.M{}
	query2 = append(query2, d.CommonLookup(constants.COLLECTIONLEAVEMASTER, "leavemasterId", "uniqueId", "leavemaster", "leavemaster")...)
	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"from": constants.COLLECTIONPOLICYRULE,
			"as":   "ref.leavemasters",
			"let":  bson.M{"leavePolicy": "$uniqueId"},
			"pipeline": []bson.M{
				{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
					{"$eq": []string{"$$leavePolicy", "$leavepolicyId"}},
					{"$eq": []string{"Active", "$status"}},
				}}}},
				query2[0],
				{"$addFields": bson.M{"leavemaster": bson.M{"$arrayElemAt": []interface{}{"$leavemaster", 0}}}},
				{"$addFields": bson.M{"uniqueId": "$leavemaster.uniqueId"}},
				{"$addFields": bson.M{"name": "$leavemaster.name"}},
				{"$addFields": bson.M{"leaveType": "$leavemaster.leaveType"}},
				{"$project": bson.M{"uniqueId": 1, "name": 1, "value": 1, "leaveType": 1}},
				//  {"$project": bson.M{"leavemaster": 1}},
			},
		},
	})
	//mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"ref.leavemasters": "$ref.leavemaster.leavemaster"}})

	// mainPipeline = append(mainPipeline, []bson.M{
	// 	bson.M{"$lookup": bson.M{
	// 		"from": "policyrule", "as": "ref.policyrule", "let": bson.M{"policyId": "$uniqueId"},
	// 		"pipeline": []bson.M{
	// 			bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{

	// 				bson.M{"$eq": []string{"$leavepolicyId", "$$policyId"}},
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
	cursor, err := ctx.DB.Collection(constants.COLLECTIONLEAVEPOLICY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var leavepolicys []models.RefLeavePolicy
	var leavepolicy *models.RefLeavePolicy
	if err = cursor.All(ctx.CTX, &leavepolicys); err != nil {
		return nil, err
	}
	if len(leavepolicys) > 0 {
		leavepolicy = &leavepolicys[0]
	}
	return leavepolicy, err
}
