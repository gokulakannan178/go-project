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

//SavePlan :""
func (d *Daos) SavePlan(ctx *models.Context, Plan *models.Plan) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONPLAN).InsertOne(ctx.CTX, Plan)
	return err
}

//GetSinglePlan : ""
func (d *Daos) GetSinglePlan(ctx *models.Context, UniqueID string) (*models.RefPlan, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})

	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPLANREGISTRATIONTYPE, "regType", "uniqueId", "ref.regType", "ref.regType")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONAPPLICANT, "creator.id", "userName", "ref.creator", "ref.creator")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "address.stateCode", "code", "ref.address.state", "ref.address.state")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "address.districtCode", "code", "ref.address.district", "ref.address.district")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPLAN).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var plans []models.RefPlan
	var Plan *models.RefPlan
	if err = cursor.All(ctx.CTX, &plans); err != nil {
		return nil, err
	}
	if len(plans) > 0 {
		Plan = &plans[0]
	}
	return Plan, nil
}

//UpdatePlan : ""
func (d *Daos) UpdatePlan(ctx *models.Context, Plan *models.Plan) error {
	selector := bson.M{"uniqueId": Plan.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": Plan, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPLAN).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterPlan : ""
func (d *Daos) FilterPlan(ctx *models.Context, planfilter *models.PlanFilter, pagination *models.Pagination) ([]models.RefPlan, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if planfilter != nil {

		if len(planfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": planfilter.Status}})
		}

		if len(planfilter.Applicant) > 0 {
			query = append(query, bson.M{"creator.id": bson.M{"$in": planfilter.Applicant}})
		}
		if len(planfilter.RegType) > 0 {
			query = append(query, bson.M{"regType": bson.M{"$in": planfilter.RegType}})
		}

		if planfilter.Address != nil {
			if len(planfilter.Address.StateCode) > 0 {
				query = append(query, bson.M{"address.stateCode": bson.M{"$in": planfilter.Address.StateCode}})
			}
			if len(planfilter.Address.DistrictCode) > 0 {
				query = append(query, bson.M{"address.districtCode": bson.M{"$in": planfilter.Address.DistrictCode}})
			}
			if len(planfilter.Address.VillageCode) > 0 {
				query = append(query, bson.M{"address.villageCode": bson.M{"$in": planfilter.Address.VillageCode}})
			}
			if len(planfilter.Address.ZoneCode) > 0 {
				query = append(query, bson.M{"address.zoneCode": bson.M{"$in": planfilter.Address.ZoneCode}})
			}
			if len(planfilter.Address.WardCode) > 0 {
				query = append(query, bson.M{"address.wardCode": bson.M{"$in": planfilter.Address.WardCode}})
			}
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONPLAN).CountDocuments(ctx.CTX, func() bson.M {
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
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPLANREGISTRATIONTYPE, "regType", "uniqueId", "ref.regType", "ref.regType")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONAPPLICANT, "creator.id", "userName", "ref.creator", "ref.creator")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "address.stateCode", "code", "ref.address.state", "ref.address.state")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "address.districtCode", "code", "ref.address.district", "ref.address.district")...)
	//Aggregation
	d.Shared.BsonToJSONPrintTag("Plan query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPLAN).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var plans []models.RefPlan
	if err = cursor.All(context.TODO(), &plans); err != nil {
		return nil, err
	}
	return plans, nil
}

//EnablePlan :""
func (d *Daos) EnablePlan(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PLANSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPLAN).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisablePlan :""
func (d *Daos) DisablePlan(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PLANSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPLAN).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeletePlan :""
func (d *Daos) DeletePlan(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PLANSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPLAN).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//PlanFlowUpdate : ""
func (d *Daos) PlanFlowUpdate(ctx *models.Context, planID string, data interface{}, timeline models.PlanTimeline) error {
	selector := bson.M{"uniqueId": planID}
	updateData := bson.M{"$set": data, "$push": bson.M{"log": timeline}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPLAN).UpdateOne(ctx.CTX, selector, updateData)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}
