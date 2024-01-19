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

// MySurvey : ""
func (d *Daos) SaveMySurvey(ctx *models.Context, MySurvey *models.MySurvey) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONMYSURVEY).InsertOne(ctx.CTX, MySurvey)
	if err != nil {
		return err
	}
	//MySurvey.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

// GetSingleMySurvey : ""
func (d *Daos) GetSingleMySurvey(ctx *models.Context, uniqueID string) (*models.RefMySurvey, error) {
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
	// 		"let":  bson.M{"MySurvey": "$uniqueId"},
	// 		"pipeline": []bson.M{
	// 			{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
	// 				{"$eq": []string{"$$MySurvey", "$MySurveyId"}},
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

	// 				bson.M{"$eq": []string{"$MySurveyId", "$$policyId"}},
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
	cursor, err := ctx.DB.Collection(constants.COLLECTIONMYSURVEY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var MySurveys []models.RefMySurvey
	var MySurvey *models.RefMySurvey
	if err = cursor.All(ctx.CTX, &MySurveys); err != nil {
		return nil, err
	}
	if len(MySurveys) > 0 {
		MySurvey = &MySurveys[0]
	}
	return MySurvey, err
}

//UpdateMySurvey : ""
func (d *Daos) UpdateMySurvey(ctx *models.Context, MySurvey *models.MySurvey) error {
	selector := bson.M{"uniqueId": MySurvey.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": MySurvey}
	_, err := ctx.DB.Collection(constants.COLLECTIONMYSURVEY).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// EnableMySurvey : ""
func (d *Daos) EnableMySurvey(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.MYSURVEYSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONMYSURVEY).UpdateOne(ctx.CTX, selector, data)
	return err
}

// DisableMySurvey : ""
func (d *Daos) DisableMySurvey(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.MYSURVEYSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONMYSURVEY).UpdateOne(ctx.CTX, selector, data)
	return err
}

//DeleteMySurvey :""
func (d *Daos) DeleteMySurvey(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.MYSURVEYSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONMYSURVEY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterMySurvey : ""
func (d *Daos) FilterMySurvey(ctx *models.Context, mySurvey *models.FilterMySurvey, pagination *models.Pagination) ([]models.RefMySurvey, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if mySurvey != nil {
		if len(mySurvey.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": mySurvey.Status}})
		}
		if len(mySurvey.GCID) > 0 {
			query = append(query, bson.M{"gcUser.id": bson.M{"$in": mySurvey.GCID}})
		}
		if len(mySurvey.ManagerID) > 0 {
			query = append(query, bson.M{"minUser.id": bson.M{"$in": mySurvey.ManagerID}})
		}
		//Regex
		if mySurvey.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: mySurvey.Regex.Name, Options: "xi"}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	if mySurvey != nil {
		if mySurvey.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{mySurvey.SortBy: mySurvey.SortOrder}})
		}
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONMYSURVEY).CountDocuments(ctx.CTX, func() bson.M {
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
	//	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)
	//Aggregation
	d.Shared.BsonToJSONPrintTag("Feature query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONMYSURVEY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var MySurvey []models.RefMySurvey
	if err = cursor.All(context.TODO(), &MySurvey); err != nil {
		return nil, err
	}
	return MySurvey, nil
}

func (d *Daos) UpdateCitizenProperty(ctx *models.Context, citizenProperty *models.MySurvey) error {
	selector := bson.M{"uniqueId": citizenProperty.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": bson.M{
		"citizen":       citizenProperty.GCUser,
		"nfcId":         citizenProperty.NfcID,
		"holdingNumber": citizenProperty.HoldingNumber,
		"mobile":        citizenProperty.Mobile,
	}}
	_, err := ctx.DB.Collection(constants.COLLECTIONMYSURVEY).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}
