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

//SavePlanReqDocument :""
func (d *Daos) SavePlanReqDocument(ctx *models.Context, PlanReqDocument *models.PlanReqDocument) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONPLANREQDOCUMENT).InsertOne(ctx.CTX, PlanReqDocument)
	return err
}

//GetSinglePlanReqDocument : ""
func (d *Daos) GetSinglePlanReqDocument(ctx *models.Context, UniqueID string) (*models.RefPlanReqDocument, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONULB, "orgId", "uniqueId", "ref.tempULB", "ref.tempULB")...)
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDEPARTMENT, "orgId", "uniqueId", "ref.dept", "ref.tempDept")...)
	// mainPipeline = append(mainPipeline,
	// 	bson.M{"$addFields": bson.M{"ref.ulb": bson.M{
	// 		"$cond": bson.M{"if": bson.M{"$eq": []string{"$orgType", "ULB"}}, "then": "$ref.tempULB", "else": nil},
	// 	},

	// 		"ref.department": bson.M{
	// 			"$cond": bson.M{"if": bson.M{"$eq": []string{"$orgType", "Department"}}, "then": "$ref.tempDept", "else": nil},
	// 		},
	// 	}})
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONULB, "departmentId", "uniqueId", "ref.department", "ref.department")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("PlanReqDocument query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPLANREQDOCUMENT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var planReqDocuments []models.RefPlanReqDocument
	var PlanReqDocument *models.RefPlanReqDocument
	if err = cursor.All(ctx.CTX, &planReqDocuments); err != nil {
		return nil, err
	}
	if len(planReqDocuments) > 0 {
		PlanReqDocument = &planReqDocuments[0]
	}
	return PlanReqDocument, nil
}

//UpdatePlanReqDocument : ""
func (d *Daos) UpdatePlanReqDocument(ctx *models.Context, PlanReqDocument *models.PlanReqDocument) error {
	selector := bson.M{"uniqueId": PlanReqDocument.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": PlanReqDocument, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPLANREQDOCUMENT).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterPlanReqDocument : ""
func (d *Daos) FilterPlanReqDocument(ctx *models.Context, planReqDocumentfilter *models.PlanReqDocumentFilter, pagination *models.Pagination) ([]models.RefPlanReqDocument, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if planReqDocumentfilter != nil {

		if len(planReqDocumentfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": planReqDocumentfilter.Status}})
		}
		if len(planReqDocumentfilter.Org) > 0 {
			query = append(query, bson.M{"orgId": bson.M{"$in": planReqDocumentfilter.Org}})
		}
		if len(planReqDocumentfilter.OrgType) > 0 {
			query = append(query, bson.M{"orgType": bson.M{"$in": planReqDocumentfilter.OrgType}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONPLANREQDOCUMENT).CountDocuments(ctx.CTX, func() bson.M {
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
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONULB, "orgId", "uniqueId", "ref.tempULB", "ref.tempULB")...)
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDEPARTMENT, "orgId", "uniqueId", "ref.dept", "ref.tempDept")...)
	// mainPipeline = append(mainPipeline,
	// 	bson.M{"$addFields": bson.M{"ref.ulb": bson.M{
	// 		"$cond": bson.M{"if": bson.M{"$eq": []string{"$orgType", "ULB"}}, "then": "$ref.tempULB", "else": nil},
	// 	},

	// 		"ref.department": bson.M{
	// 			"$cond": bson.M{"if": bson.M{"$eq": []string{"$orgType", "Department"}}, "then": "$ref.tempDept", "else": nil},
	// 		},
	// 	}})
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONULB, "departmentId", "uniqueId", "ref.department", "ref.department")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDEPARTMENT, "departmentId", "uniqueId", "ref.department", "ref.department")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("PlanReqDocument query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPLANREQDOCUMENT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var planReqDocuments []models.RefPlanReqDocument
	if err = cursor.All(context.TODO(), &planReqDocuments); err != nil {
		return nil, err
	}
	return planReqDocuments, nil
}

//EnablePlanReqDocument :""
func (d *Daos) EnablePlanReqDocument(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PLANREQDOCUMENTSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPLANREQDOCUMENT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisablePlanReqDocument :""
func (d *Daos) DisablePlanReqDocument(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PLANREQDOCUMENTSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPLANREQDOCUMENT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeletePlanReqDocument :""
func (d *Daos) DeletePlanReqDocument(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PLANREQDOCUMENTSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPLANREQDOCUMENT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
