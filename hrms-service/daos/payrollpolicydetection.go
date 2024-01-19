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
	"go.mongodb.org/mongo-driver/mongo/options"
)

//SavePayrollPolicyDetection : ""
func (d *Daos) SavePayrollPolicyDetection(ctx *models.Context, payrollPolicyDetection *models.PayrollPolicyDetection) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONPAYROLLPOLICYDETECTION).InsertOne(ctx.CTX, payrollPolicyDetection)
	if err != nil {
		return err
	}
	payrollPolicyDetection.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//UpdatePayrollPolicyDetection : ""
func (d *Daos) UpdatePayrollPolicyDetection(ctx *models.Context, payrollPolicyDetection *models.PayrollPolicyDetection) error {
	selector := bson.M{"uniqueId": payrollPolicyDetection.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": payrollPolicyDetection}
	_, err := ctx.DB.Collection(constants.COLLECTIONPAYROLLPOLICYDETECTION).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}
func (d *Daos) UpdatePayrollPolicyEarningWithDetection(ctx *models.Context, payrollPolicyDetection *models.PayrollPolicyDetection) error {
	opts := options.Update().SetUpsert(true)
	selector := bson.M{"payRollPolicyId": payrollPolicyDetection.PayRollPolicyId, "detectionMasterId": payrollPolicyDetection.DetectionMasterId}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": payrollPolicyDetection}
	_, err := ctx.DB.Collection(constants.COLLECTIONPAYROLLPOLICYDETECTION).UpdateOne(ctx.CTX, selector, updateInterface, opts)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// GetSinglePayrollPolicyDetection : ""
func (d *Daos) GetSinglePayrollPolicyDetection(ctx *models.Context, uniqueID string) (*models.RefPayrollPolicyDetection, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	//LookUp
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPAYROLLPOLICYDETECTION).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var PayrollPolicyDetections []models.RefPayrollPolicyDetection
	var PayrollPolicyDetection *models.RefPayrollPolicyDetection
	if err = cursor.All(ctx.CTX, &PayrollPolicyDetections); err != nil {
		return nil, err
	}
	if len(PayrollPolicyDetections) > 0 {
		PayrollPolicyDetection = &PayrollPolicyDetections[0]
	}
	return PayrollPolicyDetection, err
}

// EnablePayrollPolicyDetection : ""
func (d *Daos) EnablePayrollPolicyDetection(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.PAYROLLPOLICYDETECTIONSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPAYROLLPOLICYDETECTION).UpdateOne(ctx.CTX, selector, data)
	return err
}

// DisablePayrollPolicyDetection : ""
func (d *Daos) DisablePayrollPolicyDetection(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.PAYROLLPOLICYDETECTIONSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPAYROLLPOLICYDETECTION).UpdateOne(ctx.CTX, selector, data)
	return err
}

//DeletePayrollPolicyDetection :""
func (d *Daos) DeletePayrollPolicyDetection(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PAYROLLPOLICYDETECTIONSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPAYROLLPOLICYDETECTION).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterPayrollPolicyDetection : ""
func (d *Daos) FilterPayrollPolicyDetection(ctx *models.Context, payrollPolicyDetection *models.FilterPayrollPolicyDetection, pagination *models.Pagination) ([]models.RefPayrollPolicyDetection, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if payrollPolicyDetection != nil {
		if len(payrollPolicyDetection.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": payrollPolicyDetection.Status}})
		}
		if len(payrollPolicyDetection.OrganisationID) > 0 {
			query = append(query, bson.M{"organisationId": bson.M{"$in": payrollPolicyDetection.OrganisationID}})
		}
		//Regex
		if payrollPolicyDetection.Regex.Name != "" {
			query = append(query, bson.M{"title": primitive.Regex{Pattern: payrollPolicyDetection.Regex.Name, Options: "xi"}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONPAYROLLPOLICYDETECTION).CountDocuments(ctx.CTX, func() bson.M {
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

	//Aggregation
	d.Shared.BsonToJSONPrintTag("Feature query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPAYROLLPOLICYDETECTION).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var payrollPolicyDetectionFilter []models.RefPayrollPolicyDetection
	if err = cursor.All(context.TODO(), &payrollPolicyDetectionFilter); err != nil {
		return nil, err
	}
	return payrollPolicyDetectionFilter, nil
}
