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

//SaveDeptChecklist :""
func (d *Daos) SaveDeptChecklist(ctx *models.Context, DeptChecklist *models.DeptChecklist) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONDEPTCHECKLIST).InsertOne(ctx.CTX, DeptChecklist)
	return err
}

//GetSingleDeptChecklist : ""
func (d *Daos) GetSingleDeptChecklist(ctx *models.Context, UniqueID string) (*models.RefDeptChecklist, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDEPARTMENT, "deptId", "uniqueId", "ref.department", "ref.department")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONULB, "ref.department.organisationId", "uniqueId", "ref.ulb", "ref.ulb")...)
	d.Shared.BsonToJSONPrintTag("DeptChecklist query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDEPTCHECKLIST).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var deptChecklists []models.RefDeptChecklist
	var DeptChecklist *models.RefDeptChecklist
	if err = cursor.All(ctx.CTX, &deptChecklists); err != nil {
		return nil, err
	}
	if len(deptChecklists) > 0 {
		DeptChecklist = &deptChecklists[0]
	}
	return DeptChecklist, nil
}

//UpdateDeptChecklist : ""
func (d *Daos) UpdateDeptChecklist(ctx *models.Context, DeptChecklist *models.DeptChecklist) error {
	selector := bson.M{"uniqueId": DeptChecklist.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": DeptChecklist, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONDEPTCHECKLIST).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterDeptChecklist : ""
func (d *Daos) FilterDeptChecklist(ctx *models.Context, deptChecklistfilter *models.DeptChecklistFilter, pagination *models.Pagination) ([]models.RefDeptChecklist, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if deptChecklistfilter != nil {

		if len(deptChecklistfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": deptChecklistfilter.Status}})
		}
		if len(deptChecklistfilter.Dept) > 0 {
			query = append(query, bson.M{"dept": bson.M{"$in": deptChecklistfilter.Dept}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONDEPTCHECKLIST).CountDocuments(ctx.CTX, func() bson.M {
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
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONULB, "ref.department.organisationId", "uniqueId", "ref.ulb", "ref.ulb")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("DeptChecklist query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDEPTCHECKLIST).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var deptChecklists []models.RefDeptChecklist
	if err = cursor.All(context.TODO(), &deptChecklists); err != nil {
		return nil, err
	}
	return deptChecklists, nil
}

//EnableDeptChecklist :""
func (d *Daos) EnableDeptChecklist(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.DEPTCHECKLISTSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONDEPTCHECKLIST).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableDeptChecklist :""
func (d *Daos) DisableDeptChecklist(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.DEPTCHECKLISTSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONDEPTCHECKLIST).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteDeptChecklist :""
func (d *Daos) DeleteDeptChecklist(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.DEPTCHECKLISTSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONDEPTCHECKLIST).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//GetChecklistsOfDept : ""
func (d *Daos) GetChecklistsOfDept(ctx *models.Context, deptID string) ([]models.RefDeptChecklist, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}

	query = append(query, bson.M{"deptId": bson.M{"$in": []string{deptID}}})

	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	//Aggregation
	d.Shared.BsonToJSONPrintTag("DeptChecklist query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDEPTCHECKLIST).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var deptChecklists []models.RefDeptChecklist
	if err = cursor.All(context.TODO(), &deptChecklists); err != nil {
		return nil, err
	}
	return deptChecklists, nil
}
