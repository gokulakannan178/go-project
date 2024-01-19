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

// SaveOffboardingPolicy : ""
func (d *Daos) SaveOffboardingPolicy(ctx *models.Context, offboardingpolicy *models.OffboardingPolicy) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONOFFBOARDINGPOLICY).InsertOne(ctx.CTX, offboardingpolicy)
	if err != nil {
		return err
	}
	offboardingpolicy.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//UpdateOffboardingPolicy : ""
func (d *Daos) UpdateOffboardingPolicy(ctx *models.Context, OffboardingPolicy *models.OffboardingPolicy) error {
	selector := bson.M{"uniqueId": OffboardingPolicy.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": OffboardingPolicy}
	_, err := ctx.DB.Collection(constants.COLLECTIONOFFBOARDINGPOLICY).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// GetSingleOffboardingPolicy : ""
func (d *Daos) GetSingleOffboardingPolicy(ctx *models.Context, uniqueID string) (*models.RefOffboardingPolicy, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	//LookUp
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)
	//	mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONOFFBOARDINGCHECKLIST, "uniqueId", "offboardingpolicyId", "ref.offboardingchecklistId", "ref.offboardingchecklistId")...)

	//LookUp
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)
	//	mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONOFFBOARDINGCHECKLIST, "uniqueId", "offboardingpolicyId", "ref.offboardingchecklistId", "ref.offboardingchecklistId")...)
	query2 := []bson.M{}
	query2 = append(query2, d.CommonLookup(constants.COLLECTIONOFFBOARDINGCHECKLISTMASTER, "offboardingchecklistmasterId", "uniqueId", "offboardingchecklistmasterId", "offboardingchecklistmasterId")...)

	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"from": constants.COLLECTIONOFFBOARDINGCHECKLIST,
			"as":   "ref.offboardingchecklistId",
			"let":  bson.M{"offboardingpolicy": "$uniqueId"},
			"pipeline": []bson.M{
				bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
					{"$eq": []string{"$offboardingpolicyId", "$$offboardingpolicy"}},
					{"$eq": []string{"$status", "Active"}},
				}}}},
				query2[0],
				{"$addFields": bson.M{"offboardingchecklistmasterId": bson.M{"$arrayElemAt": []interface{}{"$offboardingchecklistmasterId", 0}}}},
				{"$project": bson.M{"offboardingchecklistmasterId": 1}},
			},
		},
	})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"ref.offboardingchecklistId": "$ref.offboardingchecklistId.offboardingchecklistmasterId"}})

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONOFFBOARDINGPOLICY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var offboardingpolicys []models.RefOffboardingPolicy
	var offboardingpolicy *models.RefOffboardingPolicy
	if err = cursor.All(ctx.CTX, &offboardingpolicys); err != nil {
		return nil, err
	}
	if len(offboardingpolicys) > 0 {
		offboardingpolicy = &offboardingpolicys[0]
	}
	return offboardingpolicy, err
}

// EnableOffboardingPolicy : ""
func (d *Daos) EnableOffboardingPolicy(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.OFFBOARDINGPOLICYSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONOFFBOARDINGPOLICY).UpdateOne(ctx.CTX, selector, data)
	return err
}

// DisableOffboardingPolicy : ""
func (d *Daos) DisableOffboardingPolicy(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.OFFBOARDINGPOLICYSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONOFFBOARDINGPOLICY).UpdateOne(ctx.CTX, selector, data)
	return err
}

//DeleteOffboardingPolicy :""
func (d *Daos) DeleteOffboardingPolicy(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.OFFBOARDINGPOLICYSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONOFFBOARDINGPOLICY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterOffboardingPolicy : ""
func (d *Daos) FilterOffboardingPolicy(ctx *models.Context, offboardingpolicy *models.FilterOffboardingPolicy, pagination *models.Pagination) ([]models.RefOffboardingPolicy, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if offboardingpolicy != nil {
		if len(offboardingpolicy.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": offboardingpolicy.Status}})
		}
		if len(offboardingpolicy.OrganisationID) > 0 {
			query = append(query, bson.M{"organisationId": bson.M{"$in": offboardingpolicy.OrganisationID}})
		}
		//Regex
		if offboardingpolicy.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: offboardingpolicy.Regex.Name, Options: "xi"}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	if offboardingpolicy != nil {
		if offboardingpolicy.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{offboardingpolicy.SortBy: offboardingpolicy.SortOrder}})
		}
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONOFFBOARDINGPOLICY).CountDocuments(ctx.CTX, func() bson.M {
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
	query2 := []bson.M{}
	query2 = append(query2, d.CommonLookup(constants.COLLECTIONOFFBOARDINGCHECKLISTMASTER, "offboardingchecklistmasterId", "uniqueId", "offboardingchecklistmasterId", "offboardingchecklistmasterId")...)

	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"from": constants.COLLECTIONOFFBOARDINGCHECKLIST,
			"as":   "ref.offboardingchecklistId",
			"let":  bson.M{"offboardingpolicy": "$uniqueId"},
			"pipeline": []bson.M{
				bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
					{"$eq": []string{"$offboardingpolicyId", "$$offboardingpolicy"}},
					{"$eq": []string{"$status", "Active"}},
				}}}},
				query2[0],
				{"$addFields": bson.M{"offboardingchecklistmasterId": bson.M{"$arrayElemAt": []interface{}{"$offboardingchecklistmasterId", 0}}}},
				{"$project": bson.M{"offboardingchecklistmasterId": 1}},
			},
		},
	})
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"ref.offboardingchecklistId": "$ref.offboardingchecklistId.offboardingchecklistmasterId"}})

	//Aggregation
	d.Shared.BsonToJSONPrintTag("Feature query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONOFFBOARDINGPOLICY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var offboardingpolicyFilter []models.RefOffboardingPolicy
	if err = cursor.All(context.TODO(), &offboardingpolicyFilter); err != nil {
		return nil, err
	}
	return offboardingpolicyFilter, nil
}
