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

//SavePayrollPolicyEarning : ""
func (d *Daos) SavePayrollPolicyEarning(ctx *models.Context, payrollPolicyEarning *models.PayrollPolicyEarning) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONPAYROLLPOLICYEARNING).InsertOne(ctx.CTX, payrollPolicyEarning)
	if err != nil {
		return err
	}
	payrollPolicyEarning.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//UpdatePayrollPolicyEarning : ""
func (d *Daos) UpdatePayrollPolicyEarning(ctx *models.Context, payrollPolicyEarning *models.PayrollPolicyEarning) error {
	selector := bson.M{"uniqueId": payrollPolicyEarning.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": payrollPolicyEarning}
	_, err := ctx.DB.Collection(constants.COLLECTIONPAYROLLPOLICYEARNING).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}
func (d *Daos) UpdatePayrollPolicyEarningWithEaringId(ctx *models.Context, payrollPolicyEarning *models.PayrollPolicyEarning) error {
	opts := options.Update().SetUpsert(true)
	selector := bson.M{"payRollPolicyId": payrollPolicyEarning.PayRollPolicyId, "earningMasterId": payrollPolicyEarning.EarningMasterId}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": payrollPolicyEarning}
	_, err := ctx.DB.Collection(constants.COLLECTIONPAYROLLPOLICYEARNING).UpdateOne(ctx.CTX, selector, updateInterface, opts)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// GetSinglePayrollPolicyEarning : ""
func (d *Daos) GetSinglePayrollPolicyEarning(ctx *models.Context, uniqueID string) (*models.RefPayrollPolicyEarning, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	//LookUp
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPAYROLLPOLICYEARNING).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var PayrollPolicyEarnings []models.RefPayrollPolicyEarning
	var PayrollPolicyEarning *models.RefPayrollPolicyEarning
	if err = cursor.All(ctx.CTX, &PayrollPolicyEarnings); err != nil {
		return nil, err
	}
	if len(PayrollPolicyEarnings) > 0 {
		PayrollPolicyEarning = &PayrollPolicyEarnings[0]
	}
	return PayrollPolicyEarning, err
}

// EnablePayrollPolicyEarning : ""
func (d *Daos) EnablePayrollPolicyEarning(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.PAYROLLPOLICYEARNINGSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPAYROLLPOLICYEARNING).UpdateOne(ctx.CTX, selector, data)
	return err
}

// DisablePayrollPolicyEarning : ""
func (d *Daos) DisablePayrollPolicyEarning(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.PAYROLLPOLICYEARNINGSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPAYROLLPOLICYEARNING).UpdateOne(ctx.CTX, selector, data)
	return err
}

//DeletePayrollPolicyEarning :""
func (d *Daos) DeletePayrollPolicyEarning(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PAYROLLPOLICYEARNINGSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPAYROLLPOLICYEARNING).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterPayrollPolicyEarning : ""
func (d *Daos) FilterPayrollPolicyEarning(ctx *models.Context, payrollPolicyEarning *models.FilterPayrollPolicyEarning, pagination *models.Pagination) ([]models.RefPayrollPolicyEarning, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if payrollPolicyEarning != nil {
		if len(payrollPolicyEarning.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": payrollPolicyEarning.Status}})
		}
		if len(payrollPolicyEarning.OrganisationID) > 0 {
			query = append(query, bson.M{"organisationId": bson.M{"$in": payrollPolicyEarning.OrganisationID}})
		}
		//Regex
		if payrollPolicyEarning.Regex.Name != "" {
			query = append(query, bson.M{"title": primitive.Regex{Pattern: payrollPolicyEarning.Regex.Name, Options: "xi"}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONPAYROLLPOLICYEARNING).CountDocuments(ctx.CTX, func() bson.M {
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
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPAYROLLPOLICYEARNING).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var payrollPolicyEarningFilter []models.RefPayrollPolicyEarning
	if err = cursor.All(context.TODO(), &payrollPolicyEarningFilter); err != nil {
		return nil, err
	}
	return payrollPolicyEarningFilter, nil
}
