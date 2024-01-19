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
	"go.mongodb.org/mongo-driver/mongo/options"
)

//SavePlanDocument :""
func (d *Daos) SavePlanDocument(ctx *models.Context, PlanDocument *models.PlanDocument) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONPLANDOCUMENT).InsertOne(ctx.CTX, PlanDocument)
	return err
}

//UpsertPlanDocument :""
func (d *Daos) UpsertPlanDocument(ctx *models.Context, PlanDocument *models.PlanDocument) error {
	selector := bson.M{"planId": PlanDocument.PlanID, "orgId": PlanDocument.OrgID, "orgType": PlanDocument.OrgType, "docId": PlanDocument.DocID}
	data := bson.M{"$set": PlanDocument}
	opts := options.FindOneAndUpdate().SetUpsert(true)
	d.Shared.BsonToJSONPrintTag("upsert query =>", selector)
	ctx.DB.Collection(constants.COLLECTIONPLANDOCUMENT).FindOneAndUpdate(ctx.CTX, selector, data, opts)
	return nil
}

//GetSinglePlanDocument : ""
func (d *Daos) GetSinglePlanDocument(ctx *models.Context, UniqueID string) (*models.RefPlanDocument, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPLANREQDOCUMENT, "docId", "uniqueId", "ref.doc", "ref.doc")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPLAN, "planId", "uniqueId", "ref.plan", "ref.plan")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONULB, "orgId", "uniqueId", "ref.tempULB", "ref.tempULB")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDEPARTMENT, "orgId", "uniqueId", "ref.dept", "ref.tempDept")...)
	mainPipeline = append(mainPipeline,
		bson.M{"$addFields": bson.M{"ref.ulb": bson.M{
			"$cond": bson.M{"if": bson.M{"$eq": []string{"$orgType", "ULB"}}, "then": "$ref.tempULB", "else": nil},
		},

			"ref.department": bson.M{
				"$cond": bson.M{"if": bson.M{"$eq": []string{"$orgType", "Department"}}, "then": "$ref.tempDept", "else": nil},
			},
		}})
	//Aggregation
	d.Shared.BsonToJSONPrintTag("PlanDocument query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPLANDOCUMENT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var planDocuments []models.RefPlanDocument
	var PlanDocument *models.RefPlanDocument
	if err = cursor.All(ctx.CTX, &planDocuments); err != nil {
		return nil, err
	}
	if len(planDocuments) > 0 {
		PlanDocument = &planDocuments[0]
	}
	return PlanDocument, nil
}

//UpdatePlanDocument : ""
func (d *Daos) UpdatePlanDocument(ctx *models.Context, PlanDocument *models.PlanDocument) error {
	selector := bson.M{"uniqueId": PlanDocument.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": PlanDocument, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPLANDOCUMENT).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterPlanDocument : ""
func (d *Daos) FilterPlanDocument(ctx *models.Context, planDocumentfilter *models.PlanDocumentFilter, pagination *models.Pagination) ([]models.RefPlanDocument, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if planDocumentfilter != nil {

		if len(planDocumentfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": planDocumentfilter.Status}})
		}
		if len(planDocumentfilter.Org) > 0 {
			query = append(query, bson.M{"orgId": bson.M{"$in": planDocumentfilter.Org}})
		}
		if len(planDocumentfilter.OrgType) > 0 {
			query = append(query, bson.M{"orgType": bson.M{"$in": planDocumentfilter.OrgType}})
		}
		if len(planDocumentfilter.Plan) > 0 {
			query = append(query, bson.M{"planId": bson.M{"$in": planDocumentfilter.Plan}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONPLANDOCUMENT).CountDocuments(ctx.CTX, func() bson.M {
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
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPLAN, "planId", "uniqueId", "ref.plan", "ref.plan")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPLANREQDOCUMENT, "docId", "uniqueId", "ref.doc", "ref.doc")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONULB, "orgId", "uniqueId", "ref.tempULB", "ref.tempULB")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDEPARTMENT, "orgId", "uniqueId", "ref.dept", "ref.tempDept")...)
	mainPipeline = append(mainPipeline,
		bson.M{"$addFields": bson.M{"ref.ulb": bson.M{
			"$cond": bson.M{"if": bson.M{"$eq": []string{"$orgType", "ULB"}}, "then": "$ref.tempULB", "else": nil},
		},

			"ref.department": bson.M{
				"$cond": bson.M{"if": bson.M{"$eq": []string{"$orgType", "Department"}}, "then": "$ref.tempDept", "else": nil},
			},
		}})
	//Aggregation
	d.Shared.BsonToJSONPrintTag("PlanDocument query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPLANDOCUMENT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var planDocuments []models.RefPlanDocument
	if err = cursor.All(context.TODO(), &planDocuments); err != nil {
		return nil, err
	}
	return planDocuments, nil
}

//EnablePlanDocument :""
func (d *Daos) EnablePlanDocument(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PLANDOCUMENTSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPLANDOCUMENT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisablePlanDocument :""
func (d *Daos) DisablePlanDocument(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PLANDOCUMENTSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPLANDOCUMENT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeletePlanDocument :""
func (d *Daos) DeletePlanDocument(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PLANDOCUMENTSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPLANDOCUMENT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//GetPendingDocuments : ""
func (d *Daos) GetPendingDocuments(ctx *models.Context, planDocumentfilter *models.GetPendingPlanDocumentFilter, pagination *models.Pagination) ([]models.RefPlanDocument, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if planDocumentfilter != nil {

		if len(planDocumentfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": planDocumentfilter.Status}})
		}
		if len(planDocumentfilter.Org) > 0 {
			query = append(query, bson.M{"orgId": planDocumentfilter.Org})
		}
		if len(planDocumentfilter.OrgType) > 0 {
			query = append(query, bson.M{"orgType": planDocumentfilter.OrgType})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	mainPipeline = append(mainPipeline, bson.M{
		"$lookup": bson.M{
			"from": constants.COLLECTIONPLANDOCUMENT,
			"as":   "ref.doc",
			"let":  bson.M{"orgId": "$orgId", "orgType": "$orgType", "docId": "$uniqueId", "planId": planDocumentfilter.Plan},
			"pipeline": []bson.M{
				bson.M{"$match": bson.M{
					"$expr": bson.M{
						"$and": []bson.M{
							bson.M{"$eq": []string{"$orgId", "$$orgId"}},
							bson.M{"$eq": []string{"$orgType", "$$orgType"}},
							bson.M{"$eq": []string{"$planId", "$$planId"}},
							bson.M{"$eq": []string{"$docId", "$$docId"}},
						}},
				},
				},
			},
		},
	})
	mainPipeline = append(mainPipeline, bson.M{
		"$addFields": bson.M{"ref.doc": bson.M{"$arrayElemAt": []interface{}{"$ref.doc", 0}}},
	})

	mainPipeline = append(mainPipeline, bson.M{
		"$match": bson.M{"ref.doc.status": bson.M{"$nin": []string{constants.PLANDOCUMENTSTATUSACTIVE}}},
	})

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONPLANDOCUMENT).CountDocuments(ctx.CTX, func() bson.M {
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
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONULB, "orgId", "uniqueId", "ref.tempULB", "ref.tempULB")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDEPARTMENT, "orgId", "uniqueId", "ref.dept", "ref.tempDept")...)
	mainPipeline = append(mainPipeline,
		bson.M{"$addFields": bson.M{"ref.ulb": bson.M{
			"$cond": bson.M{"if": bson.M{"$eq": []string{"$orgType", "ULB"}}, "then": "$ref.tempULB", "else": nil},
		},

			"ref.department": bson.M{
				"$cond": bson.M{"if": bson.M{"$eq": []string{"$orgType", "Department"}}, "then": "$ref.tempDept", "else": nil},
			},
		}})
	//Aggregation
	d.Shared.BsonToJSONPrintTag("PlanDocument query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPLANREQDOCUMENT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var planDocuments []models.RefPlanDocument
	if err = cursor.All(context.TODO(), &planDocuments); err != nil {
		return nil, err
	}
	return planDocuments, nil
}
