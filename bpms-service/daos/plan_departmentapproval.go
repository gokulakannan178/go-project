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

//SavePlanDepartmentApproval :""
func (d *Daos) SavePlanDepartmentApproval(ctx *models.Context, PlanDepartmentApproval *models.PlanDepartmentApproval) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONPLANDEPARTMENTAPPROVAL).InsertOne(ctx.CTX, PlanDepartmentApproval)
	return err
}

//SaveMultiplePlanDepartmentApproval :""
func (d *Daos) SaveMultiplePlanDepartmentApproval(ctx *models.Context, PlanDepartmentApproval []models.PlanDepartmentApproval) error {
	for _, v := range PlanDepartmentApproval {
		opts := options.Update().SetUpsert(true)
		// updateQuery := bson.M{"planId": v.PlanID, "departmentId": v.DepartmentID}
		updateQuery := bson.M{"planId": v.PlanID, "departmentTypeId": v.DepartmentTypeID}
		updateData := bson.M{"$set": v}
		if _, err := ctx.DB.Collection(constants.COLLECTIONPLANDEPARTMENTAPPROVAL).UpdateOne(ctx.CTX, updateQuery, updateData, opts); err != nil {
			return errors.New("Error in updating log - " + err.Error())
		}
	}
	return nil
}

//GetSinglePlanDepartmentApproval : ""
func (d *Daos) GetSinglePlanDepartmentApproval(ctx *models.Context, UniqueID string) (*models.RefPlanDepartmentApproval, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPLANDEPARTMENTAPPROVAL).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var planDepartmentApprovals []models.RefPlanDepartmentApproval
	var PlanDepartmentApproval *models.RefPlanDepartmentApproval
	if err = cursor.All(ctx.CTX, &planDepartmentApprovals); err != nil {
		return nil, err
	}
	if len(planDepartmentApprovals) > 0 {
		PlanDepartmentApproval = &planDepartmentApprovals[0]
	}
	return PlanDepartmentApproval, nil
}

//UpdatePlanDepartmentApproval : ""
func (d *Daos) UpdatePlanDepartmentApproval(ctx *models.Context, PlanDepartmentApproval *models.PlanDepartmentApproval) error {
	selector := bson.M{"uniqueId": PlanDepartmentApproval.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": PlanDepartmentApproval, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPLANDEPARTMENTAPPROVAL).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterPlanDepartmentApproval : ""
func (d *Daos) FilterPlanDepartmentApproval(ctx *models.Context, planDepartmentApprovalfilter *models.PlanDepartmentApprovalFilter, pagination *models.Pagination) ([]models.RefPlanDepartmentApproval, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if planDepartmentApprovalfilter != nil {

		if len(planDepartmentApprovalfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": planDepartmentApprovalfilter.Status}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONPLANDEPARTMENTAPPROVAL).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("PlanDepartmentApproval query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPLANDEPARTMENTAPPROVAL).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var planDepartmentApprovals []models.RefPlanDepartmentApproval
	if err = cursor.All(context.TODO(), &planDepartmentApprovals); err != nil {
		return nil, err
	}
	return planDepartmentApprovals, nil
}

//EnablePlanDepartmentApproval :""
func (d *Daos) EnablePlanDepartmentApproval(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PLANDEPTAPPROVALSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPLANDEPARTMENTAPPROVAL).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisablePlanDepartmentApproval :""
func (d *Daos) DisablePlanDepartmentApproval(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PLANDEPTAPPROVALSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPLANDEPARTMENTAPPROVAL).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeletePlanDepartmentApproval :""
func (d *Daos) DeletePlanDepartmentApproval(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.PLANDEPTAPPROVALSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPLANDEPARTMENTAPPROVAL).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//GetAPlanDeptsApproval : ""
func (d *Daos) GetAPlanDeptsApproval(ctx *models.Context, deptID string) (*models.GetAPlanDeptsApproval, error) {
	query := bson.M{"uniqueId": deptID}
	pipeline := []bson.M{}

	//Pipe 1
	pipeline = append(pipeline, bson.M{"$match": query})
	//Pipe 2
	pipeline = append(pipeline, bson.M{"$lookup": bson.M{
		"from": "planregtypes",
		"as":   "planregtypes",
		"let":  bson.M{"deptId": "$uniqueId"},
		"pipeline": []bson.M{
			bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
				bson.M{"$eq": []string{"$status", "Active"}},
			}}}},
			bson.M{"$lookup": bson.M{
				"from": "plandeptapprovals",
				"as":   "plandeptapproval",
				"let":  bson.M{"deptUniqueId": "$$deptId", "planUniqueId": "$uniqueId"},
				"pipeline": []bson.M{
					bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
						bson.M{"$eq": []string{"$departmentId", "$$deptUniqueId"}},
						bson.M{"$eq": []string{"$planId", "$$planUniqueId"}},
					}}}},
				},
			}},
			bson.M{"$addFields": bson.M{"plandeptapproval": bson.M{"$arrayElemAt": []interface{}{"$plandeptapproval", 0}}}},
		},
	}})

	//Aggregation
	d.Shared.BsonToJSONPrintTag("GetAPlanDeptsApproval query =>", pipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDEPARTMENT).Aggregate(ctx.CTX, pipeline, nil)
	if err != nil {
		return nil, err
	}
	var GetAPlanDeptsApprovals []models.GetAPlanDeptsApproval
	var GetAPlanDeptsApproval *models.GetAPlanDeptsApproval
	if err = cursor.All(context.TODO(), &GetAPlanDeptsApprovals); err != nil {
		return nil, err
	}
	if len(GetAPlanDeptsApprovals) > 0 {
		GetAPlanDeptsApproval = &GetAPlanDeptsApprovals[0]
	}
	return GetAPlanDeptsApproval, nil

}

//GetAPlanDeptsApprovalV2 : ""
func (d *Daos) GetAPlanDeptsApprovalV2(ctx *models.Context, planID string) (*models.GetAPlanDeptsApprovalV2, error) {
	query := bson.M{"uniqueId": planID}
	pipeline := []bson.M{}

	//Pipe 1
	pipeline = append(pipeline, bson.M{"$match": query})
	//Pipe 2
	pipeline = append(pipeline, bson.M{"$lookup": bson.M{
		"from": constants.COLLECTIONDEPARTMENT,
		"as":   "departments",
		"let":  bson.M{"planId": "$uniqueId"},
		"pipeline": []bson.M{
			bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
				bson.M{"$eq": []string{"$status", "Active"}},
			}}}},
			bson.M{"$lookup": bson.M{
				"from": constants.COLLECTIONPLANDEPARTMENTAPPROVAL,
				"as":   "plandeptapproval",
				"let":  bson.M{"deptUniqueId": "$uniqueId", "planUniqueId": "$$planId"},
				"pipeline": []bson.M{
					bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
						bson.M{"$eq": []string{"$departmentId", "$$deptUniqueId"}},
						bson.M{"$eq": []string{"$planId", "$$planUniqueId"}},
					}}}},
				},
			}},
			bson.M{"$addFields": bson.M{"plandeptapproval": bson.M{"$arrayElemAt": []interface{}{"$plandeptapproval", 0}}}},
		},
	}})

	//Aggregation
	d.Shared.BsonToJSONPrintTag("GetAPlanDeptsApproval query =>", pipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPLANREGISTRATIONTYPE).Aggregate(ctx.CTX, pipeline, nil)
	if err != nil {
		return nil, err
	}
	var GetAPlanDeptsApprovals []models.GetAPlanDeptsApprovalV2
	var GetAPlanDeptsApproval *models.GetAPlanDeptsApprovalV2
	if err = cursor.All(context.TODO(), &GetAPlanDeptsApprovals); err != nil {
		return nil, err
	}
	if len(GetAPlanDeptsApprovals) > 0 {
		GetAPlanDeptsApproval = &GetAPlanDeptsApprovals[0]
	}
	return GetAPlanDeptsApproval, nil

}

//GetAPlanDeptsApprovalV3 : ""
func (d *Daos) GetAPlanDeptsApprovalV3(ctx *models.Context, planID string) (*models.GetAPlanDeptsApprovalV3, error) {
	query := bson.M{"uniqueId": planID}
	pipeline := []bson.M{}

	//Pipe 1
	pipeline = append(pipeline, bson.M{"$match": query})
	//Pipe 2
	pipeline = append(pipeline, bson.M{"$lookup": bson.M{
		"from": constants.COLLECTIONDEPARTMENTTYPE,
		"as":   "departments",
		"let":  bson.M{"planId": "$uniqueId"},
		"pipeline": []bson.M{
			bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
				bson.M{"$eq": []string{"$status", "Active"}},
			}}}},
			bson.M{"$lookup": bson.M{
				"from": constants.COLLECTIONPLANDEPARTMENTAPPROVAL,
				"as":   "plandeptapproval",
				"let":  bson.M{"deptTypeId": "$uniqueId", "planUniqueId": "$$planId"},
				"pipeline": []bson.M{
					bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
						bson.M{"$eq": []string{"$departmentTypeId", "$$deptTypeId"}},
						bson.M{"$eq": []string{"$planId", "$$planUniqueId"}},
					}}}},
				},
			}},
			bson.M{"$addFields": bson.M{"plandeptapproval": bson.M{"$arrayElemAt": []interface{}{"$plandeptapproval", 0}}}},
		},
	}})

	//Aggregation
	d.Shared.BsonToJSONPrintTag("GetAPlanDeptsApproval query =>", pipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPLANREGISTRATIONTYPE).Aggregate(ctx.CTX, pipeline, nil)
	if err != nil {
		return nil, err
	}
	var GetAPlanDeptsApprovals []models.GetAPlanDeptsApprovalV3
	var GetAPlanDeptsApproval *models.GetAPlanDeptsApprovalV3
	if err = cursor.All(context.TODO(), &GetAPlanDeptsApprovals); err != nil {
		return nil, err
	}
	if len(GetAPlanDeptsApprovals) > 0 {
		GetAPlanDeptsApproval = &GetAPlanDeptsApprovals[0]
	}
	return GetAPlanDeptsApproval, nil

}

//GetSinglePlanDepartmentApprovalWithPlantype:""
func (d *Daos) GetSinglePlanDepartmentApprovalWithPlantype(ctx *models.Context, UniqueID string) ([]models.RefPlanDepartmentApproval, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"planId": UniqueID, "check": "Yes"}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPLANDEPARTMENTAPPROVAL).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var planDepartmentApprovals []models.RefPlanDepartmentApproval
	//	var PlanDepartmentApproval *models.RefPlanDepartmentApproval
	if err = cursor.All(ctx.CTX, &planDepartmentApprovals); err != nil {
		return nil, err
	}

	return planDepartmentApprovals, nil
}
