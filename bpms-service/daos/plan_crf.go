package daos

import (
	"bpms-service/constants"
	"bpms-service/models"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

//SaveCRF :""
func (d *Daos) SaveCRF(ctx *models.Context, crfs []models.CRF) error {
	data := []interface{}{}
	for _, v := range crfs {
		data = append(data, v)
	}
	_, err := ctx.DB.Collection(constants.COLLECTIONCRF).InsertMany(ctx.CTX, data)
	return err
}
func (d *Daos) SaveSingleCRF(ctx *models.Context, crfs *models.CRF) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONCRF).InsertOne(ctx.CTX, crfs)
	return err
}

//GetSingleCRF : ""
func (d *Daos) GetSingleCRF(ctx *models.Context, UniqueID string) (*models.RefCRF, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDEPARTMENT, "deptId", "uniqueId", "ref.department", "ref.department")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPLAN, "planId", "uniqueId", "ref.plan", "ref.plan")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPLANREGISTRATIONTYPE, "planRegTypeId", "uniqueId", "ref.planRegType", "ref.planRegType")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONCRF).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var crfs []models.RefCRF
	var CRF *models.RefCRF
	if err = cursor.All(ctx.CTX, &crfs); err != nil {
		return nil, err
	}
	if len(crfs) > 0 {
		CRF = &crfs[0]
	}
	return CRF, nil
}

//FilterCRF : ""
func (d *Daos) FilterCRF(ctx *models.Context, crffilter *models.CRFFilter, pagination *models.Pagination) ([]models.RefCRF, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if crffilter != nil {

		if len(crffilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": crffilter.Status}})
		}
		if len(crffilter.Dept) > 0 {
			query = append(query, bson.M{"deptId": bson.M{"$in": crffilter.Dept}})
		}
		if len(crffilter.PlanRegType) > 0 {
			query = append(query, bson.M{"planRegTypeId": bson.M{"$in": crffilter.PlanRegType}})
		}
		if len(crffilter.Plan) > 0 {
			query = append(query, bson.M{"planId": bson.M{"$in": crffilter.Plan}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONCRF).CountDocuments(ctx.CTX, func() bson.M {
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
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDEPARTMENT, "deptId", "uniqueId", "ref.department", "ref.department")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPLAN, "planId", "uniqueId", "ref.plan", "ref.plan")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPLANREGISTRATIONTYPE, "planRegTypeId", "uniqueId", "ref.planRegType", "ref.planRegType")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONULB, "ref.department.organisationId", "uniqueId", "ref.ulb", "ref.ulb")...)
	//ULB Address Lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "ref.department.address.stateCode", "code", "ref.ulbAddress.state", "ref.ulbAddress.state")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "ref.department.address.districtCode", "code", "ref.ulbAddress.district", "ref.ulbAddress.district")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("CRF query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONCRF).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var crfs []models.RefCRF
	if err = cursor.All(context.TODO(), &crfs); err != nil {
		return nil, err
	}
	return crfs, nil
}

//CRF Inspection APIS

//SaveCRFInspection :""
func (d *Daos) SaveCRFInspection(ctx *models.Context, crfsInspections []models.CRFInspection) error {
	data := []interface{}{}
	for _, v := range crfsInspections {
		data = append(data, v)
	}
	_, err := ctx.DB.Collection(constants.COLLECTIONCRFINSPECTION).InsertMany(ctx.CTX, data)
	return err
}

//GetCRFInspectionOfPlan : ""
func (d *Daos) GetCRFInspectionOfPlan(ctx *models.Context, planID, deptID string) ([]models.RefCRFInspection, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"planId": planID, "deptId": deptID}})
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDEPTCHECKLIST, "checkListId", "uniqueId", "ref.checklist", "ref.checklist")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("CRF query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONCRFINSPECTION).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var crfis []models.RefCRFInspection
	if err = cursor.All(ctx.CTX, &crfis); err != nil {
		return nil, err
	}
	return crfis, nil
}

//SubmitCRFInspection : ""
func (d *Daos) SubmitCRFInspection(ctx *models.Context, crfsInspection *models.CRFInspection) error {
	selector := bson.M{"uniqueId": crfsInspection.UniqueID}
	data := bson.M{"$set": crfsInspection}
	_, err := ctx.DB.Collection(constants.COLLECTIONCRFINSPECTION).UpdateOne(ctx.CTX, selector, data, nil)
	return err
}

//PlanCRFFlowUpdate : ""
func (d *Daos) PlanCRFFlowUpdate(ctx *models.Context, crfID string, data interface{}, timeline models.PlanTimeline) error {
	selector := bson.M{"uniqueId": crfID}
	d.Shared.BsonToJSONPrintTag("CRF query =>", selector)
	updateData := bson.M{"$set": data, "$push": bson.M{"log": timeline}}
	d.Shared.BsonToJSONPrintTag("CRF data =>", updateData)
	_, err := ctx.DB.Collection(constants.COLLECTIONCRF).UpdateOne(ctx.CTX, selector, updateData)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}
