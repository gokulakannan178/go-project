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

//SaveDepartmentType :""
func (d *Daos) SaveDepartmentType(ctx *models.Context, DepartmentType *models.DepartmentType) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONDEPARTMENTTYPE).InsertOne(ctx.CTX, DepartmentType)
	return err
}

//GetSingleDepartmentType : ""
func (d *Daos) GetSingleDepartmentType(ctx *models.Context, UniqueID string) (*models.RefDepartmentType, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDEPARTMENTTYPE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var departmentTypes []models.RefDepartmentType
	var DepartmentType *models.RefDepartmentType
	if err = cursor.All(ctx.CTX, &departmentTypes); err != nil {
		return nil, err
	}
	if len(departmentTypes) > 0 {
		DepartmentType = &departmentTypes[0]
	}
	return DepartmentType, nil
}

//UpdateDepartmentType : ""
func (d *Daos) UpdateDepartmentType(ctx *models.Context, DepartmentType *models.DepartmentType) error {
	selector := bson.M{"uniqueId": DepartmentType.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": DepartmentType, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONDEPARTMENTTYPE).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterDepartmentType : ""
func (d *Daos) FilterDepartmentType(ctx *models.Context, departmentTypefilter *models.DepartmentTypeFilter, pagination *models.Pagination) ([]models.RefDepartmentType, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if departmentTypefilter != nil {

		if len(departmentTypefilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": departmentTypefilter.Status}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONDEPARTMENTTYPE).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("DepartmentType query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDEPARTMENTTYPE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var departmentTypes []models.RefDepartmentType
	if err = cursor.All(context.TODO(), &departmentTypes); err != nil {
		return nil, err
	}
	return departmentTypes, nil
}

//EnableDepartmentType :""
func (d *Daos) EnableDepartmentType(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.DEPARTMENTTYPESTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONDEPARTMENTTYPE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableDepartmentType :""
func (d *Daos) DisableDepartmentType(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.DEPARTMENTTYPESTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONDEPARTMENTTYPE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteDepartmentType :""
func (d *Daos) DeleteDepartmentType(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.DEPARTMENTTYPESTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONDEPARTMENTTYPE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
