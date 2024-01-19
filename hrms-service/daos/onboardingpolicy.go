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

// SaveOnboardingPolicy : ""
func (d *Daos) SaveOnboardingPolicy(ctx *models.Context, onboardingpolicy *models.OnboardingPolicy) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONONBOARDINGPOLICY).InsertOne(ctx.CTX, onboardingpolicy)
	if err != nil {
		return err
	}
	onboardingpolicy.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//UpdateOnboardingPolicy : ""
func (d *Daos) UpdateOnboardingPolicy(ctx *models.Context, OnboardingPolicy *models.OnboardingPolicy) error {
	selector := bson.M{"uniqueId": OnboardingPolicy.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": OnboardingPolicy}
	_, err := ctx.DB.Collection(constants.COLLECTIONONBOARDINGPOLICY).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// GetSingleOnboardingPolicy : ""
func (d *Daos) GetSingleOnboardingPolicy(ctx *models.Context, uniqueID string) (*models.RefOnboardingPolicy, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	//LookUp
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)
	mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONONBOARDINGCHECKLIST, "uniqueId", "onboardingpolicyId", "ref.onboardingchecklist", "ref.onboardingchecklist")...)
	// mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{
	// 	"from": "onboardingchecklist",
	// 	"as":   "ref.onboardingchecklist",
	// 	"let":  bson.M{"policyId": "$uniqueId"},
	// 	"pipeline": []bson.M{
	// 		bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
	// 			bson.M{"$eq": []string{"$onboardingpolicyId", "$$policyId"}},
	// 		}}}},
	// 		bson.M{"$group": bson.M{"_id": nil, "checklistMasterIds": bson.M{"$push": "$onboardingchecklistmasterId"}}},
	// 		bson.M{"$lookup": bson.M{
	// 			"from": "onboardingchecklistmaster",
	// 			"as":   "checklistMasterIds",

	// 			"let": bson.M{"checklistMasterIds": "$checklistMasterIds"},
	// 			"pipeline": []bson.M{
	// 				bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
	// 					bson.M{"$in": []string{"$uniqueId", "$$checklistMasterIds"}},
	// 				}}}},
	// 			},
	// 		}},
	// 	},
	// }})
	// mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"ref.onboardingchecklist": bson.M{"$arrayElemAt": []interface{}{"$ref.onboardingchecklist", 0}}}})
	// mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"ref.onboardingchecklist": "$ref.onboardingchecklist.checklistMasterIds"}})
	query2 := []bson.M{}
	query2 = append(query2, d.CommonLookup(constants.COLLECTIONONBOARDINGCHECKLISTMASTER, "onboardingchecklistmasterId", "uniqueId", "onboardingchecklistmaster", "onboardingchecklistmaster")...)
	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"from": constants.COLLECTIONONBOARDINGCHECKLIST,
			"as":   "ref.onboardingchecklist",
			"let":  bson.M{"onboardingpolicyId": "$uniqueId"},
			"pipeline": []bson.M{
				{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
					{"$eq": []string{"$$onboardingpolicyId", "$onboardingpolicyId"}},
				}}}},
				query2[0],
				{"$addFields": bson.M{"onboardingchecklistmaster": bson.M{"$arrayElemAt": []interface{}{"$onboardingchecklistmaster", 0}}}},
				{"$project": bson.M{"onboardingchecklistmaster": 1}},
			},
		},
	})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"ref.onboardingchecklist": "$ref.onboardingchecklist.onboardingchecklistmaster"}})

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONONBOARDINGPOLICY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var onboardingpolicys []models.RefOnboardingPolicy
	var onboardingpolicy *models.RefOnboardingPolicy
	if err = cursor.All(ctx.CTX, &onboardingpolicys); err != nil {
		return nil, err
	}
	if len(onboardingpolicys) > 0 {
		onboardingpolicy = &onboardingpolicys[0]
	}
	return onboardingpolicy, err
}

// GetSingleOnboardingPolicy : ""
func (d *Daos) GetSingleActiveOnboardingPolicyWithName(ctx *models.Context, uniqueID string) (*models.RefOnboardingPolicy, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"name": uniqueID, "status": constants.ONBOARDINGPOLICYSTATUSACTIVE}})

	d.Shared.BsonToJSONPrintTag("GetSingleActiveOnboardingPolicyWithName query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONONBOARDINGPOLICY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var onboardingpolicys []models.RefOnboardingPolicy
	var onboardingpolicy *models.RefOnboardingPolicy
	if err = cursor.All(ctx.CTX, &onboardingpolicys); err != nil {
		return nil, err
	}
	if len(onboardingpolicys) > 0 {
		onboardingpolicy = &onboardingpolicys[0]
	}
	return onboardingpolicy, err
}

// GetSingleOnboardingPolicy : ""
func (d *Daos) GetSingleOnboardingPolicyWithActiveStatus(ctx *models.Context, uniqueID string, Status string) (*models.RefOnboardingPolicy, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID, "status": Status}})
	//LookUp
	//mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)
	mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONONBOARDINGCHECKLIST, "uniqueId", "onboardingpolicyId", "ref.onboardingchecklist", "ref.onboardingchecklist")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONONBOARDINGPOLICY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var onboardingpolicys []models.RefOnboardingPolicy
	var onboardingpolicy *models.RefOnboardingPolicy
	if err = cursor.All(ctx.CTX, &onboardingpolicys); err != nil {
		return nil, err
	}
	if len(onboardingpolicys) > 0 {
		onboardingpolicy = &onboardingpolicys[0]
	}
	return onboardingpolicy, err
}

// EnableOnboardingPolicy : ""
func (d *Daos) EnableOnboardingPolicy(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.ONBOARDINGPOLICYSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONONBOARDINGPOLICY).UpdateOne(ctx.CTX, selector, data)
	return err
}

// DisableOnboardingPolicy : ""
func (d *Daos) DisableOnboardingPolicy(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.ONBOARDINGPOLICYSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONONBOARDINGPOLICY).UpdateOne(ctx.CTX, selector, data)
	return err
}

//DeleteOnboardingPolicy :""
func (d *Daos) DeleteOnboardingPolicy(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.ONBOARDINGPOLICYSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONONBOARDINGPOLICY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterOnboardingPolicy : ""
func (d *Daos) FilterOnboardingPolicy(ctx *models.Context, onboardingPolicy *models.FilterOnboardingPolicy, pagination *models.Pagination) ([]models.RefOnboardingPolicy, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if onboardingPolicy != nil {
		if len(onboardingPolicy.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": onboardingPolicy.Status}})
		}
		if len(onboardingPolicy.OrganisationID) > 0 {
			query = append(query, bson.M{"organisationId": bson.M{"$in": onboardingPolicy.OrganisationID}})
		}
		//Regex
		if onboardingPolicy.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: onboardingPolicy.Regex.Name, Options: "xi"}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if onboardingPolicy != nil {
		if onboardingPolicy.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{onboardingPolicy.SortBy: onboardingPolicy.SortOrder}})
		}
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONONBOARDINGPOLICY).CountDocuments(ctx.CTX, func() bson.M {
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
	// mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{
	// 	"from": "onboardingchecklist",
	// 	"as":   "ref.onboardingchecklist",
	// 	"let":  bson.M{"policyId": "$uniqueId"},
	// 	"pipeline": []bson.M{
	// 		bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
	// 			bson.M{"$eq": []string{"$onboardingpolicyId", "$$policyId"}},
	// 		}}}},
	// 		bson.M{"$group": bson.M{"_id": nil, "checklistMasterIds": bson.M{"$push": "$onboardingchecklistmasterId"}}},
	// 		bson.M{"$lookup": bson.M{
	// 			"from": "onboardingchecklistmaster",
	// 			"as":   "checklistMasterIds",

	// 			"let": bson.M{"checklistMasterIds": "$checklistMasterIds"},
	// 			"pipeline": []bson.M{
	// 				bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
	// 					bson.M{"$in": []string{"$uniqueId", "$$checklistMasterIds"}},
	// 				}}}},
	// 			},
	// 		}},
	// 	},
	// }})
	// mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"ref.onboardingchecklist": bson.M{"$arrayElemAt": []interface{}{"$ref.onboardingchecklist", 0}}}})
	// mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"ref.onboardingchecklist": "$ref.onboardingchecklist.checklistMasterIds"}})
	query2 := []bson.M{}
	query2 = append(query2, d.CommonLookup(constants.COLLECTIONONBOARDINGCHECKLISTMASTER, "onboardingchecklistmasterId", "uniqueId", "onboardingchecklistmaster", "onboardingchecklistmaster")...)
	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"from": constants.COLLECTIONONBOARDINGCHECKLIST,
			"as":   "ref.onboardingchecklist",
			"let":  bson.M{"onboardingpolicyId": "$uniqueId"},
			"pipeline": []bson.M{
				{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
					{"$eq": []string{"$$onboardingpolicyId", "$onboardingpolicyId"}},
				}}}},
				query2[0],
				{"$addFields": bson.M{"onboardingchecklistmaster": bson.M{"$arrayElemAt": []interface{}{"$onboardingchecklistmaster", 0}}}},
				{"$project": bson.M{"onboardingchecklistmaster": 1}},
			},
		},
	})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"ref.onboardingchecklist": "$ref.onboardingchecklist.onboardingchecklistmaster"}})

	//Aggregation
	d.Shared.BsonToJSONPrintTag("Feature query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONONBOARDINGPOLICY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var onboardingpolicyFilter []models.RefOnboardingPolicy
	if err = cursor.All(context.TODO(), &onboardingpolicyFilter); err != nil {
		return nil, err
	}
	return onboardingpolicyFilter, nil
}
