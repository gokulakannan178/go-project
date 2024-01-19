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

//SavePayrollPolicy : ""
func (d *Daos) SavePayrollPolicy(ctx *models.Context, payrollPolicy *models.PayrollPolicy) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONPAYROLLPOLICY).InsertOne(ctx.CTX, payrollPolicy)
	if err != nil {
		return err
	}
	payrollPolicy.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//UpdatePayrollPolicy : ""
func (d *Daos) UpdatePayrollPolicy(ctx *models.Context, payrollPolicy *models.PayrollPolicy) error {
	selector := bson.M{"uniqueId": payrollPolicy.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": payrollPolicy}
	_, err := ctx.DB.Collection(constants.COLLECTIONPAYROLLPOLICY).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// GetSinglePayrollPolicy : ""
func (d *Daos) GetSinglePayrollPolicy(ctx *models.Context, uniqueID string) (*models.RefPayrollPolicy, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	//LookUp
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)
	mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{
		"from": constants.COLLECTIONPAYROLLPOLICYEARNING,
		"as":   "ref.earningMaster",
		"let":  bson.M{"payroll": "$uniqueId", "earningId": "$earningMaster.uniqueId"},
		"pipeline": []bson.M{
			{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
				{"$eq": []string{"$status", "Active"}},
				{"$eq": []string{"$$payroll", "$payRollPolicyId"}},
				// {"$eq":["$earningMasterId","$$earningId"]}
			}}}},
			bson.M{"$lookup": bson.M{
				"from":         constants.COLLECTIONEMPLOYEEEARNINGMASTER,
				"as":           "ref.earningMasterId",
				"localField":   "earningMasterId",
				"foreignField": "uniqueId",
			}},
			bson.M{"$addFields": bson.M{"ref.earningMasterId": bson.M{"$arrayElemAt": []interface{}{"$ref.earningMasterId", 0}}}},
			bson.M{"$addFields": bson.M{"name": "$ref.earningMasterId.title"}},
		}}})
	mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{
		"from": constants.COLLECTIONPAYROLLPOLICYDETECTION,
		"as":   "ref.detectionMaster",
		"let":  bson.M{"payroll": "$uniqueId", "earningId": "$detectionMaster.uniqueId"},
		"pipeline": []bson.M{
			{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
				{"$eq": []string{"$status", "Active"}},
				{"$eq": []string{"$$payroll", "$payRollPolicyId"}},
				// {"$eq":["$earningMasterId","$$earningId"]}
			}}}},
			bson.M{"$lookup": bson.M{
				"from":         constants.COLLECTIONEMPLOYEEDEDUCTIONMASTER,
				"as":           "ref.detectionMasterId",
				"localField":   "detectionMasterId",
				"foreignField": "uniqueId",
			}},
			bson.M{"$addFields": bson.M{"ref.detectionMasterId": bson.M{"$arrayElemAt": []interface{}{"$ref.detectionMasterId", 0}}}},
			bson.M{"$addFields": bson.M{"name": "$ref.detectionMasterId.title"}},
		}}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPAYROLLPOLICY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var PayrollPolicys []models.RefPayrollPolicy
	var PayrollPolicy *models.RefPayrollPolicy
	if err = cursor.All(ctx.CTX, &PayrollPolicys); err != nil {
		return nil, err
	}
	if len(PayrollPolicys) > 0 {
		PayrollPolicy = &PayrollPolicys[0]
	}
	return PayrollPolicy, err
}

// EnablePayrollPolicy : ""
func (d *Daos) EnablePayrollPolicy(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.PAYROLLPOLICYSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPAYROLLPOLICY).UpdateOne(ctx.CTX, selector, data)
	return err
}

// DisablePayrollPolicy : ""
func (d *Daos) DisablePayrollPolicy(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.PAYROLLPOLICYSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPAYROLLPOLICY).UpdateOne(ctx.CTX, selector, data)
	return err
}

//DeletePayrollPolicy :""
func (d *Daos) DeletePayrollPolicy(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PAYROLLPOLICYSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPAYROLLPOLICY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterPayrollPolicy : ""
func (d *Daos) FilterPayrollPolicy(ctx *models.Context, payrollPolicy *models.FilterPayrollPolicy, pagination *models.Pagination) ([]models.RefPayrollPolicy, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if payrollPolicy != nil {
		if len(payrollPolicy.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": payrollPolicy.Status}})
		}
		if len(payrollPolicy.OrganisationID) > 0 {
			query = append(query, bson.M{"organisationId": bson.M{"$in": payrollPolicy.OrganisationID}})
		}
		//Regex
		if payrollPolicy.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: payrollPolicy.Regex.Name, Options: "xi"}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONPAYROLLPOLICY).CountDocuments(ctx.CTX, func() bson.M {
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

	mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{
		"from": constants.COLLECTIONPAYROLLPOLICYEARNING,
		"as":   "ref.earningMaster",
		"let":  bson.M{"payroll": "$uniqueId", "earningId": "$earningMaster.uniqueId"},
		"pipeline": []bson.M{
			{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
				{"$eq": []string{"$status", "Active"}},
				{"$eq": []string{"$$payroll", "$payRollPolicyId"}},
				// {"$eq":["$earningMasterId","$$earningId"]}
			}}}},
			bson.M{"$lookup": bson.M{
				"from":         constants.COLLECTIONEMPLOYEEEARNINGMASTER,
				"as":           "ref.earningMasterId",
				"localField":   "earningMasterId",
				"foreignField": "uniqueId",
			}},
			bson.M{"$addFields": bson.M{"ref.earningMasterId": bson.M{"$arrayElemAt": []interface{}{"$ref.earningMasterId", 0}}}},
			bson.M{"$addFields": bson.M{"name": "$ref.earningMasterId.title"}},
		}}})
	mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{
		"from": constants.COLLECTIONPAYROLLPOLICYDETECTION,
		"as":   "ref.detectionMaster",
		"let":  bson.M{"payroll": "$uniqueId", "earningId": "$detectionMaster.uniqueId"},
		"pipeline": []bson.M{
			{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
				{"$eq": []string{"$status", "Active"}},
				{"$eq": []string{"$$payroll", "$payRollPolicyId"}},
				// {"$eq":["$earningMasterId","$$earningId"]}
			}}}},
			bson.M{"$lookup": bson.M{
				"from":         constants.COLLECTIONEMPLOYEEDEDUCTIONMASTER,
				"as":           "ref.detectionMasterId",
				"localField":   "detectionMasterId",
				"foreignField": "uniqueId",
			}},
			bson.M{"$addFields": bson.M{"ref.detectionMasterId": bson.M{"$arrayElemAt": []interface{}{"$ref.detectionMasterId", 0}}}},
			bson.M{"$addFields": bson.M{"name": "$ref.detectionMasterId.title"}},
		}}})
	//Aggregation
	d.Shared.BsonToJSONPrintTag("Feature query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPAYROLLPOLICY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var payrollPolicyFilter []models.RefPayrollPolicy
	if err = cursor.All(context.TODO(), &payrollPolicyFilter); err != nil {
		return nil, err
	}
	return payrollPolicyFilter, nil
}
