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

//SaveEmployeeSalary : ""
func (d *Daos) SaveEmployeeSalary(ctx *models.Context, employeeSalary *models.EmployeeSalary) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEESALARY).InsertOne(ctx.CTX, employeeSalary)
	if err != nil {
		return err
	}
	employeeSalary.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}
func (d *Daos) SaveEmployeeSalaryWithUpsert(ctx *models.Context, attendance *models.EmployeeSalary) error {
	opts := options.Update().SetUpsert(true)
	updateQuery := bson.M{"uniqueId": attendance.UniqueID, "employeeId": attendance.EmployeeId}
	updateData := bson.M{"$set": attendance}
	if _, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEESALARY).UpdateOne(ctx.CTX, updateQuery, updateData, opts); err != nil {
		return errors.New("Error in updating log - " + err.Error())
	}

	return nil
}

//UpdateEmployeeSalary : ""
func (d *Daos) UpdateEmployeeSalary(ctx *models.Context, employeeSalary *models.EmployeeSalary) error {
	selector := bson.M{"uniqueId": employeeSalary.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": employeeSalary}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEESALARY).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

// GetSingleEmployeeSalary : ""
func (d *Daos) GetSingleEmployeeSalary(ctx *models.Context, uniqueID string) (*models.RefEmployeeSalary, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": uniqueID}})
	//LookUp
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisationId", "ref.organisationId")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEESALARY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var EmployeeSalarys []models.RefEmployeeSalary
	var EmployeeSalary *models.RefEmployeeSalary
	if err = cursor.All(ctx.CTX, &EmployeeSalarys); err != nil {
		return nil, err
	}
	if len(EmployeeSalarys) > 0 {
		EmployeeSalary = &EmployeeSalarys[0]
	}
	return EmployeeSalary, err
}

// EnableEmployeeSalary : ""
func (d *Daos) EnableEmployeeSalary(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.EMPLOYEESALARYSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEESALARY).UpdateOne(ctx.CTX, selector, data)
	return err
}

// DisableEmployeeSalary : ""
func (d *Daos) DisableEmployeeSalary(ctx *models.Context, uniqueID string) error {
	selector := bson.M{"uniqueId": uniqueID}
	data := bson.M{"$set": bson.M{"status": constants.EMPLOYEESALARYSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEESALARY).UpdateOne(ctx.CTX, selector, data)
	return err
}

//DeleteEmployeeSalary :""
func (d *Daos) DeleteEmployeeSalary(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.EMPLOYEESALARYSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEESALARY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

// FilterEmployeeSalary : ""
func (d *Daos) FilterEmployeeSalary(ctx *models.Context, employeeSalary *models.FilterEmployeeSalary, pagination *models.Pagination) ([]models.RefEmployeeSalary, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if employeeSalary != nil {
		if len(employeeSalary.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": employeeSalary.Status}})
		}
		if len(employeeSalary.OrganisationId) > 0 {
			query = append(query, bson.M{"organisationId": bson.M{"$in": employeeSalary.OrganisationId}})
		}
		if len(employeeSalary.EmployeeId) > 0 {
			query = append(query, bson.M{"employeeId": bson.M{"$in": employeeSalary.EmployeeId}})
		}

	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if employeeSalary != nil {
		if employeeSalary.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{employeeSalary.SortBy: employeeSalary.SortOrder}})
		}
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEESALARY).CountDocuments(ctx.CTX, func() bson.M {
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
	cursor, err := ctx.DB.Collection(constants.COLLECTIONEMPLOYEESALARY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var employeeSalaryFilter []models.RefEmployeeSalary
	if err = cursor.All(context.TODO(), &employeeSalaryFilter); err != nil {
		return nil, err
	}
	return employeeSalaryFilter, nil
}
