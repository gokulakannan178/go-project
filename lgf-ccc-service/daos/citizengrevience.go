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

// CitizenGrevience : ""
func (d *Daos) SaveCitizenGrevience(ctx *models.Context, CitizenGrevience *models.CitizenGrevience) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONCITIZENGREVIENCE).InsertOne(ctx.CTX, CitizenGrevience)
	if err != nil {
		return err
	}
	//CitizenGrevience.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

// GetSingleCitizenGrevience : ""
func (d *Daos) GetSingleCitizenGrevience(ctx *models.Context, uniqueID string) (*models.RefCitizenGrevience, error) {
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
	// 		"let":  bson.M{"CitizenGrevience": "$uniqueId"},
	// 		"pipeline": []bson.M{
	// 			{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
	// 				{"$eq": []string{"$$CitizenGrevience", "$CitizenGrevienceId"}},
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

	// 				bson.M{"$eq": []string{"$CitizenGrevienceId", "$$policyId"}},
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
	cursor, err := ctx.DB.Collection(constants.COLLECTIONCITIZENGREVIENCE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var CitizenGreviences []models.RefCitizenGrevience
	var CitizenGrevience *models.RefCitizenGrevience
	if err = cursor.All(ctx.CTX, &CitizenGreviences); err != nil {
		return nil, err
	}
	if len(CitizenGreviences) > 0 {
		CitizenGrevience = &CitizenGreviences[0]
	}
	return CitizenGrevience, err
}

// UpdateCitizenGrevience : ""
func (d *Daos) UpdateCitizenGrevience(ctx *models.Context, CitizenGrevience *models.CitizenGrevience) error {
	selector := bson.M{"uniqueId": CitizenGrevience.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": CitizenGrevience}
	_, err := ctx.DB.Collection(constants.COLLECTIONCITIZENGREVIENCE).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// EnableCitizenGrevience : ""
func (d *Daos) EnableCitizenGrevience(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.CITIZENGREVIENCESTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONCITIZENGREVIENCE).UpdateOne(ctx.CTX, selector, data)
	return err
}

// DisableCitizenGrevience : ""
func (d *Daos) DisableCitizenGrevience(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.CITIZENGREVIENCESTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONCITIZENGREVIENCE).UpdateOne(ctx.CTX, selector, data)
	return err
}

// DeleteCitizenGrevience :""
func (d *Daos) DeleteCitizenGrevience(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.CITIZENGREVIENCESTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONCITIZENGREVIENCE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterCitizenGrevience : ""
func (d *Daos) FilterCitizenGrevience(ctx *models.Context, CitizenGrevience *models.FilterCitizenGrevience, pagination *models.Pagination) ([]models.RefCitizenGrevience, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if CitizenGrevience != nil {
		if len(CitizenGrevience.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": CitizenGrevience.Status}})
		}
		if len(CitizenGrevience.GCID) > 0 {
			query = append(query, bson.M{"gcUser.id": bson.M{"$in": CitizenGrevience.GCID}})
		}
		if len(CitizenGrevience.ManagerID) > 0 {
			query = append(query, bson.M{"minUser.id": bson.M{"$in": CitizenGrevience.ManagerID}})
		}
		//Regex
		if CitizenGrevience.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: CitizenGrevience.Regex.Name, Options: "xi"}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	if CitizenGrevience != nil {
		if CitizenGrevience.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{CitizenGrevience.SortBy: CitizenGrevience.SortOrder}})
		}
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONCITIZENGREVIENCE).CountDocuments(ctx.CTX, func() bson.M {
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
	cursor, err := ctx.DB.Collection(constants.COLLECTIONCITIZENGREVIENCE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var Citizengrevience []models.RefCitizenGrevience
	if err = cursor.All(context.TODO(), &CitizenGrevience); err != nil {
		return nil, err
	}
	return Citizengrevience, nil
}

// func (d *Daos) UpdateCitizenGrevience(ctx *models.Context, citizenProperty *models.CitizenGrevience) error {
// 	selector := bson.M{"uniqueId": citizenProperty.UniqueID}
// 	t := time.Now()
// 	update := models.Updated{}
// 	update.On = &t
// 	update.By = constants.SYSTEM
// 	updateInterface := bson.M{"$set": bson.M{
// 		"citizen":       citizenProperty.GCUser,
// 		"nfcId":         citizenProperty.NfcID,
// 		"holdingNumber": citizenProperty.HoldingNumber,
// 		"mobile":        citizenProperty.Mobile,
// 	}}
// 	_, err := ctx.DB.Collection(constants.COLLECTIONCITIZENGREVIENCE).UpdateOne(ctx.CTX, selector, updateInterface)
// 	if err != nil {
// 		fmt.Println("Not changed", err.Error())
// 		return err
// 	}
// 	return err
// }
