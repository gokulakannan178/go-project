package daos

import (
	"bpms-service/constants"
	"bpms-service/models"
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

//SavePlanRegistrationType :""
func (d *Daos) SavePlanRegistrationType(ctx *models.Context, PlanRegistrationType *models.PlanRegistrationType) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONPLANREGISTRATIONTYPE).InsertOne(ctx.CTX, PlanRegistrationType)
	return err
}

//GetSinglePlanRegistrationType : ""
func (d *Daos) GetSinglePlanRegistrationType(ctx *models.Context, UniqueID string) (*models.RefPlanRegistrationType, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPLANREGISTRATIONTYPE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var planRegistrationTypes []models.RefPlanRegistrationType
	var PlanRegistrationType *models.RefPlanRegistrationType
	if err = cursor.All(ctx.CTX, &planRegistrationTypes); err != nil {
		return nil, err
	}
	if len(planRegistrationTypes) > 0 {
		PlanRegistrationType = &planRegistrationTypes[0]
	}
	return PlanRegistrationType, nil
}

//UpdatePlanRegistrationType : ""
func (d *Daos) UpdatePlanRegistrationType(ctx *models.Context, PlanRegistrationType *models.PlanRegistrationType) error {
	selector := bson.M{"uniqueId": PlanRegistrationType.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": PlanRegistrationType, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPLANREGISTRATIONTYPE).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterPlanRegistrationType : ""
func (d *Daos) FilterPlanRegistrationType(ctx *models.Context, planRegistrationTypefilter *models.PlanRegistrationTypeFilter, pagination *models.Pagination) ([]models.RefPlanRegistrationType, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if planRegistrationTypefilter != nil {

		if len(planRegistrationTypefilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": planRegistrationTypefilter.Status}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONPLANREGISTRATIONTYPE).CountDocuments(ctx.CTX, func() bson.M {
			if query != nil {
				if len(query) > 0 {
					return bson.M{"$and": query}
				}
			}
			return bson.M{}
		}())
		if err != nil {
			log.Println("Error in geting pagination")
		}
		fmt.Println("count", totalCount)
		pagination.Count = int(totalCount)
		d.Shared.PaginationData(pagination)
	}

	//Aggregation
	d.Shared.BsonToJSONPrintTag("PlanRegistrationType query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPLANREGISTRATIONTYPE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var planRegistrationTypes []models.RefPlanRegistrationType
	if err = cursor.All(context.TODO(), &planRegistrationTypes); err != nil {
		return nil, err
	}
	return planRegistrationTypes, nil
}

//EnablePlanRegistrationType :""
func (d *Daos) EnablePlanRegistrationType(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PLANREGISTRATIONTYPESTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPLANREGISTRATIONTYPE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisablePlanRegistrationType :""
func (d *Daos) DisablePlanRegistrationType(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PLANREGISTRATIONTYPESTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPLANREGISTRATIONTYPE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeletePlanRegistrationType :""
func (d *Daos) DeletePlanRegistrationType(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PLANREGISTRATIONTYPESTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPLANREGISTRATIONTYPE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
