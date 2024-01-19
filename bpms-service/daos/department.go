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

//SaveDepartment :""
func (d *Daos) SaveDepartment(ctx *models.Context, Department *models.Department) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONDEPARTMENT).InsertOne(ctx.CTX, Department)
	return err
}

//GetSingleDepartment : ""
func (d *Daos) GetSingleDepartment(ctx *models.Context, UniqueID string) (*models.RefDepartment, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})
	//Lookups
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "address.stateCode", "code", "ref.address.state", "ref.address.state")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "address.districtCode", "code", "ref.address.district", "ref.address.district")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONULB, "organisationId", "uniqueId", "ref.organisation", "ref.organisation")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDEPARTMENTTYPE, "departmentTypeId", "uniqueId", "ref.departmentType", "ref.departmentType")...)
	//Aggregation
	d.Shared.BsonToJSONPrintTag("Department query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDEPARTMENT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}

	var departments []models.RefDepartment
	var Department *models.RefDepartment
	if err = cursor.All(ctx.CTX, &departments); err != nil {
		return nil, err
	}
	if len(departments) > 0 {
		Department = &departments[0]
	}
	return Department, nil
}

//UpdateDepartment : ""
func (d *Daos) UpdateDepartment(ctx *models.Context, Department *models.Department) error {
	selector := bson.M{"uniqueId": Department.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": Department, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONDEPARTMENT).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterDepartment : ""
func (d *Daos) FilterDepartment(ctx *models.Context, departmentfilter *models.DepartmentFilter, pagination *models.Pagination) ([]models.RefDepartment, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if departmentfilter != nil {

		if len(departmentfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": departmentfilter.Status}})
		}
		if len(departmentfilter.Organisation) > 0 {
			query = append(query, bson.M{"organisationId": bson.M{"$in": departmentfilter.Organisation}})
		}
		if len(departmentfilter.DistrictCode) > 0 {
			query = append(query, bson.M{"address.districtCode": bson.M{"$in": departmentfilter.DistrictCode}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONDEPARTMENT).CountDocuments(ctx.CTX, func() bson.M {
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
	//Lookups
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "address.stateCode", "code", "ref.address.state", "ref.address.state")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "address.districtCode", "code", "ref.address.district", "ref.address.district")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONULB, "organisationId", "uniqueId", "ref.organisation", "ref.organisation")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDEPARTMENTTYPE, "departmentTypeId", "uniqueId", "ref.departmentType", "ref.departmentType")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("Department query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDEPARTMENT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var departments []models.RefDepartment
	if err = cursor.All(context.TODO(), &departments); err != nil {
		return nil, err
	}
	return departments, nil
}

//EnableDepartment :""
func (d *Daos) EnableDepartment(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.EDUCATIONTYPESTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONDEPARTMENT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableDepartment :""
func (d *Daos) DisableDepartment(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.EDUCATIONTYPESTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONDEPARTMENT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteDepartment :""
func (d *Daos) DeleteDepartment(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.EDUCATIONTYPESTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONDEPARTMENT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//GetSingleDepartment : ""
func (d *Daos) GetSingleDepartmentWithDistrictAndDepartmentType(ctx *models.Context, departmentType string, distric string) (*models.RefDepartment, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"address.districtCode": distric, "departmentTypeId": departmentType}})
	//Lookups
	//Aggregation
	d.Shared.BsonToJSONPrintTag("Department query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDEPARTMENT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}

	var departments []models.RefDepartment
	var Department *models.RefDepartment
	if err = cursor.All(ctx.CTX, &departments); err != nil {
		return nil, err
	}
	if len(departments) > 0 {
		Department = &departments[0]
	}
	return Department, nil
}
